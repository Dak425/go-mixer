package structs

type ClipPropertiesCollection []ClipProperties
type ClipPropertiesStringMap map[string]ClipProperties
type ClipPropertiesIntMap map[int]ClipProperties
type ClipPropertiesUIntMap map[uint]ClipProperties

// Information about a clip.
type ClipProperties struct {
	ContentId         string   `json:"contentId"`
	ContentLocators   Locator  `json:"contentLocators"`
	ContentMaturity   uint     `json:"contentMaturity"`
	DurationInSeconds uint     `json:"durationInSeconds"`
	ExpirationDate    string   `json:"expirationDate"`
	OwnerChannelId    uint     `json:"ownerChannelId"`
	ShareableId       string   `json:"shareableId"`
	StreamerChannelId uint     `json:"streamerChannelId"`
	Tags              []string `json:"tags"`
	Title             string   `json:"title"`
	TypeId            uint     `json:"typeId"`
	UploadDate        string   `json:"uploadDate"`
	ViewCount         uint     `json:"viewCount"`
}
