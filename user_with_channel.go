package mixer

type UserWithChannelCollection []UserWithChannel
type UserWithChannelStringMap map[string]UserWithChannel
type UserWithChannelIntMap map[int]UserWithChannel
type UserWithChannelUIntMap map[uint]UserWithChannel

// A User object with an embedded Channel.
type UserWithChannel struct {
	User
	Channel Channel `json:"channel"`
}
