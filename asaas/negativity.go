package asaas

import (
	"context"
	"fmt"
	"net/http"
	"os"
)

type CreateNegativityRequest struct {
	// Identificador único da cobrança a ser recuperada no Asaas (REQUIRED)
	Payment string `json:"payment,omitempty"`
	// Tipo de negativação (REQUIRED)
	Type NegativityType `json:"type,omitempty"`
	// Descrição do produto ou serviço prestado
	Description string `json:"description,omitempty"`
	// Nome do cliente (REQUIRED)
	CustomerName string `json:"customerName,omitempty"`
	// CPF ou CNPJ do cliente (REQUIRED)
	CustomerCpfCnpj string `json:"customerCpfCnpj,omitempty"`
	// Telefone principal do cliente (REQUIRED)
	CustomerPrimaryPhone string `json:"customerPrimaryPhone,omitempty"`
	// Telefone secundário do cliente
	CustomerSecondaryPhone string `json:"customerSecondaryPhone,omitempty"`
	// CEP do endereço do cliente (REQUIRED)
	CustomerPostalCode string `json:"customerPostalCode,omitempty"`
	// Logradouro do cliente (REQUIRED)
	CustomerAddress string `json:"customerAddress,omitempty"`
	// Número do endereço do cliente (REQUIRED)
	CustomerAddressNumber string `json:"customerAddressNumber,omitempty"`
	// Complemento do endereço do cliente
	CustomerComplement string `json:"customerComplement,omitempty"`
	// Bairro do cliente (REQUIRED)
	CustomerProvince string `json:"customerProvince,omitempty"`
	// Nota fiscal e/ou contrato com firma reconhecida em cartório
	Documents *FileRequest `json:"documents,omitempty"`
}

type GetAllNegativitiesRequest struct {
	// Status da negativação
	Status NegativityStatus `json:"status,omitempty"`
	// Tipo de negativação
	Type NegativityType `json:"type,omitempty"`
	// Filtrar por negativações de uma determinada cobrança
	Payment string `json:"payment,omitempty"`
	// Filtrar a partir da data de solicitação inicial
	RequestStartDate Date `json:"requestStartDate,omitempty"`
	// Filtrar a partir da data de solicitação final
	RequestEndDate Date `json:"requestEndDate,omitempty"`
	// Elemento inicial da lista
	Offset int `json:"offset,omitempty"`
	// Número de elementos da lista (max: 100)
	Limit int `json:"limit,omitempty"`
}

type NegativityResendDocumentsRequest struct {
	// Nota fiscal e/ou contrato com firma reconhecida em cartório (REQUIRED)
	Documents *os.File `json:"documents,omitempty"`
}

type NegativityResponse struct {
	Id                             string           `json:"id,omitempty"`
	Payment                        string           `json:"payment,omitempty"`
	DunningNumber                  int              `json:"dunningNumber,omitempty"`
	Type                           NegativityType   `json:"type,omitempty"`
	Status                         NegativityStatus `json:"status,omitempty"`
	RequestDate                    Date             `json:"requestDate,omitempty"`
	Description                    string           `json:"description,omitempty"`
	Value                          float64          `json:"value,omitempty"`
	FeeValue                       float64          `json:"feeValue,omitempty"`
	NetValue                       float64          `json:"netValue,omitempty"`
	ReceivedInCashFeeValue         float64          `json:"receivedInCashFeeValue,omitempty"`
	CancellationFeeValue           float64          `json:"cancellationFeeValue,omitempty"`
	DenialReason                   string           `json:"denialReason,omitempty"`
	CanBeCancelled                 bool             `json:"canBeCancelled,omitempty"`
	CannotBeCancelledReason        string           `json:"cannotBeCancelledReason,omitempty"`
	IsNecessaryResendDocumentation bool             `json:"isNecessaryResendDocumentation,omitempty"`
	Errors                         []ErrorResponse  `json:"errors,omitempty"`
}

type NegativitySimulateResponse struct {
	Payment         string                             `json:"payment,omitempty"`
	Value           float64                            `json:"value,omitempty"`
	TypeSimulations []NegativityTypeSimulationResponse `json:"typeSimulations,omitempty"`
	Errors          []ErrorResponse                    `json:"errors,omitempty"`
}

type NegativityTypeSimulationResponse struct {
	Type             NegativityType `json:"type,omitempty"`
	IsAllowed        bool           `json:"isAllowed,omitempty"`
	NotAllowedReason string         `json:"notAllowedReason,omitempty"`
	FeeValue         float64        `json:"feeValue,omitempty"`
	NetValue         float64        `json:"netValue,omitempty"`
	StartDate        Date           `json:"startDate,omitempty"`
}

