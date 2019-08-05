package mixer

type SubscriptionWithGroupCollection []SubscriptionWithGroup
type SubscriptionWithGroupStringMap map[string]SubscriptionWithGroup
type SubscriptionWithGroupIntMap map[int]SubscriptionWithGroup
type SubscriptionWithGroupUIntMap map[uint]SubscriptionWithGroup

// A subscription with added group information as a nested object.
type SubscriptionWithGroup struct {
	Subscription
	GroupStruct UserGroup `json:"Group"`
	Group       uint      `json:"group"`
}
