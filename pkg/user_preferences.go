package mixer

type UserPreferencesCollection []UserPreferences
type UserPreferencesStringMap map[string]UserPreferences
type UserPreferencesIntMap map[int]UserPreferences
type UserPreferencesUIntMap map[uint]UserPreferences

// Object containing user preferences.
type UserPreferences struct {
	ChannelMatureAllowed    bool     `json:"channel:mature:allowed"`
	ChannelNotifications    struct{} `json:"channel:notifications"`
	ChannelPlayerForceflash bool     `json:"channel:player:forceflash"`
	ChatChromakey           bool     `json:"chat:chromakey"`
	ChatColors              bool     `json:"chat:colors"`
	ChatLurkmode            bool     `json:"chat:lurkmode"`
	ChatSoundsHtml5         bool     `json:"chat:sounds:html5"`
	ChatSoundsPlay          bool     `json:"chat:sounds:play"`
	ChatSoundsVolume        int      `json:"chat:sounds:volume"`
	ChatTagging             bool     `json:"chat:tagging"`
	ChatTimestamps          bool     `json:"chat:timestamps"`
	ChatWhispers            bool     `json:"chat:whispers"`
}
