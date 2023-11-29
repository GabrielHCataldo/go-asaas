package asaas

import (
	"context"
	"github.com/GabrielHCataldo/go-asaas/internal/test"
	"testing"
	"time"
)

func TestPaymentLinkSendImageByID(t *testing.T) {
	accessToken, err := test.GetAccessTokenByEnv()
	assertFatalErrorNonnull(t, err)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	f, err := test.GetSimpleFile()
	assertFatalErrorNonnull(t, err)
	nPaymentLink := NewPaymentLink(EnvSandbox, *accessToken)
	resp, errAsaas := nPaymentLink.SendImageByID(ctx, test.GetPaymentLinkIdDefault(), SendImagePaymentLinksRequest{
		Main:  true,
		Image: f,
	})
	assertResponseSuccess(t, resp, errAsaas)
}

func TestPaymentLinkGetByID(t *testing.T) {
	accessToken, err := test.GetAccessTokenByEnv()
	assertFatalErrorNonnull(t, err)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nPaymentLink := NewPaymentLink(EnvSandbox, *accessToken)
	resp, errAsaas := nPaymentLink.GetByID(ctx, test.GetPaymentLinkIdDefault())
	assertResponseSuccess(t, resp, errAsaas)
}
func TestPaymentLinkGetAll(t *testing.T) {
	accessToken, err := test.GetAccessTokenByEnv()
	assertFatalErrorNonnull(t, err)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nPaymentLink := NewPaymentLink(EnvSandbox, *accessToken)
	resp, errAsaas := nPaymentLink.GetAll(ctx, GetAllPaymentLinksRequest{})
	assertResponseSuccess(t, resp, errAsaas)
}
