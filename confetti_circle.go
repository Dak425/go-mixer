package mixer

type ConfettiCircleCollection []ConfettiCircle
type ConfettiCircleStringMap map[string]ConfettiCircle
type ConfettiCircleIntMap map[int]ConfettiCircle
type ConfettiCircleUIntMap map[uint]ConfettiCircle

// A definition for a circle within a Confetti configuration.
type ConfettiCircle struct {
	Colors     []struct{} `json:"colors"`
	FlipPeriod string     `json:"flipPeriod"`
	Shape      string     `json:"shape"`
	Size       string     `json:"size"`
}
