package asaas

type SplitStatus string

const (
	SPLIT_PENDING         SplitStatus = "PENDING"
	SPLIT_AWAITING_CREDIT SplitStatus = "AWAITING_CREDIT"
	SPLIT_CANCELLED       SplitStatus = "CANCELLED"
	SPLIT_DONE            SplitStatus = "DONE"
	SPLIT_REFUSED         SplitStatus = "REFUSED"
)

func (s SplitStatus) IsEnumValid() bool {
	switch s {
	case SPLIT_PENDING, SPLIT_AWAITING_CREDIT, SPLIT_CANCELLED, SPLIT_DONE, SPLIT_REFUSED:
		return true
	}
	return false
}
