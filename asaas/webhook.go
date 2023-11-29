package asaas

import (
	"context"
	"net/http"
)

type SaveWebhookSettingRequest struct {
	Url         string `json:"url,omitempty" validate:"required,url"`
	Email       string `json:"email,omitempty" validate:"required,email"`
	ApiVersion  string `json:"apiVersion,omitempty" validate:"required,numeric,max=4"`
	Enabled     bool   `json:"enabled,omitempty"`
	Interrupted bool   `json:"interrupted,omitempty"`
	AuthToken   string `json:"authToken,omitempty"`
}

type WebhookResponse struct {
	Type        TypeOfWebhook   `json:"type,omitempty"`
	Url         string          `json:"url,omitempty"`
	Email       string          `json:"email,omitempty"`
	ApiVersion  string          `json:"apiVersion,omitempty"`
	Enabled     bool            `json:"enabled,omitempty"`
	Interrupted bool            `json:"interrupted,omitempty"`
	AuthToken   string          `json:"authToken,omitempty"`
	Errors      []ErrorResponse `json:"errors,omitempty"`
}

type webhook struct {
	env         Env
	accessToken string
}

type Webhook interface {
	SaveSetting(ctx context.Context, typeWebhook TypeOfWebhook, body SaveWebhookSettingRequest) (*WebhookResponse, Error)
	GetSetting(ctx context.Context, typeWebhook TypeOfWebhook) (*WebhookResponse, Error)
}

func NewWebhook(env Env, accessToken string) Webhook {
	logWarning("Webhook service running on", env.String())
	return webhook{
		env:         env,
		accessToken: accessToken,
	}
}

func (w webhook) SaveSetting(ctx context.Context, typeWebhook TypeOfWebhook, body SaveWebhookSettingRequest) (
	*WebhookResponse, Error) {
	if !typeWebhook.IsEnumValid() {
		return nil, NewError(ErrorTypeValidation, "invalid typeWebhook")
	} else if err := Validate().Struct(body); err != nil {
		return nil, NewError(ErrorTypeValidation, err)
	}
	req := NewRequest[WebhookResponse](ctx, w.env, w.accessToken)
	return req.make(http.MethodPost, "/v3/webhook"+typeWebhook.PathUrl(), body)
}

func (w webhook) GetSetting(ctx context.Context, typeWebhook TypeOfWebhook) (*WebhookResponse, Error) {
	if !typeWebhook.IsEnumValid() {
		return nil, NewError(ErrorTypeValidation, "invalid typeWebhook")
	}
	req := NewRequest[WebhookResponse](ctx, w.env, w.accessToken)
	return req.make(http.MethodGet, "/v3/webhook"+typeWebhook.PathUrl(), nil)
}
