package mixer

type ProgressionPlayerStorageListCollection []ProgressionPlayerStorageList
type ProgressionPlayerStorageListStringMap map[string]ProgressionPlayerStorageList
type ProgressionPlayerStorageListIntMap map[int]ProgressionPlayerStorageList
type ProgressionPlayerStorageListUIntMap map[uint]ProgressionPlayerStorageList

// A type that contains a page of results from an progression storage query.
type ProgressionPlayerStorageList struct {
	ContinuationToken string                   `json:"continuationToken"`
	Results           ProgressionPlayerStorage `json:"results"`
}
