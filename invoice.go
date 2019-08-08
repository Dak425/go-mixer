package mixerstructs

type InvoiceCollection []Invoice
type InvoiceStringMap map[string]Invoice
type InvoiceIntMap map[int]Invoice
type InvoiceUIntMap map[uint]Invoice

// A user invoice.
type Invoice struct {
	Amount    int         `json:"amount"`
	CreatedAt string      `json:"createdAt"`
	Currency  string      `json:"currency"`
	ID        uint        `json:"id"`
	Items     InvoiceItem `json:"items"`
	Status    string      `json:"status"`
	User      uint        `json:"user"`
}
