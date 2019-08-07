package structs

type NotificationCollection []Notification
type NotificationStringMap map[string]Notification
type NotificationIntMap map[int]Notification
type NotificationUIntMap map[uint]Notification

// A single notification addressed to a user.
type Notification struct {
	Payload struct{} `json:"payload"`
	SentAt  string   `json:"sentAt"`
	Trigger string   `json:"trigger"`
	UserId  uint     `json:"userId"`
}
