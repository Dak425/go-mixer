package mixerstructs

type StreamSessionsAnalyticCollection []StreamSessionsAnalytic
type StreamSessionsAnalyticStringMap map[string]StreamSessionsAnalytic
type StreamSessionsAnalyticIntMap map[int]StreamSessionsAnalytic
type StreamSessionsAnalyticUIntMap map[uint]StreamSessionsAnalytic

// Metric showing a channel going online or offline.
type StreamSessionsAnalytic struct {
	ChannelAnalytic
	Duration  uint `json:"duration"`
	Online    bool `json:"online"`
	Partnered bool `json:"partnered"`
	Type      uint `json:"type"`
}
