package structs

type TeamCollection []Team
type TeamStringMap map[string]Team
type TeamIntMap map[int]Team
type TeamUIntMap map[uint]Team

// A team is a group of users.
type Team struct {
	TimeStamped
	BackgroundUrl string     `json:"backgroundUrl"`
	Description   string     `json:"description"`
	Id            uint       `json:"id"`
	LogoUrl       string     `json:"logoUrl"`
	Name          string     `json:"name"`
	OwnerId       uint       `json:"ownerId"`
	Social        SocialInfo `json:"social"`
	Token         string     `json:"token"`
}
