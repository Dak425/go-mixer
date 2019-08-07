package structs

type RedeemableCollection []Redeemable
type RedeemableStringMap map[string]Redeemable
type RedeemableIntMap map[int]Redeemable
type RedeemableUIntMap map[uint]Redeemable

// A redeemable is an item granted to a user that can be exchanged for an item or activity. These are currently only used for Mixer Pro membership and channel subscriptions.
type Redeemable struct {
	TimeStamped
	Code         string   `json:"code"`
	Id           uint     `json:"id"`
	Meta         struct{} `json:"meta"`
	OwnerId      uint     `json:"ownerId"`
	RedeemedAt   string   `json:"redeemedAt"`
	RedeemedById uint     `json:"redeemedById"`
	State        string   `json:"state"`
	Type         string   `json:"type"`
}
