package asaas

type Pageable[T any] struct {
	HasMore    bool `json:"hasMore"`
	TotalCount int  `json:"totalCount"`
	Limit      int  `json:"limit"`
	Offset     int  `json:"offset"`
	Data       []T  `json:"data"`
}

type PageableDefaultRequest struct {
	Offset int `json:"offset,omitempty"`
	Limit  int `json:"limit,omitempty" validate:"omitempty,lte=100"`
}
