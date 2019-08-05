package mixer

type InteractiveScreenAnalysisCollection []InteractiveScreenAnalysis
type InteractiveScreenAnalysisStringMap map[string]InteractiveScreenAnalysis
type InteractiveScreenAnalysisIntMap map[int]InteractiveScreenAnalysis
type InteractiveScreenAnalysisUIntMap map[uint]InteractiveScreenAnalysis

// Analysis settings for a Screen control. If an analysis type is enabled results will be generated and sent to the Interactive Robot.
type InteractiveScreenAnalysis struct {
	Clicks   bool                        `json:"clicks"`
	Position InteractiveAnalyticSettings `json:"position"`
}
