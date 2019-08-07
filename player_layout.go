package structs

type PlayerLayoutCollection []PlayerLayout
type PlayerLayoutStringMap map[string]PlayerLayout
type PlayerLayoutIntMap map[int]PlayerLayout
type PlayerLayoutUIntMap map[uint]PlayerLayout

type PlayerLayoutSettingsCollection []PlayerLayoutSettings

type PlayerLayoutSettings struct {
	ChannelID         int     `json:"channelId"`
	LowLatency        bool    `json:"lowLatency"`
	InteractivePinned bool    `json:"interactivePinned"`
	Volume            float64 `json:"volumne"`
	Left              float64 `json:"left"`
	Top               float64 `json:"top"`
	Width             float64 `json:"width"`
	Height            float64 `json:"height"`
}

// The PlayerLayout configures the arrangement and positioning of the channels within a costream or hosting session.
type PlayerLayout struct {
	Name    string                         `json:"name"`
	Order   []int                          `json:"order"` // Deprecated
	Players PlayerLayoutSettingsCollection `json:"players"`
}
