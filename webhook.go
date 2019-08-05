package mixer

type WebhookCollection []Webhook
type WebhookStringMap map[string]Webhook
type WebhookIntMap map[int]Webhook
type WebhookUIntMap map[uint]Webhook

type DeactivationReason struct {
	Code   uint   `json:"code"`
	Reason string `json:"reason"`
}

// A registered webhook. For more information about webhooks, see aka.ms/mixrpokr.
type Webhook struct {
	DeactivationReason DeactivationReason `json:"deactivationReason"`
	Events             []string           `json:"events"`
	ExpiresAt          string             `json:"expiresAt"`
	Id                 string             `json:"id"`
	IsActive           bool               `json:"isActive"`
	Kind               string             `json:"kind"`
	RenewUrl           string             `json:"renewUrl"`
	Url                string             `json:"url"`
}
