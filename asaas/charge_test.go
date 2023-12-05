package asaas

import (
	"context"
	"os"
	"testing"
	"time"
)

func TestChargeCreate(t *testing.T) {
	initCustomer(false)
	accessToken := getEnvValue(EnvAccessToken)
	customerId := getEnvValue(EnvCustomerId)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
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

func TestChargePayWithCreditCard(t *testing.T) {
	initUndefinedCharge()
	accessToken := getEnvValue(EnvAccessToken)
	chargeId := getEnvValue(EnvUndefinedChargeId)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nCharge := NewCharge(EnvSandbox, accessToken)
	resp, err := nCharge.PayWithCreditCard(ctx, chargeId, PayWithCreditCardRequest{
		CreditCard: &CreditCardRequest{
			HolderName:  "unit test go",
			Number:      "4000000000000010",
			ExpiryMonth: "05",
			ExpiryYear:  "2035",
			Ccv:         "318",
		},
		CreditCardHolderInfo: &CreditCardHolderInfoRequest{
			Name:          "Unit Test Go",
			CpfCnpj:       "24971563792",
			Email:         "unittest@gmail.com",
			Phone:         "4738010919",
			MobilePhone:   "47998781877",
			PostalCode:    "89223-005",
			AddressNumber: "277",
		},
		CreditCardToken: "",
	})
	assertResponseSuccess(t, resp, err)
}

func TestChargeUpdateById(t *testing.T) {
	initUndefinedCharge()
	accessToken := getEnvValue(EnvAccessToken)
	chargeId := getEnvValue(EnvUndefinedChargeId)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
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

func TestChargeDeleteById(t *testing.T) {
	initUndefinedCharge()
	accessToken := getEnvValue(EnvAccessToken)
	chargeId := getEnvValue(EnvUndefinedChargeId)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nCharge := NewCharge(EnvSandbox, accessToken)
	resp, err := nCharge.DeleteById(ctx, chargeId)
	assertResponseSuccess(t, resp, err)
}

func TestChargeRestoreById(t *testing.T) {
	initChargeDeleted()
	accessToken := getEnvValue(EnvAccessToken)
	chargeId := getEnvValue(EnvChargeDeletedId)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nCharge := NewCharge(EnvSandbox, accessToken)
	resp, err := nCharge.RestoreById(ctx, chargeId)
	assertResponseSuccess(t, resp, err)
}

func TestChargeRefundById(t *testing.T) {
	initCreditCardCharge(false, false)
	accessToken := getEnvValue(EnvAccessToken)
	chargeId := getEnvValue(EnvCreditCardChargeId)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nCharge := NewCharge(EnvSandbox, accessToken)
	resp, err := nCharge.RefundById(ctx, chargeId, RefundRequest{
		Value:       0,
		Description: "",
	})
	assertResponseSuccess(t, resp, err)
}

func TestChargeReceiveInCashById(t *testing.T) {
	initUndefinedCharge()
	accessToken := getEnvValue(EnvAccessToken)
	chargeId := getEnvValue(EnvUndefinedChargeId)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nCharge := NewCharge(EnvSandbox, accessToken)
	now := DateNow()
	resp, err := nCharge.ReceiveInCashById(ctx, chargeId, ChargeReceiveInCashRequest{
		PaymentDate:    NewDate(now.Year(), now.Month(), now.Day(), now.Location()),
		Value:          5,
		NotifyCustomer: false,
	})
	assertResponseSuccess(t, resp, err)
}

func TestChargeUndoReceivedInCashById(t *testing.T) {
	initChargeReceivedInCash()
	accessToken := getEnvValue(EnvAccessToken)
	chargeId := getEnvValue(EnvChargeReceivedInCashId)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nCharge := NewCharge(EnvSandbox, accessToken)
	resp, err := nCharge.UndoReceivedInCashById(ctx, chargeId)
	assertResponseSuccess(t, resp, err)
}

func TestChargeUploadDocumentById(t *testing.T) {
	initUndefinedCharge()
	accessToken := getEnvValue(EnvAccessToken)
	chargeId := getEnvValue(EnvUndefinedChargeId)
	f, _ := os.Open(getEnvValue(EnvFileName))
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
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
	initChargeDocumentId()
	accessToken := getEnvValue(EnvAccessToken)
	chargeId := getEnvValue(EnvUndefinedChargeId)
	docId := getEnvValue(EnvChargeDocumentId)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nCharge := NewCharge(EnvSandbox, accessToken)
	resp, err := nCharge.UpdateDocumentDefinitionsById(ctx, chargeId, docId, UpdateChargeDocumentDefinitionsRequest{
		AvailableAfterPayment: false,
		Type:                  DocumentTypeContract,
	})
	assertResponseSuccess(t, resp, err)
}

func TestChargeDeleteDocumentById(t *testing.T) {
	initChargeDocumentId()
	accessToken := getEnvValue(EnvAccessToken)
	chargeId := getEnvValue(EnvUndefinedChargeId)
	docId := getEnvValue(EnvChargeDocumentId)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nCharge := NewCharge(EnvSandbox, accessToken)
	resp, err := nCharge.DeleteDocumentById(ctx, chargeId, docId)
	assertResponseSuccess(t, resp, err)
}

func TestChargeGetById(t *testing.T) {
	initUndefinedCharge()
	accessToken := getEnvValue(EnvAccessToken)
	chargeId := getEnvValue(EnvUndefinedChargeId)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nCharge := NewCharge(EnvSandbox, accessToken)
	resp, err := nCharge.GetById(ctx, chargeId)
	assertResponseSuccess(t, resp, err)
}

func TestChargeGetCreationLimit(t *testing.T) {
	accessToken := getEnvValue(EnvAccessToken)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nCharge := NewCharge(EnvSandbox, accessToken)
	resp, err := nCharge.GetCreationLimit(ctx)
	assertResponseSuccess(t, resp, err)
}

func TestChargeGetStatusById(t *testing.T) {
	initUndefinedCharge()
	accessToken := getEnvValue(EnvAccessToken)
	chargeId := getEnvValue(EnvUndefinedChargeId)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nCharge := NewCharge(EnvSandbox, accessToken)
	resp, err := nCharge.GetStatusById(ctx, chargeId)
	assertResponseSuccess(t, resp, err)
}

func TestChargeGetIdentificationFieldById(t *testing.T) {
	initBankSlipCharge(false)
	accessToken := getEnvValue(EnvAccessToken)
	chargeId := getEnvValue(EnvBankSlipChargeId)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nCharge := NewCharge(EnvSandbox, accessToken)
	resp, err := nCharge.GetIdentificationFieldById(ctx, chargeId)
	assertResponseSuccess(t, resp, err)
}

func TestChargeGetPixQrCodeById(t *testing.T) {
	initPixCharge()
	accessToken := getEnvValue(EnvAccessToken)
	chargeId := getEnvValue(EnvPixChargeId)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nCharge := NewCharge(EnvSandbox, accessToken)
	resp, err := nCharge.GetPixQrCodeById(ctx, chargeId)
	assertResponseSuccess(t, resp, err)
}

func TestChargeGetDocumentById(t *testing.T) {
	initChargeDocumentId()
	accessToken := getEnvValue(EnvAccessToken)
	chargeId := getEnvValue(EnvUndefinedChargeId)
	docId := getEnvValue(EnvChargeDocumentId)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nCharge := NewCharge(EnvSandbox, accessToken)
	resp, err := nCharge.GetDocumentById(ctx, chargeId, docId)
	assertResponseSuccess(t, resp, err)
}

func TestChargeGetAll(t *testing.T) {
	initChargeDocumentId()
	accessToken := getEnvValue(EnvAccessToken)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
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
	initChargeDocumentId()
	accessToken := getEnvValue(EnvAccessToken)
	chargeId := getEnvValue(EnvUndefinedChargeId)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nCharge := NewCharge(EnvSandbox, accessToken)
	resp, err := nCharge.GetAllDocumentsById(ctx, chargeId, PageableDefaultRequest{
		Offset: 0,
		Limit:  10,
	})
	assertResponseSuccess(t, resp, err)
}
