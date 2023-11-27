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
type DocumentType string
type NotificationEvent string
type SubscriptionCycle string
type SubscriptionStatus string
type Order string
type SortSubscriptionField string
type InvoiceStatus string
type InvoiceDatePeriod string
type InvoiceDaysBeforeDueDate int
type PixKeyStatus string
type PixKeyType string
type QrCodeFormat string
type PixQrCodeType string
type PixTransactionStatus string
type PixTransactionFinality string
type PixTransactionType string
type PixTransactionOriginType string
type BankAccountType string
type SortPaymentBookField string
type PersonType string

const (
	PersonTypePhysical  PersonType = "FISICA"
	PersonTypeJuridical PersonType = "JURIDICA"
)
const (
	HttpContentTypeJSON = "application/json"
	HttpContentTypeText = "text/plain; charset=utf-8"
)
const (
	SortPaymentBookDueDate = "dueDate"
)
const (
	BankTypeCheckingAccount    BankAccountType = "CHECKING_ACCOUNT"
	BankTypeSalaryAccount      BankAccountType = "SALARY_ACCOUNT"
	BankTypeInvestimentAccount BankAccountType = "INVESTIMENT_ACCOUNT"
	BankTypePaymentAccount     BankAccountType = "PAYMENT_ACCOUNT"
)
const (
	PixQrCodeTypeStatic  PixQrCodeType = "STATIC"
	PixQrCodeTypeDynamic PixQrCodeType = "DYNAMIC"
)
const (
	PixTransactionOriginManual        PixTransactionOriginType = "MANUAL"
	PixTransactionOriginAddressKey    PixTransactionOriginType = "ADDRESS_KEY"
	PixTransactionOriginStaticQrcode  PixTransactionOriginType = "STATIC_QRCODE"
	PixTransactionOriginDynamicQrcode PixTransactionOriginType = "DYNAMIC_QRCODE"
	PixTransactionOriginExternalDebit PixTransactionOriginType = "EXTERNAL_DEBIT"
)
const (
	PixTransactionTypeDebit                   PixTransactionType = "DEBIT"
	PixTransactionTypeCredit                  PixTransactionType = "CREDIT"
	PixTransactionTypeCreditRefund            PixTransactionType = "CREDIT_REFUND"
	PixTransactionTypeDebitRefund             PixTransactionType = "DEBIT_REFUND"
	PixTransactionTypeDebitRefundCancellation PixTransactionType = "DEBIT_REFUND_CANCELLATION"
)
const (
	PixTransactionFinalityWithdrawal PixTransactionFinality = "WITHDRAWAL"
	PixTransactionFinalityChange     PixTransactionFinality = "CHANGE"
)
const (
	PixTransactionAwaitingRequest PixTransactionStatus = "AWAITING_REQUEST"
	PixTransactionDone            PixTransactionStatus = "DONE"
	PixTransactionRequested       PixTransactionStatus = "REQUESTED"
	PixTransactionScheduled       PixTransactionStatus = "SCHEDULED"
	PixTransactionRefused         PixTransactionStatus = "REFUSED"
	PixTransactionError           PixTransactionStatus = "ERROR"
	PixTransactionCancelled       PixTransactionStatus = "CANCELLED"
)
const (
	QrCodeFormatAll     QrCodeFormat = "ALL"
	QrCodeFormatImage   QrCodeFormat = "IMAGE"
	QrCodeFormatPayload QrCodeFormat = "PAYLOAD"
)
const (
	PixKeyTypeCpf   PixKeyType = "CPF"
	PixKeyTypeCnpj  PixKeyType = "CNPJ"
	PixKeyTypeEmail PixKeyType = "EMAIL"
	PixKeyTypePhone PixKeyType = "PHONE"
	PixKeyTypeEvp   PixKeyType = "EVP"
)
const (
	PixKeyAwaitingActivation      PixKeyStatus = "AWAITING_ACTIVATION"
	PixKeyAwaitingActive          PixKeyStatus = "ACTIVE"
	PixKeyAwaitingDeletion        PixKeyStatus = "AWAITING_DELETION"
	PixKeyAwaitingAccountDeletion PixKeyStatus = "AWAITING_ACCOUNT_DELETION"
	PixKeyAwaitingDeleted         PixKeyStatus = "AWAITING_DELETED"
	PixKeyAwaitingError           PixKeyStatus = "AWAITING_ERROR"
)
const (
	InvoiceDaysBeforeDuedateFive    InvoiceDaysBeforeDueDate = 5
	InvoiceDaysBeforeDuedateTen     InvoiceDaysBeforeDueDate = 10
	InvoiceDaysBeforeDuedateFifteen InvoiceDaysBeforeDueDate = 15
	InvoiceDaysBeforeDuedateThirty  InvoiceDaysBeforeDueDate = 30
	InvoiceDaysBeforeDuedateSixty   InvoiceDaysBeforeDueDate = 60
)
const (
	InvoiceDatePeriodOnPaymentConfirmation InvoiceDatePeriod = "ON_PAYMENT_CONFIRMATION"
	InvoiceDatePeriodOnPaymentDueDate      InvoiceDatePeriod = "ON_PAYMENT_DUE_DATE"
	InvoiceDatePeriodBeforePaymentDueDate  InvoiceDatePeriod = "BEFORE_PAYMENT_DUE_DATE"
	InvoiceDatePeriodOnDueDateMonth        InvoiceDatePeriod = "ON_DUE_DATE_MONTH"
	InvoiceDatePeriodOnNextMonth           InvoiceDatePeriod = "ON_NEXT_MONTH"
)
const (
	InvoiceScheduled              InvoiceStatus = "SCHEDULED"
	InvoiceSynchronized           InvoiceStatus = "SYNCHRONIZED"
	InvoiceAuthorized             InvoiceStatus = "AUTHORIZED"
	InvoiceProcessingCancellation InvoiceStatus = "PROCESSING_CANCELLATION"
	InvoiceCanceled               InvoiceStatus = "CANCELED"
	InvoiceCancellationDenied     InvoiceStatus = "CANCELLATION_DENIED"
	InvoiceError                  InvoiceStatus = "ERROR"
)
const (
	SortSubscriptionDateCreated SortSubscriptionField = "dateCreated"
)
const (
	OrderDesc Order = "desc"
	OrderAsc  Order = "asc"
)
const (
	SubscriptionActive   SubscriptionStatus = "ACTIVE"
	SubscriptionInactive SubscriptionStatus = "INACTIVE"
	SubscriptionExpired  SubscriptionStatus = "EXPIRED"
)
const (
	Weekly       SubscriptionCycle = "WEEKLY"
	Biweekly     SubscriptionCycle = "BIWEEKLY"
	Monthly      SubscriptionCycle = "MONTHLY"
	Bimonthly    SubscriptionCycle = "BIMONTHLY"
	Quarterly    SubscriptionCycle = "QUARTERLY"
	Semiannually SubscriptionCycle = "SEMIANNUALLY"
	Yearly       SubscriptionCycle = "YEARLY"
)
const (
	BillingTypeBoleto     BillingType = "BOLETO"
	BillingTypeCreditCard BillingType = "CREDIT_CARD"
	BillingTypeUndefined  BillingType = "UNDEFINED"
	BillingTypeDebitCard  BillingType = "DEBIT_CARD"
	BillingTypeTransfer   BillingType = "TRANSFER"
	BillingTypeDeposit    BillingType = "DEPOSIT"
	BillingTypePix        BillingType = "PIX"
)
const (
	ChargeStatusPending                    ChargeStatus = "PENDING"
	ChargeStatusReceived                   ChargeStatus = "RECEIVED"
	ChargeStatusConfirmed                  ChargeStatus = "CONFIRMED"
	ChargeStatusOverdue                    ChargeStatus = "OVERDUE"
	ChargeStatusRefunded                   ChargeStatus = "REFUNDED"
	ChargeStatusReceivedInCash             ChargeStatus = "RECEIVED_IN_CASH"
	ChargeStatusRefundRequested            ChargeStatus = "REFUND_REQUESTED"
	ChargeStatusRefundInProgress           ChargeStatus = "REFUND_IN_PROGRESS"
	ChargeStatusChargebackRequested        ChargeStatus = "CHARGEBACK_REQUESTED"
	ChargeStatusChargebackDispute          ChargeStatus = "CHARGEBACK_DISPUTE"
	ChargeStatusAwaitingChargebackReversal ChargeStatus = "AWAITING_CHARGEBACK_REVERSAL"
	ChargeStatusDunningRequested           ChargeStatus = "DUNNING_REQUESTED"
	ChargeStatusDunningReceived            ChargeStatus = "DUNNING_RECEIVED"
	ChargeStatusAwaitingRiskAnalysis       ChargeStatus = "AWAITING_RISK_ANALYSIS"
)
const (
	ChargebackReasonAbsenceOfPrint                        ChargebackReason = "ABSENCE_OF_PRINT"
	ChargebackReasonAbsentCardFraud                       ChargebackReason = "ABSENT_CARD_FRAUD"
	ChargebackReasonCardActivatedPhoneTransaction         ChargebackReason = "CARD_ACTIVATED_PHONE_TRANSACTION"
	ChargebackReasonCardFraud                             ChargebackReason = "CARD_FRAUD"
	ChargebackReasonCardRecoveryBulletin                  ChargebackReason = "CARD_RECOVERY_BULLETIN"
	ChargebackReasonCommercialDisagreement                ChargebackReason = "COMMERCIAL_DISAGREEMENT"
	ChargebackReasonCopyNotReceived                       ChargebackReason = "COPY_NOT_RECEIVED"
	ChargebackReasonCreditOrDebitPresentationError        ChargebackReason = "CREDIT_OR_DEBIT_PRESENTATION_ERROR"
	ChargebackReasonDifferentPayMethod                    ChargebackReason = "DIFFERENT_PAY_METHOD"
	ChargebackReasonFraud                                 ChargebackReason = "FRAUD"
	ChargebackReasonIncorrectTransactionValue             ChargebackReason = "INCORRECT_TRANSACTION_VALUE"
	ChargebackReasonInvalidCurrency                       ChargebackReason = "INVALID_CURRENCY"
	ChargebackReasonInvalidData                           ChargebackReason = "INVALID_DATA"
	ChargebackReasonLatePresentation                      ChargebackReason = "LATE_PRESENTATION"
	ChargebackReasonLocalRegulatoryOrLegalDispute         ChargebackReason = "LOCAL_REGULATORY_OR_LEGAL_DISPUTE"
	ChargebackReasonMultipleRocs                          ChargebackReason = "MULTIPLE_ROCS"
	ChargebackReasonOriginalCreditTransactionNotAccepted  ChargebackReason = "ORIGINAL_CREDIT_TRANSACTION_NOT_ACCEPTED"
	ChargebackReasonOtherAbsentCardFraud                  ChargebackReason = "OTHER_ABSENT_CARD_FRAUD"
	ChargebackReasonProcessError                          ChargebackReason = "PROCESS_ERROR"
	ChargebackReasonReceivedCopyIllegibleOrIncomplete     ChargebackReason = "RECEIVED_COPY_ILLEGIBLE_OR_INCOMPLETE"
	ChargebackReasonRecurrenceCanceled                    ChargebackReason = "RECURRENCE_CANCELED"
	ChargebackReasonRequiredAuthorizationNotGranted       ChargebackReason = "REQUIRED_AUTHORIZATION_NOT_GRANTED"
	ChargebackReasonRightOfFullRecourseForFraud           ChargebackReason = "RIGHT_OF_FULL_RECOURSE_FOR_FRAUD"
	ChargebackReasonSaleCanceled                          ChargebackReason = "SALE_CANCELED"
	ChargebackReasonServiceDisagreementOrDefectiveProduct ChargebackReason = "SERVICE_DISAGREEMENT_OR_DEFECTIVE_PRODUCT"
	ChargebackReasonServiceNotReceived                    ChargebackReason = "SERVICE_NOT_RECEIVED"
	ChargebackReasonSplitSale                             ChargebackReason = "SPLIT_SALE"
	ChargebackReasonTransfersOfDiverseResponsibilities    ChargebackReason = "TRANSFERS_OF_DIVERSE_RESPONSIBILITIES"
	ChargebackReasonUnqualifiedCarRentalDebit             ChargebackReason = "UNQUALIFIED_CAR_RENTAL_DEBIT"
	ChargebackReasonUsaCardholderDispute                  ChargebackReason = "USA_CARDHOLDER_DISPUTE"
	ChargebackReasonVisaFraudMonitoringProgram            ChargebackReason = "VISA_FRAUD_MONITORING_PROGRAM"
	ChargebackReasonWarningBulletinFile                   ChargebackReason = "WARNING_BULLETIN_FILE"
)
const (
	ChargebackStatusRequested ChargebackStatus = "REQUESTED"
	ChargebackStatusInDispute ChargebackStatus = "IN_DISPUTE"
	ChargebackStatusLost      ChargebackStatus = "LOST"
	ChargebackStatusReversed  ChargebackStatus = "REVERSED"
	ChargebackStatusDone      ChargebackStatus = "DONE"
)
const (
	DiscountTypeFixed      DiscountType = "FIXED"
	DiscountTypePercentage DiscountType = "PERCENTAGE"
)
const (
	EnvSandbox    Env = iota
	EnvProduction Env = iota
)
const (
	ErrorTypeValidation ErrorType = "VALIDATION"
	ErrorTypeUnexpected ErrorType = "UNEXPECTED"
)
const (
	FineTypeFixed      FineType = "FIXED"
	FineTypePercentage FineType = "PERCENTAGE"
)
const (
	InterestTypeFixed      InterestType = "FIXED"
	InterestTypePercentage InterestType = "PERCENTAGE"
)
const (
	RefundStatusPending   RefundStatus = "PENDING"
	RefundStatusCancelled RefundStatus = "CANCELLED"
	RefundStatusDone      RefundStatus = "DONE"
)
const (
	SplitRefusalReason1 SplitRefusalReason = "RECEIVABLE_UNIT_AFFECTED_BY_EXTERNAL_CONTRACTUAL_EFFECT"
)
const (
	SplitStatusPending        SplitStatus = "PENDING"
	SplitStatusAwaitingCredit SplitStatus = "AWAITING_CREDIT"
	SplitStatusCancelled      SplitStatus = "CANCELLED"
	SplitStatusDone           SplitStatus = "DONE"
	SplitStatusRefused        SplitStatus = "REFUSED"
)
const (
	DocumentTypeInvoice     DocumentType = "INVOICE"
	DocumentTypeContract    DocumentType = "CONTRACT"
	DocumentTypeDocument    DocumentType = "DOCUMENT"
	DocumentTypeSpreadsheet DocumentType = "SPREADSHEET"
	DocumentTypeProgram     DocumentType = "PROGRAM"
	DocumentTypeOther       DocumentType = "OTHER"
)
const (
	NotificationEventPaymentCreated        NotificationEvent = "PAYMENT_CREATED"
	NotificationEventPaymentDuedateWarning NotificationEvent = "PAYMENT_DUEDATE_WARNING"
	NotificationEventPaymentReceived       NotificationEvent = "PAYMENT_RECEIVED"
	NotificationEventSendLinhaDigitavel    NotificationEvent = "SEND_LINHA_DIGITAVEL"
	NotificationEventPaymentOverdue        NotificationEvent = "PAYMENT_OVERDUE"
	NotificationEventPaymentUpdated        NotificationEvent = "PAYMENT_UPDATED"
)

