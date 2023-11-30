package asaas

import "github.com/GabrielHCataldo/go-asaas/internal/util"

type DeleteResponse struct {
	Id      string          `json:"id,omitempty"`
	Deleted bool            `json:"deleted,omitempty"`
	Errors  []ErrorResponse `json:"errors,omitempty"`
}

type response interface {
	IsSuccess() bool
	IsFailure() bool
	IsNoContent() bool
}

func (p PaymentLinkImageResponse) IsSuccess() bool {
	return len(p.Errors) == 0 && util.IsNotBlank(&p.Id)
}

func (p PaymentLinkImageResponse) IsFailure() bool {
	return !p.IsSuccess()
}

func (p PaymentLinkImageResponse) IsNoContent() bool {
	return len(p.Errors) == 0 && util.IsBlank(&p.Id)
}

func (p PaymentLinkResponse) IsSuccess() bool {
	return len(p.Errors) == 0 && util.IsNotBlank(&p.Id)
}

func (p PaymentLinkResponse) IsFailure() bool {
	return !p.IsSuccess()
}

func (p PaymentLinkResponse) IsNoContent() bool {
	return len(p.Errors) == 0 && util.IsBlank(&p.Id)
}

func (s SubaccountDocumentsResponse) IsSuccess() bool {
	return len(s.Errors) == 0 && len(s.Data) > 0
}

func (s SubaccountDocumentsResponse) IsFailure() bool {
	return !s.IsSuccess()
}

func (s SubaccountDocumentsResponse) IsNoContent() bool {
	return len(s.Errors) == 0 && len(s.Data) == 0
}

func (w WebhookResponse) IsSuccess() bool {
	return len(w.Errors) == 0 && util.IsNotBlank(&w.Url)
}

func (w WebhookResponse) IsFailure() bool {
	return !w.IsSuccess()
}

func (w WebhookResponse) IsNoContent() bool {
	return len(w.Errors) == 0 && util.IsBlank(&w.Url)
}

func (s SplitStatisticResponse) IsSuccess() bool {
	return len(s.Errors) == 0
}

func (s SplitStatisticResponse) IsFailure() bool {
	return !s.IsSuccess()
}

func (s SplitStatisticResponse) IsNoContent() bool {
	return false
}

func (s PaymentStatisticResponse) IsSuccess() bool {
	return len(s.Errors) == 0
}

func (s PaymentStatisticResponse) IsFailure() bool {
	return !s.IsSuccess()
}

func (s PaymentStatisticResponse) IsNoContent() bool {
	return false
}

func (a AccountRegistrationStatusResponse) IsSuccess() bool {
	return len(a.Errors) == 0 && util.IsNotBlank(&a.Id)
}

func (a AccountRegistrationStatusResponse) IsFailure() bool {
	return !a.IsSuccess()
}

func (a AccountRegistrationStatusResponse) IsNoContent() bool {
	return len(a.Errors) == 0 && util.IsBlank(&a.Id)
}

func (a AccountBalanceResponse) IsSuccess() bool {
	return len(a.Errors) == 0
}

func (a AccountBalanceResponse) IsFailure() bool {
	return !a.IsSuccess()
}

func (a AccountBalanceResponse) IsNoContent() bool {
	return false
}

func (a AccountFeesResponse) IsSuccess() bool {
	return len(a.Errors) == 0
}

func (a AccountFeesResponse) IsFailure() bool {
	return !a.IsSuccess()
}

func (a AccountFeesResponse) IsNoContent() bool {
	return false
}

func (a AccountBankInfoResponse) IsSuccess() bool {
	return len(a.Errors) == 0
}

func (a AccountBankInfoResponse) IsFailure() bool {
	return !a.IsSuccess()
}

func (a AccountBankInfoResponse) IsNoContent() bool {
	return false
}

func (d DeleteWhiteLabelSubaccountResponse) IsSuccess() bool {
	return len(d.Errors) == 0
}

func (d DeleteWhiteLabelSubaccountResponse) IsFailure() bool {
	return !d.IsSuccess()
}

func (d DeleteWhiteLabelSubaccountResponse) IsNoContent() bool {
	return false
}

func (i InvoiceCustomizationResponse) IsSuccess() bool {
	return len(i.Errors) == 0 && util.IsNotBlank(&i.LogoBackgroundColor)
}

func (i InvoiceCustomizationResponse) IsFailure() bool {
	return !i.IsSuccess()
}

