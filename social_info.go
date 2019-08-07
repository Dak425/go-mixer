package structs

type SocialInfoCollection []SocialInfo
type SocialInfoStringMap map[string]SocialInfo
type SocialInfoIntMap map[int]SocialInfo
type SocialInfoUIntMap map[uint]SocialInfo

// The social information for a channel.
type SocialInfo struct {
	Discord  string   `json:"discord"`
	Facebook string   `json:"facebook"`
	Player   string   `json:"player"`
	Twitter  string   `json:"twitter"`
	Verified []string `json:"verified"`
	Youtube  string   `json:"youtube"`
}