func (p PersonType) IsEnumValid() bool {
	switch p {
	case PersonTypePhysical, PersonTypeJuridical:
		return true
	}
	return false
}

func (b BankAccountType) IsEnumValid() bool {
	switch b {
	case BankTypeSalaryAccount, BankTypePaymentAccount, BankTypeInvestimentAccount, BankTypeCheckingAccount:
		return true
	}
	return false
}

func (p PixTransactionFinality) IsEnumValid() bool {
	switch p {
	case PixTransactionFinalityChange, PixTransactionFinalityWithdrawal:
		return true
	}
	return false
}

func (p PixTransactionStatus) IsEnumValid() bool {
	switch p {
	case PixTransactionAwaitingRequest, PixTransactionDone, PixTransactionRequested, PixTransactionScheduled,
		PixTransactionRefused, PixTransactionError, PixTransactionCancelled:
		return true
	}
	return false
}

func (p PixTransactionOriginType) IsEnumValid() bool {
	switch p {
	case PixTransactionOriginManual, PixTransactionOriginAddressKey, PixTransactionOriginStaticQrcode,
		PixTransactionOriginDynamicQrcode, PixTransactionOriginExternalDebit:
		return true
	}
	return false
}

func (p PixTransactionType) IsEnumValid() bool {
	switch p {
	case PixTransactionTypeCredit, PixTransactionTypeCreditRefund, PixTransactionTypeDebit,
		PixTransactionTypeDebitRefund, PixTransactionTypeDebitRefundCancellation:
		return true
	}
	return false
}

