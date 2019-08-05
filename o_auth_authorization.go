package mixer

type OAuthAuthorizationCollection []OAuthAuthorization
type OAuthAuthorizationStringMap map[string]OAuthAuthorization
type OAuthAuthorizationIntMap map[int]OAuthAuthorization
type OAuthAuthorizationUIntMap map[uint]OAuthAuthorization

// An OAuthClient which has been authorized by the user.
type OAuthAuthorization struct {
	Client      OAuthClient `json:"client"`
	Permissions []string    `json:"permissions"`
	UserID      uint        `json:"userId"`
}
