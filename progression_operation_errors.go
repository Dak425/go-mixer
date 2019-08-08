package mixerstructs

type ProgressionOperationErrorsCollection []ProgressionOperationErrors
type ProgressionOperationErrorsStringMap map[string]ProgressionOperationErrors
type ProgressionOperationErrorsIntMap map[int]ProgressionOperationErrors
type ProgressionOperationErrorsUIntMap map[uint]ProgressionOperationErrors

// A type that contains information about failures in a partial success operation.
type ProgressionOperationErrors struct {
	Errors ProgressionOperationError `json:"errors"`
}
