package mixerstructs

type ClipErrorCollection []ClipError
type ClipErrorStringMap map[string]ClipError
type ClipErrorIntMap map[int]ClipError
type ClipErrorUIntMap map[uint]ClipError

// Error information for an unsuccesful clip request
type ClipError struct {
	ErrorCode    uint   `json:"errorCode"`
	ErrorMessage string `json:"errorMessage"`
}
