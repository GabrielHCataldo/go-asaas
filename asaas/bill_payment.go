package asaas

import (
	"context"
	berrors "errors"
	"fmt"
	"github.com/GabrielHCataldo/go-asaas/internal/util"
	"net/http"
)

type BillPaymentRequest struct {
	IdentificationField string  `json:"identificationField,omitempty" validate:"required"`
	ScheduleDate        *Date   `json:"scheduleDate,omitempty" validate:"omitempty,after_now"`
	Description         string  `json:"description,omitempty"`
	Discount            float64 `json:"discount,omitempty" validate:"omitempty,gte=0"`
	Interest            float64 `json:"interest,omitempty" validate:"omitempty,gte=0"`
	Fine                float64 `json:"fine,omitempty" validate:"omitempty,gte=0"`
	Value               float64 `json:"value,omitempty" validate:"omitempty,gte=0"`
	DueDate             *Date   `json:"dueDate,omitempty"`
}

type BillPaymentSimulateRequest struct {
	IdentificationField string `json:"identificationField,omitempty"`
	BarCode             string `json:"barCode,omitempty"`
}

type BillPaymentResponse struct {
	Id                    string            `json:"id,omitempty"`
	IdentificationField   string            `json:"identificationField,omitempty"`
	Status                BillPaymentStatus `json:"status,omitempty"`
	Discount              float64           `json:"discount,omitempty"`
	Interest              float64           `json:"interest,omitempty"`
	Fine                  float64           `json:"fine,omitempty"`
	Value                 float64           `json:"value,omitempty"`
	Fee                   float64           `json:"fee,omitempty" `
	Description           string            `json:"description,omitempty"`
	CompanyName           string            `json:"companyName,omitempty"`
	TransactionReceiptUrl string            `json:"transactionReceiptUrl,omitempty"`
	CanBeCancelled        bool              `json:"canBeCancelled,omitempty"`
	FailReasons           string            `json:"failReasons,omitempty"`
	DueDate               *Date             `json:"dueDate,omitempty"`
	ScheduleDate          *Date             `json:"scheduleDate,omitempty"`
	PaymentDate           *Date             `json:"paymentDate,omitempty"`
	Errors                []ErrorResponse   `json:"errors,omitempty"`
}

type BillPaymentSimulateResponse struct {
	MinimumScheduleDate *Date                           `json:"minimumScheduleDate,omitempty"`
	Fee                 float64                         `json:"fee,omitempty" `
	BankSlipInfo        BillPaymentBankSlipInfoResponse `json:"bankSlipInfo,omitempty" `
	Errors              []ErrorResponse                 `json:"errors,omitempty"`
}

type BillPaymentBankSlipInfoResponse struct {
	IdentificationField  string  `json:"identificationField,omitempty"`
	Value                float64 `json:"value,omitempty"`
	DueDate              *Date   `json:"dueDate,omitempty"`
	CompanyName          string  `json:"companyName,omitempty"`
	Bank                 string  `json:"bank,omitempty"`
	BeneficiaryName      string  `json:"beneficiaryName,omitempty"`
	BeneficiaryCpfCnpj   string  `json:"beneficiaryCpfCnpj,omitempty"`
	AllowChangeValue     bool    `json:"allowChangeValue,omitempty"`
	MinValue             float64 `json:"minValue,omitempty"`
	MaxValue             float64 `json:"maxValue,omitempty"`
	DiscountValue        float64 `json:"discountValue,omitempty"`
	InterestValue        float64 `json:"interestValue,omitempty"`
	FineValue            float64 `json:"fineValue,omitempty"`
	OriginalValue        float64 `json:"originalValue,omitempty"`
	TotalDiscountValue   float64 `json:"totalDiscountValue,omitempty"`
	TotalAdditionalValue float64 `json:"totalAdditionalValue,omitempty"`
	IsOverdue            bool    `json:"isOverdue,omitempty"`
}

type billPayment struct {
	env         Env
	accessToken string
}

type BillPayment interface {
	Create(ctx context.Context, body BillPaymentRequest) (*BillPaymentResponse, Error)
	Simulate(ctx context.Context, body BillPaymentSimulateRequest) (*BillPaymentSimulateResponse, Error)
	CancelById(ctx context.Context, billPaymentId string) (*BillPaymentResponse, Error)
	GetById(ctx context.Context, billPaymentId string) (*BillPaymentResponse, Error)
	GetAll(ctx context.Context, filter PageableDefaultRequest) (*Pageable[BillPaymentResponse], Error)
}

func NewBillPayment(env Env, accessToken string) BillPayment {
	logWarning("BillPayment service running on", env.String())
	return billPayment{
		env:         env,
		accessToken: accessToken,
	}
}

func (b billPayment) Create(ctx context.Context, body BillPaymentRequest) (*BillPaymentResponse, Error) {
	if err := Validate().Struct(body); err != nil {
		return nil, NewError(ErrorTypeValidation, err)
	}
	req := NewRequest[BillPaymentResponse](ctx, b.env, b.accessToken)
	return req.make(http.MethodPost, "/v3/bill", body)
}

func (b billPayment) Simulate(ctx context.Context, body BillPaymentSimulateRequest) (*BillPaymentSimulateResponse, Error) {
	if err := b.validateBodySimulateRequest(&body.IdentificationField, &body.BarCode); err != nil {
		return nil, NewError(ErrorTypeValidation, "idenficationField ")
	}
	req := NewRequest[BillPaymentSimulateResponse](ctx, b.env, b.accessToken)
	return req.make(http.MethodPost, "/v3/bill/simulate", body)
}

func (b billPayment) CancelById(ctx context.Context, billPaymentId string) (*BillPaymentResponse, Error) {
	req := NewRequest[BillPaymentResponse](ctx, b.env, b.accessToken)
	return req.make(http.MethodPost, fmt.Sprintf("/v3/bill/%s/cancel", billPaymentId), nil)
}

func (b billPayment) GetById(ctx context.Context, billPaymentId string) (*BillPaymentResponse, Error) {
	req := NewRequest[BillPaymentResponse](ctx, b.env, b.accessToken)
	return req.make(http.MethodPost, fmt.Sprintf("/v3/bill/%s", billPaymentId), nil)
}

func (b billPayment) GetAll(ctx context.Context, filter PageableDefaultRequest) (*Pageable[BillPaymentResponse], Error) {
	req := NewRequest[Pageable[BillPaymentResponse]](ctx, b.env, b.accessToken)
	return req.make(http.MethodPost, "/v3/bill", filter)
}

func (b billPayment) validateBodySimulateRequest(identificationField, barCode *string) error {
	if util.IsBlank(barCode) && util.IsBlank(identificationField) {
		return berrors.New("inform barCode or identificationField")
	}
	return nil
}
