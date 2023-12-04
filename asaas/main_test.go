package asaas

import (
	"context"
	"encoding/json"
	"github.com/GabrielHCataldo/go-asaas/internal/test"
	"github.com/GabrielHCataldo/go-asaas/internal/util"
	"github.com/mvrilo/go-cpf"
	"os"
	"testing"
	"time"
)

const EnvAccessToken = "ASAAS_ACCESS_TOKEN"
const EnvAccessTokenSecondary = "ASAAS_ACCESS_TOKEN_SECONDARY"
const EnvFileName = "ASAAS_FILE_NAME"
const EnvImageName = "ASAAS_IMAGE_NAME"
const EnvCustomerId = "ASAAS_CUSTOMER_ID"
const EnvCustomerDeletedId = "ASAAS_CUSTOMER_DELETED_ID"
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

func init() {
	initFile()
	initImage()
	//initFiscalInfo()
}

func TestMain(m *testing.M) {
	code := m.Run()
	clearCustomerId()
	clearCreditCardChargeId()
	clearPixChargeId()
	clearUndefinedChargeId()
	clearBillPaymentId()
	clearFileName()
	os.Exit(code)
}

func getEnvValue(env string) string {
	v := os.Getenv(env)
	if util.IsBlank(&v) {
		logError("error getEnvValue:", env, " is required env")
	}
	return v
}

func setEnv(env, v string) bool {
	err := os.Setenv(env, v)
	if err != nil {
		logError("error set", env, ":", err)
	} else {
		logInfo(EnvSandbox, "set", env, "successfully")
	}
	return err != nil
}

func initCustomer() {
	accessToken := getEnvValue(EnvAccessToken)
	if util.IsBlank(&accessToken) {
		return
	}
	clearCustomerId()
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	customerAsaas := NewCustomer(EnvSandbox, accessToken)
	resp, err := customerAsaas.Create(ctx, CreateCustomerRequest{
		Name:    "Unit test go",
		CpfCnpj: cpf.Generate(),
	})
	if err != nil || resp.IsFailure() {
		logError("error resp:", resp, "err: ", err)
		return
	}
	setEnv(EnvCustomerId, resp.Id)
}

func initCreditCardCharge(authorizeOnly bool, withInstallment bool) {
	accessToken := getEnvValue(EnvAccessToken)
	if util.IsBlank(&accessToken) {
		return
	}
	clearCreditCardChargeId()
	initCustomer()
	customerId := getEnvValue(EnvCustomerId)
	if util.IsBlank(&customerId) {
		return
	}
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	cReq := &CreateChargeRequest{}
	err := json.Unmarshal(test.GetCreateCreditCardChargeRequestDefault(), cReq)
	if err != nil {
		logError("error json.Unmarshal:", err)
		return
	}
	cReq.Customer = customerId
	cReq.AuthorizeOnly = authorizeOnly
	if withInstallment {
		cReq.InstallmentCount = 2
		cReq.InstallmentValue = cReq.Value / float64(cReq.InstallmentCount)
	}
	chargeAsaas := NewCharge(EnvSandbox, accessToken)
	resp, err := chargeAsaas.Create(ctx, *cReq)
	if err != nil || resp.IsFailure() {
		logError("error resp:", resp, "err: ", err)
		return
	}
	success := setEnv(EnvCreditCardChargeId, resp.Id)
	if !success {
		return
	}
	setEnv(EnvChargeInstallmentId, resp.Installment)
}

func initPixCharge() {
	accessToken := getEnvValue(EnvAccessToken)
	if util.IsBlank(&accessToken) {

		return
	}
	clearPixChargeId()
	initCustomer()
	customerId := getEnvValue(EnvCustomerId)
	if util.IsBlank(&customerId) {
		return
	}
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	cReq := &CreateChargeRequest{}
	err := json.Unmarshal(test.GetCreatePixChargeRequestDefault(), cReq)
	if err != nil {
		logError("error json.Unmarshal:", err)
		return
	}
	cReq.Customer = customerId
	chargeAsaas := NewCharge(EnvSandbox, accessToken)
	resp, err := chargeAsaas.Create(ctx, *cReq)
	if err != nil || resp.IsFailure() {
		logError("error resp:", resp, "err: ", err)
		return
	}
	success := setEnv(EnvPixChargeId, resp.Id)
	if !success {
		return
	}
	pixQrCodeResp, err := chargeAsaas.GetPixQrCodeById(ctx, resp.Id)
	if err != nil || pixQrCodeResp.IsFailure() {
		logError("error resp:", pixQrCodeResp, "err: ", err)
		return
	}
	setEnv(EnvChargePixQrCodePayload, pixQrCodeResp.Payload)
}

