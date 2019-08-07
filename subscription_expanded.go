package structs

type SubscriptionExpandedCollection []SubscriptionExpanded
type SubscriptionExpandedStringMap map[string]SubscriptionExpanded
type SubscriptionExpandedIntMap map[int]SubscriptionExpanded
type SubscriptionExpandedUIntMap map[uint]SubscriptionExpanded

// The subscription belonging to this payment.
type SubscriptionExpanded struct {
	Subscription
	Group UserGroup `json:"group"`
	User  User      `json:"user"`
}
