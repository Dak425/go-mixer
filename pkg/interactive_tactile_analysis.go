package mixer

type InteractiveTactileAnalysisCollection []InteractiveTactileAnalysis
type InteractiveTactileAnalysisStringMap map[string]InteractiveTactileAnalysis
type InteractiveTactileAnalysisIntMap map[int]InteractiveTactileAnalysis
type InteractiveTactileAnalysisUIntMap map[uint]InteractiveTactileAnalysis

// Analysis settings for a Tactile control. If an analysis type is enabled results will be generated and sent to the Interactive Robot.
type InteractiveTactileAnalysis struct {
	Frequency bool `json:"frequency"`
	Holding   bool `json:"holding"`
}
