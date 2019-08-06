package mixer

type ExpandedTranscodingProfileCollection []ExpandedTranscodingProfile
type ExpandedTranscodingProfileStringMap map[string]ExpandedTranscodingProfile
type ExpandedTranscodingProfileIntMap map[int]ExpandedTranscodingProfile
type ExpandedTranscodingProfileUIntMap map[uint]ExpandedTranscodingProfile

// Augments a regular transcoding profile by listing the transcodes available on this profile.
type ExpandedTranscodingProfile struct {
	TranscodingProfile
	Transcodes TranscodingProfileTranscode `json:"transcodes"`
}
