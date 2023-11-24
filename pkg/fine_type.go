package asaas

type FineType string

const (
	FINE_FIXED      FineType = "FIXED"
	FINE_PERCENTAGE FineType = "PERCENTAGE"
)

func (f FineType) IsEnumValid() bool {
	switch f {
	case FINE_FIXED, FINE_PERCENTAGE:
		return true
	}
	return false
}
