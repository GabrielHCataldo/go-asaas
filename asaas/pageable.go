package asaas

type Pageable[T any] struct {
	HasMore    bool            `json:"hasMore"`
	TotalCount int             `json:"totalCount"`
	Limit      int             `json:"limit"`
	Offset     int             `json:"offset"`
	Data       []T             `json:"data"`
	Errors     []ErrorResponse `json:"errors,omitempty"`
}

type PageableDefaultRequest struct {
	// Elemento inicial da lista
	Offset int `json:"offset,omitempty"`
	// Número de elementos da lista (max: 100)
	Limit int `json:"limit,omitempty"`
}
