package asaas

import (
	"context"
	"testing"
	"time"
)

func TestPixPayQrCode(t *testing.T) {
	accessToken := getEnvValue(EnvAccessToken)
	assertFatalStringBlank(t, accessToken)
	initPixCharge()
	pixQrCodePayload := getEnvValue(EnvChargePixQrCodePayload)
	assertFatalStringBlank(t, pixQrCodePayload)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nPix := NewPix(EnvSandbox, accessToken)
	resp, err := nPix.PayQrCode(ctx, PayPixQrCodeRequest{
		QrCode: PixQrCodeRequest{
			Payload:     pixQrCodePayload,
			ChangeValue: 0,
		},
		Value:        100,
		Description:  "",
		ScheduleDate: Date{},
	})
	assertResponseSuccess(t, resp, err)
}

func TestPixDecodeQrCode(t *testing.T) {
	accessToken := getEnvValue(EnvAccessToken)
	assertFatalStringBlank(t, accessToken)
	initPixCharge()
	pixQrCodePayload := getEnvValue(EnvChargePixQrCodePayload)
	assertFatalStringBlank(t, pixQrCodePayload)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nPix := NewPix(EnvSandbox, accessToken)
	resp, err := nPix.DecodeQrCode(ctx, PixQrCodeRequest{
		Payload:     pixQrCodePayload,
		ChangeValue: 0,
	})
	assertResponseSuccess(t, resp, err)
}

func TestPixCancelTransactionById(t *testing.T) {
	accessToken := getEnvValue(EnvAccessToken)
	assertFatalStringBlank(t, accessToken)
	initPixTransaction()
	pixTransactionId := getEnvValue(EnvPixTransactionId)
	assertFatalStringBlank(t, pixTransactionId)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nPix := NewPix(EnvSandbox, accessToken)
	resp, err := nPix.CancelTransactionById(ctx, pixTransactionId)
	assertResponseSuccess(t, resp, err)
}

func TestPixCreateKey(t *testing.T) {
	accessToken := getEnvValue(EnvAccessToken)
	assertFatalStringBlank(t, accessToken)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nPix := NewPix(EnvSandbox, accessToken)
	resp, err := nPix.CreateKey(ctx)
	assertResponseSuccess(t, resp, err)
}

func TestPixCreateStaticKey(t *testing.T) {
	accessToken := getEnvValue(EnvAccessToken)
	assertFatalStringBlank(t, accessToken)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	now := time.Now()
	nPix := NewPix(EnvSandbox, accessToken)
	resp, err := nPix.CreateStaticKey(ctx, CreatePixKeyStaticRequest{
		AddressKey:  "",
		Description: "",
		Value:       0,
		Format:      QrCodeFormatPayload,
		ExpirationDate: NewDatetimePointer(now.Year(), now.Month(), now.Day()+1, now.Hour(), now.Minute(), 0, 0,
			now.Location()),
		ExpirationSeconds:      0,
		AllowsMultiplePayments: Pointer(false),
	})
	assertResponseSuccess(t, resp, err)
}

func TestPixDeleteKeyById(t *testing.T) {
	accessToken := getEnvValue(EnvAccessToken)
	assertFatalStringBlank(t, accessToken)
	initPixKey()
	pixKeyId := getEnvValue(EnvPixKeyId)
	assertFatalStringBlank(t, pixKeyId)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nPix := NewPix(EnvSandbox, accessToken)
	resp, err := nPix.DeleteKeyById(ctx, pixKeyId)
	assertResponseSuccess(t, resp, err)
}

func TestPixGetTransactionById(t *testing.T) {
	accessToken := getEnvValue(EnvAccessToken)
	assertFatalStringBlank(t, accessToken)
	initPixTransaction()
	pixTransactionId := getEnvValue(EnvPixTransactionId)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nPix := NewPix(EnvSandbox, accessToken)
	resp, err := nPix.GetTransactionById(ctx, pixTransactionId)
	assertResponseSuccess(t, resp, err)
}

func TestPixGetKeyById(t *testing.T) {
	accessToken := getEnvValue(EnvAccessToken)
	assertFatalStringBlank(t, accessToken)
	initPixKey()
	pixKeyId := getEnvValue(EnvPixKeyId)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nPix := NewPix(EnvSandbox, accessToken)
	resp, err := nPix.GetKeyById(ctx, pixKeyId)
	assertResponseSuccess(t, resp, err)
}

func TestPixGetAllTransactions(t *testing.T) {
	accessToken := getEnvValue(EnvAccessToken)
	assertFatalStringBlank(t, accessToken)
	initPixTransaction()
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nPix := NewPix(EnvSandbox, accessToken)
	resp, err := nPix.GetAllTransactions(ctx, GetAllPixTransactionsRequest{
		Status:             "",
		Type:               "",
		EndToEndIdentifier: "",
		Offset:             0,
		Limit:              10,
	})
	assertResponseSuccess(t, resp, err)
}

func TestPixGetAllKeys(t *testing.T) {
	accessToken := getEnvValue(EnvAccessToken)
	assertFatalStringBlank(t, accessToken)
	initPixKey()
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nPix := NewPix(EnvSandbox, accessToken)
	resp, errAsaas := nPix.GetAllKeys(ctx, GetAllPixKeysRequest{
		Status:     "",
		StatusList: "",
		Offset:     0,
		Limit:      10,
	})
	assertResponseSuccess(t, resp, errAsaas)
}
