package asaas

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/GabrielHCataldo/go-asaas/internal/test"
	"github.com/GabrielHCataldo/go-asaas/internal/util"
	"github.com/mvrilo/go-cpf"
	"os"
	"testing"
	"time"
)

const EnvAccessToken = "ASAAS_ACCESS_TOKEN"
const EnvCustomerId = "ASAAS_CUSTOMER_ID"
const EnvCreditCardChargeId = "ASAAS_CREDIT_CARD_CHARGE_ID"
const EnvPixChargeId = "ASAAS_PIX_CHARGE_ID"
const EnvBankSlipChargeId = "ASAAS_BANK_SLIP_CHARGE_ID"
const EnvUndefinedChargeId = "ASAAS_UNDEFINED_CHARGE_ID"
const EnvChargeInstallmentId = "ASAAS_CHARGE_INSTALLMENT_ID"
const EnvChargeIdentificationField = "ASAAS_CHARGE_IDENTIFICATION_FIELD"
const EnvChargePixQrCodePayload = "ASAAS_CHARGE_PIX_QRCODE_PAYLOAD"
const EnvChargeDeletedId = "ASAAS_CHARGE_DELETED_ID"
const EnvChargeReceivedInCashId = "ASAAS_CHARGE_RECEIVED_IN_CASH_ID"
const EnvChargeDocumentId = "ASAAS_CHARGE_DOCUMENT_ID"
const EnvAnticipationId = "ASAAS_ANTICIPATION_ID"
const EnvBillPaymentId = "ASAAS_BILL_PAYMENT_ID"
const MessageAccessTokenRequired = "ASAAS_ACCESS_TOKEN env is required"
const MessageCustomerIdRequired = "ASAAS_CUSTOMER_ID env is required"
const MessageCreditCardChargeIdRequired = "ASAAS_CREDIT_CARD_CHARGE_ID env is required"
const MessagePixChargeIdRequired = "ASAAS_PIX_CHARGE_ID env is required"
const MessageBankSlipChargeIdRequired = "ASAAS_BANK_SLIP_CHARGE_ID env is required"
const MessageUndefinedChargeIdRequired = "ASAAS_UNDEFINED_CHARGE_ID env is required"
const MessageChargeInstallmentIdRequired = "ASAAS_CHARGE_INSTALLMENT_ID env is required"
const MessageChargeIdentificationFieldRequired = "ASAAS_CHARGE_IDENTIFICATION_FIELD env is required"
const MessageChargePixQrCodePayloadRequired = "ASAAS_CHARGE_PIX_QRCODE_PAYLOAD env is required"
const MessageChargeDeletedId = "ASAAS_CHARGE_DELETED_ID env is required"
const MessageChargeReceivedInCashId = "ASAAS_CHARGE_RECEIVED_IN_CASH_ID env is required"
const MessageChargeDocumentId = "ASAAS_CHARGE_DOCUMENT_ID env is required"
const MessageAnticipationIdRequired = "ASAAS_ANTICIPATION_ID env is required"
const MessageBillPaymentIdRequired = "ASAAS_BILL_PAYMENT_ID env is required"

func TestMain(m *testing.M) {
	code := m.Run()
	clearCustomerId()
	clearCreditCardChargeId()
	clearPixChargeId()
	clearBillPaymentId()
	os.Exit(code)
}

func getAccessToken() (string, error) {
	accessToken := os.Getenv(EnvAccessToken)
	if util.IsBlank(&accessToken) {
		return accessToken, errors.New(MessageAccessTokenRequired)
	}
	return accessToken, nil
}

func getCustomerId() (string, error) {
	customerId := os.Getenv(EnvCustomerId)
	if util.IsBlank(&customerId) {
		return customerId, errors.New(MessageCustomerIdRequired)
	}
	return customerId, nil
}

func getCreditCardChargeId() (string, error) {
	chargeId := os.Getenv(EnvCreditCardChargeId)
	if util.IsBlank(&chargeId) {
		return chargeId, errors.New(MessageCreditCardChargeIdRequired)
	}
	return chargeId, nil
}

