package structs

type InteractiveAnalyticCollection []InteractiveAnalytic
type InteractiveAnalyticStringMap map[string]InteractiveAnalytic
type InteractiveAnalyticIntMap map[int]InteractiveAnalytic
type InteractiveAnalyticUIntMap map[uint]InteractiveAnalytic

// A channel metric showing details about the type this channel is broadcasting.
type InteractiveAnalytic struct {
	ChannelAnalytic
	InputsGiven       uint     `json:"inputsGiven"`
	Playtime          uint     `json:"playtime"`
	SparksSpent       struct{} `json:"sparksSpent"`
	StreamLength      uint     `json:"streamLength"`
	Streams           uint     `json:"streams"`
	TimeToInteractive struct{} `json:"timeToInteractive"`
	Viewers           uint     `json:"viewers"`
}
