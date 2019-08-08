package mixerstructs

type TypeTitleCollection []TypeTitle
type TypeTitleStringMap map[string]TypeTitle
type TypeTitleIntMap map[int]TypeTitle
type TypeTitleUIntMap map[uint]TypeTitle

// Describes a title associated with a type.
type TypeTitle struct {
	AumId       string `json:"aumId"`
	KglId       string `json:"kglId"`
	ProcessName string `json:"processName"`
	ProductId   string `json:"productId"`
	TitleId     int    `json:"titleId"`
}
