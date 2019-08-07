package structs

type LightVideoManifestCollection []LightVideoManifest
type LightVideoManifestStringMap map[string]LightVideoManifest
type LightVideoManifestIntMap map[int]LightVideoManifest
type LightVideoManifestUIntMap map[uint]LightVideoManifest

// Describes a broadcast using the Light manifest structure.
type LightVideoManifest struct {
	Resolutions VideoManifestResolution `json:"resolutions"`
	Since       string                  `json:"since"`
}
