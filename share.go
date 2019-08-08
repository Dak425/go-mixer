package mixerstructs

type ShareCollection []Share
type ShareStringMap map[string]Share
type ShareIntMap map[int]Share
type ShareUIntMap map[uint]Share

// A share allows a user to share an entity with another user. Currently this is only used for InteractiveGames.
type Share struct {
	TimeStamped
	Code   string `json:"code"`
	Id     uint   `json:"id"`
	Relid  string `json:"relid"`
	Type   string `json:"type"`
	UserId uint   `json:"userId"`
}
