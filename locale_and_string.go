package structs

type LocaleAndStringCollection []LocaleAndString
type LocaleAndStringStringMap map[string]LocaleAndString
type LocaleAndStringIntMap map[int]LocaleAndString
type LocaleAndStringUIntMap map[uint]LocaleAndString

// An object representing a language specific string.
type LocaleAndString struct {
	Language string `json:"language"`
	Value    string `json:"value"`
}
