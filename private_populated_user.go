package mixerstructs

type PrivatePopulatedUserCollection []PrivatePopulatedUser
type PrivatePopulatedUserStringMap map[string]PrivatePopulatedUser
type PrivatePopulatedUserIntMap map[int]PrivatePopulatedUser
type PrivatePopulatedUserUIntMap map[uint]PrivatePopulatedUser

// A fully populater user with channel, preferences, groups and private details.
type PrivatePopulatedUser struct {
	PrivateUser
	Channel     Channel         `json:"channel"`
	Groups      UserGroup       `json:"groups"`
	Preferences UserPreferences `json:"preferences"`
}
