package structs

type ExpandedShareCollection []ExpandedShare
type ExpandedShareStringMap map[string]ExpandedShare
type ExpandedShareIntMap map[int]ExpandedShare
type ExpandedShareUIntMap map[uint]ExpandedShare

// Augmented share with additional properties.
type ExpandedShare struct {
	Share
	User User `json:"user"`
}
