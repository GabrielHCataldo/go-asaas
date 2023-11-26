package asaas

type Pageable[T any] struct {
	HasMore    bool `json:"hasMore"`
	TotalCount int  `json:"totalCount"`
	Limit      int  `json:"limit"`
	Offset     int  `json:"offset"`
	Data       []T  `json:"data"`
}
