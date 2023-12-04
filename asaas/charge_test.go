package asaas

import (
	"context"
	"encoding/json"
	"github.com/GabrielHCataldo/go-asaas/internal/test"
	"os"
	"testing"
	"time"
)

func TestChangeCreateSuccess(t *testing.T) {
	accessToken, err := getAccessToken()
	assertFatalErrorNonnull(t, err)
	customerId, err := getCustomerId()
	assertFatalErrorNonnull(t, err)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	req := &CreateChargeRequest{
		Customer:             "",
		BillingType:          "",
		Value:                0,
		DueDate:              Date{},
		Description:          "",
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
	}
	req.Customer = customerId
	err = json.Unmarshal(test.GetCreatePixChargeRequestDefault(), req)
	assertFatalErrorNonnull(t, err)
	nCharge := NewCharge(EnvSandbox, accessToken)
	resp, errAsaas := nCharge.Create(ctx, *req)
	assertResponseSuccess(t, resp, errAsaas)
}

func TestChargeCreateError(t *testing.T) {
	nCharge := NewCharge(EnvSandbox, "")
	_, errAsaas := nCharge.Create(context.TODO(), CreateChargeRequest{})
	assertSuccessNonnull(t, errAsaas)
}

func TestChangePayWithCreditCardSuccess(t *testing.T) {
	accessToken, err := getAccessToken()
	assertFatalErrorNonnull(t, err)
	initUndefinedCharge()
	chargeId, err := getUndefinedChargeId()
	assertFatalErrorNonnull(t, err)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	req := &CreditCardRequest{
		HolderName:  "",
		Number:      "",
		ExpiryMonth: "",
		ExpiryYear:  "",
		Ccv:         "",
	}
	err = json.Unmarshal(test.GetCreditCardRequestDefault(), req)
	assertFatalErrorNonnull(t, err)
	nCharge := NewCharge(EnvSandbox, accessToken)
	resp, errAsaas := nCharge.PayWithCreditCard(ctx, chargeId, *req)
	assertResponseSuccess(t, resp, errAsaas)
}

func TestChargePayWithCreditCardError(t *testing.T) {
	nCharge := NewCharge(EnvSandbox, "")
	_, errAsaas := nCharge.PayWithCreditCard(context.TODO(), "test", CreditCardRequest{})
	assertSuccessNonnull(t, errAsaas)
}

func TestChangeUpdateByIdSuccess(t *testing.T) {
	accessToken, err := getAccessToken()
	assertFatalErrorNonnull(t, err)
	customerId, err := getCustomerId()
	assertFatalErrorNonnull(t, err)
	initUndefinedCharge()
	chargeId, err := getUndefinedChargeId()
	assertFatalErrorNonnull(t, err)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	req := &UpdateChargeRequest{
		Customer:          "",
		BillingType:       "",
		Value:             0,
		DueDate:           Date{},
		Description:       nil,
		ExternalReference: nil,
		Discount:          nil,
		Interest:          nil,
		Fine:              nil,
		PostalService:     nil,
		Split:             nil,
		Callback:          nil,
		InstallmentCount:  0,
		InstallmentValue:  0,
	}
	err = json.Unmarshal(test.GetCreateCreditCardChargeRequestDefault(), req)
	assertFatalErrorNonnull(t, err)
	req.Customer = customerId
	nCharge := NewCharge(EnvSandbox, accessToken)
	resp, errAsaas := nCharge.UpdateById(ctx, chargeId, *req)
	assertResponseSuccess(t, resp, errAsaas)
}

func TestChargeUpdateByIdError(t *testing.T) {
	nCharge := NewCharge(EnvSandbox, "")
	_, errAsaas := nCharge.UpdateById(context.TODO(), "test", UpdateChargeRequest{
		BillingType: "asa",
	})
	assertSuccessNonnull(t, errAsaas)
}

func TestChangeDeleteByIdSuccess(t *testing.T) {
	accessToken, err := getAccessToken()
	assertFatalErrorNonnull(t, err)
	initUndefinedCharge()
	chargeId, err := getUndefinedChargeId()
	assertFatalErrorNonnull(t, err)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nCharge := NewCharge(EnvSandbox, accessToken)
	resp, errAsaas := nCharge.DeleteById(ctx, chargeId)
	assertResponseSuccess(t, resp, errAsaas)
}

