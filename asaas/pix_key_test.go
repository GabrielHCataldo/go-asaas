package asaas

import (
	"context"
	"github.com/GabrielHCataldo/go-asaas/internal/test"
	"testing"
	"time"
)

func TestPixKeyGetAllSuccess(t *testing.T) {
	accessToken, err := test.GetAccessTokenByEnv()
	assertFatalErrorNonnull(t, err)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nPixKey := NewPixKey(SANDBOX, *accessToken)
	resp, errAsaas := nPixKey.GetAll(ctx, GetAllPixKeysRequest{})
	assertResponseSuccess(t, resp, errAsaas)
}
