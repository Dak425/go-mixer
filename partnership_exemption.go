package mixer

type PartnershipExemptionCollection []PartnershipExemption
type PartnershipExemptionStringMap map[string]PartnershipExemption
type PartnershipExemptionIntMap map[int]PartnershipExemption
type PartnershipExemptionUIntMap map[uint]PartnershipExemption

// An exemption that allows a specific channel to bypass the minimum requirements check for partnership.
type PartnershipExemption struct {
	TimeStamped
	Agency     string `json:"agency"`
	ApproverID int    `json:"approverId"`
	ChannelID  int    `json:"channelId"`
	ID         int    `json:"id"`
	Reason     string `json:"reason"`
	Revoked    bool   `json:"revoked"`
	RevokerID  int    `json:"revokerId"`
}
