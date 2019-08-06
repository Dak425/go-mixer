package mixer

type InteractiveGameListingCollection []InteractiveGameListing
type InteractiveGameListingStringMap map[string]InteractiveGameListing
type InteractiveGameListingIntMap map[int]InteractiveGameListing
type InteractiveGameListingUIntMap map[uint]InteractiveGameListing

// An interactive game within a list of Interactive games. Contains additional properties to avoid extra requests.
type InteractiveGameListing struct {
	InteractiveGame
	Owner    User                             `json:"owner"`
	Versions InteractiveGameVersionCollection `json:"versions"`
}
