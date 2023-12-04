package asaas

import (
	"context"
	"os"
	"testing"
	"time"
)

func TestChangeCreate(t *testing.T) {
	accessToken := getEnvValue(EnvAccessToken)
	assertFatalStringBlank(t, accessToken)
	initCustomer()
	customerId := getEnvValue(EnvCustomerId)
	assertFatalStringBlank(t, customerId)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nCharge := NewCharge(EnvSandbox, accessToken)
	resp, err := nCharge.Create(ctx, CreateChargeRequest{
		Customer:             customerId,
		BillingType:          BillingTypeUndefined,
		Value:                10,
		DueDate:              NewDate(2026, 12, 1, time.Local),
		Description:          "Unit test golang",
		ExternalReference:    "",
		Discount:             nil,
		Interest:             nil,
		Fine:                 nil,
		PostalService:        false,
		Split:                nil,
		Callback:             nil,
		CreditCard:           nil,
		CreditCardHolderInfo: nil,
		CreditCardToken:      "",
		InstallmentCount:     0,
		InstallmentValue:     0,
		AuthorizeOnly:        false,
		RemoteIp:             "",
	})
	assertResponseSuccess(t, resp, err)
}

func TestChangePayWithCreditCard(t *testing.T) {
	accessToken := getEnvValue(EnvAccessToken)
	assertFatalStringBlank(t, accessToken)
	initUndefinedCharge()
	chargeId := getEnvValue(EnvUndefinedChargeId)
	assertFatalStringBlank(t, chargeId)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nCharge := NewCharge(EnvSandbox, accessToken)
	resp, err := nCharge.PayWithCreditCard(ctx, chargeId, CreditCardRequest{
		HolderName:  "unit test go",
		Number:      "4000000000000010",
		ExpiryMonth: "12",
		ExpiryYear:  "2036",
		Ccv:         "123",
	})
	assertResponseSuccess(t, resp, err)
}

func TestChangeUpdateById(t *testing.T) {
	accessToken := getEnvValue(EnvAccessToken)
	assertFatalStringBlank(t, accessToken)
	initUndefinedCharge()
	chargeId := getEnvValue(EnvUndefinedChargeId)
	assertFatalStringBlank(t, chargeId)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nCharge := NewCharge(EnvSandbox, accessToken)
	resp, err := nCharge.UpdateById(ctx, chargeId, UpdateChargeRequest{
		Value:            15,
		DueDate:          Date{},
		Description:      Pointer("update from unit test golang"),
		InstallmentCount: 1,
		InstallmentValue: 7.5,
	})
	assertResponseSuccess(t, resp, err)
}

func TestChangeDeleteById(t *testing.T) {
	accessToken := getEnvValue(EnvAccessToken)
	assertFatalStringBlank(t, accessToken)
	initUndefinedCharge()
	chargeId := getEnvValue(EnvUndefinedChargeId)
	assertFatalStringBlank(t, chargeId)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nCharge := NewCharge(EnvSandbox, accessToken)
	resp, err := nCharge.DeleteById(ctx, chargeId)
	assertResponseSuccess(t, resp, err)
}

func TestChangeRestoreById(t *testing.T) {
	accessToken := getEnvValue(EnvAccessToken)
	assertFatalStringBlank(t, accessToken)
	initChargeDeleted()
	chargeId := getEnvValue(EnvChargeDeletedId)
	assertFatalStringBlank(t, chargeId)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nCharge := NewCharge(EnvSandbox, accessToken)
	resp, err := nCharge.RestoreById(ctx, chargeId)
	assertResponseSuccess(t, resp, err)
}

func TestChangeRefundById(t *testing.T) {
	accessToken := getEnvValue(EnvAccessToken)
	assertFatalStringBlank(t, accessToken)
	initCreditCardCharge(true, false)
	chargeId := getEnvValue(EnvCreditCardChargeId)
	assertFatalStringBlank(t, chargeId)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nCharge := NewCharge(EnvSandbox, accessToken)
	resp, err := nCharge.RefundById(ctx, chargeId, RefundRequest{
		Value:       0,
		Description: "",
	})
	assertResponseSuccess(t, resp, err)
}