func (i InvoiceCustomizationResponse) IsNoContent() bool {
	return len(i.Errors) == 0 && util.IsBlank(&i.LogoBackgroundColor)
}

func (a AccountResponse) IsSuccess() bool {
	return len(a.Errors) == 0 && util.IsNotBlank(&a.CpfCnpj)
}

func (a AccountResponse) IsFailure() bool {
	return !a.IsSuccess()
}

func (a AccountResponse) IsNoContent() bool {
	return len(a.Errors) == 0 && util.IsBlank(&a.CpfCnpj)
}

func (c CreditBureauReportResponse) IsSuccess() bool {
	return len(c.Errors) == 0 && util.IsNotBlank(&c.Id)
}

func (c CreditBureauReportResponse) IsFailure() bool {
	return !c.IsSuccess()
}

func (c CreditBureauReportResponse) IsNoContent() bool {
	return len(c.Errors) == 0 && util.IsBlank(&c.Id)
}

func (m MobilePhoneProviderResponse) IsSuccess() bool {
	return len(m.Errors) == 0 && len(m.Values) > 0
}

func (m MobilePhoneProviderResponse) IsFailure() bool {
	return !m.IsSuccess()
}

func (m MobilePhoneProviderResponse) IsNoContent() bool {
	return len(m.Errors) == 0 && len(m.Values) == 0
}

func (m MobilePhoneRechargeResponse) IsSuccess() bool {
	return len(m.Errors) == 0 && util.IsNotBlank(&m.Id)
}

func (m MobilePhoneRechargeResponse) IsFailure() bool {
	return !m.IsSuccess()
}

func (m MobilePhoneRechargeResponse) IsNoContent() bool {
	return len(m.Errors) == 0 && util.IsBlank(&m.Id)
}

func (b BillPaymentSimulateResponse) IsSuccess() bool {
	return len(b.Errors) == 0 && util.IsNotBlank(&b.BankSlipInfo.IdentificationField)
}

func (b BillPaymentSimulateResponse) IsFailure() bool {
	return !b.IsSuccess()
}

func (b BillPaymentSimulateResponse) IsNoContent() bool {
	return len(b.Errors) == 0 && util.IsBlank(&b.BankSlipInfo.IdentificationField)
}

func (b BillPaymentResponse) IsSuccess() bool {
	return len(b.Errors) == 0 && util.IsNotBlank(&b.Id)
}

func (b BillPaymentResponse) IsFailure() bool {
	return !b.IsSuccess()
}

func (b BillPaymentResponse) IsNoContent() bool {
	return len(b.Errors) == 0 && util.IsBlank(&b.Id)
}

func (n NegativitySimulateResponse) IsSuccess() bool {
	return len(n.Errors) == 0 && util.IsNotBlank(&n.Payment)
}

func (n NegativitySimulateResponse) IsFailure() bool {
	return !n.IsSuccess()
}

func (n NegativitySimulateResponse) IsNoContent() bool {
	return len(n.Errors) == 0 && util.IsBlank(&n.Payment)
}

func (n NegativityResponse) IsSuccess() bool {
	return len(n.Errors) == 0 && util.IsNotBlank(&n.Id)
}

func (n NegativityResponse) IsFailure() bool {
	return !n.IsSuccess()
}

func (n NegativityResponse) IsNoContent() bool {
	return len(n.Errors) == 0 && util.IsBlank(&n.Id)
}

func (a AnticipationResponse) IsSuccess() bool {
	return len(a.Errors) == 0 && util.IsNotBlank(&a.Id)
}

func (a AnticipationResponse) IsFailure() bool {
	return !a.IsSuccess()
}

func (a AnticipationResponse) IsNoContent() bool {
	return len(a.Errors) == 0 && util.IsBlank(&a.Id)
}

func (a AnticipationSimulateResponse) IsSuccess() bool {
	return len(a.Errors) == 0 && a.Value > 0
}

func (a AnticipationSimulateResponse) IsFailure() bool {
	return !a.IsSuccess()
}

func (a AnticipationSimulateResponse) IsNoContent() bool {
	return len(a.Errors) == 0 && a.Value == 0
}

func (a AgreementSignResponse) IsSuccess() bool {
	return len(a.Errors) == 0
}

func (a AgreementSignResponse) IsFailure() bool {
	return !a.IsSuccess()
}

func (a AgreementSignResponse) IsNoContent() bool {
	return false
}

func (a AnticipationLimitsResponse) IsSuccess() bool {
	return len(a.Errors) == 0
}

