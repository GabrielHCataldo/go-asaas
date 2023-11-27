package asaas

type SplitRequest struct {
	WalletID        string  `json:"walletId,omitempty" validate:"required"`
	FixedValue      float64 `json:"fixedValue,omitempty" validate:"omitempty,gt=0"`
	PercentualValue float64 `json:"percentualValue,omitempty" validate:"omitempty,gt=0"`
	TotalFixedValue float64 `json:"totalFixedValue,omitempty" validate:"omitempty,gt=0"`
}

type SplitResponse struct {
	ID              string             `json:"id,omitempty"`
	WalletID        string             `json:"walletId,omitempty"`
	FixedValue      float64            `json:"fixedValue,omitempty"`
	PercentualValue float64            `json:"percentualValue,omitempty"`
	TotalValue      float64            `json:"totalValue,omitempty"`
	RefusalReason   SplitRefusalReason `json:"refusalReason,omitempty"`
	Status          SplitStatus        `json:"status,omitempty"`
}
