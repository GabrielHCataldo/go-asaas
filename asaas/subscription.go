package asaas

import (
	"context"
	"fmt"
	"net/http"
)

type CreateSubscriptionRequest struct {
	Customer             string                       `json:"customer,omitempty" validate:"required"`
	BillingType          BillingType                  `json:"billingType,omitempty" validate:"required,enum"`
	Value                float64                      `json:"value,omitempty" validate:"required,gt=0"`
	NextDueDate          Date                         `json:"nextDueDate,omitempty" validate:"required,after_now"`
	Discount             *DiscountRequest             `json:"discount,omitempty"`
	Interest             *InterestRequest             `json:"interest,omitempty"`
	Fine                 *FineRequest                 `json:"fine,omitempty"`
	Cycle                SubscriptionCycle            `json:"cycle,omitempty" validate:"required,enum"`
	Description          string                       `json:"description,omitempty" validate:"omitempty,lte=500"`
	CreditCard           *CreditCardRequest           `json:"creditCard,omitempty"`
	CreditCardHolderInfo *CreditCardHolderInfoRequest `json:"creditCardHolderInfo,omitempty"`
	CreditCardToken      string                       `json:"creditCardToken,omitempty"`
	EndDate              *Date                        `json:"endDate,omitempty" validate:"omitempty,after_now"`
	MaxPayments          int                          `json:"maxPayments,omitempty" validate:"omitempty,gt=0"`
	ExternalReference    string                       `json:"externalReference,omitempty"`
	Split                []SplitRequest               `json:"split,omitempty"`
	RemoteIp             string                       `json:"remoteIp,omitempty"`
}

type UpdateSubscriptionRequest struct {
	BillingType           BillingType                  `json:"billingType,omitempty" validate:"omitempty,enum"`
	Value                 float64                      `json:"value,omitempty" validate:"omitempty,gt=0"`
	Status                SubscriptionStatus           `json:"status,omitempty" validate:"omitempty,enum"`
	NextDueDate           *Date                        `json:"nextDueDate,omitempty"`
	Discount              *DiscountRequest             `json:"discount,omitempty"`
	Interest              *InterestRequest             `json:"interest,omitempty"`
	Fine                  *FineRequest                 `json:"fine,omitempty"`
	Cycle                 SubscriptionCycle            `json:"cycle,omitempty" validate:"omitempty,enum"`
	Description           string                       `json:"description,omitempty" validate:"omitempty,lte=500"`
	CreditCard            *CreditCardRequest           `json:"creditCard,omitempty"`
	CreditCardHolderInfo  *CreditCardHolderInfoRequest `json:"creditCardHolderInfo,omitempty"`
	CreditCardToken       string                       `json:"creditCardToken,omitempty"`
	EndDate               *Date                        `json:"endDate,omitempty"`
	UpdatePendingPayments bool                         `json:"updatePendingPayments,omitempty"`
	ExternalReference     string                       `json:"externalReference,omitempty"`
}

type GetAllSubscriptionsRequest struct {
	Customer          string                `json:"customer,omitempty"`
	CustomerGroupName string                `json:"customerGroupName,omitempty"`
	BillingType       BillingType           `json:"billingType,omitempty"`
	Status            SubscriptionStatus    `json:"status,omitempty"`
	DeletedOnly       bool                  `json:"deletedOnly,omitempty"`
	IncludeDeleted    bool                  `json:"includeDeleted,omitempty"`
	ExternalReference string                `json:"externalReference,omitempty"`
	Order             Order                 `json:"order,omitempty"`
	Sort              SortSubscriptionField `json:"sort,omitempty"`
	Offset            int                   `json:"offset,omitempty"`
	Limit             int                   `json:"limit,omitempty"`
}

type GetAllChargesBySubscriptionRequest struct {
	Status ChargeStatus `json:"status,omitempty"`
}

type SubscriptionPaymentBookRequest struct {
	Month int                  `json:"month,omitempty" validate:"required,gte=1,lte=12"`
	Year  int                  `json:"year,omitempty" validate:"required,gt=0"`
	Sort  SortPaymentBookField `json:"sort,omitempty" validate:"omitempty,enum"`
	Order Order                `json:"order,omitempty" validate:"omitempty,enum"`
}

