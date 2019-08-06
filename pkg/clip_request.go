package mixer

type ClipRequestCollection []ClipRequest
type ClipRequestStringMap map[string]ClipRequest
type ClipRequestIntMap map[int]ClipRequest
type ClipRequestUIntMap map[uint]ClipRequest

// Information required to create a clip.
type ClipRequest struct {
	BroadcastId           string `json:"broadcastId"`
	ClipDurationInSeconds int    `json:"clipDurationInSeconds"`
	FinalStreamTimeMs     int    `json:"finalStreamTimeMs"`
	HighlightTitle        string `json:"highlightTitle"`
}
