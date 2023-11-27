package asaas

type InterestRequest struct {
	Value float64 `json:"value,omitempty" validate:"required,gt=0"`
}

type InterestResponse struct {
	Value float64      `json:"value,omitempty"`
	Type  InterestType `json:"type,omitempty"`
}
