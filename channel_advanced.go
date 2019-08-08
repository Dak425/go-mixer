package mixerstructs

type ChannelAdvancedCollection []ChannelAdvanced
type ChannelAdvancedStringMap map[string]ChannelAdvanced
type ChannelAdvancedIntMap map[int]ChannelAdvanced
type ChannelAdvancedUIntMap map[uint]ChannelAdvanced

// Augmented regular channel with additional data.
type ChannelAdvanced struct {
	Channel
	Type GameType       `json:"type"`
	User UserWithGroups `json:"user"`
}
