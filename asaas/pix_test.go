package asaas

import (
	"context"
	"testing"
	"time"
)

func TestPixGetAllKeys(t *testing.T) {
	accessToken := getEnvValue(EnvAccessToken)
	assertFatalStringBlank(t, accessToken)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nPix := NewPix(EnvSandbox, accessToken)
	resp, errAsaas := nPix.GetAllKeys(ctx, GetAllPixKeysRequest{})
	assertResponseSuccess(t, resp, errAsaas)
}
