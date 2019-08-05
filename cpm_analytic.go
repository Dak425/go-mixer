package mixer

type CPMAnalyticCollection []CPMAnalytic
type CPMAnalyticStringMap map[string]CPMAnalytic
type CPMAnalyticIntMap map[int]CPMAnalytic
type CPMAnalyticUIntMap map[uint]CPMAnalytic

// A metric showing the impressions and payout for a period.
type CPMAnalytic struct {
	ChannelAnalytic
	Impressions uint `json:"impressions"`
	Payout      int  `json:"payout"`
}
