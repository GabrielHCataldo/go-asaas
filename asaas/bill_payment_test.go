package asaas

import (
	"context"
	"testing"
	"time"
)

func TestBillPaymentSimulate(t *testing.T) {
	initBankSlipCharge(false)
	accessToken := getEnvValue(EnvAccessToken)
	identificationField := getEnvValue(EnvChargeIdentificationField)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nBillPayment := NewBillPayment(EnvSandbox, accessToken)
	resp, err := nBillPayment.Simulate(ctx, BillPaymentSimulateRequest{
		IdentificationField: identificationField,
		BarCode:             "",
	})
	assertResponseSuccess(t, resp, err)
}

func TestBillPaymentCreate(t *testing.T) {
	initBankSlipCharge(false)
	accessToken := getEnvValue(EnvAccessToken)
	identificationField := getEnvValue(EnvChargeIdentificationField)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nBillPayment := NewBillPayment(EnvSandbox, accessToken)
	resp, err := nBillPayment.Create(ctx, CreateBillPaymentRequest{
		IdentificationField: identificationField,
		ScheduleDate:        Date{},
		Description:         "unit test",
		Discount:            0,
		Interest:            0,
		Fine:                0,
		Value:               0,
		DueDate:             Date{},
	})
	assertResponseSuccess(t, resp, err)
}

func TestBillPaymentCancelById(t *testing.T) {
	initBillPayment()
	accessToken := getEnvValue(EnvAccessToken)
	billPaymentId := getEnvValue(EnvBillPaymentId)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nBillPayment := NewBillPayment(EnvSandbox, accessToken)
	resp, err := nBillPayment.CancelById(ctx, billPaymentId)
	assertResponseSuccess(t, resp, err)
}

func TestBillPaymentGetById(t *testing.T) {
	initBillPayment()
	accessToken := getEnvValue(EnvAccessToken)
	billPaymentId := getEnvValue(EnvBillPaymentId)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nBillPayment := NewBillPayment(EnvSandbox, accessToken)
	resp, err := nBillPayment.GetById(ctx, billPaymentId)
	assertResponseSuccess(t, resp, err)
}

func TestBillPaymentGetAll(t *testing.T) {
	initBillPayment()
	accessToken := getEnvValue(EnvAccessToken)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nBillPayment := NewBillPayment(EnvSandbox, accessToken)
	resp, err := nBillPayment.GetAll(ctx, PageableDefaultRequest{
		Offset: 0,
		Limit:  10,
	})
	assertResponseSuccess(t, resp, err)
}
