package structs

type ViewerAnalyticCollection []ViewerAnalytic
type ViewerAnalyticStringMap map[string]ViewerAnalytic
type ViewerAnalyticIntMap map[int]ViewerAnalytic
type ViewerAnalyticUIntMap map[uint]ViewerAnalytic

// A channel metric, which contains the number of anonymous and authenticated Mixer users at a point in time.
type ViewerAnalytic struct {
	ChannelAnalytic
	Anon   uint `json:"anon"`
	Authed uint `json:"authed"`
}
