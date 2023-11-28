package asaas

type RefundRequest struct {
	Value       float64 `json:"value,omitempty"`
	Description string  `json:"description,omitempty"`
}

type RefundResponse struct {
	Status                RefundStatus `json:"status,omitempty"`
	Value                 float64      `json:"value,omitempty"`
	Description           string       `json:"description,omitempty"`
	TransactionReceiptUrl string       `json:"transactionReceiptUrl,omitempty"`
	DateCreated           Date         `json:"dateCreated,omitempty"`
}
