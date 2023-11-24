package asaas

type DiscountType string

const (
	DISCOUNT_FIXED      DiscountType = "FIXED"
	DISCOUNT_PERCENTAGE DiscountType = "PERCENTAGE"
)

func (d DiscountType) IsEnumValid() bool {
	switch d {
	case DISCOUNT_FIXED, DISCOUNT_PERCENTAGE:
		return true
	}
	return false
}
