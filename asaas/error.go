package asaas

type ErrorResponse struct {
	Code        string `json:"code,omitempty"`
	Description string `json:"description,omitempty"`
}
