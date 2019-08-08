package mixerstructs

type PartnershipAppCollection []PartnershipApp
type PartnershipAppStringMap map[string]PartnershipApp
type PartnershipAppIntMap map[int]PartnershipApp
type PartnershipAppUIntMap map[uint]PartnershipApp

// Contains details about a particular partnership application.
type PartnershipApp struct {
	TimeStamped
	Reapplies string `json:"reapplies"`
	Reason    string `json:"reason"`
	Status    string `json:"status"`
}
