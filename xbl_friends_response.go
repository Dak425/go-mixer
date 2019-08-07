package structs

type XblFriendsResponseCollection []XblFriendsResponse
type XblFriendsResponseStringMap map[string]XblFriendsResponse
type XblFriendsResponseIntMap map[int]XblFriendsResponse
type XblFriendsResponseUIntMap map[uint]XblFriendsResponse

// Response data for call to get Xbox Live friends
type XblFriendsResponse struct {
	AvatarUrl string `json:"avatarUrl"`
	ChannelId uint   `json:"channelId"`
	Level     uint   `json:"level"`
	Online    bool   `json:"online"`
	UserId    uint   `json:"userId"`
	Username  string `json:"username"`
}
