package asaas

import (
	"context"
	"fmt"
	"net/http"
)

type CreateInvoiceSettingRequest struct {
	MunicipalServiceId   string                   `json:"municipalServiceId,omitempty"`
	MunicipalServiceCode string                   `json:"municipalServiceCode,omitempty"`
	MunicipalServiceName string                   `json:"municipalServiceName,omitempty"`
	UpdatePayment        bool                     `json:"updatePayment,omitempty"`
	Deductions           float64                  `json:"deductions,omitempty"`
	EffectiveDatePeriod  InvoiceDatePeriod        `json:"effectiveDatePeriod,omitempty" validate:"omitempty,enum"`
	ReceivedOnly         bool                     `json:"receivedOnly,omitempty"`
	DaysBeforeDueDate    InvoiceDaysBeforeDueDate `json:"daysBeforeDueDate,omitempty" validate:"omitempty,enum"`
	Observations         string                   `json:"observations,omitempty"`
	Taxes                *InvoiceTaxesRequest     `json:"taxes,omitempty"`
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
	Deductions          float64                  `json:"deductions,omitempty"`
	EffectiveDatePeriod InvoiceDatePeriod        `json:"effectiveDatePeriod,omitempty" validate:"omitempty,enum"`
	ReceivedOnly        bool                     `json:"receivedOnly,omitempty"`
	DaysBeforeDueDate   InvoiceDaysBeforeDueDate `json:"daysBeforeDueDate,omitempty" validate:"omitempty,enum"`
	Observations        string                   `json:"observations,omitempty"`
	Taxes               *InvoiceTaxesRequest     `json:"taxes,omitempty"`
}

type InvoiceTaxesRequest struct {
	RetainIss bool    `json:"retainIss,omitempty"`
	Iss       float64 `json:"iss,omitempty" validate:"required,gt=0"`
	Confins   float64 `json:"cofins,omitempty" validate:"required,gt=0"`
	Csll      float64 `json:"csll,omitempty" validate:"required,gt=0"`
	Inss      float64 `json:"inss,omitempty" validate:"required,gt=0"`
	Ir        float64 `json:"ir,omitempty" validate:"required,gt=0"`
	Pis       float64 `json:"pis,omitempty" validate:"required,gt=0"`
}

type GetAllInvoicesRequest struct {
	EffectiveDateGE   *Date         `json:"effectiveDate[ge],omitempty"`
	EffectiveDateLE   *Date         `json:"effectiveDate[le],omitempty"`
	Payment           string        `json:"payment,omitempty"`
	Installment       string        `json:"installment,omitempty"`
	Customer          string        `json:"customer,omitempty"`
	ExternalReference string        `json:"externalReference,omitempty"`
	Status            InvoiceStatus `json:"status,omitempty"`
	Offset            int           `json:"offset,omitempty"`
	Limit             int           `json:"limit,omitempty"`
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
	UpdateByID(ctx context.Context, invoiceId string, body UpdateInvoiceRequest) (*InvoiceResponse, Error)
	GetByID(ctx context.Context, invoiceId string) (*InvoiceResponse, Error)
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

func (i invoice) UpdateByID(ctx context.Context, invoiceId string, body UpdateInvoiceRequest) (*InvoiceResponse, Error) {
	req := NewRequest[InvoiceResponse](ctx, i.env, i.accessToken)
	return req.make(http.MethodPost, fmt.Sprintf("/v3/invoices/%s", invoiceId), body)
}

func (i invoice) GetByID(ctx context.Context, invoiceId string) (*InvoiceResponse, Error) {
	req := NewRequest[InvoiceResponse](ctx, i.env, i.accessToken)
	return req.make(http.MethodGet, fmt.Sprintf("/v3/invoices/%s", invoiceId), nil)
}

func (i invoice) GetAll(ctx context.Context, filter GetAllInvoicesRequest) (*Pageable[InvoiceResponse], Error) {
	req := NewRequest[Pageable[InvoiceResponse]](ctx, i.env, i.accessToken)
	return req.make(http.MethodGet, "/v3/invoices/", filter)
}