func getPixChargeId() (string, error) {
	chargeId := os.Getenv(EnvPixChargeId)
	if util.IsBlank(&chargeId) {
		return chargeId, errors.New(MessagePixChargeIdRequired)
	}
	return chargeId, nil
}

func getBankSlipChargeId() (string, error) {
	chargeId := os.Getenv(EnvBankSlipChargeId)
	if util.IsBlank(&chargeId) {
		return chargeId, errors.New(MessageBankSlipChargeIdRequired)
	}
	return chargeId, nil
}

func getUndefinedChargeId() (string, error) {
	chargeId := os.Getenv(EnvUndefinedChargeId)
	if util.IsBlank(&chargeId) {
		return chargeId, errors.New(MessageUndefinedChargeIdRequired)
	}
	return chargeId, nil
}

func getChargeInstallmentId() (string, error) {
	installmentId := os.Getenv(EnvChargeInstallmentId)
	if util.IsBlank(&installmentId) {
		return installmentId, errors.New(MessageChargeInstallmentIdRequired)
	}
	return installmentId, nil
}

func getChargeIdentificationField() (string, error) {
	identificationField := os.Getenv(EnvChargeIdentificationField)
	if util.IsBlank(&identificationField) {
		return identificationField, errors.New(MessageChargeIdentificationFieldRequired)
	}
	return identificationField, nil
}

func getChargePixQrCodePayload() (string, error) {
	pixQrCodePayload := os.Getenv(EnvChargePixQrCodePayload)
	if util.IsBlank(&pixQrCodePayload) {
		return pixQrCodePayload, errors.New(MessageChargePixQrCodePayloadRequired)
	}
	return pixQrCodePayload, nil
}

func getChargeDeletedId() (string, error) {
	chargeId := os.Getenv(EnvChargeDeletedId)
	if util.IsBlank(&chargeId) {
		return chargeId, errors.New(MessageChargeDeletedId)
	}
	return chargeId, nil
}

func getChargeReceivedInCashId() (string, error) {
	chargeId := os.Getenv(EnvChargeReceivedInCashId)
	if util.IsBlank(&chargeId) {
		return chargeId, errors.New(MessageChargeReceivedInCashId)
	}
	return chargeId, nil
}

func getChargeDocumentId() (string, error) {
	chargeId := os.Getenv(EnvChargeDocumentId)
	if util.IsBlank(&chargeId) {
		return chargeId, errors.New(MessageChargeDocumentId)
	}
	return chargeId, nil
}

func getAnticipationId() (string, error) {
	anticipationId := os.Getenv(EnvAnticipationId)
	if util.IsBlank(&anticipationId) {
		return anticipationId, errors.New(MessageAnticipationIdRequired)
	}
	return anticipationId, nil
}

func getBillPaymentId() (string, error) {
	billPaymentId := os.Getenv(EnvBillPaymentId)
	if util.IsBlank(&billPaymentId) {
		return billPaymentId, errors.New(MessageBillPaymentIdRequired)
	}
	return billPaymentId, nil
}

func initCustomer() {
	clearCustomerId()
	accessToken, err := getAccessToken()
	if err != nil {
		logError("error getAccessToken:", err)
		return
	}
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	customerAsaas := NewCustomer(EnvSandbox, accessToken)
	customerResp, errAsaas := customerAsaas.Create(ctx, CustomerRequest{
		Name:    "Unit test go",
		CpfCnpj: cpf.Generate(),
	})
	if errAsaas != nil || customerResp.IsFailure() {
		logError("error create customer resp:", customerResp, "err: ", err)
		return
	}
	err = os.Setenv(EnvCustomerId, customerResp.Id)
	logError("error set customer id env:", err)
}

