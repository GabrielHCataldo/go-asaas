package asaas

type ErrorType string

const (
	ERROR_VALIDATION ErrorType = "VALIDATION"
	ERROR_UNEXPECTED ErrorType = "UNEXPECTED"
)

func (e ErrorType) IsEnumValid() bool {
	switch e {
	case ERROR_VALIDATION, ERROR_UNEXPECTED:
		return true
	}
	return false
}
