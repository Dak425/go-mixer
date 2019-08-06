package mixer

type GameRankAnalyticCollection []GameRankAnalytic
type GameRankAnalyticStringMap map[string]GameRankAnalytic
type GameRankAnalyticIntMap map[int]GameRankAnalytic
type GameRankAnalyticUIntMap map[uint]GameRankAnalytic

// A channel metric showing details about the type this channel is broadcasting.
type GameRankAnalytic struct {
	ChannelAnalytic
	Shared  uint `json:"shared"`
	Streams uint `json:"streams"`
	Views   uint `json:"views"`
}