func initCreditCardCharge() {
	clearCreditCardChargeId()
	accessToken, err := getAccessToken()
	if err != nil {
		logError("error getAccessToken:", err)
		return
	}
	customerId, err := getCustomerId()
	if err != nil {
		logError("error get customer id env:", err)
		return
	}
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	cReq := &CreateChargeRequest{}
	err = json.Unmarshal(test.GetCreateCreditCardChargeRequestDefault(), cReq)
	if err != nil {
		logError("error parse charge request:", err)
		return
	}
	cReq.Customer = customerId
	chargeAsaas := NewCharge(EnvSandbox, accessToken)
	chargeResp, errAsaas := chargeAsaas.Create(ctx, *cReq)
	if errAsaas != nil || chargeResp.IsFailure() {
		logError("error create charge resp:", chargeResp, "err: ", err)
		return
	}
	err = os.Setenv(EnvCreditCardChargeId, chargeResp.Id)
	if err != nil {
		logError("error set", EnvCreditCardChargeId, ":", err)
		return
	}
	err = os.Setenv(EnvChargeInstallmentId, chargeResp.Installment)
	logError("error set", EnvChargeInstallmentId, ":", err)
}

func initPixCharge() {
	clearPixChargeId()
	accessToken, err := getAccessToken()
	if err != nil {
		logError("error getAccessToken:", err)
		return
	}
	customerId, err := getCustomerId()
	if err != nil {
		logError("error get customer id env:", err)
		return
	}
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	cReq := &CreateChargeRequest{}
	err = json.Unmarshal(test.GetCreatePixChargeRequestDefault(), cReq)
	if err != nil {
		logError("error parse charge request:", err)
		return
	}
	cReq.Customer = customerId
	chargeAsaas := NewCharge(EnvSandbox, accessToken)
	chargeResp, errAsaas := chargeAsaas.Create(ctx, *cReq)
	if errAsaas != nil || chargeResp.IsFailure() {
		logError("error create charge resp:", chargeResp, "err: ", err)
		return
	}
	err = os.Setenv(EnvPixChargeId, chargeResp.Id)
	if err != nil {
		logError("error set", EnvPixChargeId, ":", err)
		return
	}
	pixQrCodeResp, errAsaas := chargeAsaas.GetPixQrCodeById(ctx, chargeResp.Id)
	if errAsaas != nil || pixQrCodeResp.IsFailure() {
		logError("error get charge pix qrcode resp:", pixQrCodeResp, "err: ", err)
		return
	}
	err = os.Setenv(EnvChargePixQrCodePayload, pixQrCodeResp.Payload)
	logError("error set", EnvChargePixQrCodePayload, ":", err)
}

func initBankSlipCharge() {
	clearBankSlipChargeId()
	accessToken, err := getAccessToken()
	if err != nil {
		logError("error getAccessToken:", err)
		return
	}
	customerId, err := getCustomerId()
	if err != nil {
		logError("error get customer id env:", err)
		return
	}
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	cReq := &CreateChargeRequest{}
	err = json.Unmarshal(test.GetCreateBankSlipChargeRequestDefault(), cReq)
	if err != nil {
		logError("error parse charge request:", err)
		return
	}
	cReq.Customer = customerId
	chargeAsaas := NewCharge(EnvSandbox, accessToken)
	chargeResp, errAsaas := chargeAsaas.Create(ctx, *cReq)
	if errAsaas != nil || chargeResp.IsFailure() {
		logError("error create charge resp:", chargeResp, "err: ", err)
		return
	}
	err = os.Setenv(EnvBankSlipChargeId, chargeResp.Id)
	if err != nil {
		logError("error set", EnvBankSlipChargeId, ":", err)
		return
	}
	identificationFieldResp, errAsaas := chargeAsaas.GetIdentificationFieldById(ctx, chargeResp.Id)
	if errAsaas != nil || identificationFieldResp.IsFailure() {
		logError("error get charge identification field resp:", identificationFieldResp, "err: ", err)
		return
	}
	err = os.Setenv(EnvChargeIdentificationField, identificationFieldResp.IdentificationField)
	logError("error set", EnvChargeIdentificationField, ":", err)
}

func initUndefinedCharge() {
	clearUndefinedChargeId()
	accessToken, err := getAccessToken()
	if err != nil {
		logError("error getAccessToken:", err)
		return
	}
	customerId, err := getCustomerId()
	if err != nil {
		logError("error get customer id env:", err)
		return
	}
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	cReq := &CreateChargeRequest{}
	err = json.Unmarshal(test.GetCreateUndefinedChargeRequestDefault(), cReq)
	if err != nil {
		logError("error parse charge request:", err)
		return
	}
	cReq.Customer = customerId
	chargeAsaas := NewCharge(EnvSandbox, accessToken)
	chargeResp, errAsaas := chargeAsaas.Create(ctx, *cReq)
	if errAsaas != nil || chargeResp.IsFailure() {
		logError("error create charge resp:", chargeResp, "err: ", err)
		return
	}
	err = os.Setenv(EnvUndefinedChargeId, chargeResp.Id)
	logError("error set", EnvUndefinedChargeId, ":", err)
}

