package asaas

type BaseEnum interface {
	IsEnumValid() bool
}

type BillingType string
type ChargeStatus string
type ChargebackReason string
type ChargebackStatus string
type DiscountType string
type Env int
type ErrorType string
type FineType string
type InterestType string
type RefundStatus string
type SplitRefusalReason string
type SplitStatus string
type TypeOfDocument string
type NotificationEvent string
type SubscriptionCycle string
type SubscriptionStatus string
type Order string
type SortSubscriptionField string
type InvoiceStatus string
type InvoiceDatePeriod string
type InvoiceDaysBeforeDueDate int

const (
	INVOICE_DAYS_BEFORE_DUEDATE_FIVE    InvoiceDaysBeforeDueDate = 5
	INVOICE_DAYS_BEFORE_DUEDATE_TEN     InvoiceDaysBeforeDueDate = 10
	INVOICE_DAYS_BEFORE_DUEDATE_FIFTEEN InvoiceDaysBeforeDueDate = 15
	INVOICE_DAYS_BEFORE_DUEDATE_THIRTY  InvoiceDaysBeforeDueDate = 30
	INVOICE_DAYS_BEFORE_DUEDATE_SIXTY   InvoiceDaysBeforeDueDate = 60
)
const (
	INVOICE_DATE_PERIOD_ON_PAYMENT_CONFIRMATION InvoiceDatePeriod = "ON_PAYMENT_CONFIRMATION"
	INVOICE_DATE_PERIOD_ON_PAYMENT_DUE_DATE     InvoiceDatePeriod = "ON_PAYMENT_DUE_DATE"
	INVOICE_DATE_PERIOD_BEFORE_PAYMENT_DUE_DATE InvoiceDatePeriod = "BEFORE_PAYMENT_DUE_DATE"
	INVOICE_DATE_PERIOD_ON_DUE_DATE_MONTH       InvoiceDatePeriod = "ON_DUE_DATE_MONTH"
	INVOICE_DATE_PERIOD_ON_NEXT_MONTH           InvoiceDatePeriod = "ON_NEXT_MONTH"
)
const (
	INVOICE_SCHEDULED               InvoiceStatus = "SCHEDULED"
	INVOICE_SYNCHRONIZED            InvoiceStatus = "SYNCHRONIZED"
	INVOICE_AUTHORIZED              InvoiceStatus = "AUTHORIZED"
	INVOICE_PROCESSING_CANCELLATION InvoiceStatus = "PROCESSING_CANCELLATION"
	INVOICE_CANCELED                InvoiceStatus = "CANCELED"
	INVOICE_CANCELLATION_DENIED     InvoiceStatus = "CANCELLATION_DENIED"
	INVOICE_ERROR                   InvoiceStatus = "ERROR"
)
const (
	SORT_SUBSCRIPTION_DATE_CREATED SortSubscriptionField = "dateCreated"
)
const (
	ORDER_DESC Order = "desc"
	ORDER_ASC  Order = "asc"
)
const (
	SUBSCRIPTION_ACTIVE   SubscriptionStatus = "ACTIVE"
	SUBSCRIPTION_INACTIVE SubscriptionStatus = "INACTIVE"
	SUBSCRIPTION_EXPIRED  SubscriptionStatus = "EXPIRED"
)
const (
	WEEKLY       SubscriptionCycle = "WEEKLY"
	BIWEEKLY     SubscriptionCycle = "BIWEEKLY"
	MONTHLY      SubscriptionCycle = "MONTHLY"
	BIMONTHLY    SubscriptionCycle = "BIMONTHLY"
	QUARTERLY    SubscriptionCycle = "QUARTERLY"
	SEMIANNUALLY SubscriptionCycle = "SEMIANNUALLY"
	YEARLY       SubscriptionCycle = "YEARLY"
)
const (
	BOLETO      BillingType = "BOLETO"
	CREDIT_CARD BillingType = "CREDIT_CARD"
	UNDEFINED   BillingType = "UNDEFINED"
	DEBIT_CARD  BillingType = "DEBIT_CARD"
	TRANSFER    BillingType = "TRANSFER"
	DEPOSIT     BillingType = "DEPOSIT"
	PIX         BillingType = "PIX"
)
const (
	CHARGE_PENDING                      ChargeStatus = "PENDING"
	CHARGE_RECEIVED                     ChargeStatus = "RECEIVED"
	CHARGE_CONFIRMED                    ChargeStatus = "CONFIRMED"
	CHARGE_OVERDUE                      ChargeStatus = "OVERDUE"
	CHARGE_REFUNDED                     ChargeStatus = "REFUNDED"
	CHARGE_RECEIVED_IN_CASH             ChargeStatus = "RECEIVED_IN_CASH"
	CHARGE_REFUND_REQUESTED             ChargeStatus = "REFUND_REQUESTED"
	CHARGE_REFUND_IN_PROGRESS           ChargeStatus = "REFUND_IN_PROGRESS"
	CHARGE_CHARGEBACK_REQUESTED         ChargeStatus = "CHARGEBACK_REQUESTED"
	CHARGE_CHARGEBACK_DISPUTE           ChargeStatus = "CHARGEBACK_DISPUTE"
	CHARGE_AWAITING_CHARGEBACK_REVERSAL ChargeStatus = "AWAITING_CHARGEBACK_REVERSAL"
	CHARGE_DUNNING_REQUESTED            ChargeStatus = "DUNNING_REQUESTED"
	CHARGE_DUNNING_RECEIVED             ChargeStatus = "DUNNING_RECEIVED"
	CHARGE_AWAITING_RISK_ANALYSIS       ChargeStatus = "AWAITING_RISK_ANALYSIS"
)
const (
	ABSENCE_OF_PRINT                          ChargebackReason = "ABSENCE_OF_PRINT"
	ABSENT_CARD_FRAUD                         ChargebackReason = "ABSENT_CARD_FRAUD"
	CARD_ACTIVATED_PHONE_TRANSACTION          ChargebackReason = "CARD_ACTIVATED_PHONE_TRANSACTION"
	CARD_FRAUD                                ChargebackReason = "CARD_FRAUD"
	CARD_RECOVERY_BULLETIN                    ChargebackReason = "CARD_RECOVERY_BULLETIN"
	COMMERCIAL_DISAGREEMENT                   ChargebackReason = "COMMERCIAL_DISAGREEMENT"
	COPY_NOT_RECEIVED                         ChargebackReason = "COPY_NOT_RECEIVED"
	CREDIT_OR_DEBIT_PRESENTATION_ERROR        ChargebackReason = "CREDIT_OR_DEBIT_PRESENTATION_ERROR"
	DIFFERENT_PAY_METHOD                      ChargebackReason = "DIFFERENT_PAY_METHOD"
	FRAUD                                     ChargebackReason = "FRAUD"
	INCORRECT_TRANSACTION_VALUE               ChargebackReason = "INCORRECT_TRANSACTION_VALUE"
	INVALID_CURRENCY                          ChargebackReason = "INVALID_CURRENCY"
	INVALID_DATA                              ChargebackReason = "INVALID_DATA"
	LATE_PRESENTATION                         ChargebackReason = "LATE_PRESENTATION"
	LOCAL_REGULATORY_OR_LEGAL_DISPUTE         ChargebackReason = "LOCAL_REGULATORY_OR_LEGAL_DISPUTE"
	MULTIPLE_ROCS                             ChargebackReason = "MULTIPLE_ROCS"
	ORIGINAL_CREDIT_TRANSACTION_NOT_ACCEPTED  ChargebackReason = "ORIGINAL_CREDIT_TRANSACTION_NOT_ACCEPTED"
	OTHER_ABSENT_CARD_FRAUD                   ChargebackReason = "OTHER_ABSENT_CARD_FRAUD"
	PROCESS_ERROR                             ChargebackReason = "PROCESS_ERROR"
	RECEIVED_COPY_ILLEGIBLE_OR_INCOMPLETE     ChargebackReason = "RECEIVED_COPY_ILLEGIBLE_OR_INCOMPLETE"
	RECURRENCE_CANCELED                       ChargebackReason = "RECURRENCE_CANCELED"
	REQUIRED_AUTHORIZATION_NOT_GRANTED        ChargebackReason = "REQUIRED_AUTHORIZATION_NOT_GRANTED"
	RIGHT_OF_FULL_RECOURSE_FOR_FRAUD          ChargebackReason = "RIGHT_OF_FULL_RECOURSE_FOR_FRAUD"
	SALE_CANCELED                             ChargebackReason = "SALE_CANCELED"
	SERVICE_DISAGREEMENT_OR_DEFECTIVE_PRODUCT ChargebackReason = "SERVICE_DISAGREEMENT_OR_DEFECTIVE_PRODUCT"
	SERVICE_NOT_RECEIVED                      ChargebackReason = "SERVICE_NOT_RECEIVED"
	SPLIT_SALE                                ChargebackReason = "SPLIT_SALE"
	TRANSFERS_OF_DIVERSE_RESPONSIBILITIES     ChargebackReason = "TRANSFERS_OF_DIVERSE_RESPONSIBILITIES"
	UNQUALIFIED_CAR_RENTAL_DEBIT              ChargebackReason = "UNQUALIFIED_CAR_RENTAL_DEBIT"
	USA_CARDHOLDER_DISPUTE                    ChargebackReason = "USA_CARDHOLDER_DISPUTE"
	VISA_FRAUD_MONITORING_PROGRAM             ChargebackReason = "VISA_FRAUD_MONITORING_PROGRAM"
	WARNING_BULLETIN_FILE                     ChargebackReason = "WARNING_BULLETIN_FILE"
)
const (
	CHARGEBACK_REQUESTED  ChargebackStatus = "REQUESTED"
	CHARGEBACK_IN_DISPUTE ChargebackStatus = "IN_DISPUTE"
	CHARGEBACK_LOST       ChargebackStatus = "LOST"
	CHARGEBACK_REVERSED   ChargebackStatus = "REVERSED"
	CHARGEBACK_DONE       ChargebackStatus = "DONE"
)
const (
	DISCOUNT_FIXED      DiscountType = "FIXED"
	DISCOUNT_PERCENTAGE DiscountType = "PERCENTAGE"
)
const (
	SANDBOX    Env = iota
	PRODUCTION Env = iota
)
const (
	ERROR_VALIDATION ErrorType = "VALIDATION"
	ERROR_UNEXPECTED ErrorType = "UNEXPECTED"
)
const (
	FINE_FIXED      FineType = "FIXED"
	FINE_PERCENTAGE FineType = "PERCENTAGE"
)
const (
	INTEREST_FIXED      InterestType = "FIXED"
	INTEREST_PERCENTAGE InterestType = "PERCENTAGE"
)
const (
	REFUND_PENDING   RefundStatus = "PENDING"
	REFUND_CANCELLED RefundStatus = "CANCELLED"
	REFUND_DONE      RefundStatus = "DONE"
)
const (
	RECEIVABLE_UNIT_AFFECTED_BY_EXTERNAL_CONTRACTUAL_EFFECT SplitRefusalReason = "RECEIVABLE_UNIT_AFFECTED_BY_EXTERNAL_CONTRACTUAL_EFFECT"
)
const (
	SPLIT_PENDING         SplitStatus = "PENDING"
	SPLIT_AWAITING_CREDIT SplitStatus = "AWAITING_CREDIT"
	SPLIT_CANCELLED       SplitStatus = "CANCELLED"
	SPLIT_DONE            SplitStatus = "DONE"
	SPLIT_REFUSED         SplitStatus = "REFUSED"
)
const (
	INVOICE     TypeOfDocument = "INVOICE"
	CONTRACT    TypeOfDocument = "CONTRACT"
	DOCUMENT    TypeOfDocument = "DOCUMENT"
	SPREADSHEET TypeOfDocument = "SPREADSHEET"
	PROGRAM     TypeOfDocument = "PROGRAM"
	OTHER       TypeOfDocument = "OTHER"
)
const (
	PAYMENT_CREATED         NotificationEvent = "PAYMENT_CREATED"
	PAYMENT_DUEDATE_WARNING NotificationEvent = "PAYMENT_DUEDATE_WARNING"
	PAYMENT_RECEIVED        NotificationEvent = "PAYMENT_RECEIVED"
	SEND_LINHA_DIGITAVEL    NotificationEvent = "SEND_LINHA_DIGITAVEL"
	PAYMENT_OVERDUE         NotificationEvent = "PAYMENT_OVERDUE"
	PAYMENT_UPDATED         NotificationEvent = "PAYMENT_UPDATED"
)

