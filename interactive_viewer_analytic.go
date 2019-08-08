package mixerstructs

type InteractiveViewerAnalyticCollection []InteractiveViewerAnalytic
type InteractiveViewerAnalyticStringMap map[string]InteractiveViewerAnalytic
type InteractiveViewerAnalyticIntMap map[int]InteractiveViewerAnalytic
type InteractiveViewerAnalyticUIntMap map[uint]InteractiveViewerAnalytic

// A channel metric showing details about a participant in the Interactive game.
type InteractiveViewerAnalytic struct {
	ChannelAnalytic
	Browser   string `json:"browser"`
	ChannelId int    `json:"channelId"`
	Country   string `json:"country"`
	Platform  string `json:"platform"`
	Referrer  string `json:"referrer"`
}
