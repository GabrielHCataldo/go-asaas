package asaas

import (
	"context"
	"fmt"
	"net/http"
)

type GetReportRequest struct {
	// Identificador único do cliente no Asaas (REQUIRED se CpfCnpj não for informado)
	Customer string `json:"customer,omitempty"`
	// CPF ou CNPJ do cliente. Informe este campo caso seu cliente não esteja cadastrado no Asaas (REQUIRED se Customer não for informado)
	CpfCnpj string `json:"cpfCnpj,omitempty"`
	// Estado em que deseja realizar a consulta. (REQUIRED)
	State string `json:"state,omitempty"`
}

type GetAllReportsRequest struct {
	// Filtrar a partir da data de criação
	StartDate *Date `json:"startDate,omitempty"`
	// Filtrar até uma data de criação
	EndDate *Date `json:"endDate,omitempty"`
	// Elemento inicial da lista
	Offset int `json:"offset,omitempty"`
	// Número de elementos da lista (max: 100)
	Limit int `json:"limit,omitempty"`
}

type CreditBureauReportResponse struct {
	Id          string          `json:"id,omitempty"`
	Customer    string          `json:"customer,omitempty"`
	CpfCnpj     string          `json:"cpfCnpj,omitempty"`
	State       string          `json:"state,omitempty"`
	DownloadUrl string          `json:"downloadUrl,omitempty"`
	ReportFile  string          `json:"reportFile,omitempty"`
	DateCreated *Date           `json:"dateCreated,omitempty"`
	Errors      []ErrorResponse `json:"errors,omitempty"`
}

type creditBureau struct {
	env         Env
	accessToken string
}

type CreditBureau interface {
	// GetReport (Realizar consulta)
	//
	// As consultas junto ao Serasa Experian são realizadas no momento da solicitação, para evitar possíveis
	// perdas de conexão, sugerimos um timeout de 30 segundos ou mais.
	//
	// Ao realizar a consulta será retornado o atributo CreditBureauReportResponse.ReportFile contendo o PDF
	// da consulta em Base64, este campo apenas é retornado no momento da criação da consulta, caso precise obter
	// novamente será necessário realizar o download por meio da url presente no campo
	// CreditBureauReportResponse.DownloadUrl.
	//
	// Para realizar a consulta você terá que informar um CPF ou CNPJ e o estado onde deseja realizar a consulta.
	//
	// Caso queira informar um cliente já cadastrado na sua conta Asaas:
	//
	// - Este deverá possuir um CPF ou CNPJ já cadastrado
	//
	// - O envio do estado se torna opcional caso já conste no cadastro do cliente
	//
	// # Resposta: 200
	//
	// CreditBureauReportResponse = not nil
	//
	// Error = nil
	//
	// CreditBureauReportResponse.IsSuccess() = true
	//
	// Possui os valores de resposta de sucesso segunda a documentação.
	//
	// # Resposta: 400/401/500
	//
	// CreditBureauReportResponse = not nil
	//
	// Error = nil
	//
	// CreditBureauReportResponse.IsFailure() = true
	//
	// Para qualquer outra resposta inesperada da API, possuímos o campo CreditBureauReportResponse.Errors preenchido com as informações
	// de erro, sendo 400 retornado da API Asaas com as instruções de requisição conforme a documentação,
	// diferente disso retornará uma mensagem padrão no index 0 do slice com campo ErrorResponse.Code retornando a
	// descrição status http (Ex: "401 Unauthorized") e no campo ErrorResponse.Description retornará com o valor
	// "response status code not expected".
	//
	// # Error
	//
	// CreditBureauReportResponse = nil
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
	// Realizar consulta: https://docs.asaas.com/reference/realizar-consulta
	GetReport(ctx context.Context, body GetReportRequest) (*CreditBureauReportResponse, error)
	// GetReportById (Recuperar uma consulta)
	//
	// Para recuperar uma consulta específica é necessário que você tenha o ID que o Asaas retornou no momento
	// da criação dela.
	//
	// # Resposta: 200
	//
	// CreditBureauReportResponse = not nil
	//
	// Error = nil
	//
	// CreditBureauReportResponse.IsSuccess() = true
	//
	// Possui os valores de resposta de sucesso segunda a documentação.
	//
	// # Resposta: 404
	//
	// CreditBureauReportResponse = not nil
	//
	// Error = nil
	//
	// CreditBureauReportResponse.IsNoContent() = true
	//
	// ID(s) informado no parâmetro não foi encontrado.
	//
	// # Resposta: 401/500
	//
	// CreditBureauReportResponse = not nil
	//
	// Error = nil
	//
	// CreditBureauReportResponse.IsFailure() = true
	//
	// Para qualquer outra resposta inesperada da API, possuímos o campo CreditBureauReportResponse.Errors preenchido com as informações
	// de erro, o index 0 do slice com campo ErrorResponse.Code retornando a descrição
	// status http (Ex: "401 Unauthorized") e no campo ErrorResponse.Description retornará com o valor
	// "response status code not expected".
	//
	// # Error
	//
	// CreditBureauReportResponse = nil
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
	// Recuperar uma consulta: https://docs.asaas.com/reference/recuperar-uma-consulta
	GetReportById(ctx context.Context, creditBureauReportId string) (*CreditBureauReportResponse, error)
	// GetAllReports (Listar consultas)
	//
	// Diferente da recuperação de uma consulta específica, este método retorna uma lista paginada com todas as
	// consultas para os filtros informados.
	//
	// # Resposta: 200
	//
	// Pageable(CreditBureauReportResponse) = not nil
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
	// Pageable(CreditBureauReportResponse) = not nil
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
	// Pageable(CreditBureauReportResponse) = nil
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
	// Listar consultas: https://docs.asaas.com/reference/listar-consultas
	GetAllReports(ctx context.Context, filter GetAllReportsRequest) (*Pageable[CreditBureauReportResponse], error)
}

func NewCreditBureau(env Env, accessToken string) CreditBureau {
	logWarning("CreditBureau service running on", env.String())
	return creditBureau{
		env:         env,
		accessToken: accessToken,
	}
}

func (c creditBureau) GetReport(ctx context.Context, body GetReportRequest) (*CreditBureauReportResponse, error) {
	req := NewRequest[CreditBureauReportResponse](ctx, c.env, c.accessToken)
	return req.make(http.MethodPost, "/v3/creditBureauReport", body)
}

func (c creditBureau) GetReportById(ctx context.Context, creditBureauReportId string) (*CreditBureauReportResponse, error) {
	req := NewRequest[CreditBureauReportResponse](ctx, c.env, c.accessToken)
	return req.make(http.MethodGet, fmt.Sprintf("/v3/creditBureauReport/%s", creditBureauReportId), nil)
}

func (c creditBureau) GetAllReports(ctx context.Context, filter GetAllReportsRequest) (
	*Pageable[CreditBureauReportResponse], error) {
	req := NewRequest[Pageable[CreditBureauReportResponse]](ctx, c.env, c.accessToken)
	return req.make(http.MethodGet, "/v3/creditBureauReport", filter)
}
