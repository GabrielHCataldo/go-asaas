package asaas

import (
	"context"
	"testing"
	"time"
)

func TestWebhookSaveSetting(t *testing.T) {
	accessToken := getEnvValue(EnvAccessTokenSecondary)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nWebhook := NewWebhook(EnvSandbox, accessToken)
	resp, err := nWebhook.SaveSetting(ctx, WebhookTypePayment, SaveWebhookSettingRequest{
		Url:         "https://test.com",
		Email:       "test@gmail.com",
		ApiVersion:  "3",
		Enabled:     Pointer(false),
		Interrupted: Pointer(false),
		AuthToken:   "",
	})
	assertResponseSuccess(t, resp, err)
}

func TestWebhookGetSetting(t *testing.T) {
	initWebhook()
	accessToken := getEnvValue(EnvAccessTokenSecondary)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nWebhook := NewWebhook(EnvSandbox, accessToken)
	resp, err := nWebhook.GetSetting(ctx, WebhookTypePayment)
	assertResponseSuccess(t, resp, err)
}