func (p PixQrCodeType) IsEnumValid() bool {
	switch p {
	case PixQrCodeTypeStatic, PixQrCodeTypeDynamic:
		return true
	}
	return false
}

func (q QrCodeFormat) IsEnumValid() bool {
	switch q {
	case QrCodeFormatAll, QrCodeFormatImage, QrCodeFormatPayload:
		return true
	}
	return false
}

func (p PixKeyStatus) IsEnumValid() bool {
	switch p {
	case PixKeyAwaitingActive, PixKeyAwaitingActivation, PixKeyAwaitingDeleted, PixKeyAwaitingAccountDeletion,
		PixKeyAwaitingDeletion, PixKeyAwaitingError:
		return true
	}
	return false
}

func (p PixKeyType) IsEnumValid() bool {
	switch p {
	case PixKeyTypeCnpj, PixKeyTypeCpf, PixKeyTypeEmail, PixKeyTypePhone, PixKeyTypeEvp:
		return true
	}
	return false
}

func (i InvoiceDaysBeforeDueDate) IsEnumValid() bool {
	switch i {
	case InvoiceDaysBeforeDuedateFive, InvoiceDaysBeforeDuedateTen, InvoiceDaysBeforeDuedateFifteen,
		InvoiceDaysBeforeDuedateThirty, InvoiceDaysBeforeDuedateSixty:
		return true
	}
	return false
}