func (i InvoiceDaysBeforeDueDate) IsEnumValid() bool {
	switch i {
	case INVOICE_DAYS_BEFORE_DUEDATE_FIVE, INVOICE_DAYS_BEFORE_DUEDATE_TEN, INVOICE_DAYS_BEFORE_DUEDATE_FIFTEEN,
		INVOICE_DAYS_BEFORE_DUEDATE_THIRTY, INVOICE_DAYS_BEFORE_DUEDATE_SIXTY:
		return true
	}
	return false
}

func (i InvoiceStatus) IsEnumValid() bool {
	switch i {
	case INVOICE_AUTHORIZED, INVOICE_CANCELED, INVOICE_CANCELLATION_DENIED, INVOICE_PROCESSING_CANCELLATION,
		INVOICE_ERROR, INVOICE_SCHEDULED, INVOICE_SYNCHRONIZED:
		return true
	}
	return false
}

func (i InvoiceDatePeriod) IsEnumValid() bool {
	switch i {
	case INVOICE_DATE_PERIOD_ON_PAYMENT_CONFIRMATION, INVOICE_DATE_PERIOD_ON_PAYMENT_DUE_DATE,
		INVOICE_DATE_PERIOD_BEFORE_PAYMENT_DUE_DATE, INVOICE_DATE_PERIOD_ON_DUE_DATE_MONTH,
		INVOICE_DATE_PERIOD_ON_NEXT_MONTH:
		return true
	}
	return false
}

