package mixerstructs

type ProgressionPlayerStorageCollection []ProgressionPlayerStorage
type ProgressionPlayerStorageStringMap map[string]ProgressionPlayerStorage
type ProgressionPlayerStorageIntMap map[int]ProgressionPlayerStorage
type ProgressionPlayerStorageUIntMap map[uint]ProgressionPlayerStorage

// A type that contains information about progression storage for a player.
type ProgressionPlayerStorage struct {
	ProjectId  string               `json:"projectId"`
	Statistics ProgressionStatistic `json:"statistics"`
}
