package asaas

import (
	"context"
	"github.com/GabrielHCataldo/go-asaas/internal/test"
	"os"
	"testing"
	"time"
)

func TestPaymentLinkSendImageById(t *testing.T) {
	accessToken := getEnvValue(EnvAccessToken)
	assertFatalIsBlank(t, accessToken)
	f, err := os.Open(getEnvValue(EnvFileName))
	assertFatalErrorNonnull(t, err)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nPaymentLink := NewPaymentLink(EnvSandbox, accessToken)
	resp, errAsaas := nPaymentLink.SendImageById(ctx, test.GetPaymentLinkIdDefault(), SendImagePaymentLinksRequest{
		Main:  true,
		Image: f,
	})
	assertResponseSuccess(t, resp, errAsaas)
}

func TestPaymentLinkGetById(t *testing.T) {
	accessToken := getEnvValue(EnvAccessToken)
	assertFatalIsBlank(t, accessToken)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nPaymentLink := NewPaymentLink(EnvSandbox, accessToken)
	resp, errAsaas := nPaymentLink.GetById(ctx, test.GetPaymentLinkIdDefault())
	assertResponseSuccess(t, resp, errAsaas)
}
func TestPaymentLinkGetAll(t *testing.T) {
	accessToken := getEnvValue(EnvAccessToken)
	assertFatalIsBlank(t, accessToken)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nPaymentLink := NewPaymentLink(EnvSandbox, accessToken)
	resp, errAsaas := nPaymentLink.GetAll(ctx, GetAllPaymentLinksRequest{})
	assertResponseSuccess(t, resp, errAsaas)
}
