package mixer

type PartnershipEvaluationResultCollection []PartnershipEvaluationResult
type PartnershipEvaluationResultStringMap map[string]PartnershipEvaluationResult
type PartnershipEvaluationResultIntMap map[int]PartnershipEvaluationResult
type PartnershipEvaluationResultUIntMap map[uint]PartnershipEvaluationResult

// The result of evaluating whether or not a channel is eligible for partnership.
type PartnershipEvaluationResult struct {
	AccountAge             PartnershipCriterionResult `json:"accountAge"`
	DaysStreamedLastMonth  PartnershipCriterionResult `json:"daysStreamedLastMonth"`
	Followers              PartnershipCriterionResult `json:"followers"`
	HoursStreamedLastMonth PartnershipCriterionResult `json:"hoursStreamedLastMonth"`
	Succeeded              bool                       `json:"succeeded"`
	TotalViewers           PartnershipCriterionResult `json:"totalViewers"`
}
