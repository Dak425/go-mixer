package mixer

type FTLVideoManifestCollection []FTLVideoManifest
type FTLVideoManifestStringMap map[string]FTLVideoManifest
type FTLVideoManifestIntMap map[int]FTLVideoManifest
type FTLVideoManifestUIntMap map[uint]FTLVideoManifest

// Provides details and information about an FTL Broadcast.
type FTLVideoManifest struct {
	Resolutions VideoManifestResolution `json:"resolutions"`
	Since       string                  `json:"since"`
}
