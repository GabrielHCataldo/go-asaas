package asaas

type RefundRequest struct {
	// Valor a ser estornado. Caso não informado sera utilizado o valor integral da cobrança
	Value float64 `json:"value,omitempty"`
	// Motivo do estorno
	Description string `json:"description,omitempty"`
}

type RefundResponse struct {
	Status                RefundStatus `json:"status,omitempty"`
	Value                 float64      `json:"value,omitempty"`
	Description           string       `json:"description,omitempty"`
	TransactionReceiptUrl string       `json:"transactionReceiptUrl,omitempty"`
	DateCreated           Date         `json:"dateCreated,omitempty"`
}
