package asaas

import (
	"context"
	"fmt"
	"net/http"
)

type CustomerRequest struct {
	// Nome do cliente (REQUIRED)
	Name string `json:"name,omitempty"`
	// CPF ou CNPJ do cliente (REQUIRED)
	CpfCnpj string `json:"cpfCnpj,omitempty"`
	// Email do cliente
	Email string `json:"email,omitempty"`
	// Fone fixo
	Phone string `json:"phone,omitempty"`
	// Fone celular
	MobilePhone string `json:"mobilePhone,omitempty"`
	// Logradouro
	Address string `json:"address,omitempty"`
	// Número do endereço
	AddressNumber string `json:"addressNumber,omitempty"`
	// Complemento do endereço
	Complement string `json:"complement,omitempty"`
	// Bairro
	Province string `json:"province,omitempty"`
	// CEP do endereço
	PostalCode string `json:"postalCode,omitempty"`
	// Identificador do cliente no seu sistema
	ExternalReference string `json:"externalReference,omitempty"`
	// True para desabilitar o envio de notificações de cobrança
	NotificationDisabled bool `json:"notificationDisabled,omitempty"`
	// Emails adicionais para envio de notificações de cobrança separados por ","
	AdditionalEmails string `json:"additionalEmails,omitempty"`
	// Inscrição municipal do cliente
	MunicipalInscription string `json:"municipalInscription,omitempty"`
	// Inscrição estadual do cliente
	StateInscription string `json:"stateInscription,omitempty"`
	// Observações adicionais
	Observations string `json:"observations,omitempty"`
	// Nome do grupo ao qual o cliente pertence
	GroupName string `json:"groupName,omitempty"`
	// Empresa
	Company string `json:"company,omitempty"`
}

type GetAllCustomersRequest struct {
	// Filtrar por nome
	Name string `json:"name,omitempty"`
	// Filtrar por email
	Email string `json:"email,omitempty"`
	// Filtrar por CPF ou CNPJ
	CpfCnpj string `json:"cpfCnpj,omitempty"`
	// Filtrar por grupo
	GroupName string `json:"groupName,omitempty"`
	// Filtrar pelo Identificador do seu sistema
	ExternalReference string `json:"externalReference,omitempty"`
	// Elemento inicial da lista
	Offset int `json:"offset,omitempty"`
	// Número de elementos da lista (max: 100)
	Limit int `json:"limit,omitempty"`
}

type CustomerResponse struct {
	Id                    string          `json:"id,omitempty"`
	Name                  string          `json:"name,omitempty"`
	Email                 string          `json:"email,omitempty"`
	Phone                 string          `json:"phone,omitempty"`
	MobilePhone           string          `json:"mobilePhone,omitempty"`
	Address               string          `json:"address,omitempty"`
	AddressNumber         string          `json:"addressNumber,omitempty"`
	Complement            string          `json:"complement,omitempty"`
	Province              string          `json:"province,omitempty"`
	PostalCode            string          `json:"postalCode,omitempty"`
	CpfCnpj               string          `json:"cpfCnpj,omitempty"`
	PersonType            PersonType      `json:"personType,omitempty"`
	Deleted               bool            `json:"deleted,omitempty"`
	AdditionalEmails      string          `json:"additionalEmails,omitempty"`
	ExternalReference     string          `json:"externalReference,omitempty"`
	NotificationDisabled  bool            `json:"notificationDisabled,omitempty"`
	MunicipalInscription  string          `json:"municipalInscription,omitempty"`
	StateInscription      string          `json:"stateInscription,omitempty"`
	CanDelete             bool            `json:"canDelete,omitempty"`
	CannotBeDeletedReason string          `json:"cannotBeDeletedReason,omitempty"`
	CanEdit               bool            `json:"canEdit,omitempty"`
	CannotEditReason      string          `json:"cannotEditReason,omitempty"`
	ForeignCustomer       bool            `json:"foreignCustomer,omitempty"`
	City                  int             `json:"city,omitempty"`
	State                 string          `json:"state,omitempty"`
	Country               string          `json:"country,omitempty"`
	Observations          string          `json:"observations,omitempty"`
	Errors                []ErrorResponse `json:"errors,omitempty"`
	DateCreated           Date            `json:"dateCreated,omitempty"`
}

type customer struct {
	env         Env
	accessToken string
}

