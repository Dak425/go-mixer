package structs

type PartnershipCriterionCollection []PartnershipCriterion
type PartnershipCriterionStringMap map[string]PartnershipCriterion
type PartnershipCriterionIntMap map[int]PartnershipCriterion
type PartnershipCriterionUIntMap map[uint]PartnershipCriterion

// A criterion that a channel must meet in order to apply for partnership.
type PartnershipCriterion struct {
	TimeStamped
	Comparison  string `json:"comparison"`
	Description string `json:"description"`
	ID          int    `json:"id"`
	IsEnabled   bool   `json:"isEnabled"`
	Name        string `json:"name"`
	Target      int    `json:"target"`
	Type        string `json:"type"`
}