func TestChargeDeleteByIdError(t *testing.T) {
	nCharge := NewCharge(EnvSandbox, "")
	resp, errAsaas := nCharge.DeleteById(context.TODO(), "test")
	assertResponseFailure(t, resp, errAsaas)
}

func TestChangeRestoreByIdSuccess(t *testing.T) {
	accessToken, err := getAccessToken()
	assertFatalErrorNonnull(t, err)
	initChargeDeleted()
	chargeId, err := getChargeDeletedId()
	assertFatalErrorNonnull(t, err)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nCharge := NewCharge(EnvSandbox, accessToken)
	resp, errAsaas := nCharge.RestoreById(ctx, chargeId)
	assertResponseSuccess(t, resp, errAsaas)
}

func TestChargeRestoreByIdError(t *testing.T) {
	nCharge := NewCharge(EnvSandbox, "")
	resp, errAsaas := nCharge.RestoreById(context.TODO(), "test")
	assertResponseFailure(t, resp, errAsaas)
}

func TestChangeRefundByIdSuccess(t *testing.T) {
	accessToken, err := getAccessToken()
	assertFatalErrorNonnull(t, err)
	initCreditCardCharge()
	chargeId, err := getCreditCardChargeId()
	assertFatalErrorNonnull(t, err)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nCharge := NewCharge(EnvSandbox, accessToken)
	resp, errAsaas := nCharge.RefundById(ctx, chargeId, RefundRequest{
		Value:       0,
		Description: "",
	})
	assertResponseSuccess(t, resp, errAsaas)
}

func TestChargeRefundByIdError(t *testing.T) {
	nCharge := NewCharge(EnvSandbox, "")
	resp, errAsaas := nCharge.RefundById(context.TODO(), "test", RefundRequest{})
	assertResponseFailure(t, resp, errAsaas)
}

func TestChangeReceiveInCashByIdSuccess(t *testing.T) {
	accessToken, err := getAccessToken()
	assertFatalErrorNonnull(t, err)
	initUndefinedCharge()
	chargeId, err := getUndefinedChargeId()
	assertFatalErrorNonnull(t, err)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	assertFatalErrorNonnull(t, err)
	nCharge := NewCharge(EnvSandbox, accessToken)
	now := time.Now()
	resp, errAsaas := nCharge.ReceiveInCashById(ctx, chargeId, ChargeReceiveInCashRequest{
		PaymentDate:    NewDate(now.Year(), now.Month(), now.Day(), now.Location()),
		Value:          100,
		NotifyCustomer: false,
	})
	assertResponseSuccess(t, resp, errAsaas)
}

func TestChargeReceiveInCashByIdError(t *testing.T) {
	nCharge := NewCharge(EnvSandbox, "")
	_, errAsaas := nCharge.ReceiveInCashById(context.TODO(), "test", ChargeReceiveInCashRequest{})
	assertSuccessNonnull(t, errAsaas)
}

func TestChangeUndoReceivedInCashByIdSuccess(t *testing.T) {
	accessToken, err := getAccessToken()
	assertFatalErrorNonnull(t, err)
	initChargeReceivedInCash()
	chargeId, err := getChargeReceivedInCashId()
	assertFatalErrorNonnull(t, err)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nCharge := NewCharge(EnvSandbox, accessToken)
	resp, errAsaas := nCharge.UndoReceivedInCashById(ctx, chargeId)
	assertResponseSuccess(t, resp, errAsaas)
}

func TestChargeUndoReceivedInCashByIdError(t *testing.T) {
	nCharge := NewCharge(EnvSandbox, "")
	resp, errAsaas := nCharge.UndoReceivedInCashById(context.TODO(), "")
	assertResponseFailure(t, resp, errAsaas)
}

func TestChargeUploadDocumentByIdSuccess(t *testing.T) {
	accessToken, err := getAccessToken()
	assertFatalErrorNonnull(t, err)
	initUndefinedCharge()
	chargeId, err := getUndefinedChargeId()
	assertFatalErrorNonnull(t, err)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	f, err := test.GetSimpleFile()
	assertFatalErrorNonnull(t, err)
	defer func(name string) {
		err = os.Remove(name)
		if err != nil {
			logError("error remove file test:", err)
		}
	}(f.Name())
	nCharge := NewCharge(EnvSandbox, accessToken)
	resp, errAsaas := nCharge.UploadDocumentById(ctx, chargeId, UploadChargeDocumentRequest{
		AvailableAfterPayment: false,
		Type:                  DocumentTypeDocument,
		File:                  f,
	})
	assertResponseSuccess(t, resp, errAsaas)
}

