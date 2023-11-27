package asaas

import "github.com/GabrielHCataldo/go-asaas/internal/util"

type response interface {
	IsSuccess() bool
	IsNoContent() bool
}

func (p Pageable[T]) IsSuccess() bool {
	return len(p.Data) > 0
}

func (p Pageable[T]) IsNoContent() bool {
	return len(p.Data) == 0
}

func (c ChargeResponse) IsSuccess() bool {
	return len(c.Errors) == 0 && util.IsNotBlank(&c.ID)
}

func (c ChargeResponse) IsNoContent() bool {
	return len(c.Errors) == 0 && util.IsBlank(&c.ID)
}

func (c ChargeDeleteResponse) IsSuccess() bool {
	return c.Deleted && util.IsNotBlank(&c.ID)
}

func (c ChargeDeleteResponse) IsNoContent() bool {
	return false
}

func (c ChargeDocumentResponse) IsSuccess() bool {
	return len(c.Errors) == 0 && util.IsNotBlank(&c.ID) && c.File != nil
}

func (c ChargeDocumentResponse) IsNoContent() bool {
	return len(c.Errors) == 0 && util.IsBlank(&c.ID)
}

func (i IdentificationFieldResponse) IsSuccess() bool {
	return util.IsNotBlank(&i.IdentificationField) && util.IsNotBlank(&i.BarCode) && util.IsNotBlank(&i.NossoNumero)
}

func (i IdentificationFieldResponse) IsNoContent() bool {
	return util.IsBlank(&i.IdentificationField) && util.IsBlank(&i.BarCode) && util.IsBlank(&i.NossoNumero)
}

func (p PixQRCodeResponse) IsSuccess() bool {
	return util.IsNotBlank(&p.EncodedImage) && util.IsNotBlank(&p.Payload) && !p.ExpirationDate.IsZero()
}

func (p PixQRCodeResponse) IsNoContent() bool {
	return util.IsBlank(&p.EncodedImage) && util.IsBlank(&p.Payload) && p.ExpirationDate.IsZero()
}
