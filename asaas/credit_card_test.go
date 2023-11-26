package asaas

import (
	"context"
	"encoding/json"
	"github.com/GabrielHCataldo/go-asaas/internal/test"
	"testing"
	"time"
)

func TestCreditCardTokenizeSuccess(t *testing.T) {
	accessToken, err := test.GetAccessTokenByEnv()
	AssertFatalErrorNonnull(t, err)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	req := &CreditCardTokenizeRequest{}
	err = json.Unmarshal(test.GetCreditCardRequestDefault(), req)
	AssertFatalErrorNonnull(t, err)
	nCreditCard := NewCreditCard(SANDBOX, *accessToken)
	resp, errAsaas := nCreditCard.Tokenize(ctx, *req)
	AssertAsaasResponseSuccess(t, resp, errAsaas)
}

func TestCreditCardTokenizeFailure(t *testing.T) {
	accessToken, err := test.GetAccessTokenByEnv()
	AssertFatalErrorNonnull(t, err)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	req := &CreditCardTokenizeRequest{}
	err = json.Unmarshal(test.GetCreditCardFailureRequestDefault(), req)
	AssertFatalErrorNonnull(t, err)
	nCreditCard := NewCreditCard(SANDBOX, *accessToken)
	resp, errAsaas := nCreditCard.Tokenize(ctx, *req)
	AssertAsaasResponseFailure(t, resp, errAsaas)
}
