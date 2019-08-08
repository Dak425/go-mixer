package mixerstructs

type ProgressionProjectCollection []ProgressionProject
type ProgressionProjectStringMap map[string]ProgressionProject
type ProgressionProjectIntMap map[int]ProgressionProject
type ProgressionProjectUIntMap map[uint]ProgressionProject

// A type that contains information about progression storage across users for a specific game.
type ProgressionProject struct {
	ApplicationAccess             struct{}             `json:"applicationAccess"`
	FiveMinuteHeartbeatOperations ProgressionOperation `json:"fiveMinuteHeartbeatOperations"`
	GameTypeId                    uint                 `json:"gameTypeId"`
	Id                            string               `json:"id"`
	IsPublished                   bool                 `json:"isPublished"`
	Name                          string               `json:"name"`
	OwningUsers                   uint                 `json:"owningUsers"`
	XboxDecimalTitleIdAccess      struct{}             `json:"xboxDecimalTitleIdAccess"`
}