func (i InvoiceStatus) IsEnumValid() bool {
	switch i {
	case InvoiceAuthorized, InvoiceCanceled, InvoiceCancellationDenied, InvoiceProcessingCancellation,
		InvoiceError, InvoiceScheduled, InvoiceSynchronized:
		return true
	}
	return false
}

func (i InvoiceDatePeriod) IsEnumValid() bool {
	switch i {
	case InvoiceDatePeriodOnPaymentConfirmation, InvoiceDatePeriodOnPaymentDueDate,
		InvoiceDatePeriodBeforePaymentDueDate, InvoiceDatePeriodOnDueDateMonth,
		InvoiceDatePeriodOnNextMonth:
		return true
	}
	return false
}

func (s SortPaymentBookField) IsEnumValid() bool {
	switch s {
	case SortPaymentBookDueDate:
		return true
	}
	return false
}

func (s SortSubscriptionField) IsEnumValid() bool {
	switch s {
	case SortSubscriptionDateCreated:
		return true
	}
	return false
}

func (s SubscriptionStatus) IsEnumValid() bool {
	switch s {
	case SubscriptionActive, SubscriptionInactive, SubscriptionExpired:
		return true
	}
	return false
}

func (o Order) IsEnumValid() bool {
	switch o {
	case OrderDesc, OrderAsc:
		return true
	}
	return false
}

