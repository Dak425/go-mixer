package mixerstructs

type ProgressionOperationsCollection []ProgressionOperations
type ProgressionOperationsStringMap map[string]ProgressionOperations
type ProgressionOperationsIntMap map[int]ProgressionOperations
type ProgressionOperationsUIntMap map[uint]ProgressionOperations

// A type that contains information about progression operations to run against a player or channel viewers.
type ProgressionOperations struct {
	Operations ProgressionOperation `json:"operations"`
}
