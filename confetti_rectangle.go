package structs

type ConfettiRectangleCollection []ConfettiRectangle
type ConfettiRectangleStringMap map[string]ConfettiRectangle
type ConfettiRectangleIntMap map[int]ConfettiRectangle
type ConfettiRectangleUIntMap map[uint]ConfettiRectangle

// A definition for a rectangle within a Confetti configuration.
type ConfettiRectangle struct {
	Colors     []struct{} `json:"colors"`
	FlipPeriod string     `json:"flipPeriod"`
	Height     string     `json:"height"`
	Shape      string     `json:"shape"`
	Width      string     `json:"width"`
}
