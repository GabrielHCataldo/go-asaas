package asaas

import (
	"context"
	"github.com/GabrielHCataldo/go-asaas/internal/test"
	"testing"
	"time"
)

func TestSubscriptionGetAllNoContent(t *testing.T) {
	accessToken, err := test.GetAccessTokenByEnv()
	assertFatalErrorNonnull(t, err)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nSubscription := NewSubscription(EnvSandbox, accessToken)
	resp, errAsaas := nSubscription.GetAll(ctx, GetAllSubscriptionsRequest{})
	assertResponseNoContent(t, resp, errAsaas)
}