func initChargeDeleted() {
	accessToken, err := getAccessToken()
	if err != nil {
		logError("error getAccessToken:", err)
		return
	}
	initUndefinedCharge()
	chargeId, err := getUndefinedChargeId()
	if err != nil {
		logError("error getUndefinedChargeId:", err)
		return
	}
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	chargeAsaas := NewCharge(EnvSandbox, accessToken)
	deleteResp, errAsaas := chargeAsaas.DeleteById(ctx, chargeId)
	if errAsaas != nil || deleteResp.IsFailure() {
		logError("error delete charge resp:", deleteResp, "err: ", err)
		return
	}
	err = os.Setenv(EnvChargeDeletedId, deleteResp.Id)
	logError("error set", EnvChargeDeletedId, ":", err)
}

func initChargeReceivedInCash() {
	accessToken, err := getAccessToken()
	if err != nil {
		logError("error getAccessToken:", err)
		return
	}
	initUndefinedCharge()
	chargeId, err := getUndefinedChargeId()
	if err != nil {
		logError("error getUndefinedChargeId:", err)
		return
	}
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	now := time.Now()
	chargeAsaas := NewCharge(EnvSandbox, accessToken)
	receiveInCashResp, errAsaas := chargeAsaas.ReceiveInCashById(ctx, chargeId, ChargeReceiveInCashRequest{
		PaymentDate: NewDate(now.Year(), now.Month(), now.Day(), now.Location()),
		Value:       100,
	})
	if errAsaas != nil || receiveInCashResp.IsFailure() {
		logError("error receive in cash resp:", receiveInCashResp, "err: ", err)
		return
	}
	err = os.Setenv(EnvChargeReceivedInCashId, receiveInCashResp.Id)
	logError("error set", EnvChargeReceivedInCashId, ":", err)
}

func initChargeDocumentId() {
	accessToken, err := getAccessToken()
	if err != nil {
		logError("error getAccessToken:", err)
		return
	}
	initUndefinedCharge()
	chargeId, err := getUndefinedChargeId()
	if err != nil {
		logError("error getUndefinedChargeId:", err)
		return
	}
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	f, err := test.GetSimpleFile()
	if err != nil {
		logError("error GetSimpleFile:", err)
		return
	}
	defer func(name string) {
		err = os.Remove(name)
		if err != nil {
			logError("error remove file test:", err)
		}
	}(f.Name())
	nCharge := NewCharge(EnvSandbox, accessToken)
	resp, errAsaas := nCharge.UploadDocumentById(ctx, chargeId, UploadChargeDocumentRequest{
		Type: DocumentTypeDocument,
		File: f,
	})
	if errAsaas != nil || resp.IsFailure() {
		logError("error upload document charge resp:", resp, "err: ", err)
		return
	}
	err = os.Setenv(EnvChargeDocumentId, resp.Id)
	logError("error set", EnvChargeDocumentId, ":", err)
}

func initAnticipation() {
	accessToken, err := getAccessToken()
	if err != nil {
		logError("error getAccessToken:", err)
		return
	}
	initCreditCardCharge()
	chargeId, err := getCreditCardChargeId()
	if err != nil {
		logError("error getCreditCardChargeId:", err)
		return
	}
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	anticipationAsaas := NewAnticipation(EnvSandbox, accessToken)
	anticipationResp, errAsaas := anticipationAsaas.Request(ctx, AnticipationRequest{
		Payment: chargeId,
	})
	if errAsaas != nil || anticipationResp.IsFailure() {
		logError("error create anticipation resp:", anticipationResp, "err: ", err)
		return
	}
	err = os.Setenv(EnvAnticipationId, anticipationResp.Id)
	logError("error set", EnvAnticipationId, ":", err)
}

