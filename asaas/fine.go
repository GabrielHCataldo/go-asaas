package asaas

type FineRequest struct {
	Value            float64  `json:"value,omitempty" validate:"required,gt=0"`
	DueDateLimitDays int      `json:"dueDateLimitDays,omitempty" validate:"omitempty,gte=0"`
	Type             FineType `json:"type,omitempty" validate:"omitempty,enum"`
}

type FineResponse struct {
	Value float64  `json:"value,omitempty"`
	Type  FineType `json:"type,omitempty"`
}
