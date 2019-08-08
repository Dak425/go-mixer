package mixerstructs

type PartnershipCriterionResultCollection []PartnershipCriterionResult
type PartnershipCriterionResultStringMap map[string]PartnershipCriterionResult
type PartnershipCriterionResultIntMap map[int]PartnershipCriterionResult
type PartnershipCriterionResultUIntMap map[uint]PartnershipCriterionResult

// The result of evaluating a partnership criterion.
type PartnershipCriterionResult struct {
	Actual      int    `json:"actual"`
	Comparison  string `json:"comparison"`
	CriterionID int    `json:"criterionId"`
	Succeeded   bool   `json:"succeeded"`
	Target      int    `json:"target"`
	Type        string `json:"type"`
}
