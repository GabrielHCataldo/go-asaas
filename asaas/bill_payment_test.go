package asaas

import (
	"context"
	"github.com/GabrielHCataldo/go-asaas/internal/test"
	"testing"
	"time"
)

func TestBillCreateFailure(t *testing.T) {
	accessToken, err := test.GetAccessTokenByEnv()
	assertFatalErrorNonnull(t, err)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nBillPayment := NewBillPayment(EnvSandbox, *accessToken)
	_, errAsaas := nBillPayment.Create(ctx, BillPaymentRequest{})
	assertSuccessNonnull(t, errAsaas)
}
