package asaas

type InterestRequest struct {
	// Percentual de juros ao mês sobre o valor da cobrança para pagamento após o vencimento
	Value float64 `json:"value,omitempty" validate:"required,gt=0"`
}

type InterestResponse struct {
	Value float64      `json:"value,omitempty"`
	Type  InterestType `json:"type,omitempty"`
}
