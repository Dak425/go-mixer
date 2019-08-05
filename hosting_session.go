package mixer

type HostingSessionCollection []HostingSession
type HostingSessionStringMap map[string]HostingSession
type HostingSessionIntMap map[int]HostingSession
type HostingSessionUIntMap map[uint]HostingSession

// Base game type.
type HostingSession struct {
	Hostees uint         `json:"hostees"`
	Layout  PlayerLayout `json:"layout"`
}
