package asaas

import (
	"context"
	"github.com/GabrielHCataldo/go-asaas/internal/test"
	"testing"
	"time"
)

func TestNotificationGetAllByCustomerSuccess(t *testing.T) {
	accessToken, err := test.GetAccessTokenByEnv()
	assertFatalErrorNonnull(t, err)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nNotification := NewNotification(SANDBOX, *accessToken)
	resp, errAsaas := nNotification.GetAllByCustomer(ctx, test.GetCustomerIdDefault())
	assertResponseSuccess(t, resp, errAsaas)
}

func TestNotificationGetAllByCustomerNoContent(t *testing.T) {
	accessToken, err := test.GetAccessTokenByEnv()
	assertFatalErrorNonnull(t, err)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nNotification := NewNotification(SANDBOX, *accessToken)
	resp, errAsaas := nNotification.GetAllByCustomer(ctx, "test")
	assertResponseNoContent(t, resp, errAsaas)
}
