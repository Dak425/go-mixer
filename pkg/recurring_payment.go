package mixer

type RecurringPaymentCollection []RecurringPayment
type RecurringPaymentStringMap map[string]RecurringPayment
type RecurringPaymentIntMap map[int]RecurringPayment
type RecurringPaymentUIntMap map[uint]RecurringPayment

// A recurring payment is a payment which will continue to be charged to the user every period until it is canceled. These are used for subscriptions and Mixer Pro.
type RecurringPayment struct {
	TimeStamped
	Cancelled bool   `json:"cancelled"`
	Gateway   string `json:"gateway"`
	Id        uint   `json:"id"`
	Relid     uint   `json:"relid"`
	Status    string `json:"status"`
	TimesPaid uint   `json:"timesPaid"`
	Type      string `json:"type"`
	User      uint   `json:"user"`
}
