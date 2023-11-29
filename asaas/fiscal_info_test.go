package asaas

import (
	"context"
	"github.com/GabrielHCataldo/go-asaas/internal/test"
	"testing"
	"time"
)

func TestFiscalInfoGetAllServices(t *testing.T) {
	accessToken, err := test.GetAccessTokenByEnv()
	assertFatalErrorNonnull(t, err)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nFiscalInfo := NewFiscalInfo(EnvSandbox, *accessToken)
	resp, errAsaas := nFiscalInfo.GetAllServices(ctx, GetAllServicesRequest{})
	assertResponseSuccess(t, resp, errAsaas)
}
