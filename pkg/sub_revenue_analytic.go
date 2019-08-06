package mixer

type SubRevenueAnalyticCollection []SubRevenueAnalytic
type SubRevenueAnalyticStringMap map[string]SubRevenueAnalytic
type SubRevenueAnalyticIntMap map[int]SubRevenueAnalytic
type SubRevenueAnalyticUIntMap map[uint]SubRevenueAnalytic

// A channel metric that shows when a channel's subscription revenue changes.
type SubRevenueAnalytic struct {
	ChannelAnalytic
	Count   uint   `json:"count"`
	Gateway string `json:"gateway"`
	Gross   int    `json:"gross"`
	Total   int    `json:"total"`
}