func initBankSlipCharge() {
	accessToken := getEnvValue(EnvAccessToken)
	if util.IsBlank(&accessToken) {

		return
	}
	clearBankSlipChargeId()
	initCustomer()
	customerId := getEnvValue(EnvCustomerId)
	if util.IsBlank(&customerId) {
		return
	}
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	cReq := &CreateChargeRequest{}
	err := json.Unmarshal(test.GetCreateBankSlipChargeRequestDefault(), cReq)
	if err != nil {
		logError("error json.Unmarshal:", err)
		return
	}
	cReq.Customer = customerId
	chargeAsaas := NewCharge(EnvSandbox, accessToken)
	resp, err := chargeAsaas.Create(ctx, *cReq)
	if err != nil || resp.IsFailure() {
		logError("error resp:", resp, "err: ", err)
		return
	}
	success := setEnv(EnvBankSlipChargeId, resp.Id)
	if !success {
		return
	}
	identificationFieldResp, err := chargeAsaas.GetIdentificationFieldById(ctx, resp.Id)
	if err != nil || identificationFieldResp.IsFailure() {
		logError("error resp:", identificationFieldResp, "err: ", err)
		return
	}
	setEnv(EnvChargeIdentificationField, identificationFieldResp.IdentificationField)
}

func initUndefinedCharge() {
	accessToken := getEnvValue(EnvAccessToken)
	if util.IsBlank(&accessToken) {

		return
	}
	clearUndefinedChargeId()
	initCustomer()
	customerId := getEnvValue(EnvCustomerId)
	if util.IsBlank(&customerId) {
		return
	}
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	cReq := &CreateChargeRequest{}
	err := json.Unmarshal(test.GetCreateUndefinedChargeRequestDefault(), cReq)
	if err != nil {
		logError("error json.Unmarshal:", err)
		return
	}
	cReq.Customer = customerId
	chargeAsaas := NewCharge(EnvSandbox, accessToken)
	resp, err := chargeAsaas.Create(ctx, *cReq)
	if err != nil || resp.IsFailure() {
		logError("error resp:", resp, "err: ", err)
		return
	}
	setEnv(EnvUndefinedChargeId, resp.Id)
}

func initChargeDeleted() {
	accessToken := getEnvValue(EnvAccessToken)
	if util.IsBlank(&accessToken) {

		return
	}
	initUndefinedCharge()
	chargeId := getEnvValue(EnvUndefinedChargeId)
	if util.IsBlank(&chargeId) {
		return
	}
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	chargeAsaas := NewCharge(EnvSandbox, accessToken)
	resp, err := chargeAsaas.DeleteById(ctx, chargeId)
	if err != nil || resp.IsFailure() {
		logError("error resp:", resp, "err: ", err)
		return
	}
	setEnv(EnvChargeDeletedId, resp.Id)
}

func initChargeReceivedInCash() {
	accessToken := getEnvValue(EnvAccessToken)
	if util.IsBlank(&accessToken) {

		return
	}
	initUndefinedCharge()
	chargeId := getEnvValue(EnvUndefinedChargeId)
	if util.IsBlank(&chargeId) {
		return
	}
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	now := time.Now()
	chargeAsaas := NewCharge(EnvSandbox, accessToken)
	resp, err := chargeAsaas.ReceiveInCashById(ctx, chargeId, ChargeReceiveInCashRequest{
		PaymentDate: NewDate(now.Year(), now.Month(), now.Day(), now.Location()),
		Value:       100,
	})
	if err != nil || resp.IsFailure() {
		logError("error resp:", resp, "err: ", err)
		return
	}
	setEnv(EnvChargeReceivedInCashId, resp.Id)
}

func initChargeDocumentId() {
	accessToken := getEnvValue(EnvAccessToken)
	if util.IsBlank(&accessToken) {
		return
	}
	initUndefinedCharge()
	chargeId := getEnvValue(EnvUndefinedChargeId)
	if util.IsBlank(&chargeId) {
		return
	}
	f, err := os.Open(getEnvValue(EnvFileName))
	if err != nil {
		return
	}
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nCharge := NewCharge(EnvSandbox, accessToken)
	resp, err := nCharge.UploadDocumentById(ctx, chargeId, UploadChargeDocumentRequest{
		Type: DocumentTypeDocument,
		File: f,
	})
	if err != nil || resp.IsFailure() {
		logError("error resp:", resp, "err: ", err)
		return
	}
	setEnv(EnvChargeDocumentId, resp.Id)
}

func initAnticipation() {
	accessToken := getEnvValue(EnvAccessToken)
	if util.IsBlank(&accessToken) {

		return
	}
	initCreditCardCharge(true, false)
	chargeId := getEnvValue(EnvCreditCardChargeId)
	if util.IsBlank(&chargeId) {
		return
	}
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	anticipationAsaas := NewAnticipation(EnvSandbox, accessToken)
	resp, err := anticipationAsaas.Request(ctx, AnticipationRequest{
		Payment: chargeId,
	})
	if err != nil || resp.IsFailure() {
		logError("error resp:", resp, "err: ", err)
		return
	}
	setEnv(EnvAnticipationId, resp.Id)
}

