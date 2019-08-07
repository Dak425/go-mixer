package structs

type ProgressionActiveStorageProjectListCollection []ProgressionActiveStorageProjectList
type ProgressionActiveStorageProjectListStringMap map[string]ProgressionActiveStorageProjectList
type ProgressionActiveStorageProjectListIntMap map[int]ProgressionActiveStorageProjectList
type ProgressionActiveStorageProjectListUIntMap map[uint]ProgressionActiveStorageProjectList

// A type that contains a page of results from an active progression project query.
type ProgressionActiveStorageProjectList struct {
	ContinuationToken string                          `json:"continuationToken"`
	Results           ProgressionActiveStorageProject `json:"results"`
}
