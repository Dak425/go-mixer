package structs

type RedeemablePartnerCollection []RedeemablePartner
type RedeemablePartnerStringMap map[string]RedeemablePartner
type RedeemablePartnerIntMap map[int]RedeemablePartner
type RedeemablePartnerUIntMap map[uint]RedeemablePartner

// Extra information added to a redeemable when it is a Partner's monthly subscription / Pro codes.
type RedeemablePartner struct {
	TimeStamped
	Id           uint `json:"id"`
	PartnerId    uint `json:"partnerId"`
	RedeemableId uint `json:"redeemableId"`
}
