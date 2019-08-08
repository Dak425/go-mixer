package mixerstructs

type StreamViewerCollection []StreamViewer
type StreamViewerStringMap map[string]StreamViewer
type StreamViewerIntMap map[int]StreamViewer
type StreamViewerUIntMap map[uint]StreamViewer

// A stream viewer represents a record for every time a viewer watches a stream for any period of time
type StreamViewer struct {
	BroadcastId       string `json:"broadcastId"`
	Browser           string `json:"browser"`
	ChannelId         string `json:"channelId"`
	CostreamDepth     string `json:"costreamDepth"`
	Country           string `json:"country"`
	DurationInSeconds int    `json:"durationInSeconds"`
	Id                string `json:"id"`
	Partnered         bool   `json:"partnered"`
	Platform          string `json:"platform"`
	StartedAt         string `json:"startedAt"`
	TypeId            string `json:"typeId"`
	UserId            string `json:"userId"`
}
