package asaas

import (
	"context"
	"github.com/GabrielHCataldo/go-asaas/internal/test"
	"testing"
	"time"
)

func TestWebhookGetSettingNoContent(t *testing.T) {
	accessToken, err := test.GetAccessTokenByEnv()
	assertFatalErrorNonnull(t, err)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nWebhook := NewWebhook(EnvSandbox, *accessToken)
	resp, errAsaas := nWebhook.GetSetting(ctx, WebhookTypePayment)
	assertResponseNoContent(t, resp, errAsaas)
}