type Customer interface {
	// Create (Criar novo cliente)
	//
	// Possibilita criar um novo cliente. Para ser possível criar uma cobrança, antes é necessário criar o cliente
	// ao qual ela irá pertencer. Você deve utilizar o ID retornado nesta requisição na criação da cobrança.
	//
	// # Resposta: 200
	//
	// CustomerResponse = not nil
	//
	// Error = nil
	//
	// CustomerResponse.IsSuccess() = true
	//
	// Possui os valores de resposta de sucesso segunda a documentação.
	//
	// # Resposta: 400/401/500
	//
	// CustomerResponse = not nil
	//
	// Error = nil
	//
	// CustomerResponse.IsFailure() = true
	//
	// Para qualquer outra resposta inesperada da API, possuímos o campo CustomerResponse.Errors preenchido com as informações
	// de erro, sendo 400 retornado da API Asaas com as instruções de requisição conforme a documentação,
	// diferente disso retornará uma mensagem padrão no index 0 do slice com campo ErrorResponse.Code retornando a
	// descrição status http (Ex: "401 Unauthorized") e no campo ErrorResponse.Description retornará com o valor
	// "response status code not expected".
	//
	// # Error
	//
	// CustomerResponse = nil
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
	// Criar novo cliente: https://docs.asaas.com/reference/criar-novo-cliente
	Create(ctx context.Context, body CustomerRequest) (*CustomerResponse, error)
	// UpdateById (Atualizar cliente existente)
	//
	// Permite atualizar as informações de um cliente já existente.
	//
	// # Resposta: 200
	//
	// CustomerResponse = not nil
	//
	// Error = nil
	//
	// CustomerResponse.IsSuccess() = true
	//
	// Possui os valores de resposta de sucesso segunda a documentação.
	//
	// # Resposta: 404
	//
	// CustomerResponse = not nil
	//
	// Error = nil
	//
	// CustomerResponse.IsNoContent() = true
	//
	// ID(s) informado no parâmetro não foi encontrado.
	//
	// # Resposta: 400/401/500
	//
	// CustomerResponse = not nil
	//
	// Error = nil
	//
	// CustomerResponse.IsFailure() = true
	//
	// Para qualquer outra resposta inesperada da API, possuímos o campo CustomerResponse.Errors preenchido com as informações
	// de erro, sendo 400 retornado da API Asaas com as instruções de requisição conforme a documentação,
	// diferente disso retornará uma mensagem padrão no index 0 do slice com campo ErrorResponse.Code retornando a
	// descrição status http (Ex: "401 Unauthorized") e no campo ErrorResponse.Description retornará com o valor
	// "response status code not expected".
	//
	// # Error
	//
	// CustomerResponse = nil
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
	// Atualizar cliente existente: https://docs.asaas.com/reference/atualizar-cliente-existente
	UpdateById(ctx context.Context, customerId string, body CustomerRequest) (*CustomerResponse, error)
	// DeleteById (Remover cliente)
	//
	// Ao remover um cliente, as assinaturas e cobranças aguardando pagamento ou vencidas pertencentes a ela também
	// são removidas.
	//
	// # Resposta: 200
	//
	// DeleteResponse = not nil
	//
	// Error = nil
	//
	// Se DeleteResponse.IsSuccess() for true quer dizer que foi excluída.
	//
	// Se caso DeleteResponse.IsFailure() for true quer dizer que não foi excluída.
	//
	// # Resposta: 404
	//
	// DeleteResponse = not nil
	//
	// Error = nil
	//
	// DeleteResponse.IsNoContent() = true
	//
	// ID(s) informado no parâmetro não foi encontrado.
	//
	// # Resposta: 400/401/500
	//
	// DeleteResponse = not nil
	//
	// Error = nil
	//
	// DeleteResponse.IsFailure() = true
	//
	// Para qualquer outra resposta inesperada da API, possuímos o campo DeleteResponse.Errors preenchido com as informações
	// de erro, sendo 400 retornado da API Asaas com as instruções de requisição conforme a documentação,
	// diferente disso retornará uma mensagem padrão no index 0 do slice com campo ErrorResponse.Code retornando a
	// descrição status http (Ex: "401 Unauthorized") e no campo ErrorResponse.Description retornará com o valor
	// "response status code not expected".
	//
	// # Error
	//
	// DeleteResponse = nil
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
	// Remover cliente: https://docs.asaas.com/reference/remover-cliente
	DeleteById(ctx context.Context, customerId string) (*DeleteResponse, error)
	// RestoreById (Restaurar cliente removido)
	//
	// # Resposta: 200
	//
	// CustomerResponse = not nil
	//
	// Error = nil
	//
	// CustomerResponse.IsSuccess() = true
	//
	// Possui os valores de resposta de sucesso segunda a documentação.
	//
	// # Resposta: 404
	//
	// CustomerResponse = not nil
	//
	// Error = nil
	//
	// CustomerResponse.IsNoContent() = true
	//
	// ID(s) informado no parâmetro não foi encontrado.
	//
	// # Resposta: 400/401/500
	//
	// CustomerResponse = not nil
	//
	// Error = nil
	//
	// CustomerResponse.IsFailure() = true
	//
	// Para qualquer outra resposta inesperada da API, possuímos o campo CustomerResponse.Errors preenchido com as informações
	// de erro, sendo 400 retornado da API Asaas com as instruções de requisição conforme a documentação,
	// diferente disso retornará uma mensagem padrão no index 0 do slice com campo ErrorResponse.Code retornando a
	// descrição status http (Ex: "401 Unauthorized") e no campo ErrorResponse.Description retornará com o valor
	// "response status code not expected".
	//
	// # Error
	//
	// CustomerResponse = nil
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
	// Restaurar cliente removido: https://docs.asaas.com/reference/restaurar-cliente-removido
	RestoreById(ctx context.Context, customerId string) (*CustomerResponse, error)
	// GetById (Recuperar um único cliente)
	//
	// Para recuperar um cliente específico é necessário que você tenha o ID que o Asaas retornou no momento da criação dele.
	//
	// # Resposta: 200
	//
	// CustomerResponse = not nil
	//
	// Error = nil
	//
	// CustomerResponse.IsSuccess() = true
	//
	// Possui os valores de resposta de sucesso segunda a documentação.
	//
	// # Resposta: 404
	//
	// CustomerResponse = not nil
	//
	// Error = nil
	//
	// CustomerResponse.IsNoContent() = true
	//
	// ID(s) informado no parâmetro não foi encontrado.
	//
	// # Resposta: 401/500
	//
	// CustomerResponse = not nil
	//
	// Error = nil
	//
	// CustomerResponse.IsFailure() = true
	//
	// Para qualquer outra resposta inesperada da API, possuímos o campo CustomerResponse.Errors preenchido com as informações
	// de erro, o index 0 do slice com campo ErrorResponse.Code retornando a descrição
	// status http (Ex: "401 Unauthorized") e no campo ErrorResponse.Description retornará com o valor
	// "response status code not expected".
	//
	// # Error
	//
	// CustomerResponse = nil
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
	// Recuperar um único cliente: https://docs.asaas.com/reference/recuperar-um-unico-cliente
	GetById(ctx context.Context, customerId string) (*CustomerResponse, error)
	// GetAll (Listar clientes)
	//
	// Diferente da recuperação de um cliente específico, este método retorna uma lista paginada com todos os
	// clientes para os filtros informados.
	//
	// # Resposta: 200
	//
	// Pageable(CustomerResponse) = not nil
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
	// Pageable(CustomerResponse) = not nil
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
	// Pageable(CustomerResponse) = nil
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
	// Listar clientes: https://docs.asaas.com/reference/listar-clientes
	GetAll(ctx context.Context, filter GetAllCustomersRequest) (*Pageable[CustomerResponse], error)
}

