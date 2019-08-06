package mixer

type InteractiveMetricCollection []InteractiveMetric
type InteractiveMetricStringMap map[string]InteractiveMetric
type InteractiveMetricIntMap map[int]InteractiveMetric
type InteractiveMetricUIntMap map[uint]InteractiveMetric

// An interactive metric represents a record for every time a channel enables interactive for any period of time
type InteractiveMetric struct {
	Bundle             string `json:"bundle"`
	ChannelId          string `json:"channelId"`
	Id                 string `json:"id"`
	MixerTypeId        string `json:"mixerTypeId"`
	ProjectId          string `json:"projectId"`
	Published          bool   `json:"published"`
	PublisherUserId    string `json:"publisherUserId"`
	SessionLength      uint   `json:"sessionLength"`
	Timestamp          string `json:"timestamp"`
	TotalPlaytimeHours uint   `json:"totalPlaytimeHours"`
	VersionId          string `json:"versionId"`
}
