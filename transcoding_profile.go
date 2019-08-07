package structs

type TranscodingProfileCollection []TranscodingProfile
type TranscodingProfileStringMap map[string]TranscodingProfile
type TranscodingProfileIntMap map[int]TranscodingProfile
type TranscodingProfileUIntMap map[uint]TranscodingProfile

// A transcoding profile controls the transcoding options for a channel.
type TranscodingProfile struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}
