package asaas

import (
	"context"
	"fmt"
	"net/http"
)

type SaveInvoiceSettingRequest struct {
	// Identificador único do serviço municipal.
	MunicipalServiceId string `json:"municipalServiceId,omitempty"`
	// Código de serviço municipal.
	MunicipalServiceCode string `json:"municipalServiceCode,omitempty"`
	// Nome do serviço municipal. Se não for informado, será utilizado o código do serviço municipal como nome para identificação.
	MunicipalServiceName string `json:"municipalServiceName,omitempty"`
	// Atualizar o valor da cobrança com os impostos da nota já descontados.
	UpdatePayment bool `json:"updatePayment,omitempty"`
	// Deduções. As deduções não alteram o valor total da nota fiscal, mas alteram a base de cálculo do ISS.
	Deductions float64 `json:"deductions,omitempty"`
	// Quando a nota fiscal será emitida.
	EffectiveDatePeriod InvoiceDatePeriod `json:"effectiveDatePeriod,omitempty" validate:"omitempty,enum"`
	// Emitir apenas para cobranças pagas.
	ReceivedOnly bool `json:"receivedOnly,omitempty"`
	// Quantidade de dias antes do vencimento da cobrança.
	DaysBeforeDueDate InvoiceDaysBeforeDueDate `json:"daysBeforeDueDate,omitempty" validate:"omitempty,enum"`
	// Observações adicionais da nota fiscal.
	Observations string `json:"observations,omitempty"`
	// Impostos da nota fiscal.
	Taxes *InvoiceTaxesRequest `json:"taxes,omitempty"`
}

type ScheduleInvoiceRequest struct {
	Payment              string               `json:"payment,omitempty"`
	Installment          string               `json:"installment,omitempty"`
	Customer             string               `json:"customer,omitempty"`
	ServiceDescription   string               `json:"serviceDescription,omitempty" validate:"required"`
	Observations         string               `json:"observations,omitempty" validate:"required"`
	ExternalReference    string               `json:"externalReference,omitempty"`
	Value                float64              `json:"value,omitempty" validate:"required,gt=0"`
	Deductions           float64              `json:"deductions,omitempty" validate:"gte=0"`
	EffectiveDate        Date                 `json:"effectiveDate,omitempty" validate:"required"`
	MunicipalServiceId   string               `json:"municipalServiceId,omitempty"`
	MunicipalServiceCode string               `json:"municipalServiceCode,omitempty"`
	MunicipalServiceName string               `json:"municipalServiceName,omitempty" validate:"required"`
	UpdatePayment        bool                 `json:"updatePayment,omitempty"`
	Taxes                *InvoiceTaxesRequest `json:"taxes,omitempty"`
}

type UpdateInvoiceRequest struct {
	Payment              string               `json:"payment,omitempty"`
	Installment          string               `json:"installment,omitempty"`
	Customer             string               `json:"customer,omitempty"`
	ServiceDescription   string               `json:"serviceDescription,omitempty"`
	Observations         string               `json:"observations,omitempty"`
	ExternalReference    string               `json:"externalReference,omitempty"`
	Value                float64              `json:"value,omitempty" validate:"omitempty,gt=0"`
	Deductions           float64              `json:"deductions,omitempty"`
	EffectiveDate        Date                 `json:"effectiveDate,omitempty"`
	MunicipalServiceId   string               `json:"municipalServiceId,omitempty"`
	MunicipalServiceCode string               `json:"municipalServiceCode,omitempty"`
	MunicipalServiceName string               `json:"municipalServiceName,omitempty"`
	UpdatePayment        bool                 `json:"updatePayment,omitempty"`
	Taxes                *InvoiceTaxesRequest `json:"taxes,omitempty"`
}