func (a AnticipationLimitsResponse) IsFailure() bool {
	return !a.IsSuccess()
}

func (a AnticipationLimitsResponse) IsNoContent() bool {
	return false
}

func (t TransferResponse) IsSuccess() bool {
	return len(t.Errors) == 0 && util.IsNotBlank(&t.Id)
}

func (t TransferResponse) IsFailure() bool {
	return !t.IsSuccess()
}

func (t TransferResponse) IsNoContent() bool {
	return len(t.Errors) == 0 && util.IsBlank(&t.Id)
}

func (p PixTransactionResponse) IsSuccess() bool {
	return len(p.Errors) == 0 && util.IsNotBlank(&p.Id)
}

func (p PixTransactionResponse) IsFailure() bool {
	return !p.IsSuccess()
}

func (p PixTransactionResponse) IsNoContent() bool {
	return len(p.Errors) == 0 && util.IsBlank(&p.Id)
}

func (f FileTextPlainResponse) IsSuccess() bool {
	s := f.String()
	return util.IsNotBlank(&s)
}

func (f FileTextPlainResponse) IsFailure() bool {
	return !f.IsSuccess()
}

func (f FileTextPlainResponse) IsNoContent() bool {
	s := f.String()
	return util.IsBlank(&s)
}

func (p PixCancelTransactionResponse) IsSuccess() bool {
	return len(p.Errors) == 0 && util.IsNotBlank(&p.Id)
}

func (p PixCancelTransactionResponse) IsFailure() bool {
	return !p.IsSuccess()
}

func (p PixCancelTransactionResponse) IsNoContent() bool {
	return len(p.Errors) == 0 && util.IsBlank(&p.Id)
}

func (d DecodePixQrCodeResponse) IsSuccess() bool {
	return len(d.Errors) == 0 && util.IsNotBlank(&d.Payload)
}

func (d DecodePixQrCodeResponse) IsFailure() bool {
	return !d.IsSuccess()
}

func (d DecodePixQrCodeResponse) IsNoContent() bool {
	return len(d.Errors) == 0 && util.IsBlank(&d.Payload)
}

func (p PixKeyResponse) IsSuccess() bool {
	return len(p.Errors) == 0 && util.IsNotBlank(&p.Id)
}

func (p PixKeyResponse) IsFailure() bool {
	return !p.IsSuccess()
}

func (p PixKeyResponse) IsNoContent() bool {
	return len(p.Errors) == 0 && util.IsBlank(&p.Id)
}

func (q QrCodeResponse) IsSuccess() bool {
	return len(q.Errors) == 0 && util.IsNotBlank(&q.Id)
}

func (q QrCodeResponse) IsFailure() bool {
	return !q.IsSuccess()
}

func (q QrCodeResponse) IsNoContent() bool {
	return len(q.Errors) == 0 && util.IsBlank(&q.Id)
}

func (s SubscriptionResponse) IsSuccess() bool {
	return len(s.Errors) == 0 && util.IsNotBlank(&s.Id)
}

func (s SubscriptionResponse) IsFailure() bool {
	return !s.IsSuccess()
}

func (s SubscriptionResponse) IsNoContent() bool {
	return len(s.Errors) == 0 && util.IsBlank(&s.Id)
}

func (i InvoiceResponse) IsSuccess() bool {
	return len(i.Errors) == 0 && util.IsNotBlank(&i.Id)
}

func (i InvoiceResponse) IsFailure() bool {
	return !i.IsSuccess()
}

func (i InvoiceResponse) IsNoContent() bool {
	return len(i.Errors) == 0 && util.IsBlank(&i.Id)
}

func (i InvoiceSettingResponse) IsSuccess() bool {
	return len(i.Errors) == 0
}

func (i InvoiceSettingResponse) IsFailure() bool {
	return !i.IsSuccess()
}

func (i InvoiceSettingResponse) IsNoContent() bool {
	return util.IsBlank(&i.MunicipalServiceId) && util.IsBlank(&i.MunicipalServiceCode) &&
		util.IsBlank(&i.MunicipalServiceName) && util.IsBlank(&i.InvoiceCreationPeriod) &&
		util.IsBlank(&i.Observations) && i.Deductions == 0 && !i.DaysBeforeDueDate.IsEnumValid() &&
		!i.ReceivedOnly && i.Taxes == nil && len(i.Errors) == 0
}

func (n NotificationResponse) IsSuccess() bool {
	return len(n.Errors) == 0 && util.IsNotBlank(&n.Id)
}

