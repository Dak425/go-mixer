package mixer

type ProgressionEventListCollection []ProgressionEventList
type ProgressionEventListStringMap map[string]ProgressionEventList
type ProgressionEventListIntMap map[int]ProgressionEventList
type ProgressionEventListUIntMap map[uint]ProgressionEventList

// A type that contains a page of results from a progression project events query.
type ProgressionEventList struct {
	ContinuationToken string           `json:"continuationToken"`
	Results           ProgressionEvent `json:"results"`
}
