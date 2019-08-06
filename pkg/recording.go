package mixer

type RecordingCollection []Recording
type RecordingStringMap map[string]Recording
type RecordingIntMap map[int]Recording
type RecordingUIntMap map[uint]Recording

// Details about a recorded Broadcast from a channel.
type Recording struct {
	TimeStamped
	ChannelId  uint   `json:"channelId"`
	Duration   int    `json:"duration"`
	ExpiresAt  string `json:"expiresAt"`
	Id         uint   `json:"id"`
	Name       string `json:"name"`
	Seen       bool   `json:"seen"`
	State      string `json:"state"`
	TypeId     uint   `json:"typeId"`
	Viewed     bool   `json:"viewed"`
	ViewsTotal uint   `json:"viewsTotal"`
	Vods       VOD    `json:"vods"`
}
