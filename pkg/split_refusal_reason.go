package asaas

type SplitRefusalReason string

const (
	RECEIVABLE_UNIT_AFFECTED_BY_EXTERNAL_CONTRACTUAL_EFFECT SplitRefusalReason = "RECEIVABLE_UNIT_AFFECTED_BY_EXTERNAL_CONTRACTUAL_EFFECT"
)

func (s SplitRefusalReason) IsEnumValid() bool {
	switch s {
	case RECEIVABLE_UNIT_AFFECTED_BY_EXTERNAL_CONTRACTUAL_EFFECT:
		return true
	}
	return false
}
