package asaas

import (
	"context"
	"fmt"
	"github.com/GabrielHCataldo/go-asaas/internal/util"
	"net/http"
	"os"
)

type CreateNegativityRequest struct {
	Payment                string         `json:"payment,omitempty" validate:"required"`
	Type                   NegativityType `json:"type,omitempty" validate:"required,enum"`
	Description            string         `json:"description,omitempty"`
	CustomerName           string         `json:"customerName,omitempty" validate:"required,full_name"`
	CustomerCpfCnpj        string         `json:"customerCpfCnpj,omitempty" validate:"required,document"`
	CustomerPrimaryPhone   string         `json:"customerPrimaryPhone,omitempty" validate:"required,phone"`
	CustomerSecondaryPhone string         `json:"customerSecondaryPhone,omitempty" validate:"omitempty,phone"`
	CustomerPostalCode     string         `json:"customerPostalCode,omitempty" validate:"required,postal_code"`
	CustomerAddress        string         `json:"customerAddress,omitempty" validate:"required"`
	CustomerAddressNumber  string         `json:"customerAddressNumber,omitempty" validate:"required"`
	CustomerComplement     string         `json:"customerComplement,omitempty"`
	CustomerProvince       string         `json:"customerProvince,omitempty" validate:"required"`
	Documents              *FileRequest   `json:"documents,omitempty"`
}

type GetAllNegativitiesRequest struct {
	Status           NegativityStatus `json:"status,omitempty"`
	Type             NegativityType   `json:"type,omitempty"`
	Payment          string           `json:"payment,omitempty"`
	RequestStartDate *Date            `json:"requestStartDate,omitempty"`
}

type NegativityResendDocumentsRequest struct {
	Documents *os.File `json:"documents,omitempty" validate:"required"`
}

type NegativityResponse struct {
	Id                             string           `json:"id,omitempty"`
	Payment                        string           `json:"payment,omitempty"`
	DunningNumber                  int              `json:"dunningNumber,omitempty"`
	Type                           NegativityType   `json:"type,omitempty"`
	Status                         NegativityStatus `json:"status,omitempty"`
	RequestDate                    *Date            `json:"requestDate,omitempty"`
	Description                    string           `json:"description,omitempty"`
	Value                          float64          `json:"value,omitempty"`
	FeeValue                       float64          `json:"feeValue,omitempty"`
	NetValue                       float64          `json:"netValue,omitempty"`
	ReceivedInCashFeeValue         float64          `json:"receivedInCashFeeValue,omitempty"`
	CancellationFeeValue           float64          `json:"cancellationFeeValue,omitempty"`
	DenialReason                   string           `json:"denialReason,omitempty"`
	CanBeCancelled                 bool             `json:"canBeCancelled,omitempty"`
	CannotBeCancelledReason        string           `json:"cannotBeCancelledReason,omitempty"`
	IsNecessaryResendDocumentation bool             `json:"isNecessaryResendDocumentation,omitempty"`
	Errors                         []ErrorResponse  `json:"errors,omitempty"`
}

type NegativitySimulateResponse struct {
	Payment         string                             `json:"payment,omitempty"`
	Value           float64                            `json:"value,omitempty"`
	TypeSimulations []NegativityTypeSimulationResponse `json:"typeSimulations,omitempty"`
}

type NegativityTypeSimulationResponse struct {
	Type             NegativityType `json:"type,omitempty"`
	IsAllowed        bool           `json:"isAllowed,omitempty"`
	NotAllowedReason string         `json:"notAllowedReason,omitempty"`
	FeeValue         float64        `json:"feeValue,omitempty"`
	NetValue         float64        `json:"netValue,omitempty"`
	StartDate        *Date          `json:"startDate,omitempty"`
}

type NegativityHistoryResponse struct {
	Status      NegativityStatus `json:"status,omitempty"`
	Description string           `json:"description,omitempty"`
	EventDate   *Date            `json:"eventDate,omitempty"`
}

type NegativityPaymentsResponse struct {
	Value       float64 `json:"value,omitempty"`
	Description string  `json:"description,omitempty"`
	PaymentDate *Date   `json:"paymentDate,omitempty"`
}

type ChargesAvailableForDunningResponse struct {
	Payment         string                             `json:"payment,omitempty"`
	Customer        string                             `json:"customer,omitempty"`
	BillingType     BillingType                        `json:"billingType,omitempty"`
	Status          ChargeStatus                       `json:"status,omitempty"`
	Value           float64                            `json:"value,omitempty"`
	DueDate         *Date                              `json:"dueDate,omitempty"`
	TypeSimulations []NegativityTypeSimulationResponse `json:"typeSimulations,omitempty"`
}

type negativity struct {
	env         Env
	accessToken string
}

