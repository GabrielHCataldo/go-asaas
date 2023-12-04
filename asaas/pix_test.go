package asaas

import (
	"context"
	"testing"
	"time"
)

func TestPixGetAllKeysSuccess(t *testing.T) {
	accessToken, err := getAccessToken()
	assertFatalErrorNonnull(t, err)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nPix := NewPix(EnvSandbox, accessToken)
	resp, errAsaas := nPix.GetAllKeys(ctx, GetAllPixKeysRequest{})
	assertResponseSuccess(t, resp, errAsaas)
}
