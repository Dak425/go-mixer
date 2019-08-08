package mixerstructs

type StreamHostsAnalyticCollection []StreamHostsAnalytic
type StreamHostsAnalyticStringMap map[string]StreamHostsAnalytic
type StreamHostsAnalyticIntMap map[int]StreamHostsAnalytic
type StreamHostsAnalyticUIntMap map[uint]StreamHostsAnalytic

// Metric showing when a channel is hosted.
type StreamHostsAnalytic struct {
	ChannelAnalytic
	Hostee uint `json:"hostee"`
	Hoster uint `json:"hoster"`
}