func TestChargeUploadDocumentByIdError(t *testing.T) {
	nCharge := NewCharge(EnvSandbox, "")
	_, errAsaas := nCharge.UploadDocumentById(context.TODO(), "", UploadChargeDocumentRequest{})
	assertSuccessNonnull(t, errAsaas)
}

func TestChargeUpdateDocumentDefinitionsByIdSuccess(t *testing.T) {
	accessToken, err := getAccessToken()
	assertFatalErrorNonnull(t, err)
	initChargeDocumentId()
	chargeId, err := getUndefinedChargeId()
	assertFatalErrorNonnull(t, err)
	docId, err := getChargeDocumentId()
	assertFatalErrorNonnull(t, err)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nCharge := NewCharge(EnvSandbox, accessToken)
	resp, errAsaas := nCharge.UpdateDocumentDefinitionsById(ctx, chargeId, docId, UpdateChargeDocumentDefinitionsRequest{
		AvailableAfterPayment: false,
		Type:                  DocumentTypeContract,
	})
	assertResponseSuccess(t, resp, errAsaas)
}

func TestChargeUpdateDocumentDefinitionsByIdError(t *testing.T) {
	nCharge := NewCharge(EnvSandbox, "")
	ctx := context.TODO()
	_, errAsaas := nCharge.UpdateDocumentDefinitionsById(ctx, "", "", UpdateChargeDocumentDefinitionsRequest{})
	assertSuccessNonnull(t, errAsaas)
}

func TestChargeDeleteDocumentByIdSuccess(t *testing.T) {
	accessToken, err := getAccessToken()
	assertFatalErrorNonnull(t, err)
	initChargeDocumentId()
	chargeId, err := getUndefinedChargeId()
	assertFatalErrorNonnull(t, err)
	docId, err := getChargeDocumentId()
	assertFatalErrorNonnull(t, err)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nCharge := NewCharge(EnvSandbox, accessToken)
	resp, errAsaas := nCharge.DeleteDocumentById(ctx, chargeId, docId)
	assertResponseSuccess(t, resp, errAsaas)
}

func TestChargeDeleteDocumentByIdError(t *testing.T) {
	nCharge := NewCharge(EnvSandbox, "")
	ctx := context.TODO()
	resp, errAsaas := nCharge.DeleteDocumentById(ctx, "test", "test")
	assertResponseFailure(t, resp, errAsaas)
}

func TestChargeGetByIdSuccess(t *testing.T) {
	accessToken, err := getAccessToken()
	assertFatalErrorNonnull(t, err)
	initUndefinedCharge()
	chargeId, err := getUndefinedChargeId()
	assertFatalErrorNonnull(t, err)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nCharge := NewCharge(EnvSandbox, accessToken)
	resp, errAsaas := nCharge.GetById(ctx, chargeId)
	assertResponseSuccess(t, resp, errAsaas)
}

func TestChargeGetByIdError(t *testing.T) {
	nCharge := NewCharge(EnvSandbox, "")
	resp, errAsaas := nCharge.GetDocumentById(context.TODO(), test.GetChargeIdDefault(), "test")
	assertResponseFailure(t, resp, errAsaas)
}

func TestChargeGetCreationLimitSuccess(t *testing.T) {
	accessToken, err := getAccessToken()
	assertFatalErrorNonnull(t, err)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nCharge := NewCharge(EnvSandbox, accessToken)
	resp, errAsaas := nCharge.GetCreationLimit(ctx)
	assertResponseSuccess(t, resp, errAsaas)
}

func TestChargeGetCreationLimitError(t *testing.T) {
	nCharge := NewCharge(EnvSandbox, "")
	resp, errAsaas := nCharge.GetCreationLimit(context.TODO())
	assertResponseSuccess(t, resp, errAsaas)
}

func TestChargeGetStatusByIdSuccess(t *testing.T) {
	accessToken, err := getAccessToken()
	assertFatalErrorNonnull(t, err)
	initUndefinedCharge()
	chargeId, err := getUndefinedChargeId()
	assertFatalErrorNonnull(t, err)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nCharge := NewCharge(EnvSandbox, accessToken)
	resp, errAsaas := nCharge.GetStatusById(ctx, chargeId)
	assertResponseSuccess(t, resp, errAsaas)
}

