package asaas

import (
	"context"
	"os"
	"testing"
	"time"
)

func TestPaymentLinkCreate(t *testing.T) {
	accessToken := getEnvValue(EnvAccessToken)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nPaymentLink := NewPaymentLink(EnvSandbox, accessToken)
	resp, err := nPaymentLink.Create(ctx, CreatePaymentLinkRequest{
		Name:                "Unit test go",
		Description:         "",
		BillingType:         BillingTypeUndefined,
		ChargeType:          ChargeTypeDetached,
		EndDate:             Date{},
		Value:               0,
		DueDateLimitDays:    0,
		SubscriptionCycle:   "",
		MaxInstallmentCount: 0,
		NotificationEnabled: false,
		Callback:            nil,
	})
	assertResponseSuccess(t, resp, err)
}

func TestPaymentLinkSendImageById(t *testing.T) {
	initPaymentLink()
	accessToken := getEnvValue(EnvAccessToken)
	paymentLinkId := getEnvValue(EnvPaymentLinkId)
	f, _ := os.Open(getEnvValue(EnvImageName))
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nPaymentLink := NewPaymentLink(EnvSandbox, accessToken)
	resp, err := nPaymentLink.SendImageById(ctx, paymentLinkId, SendImagePaymentLinksRequest{
		Main:  false,
		Image: f,
	})
	assertResponseSuccess(t, resp, err)
}

func TestPaymentLinkUpdateById(t *testing.T) {
	initPaymentLink()
	accessToken := getEnvValue(EnvAccessToken)
	paymentLinkId := getEnvValue(EnvPaymentLinkId)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nPaymentLink := NewPaymentLink(EnvSandbox, accessToken)
	resp, err := nPaymentLink.UpdateById(ctx, paymentLinkId, UpdatePaymentLinkRequest{
		Name:                "Unit test go 2",
		Description:         nil,
		BillingType:         "",
		ChargeType:          "",
		EndDate:             Date{},
		Value:               nil,
		DueDateLimitDays:    0,
		SubscriptionCycle:   nil,
		MaxInstallmentCount: 0,
		NotificationEnabled: nil,
		Callback:            nil,
	})
	assertResponseSuccess(t, resp, err)
}

func TestPaymentLinkUpdateImageAsMainById(t *testing.T) {
	initPaymentLinkImage()
	accessToken := getEnvValue(EnvAccessToken)
	paymentLinkId := getEnvValue(EnvPaymentLinkId)
	paymentLinkImageId := getEnvValue(EnvPaymentLinkImageId)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nPaymentLink := NewPaymentLink(EnvSandbox, accessToken)
	resp, err := nPaymentLink.UpdateImageAsMainById(ctx, paymentLinkId, paymentLinkImageId)
	assertResponseSuccess(t, resp, err)
}

func TestPaymentLinkDeleteById(t *testing.T) {
	initPaymentLinkImage()
	accessToken := getEnvValue(EnvAccessToken)
	paymentLinkId := getEnvValue(EnvPaymentLinkId)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nPaymentLink := NewPaymentLink(EnvSandbox, accessToken)
	resp, err := nPaymentLink.DeleteById(ctx, paymentLinkId)
	assertResponseSuccess(t, resp, err)
}

func TestPaymentLinkRestoreById(t *testing.T) {
	initPaymentLinkDeleted()
	accessToken := getEnvValue(EnvAccessToken)
	paymentLinkId := getEnvValue(EnvPaymentLinkDeletedId)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nPaymentLink := NewPaymentLink(EnvSandbox, accessToken)
	resp, err := nPaymentLink.RestoreById(ctx, paymentLinkId)
	assertResponseSuccess(t, resp, err)
}

func TestPaymentLinkDeleteImageById(t *testing.T) {
	initPaymentLinkImage()
	accessToken := getEnvValue(EnvAccessToken)
	paymentLinkId := getEnvValue(EnvPaymentLinkId)
	paymentLinkImageId := getEnvValue(EnvPaymentLinkImageId)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nPaymentLink := NewPaymentLink(EnvSandbox, accessToken)
	resp, err := nPaymentLink.DeleteImageById(ctx, paymentLinkId, paymentLinkImageId)
	assertResponseSuccess(t, resp, err)
}

func TestPaymentLinkGetById(t *testing.T) {
	initPaymentLinkImage()
	accessToken := getEnvValue(EnvAccessToken)
	paymentLinkId := getEnvValue(EnvPaymentLinkId)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nPaymentLink := NewPaymentLink(EnvSandbox, accessToken)
	resp, err := nPaymentLink.GetById(ctx, paymentLinkId)
	assertResponseSuccess(t, resp, err)
}

func TestPaymentLinkGetImageById(t *testing.T) {
	initPaymentLinkImage()
	accessToken := getEnvValue(EnvAccessToken)
	paymentLinkId := getEnvValue(EnvPaymentLinkId)
	paymentLinkImageId := getEnvValue(EnvPaymentLinkImageId)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nPaymentLink := NewPaymentLink(EnvSandbox, accessToken)
	resp, err := nPaymentLink.GetImageById(ctx, paymentLinkId, paymentLinkImageId)
	assertResponseSuccess(t, resp, err)
}

func TestPaymentLinkGetAll(t *testing.T) {
	initPaymentLink()
	accessToken := getEnvValue(EnvAccessToken)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nPaymentLink := NewPaymentLink(EnvSandbox, accessToken)
	resp, err := nPaymentLink.GetAll(ctx, GetAllPaymentLinksRequest{
		Name:           "",
		Active:         nil,
		IncludeDeleted: nil,
		Offset:         0,
		Limit:          10,
	})
	assertResponseSuccess(t, resp, err)
}

func TestPaymentLinkGetImagesById(t *testing.T) {
	initPaymentLinkImage()
	accessToken := getEnvValue(EnvAccessToken)
	paymentLinkId := getEnvValue(EnvPaymentLinkId)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nPaymentLink := NewPaymentLink(EnvSandbox, accessToken)
	resp, err := nPaymentLink.GetImagesById(ctx, paymentLinkId)
	assertResponseSuccess(t, resp, err)
}
