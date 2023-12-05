package asaas

import (
	"context"
	"testing"
	"time"
)

func TestPixPayQrCode(t *testing.T) {
	initPixCharge()
	accessToken := getEnvValue(EnvAccessToken)
	pixQrCodePayload := getEnvValue(EnvChargePixQrCodePayload)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nPix := NewPix(EnvSandbox, accessToken)
	resp, err := nPix.PayQrCode(ctx, PayPixQrCodeRequest{
		QrCode: PixQrCodeRequest{
			Payload:     pixQrCodePayload,
			ChangeValue: 0,
		},
		Value:        5,
		Description:  "",
		ScheduleDate: Date{},
	})
	assertResponseSuccess(t, resp, err)
}

func TestPixDecodeQrCode(t *testing.T) {
	initPixCharge()
	accessToken := getEnvValue(EnvAccessToken)
	pixQrCodePayload := getEnvValue(EnvChargePixQrCodePayload)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nPix := NewPix(EnvSandbox, accessToken)
	resp, err := nPix.DecodeQrCode(ctx, PixQrCodeRequest{
		Payload:     pixQrCodePayload,
		ChangeValue: 0,
	})
	assertResponseSuccess(t, resp, err)
}

func TestPixCancelTransactionById(t *testing.T) {
	initPixTransaction()
	accessToken := getEnvValue(EnvAccessToken)
	pixTransactionId := getEnvValue(EnvPixTransactionId)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nPix := NewPix(EnvSandbox, accessToken)
	resp, err := nPix.CancelTransactionById(ctx, pixTransactionId)
	assertResponseSuccess(t, resp, err)
}

func TestPixCreateKey(t *testing.T) {
	accessToken := getEnvValue(EnvAccessToken)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nPix := NewPix(EnvSandbox, accessToken)
	resp, err := nPix.CreateKey(ctx)
	assertResponseSuccess(t, resp, err)
}

func TestPixCreateStaticKey(t *testing.T) {
	accessToken := getEnvValue(EnvAccessToken)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	now := DatetimeNow()
	nPix := NewPix(EnvSandbox, accessToken)
	resp, err := nPix.CreateStaticKey(ctx, CreatePixKeyStaticRequest{
		AddressKey:  "",
		Description: "",
		Value:       0,
		Format:      QrCodeFormatPayload,
		ExpirationDate: NewDatetimePointer(now.Year(), now.Month(), now.Day()+1, now.Hour(), now.Minute(), now.Second(),
			now.Nanosecond(), now.Location()),
		ExpirationSeconds:      0,
		AllowsMultiplePayments: Pointer(false),
	})
	assertResponseSuccess(t, resp, err)
}

func TestPixCreateStaticKeyWithoutExpiration(t *testing.T) {
	accessToken := getEnvValue(EnvAccessToken)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nPix := NewPix(EnvSandbox, accessToken)
	resp, err := nPix.CreateStaticKey(ctx, CreatePixKeyStaticRequest{
		AddressKey:             "",
		Description:            "",
		Value:                  0,
		Format:                 QrCodeFormatPayload,
		ExpirationSeconds:      0,
		AllowsMultiplePayments: Pointer(false),
	})
	assertResponseSuccess(t, resp, err)
}

func TestPixDeleteKeyById(t *testing.T) {
	initPixKey()
	accessToken := getEnvValue(EnvAccessToken)
	pixKeyId := getEnvValue(EnvPixKeyId)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nPix := NewPix(EnvSandbox, accessToken)
	resp, err := nPix.DeleteKeyById(ctx, pixKeyId)
	assertResponseSuccess(t, resp, err)
}

func TestPixGetTransactionById(t *testing.T) {
	initPixTransaction()
	accessToken := getEnvValue(EnvAccessToken)
	pixTransactionId := getEnvValue(EnvPixTransactionId)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nPix := NewPix(EnvSandbox, accessToken)
	resp, err := nPix.GetTransactionById(ctx, pixTransactionId)
	assertResponseSuccess(t, resp, err)
}

func TestPixGetKeyById(t *testing.T) {
	initPixKey()
	accessToken := getEnvValue(EnvAccessToken)
	pixKeyId := getEnvValue(EnvPixKeyId)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nPix := NewPix(EnvSandbox, accessToken)
	resp, err := nPix.GetKeyById(ctx, pixKeyId)
	assertResponseSuccess(t, resp, err)
}

func TestPixGetAllTransactions(t *testing.T) {
	initPixTransaction()
	accessToken := getEnvValue(EnvAccessToken)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
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
	initPixKey()
	accessToken := getEnvValue(EnvAccessToken)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nPix := NewPix(EnvSandbox, accessToken)
	resp, err := nPix.GetAllKeys(ctx, GetAllPixKeysRequest{
		Status:     "",
		StatusList: "",
		Offset:     0,
		Limit:      10,
	})
	assertResponseSuccess(t, resp, err)
}
