package mixerstructs

type ChannelAnalyticCollection []ChannelAnalytic
type ChannelAnalyticStringMap map[string]ChannelAnalytic
type ChannelAnalyticIntMap map[int]ChannelAnalytic
type ChannelAnalyticUIntMap map[uint]ChannelAnalytic

// A base analytics type for channel based analytical metrics.
type ChannelAnalytic struct {
	Channel uint   `json:"channel"`
	Time    string `json:"time"`
}
