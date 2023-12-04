package asaas

import (
	"context"
	"encoding/json"
	"github.com/GabrielHCataldo/go-asaas/internal/test"
	"testing"
	"time"
)

func TestCustomerCreateSuccess(t *testing.T) {
	accessToken, err := test.GetAccessTokenByEnv()
	assertFatalErrorNonnull(t, err)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	req := &CreateCustomerRequest{}
	err = json.Unmarshal(test.GetCreateCustomerRequestDefault(), req)
	assertFatalErrorNonnull(t, err)
	nCustomer := NewCustomer(EnvSandbox, *accessToken)
	resp, errAsaas := nCustomer.Create(ctx, *req)
	assertResponseSuccess(t, resp, errAsaas)
}

func TestCustomerCreateFailure(t *testing.T) {
	accessToken, err := test.GetAccessTokenByEnv()
	assertFatalErrorNonnull(t, err)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nCustomer := NewCustomer(EnvSandbox, *accessToken)
	_, errAsaas := nCustomer.Create(ctx, CreateCustomerRequest{})
	assertSuccessNonnull(t, errAsaas)
}
