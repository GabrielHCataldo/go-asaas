package asaas

import (
	"context"
	"github.com/GabrielHCataldo/go-asaas/internal/test"
	"os"
	"testing"
	"time"
)

func TestAnticipationRequestSuccess(t *testing.T) {
	accessToken, err := test.GetAccessTokenByEnv()
	assertFatalErrorNonnull(t, err)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	f, err := test.GetSimpleFile()
	f2, err := test.GetSimpleFile()
	f3, err := test.GetSimpleFile()
	assertFatalErrorNonnull(t, err)
	nAnticipation := NewAnticipation(EnvSandbox, *accessToken)
	resp, errAsaas := nAnticipation.Request(ctx, AnticipationRequest{
		Payment:   test.GetChargeIdDefault(),
		Documents: []*os.File{f, f2, f3},
	})
	assertResponseSuccess(t, resp, errAsaas)
}

func TestAnticipationGetLimitsSuccess(t *testing.T) {
	accessToken, err := test.GetAccessTokenByEnv()
	assertFatalErrorNonnull(t, err)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nAnticipation := NewAnticipation(EnvSandbox, *accessToken)
	resp, errAsaas := nAnticipation.GetLimits(ctx)
	assertResponseSuccess(t, resp, errAsaas)
}
