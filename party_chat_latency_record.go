package structs

type PartyChatLatencyRecordCollection []PartyChatLatencyRecord
type PartyChatLatencyRecordStringMap map[string]PartyChatLatencyRecord
type PartyChatLatencyRecordIntMap map[int]PartyChatLatencyRecord
type PartyChatLatencyRecordUIntMap map[uint]PartyChatLatencyRecord

// Results of a latency measurement to a server.
type PartyChatLatencyRecord struct {
	ServerFqdn     string `json:"serverFqdn"`
	TargetLocation string `json:"targetLocation"`
}
