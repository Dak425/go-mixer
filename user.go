package structs

type UserCollection []User
type UserStringMap map[string]User
type UserIntMap map[int]User
type UserUIntMap map[uint]User

// A user is a person on Mixer; they can sign in and interact with the site. Each user owns a channel, which they can broadcast to.
type User struct {
	TimeStamped
	AvatarUrl   string     `json:"avatarUrl"`
	Bio         string     `json:"bio"`
	Email       string     `json:"email"`
	Experience  uint       `json:"experience"`
	Id          uint       `json:"id"`
	Level       uint       `json:"level"`
	PrimaryTeam uint       `json:"primaryTeam"`
	Social      SocialInfo `json:"social"`
	Sparks      uint       `json:"sparks"`
	Username    string     `json:"username"`
	Verified    bool       `json:"verified"`
}
