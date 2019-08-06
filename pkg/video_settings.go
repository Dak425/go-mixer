package mixer

type VideoSettingsCollection []VideoSettings
type VideoSettingsStringMap map[string]VideoSettings
type VideoSettingsIntMap map[int]VideoSettings
type VideoSettingsUIntMap map[uint]VideoSettings

// Video settings linked to a user's channel
type VideoSettings struct {
	IsLightstreamEnabled bool `json:"isLightstreamEnabled"`
}
