package mixerstructs

type CostreamCollection []Costream
type CostreamStringMap map[string]Costream
type CostreamIntMap map[int]Costream
type CostreamUIntMap map[uint]Costream

// A co-stream is a set of channels grouped together and displayed on Mixer together.
type Costream struct {
	Capacity  uint              `json:"capacity"`
	Channels  ChannelCollection `json:"channels"`
	Layout    PlayerLayout      `json:"layout"`
	StartedAt string            `json:"startedAt"`
}
