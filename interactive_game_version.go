package structs

type InteractiveGameVersionCollection []InteractiveGameVersion

type InteractiveGameVersion struct {
	ID           uint   `json:"id"`
	Version      string `json:"version"`
	State        string `json:"state"`
	VersionOrder uint   `json:"versionOrder"`
}