type NegativityHistoryResponse struct {
	Status      NegativityStatus `json:"status,omitempty"`
	Description string           `json:"description,omitempty"`
	EventDate   Date             `json:"eventDate,omitempty"`
}

type NegativityPaymentsResponse struct {
	Value       float64 `json:"value,omitempty"`
	Description string  `json:"description,omitempty"`
	PaymentDate Date    `json:"paymentDate,omitempty"`
}

type ChargesAvailableForDunningResponse struct {
	Payment         string                             `json:"payment,omitempty"`
	Customer        string                             `json:"customer,omitempty"`
	BillingType     BillingType                        `json:"billingType,omitempty"`
	Status          ChargeStatus                       `json:"status,omitempty"`
	Value           float64                            `json:"value,omitempty"`
	DueDate         Date                               `json:"dueDate,omitempty"`
	TypeSimulations []NegativityTypeSimulationResponse `json:"typeSimulations,omitempty"`
}

type negativity struct {
	env         Env
	accessToken string
}

type Negativity interface {
	// Create (Criar uma negativação)
	//
	// Possibilita criar uma negativação a partir de uma cobrança. Para ser possível criar uma negativação, antes é
	// necessário entrar em contato com o seu gerente de conta Asaas e solicitar a permissão de criar negativações via API.
	//
	// # Resposta: 200
	//
	// NegativityResponse = not nil
	//
	// Error = nil
	//
	// NegativityResponse.IsSuccess() = true
	//
	// Possui os valores de resposta de sucesso segunda a documentação.
	//
	// # Resposta: 400/401/500
	//
	// NegativityResponse = not nil
	//
	// Error = nil
	//
	// NegativityResponse.IsFailure() = true
	//
	// Para qualquer outra resposta inesperada da API, possuímos o campo NegativityResponse.Errors preenchido com as informações
	// de erro, sendo 400 retornado da API Asaas com as instruções de requisição conforme a documentação,
	// diferente disso retornará uma mensagem padrão no index 0 do slice com campo ErrorResponse.Code retornando a
	// descrição status http (Ex: "401 Unauthorized") e no campo ErrorResponse.Description retornará com o valor
	// "response status code not expected".
	//
	// # Error
	//
	// NegativityResponse = nil
	//
	// error = not nil
	//
	// Se o parâmetro de retorno error não estiver nil quer dizer que ocorreu um erro inesperado
	// na lib go-asaas.
	//
	// Se isso acontecer por favor report o erro no repositório: https://github.com/GabrielHCataldo/go-asaas
	//
	// # DOCS
	//
	// Criar uma negativação: https://docs.asaas.com/reference/criar-uma-negativacao
	Create(ctx context.Context, body CreateNegativityRequest) (*NegativityResponse, error)
	// Simulate (Simular uma negativação)
	//
	// Possibilita a simulação da taxa cobrada, valor a ser recuperado e data prevista de início da negativação.
	//
	// # Resposta: 200
	//
	// NegativitySimulateResponse = not nil
	//
	// Error = nil
	//
	// NegativitySimulateResponse.IsSuccess() = true
	//
	// Possui os valores de resposta de sucesso segunda a documentação.
	//
	// # Resposta: 404
	//
	// NegativitySimulateResponse = not nil
	//
	// Error = nil
	//
	// NegativitySimulateResponse.IsNoContent() = true
	//
	// ID(s) informado no parâmetro não foi encontrado.
	//
	// # Resposta: 401/500
	//
	// NegativitySimulateResponse = not nil
	//
	// Error = nil
	//
	// NegativitySimulateResponse.IsFailure() = true
	//
	// Para qualquer outra resposta inesperada da API, possuímos o campo NegativitySimulateResponse.Errors preenchido com as informações
	// de erro, o index 0 do slice com campo ErrorResponse.Code retornando a descrição
	// status http (Ex: "401 Unauthorized") e no campo ErrorResponse.Description retornará com o valor
	// "response status code not expected".
	//
	// # Error
	//
	// NegativitySimulateResponse = nil
	//
	// error = not nil
	//
	// Se o parâmetro de retorno error não estiver nil quer dizer que ocorreu um erro inesperado
	// na lib go-asaas.
	//
	// Se isso acontecer por favor report o erro no repositório: https://github.com/GabrielHCataldo/go-asaas
	//
	// # DOCS
	//
	// Simular uma negativação: https://docs.asaas.com/reference/simular-uma-negativacao
	Simulate(ctx context.Context, chargeId string) (*NegativitySimulateResponse, error)
	// ResendDocumentsById (Reenviar documentos)
	//
	// Permite o reenvio dos documentos de uma negativação em caso de negação. Utilize a propriedade
	// NegativityResponse.IsNecessaryResendDocumentation retornado no objeto de negativação para verificar se é preciso o
	// reenvio da documentação.
	//
	// Após o reenvio sua negativação retornará para o status de NegativityStatusAwaitingApproval.
	//
	// # Resposta: 200
	//
	// NegativityResponse = not nil
	//
	// Error = nil
	//
	// NegativityResponse.IsSuccess() = true
	//
	// Possui os valores de resposta de sucesso segunda a documentação.
	//
	// # Resposta: 404
	//
	// NegativityResponse = not nil
	//
	// Error = nil
	//
	// NegativityResponse.IsNoContent() = true
	//
	// ID(s) informado no parâmetro não foi encontrado.
	//
	// # Resposta: 400/401/500
	//
	// NegativityResponse = not nil
	//
	// Error = nil
	//
	// NegativityResponse.IsFailure() = true
	//
	// Para qualquer outra resposta inesperada da API, possuímos o campo NegativityResponse.Errors preenchido com as informações
	// de erro, sendo 400 retornado da API Asaas com as instruções de requisição conforme a documentação,
	// diferente disso retornará uma mensagem padrão no index 0 do slice com campo ErrorResponse.Code retornando a
	// descrição status http (Ex: "401 Unauthorized") e no campo ErrorResponse.Description retornará com o valor
	// "response status code not expected".
	//
	// # Error
	//
	// NegativityResponse = nil
	//
	// error = not nil
	//
	// Se o parâmetro de retorno error não estiver nil quer dizer que ocorreu um erro inesperado
	// na lib go-asaas.
	//
	// Se isso acontecer por favor report o erro no repositório: https://github.com/GabrielHCataldo/go-asaas
	//
	// # DOCS
	//
	// Reenviar documentos: https://docs.asaas.com/reference/reenviar-documentos
	ResendDocumentsById(ctx context.Context, negativityId string, body NegativityResendDocumentsRequest) (
		*NegativityResponse, error)
	// CancelById (Cancelar negativação)
	//
	// Permite o cancelamento de uma negativação. Utilize a propriedade NegativityResponse.CanBeCancelled retornado
	// no objeto de negativação para verificar se a negativação pode ser cancelada.
	//
	// Caso a negativação já tenha sido iniciada, ao solicitar o cancelamento a negativação ficará com o status
	// de NegativityStatusAwaitingCancellation até que seja efetivamente cancelada (NegativityStatusCancelled).
	//
	// # Resposta: 200
	//
	// NegativityResponse = not nil
	//
	// Error = nil
	//
	// NegativityResponse.IsSuccess() = true
	//
	// Possui os valores de resposta de sucesso segunda a documentação.
	//
	// # Resposta: 404
	//
	// NegativityResponse = not nil
	//
	// Error = nil
	//
	// NegativityResponse.IsNoContent() = true
	//
	// ID(s) informado no parâmetro não foi encontrado.
	//
	// # Resposta: 400/401/500
	//
	// NegativityResponse = not nil
	//
	// Error = nil
	//
	// NegativityResponse.IsFailure() = true
	//
	// Para qualquer outra resposta inesperada da API, possuímos o campo NegativityResponse.Errors preenchido com as informações
	// de erro, sendo 400 retornado da API Asaas com as instruções de requisição conforme a documentação,
	// diferente disso retornará uma mensagem padrão no index 0 do slice com campo ErrorResponse.Code retornando a
	// descrição status http (Ex: "401 Unauthorized") e no campo ErrorResponse.Description retornará com o valor
	// "response status code not expected".
	//
	// # Error
	//
	// NegativityResponse = nil
	//
	// error = not nil
	//
	// Se o parâmetro de retorno error não estiver nil quer dizer que ocorreu um erro inesperado
	// na lib go-asaas.
	//
	// Se isso acontecer por favor report o erro no repositório: https://github.com/GabrielHCataldo/go-asaas
	//
	// # DOCS
	//
	//Cancelar negativação: https://docs.asaas.com/reference/cancelar-negativacao
	CancelById(ctx context.Context, negativityId string) (*NegativityResponse, error)
	// GetById (Recuperar uma única cobrança)
	//
	// Para recuperar uma cobrança específica é necessário que você tenha o ID que o Asaas retornou no momento da
	// criação dela.
	//
	// # Resposta: 200
	//
	// ChargeResponse = not nil
	//
	// Error = nil
	//
	// ChargeResponse.IsSuccess() = true
	//
	// Possui os valores de resposta de sucesso segunda a documentação.
	//
	// # Resposta: 404
	//
	// ChargeResponse = not nil
	//
	// Error = nil
	//
	// ChargeResponse.IsNoContent() = true
	//
	// ID(s) informado no parâmetro não foi encontrado.
	//
	// # Resposta: 401/500
	//
	// ChargeResponse = not nil
	//
	// Error = nil
	//
	// ChargeResponse.IsFailure() = true
	//
	// Para qualquer outra resposta inesperada da API, possuímos o campo ChargeResponse.Errors preenchido com as informações
	// de erro, o index 0 do slice com campo ErrorResponse.Code retornando a descrição
	// status http (Ex: "401 Unauthorized") e no campo ErrorResponse.Description retornará com o valor
	// "response status code not expected".
	//
	// # Error
	//
	// ChargeResponse = nil
	//
	// error = not nil
	//
	// Se o parâmetro de retorno error não estiver nil quer dizer que ocorreu um erro inesperado
	// na lib go-asaas.
	//
	// Se isso acontecer por favor report o erro no repositório: https://github.com/GabrielHCataldo/go-asaas
	//
	// # DOCS
	//
	// Recuperar uma única cobrança: https://docs.asaas.com/reference/recuperar-uma-unica-cobranca
	GetById(ctx context.Context, negativityId string) (*NegativityResponse, error)
	// GetAll (Listar negativações)
	//
	// Diferente da recuperação de uma negativação específica, este método retorna uma lista paginada
	// com todas as negativações da conta.
	//
	// # Resposta: 200
	//
	// Pageable(NegativityResponse) = not nil
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
	// Pageable(NegativityResponse) = not nil
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
	// Pageable(NegativityResponse) = nil
	//
	// error = not nil
	//
	// Se o parâmetro de retorno error não estiver nil quer dizer que ocorreu um erro inesperado
	// na lib go-asaas.
	//
	// Se isso acontecer por favor report o erro no repositório: https://github.com/GabrielHCataldo/go-asaas
	//
	// # DOCS
	//
	// Listar cobranças: https://docs.asaas.com/reference/listar-cobrancas
	GetAll(ctx context.Context, filter GetAllNegativitiesRequest) (*Pageable[NegativityResponse], error)
	// GetHistoryById (Listar histórico de eventos)
	//
	// Retorna uma lista paginada com os eventos que ocorreram desde do início da negativação da cobrança.
	//
	// # Resposta: 200
	//
	// Pageable(NegativityHistoryResponse) = not nil
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
	// Pageable(NegativityHistoryResponse) = not nil
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
	// Pageable(NegativityHistoryResponse) = nil
	//
	// error = not nil
	//
	// Se o parâmetro de retorno error não estiver nil quer dizer que ocorreu um erro inesperado
	// na lib go-asaas.
	//
	// Se isso acontecer por favor report o erro no repositório: https://github.com/GabrielHCataldo/go-asaas
	//
	// # DOCS
	//
	// Listas histórico de eventos: https://docs.asaas.com/reference/listas-historico-de-eventos
	GetHistoryById(ctx context.Context, negativityId string, filter PageableDefaultRequest) (
		*Pageable[NegativityHistoryResponse], error)
	// GetPaymentsById (Listar pagamentos recebidos)
	//
	// Retorna uma lista paginada com os pagamentos recebidos por meio da renegociação da dívida.
	//
	// # Resposta: 200
	//
	// Pageable(NegativityPaymentsResponse) = not nil
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
	// Pageable(NegativityPaymentsResponse) = not nil
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
	// Pageable(NegativityPaymentsResponse) = nil
	//
	// error = not nil
	//
	// Se o parâmetro de retorno error não estiver nil quer dizer que ocorreu um erro inesperado
	// na lib go-asaas.
	//
	// Se isso acontecer por favor report o erro no repositório: https://github.com/GabrielHCataldo/go-asaas
	//
	// # DOCS
	//
	// Listar pagamentos recebidos: https://docs.asaas.com/reference/listar-pagamentos-recebidos
	GetPaymentsById(ctx context.Context, negativityId string, filter PageableDefaultRequest) (
		*Pageable[NegativityPaymentsResponse], error)
	// GetChargesAvailableForDunning (Listar cobranças disponíveis para negativação)
	//
	// Retorna uma lista paginada de cobranças possíveis de negativação em conjunto com uma simulação
	// de valores para cada tipo de negativação.
	//
	// # Resposta: 200
	//
	// Pageable(ChargesAvailableForDunningResponse) = not nil
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
	// Pageable(ChargesAvailableForDunningResponse) = not nil
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
	// Pageable(ChargesAvailableForDunningResponse) = nil
	//
	// error = not nil
	//
	// Se o parâmetro de retorno error não estiver nil quer dizer que ocorreu um erro inesperado
	// na lib go-asaas.
	//
	// Se isso acontecer por favor report o erro no repositório: https://github.com/GabrielHCataldo/go-asaas
	//
	// # DOCS
	//
	// Listar cobranças disponíveis para negativação: https://docs.asaas.com/reference/listar-cobrancas-disponiveis-para-negativacao
	GetChargesAvailableForDunning(ctx context.Context, filter PageableDefaultRequest) (
		*Pageable[ChargesAvailableForDunningResponse], error)
}

