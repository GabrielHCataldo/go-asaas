package asaas

import (
	"context"
	"testing"
	"time"
)

func TestFiscalInfoGetAllServices(t *testing.T) {
	accessToken := getEnvValue(EnvAccessToken)
	assertFatalIsBlank(t, accessToken)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nFiscalInfo := NewFiscalInfo(EnvSandbox, accessToken)
	resp, errAsaas := nFiscalInfo.GetAllServices(ctx, GetAllServicesRequest{})
	assertResponseSuccess(t, resp, errAsaas)
}
