package asaas

import (
	"context"
	berrors "errors"
	"fmt"
	"net/http"
	"os"
)

type PaymentLinkRequest struct {
	//Nome do link de pagamentos (OBRIGATÓRIO)
	Name                string            `json:"name,omitempty" validate:"required"`
	Description         string            `json:"description,omitempty"`
	BillingType         BillingType       `json:"billingType,omitempty" validate:"required,enum"`
	ChargeType          ChargeType        `json:"chargeType,omitempty" validate:"required,enum"`
	EndDate             *Date             `json:"endDate,omitempty" validate:"omitempty,after_now"`
	Value               float64           `json:"value,omitempty" validate:"omitempty,gt=0"`
	DueDateLimitDays    int               `json:"dueDateLimitDays,omitempty" validate:"omitempty,gte=0"`
	SubscriptionCycle   SubscriptionCycle `json:"subscriptionCycle,omitempty" validate:"omitempty,enum"`
	MaxInstallmentCount int               `json:"maxInstallmentCount,omitempty" validate:"omitempty,gt=0"`
	NotificationEnabled bool              `json:"notificationEnabled,omitempty"`
	Callback            *CallbackRequest  `json:"callback,omitempty"`
}

type GetAllPaymentLinksRequest struct {
	Name           string `json:"name,omitempty"`
	Active         bool   `json:"active,omitempty"`
	IncludeDeleted bool   `json:"includeDeleted,omitempty"`
	Offset         int    `json:"offset,omitempty"`
	Limit          int    `json:"limit,omitempty"`
}

type SendImagePaymentLinksRequest struct {
	Main  bool     `json:"main,omitempty"`
	Image *os.File `json:"image,omitempty"`
}

type PaymentLinkResponse struct {
	Id                  string            `json:"id,omitempty"`
	Name                string            `json:"name,omitempty"`
	Url                 string            `json:"url,omitempty"`
	Active              bool              `json:"active,omitempty"`
	BillingType         BillingType       `json:"billingType,omitempty"`
	ChargeType          ChargeType        `json:"chargeType,omitempty"`
	EndDate             *Date             `json:"endDate,omitempty"`
	Value               float64           `json:"value,omitempty"`
	SubscriptionCycle   SubscriptionCycle `json:"subscriptionCycle,omitempty"`
	Description         string            `json:"description,omitempty"`
	MaxInstallmentCount int               `json:"maxInstallmentCount,omitempty"`
	DueDateLimitDays    int               `json:"dueDateLimitDays,omitempty"`
	NotificationEnabled bool              `json:"notificationEnabled,omitempty"`
	Errors              []ErrorResponse   `json:"errors,omitempty"`
}

type PaymentLinkImageResponse struct {
	Id     string               `json:"id,omitempty"`
	Main   bool                 `json:"main,omitempty"`
	Image  PaymentImageDataLink `json:"image,omitempty"`
	Errors []ErrorResponse      `json:"errors,omitempty"`
}

type PaymentImageDataLink struct {
	OriginalName string `json:"originalName,omitempty"`
	Size         int    `json:"size,omitempty"`
	Extension    string `json:"extension,omitempty"`
	PreviewUrl   string `json:"previewUrl,omitempty"`
	DownloadUrl  string `json:"downloadUrl,omitempty"`
}

type paymentLink struct {
	env         Env
	accessToken string
}

type PaymentLink interface {
	Create(ctx context.Context, body PaymentLinkRequest) (*PaymentLinkResponse, Error)
	SendImageByID(ctx context.Context, paymentLinkId string, body SendImagePaymentLinksRequest) (
		*PaymentLinkImageResponse, Error)
	UpdateByID(ctx context.Context, paymentLinkId string, body PaymentLinkRequest) (*PaymentLinkResponse, Error)
	UpdateImageAsMainByID(ctx context.Context, paymentLinkId, imageId string) (*PaymentLinkImageResponse, Error)
	RestoreByID(ctx context.Context, paymentLinkId string) (*PaymentLinkResponse, Error)
	DeleteByID(ctx context.Context, paymentLinkId string) (*DeleteResponse, Error)
	DeleteImageByID(ctx context.Context, paymentLinkId, imageId string) (*DeleteResponse, Error)
	GetByID(ctx context.Context, paymentLinkId string) (*PaymentLinkResponse, Error)
	GetImageByID(ctx context.Context, paymentLinkId, imageId string) (*PaymentLinkImageResponse, Error)
	GetAll(ctx context.Context, filter GetAllPaymentLinksRequest) (*Pageable[PaymentLinkResponse], Error)
	GetImagesByID(ctx context.Context, paymentLinkId string) (*Pageable[PaymentLinkImageResponse], Error)
}

