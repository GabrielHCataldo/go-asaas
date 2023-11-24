package asaas

import (
	"time"
)

type CreateChargeRequest struct {
	Customer             string                       `json:"customer,omitempty" validate:"required"`
	BillingType          BillingType                  `json:"billingType,omitempty" validate:"required,enum"`
	Value                float64                      `json:"value,omitempty" validate:"required"`
	DueDate              time.Time                    `json:"dueDate,omitempty" validate:"required"`
	Description          string                       `json:"description,omitempty"`
	ExternalReference    string                       `json:"externalReference,omitempty"`
	Discount             *DiscountRequest             `json:"discount,omitempty"`
	Interest             *InterestRequest             `json:"interest,omitempty"`
	Fine                 *FineRequest                 `json:"fine,omitempty"`
	PostalService        bool                         `json:"postalService,omitempty"`
	Split                []SplitRequest               `json:"split,omitempty"`
	Callback             *CallbackRequest             `json:"callback,omitempty"`
	CreditCard           *CreditCardRequest           `json:"creditCard,omitempty"`
	CreditCardHolderInfo *CreditCardHolderInfoRequest `json:"creditCardHolderInfo,omitempty"`
	CreditCardToken      string                       `json:"creditCardToken,omitempty"`
	InstallmentCount     int                          `json:"installmentCount,omitempty" validate:"omitempty,gte=2"`
	InstallmentValue     float64                      `json:"installmentValue,omitempty" validate:"omitempty,gt=0"`
	AuthorizeOnly        bool                         `json:"authorizeOnly,omitempty"`
	RemoteIP             string                       `json:"remoteIP,omitempty" validate:"required,ip"`
}

type CreditCardRequest struct {
	HolderName  string `json:"holderName,omitempty" validate:"required,full-name"`
	Number      string `json:"number,omitempty" validate:"required,numeric,min=10,max=19"`
	ExpiryMonth string `json:"expiryMonth,omitempty" validate:"required,numeric,len=2"`
	ExpiryYear  string `json:"expiryYear,omitempty" validate:"required,numeric,len=4"`
	CVV         string `json:"cvv,omitempty" validate:"required,numeric,min=3,max=4"`
}

type CreditCardHolderInfoRequest struct {
	Name              string `json:"name,omitempty" validate:"required,full-name"`
	CpfCnpj           string `json:"cpfCnpj,omitempty" validate:"required,document"`
	Email             string `json:"email,omitempty" validate:"required,email"`
	Phone             string `json:"phone,omitempty" validate:"required,phone"`
	PostalCode        string `json:"postalCode,omitempty" validate:"required,postal-code"`
	AddressNumber     string `json:"addressNumber,omitempty" validate:"required,numeric"`
	AddressComplement string `json:"addressComplement,omitempty"`
}

type DiscountRequest struct {
	Value            float64      `json:"value,omitempty" validate:"required"`
	DueDateLimitDays int          `json:"dueDateLimitDays,omitempty" validate:"gte=0"`
	Type             DiscountType `json:"type,omitempty" validate:"required,enum"`
}

type InterestRequest struct {
	Value float64 `json:"value,omitempty" validate:"required"`
}

type FineRequest struct {
	Value            float64  `json:"value,omitempty" validate:"required"`
	DueDateLimitDays int      `json:"dueDateLimitDays,omitempty" validate:"omitempty,gte=0"`
	Type             FineType `json:"type,omitempty" validate:"required,enum"`
}

type SplitRequest struct {
	WalletID        string  `json:"walletId,omitempty" validate:"required"`
	FixedValue      float64 `json:"fixedValue,omitempty" validate:"omitempty,gt=0"`
	PercentualValue float64 `json:"percentualValue,omitempty" validate:"omitempty,gt=0"`
	TotalFixedValue float64 `json:"totalFixedValue,omitempty" validate:"omitempty,gt=0"`
}

type CallbackRequest struct {
	SuccessURL   string `json:"successUrl,omitempty" validate:"required,url"`
	AutoRedirect bool   `json:"autoRedirect,omitempty"`
}

type CreateChargeResponse struct {
	ID                    string            `json:"id,omitempty"`
	Customer              string            `json:"customer,omitempty"`
	CreditCardToken       string            `json:"creditCardToken,omitempty"`
	PaymentLink           string            `json:"paymentLink,omitempty"`
	DueDate               time.Time         `json:"dueDate,omitempty"`
	Value                 float64           `json:"value,omitempty"`
	NetValue              float64           `json:"netValue,omitempty"`
	BillingType           BillingType       `json:"billingType,omitempty"`
	CanBePaidAfterDueDate bool              `json:"canBePaidAfterDueDate,omitempty"`
	PixTransaction        string            `json:"pixTransaction,omitempty"`
	Status                ChargeStatus      `json:"status,omitempty"`
	Description           string            `json:"description,omitempty"`
	ExternalReference     string            `json:"externalReference,omitempty"`
	OriginalValue         string            `json:"originalValue,omitempty"`
	InterestValue         string            `json:"interestValue,omitempty"`
	OriginalDueDate       time.Time         `json:"originalDueDate,omitempty"`
	PaymentDate           time.Time         `json:"paymentDate,omitempty"`
	ClientPaymentDate     time.Time         `json:"clientPaymentDate,omitempty"`
	InstallmentNumber     int               `json:"installmentCount,omitempty"`
	TransactionReceiptURL string            `json:"transactionReceiptUrl,omitempty"`
	NossoNumero           string            `json:"nossoNumero,omitempty"`
	InvoiceURL            string            `json:"invoiceUrl,omitempty"`
	BankSlipUrl           string            `json:"bankSlipUrl,omitempty"`
	InvoiceNumber         string            `json:"invoiceNumber,omitempty"`
	Discount              *DiscountResponse `json:"discount,omitempty"`
	Fine                  *FineResponse     `json:"fine,omitempty"`
	Interest              *InterestResponse `json:"interest,omitempty"`
	Deleted               bool              `json:"deleted,omitempty"`
	PostalService         bool              `json:"postalService,omitempty"`
	Anticipated           bool              `json:"anticipated,omitempty"`
	Anticipable           bool              `json:"anticipable,omitempty"`
	Refunds               []RefundResponse  `json:"refunds,omitempty"`
	Errors                []ErrorResponse   `json:"errors,omitempty"`
	DateCreated           time.Time         `json:"dateCreated,omitempty"`
}

type DiscountResponse struct {
	Value            float64 `json:"value,omitempty"`
	DueDateLimitDays int     `json:"dueDateLimitDays,omitempty"`
}

type InterestResponse struct {
	Value float64 `json:"value,omitempty"`
}

type FineResponse struct {
	Value float64 `json:"value,omitempty"`
}

type RefundResponse struct {
	Status                RefundStatus `json:"status,omitempty"`
	Value                 float64      `json:"value,omitempty"`
	Description           string       `json:"description,omitempty"`
	TransactionReceiptURL string       `json:"transactionReceiptUrl,omitempty"`
	DateCreated           time.Time    `json:"dateCreated,omitempty"`
}
