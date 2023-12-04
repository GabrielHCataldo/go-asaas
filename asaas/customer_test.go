package asaas

import (
	"context"
	"encoding/json"
	"github.com/GabrielHCataldo/go-asaas/internal/test"
	"testing"
	"time"
)

func TestCustomerCreateSuccess(t *testing.T) {
	accessToken := getEnvValue(EnvAccessToken)
	assertFatalStringBlank(t, accessToken)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	req := &CreateCustomerRequest{}
	err := json.Unmarshal(test.GetCreateCustomerRequestDefault(), req)
	assertFatalErrorNonnull(t, err)
	nCustomer := NewCustomer(EnvSandbox, accessToken)
	resp, errAsaas := nCustomer.Create(ctx, *req)
	assertResponseSuccess(t, resp, errAsaas)
}

func TestCustomerCreateError(t *testing.T) {
	nCustomer := NewCustomer(EnvSandbox, "")
	resp, err := nCustomer.Create(context.TODO(), CreateCustomerRequest{})
	assertResponseFailure(t, resp, err)
}
