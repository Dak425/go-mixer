package structs

type InteractiveJoyStickAnalysisCollection []InteractiveJoyStickAnalysis
type InteractiveJoyStickAnalysisStringMap map[string]InteractiveJoyStickAnalysis
type InteractiveJoyStickAnalysisIntMap map[int]InteractiveJoyStickAnalysis
type InteractiveJoyStickAnalysisUIntMap map[uint]InteractiveJoyStickAnalysis

// Analysis settings for a Joystick. If an analysis type is enabled results will be generated and sent to the Interactive Robot.
type InteractiveJoyStickAnalysis struct {
	Coords InteractiveAnalyticSettings `json:"coords"`
}
