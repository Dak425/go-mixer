package structs

type UserInteractiveMetricCollection []UserInteractiveMetric
type UserInteractiveMetricStringMap map[string]UserInteractiveMetric
type UserInteractiveMetricIntMap map[int]UserInteractiveMetric
type UserInteractiveMetricUIntMap map[uint]UserInteractiveMetric

// A user interactive metric represents a record for every time a user views an channel with interactive running for any period of time
type UserInteractiveMetric struct {
	Browser        string `json:"browser"`
	ChannelId      string `json:"channelId"`
	Id             string `json:"id"`
	InputsGiven    uint   `json:"inputsGiven"`
	Platform       string `json:"platform"`
	Playtime       uint   `json:"playtime"`
	Referrer       string `json:"referrer"`
	Region         string `json:"region"`
	SparksSpent    uint   `json:"sparksSpent"`
	TimeToInteract uint   `json:"timeToInteract"`
	Timestamp      string `json:"timestamp"`
	UserId         string `json:"userId"`
	VersionId      string `json:"versionId"`
}
