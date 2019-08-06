package mixer

type ViewerSessionCountAnalyticCollection []ViewerSessionCountAnalytic
type ViewerSessionCountAnalyticStringMap map[string]ViewerSessionCountAnalytic
type ViewerSessionCountAnalyticIntMap map[int]ViewerSessionCountAnalytic
type ViewerSessionCountAnalyticUIntMap map[uint]ViewerSessionCountAnalytic

// A metric showing the number of viewers in a period on a channel.
type ViewerSessionCountAnalytic struct {
	ChannelAnalytic
	Count uint `json:"count"`
}
