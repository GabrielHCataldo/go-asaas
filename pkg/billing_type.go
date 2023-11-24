package asaas

type BillingType string

const (
	BOLETO      BillingType = "BOLETO"
	CREDIT_CARD BillingType = "CREDIT_CARD"
	UNDEFINED   BillingType = "UNDEFINED"
	DEBIT_CARD  BillingType = "DEBIT_CARD"
	TRANSFER    BillingType = "TRANSFER"
	DEPOSIT     BillingType = "DEPOSIT"
	PIX         BillingType = "PIX"
)

func (c BillingType) IsEnumValid() bool {
	switch c {
	case BOLETO, CREDIT_CARD, UNDEFINED, DEBIT_CARD, TRANSFER, DEPOSIT, PIX:
		return true
	}
	return false
}
