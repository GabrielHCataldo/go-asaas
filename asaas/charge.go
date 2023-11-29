package asaas

import (
	"context"
	"fmt"
	"net/http"
	"os"
)

type CreateChargeRequest struct {
	Customer             string                       `json:"customer,omitempty" validate:"required"`
	BillingType          BillingType                  `json:"billingType,omitempty" validate:"required,enum"`
	Value                float64                      `json:"value,omitempty" validate:"required"`
	DueDate              *Date                        `json:"dueDate,omitempty" validate:"required,after_now"`
	Description          string                       `json:"description,omitempty" validate:"omitempty,lte=500"`
	ExternalReference    string                       `json:"externalReference,omitempty"`
	Discount             *DiscountRequest             `json:"discount,omitempty"`
	Interest             *InterestRequest             `json:"interest,omitempty"`
	Fine                 *FineRequest                 `json:"fine,omitempty"`
	PostalService        bool                         `json:"postalService,omitempty"`
	Split                []SplitRequest               `json:"split,omitempty"`
	Callback             *ChargeCallbackRequest       `json:"callback,omitempty"`
	CreditCard           *CreditCardRequest           `json:"creditCard,omitempty"`
	CreditCardHolderInfo *CreditCardHolderInfoRequest `json:"creditCardHolderInfo,omitempty"`
	CreditCardToken      string                       `json:"creditCardToken,omitempty"`
	InstallmentCount     int                          `json:"installmentCount,omitempty" validate:"omitempty,gte=2"`
	InstallmentValue     float64                      `json:"installmentValue,omitempty" validate:"omitempty,gt=0"`
	AuthorizeOnly        bool                         `json:"authorizeOnly,omitempty"`
	RemoteIp             string                       `json:"remoteIp,omitempty"`
}

type UpdateChargeRequest struct {
	Customer          string                 `json:"customer,omitempty" validate:"required"`
	BillingType       BillingType            `json:"billingType,omitempty" validate:"required,enum"`
	Value             float64                `json:"value,omitempty" validate:"required,gt=0"`
	DueDate           Date                   `json:"dueDate,omitempty" validate:"required"`
	Description       string                 `json:"description,omitempty"`
	ExternalReference string                 `json:"externalReference,omitempty"`
	Discount          *DiscountRequest       `json:"discount,omitempty"`
	Interest          *InterestRequest       `json:"interest,omitempty"`
	Fine              *FineRequest           `json:"fine,omitempty"`
	PostalService     bool                   `json:"postalService,omitempty"`
	Split             []SplitRequest         `json:"split,omitempty"`
	Callback          *ChargeCallbackRequest `json:"callback,omitempty"`
	InstallmentCount  int                    `json:"installmentCount,omitempty" validate:"omitempty,gte=2"`
	InstallmentValue  float64                `json:"installmentValue,omitempty" validate:"omitempty,gt=0"`
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
	DateCreatedGe         *Date        `json:"dateCreated[ge],omitempty"`
	DateCreatedLe         *Date        `json:"dateCreated[le],omitempty"`
	EstimatedCreditDateGE *Date        `json:"estimatedCreditDate[ge],omitempty"`
	EstimatedCreditDateLE *Date        `json:"estimatedCreditDate[le],omitempty"`
	DueDateGE             *Date        `json:"dueDate[ge],omitempty"`
	DueDateLE             *Date        `json:"dueDate[le],omitempty"`
	User                  string       `json:"user,omitempty"`
	Offset                int          `json:"offset,omitempty"`
	Limit                 int          `json:"limit,omitempty"`
}

type ChargeReceiveInCashRequest struct {
	PaymentDate    *Date   `json:"paymentDate,omitempty" validate:"required"`
	Value          float64 `json:"value,omitempty" validate:"required,gt=0"`
	NotifyCustomer bool    `json:"notifyCustomer,omitempty"`
}

type UploadChargeDocumentRequest struct {
	AvailableAfterPayment bool         `json:"availableAfterPayment,omitempty"`
	Type                  DocumentType `json:"type,omitempty" validate:"required,enum"`
	File                  *os.File     `json:"file,omitempty" validate:"required"`
}

type UpdateChargeDocumentDefinitionsRequest struct {
	AvailableAfterPayment bool         `json:"availableAfterPayment,omitempty"`
	Type                  DocumentType `json:"type,omitempty" validate:"required,enum"`
}

type ChargeCallbackRequest struct {
	SuccessUrl   string `json:"successUrl,omitempty" validate:"required,url"`
	AutoRedirect bool   `json:"autoRedirect,omitempty"`
}

