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
type AccountType string
type SortPaymentBookField string
type PersonType string
type TransferType string
type TransferStatus string
type TransferAsaasStatus string
type TransferOperationType string
type BankAccountType string
type AnticipationStatus string
type NegativityStatus string
type NegativityType string
type FileMimeType string
type BillPaymentStatus string

const (
	BillPaymentStatusPending        BillPaymentStatus = "PENDING"
	BillPaymentStatusBankProcessing BillPaymentStatus = "BANK_PROCESSING"
	BillPaymentStatusBankPaid       BillPaymentStatus = "PAID"
	BillPaymentStatusFailed         BillPaymentStatus = "FAILED"
	BillPaymentStatusCancelled      BillPaymentStatus = "CANCELLED"
)
const (
	FileMimeTypePdf  FileMimeType = "application/pdf"
	FileMimeTypeText FileMimeType = "text/plain"
	FileMimeTypeAvif FileMimeType = "image/avif"
	FileMimeTypeCss  FileMimeType = "text/css; charset=utf-8"
	FileMimeTypeGif  FileMimeType = "image/gif"
	FileMimeTypeHtml FileMimeType = "text/html; charset=utf-8"
	FileMimeTypeJpeg FileMimeType = "image/jpeg"
	FileMimeTypeJs   FileMimeType = "text/javascript; charset=utf-8"
	FileMimeTypeJson FileMimeType = "application/json"
	FileMimeTypePng  FileMimeType = "image/png"
	FileMimeTypeSvg  FileMimeType = "image/svg+xml"
	FileMimeTypeWasm FileMimeType = "application/wasm"
	FileMimeTypeWebp FileMimeType = "image/webp"
	FileMimeTypeXml  FileMimeType = "text/xml; charset=utf-8"
)
const (
	NegativityTypeCreditBureau NegativityType = "CREDIT_BUREAU"
)
const (
	NegativityStatusPending              NegativityStatus = "PENDING"
	NegativityStatusAwaitingApproval     NegativityStatus = "AWAITING_APPROVAL"
	NegativityStatusAwaitingCancellation NegativityStatus = "AWAITING_CANCELLATION"
	NegativityStatusProcessed            NegativityStatus = "PROCESSED"
	NegativityStatusPaid                 NegativityStatus = "PAID"
	NegativityStatusPartiallyPaid        NegativityStatus = "PARTIALLY_PAID"
	NegativityStatusDenied               NegativityStatus = "DENIED"
	NegativityStatusCancelled            NegativityStatus = "CANCELLED"
)
const (
	AnticipationStatusPending   AnticipationStatus = "PENDING"
	AnticipationStatusDenied    AnticipationStatus = "DENIED"
	AnticipationStatusCredited  AnticipationStatus = "CREDITED"
	AnticipationStatusDebited   AnticipationStatus = "DEBITED"
	AnticipationStatusCancelled AnticipationStatus = "CANCELLED"
	AnticipationStatusOverdue   AnticipationStatus = "OVERDUE"
	AnticipationStatusScheduled AnticipationStatus = "SCHEDULED"
)
const (
	TransferTypeBankAccount  TransferType = "BANK_ACCOUNT"
	TransferTypeAsaasAccount TransferType = "ASAAS_ACCOUNT"
)
const (
	BankAccountTypeChecking BankAccountType = "CONTA_CORRENTE"
	BankAccountTypeSavings  BankAccountType = "CONTA_POUPANCA"
)
const (
	TransferOperationTypePix      TransferOperationType = "PIX"
	TransferOperationTypeTed      TransferOperationType = "TED"
	TransferOperationTypeInternal TransferOperationType = "INTERNAL"
)
const (
	TransferStatusPending        TransferStatus = "PENDING"
	TransferStatusBankProcessing TransferStatus = "BANK_PROCESSING"
	TransferStatusDone           TransferStatus = "DONE"
	TransferStatusCancelled      TransferStatus = "CANCELLED"
	TransferStatusFailed         TransferStatus = "FAILED"
)
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
	AccountTypeCheckingAccount    AccountType = "CHECKING_ACCOUNT"
	AccountTypeSalaryAccount      AccountType = "SALARY_ACCOUNT"
	AccountTypeInvestimentAccount AccountType = "INVESTIMENT_ACCOUNT"
	AccountTypeTypePaymentAccount AccountType = "PAYMENT_ACCOUNT"
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
	PixTransactionStatusAwaitingRequest PixTransactionStatus = "AWAITING_REQUEST"
	PixTransactionStatusDone            PixTransactionStatus = "DONE"
	PixTransactionStatusRequested       PixTransactionStatus = "REQUESTED"
	PixTransactionStatusScheduled       PixTransactionStatus = "SCHEDULED"
	PixTransactionStatusRefused         PixTransactionStatus = "REFUSED"
	PixTransactionStatusError           PixTransactionStatus = "ERROR"
	PixTransactionStatusCancelled       PixTransactionStatus = "CANCELLED"
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
	PixKeyStatusAwaitingActivation      PixKeyStatus = "AWAITING_ACTIVATION"
	PixKeyStatusAwaitingActive          PixKeyStatus = "ACTIVE"
	PixKeyStatusAwaitingDeletion        PixKeyStatus = "AWAITING_DELETION"
	PixKeyStatusAwaitingAccountDeletion PixKeyStatus = "AWAITING_ACCOUNT_DELETION"
	PixKeyStatusAwaitingDeleted         PixKeyStatus = "AWAITING_DELETED"
	PixKeyStatusAwaitingError           PixKeyStatus = "AWAITING_ERROR"
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
	InvoiceStatusScheduled              InvoiceStatus = "SCHEDULED"
	InvoiceStatusSynchronized           InvoiceStatus = "SYNCHRONIZED"
	InvoiceStatusAuthorized             InvoiceStatus = "AUTHORIZED"
	InvoiceStatusProcessingCancellation InvoiceStatus = "PROCESSING_CANCELLATION"
	InvoiceStatusCanceled               InvoiceStatus = "CANCELED"
	InvoiceStatusCancellationDenied     InvoiceStatus = "CANCELLATION_DENIED"
	InvoiceStatusError                  InvoiceStatus = "ERROR"
)
const (
	SortSubscriptionFieldDateCreated SortSubscriptionField = "dateCreated"
)
const (
	OrderDesc Order = "desc"
	OrderAsc  Order = "asc"
)
const (
	SubscriptionStatusActive   SubscriptionStatus = "ACTIVE"
	SubscriptionStatusInactive SubscriptionStatus = "INACTIVE"
	SubscriptionStatusExpired  SubscriptionStatus = "EXPIRED"
)
const (
	SubscriptionCycleWeekly       SubscriptionCycle = "WEEKLY"
	SubscriptionCycleBiweekly     SubscriptionCycle = "BIWEEKLY"
	SubscriptionCycleMonthly      SubscriptionCycle = "MONTHLY"
	SubscriptionCycleBimonthly    SubscriptionCycle = "BIMONTHLY"
	SubscriptionCycleQuarterly    SubscriptionCycle = "QUARTERLY"
	SubscriptionCycleSemiannually SubscriptionCycle = "SEMIANNUALLY"
	SubscriptionCycleYearly       SubscriptionCycle = "YEARLY"
)
const (
	BillingTypeBill       BillingType = "BOLETO"
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
	ChargebackReason1  ChargebackReason = "ABSENCE_OF_PRINT"
	ChargebackReason2  ChargebackReason = "ABSENT_CARD_FRAUD"
	ChargebackReason3  ChargebackReason = "CARD_ACTIVATED_PHONE_TRANSACTION"
	ChargebackReason4  ChargebackReason = "CARD_FRAUD"
	ChargebackReason5  ChargebackReason = "CARD_RECOVERY_BULLETIN"
	ChargebackReason6  ChargebackReason = "COMMERCIAL_DISAGREEMENT"
	ChargebackReason7  ChargebackReason = "COPY_NOT_RECEIVED"
	ChargebackReason8  ChargebackReason = "CREDIT_OR_DEBIT_PRESENTATION_ERROR"
	ChargebackReason9  ChargebackReason = "DIFFERENT_PAY_METHOD"
	ChargebackReason10 ChargebackReason = "FRAUD"
	ChargebackReason11 ChargebackReason = "INCORRECT_TRANSACTION_VALUE"
	ChargebackReason12 ChargebackReason = "INVALID_CURRENCY"
	ChargebackReason13 ChargebackReason = "INVALID_DATA"
	ChargebackReason14 ChargebackReason = "LATE_PRESENTATION"
	ChargebackReason15 ChargebackReason = "LOCAL_REGULATORY_OR_LEGAL_DISPUTE"
	ChargebackReason16 ChargebackReason = "MULTIPLE_ROCS"
	ChargebackReason17 ChargebackReason = "ORIGINAL_CREDIT_TRANSACTION_NOT_ACCEPTED"
	ChargebackReason18 ChargebackReason = "OTHER_ABSENT_CARD_FRAUD"
	ChargebackReason19 ChargebackReason = "PROCESS_ERROR"
	ChargebackReason20 ChargebackReason = "RECEIVED_COPY_ILLEGIBLE_OR_INCOMPLETE"
	ChargebackReason21 ChargebackReason = "RECURRENCE_CANCELED"
	ChargebackReason22 ChargebackReason = "REQUIRED_AUTHORIZATION_NOT_GRANTED"
	ChargebackReason23 ChargebackReason = "RIGHT_OF_FULL_RECOURSE_FOR_FRAUD"
	ChargebackReason24 ChargebackReason = "SALE_CANCELED"
	ChargebackReason25 ChargebackReason = "SERVICE_DISAGREEMENT_OR_DEFECTIVE_PRODUCT"
	ChargebackReason26 ChargebackReason = "SERVICE_NOT_RECEIVED"
	ChargebackReason27 ChargebackReason = "SPLIT_SALE"
	ChargebackReason28 ChargebackReason = "TRANSFERS_OF_DIVERSE_RESPONSIBILITIES"
	ChargebackReason29 ChargebackReason = "UNQUALIFIED_CAR_RENTAL_DEBIT"
	ChargebackReason30 ChargebackReason = "USA_CARDHOLDER_DISPUTE"
	ChargebackReason31 ChargebackReason = "VISA_FRAUD_MONITORING_PROGRAM"
	ChargebackReason32 ChargebackReason = "WARNING_BULLETIN_FILE"
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

func (f BillPaymentStatus) IsEnumValid() bool {
	switch f {
	case BillPaymentStatusPending, BillPaymentStatusBankProcessing, BillPaymentStatusBankPaid, BillPaymentStatusFailed,
		BillPaymentStatusCancelled:
		return true
	}
	return false
}

func (f FileMimeType) IsEnumValid() bool {
	switch f {
	case FileMimeTypePdf, FileMimeTypeText, FileMimeTypeAvif, FileMimeTypeCss, FileMimeTypeGif, FileMimeTypeHtml,
		FileMimeTypeJpeg, FileMimeTypeJs, FileMimeTypeJson, FileMimeTypePng, FileMimeTypeSvg, FileMimeTypeWasm,
		FileMimeTypeWebp, FileMimeTypeXml:
		return true
	}
	return false
}

func (a NegativityStatus) IsEnumValid() bool {
	switch a {
	case NegativityStatusAwaitingApproval, NegativityStatusAwaitingCancellation, NegativityStatusCancelled,
		NegativityStatusDenied, NegativityStatusPaid, NegativityStatusPending, NegativityStatusProcessed,
		NegativityStatusPartiallyPaid:
		return true
	}
	return false
}

func (a NegativityType) IsEnumValid() bool {
	switch a {
	case NegativityTypeCreditBureau:
		return true
	}
	return false
}

func (a AnticipationStatus) IsEnumValid() bool {
	switch a {
	case AnticipationStatusCredited, AnticipationStatusCancelled, AnticipationStatusDenied, AnticipationStatusDebited,
		AnticipationStatusOverdue, AnticipationStatusScheduled, AnticipationStatusPending:
		return true
	}
	return false
}

func (t TransferType) IsEnumValid() bool {
	switch t {
	case TransferTypeBankAccount, TransferTypeAsaasAccount:
		return true
	}
	return false
}

func (b BankAccountType) IsEnumValid() bool {
	switch b {
	case BankAccountTypeChecking, BankAccountTypeSavings:
		return true
	}
	return false
}

func (t TransferOperationType) IsEnumValid() bool {
	switch t {
	case TransferOperationTypePix, TransferOperationTypeTed, TransferOperationTypeInternal:
		return true
	}
	return false
}

func (t TransferStatus) IsEnumValid() bool {
	switch t {
	case TransferStatusPending, TransferStatusDone, TransferStatusCancelled, TransferStatusFailed,
		TransferStatusBankProcessing:
		return true
	}
	return false
}

func (p PersonType) IsEnumValid() bool {
	switch p {
	case PersonTypePhysical, PersonTypeJuridical:
		return true
	}
	return false
}

func (b AccountType) IsEnumValid() bool {
	switch b {
	case AccountTypeSalaryAccount, AccountTypeTypePaymentAccount, AccountTypeInvestimentAccount, AccountTypeCheckingAccount:
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
	case PixTransactionStatusAwaitingRequest, PixTransactionStatusDone, PixTransactionStatusRequested, PixTransactionStatusScheduled,
		PixTransactionStatusRefused, PixTransactionStatusError, PixTransactionStatusCancelled:
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
	case PixKeyStatusAwaitingActive, PixKeyStatusAwaitingActivation, PixKeyStatusAwaitingDeleted, PixKeyStatusAwaitingAccountDeletion,
		PixKeyStatusAwaitingDeletion, PixKeyStatusAwaitingError:
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
	case InvoiceStatusAuthorized, InvoiceStatusCanceled, InvoiceStatusCancellationDenied, InvoiceStatusProcessingCancellation,
		InvoiceStatusError, InvoiceStatusScheduled, InvoiceStatusSynchronized:
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
	case SortSubscriptionFieldDateCreated:
		return true
	}
	return false
}

func (s SubscriptionStatus) IsEnumValid() bool {
	switch s {
	case SubscriptionStatusActive, SubscriptionStatusInactive, SubscriptionStatusExpired:
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
	case SubscriptionCycleWeekly, SubscriptionCycleBiweekly, SubscriptionCycleMonthly, SubscriptionCycleBimonthly, SubscriptionCycleQuarterly, SubscriptionCycleSemiannually, SubscriptionCycleYearly:
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
	case ChargebackReason1, ChargebackReason2, ChargebackReason3, ChargebackReason4, ChargebackReason5,
		ChargebackReason6, ChargebackReason7, ChargebackReason8, ChargebackReason9, ChargebackReason10,
		ChargebackReason11, ChargebackReason12, ChargebackReason13, ChargebackReason14, ChargebackReason15,
		ChargebackReason16, ChargebackReason17, ChargebackReason18, ChargebackReason19,
		ChargebackReason20, ChargebackReason21, ChargebackReason22,
		ChargebackReason23, ChargebackReason24, ChargebackReason25, ChargebackReason26,
		ChargebackReason27, ChargebackReason28, ChargebackReason29, ChargebackReason30,
		ChargebackReason31, ChargebackReason32:
		return true
	}
	return false
}

func (c BillingType) IsEnumValid() bool {
	switch c {
	case BillingTypeBill, BillingTypeCreditCard, BillingTypeUndefined, BillingTypeDebitCard, BillingTypeTransfer, BillingTypeDeposit, BillingTypePix:
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
	case ChargebackReason1:
		return "Ausência de impressão"
	case ChargebackReason2:
		return "Fraude em ambiente de cartão não presente"
	case ChargebackReason3:
		return "Transação telefônica ativada por cartão"
	case ChargebackReason4:
		return "Fraude em ambiente de cartão presente"
	case ChargebackReason5:
		return "Boletim de negativação de cartões"
	case ChargebackReason6:
		return "Desacordo comercial"
	case ChargebackReason7:
		return "Cópia não atendida"
	case ChargebackReason8:
		return "Erro de apresentação de crédito / débito"
	case ChargebackReason9:
		return "Pagamento por outros meios"
	case ChargebackReason10:
		return "Sem autorização do portador do cartão"
	case ChargebackReason11:
		return "Valor da transação é diferente"
	case ChargebackReason12:
		return "Moeda inválida"
	case ChargebackReason13:
		return "Dados inválidos"
	case ChargebackReason14:
		return "Apresentação tardia"
	case ChargebackReason15:
		return "Contestação regulatória / legal local"
	case ChargebackReason16:
		return "ROCs múltiplos"
	case ChargebackReason17:
		return "Transação de crédito original não aceita"
	case ChargebackReason18:
		return "Outras fraudes - Cartão ausente"
	case ChargebackReason19:
		return "Erro de processamento"
	case ChargebackReason20:
		return "Cópia atendida ilegível / incompleta"
	case ChargebackReason21:
		return "Recorrência cancelada"
	case ChargebackReason22:
		return "Autorização requerida não obtida"
	case ChargebackReason23:
		return "Direito de regresso integral por fraude"
	case ChargebackReason24:
		return "Mercadoria / serviços cancelado"
	case ChargebackReason25:
		return "Mercadoria / serviço com defeito ou em desacordo"
	case ChargebackReason26:
		return "Mercadoria / serviços não recebidos"
	case ChargebackReason27:
		return "Desmembramento de venda"
	case ChargebackReason28:
		return "Transf. de responsabilidades diversas"
	case ChargebackReason29:
		return "Débito de aluguel de carro não qualificado"
	case ChargebackReason30:
		return "Contestação do portador de cartão (EUA)"
	case ChargebackReason31:
		return "Programa Visa de monitoramento de fraude"
	case ChargebackReason32:
		return "Arquivo boletim de advertência"
	}
	return ""
}

func (a Env) String() string {
	return []string{"SANDBOX", "PRODUCTION"}[a]
}

func (f FileMimeType) String() string {
	return string(f)
}

func (a Env) BaseUrl() string {
	return []string{"https://sandbox.asaas.com/api", "https://api.asaas.com"}[a]
}

func (f FileMimeType) Extension() string {
	switch f {
	case FileMimeTypePdf:
		return ".pdf"
	case FileMimeTypeText:
		return ".txt"
	case FileMimeTypeAvif:
		return ".avif"
	case FileMimeTypeCss:
		return ".css"
	case FileMimeTypeGif:
		return ".gif"
	case FileMimeTypeHtml:
		return ".html"
	case FileMimeTypeJpeg:
		return ".jpeg"
	case FileMimeTypeJs:
		return ".js"
	case FileMimeTypeJson:
		return ".json"
	case FileMimeTypePng:
		return ".png"
	case FileMimeTypeWasm:
		return ".wasm"
	case FileMimeTypeWebp:
		return ".webp"
	case FileMimeTypeXml:
		return ".xml"
	default:
		return ""
	}
}
