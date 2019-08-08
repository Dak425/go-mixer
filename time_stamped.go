package mixerstructs

type TimeStampedCollection []TimeStamped
type TimeStampedStringMap map[string]TimeStamped
type TimeStampedIntMap map[int]TimeStamped
type TimeStampedUIntMap map[uint]TimeStamped

// A type that contains information about creation, update and deletion dates.
type TimeStamped struct {
	CreatedAt string `json:"createdAt"`
	DeletedAt string `json:"deletedAt"`
	UpdatedAt string `json:"updatedAt"`
}
