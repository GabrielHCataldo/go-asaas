package asaas

import (
	"context"
	"encoding/json"
	"github.com/GabrielHCataldo/go-asaas/internal/test"
	"testing"
	"time"
)

func TestCreditCardTokenizeSuccess(t *testing.T) {
	accessToken := getEnvValue(EnvAccessToken)
	assertFatalStringBlank(t, accessToken)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	req := &CreditCardTokenizeRequest{}
	err := json.Unmarshal(test.GetCreditCardRequestDefault(), req)
	assertFatalErrorNonnull(t, err)
	nCreditCard := NewCreditCard(EnvSandbox, accessToken)
	resp, errAsaas := nCreditCard.Tokenize(ctx, *req)
	assertResponseSuccess(t, resp, errAsaas)
}

func TestCreditCardTokenizeFailure(t *testing.T) {
	accessToken := getEnvValue(EnvAccessToken)
	assertFatalStringBlank(t, accessToken)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	req := &CreditCardTokenizeRequest{}
	err := json.Unmarshal(test.GetCreditCardFailureRequestDefault(), req)
	assertFatalErrorNonnull(t, err)
	nCreditCard := NewCreditCard(EnvSandbox, accessToken)
	resp, errAsaas := nCreditCard.Tokenize(ctx, *req)
	assertResponseFailure(t, resp, errAsaas)
}
