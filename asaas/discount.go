package asaas

type DiscountRequest struct {
	Value            float64      `json:"value,omitempty" validate:"required,gt=0"`
	DueDateLimitDays int          `json:"dueDateLimitDays,omitempty" validate:"gte=0"`
	Type             DiscountType `json:"type,omitempty" validate:"omitempty,enum"`
}

type DiscountResponse struct {
	Value            float64      `json:"value,omitempty"`
	DueDateLimitDays int          `json:"dueDateLimitDays,omitempty"`
	Type             DiscountType `json:"type,omitempty"`
}
