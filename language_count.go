package structs

type LanguageCountCollection []LanguageCount
type LanguageCountStringMap map[string]LanguageCount
type LanguageCountIntMap map[int]LanguageCount
type LanguageCountUIntMap map[uint]LanguageCount

// An object describing a count for a language.
type LanguageCount struct {
	Count uint   `json:"count"`
	Id    string `json:"id"`
}
