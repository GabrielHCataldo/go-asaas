package asaas

type InterestRequest struct {
	// Percentual de juros ao mês sobre o valor da cobrança para pagamento após o vencimento
	Value float64 `json:"value,omitempty"`
}

type InterestResponse struct {
	Value float64 `json:"value,omitempty"`
}
