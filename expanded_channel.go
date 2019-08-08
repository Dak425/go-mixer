package mixerstructs

type ExpandedChannelCollection []ExpandedChannel
type ExpandedChannelStringMap map[string]ExpandedChannel
type ExpandedChannelIntMap map[int]ExpandedChannel
type ExpandedChannelUIntMap map[uint]ExpandedChannel

// Augmented ChannelAdvanced with additional properties.
type ExpandedChannel struct {
	ChannelAdvanced
	Badge       Resource           `json:"badge"`
	Cache       struct{}           `json:"cache"` // Deprecated
	Cover       Resource           `json:"cover"`
	Preferences ChannelPreferences `json:"preferences"`
	Thumbnail   Resource           `json:"thumbnail"`
}
