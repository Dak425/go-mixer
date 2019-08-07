package structs

type SubscriptionsAnalyticCollection []SubscriptionsAnalytic
type SubscriptionsAnalyticStringMap map[string]SubscriptionsAnalytic
type SubscriptionsAnalyticIntMap map[int]SubscriptionsAnalytic
type SubscriptionsAnalyticUIntMap map[uint]SubscriptionsAnalytic

// A channel metric that shows when a subscription was made to a channel.
type SubscriptionsAnalytic struct {
	ChannelAnalytic
	Delta int  `json:"delta"`
	Total uint `json:"total"`
	User  uint `json:"user"`
}
