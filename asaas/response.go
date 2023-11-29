package asaas

import "github.com/GabrielHCataldo/go-asaas/internal/util"

type DeleteResponse struct {
	ID      string `json:"id,omitempty"`
	Deleted bool   `json:"deleted,omitempty"`
}

type response interface {
	IsSuccess() bool
	IsFailure() bool
	IsNoContent() bool
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
	return len(m.Errors) == 0 && util.IsNotBlank(&m.ID)
}

func (m MobilePhoneRechargeResponse) IsFailure() bool {
	return !m.IsSuccess()
}

func (m MobilePhoneRechargeResponse) IsNoContent() bool {
	return len(m.Errors) == 0 && util.IsBlank(&m.ID)
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
	return len(b.Errors) == 0 && util.IsNotBlank(&b.ID)
}

func (b BillPaymentResponse) IsFailure() bool {
	return !b.IsSuccess()
}

func (b BillPaymentResponse) IsNoContent() bool {
	return len(b.Errors) == 0 && util.IsBlank(&b.ID)
}

func (n NegativitySimulateResponse) IsSuccess() bool {
	return util.IsNotBlank(&n.Payment)
}

func (n NegativitySimulateResponse) IsFailure() bool {
	return !n.IsSuccess()
}

func (n NegativitySimulateResponse) IsNoContent() bool {
	return util.IsBlank(&n.Payment)
}

func (n NegativityResponse) IsSuccess() bool {
	return len(n.Errors) == 0 && util.IsNotBlank(&n.ID)
}

func (n NegativityResponse) IsFailure() bool {
	return !n.IsSuccess()
}

func (n NegativityResponse) IsNoContent() bool {
	return len(n.Errors) == 0 && util.IsBlank(&n.ID)
}

func (a AnticipationResponse) IsSuccess() bool {
	return len(a.Errors) == 0 && util.IsNotBlank(&a.ID)
}

func (a AnticipationResponse) IsFailure() bool {
	return !a.IsSuccess()
}

func (a AnticipationResponse) IsNoContent() bool {
	return len(a.Errors) == 0 && util.IsBlank(&a.ID)
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
	return true
}

func (a AgreementSignResponse) IsFailure() bool {
	return false
}

func (a AgreementSignResponse) IsNoContent() bool {
	return false
}

func (a AnticipationLimitsResponse) IsSuccess() bool {
	return true
}

func (a AnticipationLimitsResponse) IsFailure() bool {
	return false
}

func (a AnticipationLimitsResponse) IsNoContent() bool {
	return false
}

func (t TransferResponse) IsSuccess() bool {
	return len(t.Errors) == 0 && util.IsNotBlank(&t.ID)
}

func (t TransferResponse) IsFailure() bool {
	return !t.IsSuccess()
}

func (t TransferResponse) IsNoContent() bool {
	return len(t.Errors) == 0 && util.IsBlank(&t.ID)
}

func (p PixTransactionResponse) IsSuccess() bool {
	return len(p.Errors) == 0 && util.IsNotBlank(&p.ID)
}

func (p PixTransactionResponse) IsFailure() bool {
	return !p.IsSuccess()
}

func (p PixTransactionResponse) IsNoContent() bool {
	return len(p.Errors) == 0 && util.IsBlank(&p.ID)
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
	return len(p.Errors) == 0 && util.IsNotBlank(&p.ID)
}

func (p PixCancelTransactionResponse) IsFailure() bool {
	return !p.IsSuccess()
}

func (p PixCancelTransactionResponse) IsNoContent() bool {
	return len(p.Errors) == 0 && util.IsBlank(&p.ID)
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
	return len(p.Errors) == 0 && util.IsNotBlank(&p.ID)
}

func (p PixKeyResponse) IsFailure() bool {
	return !p.IsSuccess()
}

func (p PixKeyResponse) IsNoContent() bool {
	return len(p.Errors) == 0 && util.IsBlank(&p.ID)
}

func (q QrCodeResponse) IsSuccess() bool {
	return len(q.Errors) == 0 && util.IsNotBlank(&q.ID)
}

func (q QrCodeResponse) IsFailure() bool {
	return !q.IsSuccess()
}

func (q QrCodeResponse) IsNoContent() bool {
	return len(q.Errors) == 0 && util.IsBlank(&q.ID)
}

func (s SubscriptionResponse) IsSuccess() bool {
	return len(s.Errors) == 0 && util.IsNotBlank(&s.ID)
}

