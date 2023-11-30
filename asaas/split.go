package asaas

type SplitRequest struct {
	// Identificador da carteira (retornado no momento da criação da conta) (REQUIRED)
	WalletId string `json:"walletId,omitempty" validate:"required"`
	// Valor fixo a ser transferido para a conta quando a cobrança for recebida
	FixedValue float64 `json:"fixedValue,omitempty" validate:"omitempty,gt=0"`
	// Percentual sobre o valor líquido da cobrança a ser transferido quando for recebida
	PercentualValue float64 `json:"percentualValue,omitempty" validate:"omitempty,gt=0"`
	// (Somente parcelamentos). Valor que será feito split referente ao valor total que será parcelado.
	TotalFixedValue float64 `json:"totalFixedValue,omitempty" validate:"omitempty,gt=0"`
}

type SplitResponse struct {
	Id              string             `json:"id,omitempty"`
	WalletId        string             `json:"walletId,omitempty"`
	Status          SplitStatus        `json:"status,omitempty"`
	FixedValue      float64            `json:"fixedValue,omitempty"`
	PercentualValue float64            `json:"percentualValue,omitempty"`
	TotalValue      float64            `json:"totalValue,omitempty"`
	RefusalReason   SplitRefusalReason `json:"refusalReason,omitempty"`
}
