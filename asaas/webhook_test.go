package asaas

import (
	"context"
	"testing"
	"time"
)

func TestWebhookSaveSetting(t *testing.T) {
	accessToken := getEnvValue(EnvAccessTokenSecondary)
	assertFatalStringBlank(t, accessToken)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nWebhook := NewWebhook(EnvSandbox, accessToken)
	resp, errAsaas := nWebhook.SaveSetting(ctx, WebhookTypePayment, SaveWebhookSettingRequest{
		Url:         "https://test.com",
		Email:       "test@gmail.com",
		ApiVersion:  "3",
		Enabled:     Pointer(false),
		Interrupted: Pointer(false),
		AuthToken:   "",
	})
	assertResponseNoContent(t, resp, errAsaas)
}

func TestWebhookGetSetting(t *testing.T) {
	accessToken := getEnvValue(EnvAccessTokenSecondary)
	assertFatalStringBlank(t, accessToken)
	initWebhook()
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nWebhook := NewWebhook(EnvSandbox, accessToken)
	resp, errAsaas := nWebhook.GetSetting(ctx, WebhookTypePayment)
	assertResponseNoContent(t, resp, errAsaas)
}