func (s SortSubscriptionField) IsEnumValid() bool {
	switch s {
	case SORT_SUBSCRIPTION_DATE_CREATED:
		return true
	}
	return false
}

func (s SubscriptionStatus) IsEnumValid() bool {
	switch s {
	case SUBSCRIPTION_ACTIVE, SUBSCRIPTION_INACTIVE, SUBSCRIPTION_EXPIRED:
		return true
	}
	return false
}

func (o Order) IsEnumValid() bool {
	switch o {
	case ORDER_DESC, ORDER_ASC:
		return true
	}
	return false
}

func (s SubscriptionCycle) IsEnumValid() bool {
	switch s {
	case WEEKLY, BIWEEKLY, MONTHLY, BIMONTHLY, QUARTERLY, SEMIANNUALLY, YEARLY:
		return true
	}
	return false
}

func (n NotificationEvent) IsEnumValid() bool {
	switch n {
	case PAYMENT_CREATED, PAYMENT_DUEDATE_WARNING, PAYMENT_RECEIVED, SEND_LINHA_DIGITAVEL, PAYMENT_OVERDUE,
		PAYMENT_UPDATED:
		return true
	}
	return false
}

func (c ChargebackReason) IsEnumValid() bool {
	switch c {
	case ABSENCE_OF_PRINT, ABSENT_CARD_FRAUD, CARD_ACTIVATED_PHONE_TRANSACTION, CARD_FRAUD, CARD_RECOVERY_BULLETIN,
		COMMERCIAL_DISAGREEMENT, COPY_NOT_RECEIVED, CREDIT_OR_DEBIT_PRESENTATION_ERROR, DIFFERENT_PAY_METHOD, FRAUD,
		INCORRECT_TRANSACTION_VALUE, INVALID_CURRENCY, INVALID_DATA, LATE_PRESENTATION, LOCAL_REGULATORY_OR_LEGAL_DISPUTE,
		MULTIPLE_ROCS, ORIGINAL_CREDIT_TRANSACTION_NOT_ACCEPTED, OTHER_ABSENT_CARD_FRAUD, PROCESS_ERROR,
		RECEIVED_COPY_ILLEGIBLE_OR_INCOMPLETE, RECURRENCE_CANCELED, REQUIRED_AUTHORIZATION_NOT_GRANTED,
		RIGHT_OF_FULL_RECOURSE_FOR_FRAUD, SALE_CANCELED, SERVICE_DISAGREEMENT_OR_DEFECTIVE_PRODUCT, SERVICE_NOT_RECEIVED,
		SPLIT_SALE, TRANSFERS_OF_DIVERSE_RESPONSIBILITIES, UNQUALIFIED_CAR_RENTAL_DEBIT, USA_CARDHOLDER_DISPUTE,
		VISA_FRAUD_MONITORING_PROGRAM, WARNING_BULLETIN_FILE:
		return true
	}
	return false
}

