package asaas

import (
	"context"
	"github.com/GabrielHCataldo/go-asaas/internal/test"
	"github.com/GabrielHCataldo/go-asaas/internal/util"
	"github.com/mvrilo/go-cpf"
	"os"
	"testing"
	"time"
)

const EnvAccessToken = "ASAAS_ACCESS_TOKEN"
const EnvAccessTokenSecondary = "ASAAS_ACCESS_TOKEN_SECONDARY"
const EnvWalletIdSecondary = "ASAAS_WALLET_ID_SECONDARY"
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
const EnvCreditBureauReportId = "ASAAS_CREDIT_BUREAU_REPORT_ID"

func init() {
	initFile()
	initImage()
	initFiscalInfo()
}

func TestMain(m *testing.M) {
	code := m.Run()
	logInfo(EnvSandbox, "cleaning all envs")
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

func getEnvValueWithoutLogger(env string) string {
	return os.Getenv(env)
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

func initCustomerDeleted() {
	accessToken := getEnvValue(EnvAccessToken)
	if util.IsBlank(&accessToken) {
		return
	}
	initCustomer()
	customerId := getEnvValue(EnvCustomerId)
	if util.IsBlank(&customerId) {
		return
	}
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	customerAsaas := NewCustomer(EnvSandbox, accessToken)
	resp, err := customerAsaas.DeleteById(ctx, customerId)
	if err != nil || resp.IsFailure() {
		logError("error resp:", resp, "err: ", err)
		return
	}
	setEnv(EnvCustomerDeletedId, resp.Id)
}

func initCreditCardCharge(capture bool, withInstallment bool) {
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
	now := time.Now()
	req := CreateChargeRequest{
		Customer:    customerId,
		BillingType: BillingTypeCreditCard,
		Value:       100,
		DueDate:     NewDate(now.Year(), now.Month(), now.Day(), now.Location()),
		Description: "Cobrança via teste unitário em Golang",
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
		AuthorizeOnly: !capture,
		RemoteIp:      "191.253.125.194",
	}
	if withInstallment {
		req.InstallmentCount = 2
		req.InstallmentValue = req.Value / float64(req.InstallmentCount)
	}
	chargeAsaas := NewCharge(EnvSandbox, accessToken)
	resp, err := chargeAsaas.Create(ctx, req)
	if err != nil || resp.IsFailure() {
		logError("error resp:", resp, "err: ", err)
		return
	}
	success := setEnv(EnvCreditCardChargeId, resp.Id)
	if !success {
		return
	}
	if withInstallment {
		setEnv(EnvChargeInstallmentId, resp.Installment)
	}
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
	now := time.Now()
	chargeAsaas := NewCharge(EnvSandbox, accessToken)
	resp, err := chargeAsaas.Create(ctx, CreateChargeRequest{
		BillingType: BillingTypePix,
		DueDate:     NewDate(now.Year(), now.Month(), now.Day(), now.Location()),
		Value:       100,
		Description: "Cobrança via teste unitário em Golang",
	})
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

func initBankSlipCharge(withInstallment bool) {
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
	now := time.Now()
	req := CreateChargeRequest{
		BillingType: BillingTypeBankSlip,
		DueDate:     NewDate(now.Year(), now.Month(), now.Day(), now.Location()),
		Value:       100,
		Description: "Cobrança via teste unitário em Golang",
	}
	if withInstallment {
		req.InstallmentCount = 2
		req.InstallmentValue = req.Value / float64(req.InstallmentCount)
	}
	chargeAsaas := NewCharge(EnvSandbox, accessToken)
	resp, err := chargeAsaas.Create(ctx, req)
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
	if withInstallment {
		setEnv(EnvChargeInstallmentId, resp.Installment)
	}
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
	now := time.Now()
	chargeAsaas := NewCharge(EnvSandbox, accessToken)
	resp, err := chargeAsaas.Create(ctx, CreateChargeRequest{
		Customer:    customerId,
		BillingType: BillingTypeUndefined,
		Value:       100,
		DueDate:     NewDate(now.Year(), now.Month(), now.Day(), now.Location()),
		Description: "Cobrança via teste unitário em Golang",
	})
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
	initBankSlipCharge(false)
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

func initCreditBureauReport() {
	accessToken := getEnvValue(EnvAccessToken)
	if util.IsBlank(&accessToken) {

		return
	}
	initCustomer()
	customerId := getEnvValue(EnvCustomerId)
	if util.IsBlank(&customerId) {
		return
	}
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nCreditBureau := NewCreditBureau(EnvSandbox, accessToken)
	resp, err := nCreditBureau.GetReport(ctx, GetReportRequest{
		Customer: customerId,
		CpfCnpj:  "",
		State:    "SP",
	})
	if err != nil || resp.IsFailure() {
		logError("error resp:", resp, "err: ", err)
		return
	}
	setEnv(EnvCreditBureauReportId, resp.Id)
}

func initFiscalInfo() {
	accessToken := getEnvValue(EnvAccessToken)
	if util.IsBlank(&accessToken) {

		return
	}
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nFiscalInfo := NewFiscalInfo(EnvSandbox, accessToken)
	resp, err := nFiscalInfo.Save(ctx, SaveFiscalInfoRequest{
		Email:                    "test@gmail.com",
		MunicipalInscription:     Pointer("15.54.74"),
		SimplesNacional:          Pointer(true),
		CulturalProjectsPromoter: nil,
		Cnae:                     Pointer("6201501"),
		SpecialTaxRegime:         nil,
		ServiceListItem:          nil,
		RpsSerie:                 nil,
		RpsNumber:                Pointer(21),
		LoteNumber:               nil,
		Username:                 nil,
		Password:                 Pointer("test"),
		AccessToken:              nil,
		CertificateFile:          nil,
		CertificatePassword:      nil,
	})
	if err != nil || resp.IsFailure() {
		logError("error resp:", resp, "err: ", err)
		return
	}
}

func clearCustomerId() {
	accessToken := getEnvValueWithoutLogger(EnvAccessToken)
	if util.IsBlank(&accessToken) {

		return
	}
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	customerId := getEnvValueWithoutLogger(EnvCustomerId)
	if util.IsBlank(&customerId) {
		return
	}
	customerAsaas := NewCustomer(EnvSandbox, accessToken)
	_, _ = customerAsaas.DeleteById(ctx, customerId)
	_ = os.Unsetenv(EnvCustomerId)
}

func clearCreditCardChargeId() {
	accessToken := getEnvValueWithoutLogger(EnvAccessToken)
	if util.IsBlank(&accessToken) {

		return
	}
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	chargeId := getEnvValueWithoutLogger(EnvCreditCardChargeId)
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
	accessToken := getEnvValueWithoutLogger(EnvAccessToken)
	if util.IsBlank(&accessToken) {

		return
	}
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	chargeId := getEnvValueWithoutLogger(EnvCreditCardChargeId)
	if util.IsBlank(&chargeId) {
		return
	}
	chargeAsaas := NewCharge(EnvSandbox, accessToken)
	_, _ = chargeAsaas.DeleteById(ctx, chargeId)
	_ = os.Unsetenv(EnvPixChargeId)
}

func clearBankSlipChargeId() {
	accessToken := getEnvValueWithoutLogger(EnvAccessToken)
	if util.IsBlank(&accessToken) {

		return
	}
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	chargeId := getEnvValueWithoutLogger(EnvBankSlipChargeId)
	if util.IsBlank(&chargeId) {
		return
	}
	chargeAsaas := NewCharge(EnvSandbox, accessToken)
	_, _ = chargeAsaas.DeleteById(ctx, chargeId)
	_ = os.Unsetenv(EnvBankSlipChargeId)
}

func clearUndefinedChargeId() {
	accessToken := getEnvValueWithoutLogger(EnvAccessToken)
	if util.IsBlank(&accessToken) {

		return
	}
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	chargeId := getEnvValueWithoutLogger(EnvBankSlipChargeId)
	if util.IsBlank(&chargeId) {
		return
	}
	chargeAsaas := NewCharge(EnvSandbox, accessToken)
	_, _ = chargeAsaas.DeleteById(ctx, chargeId)
	_ = os.Unsetenv(EnvUndefinedChargeId)
}

func clearBillPaymentId() {
	accessToken := getEnvValueWithoutLogger(EnvAccessToken)
	if util.IsBlank(&accessToken) {

		return
	}
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	billPaymentId := getEnvValueWithoutLogger(EnvBillPaymentId)
	if util.IsBlank(&billPaymentId) {
		return
	}
	billPaymentAsaas := NewBillPayment(EnvSandbox, accessToken)
	_, _ = billPaymentAsaas.CancelById(ctx, billPaymentId)
	_ = os.Unsetenv(EnvBillPaymentId)
}

func clearFileName() {
	fileName := getEnvValueWithoutLogger(EnvFileName)
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
