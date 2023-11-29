package asaas

import (
	"context"
	"fmt"
	"net/http"
)

type MobilePhoneRechargeRequest struct {
	PhoneNumber string  `json:"phoneNumber,omitempty" validate:"required,phone"`
	Value       float64 `json:"value,omitempty" validate:"required,gt=0"`
}

type MobilePhoneRechargeResponse struct {
	ID             string                    `json:"id,omitempty"`
	PhoneNumber    string                    `json:"phoneNumber,omitempty"`
	Value          float64                   `json:"value,omitempty"`
	Status         MobilePhoneRechargeStatus `json:"status,omitempty"`
	CanBeCancelled bool                      `json:"canBeCancelled,omitempty"`
	OperatorName   string                    `json:"operatorName,omitempty"`
	Errors         []ErrorResponse           `json:"errors,omitempty"`
}

type MobilePhoneProviderResponse struct {
	Name   string                             `json:"name,omitempty"`
	Values []MobilePhoneProviderValueResponse `json:"values,omitempty"`
	Errors []ErrorResponse                    `json:"errors,omitempty"`
}

type MobilePhoneProviderValueResponse struct {
	Name        string  `json:"name,omitempty"`
	Description string  `json:"description,omitempty"`
	Bonus       string  `json:"bonus,omitempty"`
	MinValue    float64 `json:"minValue,omitempty"`
	MaxValue    float64 `json:"maxValue,omitempty"`
}

type mobilePhone struct {
	env         Env
	accessToken string
}

type MobilePhone interface {
	Recharge(ctx context.Context, body MobilePhoneRechargeRequest) (*MobilePhoneRechargeResponse, Error)
	CancelRechargeByID(ctx context.Context, rechargeID string) (*MobilePhoneRechargeResponse, Error)
	GetRechargeByID(ctx context.Context, rechargeID string) (*MobilePhoneRechargeResponse, Error)
	GetProviderByPhoneNumber(ctx context.Context, phoneNumber string) (*MobilePhoneProviderResponse, Error)
	GetAllRecharges(ctx context.Context, filter PageableDefaultRequest) (*Pageable[MobilePhoneRechargeResponse], Error)
}

func NewMobilePhone(env Env, accessToken string) MobilePhone {
	return mobilePhone{
		env:         env,
		accessToken: accessToken,
	}
}

func (m mobilePhone) Recharge(ctx context.Context, body MobilePhoneRechargeRequest) (*MobilePhoneRechargeResponse, Error) {
	if err := Validate().Struct(body); err != nil {
		return nil, NewError(ErrorTypeValidation, err)
	}
	req := NewRequest[MobilePhoneRechargeResponse](ctx, m.env, m.accessToken)
	return req.make(http.MethodPost, "/v3/mobilePhoneRecharges", body)
}

func (m mobilePhone) CancelRechargeByID(ctx context.Context, rechargeID string) (*MobilePhoneRechargeResponse, Error) {
	req := NewRequest[MobilePhoneRechargeResponse](ctx, m.env, m.accessToken)
	return req.make(http.MethodPost, fmt.Sprintf("/v3/mobilePhoneRecharges/%s/cancel", rechargeID), nil)
}

func (m mobilePhone) GetRechargeByID(ctx context.Context, rechargeID string) (*MobilePhoneRechargeResponse, Error) {
	req := NewRequest[MobilePhoneRechargeResponse](ctx, m.env, m.accessToken)
	return req.make(http.MethodGet, fmt.Sprintf("/v3/mobilePhoneRecharges/%s", rechargeID), nil)
}

func (m mobilePhone) GetProviderByPhoneNumber(ctx context.Context, phoneNumber string) (*MobilePhoneProviderResponse,
	Error) {
	req := NewRequest[MobilePhoneProviderResponse](ctx, m.env, m.accessToken)
	return req.make(http.MethodGet, fmt.Sprintf("/v3/mobilePhoneRecharges/%s/provider", phoneNumber), nil)
}

func (m mobilePhone) GetAllRecharges(ctx context.Context, filter PageableDefaultRequest) (
	*Pageable[MobilePhoneRechargeResponse], Error) {
	req := NewRequest[Pageable[MobilePhoneRechargeResponse]](ctx, m.env, m.accessToken)
	return req.make(http.MethodGet, "/v3/mobilePhoneRecharges", filter)
}