func NewCustomer(env Env, accessToken string) Customer {
	logWarning("Customer service running on", env.String())
	return customer{
		env:         env,
		accessToken: accessToken,
	}
}

func (c customer) Create(ctx context.Context, body CustomerRequest) (*CustomerResponse, error) {
	req := NewRequest[CustomerResponse](ctx, c.env, c.accessToken)
	return req.make(http.MethodPost, "/v3/customers", body)
}

func (c customer) UpdateById(ctx context.Context, customerId string, body CustomerRequest) (*CustomerResponse, error) {
	req := NewRequest[CustomerResponse](ctx, c.env, c.accessToken)
	return req.make(http.MethodPost, fmt.Sprintf("/v3/customers/%s", customerId), body)
}

func (c customer) DeleteById(ctx context.Context, customerId string) (*DeleteResponse, error) {
	req := NewRequest[DeleteResponse](ctx, c.env, c.accessToken)
	return req.make(http.MethodDelete, fmt.Sprintf("/v3/customers/%s", customerId), nil)
}

func (c customer) RestoreById(ctx context.Context, customerId string) (*CustomerResponse, error) {
	req := NewRequest[CustomerResponse](ctx, c.env, c.accessToken)
	return req.make(http.MethodPost, fmt.Sprintf("/v3/customers/%s", customerId), nil)
}

func (c customer) GetById(ctx context.Context, customerId string) (*CustomerResponse, error) {
	req := NewRequest[CustomerResponse](ctx, c.env, c.accessToken)
	return req.make(http.MethodGet, fmt.Sprintf("/v3/customers/%s", customerId), nil)
}

func (c customer) GetAll(ctx context.Context, filter GetAllCustomersRequest) (*Pageable[CustomerResponse], error) {
	req := NewRequest[Pageable[CustomerResponse]](ctx, c.env, c.accessToken)
	return req.make(http.MethodGet, "/v3/customers", filter)
}