func TestChangeReceiveInCashById(t *testing.T) {
	accessToken := getEnvValue(EnvAccessToken)
	assertFatalStringBlank(t, accessToken)
	initUndefinedCharge()
	chargeId := getEnvValue(EnvUndefinedChargeId)
	assertFatalStringBlank(t, chargeId)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nCharge := NewCharge(EnvSandbox, accessToken)
	now := time.Now()
	resp, err := nCharge.ReceiveInCashById(ctx, chargeId, ChargeReceiveInCashRequest{
		PaymentDate:    NewDate(now.Year(), now.Month(), now.Day(), now.Location()),
		Value:          100,
		NotifyCustomer: false,
	})
	assertResponseSuccess(t, resp, err)
}

func TestChangeUndoReceivedInCashById(t *testing.T) {
	accessToken := getEnvValue(EnvAccessToken)
	assertFatalStringBlank(t, accessToken)
	initChargeReceivedInCash()
	chargeId := getEnvValue(EnvChargeReceivedInCashId)
	assertFatalStringBlank(t, chargeId)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nCharge := NewCharge(EnvSandbox, accessToken)
	resp, err := nCharge.UndoReceivedInCashById(ctx, chargeId)
	assertResponseSuccess(t, resp, err)
}

func TestChargeUploadDocumentById(t *testing.T) {
	accessToken := getEnvValue(EnvAccessToken)
	assertFatalStringBlank(t, accessToken)
	initUndefinedCharge()
	chargeId := getEnvValue(EnvUndefinedChargeId)
	assertFatalStringBlank(t, chargeId)
	f, err := os.Open(getEnvValue(EnvFileName))
	assertFatalErrorNonnull(t, err)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nCharge := NewCharge(EnvSandbox, accessToken)
	resp, err := nCharge.UploadDocumentById(ctx, chargeId, UploadChargeDocumentRequest{
		AvailableAfterPayment: false,
		Type:                  DocumentTypeDocument,
		File:                  f,
	})
	assertResponseSuccess(t, resp, err)
}

func TestChargeUpdateDocumentDefinitionsById(t *testing.T) {
	accessToken := getEnvValue(EnvAccessToken)
	assertFatalStringBlank(t, accessToken)
	initChargeDocumentId()
	chargeId := getEnvValue(EnvUndefinedChargeId)
	assertFatalStringBlank(t, chargeId)
	docId := getEnvValue(EnvChargeDocumentId)
	assertFatalStringBlank(t, docId)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nCharge := NewCharge(EnvSandbox, accessToken)
	resp, err := nCharge.UpdateDocumentDefinitionsById(ctx, chargeId, docId, UpdateChargeDocumentDefinitionsRequest{
		AvailableAfterPayment: false,
		Type:                  DocumentTypeContract,
	})
	assertResponseSuccess(t, resp, err)
}

func TestChargeDeleteDocumentById(t *testing.T) {
	accessToken := getEnvValue(EnvAccessToken)
	assertFatalStringBlank(t, accessToken)
	initChargeDocumentId()
	chargeId := getEnvValue(EnvUndefinedChargeId)
	assertFatalStringBlank(t, chargeId)
	docId := getEnvValue(EnvChargeDocumentId)
	assertFatalStringBlank(t, docId)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nCharge := NewCharge(EnvSandbox, accessToken)
	resp, err := nCharge.DeleteDocumentById(ctx, chargeId, docId)
	assertResponseSuccess(t, resp, err)
}

func TestChargeGetById(t *testing.T) {
	accessToken := getEnvValue(EnvAccessToken)
	assertFatalStringBlank(t, accessToken)
	initUndefinedCharge()
	chargeId := getEnvValue(EnvUndefinedChargeId)
	assertFatalStringBlank(t, chargeId)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nCharge := NewCharge(EnvSandbox, accessToken)
	resp, err := nCharge.GetById(ctx, chargeId)
	assertResponseSuccess(t, resp, err)
}

