package mixer

type ChannelCollection []Channel
type ChannelStringMap map[string]Channel
type ChannelIntMap map[int]Channel
type ChannelUIntMap map[uint]Channel

// A single channel within Mixer. Each channel is owned by a user, and a channel can be broadcasted to.
type Channel struct {
	TimeStamped
	Audience             string `json:"audience"`
	BadgeId              uint   `json:"badgeId"`
	BannerUrl            string `json:"bannerUrl"`
	CostreamId           string `json:"costreamId"`
	CoverId              uint   `json:"coverId"`
	Description          string `json:"description"`
	FeatureLevel         int    `json:"featureLevel"`
	Featured             bool   `json:"featured"`
	Ftl                  uint   `json:"ftl"`
	HasTranscodes        bool   `json:"hasTranscodes"`
	HasVod               bool   `json:"hasVod"`
	HosteeId             uint   `json:"hosteeId"`
	Id                   uint   `json:"id"`
	Interactive          bool   `json:"interactive"`
	InteractiveGameId    uint   `json:"interactiveGameId"`
	LanguageId           string `json:"languageId"`
	Name                 string `json:"name"`
	NumFollowers         uint   `json:"numFollowers"`
	Online               bool   `json:"online"`
	Partnered            bool   `json:"partnered"`
	Suspended            bool   `json:"suspended"`
	ThumbnailId          uint   `json:"thumbnailId"`
	Token                string `json:"token"`
	TranscodingProfileId uint   `json:"transcodingProfileId"`
	TypeId               uint   `json:"typeId"`
	UserId               uint   `json:"userId"`
	ViewersCurrent       uint   `json:"viewersCurrent"`
	ViewersTotal         uint   `json:"viewersTotal"`
	VodsEnabled          bool   `json:"vodsEnabled"`
}
