package mixerstructs

type ProgressionOperationErrorCollection []ProgressionOperationError
type ProgressionOperationErrorStringMap map[string]ProgressionOperationError
type ProgressionOperationErrorIntMap map[int]ProgressionOperationError
type ProgressionOperationErrorUIntMap map[uint]ProgressionOperationError

// A type that contains information a viewer session and failed operation.
type ProgressionOperationError struct {
	Error     string `json:"error"`
	SessionId string `json:"sessionId"`
}
