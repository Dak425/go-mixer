package mixer

type EmoticonPackCollection []EmoticonPack
type EmoticonPackStringMap map[string]EmoticonPack
type EmoticonPackIntMap map[int]EmoticonPack
type EmoticonPackUIntMap map[uint]EmoticonPack

// An emoticon pack is a list of emoticons that belong to a channel.
type EmoticonPack struct {
	ChannelId uint          `json:"channelId"`
	Emoticons EmoticonGroup `json:"emoticons"`
	Url       string        `json:"url"`
}
