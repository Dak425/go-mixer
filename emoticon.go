package mixerstructs

type EmoticonCollection []Emoticon
type EmoticonStringMap map[string]Emoticon
type EmoticonIntMap map[int]Emoticon
type EmoticonUIntMap map[uint]Emoticon

type Emoticon struct {
	X      uint `json:"x"`
	Y      uint `json:"y"`
	Width  uint `json:"width"`
	Height uint `json:"height"`
}
