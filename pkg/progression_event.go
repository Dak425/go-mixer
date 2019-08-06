package mixer

type ProgressionEventCollection []ProgressionEvent
type ProgressionEventStringMap map[string]ProgressionEvent
type ProgressionEventIntMap map[int]ProgressionEvent
type ProgressionEventUIntMap map[uint]ProgressionEvent

// A type that contains information about a notification event to be sent for progression storage changes.
type ProgressionEvent struct {
	AwardMessages                  LocaleAndString `json:"awardMessages"`
	AwardMessagesDetails           LocaleAndString `json:"awardMessagesDetails"`
	Id                             string          `json:"id"`
	MaxDisplayLengthInMilliseconds uint            `json:"maxDisplayLengthInMilliseconds"`
	MinDisplayLengthInMilliseconds uint            `json:"minDisplayLengthInMilliseconds"`
	Name                           string          `json:"name"`
	OwningProject                  string          `json:"owningProject"`
	Type                           string          `json:"type"`
}
