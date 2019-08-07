package structs

type ProgressionActiveStorageProjectCollection []ProgressionActiveStorageProject
type ProgressionActiveStorageProjectStringMap map[string]ProgressionActiveStorageProject
type ProgressionActiveStorageProjectIntMap map[int]ProgressionActiveStorageProject
type ProgressionActiveStorageProjectUIntMap map[uint]ProgressionActiveStorageProject

// A type that contains information about progression storage that has been loaded into a channel.
type ProgressionActiveStorageProject struct {
	ChannelId         uint   `json:"channelId"`
	GameTypeId        uint   `json:"gameTypeId"`
	ProgressionActive bool   `json:"progressionActive"`
	ProjectId         string `json:"projectId"`
	StartTime         string `json:"startTime"`
}
