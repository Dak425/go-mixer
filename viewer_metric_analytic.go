package mixerstructs

type ViewerMetricAnalyticCollection []ViewerMetricAnalytic
type ViewerMetricAnalyticStringMap map[string]ViewerMetricAnalytic
type ViewerMetricAnalyticIntMap map[int]ViewerMetricAnalytic
type ViewerMetricAnalyticUIntMap map[uint]ViewerMetricAnalytic

// A channel metric showing details about a viewer on a channel.
type ViewerMetricAnalytic struct {
	ChannelAnalytic
	Browser  string `json:"browser"`
	Country  string `json:"country"`
	Platform string `json:"platform"`
}
