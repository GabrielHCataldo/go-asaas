package asaas

import (
	"context"
	"github.com/GabrielHCataldo/go-asaas/internal/test"
	"testing"
	"time"
)

func TestPixGetAllKeysSuccess(t *testing.T) {
	accessToken, err := test.GetAccessTokenByEnv()
	assertFatalErrorNonnull(t, err)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nPix := NewPix(SANDBOX, *accessToken)
	resp, errAsaas := nPix.GetAllKeys(ctx, GetAllPixKeysRequest{})
	assertResponseSuccess(t, resp, errAsaas)
}
