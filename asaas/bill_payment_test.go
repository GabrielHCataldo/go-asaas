package asaas

import (
	"context"
	"testing"
	"time"
)

func TestBillPaymentSimulateSuccess(t *testing.T) {
	accessToken, err := getAccessToken()
	assertFatalErrorNonnull(t, err)
	identificationField, err := getChargeIdentificationField()
	assertFatalErrorNonnull(t, err)
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
	_, errAsaas := nBillPayment.Create(context.TODO(), CreateBillPaymentRequest{})
	assertSuccessNonnull(t, errAsaas)
}

func TestBillPaymentCreateSuccess(t *testing.T) {
	accessToken, err := getAccessToken()
	assertFatalErrorNonnull(t, err)
	identificationField, err := getChargeIdentificationField()
	assertFatalErrorNonnull(t, err)
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
	_, errAsaas := nBillPayment.Create(context.TODO(), CreateBillPaymentRequest{})
	assertSuccessNonnull(t, errAsaas)
}

func TestBillPaymentCancelByIdSuccess(t *testing.T) {
	accessToken, err := getAccessToken()
	assertFatalErrorNonnull(t, err)
	initBillPayment()
	billPaymentId, err := getBillPaymentId()
	assertFatalErrorNonnull(t, err)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nBillPayment := NewBillPayment(EnvSandbox, accessToken)
	resp, errAsaas := nBillPayment.CancelById(ctx, billPaymentId)
	assertResponseSuccess(t, resp, errAsaas)
}

func TestBillPaymentCancelByIdError(t *testing.T) {
	nBillPayment := NewBillPayment(EnvSandbox, "")
	resp, errAsaas := nBillPayment.CancelById(context.TODO(), "test")
	assertResponseFailure(t, resp, errAsaas)
}

func TestBillPaymentGetByIdSuccess(t *testing.T) {
	accessToken, err := getAccessToken()
	assertFatalErrorNonnull(t, err)
	initBillPayment()
	billPaymentId, err := getBillPaymentId()
	assertFatalErrorNonnull(t, err)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nBillPayment := NewBillPayment(EnvSandbox, accessToken)
	resp, errAsaas := nBillPayment.GetById(ctx, billPaymentId)
	assertResponseSuccess(t, resp, errAsaas)
}

func TestBillPaymentGetIdError(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nBillPayment := NewBillPayment(EnvSandbox, "")
	resp, errAsaas := nBillPayment.GetById(ctx, "test")
	assertResponseFailure(t, resp, errAsaas)
}

func TestBillPaymentGetAllSuccess(t *testing.T) {
	accessToken, err := getAccessToken()
	assertFatalErrorNonnull(t, err)
	initBillPayment()
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nBillPayment := NewBillPayment(EnvSandbox, accessToken)
	resp, errAsaas := nBillPayment.GetAll(ctx, PageableDefaultRequest{})
	assertResponseSuccess(t, resp, errAsaas)
}

func TestBillPaymentGetAllError(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nBillPayment := NewBillPayment(EnvSandbox, "")
	resp, errAsaas := nBillPayment.GetAll(ctx, PageableDefaultRequest{})
	assertResponseFailure(t, resp, errAsaas)
}