type UpdateInvoiceSettingRequest struct {
	// Deduções. As deduções não alteram o valor total da nota fiscal, mas alteram a base de cálculo do ISS.
	Deductions float64 `json:"deductions,omitempty"`
	// Quando a nota fiscal será emitida.
	EffectiveDatePeriod InvoiceDatePeriod `json:"effectiveDatePeriod,omitempty" validate:"omitempty,enum"`
	// Emitir apenas para cobranças pagas.
	ReceivedOnly bool `json:"receivedOnly,omitempty"`
	// Quantidade de dias antes do vencimento da cobrança.
	DaysBeforeDueDate InvoiceDaysBeforeDueDate `json:"daysBeforeDueDate,omitempty" validate:"omitempty,enum"`
	// Observações adicionais da nota fiscal.
	Observations string `json:"observations,omitempty"`
	// Impostos da nota fiscal.
	Taxes *InvoiceTaxesRequest `json:"taxes,omitempty"`
}

type InvoiceTaxesRequest struct {
	// Tomador da nota fiscal deve reter ISS ou não
	RetainIss bool `json:"retainIss,omitempty"`
	// Alíquota ISS (REQUIRED)
	Iss float64 `json:"iss,omitempty" validate:"required,gt=0"`
	// Alíquota COFINS (REQUIRED)
	Confins float64 `json:"cofins,omitempty" validate:"required,gt=0"`
	// Alíquota CSLL (REQUIRED)
	Csll float64 `json:"csll,omitempty" validate:"required,gt=0"`
	// Alíquota INSS (REQUIRED)
	Inss float64 `json:"inss,omitempty" validate:"required,gt=0"`
	// Alíquota IR (REQUIRED)
	Ir float64 `json:"ir,omitempty" validate:"required,gt=0"`
	// Alíquota PIS (REQUIRED)
	Pis float64 `json:"pis,omitempty" validate:"required,gt=0"`
}

type GetAllInvoicesRequest struct {
	// Filtrar a partir de uma data de emissão
	EffectiveDateGE *Date `json:"effectiveDate[ge],omitempty"`
	// Filtrar até uma data de emissão
	EffectiveDateLE *Date  `json:"effectiveDate[le],omitempty"`
	Payment         string `json:"payment,omitempty"`
	Installment     string `json:"installment,omitempty"`
	// Filtrar pelo identificador único do cliente
	Customer string `json:"customer,omitempty"`
	// Identificador da nota fiscal no seu sistema
	ExternalReference string `json:"externalReference,omitempty"`
	// Status da nota fiscal
	Status InvoiceStatus `json:"status,omitempty"`
	// Elemento inicial da lista
	Offset int `json:"offset,omitempty"`
	// Número de elementos da lista (max: 100)
	Limit int `json:"limit,omitempty"`
}

type InvoiceSettingResponse struct {
	MunicipalServiceId    string                   `json:"municipalServiceId,omitempty"`
	MunicipalServiceCode  string                   `json:"municipalServiceCode,omitempty"`
	MunicipalServiceName  string                   `json:"municipalServiceName,omitempty"`
	Deductions            float64                  `json:"deductions,omitempty"`
	InvoiceCreationPeriod string                   `json:"invoiceCreationPeriod,omitempty"`
	DaysBeforeDueDate     InvoiceDaysBeforeDueDate `json:"daysBeforeDueDate,omitempty"`
	ReceivedOnly          bool                     `json:"receivedOnly,omitempty"`
	Observations          string                   `json:"observations,omitempty"`
	Taxes                 *InvoiceTaxesResponse    `json:"taxes,omitempty"`
	Errors                []ErrorResponse          `json:"errors,omitempty"`
}

type InvoiceResponse struct {
	Id                        string                `json:"id,omitempty"`
	Payment                   string                `json:"payment,omitempty"`
	Installment               string                `json:"installment,omitempty"`
	Customer                  string                `json:"customer,omitempty"`
	Status                    InvoiceStatus         `json:"status,omitempty"`
	Type                      string                `json:"type,omitempty"`
	StatusDescription         string                `json:"statusDescription,omitempty"`
	ServiceDescription        string                `json:"serviceDescription,omitempty"`
	PdfUrl                    string                `json:"pdfUrl,omitempty"`
	XmlUrl                    string                `json:"xmlUrl,omitempty"`
	RpsSerie                  string                `json:"rpsSerie,omitempty"`
	RpsNumber                 string                `json:"rpsNumber,omitempty"`
	Number                    string                `json:"number,omitempty"`
	ValidationCode            string                `json:"validationCode,omitempty"`
	Value                     float64               `json:"value,omitempty"`
	Deductions                float64               `json:"deductions,omitempty"`
	EffectiveDate             *Date                 `json:"effectiveDate,omitempty"`
	Observations              string                `json:"observations,omitempty"`
	EstimatedTaxesDescription string                `json:"estimatedTaxesDescription,omitempty"`
	ExternalReference         string                `json:"externalReference,omitempty"`
	Taxes                     *InvoiceTaxesResponse `json:"taxes,omitempty"`
	MunicipalServiceId        string                `json:"municipalServiceId,omitempty"`
	MunicipalServiceCode      string                `json:"municipalServiceCode,omitempty"`
	MunicipalServiceName      string                `json:"municipalServiceName,omitempty"`
	Errors                    []ErrorResponse       `json:"errors,omitempty"`
}

