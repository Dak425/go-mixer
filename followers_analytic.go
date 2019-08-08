package mixerstructs

type FollowersAnalyticCollection []FollowersAnalytic
type FollowersAnalyticStringMap map[string]FollowersAnalytic
type FollowersAnalyticIntMap map[int]FollowersAnalytic
type FollowersAnalyticUIntMap map[uint]FollowersAnalytic

// A channel metric showing how many followers the channel had at this time and the change in that metric.
type FollowersAnalytic struct {
	ChannelAnalytic
	Delta int  `json:"delta"`
	Total uint `json:"total"`
	User  int  `json:"user"`
}
