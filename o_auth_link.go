package structs

type OAuthLinkCollection []OAuthLink
type OAuthLinkStringMap map[string]OAuthLink
type OAuthLinkIntMap map[int]OAuthLink
type OAuthLinkUIntMap map[uint]OAuthLink

// Represents a linked account on another non-mixer service that this user has linked to their account.
type OAuthLink struct {
	Service   string `json:"service"`
	ServiceId string `json:"serviceId"`
	UserId    int    `json:"userId"`
}
