package asaas

import (
	"context"
	"testing"
	"time"
)

func TestInstallmentGetAllSuccess(t *testing.T) {
	accessToken := getEnvValue(EnvAccessToken)
	assertFatalIsBlank(t, accessToken)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nInstallment := NewInstallment(EnvSandbox, accessToken)
	resp, errAsaas := nInstallment.GetAll(ctx, PageableDefaultRequest{})
	assertResponseSuccess(t, resp, errAsaas)
}