func (s SubscriptionResponse) IsFailure() bool {
	return !s.IsSuccess()
}

func (s SubscriptionResponse) IsNoContent() bool {
	return len(s.Errors) == 0 && util.IsBlank(&s.ID)
}

func (i InvoiceResponse) IsSuccess() bool {
	return util.IsNotBlank(&i.ID)
}

func (i InvoiceResponse) IsFailure() bool {
	return !i.IsSuccess()
}

func (i InvoiceResponse) IsNoContent() bool {
	return util.IsBlank(&i.ID)
}

func (i InvoiceSettingResponse) IsSuccess() bool {
	return len(i.Errors) == 0 && (util.IsNotBlank(&i.MunicipalServiceId) || util.IsNotBlank(&i.MunicipalServiceCode) ||
		util.IsNotBlank(&i.MunicipalServiceName) || util.IsNotBlank(&i.InvoiceCreationPeriod) ||
		util.IsNotBlank(&i.Observations) || i.Deductions != 0 || i.DaysBeforeDueDate.IsEnumValid() ||
		i.ReceivedOnly || i.Taxes != nil)
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
	return len(n.Errors) == 0 && util.IsNotBlank(&n.ID)
}

func (n NotificationResponse) IsFailure() bool {
	return !n.IsSuccess()
}

func (n NotificationResponse) IsNoContent() bool {
	return len(n.Errors) == 0 && util.IsBlank(&n.ID)
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
	return len(p.Data) > 0
}

func (p Pageable[T]) IsFailure() bool {
	return !p.IsSuccess()
}

func (p Pageable[T]) IsNoContent() bool {
	return len(p.Data) == 0
}

func (c ChargeResponse) IsSuccess() bool {
	return len(c.Errors) == 0 && util.IsNotBlank(&c.ID)
}

func (c ChargeResponse) IsFailure() bool {
	return !c.IsSuccess()
}

func (c ChargeResponse) IsNoContent() bool {
	return len(c.Errors) == 0 && util.IsBlank(&c.ID)
}

func (d DeleteResponse) IsSuccess() bool {
	return d.Deleted && util.IsNotBlank(&d.ID)
}

func (d DeleteResponse) IsFailure() bool {
	return !d.IsSuccess()
}

func (d DeleteResponse) IsNoContent() bool {
	return util.IsBlank(&d.ID)
}

func (c ChargeDocumentResponse) IsSuccess() bool {
	return len(c.Errors) == 0 && util.IsNotBlank(&c.ID) && c.File != nil
}

func (c ChargeDocumentResponse) IsFailure() bool {
	return !c.IsSuccess()
}

func (c ChargeDocumentResponse) IsNoContent() bool {
	return len(c.Errors) == 0 && util.IsBlank(&c.ID)
}

func (i IdentificationFieldResponse) IsSuccess() bool {
	return util.IsNotBlank(&i.IdentificationField) && util.IsNotBlank(&i.BarCode) && util.IsNotBlank(&i.NossoNumero)
}

func (i IdentificationFieldResponse) IsFailure() bool {
	return !i.IsSuccess()
}

func (i IdentificationFieldResponse) IsNoContent() bool {
	return util.IsBlank(&i.IdentificationField) && util.IsBlank(&i.BarCode) && util.IsBlank(&i.NossoNumero)
}

func (c ChargePixQrCodeResponse) IsSuccess() bool {
	return util.IsNotBlank(&c.EncodedImage) && util.IsNotBlank(&c.Payload) && !c.ExpirationDate.IsZero()
}

func (c ChargePixQrCodeResponse) IsFailure() bool {
	return !c.IsSuccess()
}

func (c ChargePixQrCodeResponse) IsNoContent() bool {
	return util.IsBlank(&c.EncodedImage) && util.IsBlank(&c.Payload) && c.ExpirationDate.IsZero()
}

func (c ChargeCreationLimitResponse) IsSuccess() bool {
	return true
}

func (c ChargeCreationLimitResponse) IsFailure() bool {
	return false
}

func (c ChargeCreationLimitResponse) IsNoContent() bool {
	return false
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
	return len(c.Errors) == 0 && util.IsNotBlank(&c.ID)
}

func (c CustomerResponse) IsFailure() bool {
	return !c.IsSuccess()
}

func (c CustomerResponse) IsNoContent() bool {
	return len(c.Errors) == 0 && util.IsBlank(&c.ID)
}