func initBillPayment() {
	clearBillPaymentId()
	accessToken, err := getAccessToken()
	if err != nil {
		logError("error getAccessToken:", err)
		return
	}
	initBankSlipCharge()
	_, err = getBankSlipChargeId()
	if err != nil {
		logError("error getBankSlipChargeId:", err)
		return
	}
	identificationField, err := getChargeIdentificationField()
	if err != nil {
		logError("error getChargeIdentificationField:", err)
		return
	}
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	now := time.Now()
	scheduleDate := NewDate(now.Year(), now.Month(), now.Day()+1, now.Location())
	billPaymentAsaas := NewBillPayment(EnvSandbox, accessToken)
	billPaymentResp, errAsaas := billPaymentAsaas.Create(ctx, CreateBillPaymentRequest{
		IdentificationField: identificationField,
		ScheduleDate:        &scheduleDate,
	})
	if errAsaas != nil || billPaymentResp.IsFailure() {
		logError("error create bill payment resp:", billPaymentResp, "err: ", err)
		return
	}
	err = os.Setenv(EnvBillPaymentId, billPaymentResp.Id)
	logError("error set", EnvBillPaymentId, ":", err)
}

func clearBillPaymentId() {
	accessToken, err := getAccessToken()
	if err != nil {
		logError("error getAccessToken:", err)
		return
	}
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	billPaymentId, err := getBillPaymentId()
	if err != nil || util.IsBlank(&billPaymentId) {
		return
	}
	billPaymentAsaas := NewBillPayment(EnvSandbox, accessToken)
	billPaymentAsaas.CancelById(ctx, billPaymentId)
	_ = os.Unsetenv(EnvBillPaymentId)
}

func clearCustomerId() {
	accessToken, err := getAccessToken()
	if err != nil {
		logError("error getAccessToken:", err)
		return
	}
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	customerId, err := getCustomerId()
	if err != nil || util.IsBlank(&customerId) {
		return
	}
	customerAsaas := NewCustomer(EnvSandbox, accessToken)
	customerAsaas.DeleteById(ctx, customerId)
	_ = os.Unsetenv(EnvCustomerId)
}

func clearCreditCardChargeId() {
	accessToken, err := getAccessToken()
	if err != nil {
		logError("error getAccessToken:", err)
		return
	}
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	chargeId, err := getCreditCardChargeId()
	if err != nil || util.IsBlank(&chargeId) {
		return
	}
	chargeAsaas := NewCharge(EnvSandbox, accessToken)
	chargeAsaas.RefundById(ctx, chargeId, RefundRequest{
		Description: "estorno test clearCreditCardChargeId",
	})
	chargeAsaas.DeleteById(ctx, chargeId)
	_ = os.Unsetenv(EnvCreditCardChargeId)
}

func clearPixChargeId() {
	accessToken, err := getAccessToken()
	if err != nil {
		logError("error getAccessToken:", err)
		return
	}
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	chargeId, err := getCreditCardChargeId()
	if err != nil || util.IsBlank(&chargeId) {
		return
	}
	chargeAsaas := NewCharge(EnvSandbox, accessToken)
	chargeAsaas.DeleteById(ctx, chargeId)
	_ = os.Unsetenv(EnvPixChargeId)
}

func clearBankSlipChargeId() {
	accessToken, err := getAccessToken()
	if err != nil {
		logError("error getAccessToken:", err)
		return
	}
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	chargeId, err := getBankSlipChargeId()
	if err != nil || util.IsBlank(&chargeId) {
		return
	}
	chargeAsaas := NewCharge(EnvSandbox, accessToken)
	chargeAsaas.DeleteById(ctx, chargeId)
	_ = os.Unsetenv(EnvBankSlipChargeId)
}

func clearUndefinedChargeId() {
	accessToken, err := getAccessToken()
	if err != nil {
		logError("error getAccessToken:", err)
		return
	}
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	chargeId, err := getBankSlipChargeId()
	if err != nil || util.IsBlank(&chargeId) {
		return
	}
	chargeAsaas := NewCharge(EnvSandbox, accessToken)
	chargeAsaas.DeleteById(ctx, chargeId)
	_ = os.Unsetenv(EnvUndefinedChargeId)
}
