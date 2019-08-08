package mixerstructs

type EmojiRankAnalyticCollection []EmojiRankAnalytic
type EmojiRankAnalyticStringMap map[string]EmojiRankAnalytic
type EmojiRankAnalyticIntMap map[int]EmojiRankAnalytic
type EmojiRankAnalyticUIntMap map[uint]EmojiRankAnalytic

// A channel metric showing how many times an emoji was used in a period.
type EmojiRankAnalytic struct {
	ChannelAnalytic
	Count uint   `json:"count"`
	Emoji string `json:"emoji"`
}
