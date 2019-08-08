package mixerstructs

type InteractiveVersionExpandedCollection []InteractiveVersionExpanded
type InteractiveVersionExpandedStringMap map[string]InteractiveVersionExpanded
type InteractiveVersionExpandedIntMap map[int]InteractiveVersionExpanded
type InteractiveVersionExpandedUIntMap map[uint]InteractiveVersionExpanded

// A versioned instance of an InteractiveGame.
type InteractiveVersionExpanded struct {
	TimeStamped
	Bundle         string              `json:"bundle"`
	Changelog      string              `json:"changelog"`
	ControlVersion string              `json:"controlVersion"`
	Controls       InteractiveControls `json:"controls"`
	Download       string              `json:"download"`
	GameId         uint                `json:"gameId"`
	Id             uint                `json:"id"`
	Installation   string              `json:"installation"`
	State          string              `json:"state"`
	Version        string              `json:"version"`
	VersionOrder   uint                `json:"versionOrder"`
}
