package structs

type IngestCollection []Ingest
type IngestStringMap map[string]Ingest
type IngestIntMap map[int]Ingest
type IngestUIntMap map[uint]Ingest

// An ingest definition. Ingests are used by streaming software to send video to Mixer.
type Ingest struct {
	Host      string             `json:"host"`
	Name      string             `json:"name"`
	PingTest  string             `json:"pingTest"`
	Protocols ProtocolCollection `json:"protocols"`
}
