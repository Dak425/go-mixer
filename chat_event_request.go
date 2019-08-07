package structs

type ChatEventRequestCollection []ChatEventRequest
type ChatEventRequestStringMap map[string]ChatEventRequest
type ChatEventRequestIntMap map[int]ChatEventRequest
type ChatEventRequestUIntMap map[uint]ChatEventRequest

// Data required to create a Chat event.
type ChatEventRequest struct {
	Data struct{} `json:"data"`
	Type string   `json:"type"`
}
