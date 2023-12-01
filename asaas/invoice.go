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
	// ID da cobrança a ser antecipada (REQUIRED se Installment, Customer não for informado)
	Payment string `json:"payment,omitempty"`
	// ID do parcelamento a ser antecipado (REQUIRED se Payment, Customer não for informado)
	Installment string `json:"installment,omitempty"`
	// Identificador único do cliente no Asaas (REQUIRED se Payment, Installment não for informado)
	Customer string `json:"customer,omitempty"`
	// Descrição dos serviços da nota fiscal (REQUIRED)
	ServiceDescription string `json:"serviceDescription,omitempty" validate:"required"`
	// Observações adicionais da nota fiscal (REQUIRED)
	Observations string `json:"observations,omitempty" validate:"required"`
	// Identificador da nota fiscal no seu sistema
	ExternalReference string `json:"externalReference,omitempty"`
	// Valor (REQUIRED)
	Value float64 `json:"value,omitempty" validate:"required,gt=0"`
	// Deduções. As deduções não alteram o valor total da nota fiscal, mas alteram a base de cálculo do ISS (REQUIRED)
	Deductions float64 `json:"deductions,omitempty" validate:"gte=0"`
	// Data de emissão da nota fiscal (REQUIRED)
	EffectiveDate Date `json:"effectiveDate,omitempty" validate:"required"`
	// Identificador único do serviço municipal.
	MunicipalServiceId string `json:"municipalServiceId,omitempty"`
	// Código de serviço municipal
	MunicipalServiceCode string `json:"municipalServiceCode,omitempty"`
	// Nome do serviço municipal. Se não for informado, será utilizado o atributo MunicipalServiceCode como nome para identificação.
	MunicipalServiceName string `json:"municipalServiceName,omitempty"`
	// Atualizar o valor da cobrança com os impostos da nota já descontados.
	UpdatePayment bool `json:"updatePayment,omitempty"`
	// Impostos da nota fiscal (REQUIRED)
	Taxes *InvoiceTaxesRequest `json:"taxes,omitempty"`
}

