package asaas

import (
	"context"
	"testing"
	"time"
)

func TestBillPaymentSimulate(t *testing.T) {
	accessToken := getEnvValue(EnvAccessToken)
	assertFatalStringBlank(t, accessToken)
	identificationField := getEnvValue(EnvChargeIdentificationField)
	assertFatalStringBlank(t, identificationField)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nBillPayment := NewBillPayment(EnvSandbox, accessToken)
	resp, errAsaas := nBillPayment.Simulate(ctx, BillPaymentSimulateRequest{
		IdentificationField: identificationField,
		BarCode:             "",
	})
	assertResponseSuccess(t, resp, errAsaas)
}

func TestBillPaymentCreate(t *testing.T) {
	accessToken := getEnvValue(EnvAccessToken)
	assertFatalStringBlank(t, accessToken)
	identificationField := getEnvValue(EnvChargeIdentificationField)
	assertFatalStringBlank(t, identificationField)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nBillPayment := NewBillPayment(EnvSandbox, accessToken)
	resp, errAsaas := nBillPayment.Create(ctx, CreateBillPaymentRequest{
		IdentificationField: identificationField,
		ScheduleDate:        Date{},
		Description:         "unit test",
		Discount:            0,
		Interest:            0,
		Fine:                0,
		Value:               0,
		DueDate:             Date{},
	})
	assertResponseSuccess(t, resp, errAsaas)
}

func TestBillPaymentCancelById(t *testing.T) {
	accessToken := getEnvValue(EnvAccessToken)
	assertFatalStringBlank(t, accessToken)
	initBillPayment()
	billPaymentId := getEnvValue(EnvBillPaymentId)
	assertFatalStringBlank(t, billPaymentId)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nBillPayment := NewBillPayment(EnvSandbox, accessToken)
	resp, errAsaas := nBillPayment.CancelById(ctx, billPaymentId)
	assertResponseSuccess(t, resp, errAsaas)
}

func TestBillPaymentGetById(t *testing.T) {
	accessToken := getEnvValue(EnvAccessToken)
	assertFatalStringBlank(t, accessToken)
	initBillPayment()
	billPaymentId := getEnvValue(EnvBillPaymentId)
	assertFatalStringBlank(t, billPaymentId)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nBillPayment := NewBillPayment(EnvSandbox, accessToken)
	resp, errAsaas := nBillPayment.GetById(ctx, billPaymentId)
	assertResponseSuccess(t, resp, errAsaas)
}

func TestBillPaymentGetAll(t *testing.T) {
	accessToken := getEnvValue(EnvAccessToken)
	assertFatalStringBlank(t, accessToken)
	initBillPayment()
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nBillPayment := NewBillPayment(EnvSandbox, accessToken)
	resp, errAsaas := nBillPayment.GetAll(ctx, PageableDefaultRequest{})
	assertResponseSuccess(t, resp, errAsaas)
}
