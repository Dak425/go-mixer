package mixer

type InvoiceItemCollection []InvoiceItem
type InvoiceItemStringMap map[string]InvoiceItem
type InvoiceItemIntMap map[int]InvoiceItem
type InvoiceItemUIntMap map[uint]InvoiceItem

// An item on a user's invoice.
type InvoiceItem struct {
	TimeStamped
	Amount      int    `json:"amount"`
	Description string `json:"description"`
	ID          uint   `json:"id"`
	Invoice     uint   `json:"invoice"`
	Quantity    uint   `json:"quantity"`
	Relid       uint   `json:"relid"`
	Status      string `json:"status"`
	Type        string `json:"type"`
	User        uint   `json:"user"`
}