func NewNegativity(env Env, accessToken string) Negativity {
	logWarning("Negativity service running on", env.String())
	return negativity{
		env:         env,
		accessToken: accessToken,
	}
}

func (n negativity) Create(ctx context.Context, body CreateNegativityRequest) (*NegativityResponse, error) {
	req := NewRequest[NegativityResponse](ctx, n.env, n.accessToken)
	return req.make(http.MethodPost, "/v3/paymentDunnings", body)
}

func (n negativity) Simulate(ctx context.Context, chargeId string) (*NegativitySimulateResponse, error) {
	req := NewRequest[NegativitySimulateResponse](ctx, n.env, n.accessToken)
	return req.make(http.MethodPost, fmt.Sprintf("/v3/paymentDunnings/simulate?payment=%s", chargeId), nil)
}

func (n negativity) ResendDocumentsById(ctx context.Context, negativityId string, body NegativityResendDocumentsRequest) (
	*NegativityResponse, error) {
	req := NewRequest[NegativityResponse](ctx, n.env, n.accessToken)
	return req.makeMultipartForm(http.MethodPost, fmt.Sprintf("/v3/paymentDunnings/%s/documents", negativityId), body)
}

func (n negativity) CancelById(ctx context.Context, negativityId string) (*NegativityResponse, error) {
	req := NewRequest[NegativityResponse](ctx, n.env, n.accessToken)
	return req.make(http.MethodPost, fmt.Sprintf("/v3/paymentDunnings/%s/cancel", negativityId), nil)
}

