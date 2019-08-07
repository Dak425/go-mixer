package structs

type ViewerCountCollection []ViewerCount
type ViewerCountStringMap map[string]ViewerCount
type ViewerCountIntMap map[int]ViewerCount
type ViewerCountUIntMap map[uint]ViewerCount

// A channel metric showing how many viewers and of what type were watching the channel at a set point in time.
type ViewerCount struct {
	Anon   uint   `json:"anon"`
	Authed uint   `json:"authed"`
	Time   string `json:"time"`
}