func initBillPayment() {
	accessToken := getEnvValue(EnvAccessToken)
	if util.IsBlank(&accessToken) {

		return
	}
	clearBillPaymentId()
	initBankSlipCharge()
	chargeId := getEnvValue(EnvBankSlipChargeId)
	if util.IsBlank(&chargeId) {
		return
	}
	identificationField := getEnvValue(EnvChargeIdentificationField)
	if util.IsBlank(&chargeId) {
		return
	}
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	now := time.Now()
	billPaymentAsaas := NewBillPayment(EnvSandbox, accessToken)
	resp, err := billPaymentAsaas.Create(ctx, CreateBillPaymentRequest{
		IdentificationField: identificationField,
		ScheduleDate:        NewDate(now.Year(), now.Month(), now.Day()+1, now.Location()),
	})
	if err != nil || resp.IsFailure() {
		logError("error create bill payment resp:", resp, "err: ", err)
		return
	}
	setEnv(EnvBillPaymentId, resp.Id)
}

func initFile() {
	f, err := test.GetSimpleFile()
	if err != nil {
		logError("error init file GetSimpleFile:", err)
		return
	}
	setEnv(EnvFileName, f.Name())
}

func initImage() {
	f, err := test.GetSimpleImage()
	if err != nil {
		logError("error init image GetSimpleImage:", err)
		return
	}
	setEnv(EnvImageName, f.Name())
}

func clearCustomerId() {
	accessToken := getEnvValue(EnvAccessToken)
	if util.IsBlank(&accessToken) {

		return
	}
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	customerId := getEnvValue(EnvCustomerId)
	if util.IsBlank(&customerId) {
		return
	}
	customerAsaas := NewCustomer(EnvSandbox, accessToken)
	_, _ = customerAsaas.DeleteById(ctx, customerId)
	_ = os.Unsetenv(EnvCustomerId)
}

func clearCreditCardChargeId() {
	accessToken := getEnvValue(EnvAccessToken)
	if util.IsBlank(&accessToken) {

		return
	}
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	chargeId := getEnvValue(EnvCreditCardChargeId)
	if util.IsBlank(&chargeId) {
		return
	}
	chargeAsaas := NewCharge(EnvSandbox, accessToken)
	_, _ = chargeAsaas.RefundById(ctx, chargeId, RefundRequest{
		Description: "unit test golang",
	})
	_, _ = chargeAsaas.DeleteById(ctx, chargeId)
	_ = os.Unsetenv(EnvCreditCardChargeId)
}

func clearPixChargeId() {
	accessToken := getEnvValue(EnvAccessToken)
	if util.IsBlank(&accessToken) {

		return
	}
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	chargeId := getEnvValue(EnvCreditCardChargeId)
	if util.IsBlank(&chargeId) {
		return
	}
	chargeAsaas := NewCharge(EnvSandbox, accessToken)
	_, _ = chargeAsaas.DeleteById(ctx, chargeId)
	_ = os.Unsetenv(EnvPixChargeId)
}

func clearBankSlipChargeId() {
	accessToken := getEnvValue(EnvAccessToken)
	if util.IsBlank(&accessToken) {

		return
	}
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	chargeId := getEnvValue(EnvBankSlipChargeId)
	if util.IsBlank(&chargeId) {
		return
	}
	chargeAsaas := NewCharge(EnvSandbox, accessToken)
	_, _ = chargeAsaas.DeleteById(ctx, chargeId)
	_ = os.Unsetenv(EnvBankSlipChargeId)
}

func clearUndefinedChargeId() {
	accessToken := getEnvValue(EnvAccessToken)
	if util.IsBlank(&accessToken) {

		return
	}
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	chargeId := getEnvValue(EnvBankSlipChargeId)
	if util.IsBlank(&chargeId) {
		return
	}
	chargeAsaas := NewCharge(EnvSandbox, accessToken)
	_, _ = chargeAsaas.DeleteById(ctx, chargeId)
	_ = os.Unsetenv(EnvUndefinedChargeId)
}

func clearBillPaymentId() {
	accessToken := getEnvValue(EnvAccessToken)
	if util.IsBlank(&accessToken) {

		return
	}
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	billPaymentId := getEnvValue(EnvBillPaymentId)
	if util.IsBlank(&billPaymentId) {
		return
	}
	billPaymentAsaas := NewBillPayment(EnvSandbox, accessToken)
	_, _ = billPaymentAsaas.CancelById(ctx, billPaymentId)
	_ = os.Unsetenv(EnvBillPaymentId)
}

func clearFileName() {
	fileName := getEnvValue(EnvFileName)
	if util.IsBlank(&fileName) {
		return
	}
	removeFileTest(fileName)
	_ = os.Unsetenv(EnvFileName)
}

func removeFileTest(fileName string) {
	err := os.Remove(fileName)
	if err != nil {
		logError("error remove file test:", err)
	}
}
