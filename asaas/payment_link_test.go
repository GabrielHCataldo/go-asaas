package asaas

import (
	"context"
	"github.com/GabrielHCataldo/go-asaas/internal/test"
	"testing"
	"time"
)

func TestPaymentLinkSendImageById(t *testing.T) {
	accessToken, err := getAccessToken()
	assertFatalErrorNonnull(t, err)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	f, err := test.GetSimpleFile()
	assertFatalErrorNonnull(t, err)
	nPaymentLink := NewPaymentLink(EnvSandbox, accessToken)
	resp, errAsaas := nPaymentLink.SendImageById(ctx, test.GetPaymentLinkIdDefault(), SendImagePaymentLinksRequest{
		Main:  true,
		Image: f,
	})
	assertResponseSuccess(t, resp, errAsaas)
}

func TestPaymentLinkGetById(t *testing.T) {
	accessToken, err := getAccessToken()
	assertFatalErrorNonnull(t, err)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nPaymentLink := NewPaymentLink(EnvSandbox, accessToken)
	resp, errAsaas := nPaymentLink.GetById(ctx, test.GetPaymentLinkIdDefault())
	assertResponseSuccess(t, resp, errAsaas)
}
func TestPaymentLinkGetAll(t *testing.T) {
	accessToken, err := getAccessToken()
	assertFatalErrorNonnull(t, err)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nPaymentLink := NewPaymentLink(EnvSandbox, accessToken)
	resp, errAsaas := nPaymentLink.GetAll(ctx, GetAllPaymentLinksRequest{})
	assertResponseSuccess(t, resp, errAsaas)
}
