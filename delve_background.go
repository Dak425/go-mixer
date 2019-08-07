package structs

type DelveBackgroundCollection []DelveBackground
type DelveBackgroundStringMap map[string]DelveBackground
type DelveBackgroundIntMap map[int]DelveBackground
type DelveBackgroundUIntMap map[uint]DelveBackground

// Dictates what background to display on the homepage.
type DelveBackground struct {
	FullWidthSrc string `json:"fullWidthSrc"`
	LeftSrc      string `json:"leftSrc"`
	RightSrc     string `json:"rightSrc"`
}