func (c BillingType) IsEnumValid() bool {
	switch c {
	case BOLETO, CREDIT_CARD, UNDEFINED, DEBIT_CARD, TRANSFER, DEPOSIT, PIX:
		return true
	}
	return false
}

func (c ChargeStatus) IsEnumValid() bool {
	switch c {
	case CHARGE_PENDING, CHARGE_RECEIVED, CHARGE_CONFIRMED, CHARGE_OVERDUE, CHARGE_REFUNDED, CHARGE_RECEIVED_IN_CASH,
		CHARGE_REFUND_REQUESTED, CHARGE_REFUND_IN_PROGRESS, CHARGE_CHARGEBACK_REQUESTED, CHARGE_CHARGEBACK_DISPUTE,
		CHARGE_AWAITING_CHARGEBACK_REVERSAL, CHARGE_DUNNING_REQUESTED, CHARGE_DUNNING_RECEIVED,
		CHARGE_AWAITING_RISK_ANALYSIS:
		return true
	}
	return false
}

func (c ChargebackStatus) IsEnumValid() bool {
	switch c {
	case CHARGEBACK_REQUESTED, CHARGEBACK_IN_DISPUTE, CHARGEBACK_LOST, CHARGEBACK_REVERSED, CHARGEBACK_DONE:
		return true
	}
	return false
}

