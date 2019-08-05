package mixer

type FollowCollection []Follow
type FollowStringMap map[string]Follow
type FollowIntMap map[int]Follow
type FollowUIntMap map[uint]Follow

// Represents a follow relationship between a user and a channel.
type Follow struct {
	TimeStamped
	Channel uint `json:"channel"`
	User    uint `json:"user"`
}
