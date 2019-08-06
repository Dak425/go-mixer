package mixer

type BroadcastCollection []Broadcast
type BroadcastStringMap map[string]Broadcast
type BroadcastIntMap map[int]Broadcast
type BroadcastUIntMap map[uint]Broadcast

// A broadcast represents a single broadcast from a Channel.
type Broadcast struct {
	ChannelId    uint   `json:"channelId"`
	Id           string `json:"id"`
	IsTestStream bool   `json:"isTestStream"`
	Online       bool   `json:"online"`
	StartedAt    string `json:"startedAt"`
}