func (d DiscountType) IsEnumValid() bool {
	switch d {
	case DISCOUNT_FIXED, DISCOUNT_PERCENTAGE:
		return true
	}
	return false
}

func (a Env) IsEnumValid() bool {
	switch a {
	case SANDBOX, PRODUCTION:
		return true
	}
	return false
}

func (e ErrorType) IsEnumValid() bool {
	switch e {
	case ERROR_VALIDATION, ERROR_UNEXPECTED:
		return true
	}
	return false
}

func (f FineType) IsEnumValid() bool {
	switch f {
	case FINE_FIXED, FINE_PERCENTAGE:
		return true
	}
	return false
}

func (i InterestType) IsEnumValid() bool {
	switch i {
	case INTEREST_FIXED, INTEREST_PERCENTAGE:
		return true
	}
	return false
}

func (r RefundStatus) IsEnumValid() bool {
	switch r {
	case REFUND_PENDING, REFUND_CANCELLED, REFUND_DONE:
		return true
	}
	return false
}

func (s SplitRefusalReason) IsEnumValid() bool {
	switch s {
	case RECEIVABLE_UNIT_AFFECTED_BY_EXTERNAL_CONTRACTUAL_EFFECT:
		return true
	}
	return false
}

func (s SplitStatus) IsEnumValid() bool {
	switch s {
	case SPLIT_PENDING, SPLIT_AWAITING_CREDIT, SPLIT_CANCELLED, SPLIT_DONE, SPLIT_REFUSED:
		return true
	}
	return false
}

