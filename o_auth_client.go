package mixer

type OAuthClientCollection []OAuthClient
type OAuthClientStringMap map[string]OAuthClient
type OAuthClientIntMap map[int]OAuthClient
type OAuthClientUIntMap map[uint]OAuthClient

// An OAuthClient is an application or tool created by a developer. They can be used to ask for access to a user's account.
type OAuthClient struct {
	ClientID string   `json:"clientId"`
	Hosts    []string `json:"hosts"`
	ID       uint     `json:"id"`
	Logo     string   `json:"logo"`
	Name     string   `json:"name"`
	Website  string   `json:"website"`
}
