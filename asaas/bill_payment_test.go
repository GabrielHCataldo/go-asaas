package asaas

import (
	"context"
	"testing"
	"time"
)

func TestBillPaymentSimulateSuccess(t *testing.T) {
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

func TestBillPaymentSimulateError(t *testing.T) {
	nBillPayment := NewBillPayment(EnvSandbox, "")
	resp, err := nBillPayment.Create(context.TODO(), CreateBillPaymentRequest{})
	assertResponseFailure(t, resp, err)
}

func TestBillPaymentCreateSuccess(t *testing.T) {
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

func TestBillPaymentCreateError(t *testing.T) {
	nBillPayment := NewBillPayment(EnvSandbox, "")
	resp, err := nBillPayment.Create(context.TODO(), CreateBillPaymentRequest{})
	assertResponseFailure(t, resp, err)
}

func TestBillPaymentCancelByIdSuccess(t *testing.T) {
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

func TestBillPaymentCancelByIdError(t *testing.T) {
	nBillPayment := NewBillPayment(EnvSandbox, "")
	resp, errAsaas := nBillPayment.CancelById(context.TODO(), "")
	assertResponseFailure(t, resp, errAsaas)
}

func TestBillPaymentGetByIdSuccess(t *testing.T) {
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

func TestBillPaymentGetIdError(t *testing.T) {
	nBillPayment := NewBillPayment(EnvSandbox, "")
	resp, errAsaas := nBillPayment.GetById(context.TODO(), "")
	assertResponseFailure(t, resp, errAsaas)
}

func TestBillPaymentGetAllSuccess(t *testing.T) {
	accessToken := getEnvValue(EnvAccessToken)
	assertFatalStringBlank(t, accessToken)
	initBillPayment()
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nBillPayment := NewBillPayment(EnvSandbox, accessToken)
	resp, errAsaas := nBillPayment.GetAll(ctx, PageableDefaultRequest{})
	assertResponseSuccess(t, resp, errAsaas)
}

func TestBillPaymentGetAllError(t *testing.T) {
	nBillPayment := NewBillPayment(EnvSandbox, "")
	resp, errAsaas := nBillPayment.GetAll(context.TODO(), PageableDefaultRequest{})
	assertResponseFailure(t, resp, errAsaas)
}
