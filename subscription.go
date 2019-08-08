package mixerstructs

type SubscriptionCollection []Subscription
type SubscriptionStringMap map[string]Subscription
type SubscriptionIntMap map[int]Subscription
type SubscriptionUIntMap map[uint]Subscription

// A subscription lists details about a user's subscription to a resource.
type Subscription struct {
	TimeStamped
	Cancelled    bool   `json:"cancelled"`
	ExpiresAt    string `json:"expiresAt"`
	Id           uint   `json:"id"`
	ResourceId   uint   `json:"resourceId"`
	ResourceType string `json:"resourceType"`
	Status       string `json:"status"`
}