type Negativity interface {
	Create(ctx context.Context, body CreateNegativityRequest) (*NegativityResponse, Error)
	Simulate(ctx context.Context, chargeId string) (*NegativitySimulateResponse, Error)
	ResendDocumentsById(ctx context.Context, negativityId string, body NegativityResendDocumentsRequest) (
		*NegativityResponse, Error)
	CancelById(ctx context.Context, negativityId string) (*NegativityResponse, Error)
	GetById(ctx context.Context, negativityId string) (*NegativityResponse, Error)
	GetHistoryById(ctx context.Context, negativityId string, filter PageableDefaultRequest) (
		*Pageable[NegativityHistoryResponse], Error)
	GetPaymentsById(ctx context.Context, negativityId string, filter PageableDefaultRequest) (
		*Pageable[NegativityPaymentsResponse], Error)
	GetChargesAvailableForDunning(ctx context.Context, filter PageableDefaultRequest) (
		*Pageable[ChargesAvailableForDunningResponse], Error)
	GetAll(ctx context.Context, filter GetAllNegativitiesRequest) (*Pageable[NegativityResponse], Error)
}

func NewNegativity(env Env, accessToken string) Negativity {
	logWarning("Negativity service running on", env.String())
	return negativity{
		env:         env,
		accessToken: accessToken,
	}
}

func (n negativity) Create(ctx context.Context, body CreateNegativityRequest) (*NegativityResponse, Error) {
	if err := Validate().Struct(body); err != nil {
		return nil, NewError(ErrorTypeValidation, err)
	}
	req := NewRequest[NegativityResponse](ctx, n.env, n.accessToken)
	return req.make(http.MethodPost, "/v3/paymentDunnings", body)
}

func (n negativity) Simulate(ctx context.Context, chargeId string) (*NegativitySimulateResponse, Error) {
	if util.IsBlank(&chargeId) {
		return nil, NewError(ErrorTypeValidation, "chargeId is required")
	}
	req := NewRequest[NegativitySimulateResponse](ctx, n.env, n.accessToken)
	return req.make(http.MethodPost, fmt.Sprintf("/v3/paymentDunnings/simulate?payment=%s", chargeId), nil)
}

func (n negativity) ResendDocumentsById(ctx context.Context, negativityId string, body NegativityResendDocumentsRequest) (
	*NegativityResponse, Error) {
	if err := Validate().Struct(body); err != nil {
		return nil, NewError(ErrorTypeValidation, err)
	}
	req := NewRequest[NegativityResponse](ctx, n.env, n.accessToken)
	return req.makeMultipartForm(http.MethodPost, fmt.Sprintf("/v3/paymentDunnings/%s/documents", negativityId), body)
}

func (n negativity) CancelById(ctx context.Context, negativityId string) (*NegativityResponse, Error) {
	req := NewRequest[NegativityResponse](ctx, n.env, n.accessToken)
	return req.make(http.MethodPost, fmt.Sprintf("/v3/paymentDunnings/%s/cancel", negativityId), nil)
}

func (n negativity) GetById(ctx context.Context, negativityId string) (*NegativityResponse, Error) {
	req := NewRequest[NegativityResponse](ctx, n.env, n.accessToken)
	return req.make(http.MethodGet, fmt.Sprintf("/v3/paymentDunnings/%s", negativityId), nil)
}

func (n negativity) GetHistoryById(ctx context.Context, negativityId string, filter PageableDefaultRequest) (
	*Pageable[NegativityHistoryResponse], Error) {
	req := NewRequest[Pageable[NegativityHistoryResponse]](ctx, n.env, n.accessToken)
	return req.make(http.MethodGet, fmt.Sprintf("/v3/paymentDunnings/%s/history", negativityId), filter)
}

func (n negativity) GetPaymentsById(ctx context.Context, negativityId string, filter PageableDefaultRequest) (
	*Pageable[NegativityPaymentsResponse], Error) {
	req := NewRequest[Pageable[NegativityPaymentsResponse]](ctx, n.env, n.accessToken)
	return req.make(http.MethodGet, fmt.Sprintf("/v3/paymentDunnings/%s/partialPayments", negativityId), filter)
}

func (n negativity) GetChargesAvailableForDunning(ctx context.Context, filter PageableDefaultRequest) (
	*Pageable[ChargesAvailableForDunningResponse], Error) {
	req := NewRequest[Pageable[ChargesAvailableForDunningResponse]](ctx, n.env, n.accessToken)
	return req.make(http.MethodGet, "/v3/paymentDunnings/paymentsAvailableForDunning", filter)
}

func (n negativity) GetAll(ctx context.Context, filter GetAllNegativitiesRequest) (*Pageable[NegativityResponse], Error) {
	req := NewRequest[Pageable[NegativityResponse]](ctx, n.env, n.accessToken)
	return req.make(http.MethodGet, "/v3/paymentDunnings", filter)
}
