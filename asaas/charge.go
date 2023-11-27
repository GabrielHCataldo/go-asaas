package asaas

import (
	"context"
	berrors "errors"
	"fmt"
	"github.com/GabrielHCataldo/go-asaas/internal/util"
	"net/http"
	"os"
	"time"
)

type CreateChargeRequest struct {
	Customer             string                       `json:"customer,omitempty" validate:"required"`
	BillingType          BillingType                  `json:"billingType,omitempty" validate:"required,enum"`
	Value                float64                      `json:"value,omitempty" validate:"required"`
	DueDate              Date                         `json:"dueDate,omitempty" validate:"required"`
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
	RemoteIP             string                       `json:"remoteIp,omitempty" validate:"required,ip"`
}

type UpdateChargeRequest struct {
	Customer          string           `json:"customer,omitempty" validate:"required"`
	BillingType       BillingType      `json:"billingType,omitempty" validate:"required,enum"`
	Value             float64          `json:"value,omitempty" validate:"required,gt=0"`
	DueDate           Date             `json:"dueDate,omitempty" validate:"required"`
	Description       string           `json:"description,omitempty"`
	ExternalReference string           `json:"externalReference,omitempty"`
	Discount          *DiscountRequest `json:"discount,omitempty"`
	Interest          *InterestRequest `json:"interest,omitempty"`
	Fine              *FineRequest     `json:"fine,omitempty"`
	PostalService     bool             `json:"postalService,omitempty"`
	Split             []SplitRequest   `json:"split,omitempty"`
	Callback          *CallbackRequest `json:"callback,omitempty"`
	InstallmentCount  int              `json:"installmentCount,omitempty" validate:"omitempty,gte=2"`
	InstallmentValue  float64          `json:"installmentValue,omitempty" validate:"omitempty,gt=0"`
}

type GetAllChargesRequest struct {
	Customer              string       `json:"customer,omitempty"`
	CustomerGroupName     string       `json:"customerGroupName,omitempty"`
	BillingType           BillingType  `json:"billingType,omitempty"`
	Status                ChargeStatus `json:"status,omitempty"`
	Subscription          string       `json:"subscription,omitempty"`
	Installment           string       `json:"installment,omitempty"`
	ExternalReference     string       `json:"externalReference,omitempty"`
	InvoiceStatus         string       `json:"invoiceStatus,omitempty"`
	EstimatedCreditDate   *Date        `json:"estimatedCreditDate,omitempty"`
	PixQrCodeId           string       `json:"pixQrCodeId,omitempty"`
	Anticipated           bool         `json:"anticipated,omitempty"`
	DateCreatedGE         *Date        `json:"dateCreated[ge],omitempty"`
	DateCreatedLE         *Date        `json:"dateCreated[le],omitempty"`
	EstimatedCreditDateGE *Date        `json:"estimatedCreditDate[ge],omitempty"`
	EstimatedCreditDateLE *Date        `json:"estimatedCreditDate[le],omitempty"`
	DueDateGE             *Date        `json:"dueDate[ge],omitempty"`
	DueDateLE             *Date        `json:"dueDate[le],omitempty"`
	User                  string       `json:"user,omitempty"`
	Offset                int          `json:"offset,omitempty"`
	Limit                 int          `json:"limit,omitempty" validate:"omitempty,lte=100"`
}

type CreditCardRequest struct {
	HolderName  string `json:"holderName,omitempty" validate:"required,full_name"`
	Number      string `json:"number,omitempty" validate:"required,numeric,min=10,max=19"`
	ExpiryMonth string `json:"expiryMonth,omitempty" validate:"required,numeric,len=2"`
	ExpiryYear  string `json:"expiryYear,omitempty" validate:"required,numeric,len=4"`
	CCV         string `json:"ccv,omitempty" validate:"required,numeric,min=3,max=4"`
}

type RefundChargeRequest struct {
	Value       float64 `json:"value,omitempty"`
	Description string  `json:"description,omitempty"`
}

type ReceiveInCashRequest struct {
	PaymentDate    Date    `json:"paymentDate,omitempty" validate:"required"`
	Value          float64 `json:"value,omitempty" validate:"required,gt=0"`
	NotifyCustomer bool    `json:"notifyCustomer,omitempty"`
}