func TestChargeGetStatusByIdError(t *testing.T) {
	nCharge := NewCharge(EnvSandbox, "")
	resp, errAsaas := nCharge.GetStatusById(context.TODO(), "test")
	assertResponseFailure(t, resp, errAsaas)
}

func TestChargeGetIdentificationFieldByIdSuccess(t *testing.T) {
	accessToken, err := getAccessToken()
	assertFatalErrorNonnull(t, err)
	initBankSlipCharge()
	chargeId, err := getBankSlipChargeId()
	assertFatalErrorNonnull(t, err)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nCharge := NewCharge(EnvSandbox, accessToken)
	resp, errAsaas := nCharge.GetIdentificationFieldById(ctx, chargeId)
	assertResponseSuccess(t, resp, errAsaas)
}

func TestChargeGetIdentificationFieldByIdError(t *testing.T) {
	nCharge := NewCharge(EnvSandbox, "")
	resp, errAsaas := nCharge.GetIdentificationFieldById(context.TODO(), "test")
	assertResponseFailure(t, resp, errAsaas)
}

func TestChargeGetPixQrCodeByIdSuccess(t *testing.T) {
	accessToken, err := getAccessToken()
	assertFatalErrorNonnull(t, err)
	initPixCharge()
	chargeId, err := getPixChargeId()
	assertFatalErrorNonnull(t, err)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nCharge := NewCharge(EnvSandbox, accessToken)
	resp, errAsaas := nCharge.GetPixQrCodeById(ctx, chargeId)
	assertResponseSuccess(t, resp, errAsaas)
}

func TestChargeGetPixQrCodeByIdError(t *testing.T) {
	nCharge := NewCharge(EnvSandbox, "")
	resp, errAsaas := nCharge.GetPixQrCodeById(context.TODO(), "test")
	assertResponseFailure(t, resp, errAsaas)
}

func TestChargeGetDocumentByIdSuccess(t *testing.T) {
	accessToken, err := getAccessToken()
	assertFatalErrorNonnull(t, err)
	initChargeDocumentId()
	chargeId, err := getUndefinedChargeId()
	assertFatalErrorNonnull(t, err)
	docId, err := getChargeDocumentId()
	assertFatalErrorNonnull(t, err)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nCharge := NewCharge(EnvSandbox, accessToken)
	resp, errAsaas := nCharge.GetDocumentById(ctx, chargeId, docId)
	assertResponseSuccess(t, resp, errAsaas)
}

func TestChargeGetDocumentByIdError(t *testing.T) {
	nCharge := NewCharge(EnvSandbox, "")
	resp, errAsaas := nCharge.GetDocumentById(context.TODO(), "test", "test")
	assertResponseFailure(t, resp, errAsaas)
}

func TestChargeGetAllSuccess(t *testing.T) {
	accessToken, err := getAccessToken()
	assertFatalErrorNonnull(t, err)
	initChargeDocumentId()
	_, err = getUndefinedChargeId()
	assertFatalErrorNonnull(t, err)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nCharge := NewCharge(EnvSandbox, accessToken)
	resp, errAsaas := nCharge.GetAll(ctx, GetAllChargesRequest{
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
	assertResponseSuccess(t, resp, errAsaas)
}

func TestChargeGetAllError(t *testing.T) {
	nCharge := NewCharge(EnvSandbox, "")
	resp, errAsaas := nCharge.GetAll(context.TODO(), GetAllChargesRequest{})
	assertResponseFailure(t, resp, errAsaas)
}

func TestChargeGetAllDocumentsByIdSuccess(t *testing.T) {
	accessToken, err := getAccessToken()
	assertFatalErrorNonnull(t, err)
	initChargeDocumentId()
	chargeId, err := getUndefinedChargeId()
	assertFatalErrorNonnull(t, err)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nCharge := NewCharge(EnvSandbox, accessToken)
	resp, errAsaas := nCharge.GetAllDocumentsById(ctx, chargeId, PageableDefaultRequest{
		Offset: 0,
		Limit:  10,
	})
	assertResponseSuccess(t, resp, errAsaas)
}

func TestChargeGetAllDocumentsByIdError(t *testing.T) {
	nCharge := NewCharge(EnvSandbox, "")
	resp, errAsaas := nCharge.GetAllDocumentsById(context.TODO(), "test", PageableDefaultRequest{})
	assertResponseFailure(t, resp, errAsaas)
}