type UpdateInvoiceRequest struct {
	// Descrição dos serviços da nota fiscal
	ServiceDescription string `json:"serviceDescription,omitempty"`
	// Observações adicionais da nota fiscal
	Observations string `json:"observations,omitempty"`
	// Identificador da nota fiscal no seu sistema
	ExternalReference string `json:"externalReference,omitempty"`
	// Valor
	Value float64 `json:"value,omitempty" validate:"omitempty,gt=0"`
	// Deduções. As deduções não alteram o valor total da nota fiscal, mas alteram a base de cálculo do ISS
	Deductions float64 `json:"deductions,omitempty"`
	// Data de emissão da nota fiscal
	EffectiveDate Date `json:"effectiveDate,omitempty"`
	// Identificador único do serviço municipal.
	MunicipalServiceId string `json:"municipalServiceId,omitempty"`
	// Código de serviço municipal
	MunicipalServiceCode string `json:"municipalServiceCode,omitempty"`
	// Nome do serviço municipal. Se não for informado, será utilizado o atributo MunicipalServiceCode como nome para identificação.
	MunicipalServiceName string `json:"municipalServiceName,omitempty"`
	// Atualizar o valor da cobrança com os impostos da nota já descontados.
	UpdatePayment bool `json:"updatePayment,omitempty"`
	// Impostos da nota fiscal
	Taxes *InvoiceTaxesRequest `json:"taxes,omitempty"`
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
	// Schedule (Agendar nota fiscal)
	//
	// # Resposta: 200
	//
	// InvoiceResponse = not nil
	//
	// Error = nil
	//
	// InvoiceResponse.IsSuccess() = true
	//
	// Possui os valores de resposta de sucesso segunda a documentação.
	//
	// # Resposta: 400/401/500
	//
	// InvoiceResponse = not nil
	//
	// Error = nil
	//
	// InvoiceResponse.IsFailure() = true
	//
	// Para qualquer outra resposta inesperada da API, possuímos o campo InvoiceResponse.Errors preenchido com as informações
	// de erro, sendo 400 retornado da API Asaas com as instruções de requisição conforme a documentação,
	// diferente disso retornará uma mensagem padrão no index 0 do slice com campo ErrorResponse.Code retornando a
	// descrição status http (Ex: "401 Unauthorized") e no campo ErrorResponse.Description retornará com o valor
	// "response status code not expected".
	//
	// # Error
	//
	// InvoiceResponse = nil
	//
	// Error = not nil
	//
	// Se o campo ErrorAsaas.Type tiver com valor ErrorTypeValidation quer dizer que não passou pela validação dos
	// parâmetros informados segundo a documentação.
	// Por fim se o campo ErrorAsaas.Type tiver com valor ErrorTypeUnexpected quer dizer que ocorreu um erro inesperado
	// na lib go-asaas.
	//
	// Para obter mais detalhes confira as colunas:
	//
	// ErrorAsaas.Msg (mensagem do erro),
	//
	// ErrorAsaas.File (Arquivo aonde ocorreu o erro),
	//
	// ErrorAsaas.Line (Linha aonde ocorreu o erro)
	//
	// Caso ocorra um erro inesperado por favor report o erro no repositório: https://github.com/GabrielHCataldo/go-asaas
	//
	// # DOCS
	//
	// Agendar nota fiscal: https://docs.asaas.com/reference/agendar-nota-fiscal
	Schedule(ctx context.Context, body ScheduleInvoiceRequest) (*InvoiceResponse, Error)
	// Authorize (Emitir uma nota fiscal)
	//
	// Para emitir uma nota fiscal específica é necessário que você tenha o ID que o Asaas retornou no momento
	// da criação dela.
	//
	// # Resposta: 200
	//
	// InvoiceResponse = not nil
	//
	// Error = nil
	//
	// InvoiceResponse.IsSuccess() = true
	//
	// Possui os valores de resposta de sucesso segunda a documentação.
	//
	// # Resposta: 404
	//
	// InvoiceResponse = not nil
	//
	// Error = nil
	//
	// InvoiceResponse.IsNoContent() = true
	//
	// ID(s) informado no parâmetro não foi encontrado.
	//
	// # Resposta: 400/401/500
	//
	// InvoiceResponse = not nil
	//
	// Error = nil
	//
	// InvoiceResponse.IsFailure() = true
	//
	// Para qualquer outra resposta inesperada da API, possuímos o campo InvoiceResponse.Errors preenchido com as informações
	// de erro, sendo 400 retornado da API Asaas com as instruções de requisição conforme a documentação,
	// diferente disso retornará uma mensagem padrão no index 0 do slice com campo ErrorResponse.Code retornando a
	// descrição status http (Ex: "401 Unauthorized") e no campo ErrorResponse.Description retornará com o valor
	// "response status code not expected".
	//
	// # Error
	//
	// InvoiceResponse = nil
	//
	// Error = not nil
	//
	// Se o campo ErrorAsaas.Type tiver com valor ErrorTypeValidation quer dizer que não passou pela validação dos
	// parâmetros informados segundo a documentação.
	// Por fim se o campo ErrorAsaas.Type tiver com valor ErrorTypeUnexpected quer dizer que ocorreu um erro inesperado
	// na lib go-asaas.
	//
	// Para obter mais detalhes confira as colunas:
	//
	// ErrorAsaas.Msg (mensagem do erro),
	//
	// ErrorAsaas.File (Arquivo aonde ocorreu o erro),
	//
	// ErrorAsaas.Line (Linha aonde ocorreu o erro)
	//
	// Caso ocorra um erro inesperado por favor report o erro no repositório: https://github.com/GabrielHCataldo/go-asaas
	//
	// # DOCS
	//
	// Cancelar uma nota fiscal: https://docs.asaas.com/reference/emitir-uma-nota-fiscal
	Authorize(ctx context.Context, invoiceId string) (*InvoiceResponse, Error)
	// CancelById (Cancelar uma nota fiscal)
	//
	// Para cancelar uma nota fiscal específica é necessário que você tenha o ID que o Asaas retornou no momento
	// da criação dela.
	//
	// # Resposta: 200
	//
	// InvoiceResponse = not nil
	//
	// Error = nil
	//
	// InvoiceResponse.IsSuccess() = true
	//
	// Possui os valores de resposta de sucesso segunda a documentação.
	//
	// # Resposta: 404
	//
	// InvoiceResponse = not nil
	//
	// Error = nil
	//
	// InvoiceResponse.IsNoContent() = true
	//
	// ID(s) informado no parâmetro não foi encontrado.
	//
	// # Resposta: 400/401/500
	//
	// InvoiceResponse = not nil
	//
	// Error = nil
	//
	// InvoiceResponse.IsFailure() = true
	//
	// Para qualquer outra resposta inesperada da API, possuímos o campo InvoiceResponse.Errors preenchido com as informações
	// de erro, sendo 400 retornado da API Asaas com as instruções de requisição conforme a documentação,
	// diferente disso retornará uma mensagem padrão no index 0 do slice com campo ErrorResponse.Code retornando a
	// descrição status http (Ex: "401 Unauthorized") e no campo ErrorResponse.Description retornará com o valor
	// "response status code not expected".
	//
	// # Error
	//
	// InvoiceResponse = nil
	//
	// Error = not nil
	//
	// Se o campo ErrorAsaas.Type tiver com valor ErrorTypeValidation quer dizer que não passou pela validação dos
	// parâmetros informados segundo a documentação.
	// Por fim se o campo ErrorAsaas.Type tiver com valor ErrorTypeUnexpected quer dizer que ocorreu um erro inesperado
	// na lib go-asaas.
	//
	// Para obter mais detalhes confira as colunas:
	//
	// ErrorAsaas.Msg (mensagem do erro),
	//
	// ErrorAsaas.File (Arquivo aonde ocorreu o erro),
	//
	// ErrorAsaas.Line (Linha aonde ocorreu o erro)
	//
	// Caso ocorra um erro inesperado por favor report o erro no repositório: https://github.com/GabrielHCataldo/go-asaas
	//
	// # DOCS
	//
	// Cancelar uma nota fiscal: https://docs.asaas.com/reference/cancelar-uma-nota-fiscal
	CancelById(ctx context.Context, invoiceId string) (*InvoiceResponse, Error)
	// UpdateById (Atualizar nota fiscal)
	//
	// É possível atualizar notas fiscais que ainda não tenham sido emitidas, ou seja, estão
	// com status InvoiceStatusScheduled ou InvoiceStatusError.
	//
	// # Resposta: 200
	//
	// InvoiceResponse = not nil
	//
	// Error = nil
	//
	// InvoiceResponse.IsSuccess() = true
	//
	// Possui os valores de resposta de sucesso segunda a documentação.
	//
	// # Resposta: 404
	//
	// InvoiceResponse = not nil
	//
	// Error = nil
	//
	// InvoiceResponse.IsNoContent() = true
	//
	// ID(s) informado no parâmetro não foi encontrado.
	//
	// # Resposta: 400/401/500
	//
	// InvoiceResponse = not nil
	//
	// Error = nil
	//
	// InvoiceResponse.IsFailure() = true
	//
	// Para qualquer outra resposta inesperada da API, possuímos o campo InvoiceResponse.Errors preenchido com as informações
	// de erro, sendo 400 retornado da API Asaas com as instruções de requisição conforme a documentação,
	// diferente disso retornará uma mensagem padrão no index 0 do slice com campo ErrorResponse.Code retornando a
	// descrição status http (Ex: "401 Unauthorized") e no campo ErrorResponse.Description retornará com o valor
	// "response status code not expected".
	//
	// # Error
	//
	// InvoiceResponse = nil
	//
	// Error = not nil
	//
	// Se o campo ErrorAsaas.Type tiver com valor ErrorTypeValidation quer dizer que não passou pela validação dos
	// parâmetros informados segundo a documentação.
	// Por fim se o campo ErrorAsaas.Type tiver com valor ErrorTypeUnexpected quer dizer que ocorreu um erro inesperado
	// na lib go-asaas.
	//
	// Para obter mais detalhes confira as colunas:
	//
	// ErrorAsaas.Msg (mensagem do erro),
	//
	// ErrorAsaas.File (Arquivo aonde ocorreu o erro),
	//
	// ErrorAsaas.Line (Linha aonde ocorreu o erro)
	//
	// Caso ocorra um erro inesperado por favor report o erro no repositório: https://github.com/GabrielHCataldo/go-asaas
	//
	// # DOCS
	//
	// Atualizar nota fiscal: https://docs.asaas.com/reference/atualizar-nota-fiscal
	UpdateById(ctx context.Context, invoiceId string, body UpdateInvoiceRequest) (*InvoiceResponse, Error)
	// GetById (Recuperar uma nota fiscal)
	//
	// # Resposta: 200
	//
	// InvoiceResponse = not nil
	//
	// Error = nil
	//
	// InvoiceResponse.IsSuccess() = true
	//
	// Possui os valores de resposta de sucesso segunda a documentação.
	//
	// # Resposta: 404
	//
	// InvoiceResponse = not nil
	//
	// Error = nil
	//
	// InvoiceResponse.IsNoContent() = true
	//
	// ID(s) informado no parâmetro não foi encontrado.
	//
	// # Resposta: 401/500
	//
	// InvoiceResponse = not nil
	//
	// Error = nil
	//
	// InvoiceResponse.IsFailure() = true
	//
	// Para qualquer outra resposta inesperada da API, possuímos o campo InvoiceResponse.Errors preenchido com as informações
	// de erro, o index 0 do slice com campo ErrorResponse.Code retornando a descrição
	// status http (Ex: "401 Unauthorized") e no campo ErrorResponse.Description retornará com o valor
	// "response status code not expected".
	//
	// # Error
	//
	// InvoiceResponse = nil
	//
	// Error = not nil
	//
	// Se o campo ErrorAsaas.Type tiver com valor ErrorTypeValidation quer dizer que não passou pela validação dos
	// parâmetros informados segundo a documentação.
	// Por fim se o campo ErrorAsaas.Type tiver com valor ErrorTypeUnexpected quer dizer que ocorreu um erro inesperado
	// na lib go-asaas.
	//
	// Para obter mais detalhes confira as colunas:
	//
	// ErrorAsaas.Msg (mensagem do erro),
	//
	// ErrorAsaas.File (Arquivo aonde ocorreu o erro),
	//
	// ErrorAsaas.Line (Linha aonde ocorreu o erro)
	//
	// Caso ocorra um erro inesperado por favor report o erro no repositório: https://github.com/GabrielHCataldo/go-asaas
	//
	// # DOCS
	//
	// Recuperar uma nota fiscal: https://docs.asaas.com/reference/recuperar-uma-nota-fiscal
	GetById(ctx context.Context, invoiceId string) (*InvoiceResponse, Error)
	// GetAll (Listar notas fiscais)
	//
	// Diferente da recuperação de uma nota fiscal específica, este método retorna uma lista paginada com todas as notas
	// fiscais para os filtros informados.
	//
	// # Resposta: 200
	//
	// Pageable(InvoiceResponse) = not nil
	//
	// Error = nil
	//
	// Se Pageable.IsSuccess() for true quer dizer que retornaram os dados conforme a documentação.
	// Se Pageable.IsNoContent() for true quer dizer que retornou os dados vazio.
	//
	// Error = nil
	//
	// Pageable.IsNoContent() = true
	//
	// Pageable.Data retornou vazio.
	//
	// # Resposta: 401/500
	//
	// Pageable(InvoiceResponse) = not nil
	//
	// Error = nil
	//
	// Pageable.IsFailure() = true
	//
	// Para qualquer outra resposta inesperada da API, possuímos o campo Pageable.Errors preenchido com
	// as informações de erro, o index 0 do slice com campo ErrorResponse.Code retornando a descrição
	// status http (Ex: "401 Unauthorized") e no campo ErrorResponse.Description retornará com o valor
	// "response status code not expected".
	//
	// # Error
	//
	// Pageable(InvoiceResponse) = nil
	//
	// Error = not nil
	//
	// Se o campo ErrorAsaas.Type tiver com valor ErrorTypeValidation quer dizer que não passou pela validação dos
	// parâmetros informados segundo a documentação.
	// Por fim se o campo ErrorAsaas.Type tiver com valor ErrorTypeUnexpected quer dizer que ocorreu um erro inesperado
	// na lib go-asaas.
	//
	// Para obter mais detalhes confira as colunas:
	//
	// ErrorAsaas.Msg (mensagem do erro),
	//
	// ErrorAsaas.File (Arquivo aonde ocorreu o erro),
	//
	// ErrorAsaas.Line (Linha aonde ocorreu o erro)
	//
	// Caso ocorra um erro inesperado por favor report o erro no repositório: https://github.com/GabrielHCataldo/go-asaas
	//
	// # DOCS
	//
	// Listar notas fiscais: https://docs.asaas.com/reference/listar-notas-fiscais
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

func (i invoice) CancelById(ctx context.Context, invoiceId string) (*InvoiceResponse, Error) {
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
