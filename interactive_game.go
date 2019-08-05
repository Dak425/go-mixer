package mixer

type InteractiveGameCollection []InteractiveGame
type InteractiveGameStringMap map[string]InteractiveGame
type InteractiveGameIntMap map[int]InteractiveGame
type InteractiveGameUIntMap map[uint]InteractiveGame

// An InteractiveGame is owned by a Channel and provides details about an Interactive experience or game on Mixer.
type InteractiveGame struct {
	TimeStamped
	CoverUrl             string `json:"coverUrl"`
	Description          string `json:"description"`
	HasPublishedVersions bool   `json:"hasPublishedVersions"`
	Id                   uint   `json:"id"`
	Installation         string `json:"installation"`
	Name                 string `json:"name"`
	OwnerId              uint   `json:"ownerId"`
	TypeId               uint   `json:"typeId"`
}
