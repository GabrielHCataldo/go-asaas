package asaas

import (
	"context"
	"fmt"
	"net/http"
)

type MobilePhoneRechargeRequest struct {
	// Número do celular (REQUIRED)
	PhoneNumber string `json:"phoneNumber,omitempty"`
	// Valor da recarga (REQUIRED)
	Value float64 `json:"value,omitempty"`
}

type MobilePhoneRechargeResponse struct {
	Id             string                    `json:"id,omitempty"`
	PhoneNumber    string                    `json:"phoneNumber,omitempty"`
	Value          float64                   `json:"value,omitempty"`
	Status         MobilePhoneRechargeStatus `json:"status,omitempty"`
	CanBeCancelled bool                      `json:"canBeCancelled,omitempty"`
	OperatorName   string                    `json:"operatorName,omitempty"`
	Errors         []ErrorResponse           `json:"errors,omitempty"`
}

type MobilePhoneProviderResponse struct {
	Name   string                             `json:"name,omitempty"`
	Values []MobilePhoneProviderValueResponse `json:"values,omitempty"`
	Errors []ErrorResponse                    `json:"errors,omitempty"`
}

type MobilePhoneProviderValueResponse struct {
	Name        string  `json:"name,omitempty"`
	Description string  `json:"description,omitempty"`
	Bonus       string  `json:"bonus,omitempty"`
	MinValue    float64 `json:"minValue,omitempty"`
	MaxValue    float64 `json:"maxValue,omitempty"`
}

type mobilePhone struct {
	env         Env
	accessToken string
}

type MobilePhone interface {
	// Recharge (Solicitar recarga)
	//
	// # Resposta: 200
	//
	// MobilePhoneRechargeResponse = not nil
	//
	// Error = nil
	//
	// MobilePhoneRechargeResponse.IsSuccess() = true
	//
	// Possui os valores de resposta de sucesso segunda a documentação.
	//
	// # Resposta: 400/401/500
	//
	// MobilePhoneRechargeResponse = not nil
	//
	// Error = nil
	//
	// MobilePhoneRechargeResponse.IsFailure() = true
	//
	// Para qualquer outra resposta inesperada da API, possuímos o campo MobilePhoneRechargeResponse.Errors preenchido com as informações
	// de erro, sendo 400 retornado da API Asaas com as instruções de requisição conforme a documentação,
	// diferente disso retornará uma mensagem padrão no index 0 do slice com campo ErrorResponse.Code retornando a
	// descrição status http (Ex: "401 Unauthorized") e no campo ErrorResponse.Description retornará com o valor
	// "response status code not expected".
	//
	// # Error
	//
	// MobilePhoneRechargeResponse = nil
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
	// Solicitar recarga: https://docs.asaas.com/reference/solicitar-recarga
	Recharge(ctx context.Context, body MobilePhoneRechargeRequest) (*MobilePhoneRechargeResponse, error)
	// CancelRechargeById (Cancelar uma recarga de celular)
	//
	// Permite o cancelamento da recarga de celular. Utilize a propriedade MobilePhoneRechargeResponse.CanBeCancelled
	// para verificar se a recarga pode ser cancelada.
	//
	// Ao ser cancelado a recarga não será realizada.
	//
	// # Resposta: 200
	//
	// MobilePhoneRechargeResponse = not nil
	//
	// Error = nil
	//
	// MobilePhoneRechargeResponse.IsSuccess() = true
	//
	// Possui os valores de resposta de sucesso segunda a documentação.
	//
	// # Resposta: 404
	//
	// MobilePhoneRechargeResponse = not nil
	//
	// Error = nil
	//
	// MobilePhoneRechargeResponse.IsNoContent() = true
	//
	// ID(s) informado no parâmetro não foi encontrado.
	//
	// # Resposta: 400/401/500
	//
	// MobilePhoneRechargeResponse = not nil
	//
	// Error = nil
	//
	// MobilePhoneRechargeResponse.IsFailure() = true
	//
	// Para qualquer outra resposta inesperada da API, possuímos o campo MobilePhoneRechargeResponse.Errors preenchido com as informações
	// de erro, sendo 400 retornado da API Asaas com as instruções de requisição conforme a documentação,
	// diferente disso retornará uma mensagem padrão no index 0 do slice com campo ErrorResponse.Code retornando a
	// descrição status http (Ex: "401 Unauthorized") e no campo ErrorResponse.Description retornará com o valor
	// "response status code not expected".
	//
	// # Error
	//
	// MobilePhoneRechargeResponse = nil
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
	// Cancelar uma recarga de celular: https://docs.asaas.com/reference/cancelar-uma-recarga-de-celular
	CancelRechargeById(ctx context.Context, rechargeId string) (*MobilePhoneRechargeResponse, error)
	// GetRechargeById (Recuperar uma única recarga de celular)
	//
	// Para recuperar uma recarga de celular em específico é necessário que você tenha o ID que o Asaas retornou no
	// momento da sua criação.
	//
	// # Resposta: 200
	//
	// MobilePhoneRechargeResponse = not nil
	//
	// Error = nil
	//
	// MobilePhoneRechargeResponse.IsSuccess() = true
	//
	// Possui os valores de resposta de sucesso segunda a documentação.
	//
	// # Resposta: 404
	//
	// MobilePhoneRechargeResponse = not nil
	//
	// Error = nil
	//
	// MobilePhoneRechargeResponse.IsNoContent() = true
	//
	// ID(s) informado no parâmetro não foi encontrado.
	//
	// # Resposta: 401/500
	//
	// MobilePhoneRechargeResponse = not nil
	//
	// Error = nil
	//
	// MobilePhoneRechargeResponse.IsFailure() = true
	//
	// Para qualquer outra resposta inesperada da API, possuímos o campo MobilePhoneRechargeResponse.Errors preenchido com as informações
	// de erro, o index 0 do slice com campo ErrorResponse.Code retornando a descrição
	// status http (Ex: "401 Unauthorized") e no campo ErrorResponse.Description retornará com o valor
	// "response status code not expected".
	//
	// # Error
	//
	// MobilePhoneRechargeResponse = nil
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
	// Recuperar uma única recarga de celular: https://docs.asaas.com/reference/recuperar-uma-unica-recarga-de-celular
	GetRechargeById(ctx context.Context, rechargeId string) (*MobilePhoneRechargeResponse, error)
	// GetProviderByPhoneNumber (Buscar qual provedor o número pertence e os valores disponíveis para recarga)
	//
	// # Resposta: 200
	//
	// MobilePhoneProviderResponse = not nil
	//
	// Error = nil
	//
	// MobilePhoneProviderResponse.IsSuccess() = true
	//
	// Possui os valores de resposta de sucesso segunda a documentação.
	//
	// # Resposta: 400/401/500
	//
	// MobilePhoneProviderResponse = not nil
	//
	// Error = nil
	//
	// MobilePhoneProviderResponse.IsFailure() = true
	//
	// Para qualquer outra resposta inesperada da API, possuímos o campo MobilePhoneProviderResponse.Errors preenchido com as informações
	// de erro, sendo 400 retornado da API Asaas com as instruções de requisição conforme a documentação,
	// diferente disso retornará uma mensagem padrão no index 0 do slice com campo ErrorResponse.Code retornando a
	// descrição status http (Ex: "401 Unauthorized") e no campo ErrorResponse.Description retornará com o valor
	// "response status code not expected".
	//
	// # Error
	//
	// MobilePhoneProviderResponse = nil
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
	// https://docs.asaas.com/reference/buscar-qual-provedor-o-numero-pertence-e-os-valores-disponiveis-para-recarga
	GetProviderByPhoneNumber(ctx context.Context, phoneNumber string) (*MobilePhoneProviderResponse, error)
	// GetAllRecharges (Listar recargas de celular)
	//
	// Diferente da recuperação de uma recarga de celular em específico, este método retorna uma lista paginada
	// com todas as recargas.
	//
	// # Resposta: 200
	//
	// Pageable(MobilePhoneRechargeResponse) = not nil
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
	// Pageable(MobilePhoneRechargeResponse) = not nil
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
	// Pageable(MobilePhoneRechargeResponse) = nil
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
	// Listar recargas de celular: https://docs.asaas.com/reference/listar-recargas-de-celular
	GetAllRecharges(ctx context.Context, filter PageableDefaultRequest) (*Pageable[MobilePhoneRechargeResponse], error)
}

