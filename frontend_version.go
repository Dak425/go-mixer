package mixerstructs

type FrontendVersionCollection []FrontendVersion
type FrontendVersionStringMap map[string]FrontendVersion
type FrontendVersionIntMap map[int]FrontendVersion
type FrontendVersionUIntMap map[uint]FrontendVersion

// Describes an available FrontendVersion. Users can sometime swap between frontend versions to test out new features on Mixer.
type FrontendVersion struct {
	DisplayName string `json:"displayName"`
	Version     string `json:"version"`
}
