package asaas

import (
	"context"
	"testing"
	"time"
)

func TestWebhookGetSettingNoContent(t *testing.T) {
	accessToken, err := getAccessToken()
	assertFatalErrorNonnull(t, err)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nWebhook := NewWebhook(EnvSandbox, accessToken)
	resp, errAsaas := nWebhook.GetSetting(ctx, WebhookTypePayment)
	assertResponseNoContent(t, resp, errAsaas)
}
