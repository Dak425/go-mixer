package mixerstructs

type InteractiveControlsCollection []InteractiveControls
type InteractiveControlsStringMap map[string]InteractiveControls
type InteractiveControlsIntMap map[int]InteractiveControls
type InteractiveControlsUIntMap map[uint]InteractiveControls

// Specifies the Interactive V1 Controls that an InteractiveVersion has.
type InteractiveControls struct {
	Joysticks      struct{} `json:"joysticks"`
	ReportInterval uint     `json:"reportInterval"`
	Screens        struct{} `json:"screens"`
	Tactiles       struct{} `json:"tactiles"`
}