type UploadDocumentRequest struct {
	AvailableAfterPayment bool           `json:"availableAfterPayment,omitempty"`
	Type                  TypeOfDocument `json:"type,omitempty" validate:"required,enum"`
	File                  *os.File       `json:"file,omitempty" validate:"required"`
}

type UpdateDocumentDefinitionsRequest struct {
	AvailableAfterPayment bool           `json:"availableAfterPayment,omitempty"`
	Type                  TypeOfDocument `json:"type,omitempty" validate:"required,enum"`
}

type CreditCardHolderInfoRequest struct {
	Name              string `json:"name,omitempty" validate:"required,full_name"`
	CpfCnpj           string `json:"cpfCnpj,omitempty" validate:"required,document"`
	Email             string `json:"email,omitempty" validate:"required,email"`
	Phone             string `json:"phone,omitempty" validate:"required,phone"`
	MobilePhone       string `json:"mobilePhone,omitempty" validate:"omitempty,phone"`
	PostalCode        string `json:"postalCode,omitempty" validate:"required,postal_code"`
	AddressNumber     string `json:"addressNumber,omitempty" validate:"required,numeric"`
	AddressComplement string `json:"addressComplement,omitempty"`
}

type DiscountRequest struct {
	Value            float64      `json:"value,omitempty" validate:"required,gt=0"`
	DueDateLimitDays int          `json:"dueDateLimitDays,omitempty" validate:"gte=0"`
	Type             DiscountType `json:"type,omitempty" validate:"required,enum"`
}

type InterestRequest struct {
	Value float64 `json:"value,omitempty" validate:"required,gt=0"`
}