func TestChargeGetCreationLimit(t *testing.T) {
	accessToken := getEnvValue(EnvAccessToken)
	assertFatalStringBlank(t, accessToken)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nCharge := NewCharge(EnvSandbox, accessToken)
	resp, err := nCharge.GetCreationLimit(ctx)
	assertResponseSuccess(t, resp, err)
}

func TestChargeGetStatusById(t *testing.T) {
	accessToken := getEnvValue(EnvAccessToken)
	assertFatalStringBlank(t, accessToken)
	initUndefinedCharge()
	chargeId := getEnvValue(EnvUndefinedChargeId)
	assertFatalStringBlank(t, chargeId)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nCharge := NewCharge(EnvSandbox, accessToken)
	resp, err := nCharge.GetStatusById(ctx, chargeId)
	assertResponseSuccess(t, resp, err)
}

func TestChargeGetIdentificationFieldById(t *testing.T) {
	accessToken := getEnvValue(EnvAccessToken)
	assertFatalStringBlank(t, accessToken)
	initBankSlipCharge(false)
	chargeId := getEnvValue(EnvBankSlipChargeId)
	assertFatalStringBlank(t, chargeId)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nCharge := NewCharge(EnvSandbox, accessToken)
	resp, err := nCharge.GetIdentificationFieldById(ctx, chargeId)
	assertResponseSuccess(t, resp, err)
}

func TestChargeGetPixQrCodeById(t *testing.T) {
	accessToken := getEnvValue(EnvAccessToken)
	assertFatalStringBlank(t, accessToken)
	initPixCharge()
	chargeId := getEnvValue(EnvPixChargeId)
	assertFatalStringBlank(t, chargeId)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nCharge := NewCharge(EnvSandbox, accessToken)
	resp, err := nCharge.GetPixQrCodeById(ctx, chargeId)
	assertResponseSuccess(t, resp, err)
}

func TestChargeGetDocumentById(t *testing.T) {
	accessToken := getEnvValue(EnvAccessToken)
	assertFatalStringBlank(t, accessToken)
	initChargeDocumentId()
	chargeId := getEnvValue(EnvUndefinedChargeId)
	assertFatalStringBlank(t, chargeId)
	docId := getEnvValue(EnvChargeDocumentId)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nCharge := NewCharge(EnvSandbox, accessToken)
	resp, err := nCharge.GetDocumentById(ctx, chargeId, docId)
	assertResponseSuccess(t, resp, err)
}

func TestChargeGetAll(t *testing.T) {
	accessToken := getEnvValue(EnvAccessToken)
	assertFatalStringBlank(t, accessToken)
	initChargeDocumentId()
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nCharge := NewCharge(EnvSandbox, accessToken)
	resp, err := nCharge.GetAll(ctx, GetAllChargesRequest{
		Customer:              "",
		Subscription:          "",
		Installment:           "",
		CustomerGroupName:     "",
		BillingType:           "",
		Status:                "",
		ExternalReference:     "",
		InvoiceStatus:         "",
		EstimatedCreditDate:   Date{},
		PixQrCodeId:           "",
		Anticipated:           nil,
		PaymentDate:           Date{},
		DateCreatedGe:         Date{},
		DateCreatedLe:         Date{},
		PaymentDateGe:         Date{},
		PaymentDateLe:         Date{},
		EstimatedCreditDateGE: Date{},
		EstimatedCreditDateLE: Date{},
		DueDateGe:             Date{},
		DueDateLe:             Date{},
		User:                  "",
		Offset:                0,
		Limit:                 10,
	})
	assertResponseSuccess(t, resp, err)
}

func TestChargeGetAllDocumentsById(t *testing.T) {
	accessToken := getEnvValue(EnvAccessToken)
	assertFatalStringBlank(t, accessToken)
	initChargeDocumentId()
	chargeId := getEnvValue(EnvUndefinedChargeId)
	assertFatalStringBlank(t, chargeId)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nCharge := NewCharge(EnvSandbox, accessToken)
	resp, err := nCharge.GetAllDocumentsById(ctx, chargeId, PageableDefaultRequest{
		Offset: 0,
		Limit:  10,
	})
	assertResponseSuccess(t, resp, err)
}