func (s SubscriptionCycle) IsEnumValid() bool {
	switch s {
	case Weekly, Biweekly, Monthly, Bimonthly, Quarterly, Semiannually, Yearly:
		return true
	}
	return false
}

func (n NotificationEvent) IsEnumValid() bool {
	switch n {
	case NotificationEventPaymentCreated, NotificationEventPaymentDuedateWarning, NotificationEventPaymentReceived, NotificationEventSendLinhaDigitavel, NotificationEventPaymentOverdue,
		NotificationEventPaymentUpdated:
		return true
	}
	return false
}

func (c ChargebackReason) IsEnumValid() bool {
	switch c {
	case ChargebackReasonAbsenceOfPrint, ChargebackReasonAbsentCardFraud, ChargebackReasonCardActivatedPhoneTransaction, ChargebackReasonCardFraud, ChargebackReasonCardRecoveryBulletin,
		ChargebackReasonCommercialDisagreement, ChargebackReasonCopyNotReceived, ChargebackReasonCreditOrDebitPresentationError, ChargebackReasonDifferentPayMethod, ChargebackReasonFraud,
		ChargebackReasonIncorrectTransactionValue, ChargebackReasonInvalidCurrency, ChargebackReasonInvalidData, ChargebackReasonLatePresentation, ChargebackReasonLocalRegulatoryOrLegalDispute,
		ChargebackReasonMultipleRocs, ChargebackReasonOriginalCreditTransactionNotAccepted, ChargebackReasonOtherAbsentCardFraud, ChargebackReasonProcessError,
		ChargebackReasonReceivedCopyIllegibleOrIncomplete, ChargebackReasonRecurrenceCanceled, ChargebackReasonRequiredAuthorizationNotGranted,
		ChargebackReasonRightOfFullRecourseForFraud, ChargebackReasonSaleCanceled, ChargebackReasonServiceDisagreementOrDefectiveProduct, ChargebackReasonServiceNotReceived,
		ChargebackReasonSplitSale, ChargebackReasonTransfersOfDiverseResponsibilities, ChargebackReasonUnqualifiedCarRentalDebit, ChargebackReasonUsaCardholderDispute,
		ChargebackReasonVisaFraudMonitoringProgram, ChargebackReasonWarningBulletinFile:
		return true
	}
	return false
}

