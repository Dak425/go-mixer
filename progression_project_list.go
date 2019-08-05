package mixer

type ProgressionProjectListCollection []ProgressionProjectList
type ProgressionProjectListStringMap map[string]ProgressionProjectList
type ProgressionProjectListIntMap map[int]ProgressionProjectList
type ProgressionProjectListUIntMap map[uint]ProgressionProjectList

// A type that contains a page of results from a progression project query.
type ProgressionProjectList struct {
	ContinuationToken string             `json:"continuationToken"`
	Results           ProgressionProject `json:"results"`
}
