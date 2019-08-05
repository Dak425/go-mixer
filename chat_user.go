package mixer

type ChatUserCollection []ChatUser
type ChatUserStringMap map[string]ChatUser
type ChatUserIntMap map[int]ChatUser
type ChatUserUIntMap map[uint]ChatUser

// Represents a user within chat for a channel and provides details about their status within the channel.
type ChatUser struct {
	User      User     `json:"user"`
	UserId    uint     `json:"userId"`
	UserName  string   `json:"userName"`
	UserRoles []string `json:"userRoles"`
}