func NewPaymentLink(env Env, accessToken string) PaymentLink {
	logWarning("PaymentLink service running on", env.String())
	return paymentLink{
		env:         env,
		accessToken: accessToken,
	}
}

func (p paymentLink) Create(ctx context.Context, body PaymentLinkRequest) (*PaymentLinkResponse, Error) {
	if err := p.validateCreateBodyRequest(body); err != nil {
		return nil, NewError(ErrorTypeValidation, err)
	}
	req := NewRequest[PaymentLinkResponse](ctx, p.env, p.accessToken)
	return req.make(http.MethodPost, "/v3/paymentLinks", body)
}

func (p paymentLink) SendImageByID(ctx context.Context, paymentLinkId string, body SendImagePaymentLinksRequest) (
	*PaymentLinkImageResponse, Error) {
	req := NewRequest[PaymentLinkImageResponse](ctx, p.env, p.accessToken)
	return req.makeMultipartForm(http.MethodPost, fmt.Sprintf("/v3/paymentLinks/%s/images", paymentLinkId), body)
}

func (p paymentLink) UpdateByID(ctx context.Context, paymentLinkId string, body PaymentLinkRequest) (
	*PaymentLinkResponse, Error) {
	req := NewRequest[PaymentLinkResponse](ctx, p.env, p.accessToken)
	return req.make(http.MethodPut, fmt.Sprintf("/v3/paymentLinks/%s", paymentLinkId), body)
}

func (p paymentLink) UpdateImageAsMainByID(ctx context.Context, paymentLinkId, imageId string) (
	*PaymentLinkImageResponse, Error) {
	req := NewRequest[PaymentLinkImageResponse](ctx, p.env, p.accessToken)
	return req.make(http.MethodPut, fmt.Sprintf("/v3/paymentLinks/%s/images/%s/setAsMain", paymentLinkId, imageId), nil)
}

func (p paymentLink) RestoreByID(ctx context.Context, paymentLinkId string) (*PaymentLinkResponse, Error) {
	req := NewRequest[PaymentLinkResponse](ctx, p.env, p.accessToken)
	return req.make(http.MethodPost, fmt.Sprintf("/v3/paymentLinks/%s", paymentLinkId), nil)
}

func (p paymentLink) DeleteByID(ctx context.Context, paymentLinkId string) (*DeleteResponse, Error) {
	req := NewRequest[DeleteResponse](ctx, p.env, p.accessToken)
	return req.make(http.MethodDelete, fmt.Sprintf("/v3/paymentLinks/%s", paymentLinkId), nil)
}

func (p paymentLink) DeleteImageByID(ctx context.Context, paymentLinkId, imageId string) (*DeleteResponse, Error) {
	req := NewRequest[DeleteResponse](ctx, p.env, p.accessToken)
	return req.make(http.MethodDelete, fmt.Sprintf("/v3/paymentLinks/%s/images/%s", paymentLinkId, imageId), nil)
}

func (p paymentLink) GetByID(ctx context.Context, paymentLinkId string) (*PaymentLinkResponse, Error) {
	req := NewRequest[PaymentLinkResponse](ctx, p.env, p.accessToken)
	return req.make(http.MethodGet, fmt.Sprintf("/v3/paymentLinks/%s", paymentLinkId), nil)
}

func (p paymentLink) GetImageByID(ctx context.Context, paymentLinkId, imageId string) (*PaymentLinkImageResponse, Error) {
	req := NewRequest[PaymentLinkImageResponse](ctx, p.env, p.accessToken)
	return req.make(http.MethodGet, fmt.Sprintf("/v3/paymentLinks/%s/images/%s", paymentLinkId, imageId), nil)
}

func (p paymentLink) GetAll(ctx context.Context, filter GetAllPaymentLinksRequest) (*Pageable[PaymentLinkResponse], Error) {
	req := NewRequest[Pageable[PaymentLinkResponse]](ctx, p.env, p.accessToken)
	return req.make(http.MethodGet, "/v3/paymentLinks", filter)
}

func (p paymentLink) GetImagesByID(ctx context.Context, paymentLinkId string) (*Pageable[PaymentLinkImageResponse], Error) {
	req := NewRequest[Pageable[PaymentLinkImageResponse]](ctx, p.env, p.accessToken)
	return req.make(http.MethodGet, fmt.Sprintf("/v3/paymentLinks/%s/images", paymentLinkId), nil)
}

func (p paymentLink) validateCreateBodyRequest(body PaymentLinkRequest) error {
	if err := Validate().Struct(body); err != nil {
		return err
	} else if body.ChargeType == ChargeTypeRecurrent && !body.SubscriptionCycle.IsEnumValid() {
		return berrors.New("subscriptionCycle is required for RECURRENT chargeType")
	}
	return nil
}
