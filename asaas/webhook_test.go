package asaas

import (
	"context"
	"testing"
	"time"
)

func TestWebhookSaveSettingPayment(t *testing.T) {
	accessToken := getEnvValue(EnvAccessTokenSecondary)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nWebhook := NewWebhook(EnvSandbox, accessToken)
	resp, err := nWebhook.SaveSetting(ctx, WebhookTypePayment, getWebhookSaveSettingRequest())
	assertResponseSuccess(t, resp, err)
}

func TestWebhookSaveSettingInvoice(t *testing.T) {
	accessToken := getEnvValue(EnvAccessTokenSecondary)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nWebhook := NewWebhook(EnvSandbox, accessToken)
	resp, err := nWebhook.SaveSetting(ctx, WebhookTypeInvoice, getWebhookSaveSettingRequest())
	assertResponseSuccess(t, resp, err)
}

func TestWebhookSaveSettingTransfer(t *testing.T) {
	accessToken := getEnvValue(EnvAccessTokenSecondary)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nWebhook := NewWebhook(EnvSandbox, accessToken)
	resp, err := nWebhook.SaveSetting(ctx, WebhookTypeTransfer, getWebhookSaveSettingRequest())
	assertResponseSuccess(t, resp, err)
}

func TestWebhookSaveSettingBill(t *testing.T) {
	accessToken := getEnvValue(EnvAccessTokenSecondary)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nWebhook := NewWebhook(EnvSandbox, accessToken)
	resp, err := nWebhook.SaveSetting(ctx, WebhookTypeBill, getWebhookSaveSettingRequest())
	assertResponseSuccess(t, resp, err)
}

func TestWebhookSaveSettingReceivableAnticipation(t *testing.T) {
	accessToken := getEnvValue(EnvAccessTokenSecondary)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nWebhook := NewWebhook(EnvSandbox, accessToken)
	resp, err := nWebhook.SaveSetting(ctx, WebhookTypeReceivableAnticipation, getWebhookSaveSettingRequest())
	assertResponseSuccess(t, resp, err)
}

func TestWebhookSaveSettingMobilePhoneRecharge(t *testing.T) {
	accessToken := getEnvValue(EnvAccessTokenSecondary)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nWebhook := NewWebhook(EnvSandbox, accessToken)
	resp, err := nWebhook.SaveSetting(ctx, WebhookTypeMobilePhoneRecharge, getWebhookSaveSettingRequest())
	assertResponseSuccess(t, resp, err)
}

func TestWebhookSaveSettingWebhookTypeAccountStatus(t *testing.T) {
	accessToken := getEnvValue(EnvAccessTokenSecondary)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nWebhook := NewWebhook(EnvSandbox, accessToken)
	resp, err := nWebhook.SaveSetting(ctx, WebhookTypeAccountStatus, getWebhookSaveSettingRequest())
	assertResponseSuccess(t, resp, err)
}

func TestWebhookGetSetting(t *testing.T) {
	initWebhook()
	accessToken := getEnvValue(EnvAccessTokenSecondary)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nWebhook := NewWebhook(EnvSandbox, accessToken)
	resp, err := nWebhook.GetSetting(ctx, WebhookTypePayment)
	assertResponseSuccess(t, resp, err)
}

func getWebhookSaveSettingRequest() SaveWebhookSettingRequest {
	return SaveWebhookSettingRequest{
		Url:         "https://test.com",
		Email:       "test@gmail.com",
		ApiVersion:  "3",
		Enabled:     Pointer(false),
		Interrupted: Pointer(false),
		AuthToken:   "",
	}
}
