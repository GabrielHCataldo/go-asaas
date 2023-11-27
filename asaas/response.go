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

func (p PixQRCodeResponse) IsSuccess() bool {
	return util.IsNotBlank(&p.EncodedImage) && util.IsNotBlank(&p.Payload) && !p.ExpirationDate.IsZero()
}

func (p PixQRCodeResponse) IsFailure() bool {
	return !p.IsSuccess()
}

func (p PixQRCodeResponse) IsNoContent() bool {
	return util.IsBlank(&p.EncodedImage) && util.IsBlank(&p.Payload) && p.ExpirationDate.IsZero()
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