type InvoiceTaxesResponse struct {
	RetainIss bool    `json:"retainIss"`
	Iss       float64 `json:"iss"`
	Confins   float64 `json:"cofins"`
	Csll      float64 `json:"csll"`
	Inss      float64 `json:"inss"`
	Ir        float64 `json:"ir"`
	Pis       float64 `json:"pis"`
}

type invoice struct {
	env         Env
	accessToken string
}

type Invoice interface {
	Schedule(ctx context.Context, body ScheduleInvoiceRequest) (*InvoiceResponse, Error)
	Authorize(ctx context.Context, invoiceId string) (*InvoiceResponse, Error)
	Cancel(ctx context.Context, invoiceId string) (*InvoiceResponse, Error)
	UpdateById(ctx context.Context, invoiceId string, body UpdateInvoiceRequest) (*InvoiceResponse, Error)
	GetById(ctx context.Context, invoiceId string) (*InvoiceResponse, Error)
	GetAll(ctx context.Context, filter GetAllInvoicesRequest) (*Pageable[InvoiceResponse], Error)
}

func NewInvoice(env Env, accessToken string) Invoice {
	logWarning("Invoice service running on", env.String())
	return invoice{
		env:         env,
		accessToken: accessToken,
	}
}

func (i invoice) Schedule(ctx context.Context, body ScheduleInvoiceRequest) (*InvoiceResponse, Error) {
	if err := Validate().Struct(body); err != nil {
		return nil, NewError(ErrorTypeValidation, err)
	}
	req := NewRequest[InvoiceResponse](ctx, i.env, i.accessToken)
	return req.make(http.MethodPost, "/v3/invoices", body)
}

func (i invoice) Authorize(ctx context.Context, invoiceId string) (*InvoiceResponse, Error) {
	req := NewRequest[InvoiceResponse](ctx, i.env, i.accessToken)
	return req.make(http.MethodPost, fmt.Sprintf("/v3/invoices/%s/authorize", invoiceId), nil)
}

func (i invoice) Cancel(ctx context.Context, invoiceId string) (*InvoiceResponse, Error) {
	req := NewRequest[InvoiceResponse](ctx, i.env, i.accessToken)
	return req.make(http.MethodPost, fmt.Sprintf("/v3/invoices/%s/cancel", invoiceId), nil)
}

func (i invoice) UpdateById(ctx context.Context, invoiceId string, body UpdateInvoiceRequest) (*InvoiceResponse, Error) {
	req := NewRequest[InvoiceResponse](ctx, i.env, i.accessToken)
	return req.make(http.MethodPost, fmt.Sprintf("/v3/invoices/%s", invoiceId), body)
}

func (i invoice) GetById(ctx context.Context, invoiceId string) (*InvoiceResponse, Error) {
	req := NewRequest[InvoiceResponse](ctx, i.env, i.accessToken)
	return req.make(http.MethodGet, fmt.Sprintf("/v3/invoices/%s", invoiceId), nil)
}

func (i invoice) GetAll(ctx context.Context, filter GetAllInvoicesRequest) (*Pageable[InvoiceResponse], Error) {
	req := NewRequest[Pageable[InvoiceResponse]](ctx, i.env, i.accessToken)
	return req.make(http.MethodGet, "/v3/invoices/", filter)
}