func NewMobilePhone(env Env, accessToken string) MobilePhone {
	logWarning("MobilePhone service running on", env.String())
	return mobilePhone{
		env:         env,
		accessToken: accessToken,
	}
}

func (m mobilePhone) Recharge(ctx context.Context, body MobilePhoneRechargeRequest) (*MobilePhoneRechargeResponse, error) {
	req := NewRequest[MobilePhoneRechargeResponse](ctx, m.env, m.accessToken)
	return req.make(http.MethodPost, "/v3/mobilePhoneRecharges", body)
}

func (m mobilePhone) CancelRechargeById(ctx context.Context, rechargeId string) (*MobilePhoneRechargeResponse, error) {
	req := NewRequest[MobilePhoneRechargeResponse](ctx, m.env, m.accessToken)
	return req.make(http.MethodPost, fmt.Sprintf("/v3/mobilePhoneRecharges/%s/cancel", rechargeId), nil)
}

func (m mobilePhone) GetRechargeById(ctx context.Context, rechargeId string) (*MobilePhoneRechargeResponse, error) {
	req := NewRequest[MobilePhoneRechargeResponse](ctx, m.env, m.accessToken)
	return req.make(http.MethodGet, fmt.Sprintf("/v3/mobilePhoneRecharges/%s", rechargeId), nil)
}

func (m mobilePhone) GetProviderByPhoneNumber(ctx context.Context, phoneNumber string) (*MobilePhoneProviderResponse,
	error) {
	req := NewRequest[MobilePhoneProviderResponse](ctx, m.env, m.accessToken)
	return req.make(http.MethodGet, fmt.Sprintf("/v3/mobilePhoneRecharges/%s/provider", phoneNumber), nil)
}

func (m mobilePhone) GetAllRecharges(ctx context.Context, filter PageableDefaultRequest) (
	*Pageable[MobilePhoneRechargeResponse], error) {
	req := NewRequest[Pageable[MobilePhoneRechargeResponse]](ctx, m.env, m.accessToken)
	return req.make(http.MethodGet, "/v3/mobilePhoneRecharges", filter)
}
