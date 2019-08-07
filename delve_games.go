package structs

type DelveGamesCollection []DelveGames
type DelveGamesStringMap map[string]DelveGames
type DelveGamesIntMap map[int]DelveGames
type DelveGamesUIntMap map[uint]DelveGames

// A list of games streamed on Mixer.
type DelveGames struct {
	Hydration struct{} `json:"hydration"`
	MorePath  string   `json:"morePath"`
	Title     struct{} `json:"title"`
	Type      string   `json:"type"`
}
