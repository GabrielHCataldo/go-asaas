package asaas

type FineRequest struct {
	// Percentual de multa sobre o valor da cobrança para pagamento após o vencimento
	Value float64  `json:"value,omitempty" validate:"required,gt=0"`
	Type  FineType `json:"type,omitempty" validate:"omitempty,enum"`
}

type FineResponse struct {
	Value float64  `json:"value,omitempty"`
	Type  FineType `json:"type,omitempty"`
}
