package asaas

type DiscountRequest struct {
	// Valor percentual ou fixo de desconto a ser aplicado sobre o valor da cobrança
	Value float64 `json:"value,omitempty"`
	// Dias antes do vencimento para aplicar desconto. Ex: 0 = até o vencimento, 1 = até um dia antes, 2 = até dois dias antes, e assim por diante
	DueDateLimitDays int          `json:"dueDateLimitDays,omitempty"`
	Type             DiscountType `json:"type,omitempty"`
}

type DiscountResponse struct {
	Value            float64      `json:"value,omitempty"`
	DueDateLimitDays int          `json:"dueDateLimitDays,omitempty"`
	Type             DiscountType `json:"type,omitempty"`
}
