package asaas

import (
	"context"
	"fmt"
	"net/http"
)

type InstallmentPaymentBookRequest struct {
	Sort  SortPaymentBookField `json:"sort,omitempty" validate:"omitempty,enum"`
	Order Order                `json:"order,omitempty" validate:"omitempty,enum"`
}

type InstallmentResponse struct {
	ID                    string                         `json:"id,omitempty"`
	Customer              string                         `json:"customer,omitempty"`
	Value                 float64                        `json:"value,omitempty"`
	NetValue              float64                        `json:"netValue,omitempty"`
	PaymentValue          float64                        `json:"paymentValue,omitempty"`
	InstallmentCount      int                            `json:"installmentCount,omitempty"`
	BillingType           BillingType                    `json:"billingType,omitempty"`
	PaymentDate           Date                           `json:"paymentDate,omitempty"`
	Description           string                         `json:"description,omitempty"`
	ExpirationDay         int                            `json:"expirationDay,omitempty"`
	Deleted               bool                           `json:"deleted,omitempty"`
	PaymentLink           string                         `json:"paymentLink,omitempty"`
	TransactionReceiptUrl string                         `json:"transactionReceiptUrl,omitempty"`
	Chargeback            *InstallmentChargebackResponse `json:"chargeback,omitempty"`
	DateCreated           Date                           `json:"dateCreated,omitempty"`
}

type InstallmentChargebackResponse struct {
	Status ChargebackStatus `json:"status,omitempty"`
	Reason ChargebackReason `json:"reason,omitempty"`
}

type UpdateInstallmentSplitsResponse struct {
	Splits []SplitResponse `json:"splits,omitempty"`
}

type installment struct {
	env         Env
	accessToken string
}

type Installment interface {
	UpdateSplitsByID(ctx context.Context, installmentID string, body []SplitRequest) (*UpdateInstallmentSplitsResponse,
		Error)
	RefundByID(ctx context.Context, installmentID string) (*InstallmentResponse, Error)
	DeleteByID(ctx context.Context, installmentID string) (*DeleteResponse, Error)
	GetByID(ctx context.Context, installmentID string) (*InstallmentResponse, Error)
	GetPaymentBookByID(ctx context.Context, installmentID string, filter InstallmentPaymentBookRequest) (
		*FileTextPlainResponse, Error)
	GetAll(ctx context.Context, filter PageableDefaultRequest) (*Pageable[InstallmentResponse], Error)
}

func NewInstallment(env Env, accessToken string) Installment {
	logWarning("Installment service running on", env.String())
	return installment{
		env:         env,
		accessToken: accessToken,
	}
}

func (i installment) UpdateSplitsByID(ctx context.Context, installmentID string, body []SplitRequest) (
	*UpdateInstallmentSplitsResponse, Error) {
	if err := Validate().Struct(body); err != nil {
		return nil, NewError(ErrorTypeValidation, err)
	}
	req := NewRequest[UpdateInstallmentSplitsResponse](ctx, i.env, i.accessToken)
	return req.make(http.MethodPost, fmt.Sprintf("/v3/installments/%s/splits", installmentID), body)
}

func (i installment) RefundByID(ctx context.Context, installmentID string) (*InstallmentResponse, Error) {
	req := NewRequest[InstallmentResponse](ctx, i.env, i.accessToken)
	return req.make(http.MethodPost, fmt.Sprintf("/v3/installments/%s/refund", installmentID), nil)
}

func (i installment) DeleteByID(ctx context.Context, installmentID string) (*DeleteResponse, Error) {
	req := NewRequest[DeleteResponse](ctx, i.env, i.accessToken)
	return req.make(http.MethodDelete, fmt.Sprintf("/v3/installments/%s", installmentID), nil)
}

func (i installment) GetByID(ctx context.Context, installmentID string) (*InstallmentResponse, Error) {
	req := NewRequest[InstallmentResponse](ctx, i.env, i.accessToken)
	return req.make(http.MethodGet, fmt.Sprintf("/v3/installments/%s", installmentID), nil)
}

func (i installment) GetPaymentBookByID(ctx context.Context, installmentID string, filter InstallmentPaymentBookRequest) (*FileTextPlainResponse, Error) {
	req := NewRequest[FileTextPlainResponse](ctx, i.env, i.accessToken)
	return req.make(http.MethodGet, fmt.Sprintf("/v3/installments/%s/paymentBook", installmentID), filter)
}

func (i installment) GetAll(ctx context.Context, filter PageableDefaultRequest) (*Pageable[InstallmentResponse], Error) {
	req := NewRequest[Pageable[InstallmentResponse]](ctx, i.env, i.accessToken)
	return req.make(http.MethodGet, "/v3/installments", filter)
}
