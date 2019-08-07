package structs

type ConfettiCollection []Confetti
type ConfettiStringMap map[string]Confetti
type ConfettiIntMap map[int]Confetti
type ConfettiUIntMap map[uint]Confetti

// Represents a confetti cannon configuration. Each channel can have is own customized confetti. The confetti cannon is triggered on subscription or direct purchase events on a channel.
type Confetti struct {
	ChannelId uint     `json:"channelId"`
	Id        uint     `json:"id"`
	Settings  struct{} `json:"settings"`
}
