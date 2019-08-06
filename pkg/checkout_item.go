package mixer

type CheckoutItemCollection []CheckoutItem
type CheckoutItemStringMap map[string]CheckoutItem
type CheckoutItemIntMap map[int]CheckoutItem
type CheckoutItemUIntMap map[uint]CheckoutItem

// One item within an user's Order.
type CheckoutItem struct {
	Id       uint `json:"id"`
	Quantity uint `json:"quantity"`
}
