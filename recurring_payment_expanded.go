package mixerstructs

type RecurringPaymentExpandedCollection []RecurringPaymentExpanded
type RecurringPaymentExpandedStringMap map[string]RecurringPaymentExpanded
type RecurringPaymentExpandedIntMap map[int]RecurringPaymentExpanded
type RecurringPaymentExpandedUIntMap map[uint]RecurringPaymentExpanded

// An augmented recurring payment with additional details, used in Details views. Provides nested objects such as the owner of the RecurringPayment.
type RecurringPaymentExpanded struct {
	RecurringPayment
	Subscription SubscriptionExpanded `json:"subscription"`
}
