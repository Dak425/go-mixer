package mixer

type LanguageCollection []Language
type LanguageStringMap map[string]Language
type LanguageIntMap map[int]Language
type LanguageUIntMap map[uint]Language

// An object representing a language.
type Language struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