func (n NotificationResponse) IsFailure() bool {
	return !n.IsSuccess()
}

func (n NotificationResponse) IsNoContent() bool {
	return len(n.Errors) == 0 && util.IsBlank(&n.Id)
}

func (u UpdateManyNotificationsResponse) IsSuccess() bool {
	return len(u.Errors) == 0 && len(u.Notifications) > 0
}

func (u UpdateManyNotificationsResponse) IsFailure() bool {
	return !u.IsSuccess()
}

func (u UpdateManyNotificationsResponse) IsNoContent() bool {
	return false
}

func (p Pageable[T]) IsSuccess() bool {
	return len(p.Errors) == 0 && len(p.Data) > 0
}

func (p Pageable[T]) IsFailure() bool {
	return !p.IsSuccess()
}

func (p Pageable[T]) IsNoContent() bool {
	return len(p.Errors) == 0 && len(p.Data) == 0
}

func (c ChargeResponse) IsSuccess() bool {
	return len(c.Errors) == 0 && util.IsNotBlank(&c.Id)
}

func (c ChargeResponse) IsFailure() bool {
	return !c.IsSuccess()
}

func (c ChargeResponse) IsNoContent() bool {
	return len(c.Errors) == 0 && util.IsBlank(&c.Id)
}

func (d DeleteResponse) IsSuccess() bool {
	return len(d.Errors) == 0 && d.Deleted && util.IsNotBlank(&d.Id)
}

func (d DeleteResponse) IsFailure() bool {
	return !d.IsSuccess()
}

func (d DeleteResponse) IsNoContent() bool {
	return len(d.Errors) == 0 && util.IsBlank(&d.Id)
}

func (c ChargeDocumentResponse) IsSuccess() bool {
	return len(c.Errors) == 0 && util.IsNotBlank(&c.Id) && c.File != nil
}

func (c ChargeDocumentResponse) IsFailure() bool {
	return !c.IsSuccess()
}

func (c ChargeDocumentResponse) IsNoContent() bool {
	return len(c.Errors) == 0 && util.IsBlank(&c.Id)
}

func (i IdentificationFieldResponse) IsSuccess() bool {
	return len(i.Errors) == 0 && util.IsNotBlank(&i.IdentificationField)
}

func (i IdentificationFieldResponse) IsFailure() bool {
	return !i.IsSuccess()
}

func (i IdentificationFieldResponse) IsNoContent() bool {
	return len(i.Errors) == 0 && util.IsBlank(&i.IdentificationField)
}

func (c ChargePixQrCodeResponse) IsSuccess() bool {
	return len(c.Errors) == 0 && util.IsNotBlank(&c.EncodedImage) && util.IsNotBlank(&c.Payload)
}

func (c ChargePixQrCodeResponse) IsFailure() bool {
	return !c.IsSuccess()
}

func (c ChargePixQrCodeResponse) IsNoContent() bool {
	return len(c.Errors) == 0 && util.IsBlank(&c.EncodedImage) && util.IsBlank(&c.Payload)
}

func (c ChargeCreationLimitResponse) IsSuccess() bool {
	return len(c.Errors) == 0
}

func (c ChargeCreationLimitResponse) IsFailure() bool {
	return !c.IsSuccess()
}

func (c ChargeCreationLimitResponse) IsNoContent() bool {
	return false
}

func (c ChargeStatusResponse) IsSuccess() bool {
	return len(c.Errors) == 0 && c.Status.IsEnumValid()
}

func (c ChargeStatusResponse) IsFailure() bool {
	return !c.IsSuccess()
}

func (c ChargeStatusResponse) IsNoContent() bool {
	return len(c.Errors) == 0 && !c.Status.IsEnumValid()
}

func (c CreditCardTokenizeResponse) IsSuccess() bool {
	return len(c.Errors) == 0 && util.IsNotBlank(&c.CreditCardToken)
}

func (c CreditCardTokenizeResponse) IsFailure() bool {
	return !c.IsSuccess()
}

func (c CreditCardTokenizeResponse) IsNoContent() bool {
	return len(c.Errors) == 0 && util.IsBlank(&c.CreditCardToken)
}

func (c CustomerResponse) IsSuccess() bool {
	return len(c.Errors) == 0 && util.IsNotBlank(&c.Id)
}

func (c CustomerResponse) IsFailure() bool {
	return !c.IsSuccess()
}

func (c CustomerResponse) IsNoContent() bool {
	return len(c.Errors) == 0 && util.IsBlank(&c.Id)
}
