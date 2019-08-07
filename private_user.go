package structs

type PrivateUserCollection []PrivateUser
type PrivateUserStringMap map[string]PrivateUser
type PrivateUserIntMap map[int]PrivateUser
type PrivateUserUIntMap map[uint]PrivateUser

type TwoFactorSettings struct {
	Enabled     bool `json:"enabled"`
	CodesViewed bool `json:"codesViewed"`
}

// A fully populater user with channel, preferences, groups and private details.
type PrivateUser struct {
	User
	Email     string            `json:"email"`
	Password  string            `json:"password"`
	TwoFactor TwoFactorSettings `json:"twoFactor"`
}
