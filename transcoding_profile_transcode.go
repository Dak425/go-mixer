package mixerstructs

type TranscodingProfileTranscodeCollection []TranscodingProfileTranscode
type TranscodingProfileTranscodeStringMap map[string]TranscodingProfileTranscode
type TranscodingProfileTranscodeIntMap map[int]TranscodingProfileTranscode
type TranscodingProfileTranscodeUIntMap map[uint]TranscodingProfileTranscode

// Represents a single transcode option for a channel.
type TranscodingProfileTranscode struct {
	Bitrate         uint   `json:"bitrate"`
	Fps             uint   `json:"fps"`
	Height          uint   `json:"height"`
	Id              uint   `json:"id"`
	Name            string `json:"name"`
	RequiresPartner bool   `json:"requiresPartner"`
	Title           string `json:"title"`
	Width           uint   `json:"width"`
}
