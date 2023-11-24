package asaas

type RefundStatus string

const (
	REFUND_PENDING   RefundStatus = "PENDING"
	REFUND_CANCELLED RefundStatus = "CANCELLED"
	REFUND_DONE      RefundStatus = "DONE"
)

func (r RefundStatus) IsEnumValid() bool {
	switch r {
	case REFUND_PENDING, REFUND_CANCELLED, REFUND_DONE:
		return true
	}
	return false
}
