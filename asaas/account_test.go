package asaas

import (
	"context"
	"github.com/GabrielHCataldo/go-asaas/internal/test"
	"testing"
	"time"
)

func TestAccountGet(t *testing.T) {
	accessToken, err := test.GetAccessTokenByEnv()
	assertFatalErrorNonnull(t, err)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nAccount := NewAccount(EnvSandbox, *accessToken)
	resp, errAsaas := nAccount.Get(ctx)
	assertResponseSuccess(t, resp, errAsaas)
}

func TestAccountGetBalance(t *testing.T) {
	accessToken, err := test.GetAccessTokenByEnv()
	assertFatalErrorNonnull(t, err)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nAccount := NewAccount(EnvSandbox, *accessToken)
	resp, errAsaas := nAccount.GetBalance(ctx)
	assertResponseSuccess(t, resp, errAsaas)
}
