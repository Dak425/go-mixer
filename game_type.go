package mixerstructs

type GameTypeCollection []GameType
type GameTypeStringMap map[string]GameType
type GameTypeIntMap map[int]GameType
type GameTypeUIntMap map[uint]GameType

// A GameType can be set on a channel and represents the title they are broadcasting.
type GameType struct {
	GameTypeSimple
	Description    string `json:"description"`
	Online         uint   `json:"online"`
	Parent         string `json:"parent"`
	Source         string `json:"source"`
	ViewersCurrent uint   `json:"viewersCurrent"`
}