type ChargeResponse struct {
	Id                    string              `json:"id,omitempty"`
	Customer              string              `json:"customer,omitempty"`
	Status                ChargeStatus        `json:"status,omitempty"`
	PaymentLink           string              `json:"paymentLink,omitempty"`
	DueDate               *Date               `json:"dueDate,omitempty"`
	Value                 float64             `json:"value,omitempty"`
	NetValue              float64             `json:"netValue,omitempty"`
	BillingType           BillingType         `json:"billingType,omitempty"`
	CanBePaidAfterDueDate bool                `json:"canBePaidAfterDueDate,omitempty"`
	PixTransaction        string              `json:"pixTransaction,omitempty"`
	Description           string              `json:"description,omitempty"`
	ExternalReference     string              `json:"externalReference,omitempty"`
	OriginalValue         string              `json:"originalValue,omitempty"`
	InterestValue         string              `json:"interestValue,omitempty"`
	OriginalDueDate       *Date               `json:"originalDueDate,omitempty"`
	PaymentDate           *Date               `json:"paymentDate,omitempty"`
	ClientPaymentDate     *Date               `json:"clientPaymentDate,omitempty"`
	InstallmentNumber     int                 `json:"installmentCount,omitempty"`
	TransactionReceiptUrl string              `json:"transactionReceiptUrl,omitempty"`
	NossoNumero           string              `json:"nossoNumero,omitempty"`
	InvoiceUrl            string              `json:"invoiceUrl,omitempty"`
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

type ChargeStatusResponse struct {
	Status ChargeStatus `json:"status,omitempty"`
}

type IdentificationFieldResponse struct {
	IdentificationField string `json:"identificationField,omitempty"`
	NossoNumero         string `json:"nossoNumero,omitempty"`
	BarCode             string `json:"barCode,omitempty"`
}

type ChargePixQrCodeResponse struct {
	EncodedImage   string `json:"encodedImage,omitempty"`
	Payload        string `json:"payload,omitempty"`
	ExpirationDate *Date  `json:"expirationDate,omitempty"`
}

type ChargeDocumentResponse struct {
	Id                    string              `json:"id,omitempty"`
	Name                  string              `json:"name,omitempty"`
	AvailableAfterPayment bool                `json:"availableAfterPayment,omitempty"`
	Type                  DocumentType        `json:"type,omitempty"`
	File                  *ChargeFileResponse `json:"file,omitempty"`
	Deleted               bool                `json:"deleted,omitempty"`
	Errors                []ErrorResponse     `json:"errors,omitempty"`
}

type ChargeFileResponse struct {
	PublicId     string `json:"publicId,omitempty"`
	OriginalName string `json:"originalName,omitempty"`
	Size         int    `json:"size,omitempty"`
	Extension    string `json:"extension,omitempty"`
	PreviewUrl   string `json:"previewUrl,omitempty"`
	DownloadUrl  string `json:"downloadUrl,omitempty"`
}

type ChargeCreationLimitResponse struct {
	Creation ChargeCreationResponse `json:"creation,omitempty"`
}

type ChargeCreationResponse struct {
	Daily DailyCreationLimitResponse `json:"daily,omitempty"`
}

type DailyCreationLimitResponse struct {
	Limit      int  `json:"limit,omitempty"`
	Used       int  `json:"used,omitempty"`
	WasReached bool `json:"wasReached,omitempty"`
}

type charge struct {
	env         Env
	accessToken string
}

type Charge interface {
	Create(ctx context.Context, body CreateChargeRequest) (*ChargeResponse, Error)
	PayWithCreditCard(ctx context.Context, chargeId string, body CreditCardRequest) (*ChargeResponse, Error)
	UpdateById(ctx context.Context, chargeId string, body UpdateChargeRequest) (*ChargeResponse, Error)
	DeleteById(ctx context.Context, chargeId string) (*DeleteResponse, Error)
	RestoreById(ctx context.Context, chargeId string) (*ChargeResponse, Error)
	RefundById(ctx context.Context, chargeId string, body RefundRequest) (*ChargeResponse, Error)
	ReceiveInCashById(ctx context.Context, chargeId string, body ChargeReceiveInCashRequest) (*ChargeResponse, Error)
	UndoReceivedInCashById(ctx context.Context, chargeId string) (*ChargeResponse, Error)
	UploadDocumentById(ctx context.Context, chargeId string, body UploadChargeDocumentRequest) (*ChargeDocumentResponse, Error)
	UpdateDocumentDefinitionsById(ctx context.Context, chargeId, docId string, body UpdateChargeDocumentDefinitionsRequest) (
		*ChargeDocumentResponse, Error)
	DeleteDocumentById(ctx context.Context, chargeId, docId string) (*DeleteResponse, Error)
	GetById(ctx context.Context, chargeId string) (*ChargeResponse, Error)
	GetCreationLimit(ctx context.Context) (*ChargeCreationLimitResponse, Error)
	GetStatusById(ctx context.Context, chargeId string) (*ChargeStatus, Error)
	GetIdentificationFieldById(ctx context.Context, chargeId string) (*IdentificationFieldResponse, Error)
	GetPixQrCodeById(ctx context.Context, chargeId string) (*ChargePixQrCodeResponse, Error)
	GetDocumentById(ctx context.Context, chargeId, docId string) (*ChargeDocumentResponse, Error)
	GetAllDocumentsById(ctx context.Context, chargeId string, filter PageableDefaultRequest) (
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
		return nil, NewError(ErrorTypeValidation, err)
	}
	c.prepareCreateBodyRequest(&body)
	req := NewRequest[ChargeResponse](ctx, c.env, c.accessToken)
	return req.make(http.MethodPost, "/v3/payments", body)
}

func (c charge) PayWithCreditCard(ctx context.Context, chargeId string, body CreditCardRequest) (*ChargeResponse,
	Error) {
	if err := Validate().Struct(body); err != nil {
		return nil, NewError(ErrorTypeValidation, err)
	}
	req := NewRequest[ChargeResponse](ctx, c.env, c.accessToken)
	return req.make(http.MethodPost, fmt.Sprintf(`/v3/payments/%s/payWithCreditCard`, chargeId), body)
}

func (c charge) UpdateById(ctx context.Context, chargeId string, body UpdateChargeRequest) (*ChargeResponse,
	Error) {
	if err := Validate().Struct(body); err != nil {
		return nil, NewError(ErrorTypeValidation, err)
	}
	req := NewRequest[ChargeResponse](ctx, c.env, c.accessToken)
	return req.make(http.MethodPut, fmt.Sprintf(`/v3/payments/%s`, chargeId), body)
}

func (c charge) DeleteById(ctx context.Context, chargeId string) (*DeleteResponse, Error) {
	req := NewRequest[DeleteResponse](ctx, c.env, c.accessToken)
	return req.make(http.MethodDelete, fmt.Sprintf(`/v3/payments/%s`, chargeId), nil)
}

func (c charge) RestoreById(ctx context.Context, chargeId string) (*ChargeResponse, Error) {
	req := NewRequest[ChargeResponse](ctx, c.env, c.accessToken)
	return req.make(http.MethodPost, fmt.Sprintf(`/v3/payments/%s/restore`, chargeId), nil)
}

func (c charge) RefundById(ctx context.Context, chargeId string, body RefundRequest) (
	*ChargeResponse, Error) {
	req := NewRequest[ChargeResponse](ctx, c.env, c.accessToken)
	return req.make(http.MethodPost, fmt.Sprintf(`/v3/payments/%s/refund`, chargeId), body)
}

func (c charge) ReceiveInCashById(ctx context.Context, chargeId string, body ChargeReceiveInCashRequest) (
	*ChargeResponse, Error) {
	if err := Validate().Struct(body); err != nil {
		return nil, NewError(ErrorTypeValidation, err)
	}
	req := NewRequest[ChargeResponse](ctx, c.env, c.accessToken)
	return req.make(http.MethodPost, fmt.Sprintf(`/v3/payments/%s/receiveInCash`, chargeId), body)
}

func (c charge) UndoReceivedInCashById(ctx context.Context, chargeId string) (*ChargeResponse, Error) {
	req := NewRequest[ChargeResponse](ctx, c.env, c.accessToken)
	return req.make(http.MethodPost, fmt.Sprintf(`/v3/payments/%s/undoReceivedInCash`, chargeId), nil)
}

func (c charge) UploadDocumentById(ctx context.Context, chargeId string, body UploadChargeDocumentRequest) (
	*ChargeDocumentResponse, Error) {
	if err := Validate().Struct(body); err != nil {
		return nil, NewError(ErrorTypeValidation, err)
	}
	req := NewRequest[ChargeDocumentResponse](ctx, c.env, c.accessToken)
	return req.makeMultipartForm(http.MethodPost, fmt.Sprintf(`/v3/payments/%s/documents`, chargeId), body)
}

func (c charge) UpdateDocumentDefinitionsById(
	ctx context.Context,
	chargeId,
	docId string,
	body UpdateChargeDocumentDefinitionsRequest,
) (*ChargeDocumentResponse, Error) {
	if err := Validate().Struct(body); err != nil {
		return nil, NewError(ErrorTypeValidation, err)
	}
	req := NewRequest[ChargeDocumentResponse](ctx, c.env, c.accessToken)
	return req.make(http.MethodPut, fmt.Sprintf(`/v3/payments/%s/documents/%v`, chargeId, docId), body)
}

func (c charge) DeleteDocumentById(ctx context.Context, chargeId, docId string) (*DeleteResponse, Error) {
	req := NewRequest[DeleteResponse](ctx, c.env, c.accessToken)
	return req.make(http.MethodDelete, fmt.Sprintf(`/v3/payments/%s/documents/%v`, chargeId, docId), nil)
}

func (c charge) GetCreationLimit(ctx context.Context) (*ChargeCreationLimitResponse, Error) {
	req := NewRequest[ChargeCreationLimitResponse](ctx, c.env, c.accessToken)
	return req.make(http.MethodGet, "/v3/payments/limits", nil)
}

func (c charge) GetById(ctx context.Context, chargeId string) (*ChargeResponse, Error) {
	req := NewRequest[ChargeResponse](ctx, c.env, c.accessToken)
	return req.make(http.MethodGet, fmt.Sprintf(`/v3/payments/%s`, chargeId), nil)
}

func (c charge) GetStatusById(ctx context.Context, chargeId string) (*ChargeStatus, Error) {
	req := NewRequest[ChargeResponse](ctx, c.env, c.accessToken)
	resp, err := req.make(http.MethodGet, fmt.Sprintf(`/v3/payments/%s/status`, chargeId), nil)
	if err != nil {
		return nil, err
	}
	return &resp.Status, nil
}

func (c charge) GetIdentificationFieldById(ctx context.Context, chargeId string) (*IdentificationFieldResponse,
	Error) {
	req := NewRequest[IdentificationFieldResponse](ctx, c.env, c.accessToken)
	return req.make(http.MethodGet, fmt.Sprintf(`/v3/payments/%s/identificationField`, chargeId), nil)
}

func (c charge) GetPixQrCodeById(ctx context.Context, chargeId string) (*ChargePixQrCodeResponse, Error) {
	req := NewRequest[ChargePixQrCodeResponse](ctx, c.env, c.accessToken)
	return req.make(http.MethodGet, fmt.Sprintf(`/v3/payments/%s/pixQrCode`, chargeId), nil)
}

func (c charge) GetDocumentById(ctx context.Context, chargeId, docId string) (*ChargeDocumentResponse, Error) {
	req := NewRequest[ChargeDocumentResponse](ctx, c.env, c.accessToken)
	return req.make(http.MethodGet, fmt.Sprintf(`/v3/payments/%s/documents/%v`, chargeId, docId), nil)
}

func (c charge) GetAllDocumentsById(ctx context.Context, chargeId string, filter PageableDefaultRequest) (
	*Pageable[ChargeDocumentResponse], Error) {
	req := NewRequest[Pageable[ChargeDocumentResponse]](ctx, c.env, c.accessToken)
	return req.make(http.MethodGet, fmt.Sprintf(`/v3/payments/%s/documents`, chargeId), filter)
}

func (c charge) GetAll(ctx context.Context, filter GetAllChargesRequest) (
	*Pageable[ChargeResponse], Error) {
	req := NewRequest[Pageable[ChargeResponse]](ctx, c.env, c.accessToken)
	return req.make(http.MethodGet, "/v3/payments", filter)
}

func (c charge) validateCreateBodyRequest(body CreateChargeRequest) error {
	if err := Validate().Struct(body); err != nil {
		return err
	}
	return validateBillingBody(body.BillingType, body.CreditCard, body.CreditCardHolderInfo, body.CreditCardToken,
		body.RemoteIp)
}

func (c charge) prepareCreateBodyRequest(body *CreateChargeRequest) {
	switch body.BillingType {
	case BillingTypeCreditCard:
		if body.Fine != nil {
			body.Fine.DueDateLimitDays = 0
		}
		break
	default:
		body.CreditCard = nil
		body.CreditCardHolderInfo = nil
		body.CreditCardToken = ""
		body.RemoteIp = ""
	}
}
