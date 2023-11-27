package asaas

import (
	"context"
	berrors "errors"
	"fmt"
	"net/http"
	"time"
)

type CreateSubscriptionRequest struct {
	Customer             string                       `json:"customer,omitempty" validate:"required"`
	BillingType          BillingType                  `json:"billingType,omitempty" validate:"required,enum"`
	Value                float64                      `json:"value,omitempty" validate:"required,gt=0"`
	NextDueDate          Date                         `json:"nextDueDate,omitempty" validate:"required"`
	Discount             *DiscountRequest             `json:"discount,omitempty"`
	Interest             *InterestRequest             `json:"interest,omitempty"`
	Fine                 *FineRequest                 `json:"fine,omitempty"`
	Cycle                SubscriptionCycle            `json:"cycle,omitempty" validate:"required,enum"`
	Description          string                       `json:"description,omitempty" validate:"omitempty,lte=500"`
	CreditCard           *CreditCardRequest           `json:"creditCard,omitempty"`
	CreditCardHolderInfo *CreditCardHolderInfoRequest `json:"creditCardHolderInfo,omitempty"`
	CreditCardToken      string                       `json:"creditCardToken,omitempty"`
	EndDate              *Date                        `json:"endDate,omitempty"`
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
	ID                string             `json:"id,omitempty"`
	Customer          string             `json:"customer,omitempty"`
	Refunds           []RefundResponse   `json:"refunds,omitempty"`
	Errors            []ErrorResponse    `json:"errors,omitempty"`
	BillingType       BillingType        `json:"billingType,omitempty"`
	Value             float64            `json:"value,omitempty"`
	NextDueDate       Date               `json:"nextDueDate,omitempty"`
	Cycle             SubscriptionCycle  `json:"cycle,omitempty"`
	Status            SubscriptionStatus `json:"status,omitempty"`
	Discount          *DiscountResponse  `json:"discount,omitempty"`
	Interest          *InterestResponse  `json:"interest,omitempty"`
	Fine              *FineResponse      `json:"fine,omitempty"`
	Description       string             `json:"description,omitempty"`
	EndDate           *Date              `json:"endDate,omitempty"`
	MaxPayments       int                `json:"maxPayments,omitempty"`
	ExternalReference string             `json:"externalReference,omitempty"`
	Deleted           bool               `json:"deleted,omitempty"`
	DateCreated       *DateTime          `json:"dateCreated,omitempty"`
}

type subscription struct {
	env         Env
	accessToken string
}