func (c BillingType) IsEnumValid() bool {
	switch c {
	case BillingTypeBoleto, BillingTypeCreditCard, BillingTypeUndefined, BillingTypeDebitCard, BillingTypeTransfer, BillingTypeDeposit, BillingTypePix:
		return true
	}
	return false
}

func (c ChargeStatus) IsEnumValid() bool {
	switch c {
	case ChargeStatusPending, ChargeStatusReceived, ChargeStatusConfirmed, ChargeStatusOverdue, ChargeStatusRefunded, ChargeStatusReceivedInCash,
		ChargeStatusRefundRequested, ChargeStatusRefundInProgress, ChargeStatusChargebackRequested, ChargeStatusChargebackDispute,
		ChargeStatusAwaitingChargebackReversal, ChargeStatusDunningRequested, ChargeStatusDunningReceived,
		ChargeStatusAwaitingRiskAnalysis:
		return true
	}
	return false
}

func (c ChargebackStatus) IsEnumValid() bool {
	switch c {
	case ChargebackStatusRequested, ChargebackStatusInDispute, ChargebackStatusLost, ChargebackStatusReversed, ChargebackStatusDone:
		return true
	}
	return false
}

func (d DiscountType) IsEnumValid() bool {
	switch d {
	case DiscountTypeFixed, DiscountTypePercentage:
		return true
	}
	return false
}

func (a Env) IsEnumValid() bool {
	switch a {
	case EnvSandbox, EnvProduction:
		return true
	}
	return false
}

func (e ErrorType) IsEnumValid() bool {
	switch e {
	case ErrorTypeValidation, ErrorTypeUnexpected:
		return true
	}
	return false
}

func (f FineType) IsEnumValid() bool {
	switch f {
	case FineTypeFixed, FineTypePercentage:
		return true
	}
	return false
}

func (i InterestType) IsEnumValid() bool {
	switch i {
	case InterestTypeFixed, InterestTypePercentage:
		return true
	}
	return false
}

