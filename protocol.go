package mixer

type ProtocolCollection []Protocol
type ProtocolStringMap map[string]Protocol
type ProtocolIntMap map[int]Protocol
type ProtocolUIntMap map[uint]Protocol

type Protocol struct {
	Type string `json:"type"`
}
