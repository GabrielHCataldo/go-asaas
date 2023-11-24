package asaas

type ChargebackStatus string

const (
	CHARGEBACK_REQUESTED  ChargebackStatus = "REQUESTED"
	CHARGEBACK_IN_DISPUTE ChargebackStatus = "IN_DISPUTE"
	CHARGEBACK_LOST       ChargebackStatus = "LOST"
	CHARGEBACK_REVERSED   ChargebackStatus = "REVERSED"
	CHARGEBACK_DONE       ChargebackStatus = "DONE"
)

func (c ChargebackStatus) IsEnumValid() bool {
	switch c {
	case CHARGEBACK_REQUESTED, CHARGEBACK_IN_DISPUTE, CHARGEBACK_LOST, CHARGEBACK_REVERSED, CHARGEBACK_DONE:
		return true
	}
	return false
}