type SubscriptionResponse struct {
	Id                string             `json:"id,omitempty"`
	Customer          string             `json:"customer,omitempty"`
	Status            SubscriptionStatus `json:"status,omitempty"`
	Refunds           []RefundResponse   `json:"refunds,omitempty"`
	BillingType       BillingType        `json:"billingType,omitempty"`
	Value             float64            `json:"value,omitempty"`
	NextDueDate       *Date              `json:"nextDueDate,omitempty"`
	Cycle             SubscriptionCycle  `json:"cycle,omitempty"`
	Discount          *DiscountResponse  `json:"discount,omitempty"`
	Interest          *InterestResponse  `json:"interest,omitempty"`
	Fine              *FineResponse      `json:"fine,omitempty"`
	Description       string             `json:"description,omitempty"`
	EndDate           *Date              `json:"endDate,omitempty"`
	MaxPayments       int                `json:"maxPayments,omitempty"`
	ExternalReference string             `json:"externalReference,omitempty"`
	Deleted           bool               `json:"deleted,omitempty"`
	Errors            []ErrorResponse    `json:"errors,omitempty"`
	DateCreated       *Date              `json:"dateCreated,omitempty"`
}

type subscription struct {
	env         Env
	accessToken string
}

type Subscription interface {
	Create(ctx context.Context, body CreateSubscriptionRequest) (*SubscriptionResponse, Error)
	CreateInvoiceSettingById(ctx context.Context, subscriptionId string, body CreateInvoiceSettingRequest) (
		*InvoiceSettingResponse, Error)
	UpdateById(ctx context.Context, subscriptionId string, body UpdateSubscriptionRequest) (*SubscriptionResponse, Error)
	UpdateInvoiceSettingsById(ctx context.Context, subscriptionId string, body UpdateInvoiceSettingRequest) (
		*InvoiceSettingResponse, Error)
	DeleteById(ctx context.Context, subscriptionId string) (*DeleteResponse, Error)
	DeleteInvoiceSettingById(ctx context.Context, subscriptionId string) (*DeleteResponse, Error)
	GetById(ctx context.Context, subscriptionId string) (*SubscriptionResponse, Error)
	GetInvoiceSettingById(ctx context.Context, subscriptionId string) (*InvoiceSettingResponse, Error)
	GetAllChargesBySubscription(ctx context.Context, subscriptionId string, filter GetAllChargesBySubscriptionRequest) (
		*Pageable[ChargeResponse], Error)
	GetAllInvoicesBySubscription(ctx context.Context, subscriptionId string, filter GetAllInvoicesRequest) (
		*Pageable[InvoiceResponse], Error)
	GetPaymentBookById(ctx context.Context, subscriptionId string, filter SubscriptionPaymentBookRequest) (
		*FileTextPlainResponse, Error)
	GetAll(ctx context.Context, filter GetAllSubscriptionsRequest) (*Pageable[SubscriptionResponse], Error)
}

func NewSubscription(env Env, accessToken string) Subscription {
	logWarning("Subscription service running on", env.String())
	return subscription{
		env:         env,
		accessToken: accessToken,
	}
}

func (s subscription) Create(ctx context.Context, body CreateSubscriptionRequest) (*SubscriptionResponse, Error) {
	if err := s.validateCreateBodyRequest(body); err != nil {
		return nil, NewError(ErrorTypeValidation, err)
	}
	s.prepareCreateBodyRequest(&body)
	req := NewRequest[SubscriptionResponse](ctx, s.env, s.accessToken)
	return req.make(http.MethodPost, "/v3/subscriptions", body)
}

func (s subscription) CreateInvoiceSettingById(ctx context.Context, subscriptionId string, body CreateInvoiceSettingRequest) (
	*InvoiceSettingResponse, Error) {
	if err := Validate().Struct(body); err != nil {
		return nil, NewError(ErrorTypeValidation, err)
	}
	req := NewRequest[InvoiceSettingResponse](ctx, s.env, s.accessToken)
	return req.make(http.MethodPost, fmt.Sprintf("/v3/subscriptions/%s/invoiceSettings", subscriptionId), body)
}

func (s subscription) UpdateById(ctx context.Context, subscriptionId string, body UpdateSubscriptionRequest) (
	*SubscriptionResponse, Error) {
	if err := Validate().Struct(body); err != nil {
		return nil, NewError(ErrorTypeValidation, err)
	}
	req := NewRequest[SubscriptionResponse](ctx, s.env, s.accessToken)
	return req.make(http.MethodPut, fmt.Sprintf("/v3/subscriptions/%s", subscriptionId), body)
}

