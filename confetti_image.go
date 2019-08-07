package structs

type ConfettiImageCollection []ConfettiImage
type ConfettiImageStringMap map[string]ConfettiImage
type ConfettiImageIntMap map[int]ConfettiImage
type ConfettiImageUIntMap map[uint]ConfettiImage

// A definition for an image within a Confetti configuration.
type ConfettiImage struct {
	Data       string `json:"data"`
	Scale      string `json:"scale"`
	Shape      string `json:"shape"`
	SpinPeriod string `json:"spinPeriod"`
}
