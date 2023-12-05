package asaas

import (
	"context"
	"github.com/GabrielHCataldo/go-asaas/internal/util"
	"github.com/mvrilo/go-cpf"
	"io"
	"os"
	"strconv"
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
const EnvInvoiceId = "ASAAS_INVOICE_ID"
const EnvMobilePhoneRechargeId = "ASAAS_MOBILE_PHONE_RECHARGE_ID"
const EnvNegativityId = "ASAAS_NEGATIVITY_ID"
const EnvNotificationId = "ASAAS_NOTIFICATION_ID"
const EnvPaymentLinkId = "ASAAS_PAYMENT_LINK_ID"
const EnvPaymentLinkDeletedId = "ASAAS_PAYMENT_LINK_DELETED_ID"
const EnvPaymentLinkImageId = "ASAAS_PAYMENT_LINK_IMAGE_ID"
const EnvPixKeyId = "ASAAS_PIX_KEY_ID"
const EnvPixTransactionId = "ASAAS_PIX_TRANSACTION_ID"
const EnvSubaccountId = "ASAAS_SUBACCOUNT_ID"
const EnvSubaccountAccessToken = "ASAAS_SUBACCOUNT_ACCESS_TOKEN"
const EnvSubaccountDocumentId = "ASAAS_SUBACCOUNT_DOCUMENT_ID"
const EnvSubaccountDocumentType = "ASAAS_SUBACCOUNT_DOCUMENT_TYPE"
const EnvSubaccountDocumentSentId = "ASAAS_SUBACCOUNT_DOCUMENT_SENT_ID"
const EnvTransferId = "ASAAS_TRANSFER_ID"
const EnvSubscriptionId = "ASAAS_SUBSCRIPTION_ID"

func init() {
	initFile()
	initImage()
	initFiscalInfo()
}

func TestMain(m *testing.M) {
	code := m.Run()
	logInfo(EnvSandbox, "cleaning all envs")
	clearCustomerId()
	clearPixChargeId()
	clearUndefinedChargeId()
	clearBillPaymentId()
	clearFileName()
	clearInvoiceId()
	clearMobilePhoneRechargeId()
	clearNegativityId()
	clearPaymentLinkId()
	clearPixTransactionId()
	clearSubaccountDocumentSentId()
	clearSubaccount()
	clearTransferId()
	clearSubscriptionId()
	logInfo(EnvSandbox, "clean all envs successfully")
	os.Exit(code)
}

func getEnvValue(env string) string {
	v := os.Getenv(env)
	if util.IsBlank(&v) {
		logErrorSkipCaller(4, "error getEnvValue:", env, "is required env")
		return "undefined"
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

func initCustomer(enableNotification bool) {
	clearCustomerId()
	accessToken := getEnvValue(EnvAccessToken)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	customerAsaas := NewCustomer(EnvSandbox, accessToken)
	resp, err := customerAsaas.Create(ctx, CreateCustomerRequest{
		Name:                 "Unit test go",
		CpfCnpj:              cpf.Generate(),
		Email:                "unittestgo@gmail.com",
		MobilePhone:          "47997576130",
		PostalCode:           "89041-001",
		Address:              "Rua General Osório",
		AddressNumber:        "1500",
		NotificationDisabled: !enableNotification,
	})
	if err != nil || resp.IsNoContent() || resp.IsFailure() {
		logError("error resp:", resp, "err: ", err)
		return
	}
	setEnv(EnvCustomerId, resp.Id)
}

func initCustomerDeleted() {
	initCustomer(false)
	accessToken := getEnvValue(EnvAccessToken)
	customerId := getEnvValue(EnvCustomerId)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	customerAsaas := NewCustomer(EnvSandbox, accessToken)
	resp, err := customerAsaas.DeleteById(ctx, customerId)
	if err != nil || resp.IsNoContent() || resp.IsFailure() {
		logError("error resp:", resp, "err: ", err)
		return
	}
	setEnv(EnvCustomerDeletedId, customerId)
}

func initCreditCardCharge(withInstallment, enableCustomerNotification bool) {
	initCustomer(enableCustomerNotification)
	accessToken := getEnvValue(EnvAccessToken)
	customerId := getEnvValue(EnvCustomerId)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	now := DateNow()
	req := CreateChargeRequest{
		Customer:    customerId,
		BillingType: BillingTypeCreditCard,
		Value:       500,
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
		AuthorizeOnly: false,
		RemoteIp:      "191.253.125.194",
	}
	if withInstallment {
		req.InstallmentCount = 2
		req.InstallmentValue = req.Value / float64(req.InstallmentCount)
	}
	chargeAsaas := NewCharge(EnvSandbox, accessToken)
	resp, err := chargeAsaas.Create(ctx, req)
	if err != nil || resp.IsNoContent() || resp.IsFailure() {
		logError("error resp:", resp, "err: ", err)
		return
	}
	setEnv(EnvCreditCardChargeId, resp.Id)
	setEnv(EnvChargeInstallmentId, resp.Installment)
}

func initPixCharge() {
	initCustomer(false)
	clearPixChargeId()
	accessToken := getEnvValue(EnvAccessToken)
	customerId := getEnvValue(EnvCustomerId)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	now := DateNow()
	chargeAsaas := NewCharge(EnvSandbox, accessToken)
	resp, err := chargeAsaas.Create(ctx, CreateChargeRequest{
		Customer:    customerId,
		BillingType: BillingTypePix,
		DueDate:     NewDate(now.Year(), now.Month(), now.Day(), now.Location()),
		Value:       5,
		Description: "Cobrança via teste unitário em Golang",
	})
	if err != nil || resp.IsNoContent() || resp.IsFailure() {
		logError("error resp:", resp, "err: ", err)
		return
	}
	setEnv(EnvPixChargeId, resp.Id)
	pixQrCodeResp, err := chargeAsaas.GetPixQrCodeById(ctx, resp.Id)
	if err != nil || pixQrCodeResp.IsNoContent() || pixQrCodeResp.IsFailure() {
		logError("error resp:", pixQrCodeResp, "err: ", err)
		return
	}
	setEnv(EnvChargePixQrCodePayload, pixQrCodeResp.Payload)
}

func initBankSlipCharge(withInstallment bool) {
	initCustomer(false)
	clearBankSlipChargeId()
	accessToken := getEnvValue(EnvAccessToken)
	customerId := getEnvValue(EnvCustomerId)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	now := DateNow()
	req := CreateChargeRequest{
		Customer:    customerId,
		BillingType: BillingTypeBankSlip,
		DueDate:     NewDate(now.Year(), now.Month(), now.Day(), now.Location()),
		Value:       5,
		Description: "Cobrança via teste unitário em Golang",
	}
	if withInstallment {
		req.InstallmentCount = 2
		req.InstallmentValue = req.Value / float64(req.InstallmentCount)
	}
	chargeAsaas := NewCharge(EnvSandbox, accessToken)
	resp, err := chargeAsaas.Create(ctx, req)
	if err != nil || resp.IsNoContent() || resp.IsFailure() {
		logError("error resp:", resp, "err: ", err)
		return
	}
	setEnv(EnvBankSlipChargeId, resp.Id)
	identificationFieldResp, err := chargeAsaas.GetIdentificationFieldById(ctx, resp.Id)
	if err != nil || identificationFieldResp.IsNoContent() || identificationFieldResp.IsFailure() {
		logError("error resp:", identificationFieldResp, "err: ", err)
		return
	}
	setEnv(EnvChargeIdentificationField, identificationFieldResp.IdentificationField)
	setEnv(EnvChargeInstallmentId, resp.Installment)
}

func initUndefinedCharge() {
	initCustomer(false)
	clearUndefinedChargeId()
	accessToken := getEnvValue(EnvAccessToken)
	customerId := getEnvValue(EnvCustomerId)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	now := DateNow()
	chargeAsaas := NewCharge(EnvSandbox, accessToken)
	resp, err := chargeAsaas.Create(ctx, CreateChargeRequest{
		Customer:    customerId,
		BillingType: BillingTypeUndefined,
		Value:       5,
		DueDate:     NewDate(now.Year(), now.Month(), now.Day(), now.Location()),
		Description: "Cobrança via teste unitário em Golang",
	})
	if err != nil || resp.IsNoContent() || resp.IsFailure() {
		logError("error resp:", resp, "err: ", err)
		return
	}
	setEnv(EnvUndefinedChargeId, resp.Id)
}

func initChargeDeleted() {
	initUndefinedCharge()
	accessToken := getEnvValue(EnvAccessToken)
	chargeId := getEnvValue(EnvUndefinedChargeId)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	chargeAsaas := NewCharge(EnvSandbox, accessToken)
	resp, err := chargeAsaas.DeleteById(ctx, chargeId)
	if err != nil || resp.IsNoContent() || resp.IsFailure() {
		logError("error resp:", resp, "err: ", err)
		return
	}
	setEnv(EnvChargeDeletedId, chargeId)
}

func initChargeReceivedInCash() {
	initUndefinedCharge()
	accessToken := getEnvValue(EnvAccessToken)
	chargeId := getEnvValue(EnvUndefinedChargeId)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	now := DateNow()
	chargeAsaas := NewCharge(EnvSandbox, accessToken)
	resp, err := chargeAsaas.ReceiveInCashById(ctx, chargeId, ChargeReceiveInCashRequest{
		PaymentDate: NewDate(now.Year(), now.Month(), now.Day(), now.Location()),
		Value:       5,
	})
	if err != nil || resp.IsNoContent() || resp.IsFailure() {
		logError("error resp:", resp, "err: ", err)
		return
	}
	setEnv(EnvChargeReceivedInCashId, resp.Id)
}

func initChargeDocumentId() {
	initUndefinedCharge()
	accessToken := getEnvValue(EnvAccessToken)
	chargeId := getEnvValue(EnvUndefinedChargeId)
	f, err := os.Open(getEnvValue(EnvFileName))
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nCharge := NewCharge(EnvSandbox, accessToken)
	resp, err := nCharge.UploadDocumentById(ctx, chargeId, UploadChargeDocumentRequest{
		Type: DocumentTypeDocument,
		File: f,
	})
	if err != nil || resp.IsNoContent() || resp.IsFailure() {
		logError("error resp:", resp, "err: ", err)
		return
	}
	setEnv(EnvChargeDocumentId, resp.Id)
}

func initAnticipation() {
	initCreditCardCharge(false, false)
	accessToken := getEnvValue(EnvAccessToken)
	chargeId := getEnvValue(EnvCreditCardChargeId)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	anticipationAsaas := NewAnticipation(EnvSandbox, accessToken)
	resp, err := anticipationAsaas.Request(ctx, AnticipationRequest{
		Payment: chargeId,
	})
	if err != nil || resp.IsNoContent() || resp.IsFailure() {
		logError("error resp:", resp, "err: ", err)
		return
	}
	setEnv(EnvAnticipationId, resp.Id)
}

func initBillPayment() {
	initBankSlipCharge(false)
	clearBillPaymentId()
	accessToken := getEnvValue(EnvAccessToken)
	identificationField := getEnvValue(EnvChargeIdentificationField)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	now := DateNow()
	billPaymentAsaas := NewBillPayment(EnvSandbox, accessToken)
	resp, err := billPaymentAsaas.Create(ctx, CreateBillPaymentRequest{
		IdentificationField: identificationField,
		ScheduleDate:        NewDate(now.Year(), now.Month(), now.Day()+1, now.Location()),
	})
	if err != nil || resp.IsNoContent() || resp.IsFailure() {
		logError("error create bill payment resp:", resp, "err: ", err)
		return
	}
	setEnv(EnvBillPaymentId, resp.Id)
}

func initFile() {
	f, err := getTestFile()
	if err != nil {
		logError("error init file GetSimpleFile:", err)
		return
	}
	setEnv(EnvFileName, f.Name())
}

func initImage() {
	f, err := getTestImage()
	if err != nil {
		logError("error init image GetSimpleImage:", err)
		return
	}
	setEnv(EnvImageName, f.Name())
}

func initCreditBureauReport() {
	initCustomer(false)
	accessToken := getEnvValue(EnvAccessToken)
	customerId := getEnvValue(EnvCustomerId)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nCreditBureau := NewCreditBureau(EnvSandbox, accessToken)
	resp, err := nCreditBureau.GetReport(ctx, GetReportRequest{
		Customer: customerId,
		CpfCnpj:  "",
		State:    "SP",
	})
	if err != nil || resp.IsNoContent() || resp.IsFailure() {
		logError("error resp:", resp, "err: ", err)
		return
	}
	setEnv(EnvCreditBureauReportId, resp.Id)
}

func initFiscalInfo() {
	accessToken := getEnvValue(EnvAccessToken)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nFiscalInfo := NewFiscalInfo(EnvSandbox, accessToken)
	resp, err := nFiscalInfo.Save(ctx, SaveFiscalInfoRequest{
		Email:                    "test@gmail.com",
		MunicipalInscription:     Pointer("15.54.74"),
		SimplesNacional:          Pointer(true),
		CulturalProjectsPromoter: nil,
		Cnae:                     Pointer("6201501"),
		SpecialTaxRegime:         Pointer("test"),
		ServiceListItem:          nil,
		RpsSerie:                 Pointer("E"),
		RpsNumber:                Pointer(21),
		LoteNumber:               Pointer(21),
		Username:                 nil,
		Password:                 Pointer("test"),
		AccessToken:              nil,
		CertificateFile:          nil,
		CertificatePassword:      nil,
	})
	if err != nil || resp.IsNoContent() || resp.IsFailure() {
		logError("error resp:", resp, "err: ", err)
		return
	}
}

func initInvoice() {
	initCreditCardCharge(false, false)
	clearInvoiceId()
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
	if err != nil || resp.IsNoContent() || resp.IsFailure() {
		logError("error resp:", resp, "err: ", err)
		return
	}
	setEnv(EnvInvoiceId, resp.Id)
}

func initMobilePhoneRecharge() {
	clearMobilePhoneRechargeId()
	accessToken := getEnvValue(EnvAccessToken)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nMobilePhone := NewMobilePhone(EnvSandbox, accessToken)
	resp, err := nMobilePhone.Recharge(ctx, MobilePhoneRechargeRequest{
		PhoneNumber: "47997576130",
		Value:       20,
	})
	if err != nil || resp.IsNoContent() || resp.IsFailure() {
		logError("error resp:", resp, "err: ", err)
		return
	}
	setEnv(EnvMobilePhoneRechargeId, resp.Id)
}

func initNegativity() {
	initBankSlipCharge(false)
	clearNegativityId()
	accessToken := getEnvValue(EnvAccessToken)
	chargeId := getEnvValue(EnvBankSlipChargeId)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nNegativity := NewNegativity(EnvSandbox, accessToken)
	resp, err := nNegativity.Create(ctx, CreateNegativityRequest{
		Payment:               chargeId,
		Type:                  NegativityTypeCreditBureau,
		Description:           "Unit test golang",
		CustomerName:          "Unit test golang",
		CustomerCpfCnpj:       "24971563792",
		CustomerPrimaryPhone:  "47999376637",
		CustomerPostalCode:    "01310-000",
		CustomerAddress:       "Av. Paulista",
		CustomerAddressNumber: "150",
		CustomerProvince:      "Centro",
	})
	if err != nil || resp.IsNoContent() || resp.IsFailure() {
		logError("error resp:", resp, "err: ", err)
		return
	}
	setEnv(EnvNegativityId, resp.Id)
}

func initNotification() {
	initCreditCardCharge(false, true)
	accessToken := getEnvValue(EnvAccessToken)
	customerId := getEnvValue(EnvCustomerId)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nNotification := NewNotification(EnvSandbox, accessToken)
	resp, err := nNotification.GetAllByCustomer(ctx, customerId)
	if err != nil || resp.IsNoContent() || resp.IsFailure() {
		logError("error resp:", resp, "err: ", err)
		return
	}
	setEnv(EnvNotificationId, resp.Data[0].Id)
}

func initPaymentLink() {
	clearPaymentLinkId()
	accessToken := getEnvValue(EnvAccessToken)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nPaymentLink := NewPaymentLink(EnvSandbox, accessToken)
	resp, err := nPaymentLink.Create(ctx, CreatePaymentLinkRequest{
		Name:        "Unit test go",
		BillingType: BillingTypeUndefined,
		ChargeType:  ChargeTypeDetached,
	})
	if err != nil || resp.IsNoContent() || resp.IsFailure() {
		logError("error resp:", resp, "err: ", err)
		return
	}
	setEnv(EnvPaymentLinkId, resp.Id)
}

func initPaymentLinkDeleted() {
	initPaymentLink()
	accessToken := getEnvValue(EnvAccessToken)
	paymentLinkId := getEnvValue(EnvPaymentLinkId)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nPaymentLink := NewPaymentLink(EnvSandbox, accessToken)
	resp, err := nPaymentLink.DeleteById(ctx, paymentLinkId)
	if err != nil || resp.IsNoContent() || resp.IsFailure() {
		logError("error resp:", resp, "err: ", err)
		return
	}
	setEnv(EnvPaymentLinkDeletedId, paymentLinkId)
}

func initPaymentLinkImage() {
	initPaymentLink()
	accessToken := getEnvValue(EnvAccessToken)
	paymentLinkId := getEnvValue(EnvPaymentLinkId)
	f, _ := os.Open(getEnvValue(EnvImageName))
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nPaymentLink := NewPaymentLink(EnvSandbox, accessToken)
	resp, err := nPaymentLink.SendImageById(ctx, paymentLinkId, SendImagePaymentLinksRequest{
		Main:  false,
		Image: f,
	})
	if err != nil || resp.IsNoContent() || resp.IsFailure() {
		logError("error resp:", resp, "err: ", err)
		return
	}
	setEnv(EnvPaymentLinkImageId, resp.Id)
}

func initPixKey() {
	clearPixKeyId()
	accessToken := getEnvValue(EnvAccessToken)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nPix := NewPix(EnvSandbox, accessToken)
	resp, err := nPix.CreateKey(ctx)
	if err != nil || resp.IsNoContent() || resp.IsFailure() {
		logError("error resp:", resp, "err: ", err)
		return
	}
	setEnv(EnvPixKeyId, resp.Id)
}

func initPixTransaction() {
	initPixCharge()
	clearPixTransactionId()
	accessToken := getEnvValue(EnvAccessToken)
	pixQrCodePayload := getEnvValue(EnvChargePixQrCodePayload)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nPix := NewPix(EnvSandbox, accessToken)
	resp, err := nPix.PayQrCode(ctx, PayPixQrCodeRequest{
		QrCode: PixQrCodeRequest{
			Payload:     pixQrCodePayload,
			ChangeValue: 0,
		},
		Value:        5,
		Description:  "",
		ScheduleDate: Date{},
	})
	if err != nil || resp.IsNoContent() || resp.IsFailure() {
		logError("error resp:", resp, "err: ", err)
		return
	}
	setEnv(EnvPixTransactionId, resp.Id)
}

func initSubaccount() {
	clearSubaccount()
	accessToken := getEnvValue(EnvAccessToken)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nSubaccount := NewSubaccount(EnvSandbox, accessToken)
	resp, err := nSubaccount.Create(ctx, CreateSubaccountRequest{
		Name:          "Unit test go",
		Email:         util.GenerateEmail(),
		CpfCnpj:       "69257172000141",
		CompanyType:   CompanyTypeLimited,
		MobilePhone:   util.GenerateMobilePhone(),
		Address:       "Rua Maria de Souza Maba",
		AddressNumber: "123",
		Province:      "Fortaleza",
		PostalCode:    "89056-220",
	})
	if err != nil || resp.IsNoContent() || resp.IsFailure() {
		logError("error resp:", resp, "err: ", err)
		return
	}
	setEnv(EnvSubaccountId, resp.Id)
	setEnv(EnvSubaccountAccessToken, resp.ApiKey)
}

func initSubaccountDocument() {
	initSubaccount()
	subaccountAccessToken := getEnvValue(EnvSubaccountAccessToken)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nSubaccount := NewSubaccount(EnvSandbox, subaccountAccessToken)
	resp, err := nSubaccount.GetPendingDocuments(ctx)
	if err != nil || resp.IsNoContent() || resp.IsFailure() {
		logError("error resp:", resp, "err: ", err)
		return
	}
	var subaccountDocument SubaccountDocumentResponse
	for _, documentToSent := range resp.Data {
		if documentToSent.Status == SubaccountDocumentStatusNotSent &&
			documentToSent.Type != SubaccountDocumentTypeIdentification {
			subaccountDocument = documentToSent
		}
	}
	if util.IsBlank(&subaccountDocument.Id) {
		logError("subaccountDocument not found")
		return
	}
	setEnv(EnvSubaccountDocumentId, subaccountDocument.Id)
	setEnv(EnvSubaccountDocumentType, string(subaccountDocument.Type))
}

func initSubaccountDocumentSent() {
	initSubaccountDocument()
	clearSubaccountDocumentSentId()
	subaccountAccessToken := getEnvValue(EnvSubaccountAccessToken)
	subaccountDocumentId := getEnvValue(EnvSubaccountDocumentId)
	subaccountDocumentType := getEnvValue(EnvSubaccountDocumentType)
	f, _ := os.Open(getEnvValue(EnvImageName))
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nSubaccount := NewSubaccount(EnvSandbox, subaccountAccessToken)
	resp, err := nSubaccount.SendWhiteLabelDocument(ctx, subaccountDocumentId, SendWhiteLabelDocumentRequest{
		Type:         SubaccountDocumentType(subaccountDocumentType),
		DocumentFile: f,
	})
	if err != nil || resp.IsNoContent() || resp.IsFailure() {
		logError("error resp:", resp, "err: ", err)
		return
	}
	setEnv(EnvSubaccountDocumentSentId, resp.Id)
}

func initTransfer() {
	clearTransferId()
	accessToken := getEnvValue(EnvAccessToken)
	walletId := getEnvValue(EnvWalletIdSecondary)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nTransfer := NewTransfer(EnvSandbox, accessToken)
	resp, err := nTransfer.TransferToAsaas(ctx, TransferToAssasRequest{
		Value:    10,
		WalletId: walletId,
	})
	if err != nil || resp.IsNoContent() || resp.IsFailure() {
		logError("error resp:", resp, "err: ", err)
		return
	}
	setEnv(EnvTransferId, resp.Id)
}

func initWebhook() {
	accessToken := getEnvValue(EnvAccessTokenSecondary)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nWebhook := NewWebhook(EnvSandbox, accessToken)
	resp, err := nWebhook.SaveSetting(ctx, WebhookTypePayment, SaveWebhookSettingRequest{
		Url:         "https://test.com",
		Email:       "test@gmail.com",
		ApiVersion:  "3",
		Enabled:     Pointer(false),
		Interrupted: Pointer(false),
		AuthToken:   "",
	})
	if err != nil || resp.IsNoContent() || resp.IsFailure() {
		logError("error resp:", resp, "err: ", err)
		return
	}
}

func initSubscription() {
	initCustomer(false)
	clearSubscriptionId()
	accessToken := getEnvValue(EnvAccessToken)
	customerId := getEnvValue(EnvCustomerId)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	now := DateNow()
	nSubscription := NewSubscription(EnvSandbox, accessToken)
	resp, err := nSubscription.Create(ctx, CreateSubscriptionRequest{
		Customer:    customerId,
		BillingType: BillingTypeBankSlip,
		Value:       5,
		NextDueDate: NewDate(now.Year(), now.Month()+1, now.Day(), now.Location()),
		Cycle:       SubscriptionCycleMonthly,
		Description: "Unit test go",
	})
	if err != nil || resp.IsNoContent() || resp.IsFailure() {
		logError("error resp:", resp, "err: ", err)
		return
	}
	setEnv(EnvSubscriptionId, resp.Id)
}

func initSubscriptionInvoiceSetting() {
	accessToken := getEnvValue(EnvAccessToken)
	subscriptionId := getEnvValue(EnvSubscriptionId)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nSubscription := NewSubscription(EnvSandbox, accessToken)
	resp, err := nSubscription.CreateInvoiceSettingById(ctx, subscriptionId, CreateInvoiceSettingRequest{
		MunicipalServiceCode: "123",
		MunicipalServiceName: "Unit test go",
		EffectiveDatePeriod:  EffectiveDatePeriodOnNextMonth,
		Observations:         "Unit test go",
	})
	if err != nil || resp.IsNoContent() || resp.IsFailure() {
		logError("error resp:", resp, "err: ", err)
		return
	}
}

func clearCustomerId() {
	accessToken := getEnvValueWithoutLogger(EnvAccessToken)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	customerId := getEnvValueWithoutLogger(EnvCustomerId)
	if util.IsBlank(&customerId) {
		return
	}
	customerAsaas := NewCustomer(EnvSandbox, accessToken)
	_, _ = customerAsaas.DeleteById(ctx, customerId)
	_ = os.Unsetenv(EnvCustomerId)
}

func clearPixChargeId() {
	accessToken := getEnvValueWithoutLogger(EnvAccessToken)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	chargeId := getEnvValueWithoutLogger(EnvPixChargeId)
	if util.IsBlank(&chargeId) {
		return
	}
	chargeAsaas := NewCharge(EnvSandbox, accessToken)
	_, _ = chargeAsaas.DeleteById(ctx, chargeId)
	_ = os.Unsetenv(EnvPixChargeId)
}

func clearBankSlipChargeId() {
	accessToken := getEnvValueWithoutLogger(EnvAccessToken)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
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
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
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
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
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

func clearInvoiceId() {
	accessToken := getEnvValueWithoutLogger(EnvAccessToken)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	invoiceId := getEnvValueWithoutLogger(EnvInvoiceId)
	if util.IsBlank(&invoiceId) {
		return
	}
	invoiceAsaas := NewInvoice(EnvSandbox, accessToken)
	_, _ = invoiceAsaas.CancelById(ctx, invoiceId)
	_ = os.Unsetenv(EnvInvoiceId)
}

func clearMobilePhoneRechargeId() {
	accessToken := getEnvValueWithoutLogger(EnvAccessToken)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	rechargeId := getEnvValueWithoutLogger(EnvMobilePhoneRechargeId)
	if util.IsBlank(&rechargeId) {
		return
	}
	mobilePhoneAsaas := NewMobilePhone(EnvSandbox, accessToken)
	_, _ = mobilePhoneAsaas.CancelRechargeById(ctx, rechargeId)
	_ = os.Unsetenv(EnvMobilePhoneRechargeId)
}

func clearNegativityId() {
	accessToken := getEnvValueWithoutLogger(EnvAccessToken)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	negativityId := getEnvValueWithoutLogger(EnvNegativityId)
	if util.IsBlank(&negativityId) {
		return
	}
	negativityAsaas := NewNegativity(EnvSandbox, accessToken)
	_, _ = negativityAsaas.CancelById(ctx, negativityId)
	_ = os.Unsetenv(EnvNegativityId)
}

func clearPaymentLinkId() {
	accessToken := getEnvValueWithoutLogger(EnvAccessToken)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	paymentLinkId := getEnvValueWithoutLogger(EnvPaymentLinkId)
	if util.IsBlank(&paymentLinkId) {
		return
	}
	paymentLinkAsaas := NewPaymentLink(EnvSandbox, accessToken)
	_, _ = paymentLinkAsaas.DeleteById(ctx, paymentLinkId)
	_ = os.Unsetenv(EnvPaymentLinkId)
}

func clearPixTransactionId() {
	accessToken := getEnvValueWithoutLogger(EnvAccessToken)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	pixTransactionId := getEnvValueWithoutLogger(EnvPixTransactionId)
	if util.IsBlank(&pixTransactionId) {
		return
	}
	pixAsaas := NewPix(EnvSandbox, accessToken)
	_, _ = pixAsaas.CancelTransactionById(ctx, pixTransactionId)
	_ = os.Unsetenv(EnvPixTransactionId)
}

func clearPixKeyId() {
	accessToken := getEnvValueWithoutLogger(EnvAccessToken)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	pixKeyId := getEnvValueWithoutLogger(EnvPixKeyId)
	if util.IsBlank(&pixKeyId) {
		return
	}
	pixAsaas := NewPix(EnvSandbox, accessToken)
	_, _ = pixAsaas.DeleteKeyById(ctx, pixKeyId)
	_ = os.Unsetenv(EnvPixKeyId)
}

func clearSubaccount() {
	subaccountAccessToken := getEnvValueWithoutLogger(EnvSubaccountAccessToken)
	if util.IsBlank(&subaccountAccessToken) {
		return
	}
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	subaccountId := getEnvValueWithoutLogger(EnvSubaccountId)
	if util.IsBlank(&subaccountId) {
		return
	}
	accountAsaas := NewAccount(EnvSandbox, subaccountAccessToken)
	_, _ = accountAsaas.DeleteWhiteLabelSubaccount(ctx, DeleteWhiteLabelSubaccountRequest{
		RemoveReason: "Unit test go",
	})
	_ = os.Unsetenv(EnvSubaccountId)
}

func clearSubaccountDocumentSentId() {
	subaccountAccessToken := getEnvValueWithoutLogger(EnvSubaccountAccessToken)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	documentSentId := getEnvValueWithoutLogger(EnvSubaccountDocumentSentId)
	if util.IsBlank(&documentSentId) {
		return
	}
	subaccountAsaas := NewSubaccount(EnvSandbox, subaccountAccessToken)
	_, _ = subaccountAsaas.DeleteWhiteLabelDocumentSentById(ctx, documentSentId)
	_ = os.Unsetenv(EnvSubaccountDocumentSentId)
}

func clearTransferId() {
	accessToken := getEnvValueWithoutLogger(EnvAccessToken)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	transferId := getEnvValueWithoutLogger(EnvTransferId)
	if util.IsBlank(&transferId) {
		return
	}
	transferAsaas := NewTransfer(EnvSandbox, accessToken)
	_, _ = transferAsaas.CancelById(ctx, transferId)
	_ = os.Unsetenv(EnvTransferId)
}

func clearSubscriptionId() {
	accessToken := getEnvValueWithoutLogger(EnvAccessToken)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	subscriptionId := getEnvValueWithoutLogger(EnvSubscriptionId)
	if util.IsBlank(&subscriptionId) {
		return
	}
	subscriptionAsaas := NewSubscription(EnvSandbox, accessToken)
	_, _ = subscriptionAsaas.DeleteById(ctx, subscriptionId)
	_ = os.Unsetenv(EnvSubscriptionId)
}

func getTestFile() (*os.File, error) {
	randomKey := strconv.FormatInt(time.Now().Unix()+int64(time.Now().Nanosecond()), 10)
	nameFile := "test " + randomKey + ".txt"
	f, err := os.Create(nameFile)
	if err != nil {
		return nil, err
	}
	_, err = io.WriteString(f, "unit test golang")
	if err != nil {
		return nil, err
	}
	err = f.Close()
	if err != nil {
		return nil, err
	}
	return os.Open(nameFile)
}

func getTestImage() (*os.File, error) {
	return os.Open("../gopher-asaas.png")
}

func removeFileTest(fileName string) {
	err := os.Remove(fileName)
	if err != nil {
		logError("error remove file test:", err)
	}
}
