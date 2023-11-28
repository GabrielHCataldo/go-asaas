package asaas

import (
	"context"
	"github.com/GabrielHCataldo/go-asaas/internal/test"
	"testing"
	"time"
)

func TestAnticipationGetLimitsSuccess(t *testing.T) {
	accessToken, err := test.GetAccessTokenByEnv()
	assertFatalErrorNonnull(t, err)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nAnticipation := NewAnticipation(EnvSandbox, *accessToken)
	resp, errAsaas := nAnticipation.GetLimits(ctx)
	assertResponseSuccess(t, resp, errAsaas)
}
