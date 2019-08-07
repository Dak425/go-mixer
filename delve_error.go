package structs

type DelveErrorCollection []DelveError
type DelveErrorStringMap map[string]DelveError
type DelveErrorIntMap map[int]DelveError
type DelveErrorUIntMap map[uint]DelveError

// Returned when we tried to get some contents for this row, but encountered an error doing so.
type DelveError struct {
	OriginalType string `json:"originalType"`
	Type         string `json:"type"`
}