func (n negativity) GetById(ctx context.Context, negativityId string) (*NegativityResponse, error) {
	req := NewRequest[NegativityResponse](ctx, n.env, n.accessToken)
	return req.make(http.MethodGet, fmt.Sprintf("/v3/paymentDunnings/%s", negativityId), nil)
}

func (n negativity) GetAll(ctx context.Context, filter GetAllNegativitiesRequest) (*Pageable[NegativityResponse], error) {
	req := NewRequest[Pageable[NegativityResponse]](ctx, n.env, n.accessToken)
	return req.make(http.MethodGet, "/v3/paymentDunnings", filter)
}

func (n negativity) GetHistoryById(ctx context.Context, negativityId string, filter PageableDefaultRequest) (
	*Pageable[NegativityHistoryResponse], error) {
	req := NewRequest[Pageable[NegativityHistoryResponse]](ctx, n.env, n.accessToken)
	return req.make(http.MethodGet, fmt.Sprintf("/v3/paymentDunnings/%s/history", negativityId), filter)
}

func (n negativity) GetPaymentsById(ctx context.Context, negativityId string, filter PageableDefaultRequest) (
	*Pageable[NegativityPaymentsResponse], error) {
	req := NewRequest[Pageable[NegativityPaymentsResponse]](ctx, n.env, n.accessToken)
	return req.make(http.MethodGet, fmt.Sprintf("/v3/paymentDunnings/%s/partialPayments", negativityId), filter)
}

func (n negativity) GetChargesAvailableForDunning(ctx context.Context, filter PageableDefaultRequest) (
	*Pageable[ChargesAvailableForDunningResponse], error) {
	req := NewRequest[Pageable[ChargesAvailableForDunningResponse]](ctx, n.env, n.accessToken)
	return req.make(http.MethodGet, "/v3/paymentDunnings/paymentsAvailableForDunning", filter)
}
