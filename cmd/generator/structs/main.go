package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
	"text/template"

	"github.com/gocolly/colly"
)

type ScrapedTypeMap map[string]ScrapedType
type ScrapedTypeCollection []ScrapedType

type ScrapedType struct {
	Name        string            `json:"name"`
	Extends     string            `json:"extends"`
	Description string            `json:"description"`
	Attributes  map[string]string `json:"attributes"`
}

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")
var matchAlpha = regexp.MustCompile("([^a-zA-Z]+)")
var matchColon = regexp.MustCompile(":")

func toSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

func sanitizeTypeText(str string) string {
	return matchAlpha.ReplaceAllString(str, "")
}

func main() {
	url := "https://dev.mixer.com/rest/index.html"

	panelSelector := "body > div.site-content > div > div > div.col-md-9 > div"
	extendsSelector := "div.panel-heading > h3 > small > a"
	descriptionSelector := "div.panel-body > p:nth-child(2)"
	propertiesSelector := "div.panel-body > ul > li"
	propKeySelector := "h5 > span"
	propValueSelector := "h5 > small"
	propValueRefSelector := "h5 > small > a"

	apiCollection := ScrapedTypeCollection{}

	ignore := map[string]bool{
		"uint":    true,
		"IsoDate": true,
		"UUID":    true,
		"Role":    true,
	}

	typeMap := map[string]string{
		"IsoDate":     "string",
		"number":      "int",
		"integer":     "int",
		"boolean":     "bool",
		"object":      "struct{}",
		"UUID":        "string",
		"Role":        "string",
		"arraystring": "[]string",
	}

	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting: ", r.URL)
	})

	c.OnHTML(panelSelector, func(e *colly.HTMLElement) {
		if e.Index >= 29 {
			// Get the name of each type from a specific link tag
			st := ScrapedType{
				Attributes: make(map[string]string),
			}
			st.Name = e.ChildAttr("a", "id")

			if _, ok := ignore[st.Name]; ok {
				return
			}
			// Check if the type extends another type
			st.Extends = e.ChildText(extendsSelector)

			// Type description
			st.Description = e.ChildText(descriptionSelector)

			// Type properties
			e.ForEach(propertiesSelector, func(index int, li *colly.HTMLElement) {
				v := li.ChildText(propValueRefSelector)
				if v == "" {
					v = li.ChildText(propValueSelector)
				}
				st.Attributes[li.ChildText(propKeySelector)] = v
			})

			// Add scraped type info to collection
			apiCollection = append(apiCollection, st)
		}

	})

	c.OnScraped(func(r *colly.Response) {
		log.Println("Complete")
		log.Printf("Compiled Map: %+v \n", apiCollection)
		data, err := json.MarshalIndent(apiCollection, "", "	")
		if err != nil {
			log.Fatalf("Error when marshalling type collection to JSON: %v", err)
		}
		// Write scraped types to file for caching
		ioutil.WriteFile("gen/typeMap.json", data, 0644)

		// Define func map used for template processing to get desired formatting
		funcMap := template.FuncMap{
			"ToTitle": strings.Title,
			"ReplaceType": func(text string) string {
				if value, found := typeMap[text]; found {
					return value
				}
				return text
			},
			"Sanitize": sanitizeTypeText,
		}

		t := template.Must(template.New("struct.tmpl").Funcs(funcMap).ParseFiles("templates/struct.tmpl"))
		for _, scraped := range apiCollection {
			file, err := os.OpenFile("gen/"+toSnakeCase(scraped.Name)+".go", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
			if err != nil {
				log.Fatalf("Unable to open file to for template processing: %v", err)
			}
			err = t.Execute(file, scraped)
			if err != nil {
				log.Fatalf("Error when processing scraped type to template: %v", err)
			}
			file.Close()
		}
	})

	c.Visit(url)
}
