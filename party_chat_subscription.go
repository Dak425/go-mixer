package structs

type PartyChatSubscriptionCollection []PartyChatSubscription
type PartyChatSubscriptionStringMap map[string]PartyChatSubscription
type PartyChatSubscriptionIntMap map[int]PartyChatSubscription
type PartyChatSubscriptionUIntMap map[uint]PartyChatSubscription

// Provided to the Mixer server, then MPSD, when the client sets up an RTA subscription. See the Book of Party Chat / MPSD reference for more information about this, and how to use it.
type PartyChatSubscription struct {
	Active         bool     `json:"active"`
	ChangeTypes    []string `json:"changeTypes"`
	Connection     string   `json:"connection"`
	SubscriptionId string   `json:"subscriptionId"`
}
