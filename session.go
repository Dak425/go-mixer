package mixerstructs

type SessionCollection []Session
type SessionStringMap map[string]Session
type SessionIntMap map[int]Session
type SessionUIntMap map[uint]Session

// Provides details about a Mixer user's session.
type Session struct {
	Expires uint     `json:"expires"`
	Id      string   `json:"id"`
	Ip      string   `json:"ip"`
	Meta    struct{} `json:"meta"`
}
