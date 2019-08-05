package mixer

type PartyChatServerCollection []PartyChatServer
type PartyChatServerStringMap map[string]PartyChatServer
type PartyChatServerIntMap map[int]PartyChatServer
type PartyChatServerUIntMap map[uint]PartyChatServer

// A server used to measure location-based latency in party chat.
type PartyChatServer struct {
	ServerFqdn     string `json:"serverFqdn"`
	TargetLocation string `json:"targetLocation"`
}
