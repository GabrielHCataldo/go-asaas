package asaas

import (
	"context"
	"os"
	"testing"
	"time"
)

func TestChangeCreateSuccess(t *testing.T) {
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

func TestChargeCreateError(t *testing.T) {
	nCharge := NewCharge(EnvSandbox, "")
	resp, err := nCharge.Create(context.TODO(), CreateChargeRequest{})
	assertResponseFailure(t, resp, err)
}

func TestChangePayWithCreditCardSuccess(t *testing.T) {
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

func TestChargePayWithCreditCardError(t *testing.T) {
	nCharge := NewCharge(EnvSandbox, "")
	resp, err := nCharge.PayWithCreditCard(context.TODO(), "", CreditCardRequest{})
	assertResponseFailure(t, resp, err)
}

func TestChangeUpdateByIdSuccess(t *testing.T) {
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

func TestChargeUpdateByIdError(t *testing.T) {
	nCharge := NewCharge(EnvSandbox, "")
	resp, err := nCharge.UpdateById(context.TODO(), "", UpdateChargeRequest{})
	assertResponseFailure(t, resp, err)
}

func TestChangeDeleteByIdSuccess(t *testing.T) {
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

func TestChargeDeleteByIdError(t *testing.T) {
	nCharge := NewCharge(EnvSandbox, "")
	resp, err := nCharge.DeleteById(context.TODO(), "")
	assertResponseFailure(t, resp, err)
}

func TestChangeRestoreByIdSuccess(t *testing.T) {
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

func TestChargeRestoreByIdError(t *testing.T) {
	nCharge := NewCharge(EnvSandbox, "")
	resp, err := nCharge.RestoreById(context.TODO(), "")
	assertResponseFailure(t, resp, err)
}

func TestChangeRefundByIdSuccess(t *testing.T) {
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

func TestChargeRefundByIdError(t *testing.T) {
	nCharge := NewCharge(EnvSandbox, "")
	resp, err := nCharge.RefundById(context.TODO(), "", RefundRequest{})
	assertResponseFailure(t, resp, err)
}

func TestChangeReceiveInCashByIdSuccess(t *testing.T) {
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

func TestChargeReceiveInCashByIdError(t *testing.T) {
	nCharge := NewCharge(EnvSandbox, "")
	resp, err := nCharge.ReceiveInCashById(context.TODO(), "", ChargeReceiveInCashRequest{})
	assertResponseFailure(t, resp, err)
}

func TestChangeUndoReceivedInCashByIdSuccess(t *testing.T) {
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

func TestChargeUndoReceivedInCashByIdError(t *testing.T) {
	nCharge := NewCharge(EnvSandbox, "")
	resp, err := nCharge.UndoReceivedInCashById(context.TODO(), "")
	assertResponseFailure(t, resp, err)
}

func TestChargeUploadDocumentByIdSuccess(t *testing.T) {
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

func TestChargeUploadDocumentByIdError(t *testing.T) {
	nCharge := NewCharge(EnvSandbox, "")
	resp, err := nCharge.UploadDocumentById(context.TODO(), "", UploadChargeDocumentRequest{})
	assertResponseFailure(t, resp, err)
}

func TestChargeUpdateDocumentDefinitionsByIdSuccess(t *testing.T) {
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

func TestChargeUpdateDocumentDefinitionsByIdError(t *testing.T) {
	nCharge := NewCharge(EnvSandbox, "")
	ctx := context.TODO()
	resp, err := nCharge.UpdateDocumentDefinitionsById(ctx, "", "", UpdateChargeDocumentDefinitionsRequest{})
	assertResponseFailure(t, resp, err)
}

func TestChargeDeleteDocumentByIdSuccess(t *testing.T) {
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

func TestChargeDeleteDocumentByIdError(t *testing.T) {
	nCharge := NewCharge(EnvSandbox, "")
	ctx := context.TODO()
	resp, err := nCharge.DeleteDocumentById(ctx, "", "")
	assertResponseFailure(t, resp, err)
}

func TestChargeGetByIdSuccess(t *testing.T) {
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

func TestChargeGetByIdError(t *testing.T) {
	nCharge := NewCharge(EnvSandbox, "")
	resp, err := nCharge.GetDocumentById(context.TODO(), "", "")
	assertResponseFailure(t, resp, err)
}

func TestChargeGetCreationLimitSuccess(t *testing.T) {
	accessToken := getEnvValue(EnvAccessToken)
	assertFatalStringBlank(t, accessToken)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nCharge := NewCharge(EnvSandbox, accessToken)
	resp, err := nCharge.GetCreationLimit(ctx)
	assertResponseSuccess(t, resp, err)
}

func TestChargeGetCreationLimitError(t *testing.T) {
	nCharge := NewCharge(EnvSandbox, "")
	resp, err := nCharge.GetCreationLimit(context.TODO())
	assertResponseSuccess(t, resp, err)
}

func TestChargeGetStatusByIdSuccess(t *testing.T) {
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

func TestChargeGetStatusByIdError(t *testing.T) {
	nCharge := NewCharge(EnvSandbox, "")
	resp, err := nCharge.GetStatusById(context.TODO(), "")
	assertResponseFailure(t, resp, err)
}

func TestChargeGetIdentificationFieldByIdSuccess(t *testing.T) {
	accessToken := getEnvValue(EnvAccessToken)
	assertFatalStringBlank(t, accessToken)
	initBankSlipCharge()
	chargeId := getEnvValue(EnvBankSlipChargeId)
	assertFatalStringBlank(t, chargeId)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nCharge := NewCharge(EnvSandbox, accessToken)
	resp, err := nCharge.GetIdentificationFieldById(ctx, chargeId)
	assertResponseSuccess(t, resp, err)
}

func TestChargeGetIdentificationFieldByIdError(t *testing.T) {
	nCharge := NewCharge(EnvSandbox, "")
	resp, err := nCharge.GetIdentificationFieldById(context.TODO(), "")
	assertResponseFailure(t, resp, err)
}

func TestChargeGetPixQrCodeByIdSuccess(t *testing.T) {
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

func TestChargeGetPixQrCodeByIdError(t *testing.T) {
	nCharge := NewCharge(EnvSandbox, "")
	resp, err := nCharge.GetPixQrCodeById(context.TODO(), "")
	assertResponseFailure(t, resp, err)
}

func TestChargeGetDocumentByIdSuccess(t *testing.T) {
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

func TestChargeGetDocumentByIdError(t *testing.T) {
	nCharge := NewCharge(EnvSandbox, "")
	resp, err := nCharge.GetDocumentById(context.TODO(), "", "")
	assertResponseFailure(t, resp, err)
}

func TestChargeGetAllSuccess(t *testing.T) {
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

func TestChargeGetAllError(t *testing.T) {
	nCharge := NewCharge(EnvSandbox, "")
	resp, err := nCharge.GetAll(context.TODO(), GetAllChargesRequest{})
	assertResponseFailure(t, resp, err)
}

func TestChargeGetAllDocumentsByIdSuccess(t *testing.T) {
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

func TestChargeGetAllDocumentsByIdError(t *testing.T) {
	nCharge := NewCharge(EnvSandbox, "")
	resp, err := nCharge.GetAllDocumentsById(context.TODO(), "", PageableDefaultRequest{})
	assertResponseFailure(t, resp, err)
}
