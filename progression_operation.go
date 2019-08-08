package mixerstructs

type ProgressionOperationCollection []ProgressionOperation
type ProgressionOperationStringMap map[string]ProgressionOperation
type ProgressionOperationIntMap map[int]ProgressionOperation
type ProgressionOperationUIntMap map[uint]ProgressionOperation

// A type that contains information about an operation to run against a player or channel viewers.
type ProgressionOperation struct {
	Amount              int    `json:"amount"`
	AwardMessage        string `json:"awardMessage"`
	AwardMessageDetails string `json:"awardMessageDetails"`
	MaxVal              int    `json:"maxVal"`
	MinVal              int    `json:"minVal"`
	NotificationEventId string `json:"notificationEventId"`
	Op                  string `json:"op"`
	SessionIds          string `json:"sessionIds"`
	Stat                string `json:"stat"`
	TimeRestriction     string `json:"timeRestriction"`
}