func (s subscription) UpdateInvoiceSettingsById(ctx context.Context, subscriptionId string, body UpdateInvoiceSettingRequest) (
	*InvoiceSettingResponse, Error) {
	if err := Validate().Struct(body); err != nil {
		return nil, NewError(ErrorTypeValidation, err)
	}
	req := NewRequest[InvoiceSettingResponse](ctx, s.env, s.accessToken)
	return req.make(http.MethodPut, fmt.Sprintf("/v3/subscriptions/%s/invoiceSettings", subscriptionId), body)
}

func (s subscription) DeleteById(ctx context.Context, subscriptionId string) (*DeleteResponse, Error) {
	req := NewRequest[DeleteResponse](ctx, s.env, s.accessToken)
	return req.make(http.MethodDelete, fmt.Sprintf("/v3/subscriptions/%s", subscriptionId), nil)
}

func (s subscription) DeleteInvoiceSettingById(ctx context.Context, subscriptionId string) (*DeleteResponse, Error) {
	req := NewRequest[DeleteResponse](ctx, s.env, s.accessToken)
	return req.make(http.MethodDelete, fmt.Sprintf("/v3/subscriptions/%s/invoiceSettings", subscriptionId), nil)
}

func (s subscription) GetById(ctx context.Context, subscriptionId string) (*SubscriptionResponse, Error) {
	req := NewRequest[SubscriptionResponse](ctx, s.env, s.accessToken)
	return req.make(http.MethodGet, fmt.Sprintf("/v3/subscriptions/%s", subscriptionId), nil)
}

func (s subscription) GetInvoiceSettingById(ctx context.Context, subscriptionId string) (*InvoiceSettingResponse, Error) {
	req := NewRequest[InvoiceSettingResponse](ctx, s.env, s.accessToken)
	return req.make(http.MethodGet, fmt.Sprintf("/v3/subscriptions/%s/invoiceSettings", subscriptionId), nil)
}

func (s subscription) GetPaymentBookById(ctx context.Context, subscriptionId string, filter SubscriptionPaymentBookRequest) (
	*FileTextPlainResponse, Error) {
	if err := Validate().Struct(filter); err != nil {
		return nil, NewError(ErrorTypeValidation, err)
	}
	req := NewRequest[FileTextPlainResponse](ctx, s.env, s.accessToken)
	return req.make(http.MethodGet, fmt.Sprintf("/v3/subscriptions/%s/paymentBook", subscriptionId), filter)
}

func (s subscription) GetAllChargesBySubscription(
	ctx context.Context,
	subscriptionId string,
	filter GetAllChargesBySubscriptionRequest) (*Pageable[ChargeResponse], Error) {
	req := NewRequest[Pageable[ChargeResponse]](ctx, s.env, s.accessToken)
	return req.make(http.MethodGet, fmt.Sprintf("/v3/subscriptions/%s/payments", subscriptionId), filter)
}

func (s subscription) GetAllInvoicesBySubscription(ctx context.Context, subscriptionId string,
	filter GetAllInvoicesRequest) (*Pageable[InvoiceResponse], Error) {
	req := NewRequest[Pageable[InvoiceResponse]](ctx, s.env, s.accessToken)
	return req.make(http.MethodGet, fmt.Sprintf("/v3/subscriptions/%s/invoices", subscriptionId), filter)
}

func (s subscription) GetAll(ctx context.Context, filter GetAllSubscriptionsRequest) (*Pageable[SubscriptionResponse],
	Error) {
	req := NewRequest[Pageable[SubscriptionResponse]](ctx, s.env, s.accessToken)
	return req.make(http.MethodGet, "/v3/subscriptions", filter)
}

func (s subscription) validateCreateBodyRequest(body CreateSubscriptionRequest) error {
	if err := Validate().Struct(body); err != nil {
		return err
	}
	return validateBillingBody(body.BillingType, body.CreditCard, body.CreditCardHolderInfo, body.CreditCardToken,
		body.RemoteIp)
}

func (s subscription) prepareCreateBodyRequest(body *CreateSubscriptionRequest) {
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
