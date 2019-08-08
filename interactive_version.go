package mixerstructs

type InteractiveVersionCollection []InteractiveVersion
type InteractiveVersionStringMap map[string]InteractiveVersion
type InteractiveVersionIntMap map[int]InteractiveVersion
type InteractiveVersionUIntMap map[uint]InteractiveVersion

// A versioned instance of an InteractiveGame.
type InteractiveVersion struct {
	TimeStamped
	Bundle         string              `json:"bundle"`
	Changelog      string              `json:"changelog"`
	ControlVersion string              `json:"controlVersion"`
	Controls       InteractiveControls `json:"controls"`
	Download       string              `json:"download"`
	GameID         uint                `json:"gameId"`
	ID             uint                `json:"id"`
	Installation   string              `json:"installation"`
	State          string              `json:"state"`
	Version        string              `json:"version"`
	VersionOrder   uint                `json:"versionOrder"`
}
