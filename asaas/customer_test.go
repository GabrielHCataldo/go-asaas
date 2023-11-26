package asaas

import (
	"context"
	"encoding/json"
	"github.com/GabrielHCataldo/go-asaas/internal/test"
	"testing"
	"time"
)

func TestCustomerCreate(t *testing.T) {
	accessToken, err := test.GetAccessTokenByEnv()
	AssertFatalErrorNonnull(t, err)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	req := &CreateCustomerRequest{}
	err = json.Unmarshal(test.GetCreateCustomerRequestDefault(), req)
	nCustomer := NewCustomer(SANDBOX, *accessToken)
	resp, errAsaas := nCustomer.Create(ctx, *req)
	AssertAsaasResponseSuccess(t, resp, errAsaas)
}
