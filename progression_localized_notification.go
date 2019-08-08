package mixerstructs

type ProgressionLocalizedNotificationCollection []ProgressionLocalizedNotification
type ProgressionLocalizedNotificationStringMap map[string]ProgressionLocalizedNotification
type ProgressionLocalizedNotificationIntMap map[int]ProgressionLocalizedNotification
type ProgressionLocalizedNotificationUIntMap map[uint]ProgressionLocalizedNotification

// A type that contains information about an operation to run against a player or channel viewers.
type ProgressionLocalizedNotification struct {
	AwardMessage                   string `json:"awardMessage"`
	AwardMessageDetails            string `json:"awardMessageDetails"`
	Id                             string `json:"id"`
	Language                       string `json:"language"`
	MaxDisplayLengthInMilliseconds uint   `json:"maxDisplayLengthInMilliseconds"`
	MinDisplayLengthInMilliseconds uint   `json:"minDisplayLengthInMilliseconds"`
	Name                           string `json:"name"`
	Type                           string `json:"type"`
}
