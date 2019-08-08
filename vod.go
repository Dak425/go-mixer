package mixerstructs

type VODCollection []VOD
type VODStringMap map[string]VOD
type VODIntMap map[int]VOD
type VODUIntMap map[uint]VOD

// A VOD is a recorded broadcast which can be watched again on Mixer.
type VOD struct {
	TimeStamped
	BaseUrl     string   `json:"baseUrl"`
	Data        struct{} `json:"data"`
	Format      string   `json:"format"`
	Id          uint     `json:"id"`
	RecordingId uint     `json:"recordingId"`
}
