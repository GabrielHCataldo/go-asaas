package asaas

type SplitRequest struct {
	WalletId        string  `json:"walletId,omitempty" validate:"required"`
	FixedValue      float64 `json:"fixedValue,omitempty" validate:"omitempty,gt=0"`
	PercentualValue float64 `json:"percentualValue,omitempty" validate:"omitempty,gt=0"`
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
