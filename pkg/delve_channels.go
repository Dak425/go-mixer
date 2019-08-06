package mixer

type DelveChannelsCollection []DelveChannels
type DelveChannelsStringMap map[string]DelveChannels
type DelveChannelsIntMap map[int]DelveChannels
type DelveChannelsUIntMap map[uint]DelveChannels

// A list of channel on Mixer.
type DelveChannels struct {
	Hydration struct{} `json:"hydration"`
	MorePath  string   `json:"morePath"`
	Style     string   `json:"style"`
	Title     struct{} `json:"title"`
	Type      string   `json:"type"`
}
