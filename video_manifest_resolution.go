package structs

type VideoManifestResolutionCollection []VideoManifestResolution
type VideoManifestResolutionStringMap map[string]VideoManifestResolution
type VideoManifestResolutionIntMap map[int]VideoManifestResolution
type VideoManifestResolutionUIntMap map[uint]VideoManifestResolution

// A resoloution available for a video manifest.
type VideoManifestResolution struct {
	Bitrate  uint   `json:"bitrate"`
	HasVideo bool   `json:"hasVideo"`
	Height   uint   `json:"height"`
	Name     string `json:"name"`
	Slug     string `json:"slug"`
	Width    uint   `json:"width"`
}