type Subscription interface {
	Create(ctx context.Context, body CreateSubscriptionRequest) (*SubscriptionResponse, Error)
	CreateInvoiceSettingByID(ctx context.Context, subscriptionID string, body CreateInvoiceSettingRequest) (
		*InvoiceSettingResponse, Error)
	UpdateByID(ctx context.Context, subscriptionID string, body UpdateSubscriptionRequest) (*SubscriptionResponse, Error)
	UpdateInvoiceSettingsByID(ctx context.Context, subscriptionID string, body UpdateInvoiceSettingRequest) (
		*InvoiceSettingResponse, Error)
	DeleteByID(ctx context.Context, subscriptionID string) (*DeleteResponse, Error)
	DeleteInvoiceSettingByID(ctx context.Context, subscriptionID string) (*DeleteResponse, Error)
	GetByID(ctx context.Context, subscriptionID string) (*SubscriptionResponse, Error)
	GetInvoiceSettingByID(ctx context.Context, subscriptionID string) (*InvoiceSettingResponse, Error)
	GetAllChargesBySubscription(ctx context.Context, subscriptionID string, filter GetAllChargesBySubscriptionRequest) (
		*Pageable[ChargeResponse], Error)
	GetAllInvoicesBySubscription(ctx context.Context, subscriptionID string, filter GetAllInvoicesRequest) (
		*Pageable[InvoiceResponse], Error)
	GetPaymentBookByID(ctx context.Context, subscriptionID string, filter SubscriptionPaymentBookRequest) (
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

func (s subscription) CreateInvoiceSettingByID(ctx context.Context, subscriptionID string, body CreateInvoiceSettingRequest) (
	*InvoiceSettingResponse, Error) {
	if err := Validate().Struct(body); err != nil {
		return nil, NewError(ErrorTypeValidation, err)
	}
	req := NewRequest[InvoiceSettingResponse](ctx, s.env, s.accessToken)
	return req.make(http.MethodPost, fmt.Sprintf("/v3/subscriptions/%s/invoiceSettings", subscriptionID), body)
}

func (s subscription) UpdateByID(ctx context.Context, subscriptionID string, body UpdateSubscriptionRequest) (
	*SubscriptionResponse, Error) {
	if err := Validate().Struct(body); err != nil {
		return nil, NewError(ErrorTypeValidation, err)
	}
	req := NewRequest[SubscriptionResponse](ctx, s.env, s.accessToken)
	return req.make(http.MethodPut, fmt.Sprintf("/v3/subscriptions/%s", subscriptionID), body)
}

func (s subscription) UpdateInvoiceSettingsByID(ctx context.Context, subscriptionID string, body UpdateInvoiceSettingRequest) (
	*InvoiceSettingResponse, Error) {
	if err := Validate().Struct(body); err != nil {
		return nil, NewError(ErrorTypeValidation, err)
	}
	req := NewRequest[InvoiceSettingResponse](ctx, s.env, s.accessToken)
	return req.make(http.MethodPut, fmt.Sprintf("/v3/subscriptions/%s/invoiceSettings", subscriptionID), body)
}

func (s subscription) DeleteByID(ctx context.Context, subscriptionID string) (*DeleteResponse, Error) {
	req := NewRequest[DeleteResponse](ctx, s.env, s.accessToken)
	return req.make(http.MethodDelete, fmt.Sprintf("/v3/subscriptions/%s", subscriptionID), nil)
}

func (s subscription) DeleteInvoiceSettingByID(ctx context.Context, subscriptionID string) (*DeleteResponse, Error) {
	req := NewRequest[DeleteResponse](ctx, s.env, s.accessToken)
	return req.make(http.MethodDelete, fmt.Sprintf("/v3/subscriptions/%s/invoiceSettings", subscriptionID), nil)
}

func (s subscription) GetByID(ctx context.Context, subscriptionID string) (*SubscriptionResponse, Error) {
	req := NewRequest[SubscriptionResponse](ctx, s.env, s.accessToken)
	return req.make(http.MethodGet, fmt.Sprintf("/v3/subscriptions/%s", subscriptionID), nil)
}

func (s subscription) GetInvoiceSettingByID(ctx context.Context, subscriptionID string) (*InvoiceSettingResponse, Error) {
	req := NewRequest[InvoiceSettingResponse](ctx, s.env, s.accessToken)
	return req.make(http.MethodGet, fmt.Sprintf("/v3/subscriptions/%s/invoiceSettings", subscriptionID), nil)
}

func (s subscription) GetPaymentBookByID(ctx context.Context, subscriptionID string, filter SubscriptionPaymentBookRequest) (
	*FileTextPlainResponse, Error) {
	if err := Validate().Struct(filter); err != nil {
		return nil, NewError(ErrorTypeValidation, err)
	}
	req := NewRequest[FileTextPlainResponse](ctx, s.env, s.accessToken)
	return req.make(http.MethodGet, fmt.Sprintf("/v3/subscriptions/%s/paymentBook", subscriptionID), filter)
}

func (s subscription) GetAllChargesBySubscription(
	ctx context.Context,
	subscriptionID string,
	filter GetAllChargesBySubscriptionRequest) (*Pageable[ChargeResponse], Error) {
	req := NewRequest[Pageable[ChargeResponse]](ctx, s.env, s.accessToken)
	return req.make(http.MethodGet, fmt.Sprintf("/v3/subscriptions/%s/payments", subscriptionID), filter)
}

func (s subscription) GetAllInvoicesBySubscription(ctx context.Context, subscriptionID string,
	filter GetAllInvoicesRequest) (*Pageable[InvoiceResponse], Error) {
	req := NewRequest[Pageable[InvoiceResponse]](ctx, s.env, s.accessToken)
	return req.make(http.MethodGet, fmt.Sprintf("/v3/subscriptions/%s/invoices", subscriptionID), filter)
}

func (s subscription) GetAll(ctx context.Context, filter GetAllSubscriptionsRequest) (*Pageable[SubscriptionResponse],
	Error) {
	req := NewRequest[Pageable[SubscriptionResponse]](ctx, s.env, s.accessToken)
	return req.make(http.MethodGet, "/v3/subscriptions", filter)
}

func (s subscription) validateCreateBodyRequest(body CreateSubscriptionRequest) error {
	if err := Validate().Struct(body); err != nil {
		return err
	} else if time.Now().After(body.NextDueDate.Time()) {
		return berrors.New("invalid nextDueDate")
	} else if body.EndDate != nil && time.Now().After(body.EndDate.Time()) {
		return berrors.New("invalid endDate")
	}
	return validateBillingBody(body.BillingType, body.CreditCard, body.CreditCardHolderInfo, body.CreditCardToken,
		body.RemoteIp)
}

func (s subscription) prepareCreateBodyRequest(body *CreateSubscriptionRequest) {
	body.NextDueDate = NewDate(body.NextDueDate.Year(), body.NextDueDate.Month(), body.NextDueDate.Day(),
		23, 59, 0, 0, body.NextDueDate.Location())
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
