package structs

type DelveOnlyOnMixerCollection []DelveOnlyOnMixer
type DelveOnlyOnMixerStringMap map[string]DelveOnlyOnMixer
type DelveOnlyOnMixerIntMap map[int]DelveOnlyOnMixer
type DelveOnlyOnMixerUIntMap map[uint]DelveOnlyOnMixer

// Returned when we tried to get some contents for this row, but encountered an error doing so.
type DelveOnlyOnMixer struct {
	Channel Channel `json:"channel"`
	Id      int     `json:"id"`
	Type    string  `json:"type"`
	TypeId  int     `json:"typeId"`
}
