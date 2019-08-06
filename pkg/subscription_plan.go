package mixer

type SubscriptionPlanCollection []SubscriptionPlan
type SubscriptionPlanStringMap map[string]SubscriptionPlan
type SubscriptionPlanIntMap map[int]SubscriptionPlan
type SubscriptionPlanUIntMap map[uint]SubscriptionPlan

// Details on a subscription plan
type SubscriptionPlan struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
	Id       string `json:"id"`
	Name     string `json:"name"`
	Type     string `json:"type"`
}