func (r RefundStatus) IsEnumValid() bool {
	switch r {
	case RefundStatusPending, RefundStatusCancelled, RefundStatusDone:
		return true
	}
	return false
}

func (s SplitRefusalReason) IsEnumValid() bool {
	switch s {
	case SplitRefusalReason1:
		return true
	}
	return false
}

func (s SplitStatus) IsEnumValid() bool {
	switch s {
	case SplitStatusPending, SplitStatusAwaitingCredit, SplitStatusCancelled, SplitStatusDone, SplitStatusRefused:
		return true
	}
	return false
}

func (t DocumentType) IsEnumValid() bool {
	switch t {
	case DocumentTypeInvoice, DocumentTypeContract, DocumentTypeDocument, DocumentTypeSpreadsheet, DocumentTypeProgram, DocumentTypeOther:
		return true
	}
	return false
}

func (c ChargebackReason) String() string {
	switch c {
	case ChargebackReasonAbsenceOfPrint:
		return "Ausência de impressão"
	case ChargebackReasonAbsentCardFraud:
		return "Fraude em ambiente de cartão não presente"
	case ChargebackReasonCardActivatedPhoneTransaction:
		return "Transação telefônica ativada por cartão"
	case ChargebackReasonCardFraud:
		return "Fraude em ambiente de cartão presente"
	case ChargebackReasonCardRecoveryBulletin:
		return "Boletim de negativação de cartões"
	case ChargebackReasonCommercialDisagreement:
		return "Desacordo comercial"
	case ChargebackReasonCopyNotReceived:
		return "Cópia não atendida"
	case ChargebackReasonCreditOrDebitPresentationError:
		return "Erro de apresentação de crédito / débito"
	case ChargebackReasonDifferentPayMethod:
		return "Pagamento por outros meios"
	case ChargebackReasonFraud:
		return "Sem autorização do portador do cartão"
	case ChargebackReasonIncorrectTransactionValue:
		return "Valor da transação é diferente"
	case ChargebackReasonInvalidCurrency:
		return "Moeda inválida"
	case ChargebackReasonInvalidData:
		return "Dados inválidos"
	case ChargebackReasonLatePresentation:
		return "Apresentação tardia"
	case ChargebackReasonLocalRegulatoryOrLegalDispute:
		return "Contestação regulatória / legal local"
	case ChargebackReasonMultipleRocs:
		return "ROCs múltiplos"
	case ChargebackReasonOriginalCreditTransactionNotAccepted:
		return "Transação de crédito original não aceita"
	case ChargebackReasonOtherAbsentCardFraud:
		return "Outras fraudes - Cartão ausente"
	case ChargebackReasonProcessError:
		return "Erro de processamento"
	case ChargebackReasonReceivedCopyIllegibleOrIncomplete:
		return "Cópia atendida ilegível / incompleta"
	case ChargebackReasonRecurrenceCanceled:
		return "Recorrência cancelada"
	case ChargebackReasonRequiredAuthorizationNotGranted:
		return "Autorização requerida não obtida"
	case ChargebackReasonRightOfFullRecourseForFraud:
		return "Direito de regresso integral por fraude"
	case ChargebackReasonSaleCanceled:
		return "Mercadoria / serviços cancelado"
	case ChargebackReasonServiceDisagreementOrDefectiveProduct:
		return "Mercadoria / serviço com defeito ou em desacordo"
	case ChargebackReasonServiceNotReceived:
		return "Mercadoria / serviços não recebidos"
	case ChargebackReasonSplitSale:
		return "Desmembramento de venda"
	case ChargebackReasonTransfersOfDiverseResponsibilities:
		return "Transf. de responsabilidades diversas"
	case ChargebackReasonUnqualifiedCarRentalDebit:
		return "Débito de aluguel de carro não qualificado"
	case ChargebackReasonUsaCardholderDispute:
		return "Contestação do portador de cartão (EUA)"
	case ChargebackReasonVisaFraudMonitoringProgram:
		return "Programa Visa de monitoramento de fraude"
	case ChargebackReasonWarningBulletinFile:
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
