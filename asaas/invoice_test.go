package asaas

import (
	"context"
	"testing"
	"time"
)

func TestInvoiceGetAll(t *testing.T) {
	accessToken, err := getAccessToken()
	assertFatalErrorNonnull(t, err)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nInvoice := NewInvoice(EnvSandbox, accessToken)
	resp, errAsaas := nInvoice.GetAll(ctx, GetAllInvoicesRequest{})
	assertResponseSuccess(t, resp, errAsaas)
}