func (t TypeOfDocument) IsEnumValid() bool {
	switch t {
	case INVOICE, CONTRACT, DOCUMENT, SPREADSHEET, PROGRAM, OTHER:
		return true
	}
	return false
}

func (c ChargebackReason) String() string {
	switch c {
	case ABSENCE_OF_PRINT:
		return "Ausência de impressão"
	case ABSENT_CARD_FRAUD:
		return "Fraude em ambiente de cartão não presente"
	case CARD_ACTIVATED_PHONE_TRANSACTION:
		return "Transação telefônica ativada por cartão"
	case CARD_FRAUD:
		return "Fraude em ambiente de cartão presente"
	case CARD_RECOVERY_BULLETIN:
		return "Boletim de negativação de cartões"
	case COMMERCIAL_DISAGREEMENT:
		return "Desacordo comercial"
	case COPY_NOT_RECEIVED:
		return "Cópia não atendida"
	case CREDIT_OR_DEBIT_PRESENTATION_ERROR:
		return "Erro de apresentação de crédito / débito"
	case DIFFERENT_PAY_METHOD:
		return "Pagamento por outros meios"
	case FRAUD:
		return "Sem autorização do portador do cartão"
	case INCORRECT_TRANSACTION_VALUE:
		return "Valor da transação é diferente"
	case INVALID_CURRENCY:
		return "Moeda inválida"
	case INVALID_DATA:
		return "Dados inválidos"
	case LATE_PRESENTATION:
		return "Apresentação tardia"
	case LOCAL_REGULATORY_OR_LEGAL_DISPUTE:
		return "Contestação regulatória / legal local"
	case MULTIPLE_ROCS:
		return "ROCs múltiplos"
	case ORIGINAL_CREDIT_TRANSACTION_NOT_ACCEPTED:
		return "Transação de crédito original não aceita"
	case OTHER_ABSENT_CARD_FRAUD:
		return "Outras fraudes - Cartão ausente"
	case PROCESS_ERROR:
		return "Erro de processamento"
	case RECEIVED_COPY_ILLEGIBLE_OR_INCOMPLETE:
		return "Cópia atendida ilegível / incompleta"
	case RECURRENCE_CANCELED:
		return "Recorrência cancelada"
	case REQUIRED_AUTHORIZATION_NOT_GRANTED:
		return "Autorização requerida não obtida"
	case RIGHT_OF_FULL_RECOURSE_FOR_FRAUD:
		return "Direito de regresso integral por fraude"
	case SALE_CANCELED:
		return "Mercadoria / serviços cancelado"
	case SERVICE_DISAGREEMENT_OR_DEFECTIVE_PRODUCT:
		return "Mercadoria / serviço com defeito ou em desacordo"
	case SERVICE_NOT_RECEIVED:
		return "Mercadoria / serviços não recebidos"
	case SPLIT_SALE:
		return "Desmembramento de venda"
	case TRANSFERS_OF_DIVERSE_RESPONSIBILITIES:
		return "Transf. de responsabilidades diversas"
	case UNQUALIFIED_CAR_RENTAL_DEBIT:
		return "Débito de aluguel de carro não qualificado"
	case USA_CARDHOLDER_DISPUTE:
		return "Contestação do portador de cartão (EUA)"
	case VISA_FRAUD_MONITORING_PROGRAM:
		return "Programa Visa de monitoramento de fraude"
	case WARNING_BULLETIN_FILE:
		return "Arquivo boletim de advertência"
	}
	return ""
}

func (a Env) String() string {
	return []string{"SANDBOX", "PRODUCTION"}[a]
}

func (a Env) BaseUrl() string {
	return []string{"https://sandbox.asaas.com/api", "https://api.asaas.com"}[a]
}
