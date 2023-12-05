package asaas

import (
	"context"
	"testing"
	"time"
)

func TestInvoiceSchedule(t *testing.T) {
	initCreditCardCharge(false, false)
	accessToken := getEnvValue(EnvAccessToken)
	chargeId := getEnvValue(EnvCreditCardChargeId)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nInvoice := NewInvoice(EnvSandbox, accessToken)
	resp, err := nInvoice.Schedule(ctx, ScheduleInvoiceRequest{
		Payment:              chargeId,
		Installment:          "",
		Customer:             "",
		ServiceDescription:   "Unit test go",
		Observations:         "Unit test go",
		ExternalReference:    "",
		Value:                5,
		Deductions:           0,
		EffectiveDate:        Date{},
		MunicipalServiceId:   "",
		MunicipalServiceCode: "",
		MunicipalServiceName: "",
		UpdatePayment:        false,
		Taxes:                InvoiceTaxesRequest{},
	})
	assertResponseSuccess(t, resp, err)
}

func TestInvoiceAuthorizeById(t *testing.T) {
	initInvoice()
	accessToken := getEnvValue(EnvAccessToken)
	invoiceId := getEnvValue(EnvCreditCardChargeId)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nInvoice := NewInvoice(EnvSandbox, accessToken)
	resp, err := nInvoice.AuthorizeById(ctx, invoiceId)
	assertResponseSuccess(t, resp, err)
}

func TestInvoiceUpdateById(t *testing.T) {
	initInvoice()
	accessToken := getEnvValue(EnvAccessToken)
	invoiceId := getEnvValue(EnvCreditCardChargeId)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nInvoice := NewInvoice(EnvSandbox, accessToken)
	resp, err := nInvoice.UpdateById(ctx, invoiceId, UpdateInvoiceRequest{
		ServiceDescription:   "Unit test golang",
		Observations:         "",
		ExternalReference:    nil,
		Value:                0,
		Deductions:           nil,
		EffectiveDate:        Date{},
		MunicipalServiceId:   nil,
		MunicipalServiceCode: nil,
		MunicipalServiceName: nil,
		UpdatePayment:        nil,
		Taxes:                nil,
	})
	assertResponseSuccess(t, resp, err)
}

func TestInvoiceCancelById(t *testing.T) {
	initInvoice()
	accessToken := getEnvValue(EnvAccessToken)
	invoiceId := getEnvValue(EnvCreditCardChargeId)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nInvoice := NewInvoice(EnvSandbox, accessToken)
	resp, err := nInvoice.CancelById(ctx, invoiceId)
	assertResponseSuccess(t, resp, err)
}

func TestInvoiceGetById(t *testing.T) {
	initInvoice()
	accessToken := getEnvValue(EnvAccessToken)
	invoiceId := getEnvValue(EnvCreditCardChargeId)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nInvoice := NewInvoice(EnvSandbox, accessToken)
	resp, err := nInvoice.GetById(ctx, invoiceId)
	assertResponseSuccess(t, resp, err)
}

func TestInvoiceGetAll(t *testing.T) {
	initInvoice()
	accessToken := getEnvValue(EnvAccessToken)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nInvoice := NewInvoice(EnvSandbox, accessToken)
	resp, err := nInvoice.GetAll(ctx, GetAllInvoicesRequest{
		EffectiveDateGE:   Date{},
		EffectiveDateLE:   Date{},
		Payment:           "",
		Installment:       "",
		Customer:          "",
		ExternalReference: "",
		Status:            "",
		Offset:            0,
		Limit:             10,
	})
	assertResponseSuccess(t, resp, err)
}
