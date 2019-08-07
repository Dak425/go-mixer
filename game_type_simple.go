package structs

type GameTypeSimpleCollection []GameTypeSimple
type GameTypeSimpleStringMap map[string]GameTypeSimple
type GameTypeSimpleIntMap map[int]GameTypeSimple
type GameTypeSimpleUIntMap map[uint]GameTypeSimple

// Base game type.
type GameTypeSimple struct {
	BackgroundUrl string `json:"backgroundUrl"`
	CoverUrl      string `json:"coverUrl"`
	Id            uint   `json:"id"`
	Name          string `json:"name"`
}
