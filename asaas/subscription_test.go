package asaas

import (
	"context"
	"testing"
	"time"
)

// TODO
func TestSubscription_Create(t *testing.T) {

}

func TestSubscription_CreateInvoiceSettingById(t *testing.T) {

}

func TestSubscription_UpdateById(t *testing.T) {

}

func TestSubscription_UpdateInvoiceSettingsById(t *testing.T) {

}

func TestSubscription_DeleteById(t *testing.T) {

}

func TestSubscription_DeleteInvoiceSettingById(t *testing.T) {

}

func TestSubscription_GetById(t *testing.T) {

}

func TestSubscription_GetInvoiceSettingById(t *testing.T) {

}

func TestSubscription_GetPaymentBookById(t *testing.T) {

}

func TestSubscription_GetAll(t *testing.T) {
	accessToken := getEnvValue(EnvAccessToken)
	assertFatalStringBlank(t, accessToken)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nSubscription := NewSubscription(EnvSandbox, accessToken)
	resp, errAsaas := nSubscription.GetAll(ctx, GetAllSubscriptionsRequest{})
	assertResponseNoContent(t, resp, errAsaas)
}

func TestSubscription_GetAllChargesBySubscription(t *testing.T) {

}

func TestSubscription_GetAllInvoicesBySubscription(t *testing.T) {

}
