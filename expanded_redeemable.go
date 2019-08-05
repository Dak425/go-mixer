package mixer

type ExpandedRedeemableCollection []ExpandedRedeemable
type ExpandedRedeemableStringMap map[string]ExpandedRedeemable
type ExpandedRedeemableIntMap map[int]ExpandedRedeemable
type ExpandedRedeemableUIntMap map[uint]ExpandedRedeemable

// Augmented Redeemable with additional properties
type ExpandedRedeemable struct {
	Redeemable
	Owner      User `json:"owner"`
	RedeemedBy User `json:"redeemedBy"`
}
