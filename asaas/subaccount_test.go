package asaas

import (
	"context"
	"github.com/GabrielHCataldo/go-asaas/internal/test"
	"testing"
	"time"
)

func TestSubaccountGetAll(t *testing.T) {
	accessToken, err := test.GetAccessTokenByEnv()
	assertFatalErrorNonnull(t, err)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nSubaccount := NewSubaccount(EnvSandbox, accessToken)
	resp, errAsaas := nSubaccount.GetAll(ctx, GetAllSubaccountsRequest{})
	assertResponseNoContent(t, resp, errAsaas)
}