type FineRequest struct {
	Value            float64  `json:"value,omitempty" validate:"required,gt=0"`
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

type ChargeResponse struct {
	ID                    string              `json:"id,omitempty"`
	Customer              string              `json:"customer,omitempty"`
	PaymentLink           string              `json:"paymentLink,omitempty"`
	DueDate               *Date               `json:"dueDate,omitempty"`
	Value                 float64             `json:"value,omitempty"`
	NetValue              float64             `json:"netValue,omitempty"`
	BillingType           BillingType         `json:"billingType,omitempty"`
	CanBePaidAfterDueDate bool                `json:"canBePaidAfterDueDate,omitempty"`
	PixTransaction        string              `json:"pixTransaction,omitempty"`
	Status                ChargeStatus        `json:"status,omitempty"`
	Description           string              `json:"description,omitempty"`
	ExternalReference     string              `json:"externalReference,omitempty"`
	OriginalValue         string              `json:"originalValue,omitempty"`
	InterestValue         string              `json:"interestValue,omitempty"`
	OriginalDueDate       *Date               `json:"originalDueDate,omitempty"`
	PaymentDate           *Date               `json:"paymentDate,omitempty"`
	ClientPaymentDate     *Date               `json:"clientPaymentDate,omitempty"`
	InstallmentNumber     int                 `json:"installmentCount,omitempty"`
	TransactionReceiptURL string              `json:"transactionReceiptUrl,omitempty"`
	NossoNumero           string              `json:"nossoNumero,omitempty"`
	InvoiceURL            string              `json:"invoiceUrl,omitempty"`
	BankSlipUrl           string              `json:"bankSlipUrl,omitempty"`
	InvoiceNumber         string              `json:"invoiceNumber,omitempty"`
	CreditCard            *CreditCardResponse `json:"creditCard,omitempty"`
	Discount              *DiscountResponse   `json:"discount,omitempty"`
	Fine                  *FineResponse       `json:"fine,omitempty"`
	Interest              *InterestResponse   `json:"interest,omitempty"`
	Deleted               bool                `json:"deleted,omitempty"`
	PostalService         bool                `json:"postalService,omitempty"`
	Anticipated           bool                `json:"anticipated,omitempty"`
	Anticipable           bool                `json:"anticipable,omitempty"`
	Refunds               []RefundResponse    `json:"refunds,omitempty"`
	Errors                []ErrorResponse     `json:"errors,omitempty"`
	DateCreated           *Date               `json:"dateCreated,omitempty"`
}

type ChargeDeleteResponse struct {
	ID      string `json:"id,omitempty"`
	Deleted bool   `json:"deleted,omitempty"`
}

type ChargeStatusResponse struct {
	Status ChargeStatus `json:"status,omitempty"`
}

type CreditCardResponse struct {
	CreditCardNumber string `json:"creditCardNumber,omitempty"`
	CreditCardBrand  string `json:"creditCardBrand,omitempty"`
	CreditCardToken  string `json:"creditCardToken,omitempty"`
}

type DiscountResponse struct {
	Value            float64      `json:"value,omitempty"`
	DueDateLimitDays int          `json:"dueDateLimitDays,omitempty"`
	Type             DiscountType `json:"type,omitempty"`
}

type InterestResponse struct {
	Value float64      `json:"value,omitempty"`
	Type  InterestType `json:"type,omitempty"`
}

type FineResponse struct {
	Value float64  `json:"value,omitempty"`
	Type  FineType `json:"type,omitempty"`
}

type RefundResponse struct {
	Status                RefundStatus `json:"status,omitempty"`
	Value                 float64      `json:"value,omitempty"`
	Description           string       `json:"description,omitempty"`
	TransactionReceiptURL string       `json:"transactionReceiptUrl,omitempty"`
	DateCreated           Date         `json:"dateCreated,omitempty"`
}

type IdentificationFieldResponse struct {
	IdentificationField string `json:"identificationField,omitempty"`
	NossoNumero         string `json:"nossoNumero,omitempty"`
	BarCode             string `json:"barCode,omitempty"`
}

type PixQRCodeResponse struct {
	EncodedImage   string `json:"encodedImage,omitempty"`
	Payload        string `json:"payload,omitempty"`
	ExpirationDate Date   `json:"expirationDate,omitempty"`
}

type ChargeDocumentResponse struct {
	ID                    string              `json:"id,omitempty"`
	Name                  string              `json:"name,omitempty"`
	AvailableAfterPayment bool                `json:"availableAfterPayment,omitempty"`
	Type                  TypeOfDocument      `json:"type,omitempty"`
	File                  *ChargeFileResponse `json:"file,omitempty"`
	Deleted               bool                `json:"deleted,omitempty"`
	Errors                []ErrorResponse     `json:"errors,omitempty"`
}

type ChargeFileResponse struct {
	PublicID     string `json:"publicId,omitempty"`
	OriginalName string `json:"originalName,omitempty"`
	Size         int    `json:"size,omitempty"`
	Extension    string `json:"extension,omitempty"`
	PreviewURL   string `json:"previewUrl,omitempty"`
	DownloadURL  string `json:"downloadUrl,omitempty"`
}

type charge struct {
	env         Env
	accessToken string
}

type Charge interface {
	Create(ctx context.Context, body CreateChargeRequest) (*ChargeResponse, Error)
	PayWithCreditCard(ctx context.Context, chargeID string, body CreditCardRequest) (*ChargeResponse, Error)
	UpdateByID(ctx context.Context, chargeID string, body UpdateChargeRequest) (*ChargeResponse, Error)
	DeleteByID(ctx context.Context, chargeID string) (*ChargeDeleteResponse, Error)
	RestoreByID(ctx context.Context, chargeID string) (*ChargeResponse, Error)
	RefundByID(ctx context.Context, chargeID string, body RefundChargeRequest) (*ChargeResponse, Error)
	ReceiveInCashByID(ctx context.Context, chargeID string, body ReceiveInCashRequest) (*ChargeResponse, Error)
	UndoReceivedInCashByID(ctx context.Context, chargeID string) (*ChargeResponse, Error)
	UploadDocumentByID(ctx context.Context, chargeID string, body UploadDocumentRequest) (*ChargeDocumentResponse, Error)
	UpdateDocumentDefinitionsByID(ctx context.Context, chargeID, docID string, body UpdateDocumentDefinitionsRequest) (
		*ChargeDocumentResponse, Error)
	GetByID(ctx context.Context, chargeID string) (*ChargeResponse, Error)
	GetStatusByID(ctx context.Context, chargeID string) (*ChargeStatus, Error)
	GetIdentificationFieldByID(ctx context.Context, chargeID string) (*IdentificationFieldResponse, Error)
	GetPixQRCodeByID(ctx context.Context, chargeID string) (*PixQRCodeResponse, Error)
	GetDocumentByID(ctx context.Context, chargeID, docID string) (*ChargeDocumentResponse, Error)
	GetAllDocumentsByID(ctx context.Context, chargeID string, filter PageableDefaultRequest) (
		*Pageable[ChargeDocumentResponse], Error)
	GetAll(ctx context.Context, filter GetAllChargesRequest) (*Pageable[ChargeResponse], Error)
}

func NewCharge(env Env, accessCode string) Charge {
	logWarning("Charge service running on", env.String())
	return charge{
		env:         env,
		accessToken: accessCode,
	}
}

func (c charge) Create(ctx context.Context, body CreateChargeRequest) (*ChargeResponse, Error) {
	if err := c.validateCreateBodyRequest(body); err != nil {
		return nil, NewError(ERROR_VALIDATION, err)
	}
	c.prepareCreateBodyRequest(&body)
	req := NewRequest[ChargeResponse](ctx, c.env, c.accessToken)
	return req.make(http.MethodPost, "/v3/payments", body)
}

func (c charge) PayWithCreditCard(ctx context.Context, chargeID string, body CreditCardRequest) (*ChargeResponse,
	Error) {
	if err := Validate().Struct(body); err != nil {
		return nil, NewError(ERROR_VALIDATION, err)
	}
	req := NewRequest[ChargeResponse](ctx, c.env, c.accessToken)
	return req.make(http.MethodPost, fmt.Sprintf(`/v3/payments/%s/payWithCreditCard`, chargeID), body)
}

func (c charge) UpdateByID(ctx context.Context, chargeID string, body UpdateChargeRequest) (*ChargeResponse,
	Error) {
	if err := Validate().Struct(body); err != nil {
		return nil, NewError(ERROR_VALIDATION, err)
	}
	req := NewRequest[ChargeResponse](ctx, c.env, c.accessToken)
	return req.make(http.MethodPut, fmt.Sprintf(`/v3/payments/%s`, chargeID), body)
}

func (c charge) DeleteByID(ctx context.Context, chargeID string) (*ChargeDeleteResponse, Error) {
	req := NewRequest[ChargeDeleteResponse](ctx, c.env, c.accessToken)
	return req.make(http.MethodDelete, fmt.Sprintf(`/v3/payments/%s`, chargeID), nil)
}

func (c charge) RestoreByID(ctx context.Context, chargeID string) (*ChargeResponse, Error) {
	req := NewRequest[ChargeResponse](ctx, c.env, c.accessToken)
	return req.make(http.MethodPost, fmt.Sprintf(`/v3/payments/%s/restore`, chargeID), nil)
}

func (c charge) RefundByID(ctx context.Context, chargeID string, body RefundChargeRequest) (
	*ChargeResponse, Error) {
	req := NewRequest[ChargeResponse](ctx, c.env, c.accessToken)
	return req.make(http.MethodPost, fmt.Sprintf(`/v3/payments/%s/refund`, chargeID), body)
}

func (c charge) ReceiveInCashByID(ctx context.Context, chargeID string, body ReceiveInCashRequest) (
	*ChargeResponse, Error) {
	if err := Validate().Struct(body); err != nil {
		return nil, NewError(ERROR_VALIDATION, err)
	}
	req := NewRequest[ChargeResponse](ctx, c.env, c.accessToken)
	return req.make(http.MethodPost, fmt.Sprintf(`/v3/payments/%s/receiveInCash`, chargeID), body)
}

func (c charge) UndoReceivedInCashByID(ctx context.Context, chargeID string) (*ChargeResponse, Error) {
	req := NewRequest[ChargeResponse](ctx, c.env, c.accessToken)
	return req.make(http.MethodPost, fmt.Sprintf(`/v3/payments/%s/undoReceivedInCash`, chargeID), nil)
}

func (c charge) UploadDocumentByID(ctx context.Context, chargeID string, body UploadDocumentRequest) (
	*ChargeDocumentResponse, Error) {
	if err := Validate().Struct(body); err != nil {
		return nil, NewError(ERROR_VALIDATION, err)
	}
	req := NewRequest[ChargeDocumentResponse](ctx, c.env, c.accessToken)
	return req.makeMultipartForm(http.MethodPost, fmt.Sprintf(`/v3/payments/%s/documents`, chargeID), body)
}

func (c charge) UpdateDocumentDefinitionsByID(
	ctx context.Context,
	chargeID,
	docID string,
	body UpdateDocumentDefinitionsRequest,
) (*ChargeDocumentResponse, Error) {
	if err := Validate().Struct(body); err != nil {
		return nil, NewError(ERROR_VALIDATION, err)
	}
	req := NewRequest[ChargeDocumentResponse](ctx, c.env, c.accessToken)
	return req.make(http.MethodPut, fmt.Sprintf(`/v3/payments/%s/documents/%v`, chargeID, docID), body)
}

func (c charge) GetByID(ctx context.Context, chargeID string) (*ChargeResponse, Error) {
	req := NewRequest[ChargeResponse](ctx, c.env, c.accessToken)
	return req.make(http.MethodGet, fmt.Sprintf(`/v3/payments/%s`, chargeID), nil)
}

func (c charge) GetStatusByID(ctx context.Context, chargeID string) (*ChargeStatus, Error) {
	req := NewRequest[ChargeResponse](ctx, c.env, c.accessToken)
	resp, err := req.make(http.MethodGet, fmt.Sprintf(`/v3/payments/%s/status`, chargeID), nil)
	if err != nil {
		return nil, err
	}
	return &resp.Status, nil
}

func (c charge) GetIdentificationFieldByID(ctx context.Context, chargeID string) (*IdentificationFieldResponse,
	Error) {
	req := NewRequest[IdentificationFieldResponse](ctx, c.env, c.accessToken)
	return req.make(http.MethodGet, fmt.Sprintf(`/v3/payments/%s/identificationField`, chargeID), nil)
}

func (c charge) GetPixQRCodeByID(ctx context.Context, chargeID string) (*PixQRCodeResponse, Error) {
	req := NewRequest[PixQRCodeResponse](ctx, c.env, c.accessToken)
	return req.make(http.MethodGet, fmt.Sprintf(`/v3/payments/%s/pixQrCode`, chargeID), nil)
}

func (c charge) GetDocumentByID(ctx context.Context, chargeID, docID string) (*ChargeDocumentResponse, Error) {
	req := NewRequest[ChargeDocumentResponse](ctx, c.env, c.accessToken)
	return req.make(http.MethodGet, fmt.Sprintf(`/v3/payments/%s/documents/%v`, chargeID, docID), nil)
}

func (c charge) GetAllDocumentsByID(ctx context.Context, chargeID string, filter PageableDefaultRequest) (
	*Pageable[ChargeDocumentResponse], Error) {
	req := NewRequest[Pageable[ChargeDocumentResponse]](ctx, c.env, c.accessToken)
	return req.make(http.MethodGet, fmt.Sprintf(`/v3/payments/%s/documents`, chargeID), filter)
}

func (c charge) GetAll(ctx context.Context, filter GetAllChargesRequest) (
	*Pageable[ChargeResponse], Error) {
	if err := Validate().Struct(filter); err != nil {
		return nil, NewError(ERROR_VALIDATION, err)
	}
	req := NewRequest[Pageable[ChargeResponse]](ctx, c.env, c.accessToken)
	return req.make(http.MethodGet, "/v3/payments", filter)
}

func (c charge) validateCreateBodyRequest(body CreateChargeRequest) error {
	if err := Validate().Struct(body); err != nil {
		return err
	} else {
		dueDate := time.Date(body.DueDate.Time().Year(), body.DueDate.Month(), body.DueDate.Day(), 23, 59, 0, 0,
			body.DueDate.Location())
		if time.Now().UTC().After(dueDate.UTC()) {
			return berrors.New("invalid due date")
		}
	}
	switch body.BillingType {
	case CREDIT_CARD:
		cCard := body.CreditCard
		cCardHolderInfoBody := body.CreditCardHolderInfo
		cCardToken := body.CreditCardToken
		if util.IsBlank(&cCardToken) && (cCard == nil || cCardHolderInfoBody == nil) {
			return berrors.New("to charge by credit card, enter the credit card or credit card token")
		} else if cCard != nil && !util.ValidateExpirationCreditCard(cCard.ExpiryYear, cCard.ExpiryMonth) {
			return berrors.New("expired card")
		}
		break
	}
	return nil
}

func (c charge) prepareCreateBodyRequest(body *CreateChargeRequest) {
	body.DueDate = NewDate(body.DueDate.Year(), body.DueDate.Month(), body.DueDate.Day(),
		23, 59, 0, 0, body.DueDate.Location())
	switch body.BillingType {
	case CREDIT_CARD:
		if body.Fine != nil {
			body.Fine.DueDateLimitDays = 0
		}
		break
	default:
		body.CreditCard = nil
		body.CreditCardHolderInfo = nil
		body.CreditCardToken = ""
	}
}
