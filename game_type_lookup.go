package structs

type GameTypeLookupCollection []GameTypeLookup
type GameTypeLookupStringMap map[string]GameTypeLookup
type GameTypeLookupIntMap map[int]GameTypeLookup
type GameTypeLookupUIntMap map[uint]GameTypeLookup

// The result from looking up a game type.
type GameTypeLookup struct {
	GameTypeSimple
	Exact bool `json:"exact"`
}
