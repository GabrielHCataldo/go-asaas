package asaas

import (
	"context"
	"net/http"
)

type CreditCardTokenizeRequest struct {
	// Identificador único do cliente no Asaas (REQUIRED)
	Customer string `json:"customer,omitempty"`
	// Informações do cartão de crédito (REQUIRED)
	CreditCard CreditCardRequest `json:"creditCard,omitempty"`
	// Informações do titular do cartão de crédito (REQUIRED)
	CreditCardHolderInfo CreditCardHolderInfoRequest `json:"creditCardHolderInfo,omitempty"`
	// IP de onde o cliente está fazendo a compra. Não deve ser informado o IP do seu servidor. (REQUIRED)
	RemoteIp string `json:"remoteIp,omitempty"`
}

type CreditCardRequest struct {
	// Nome impresso no cartão (REQUIRED)
	HolderName string `json:"holderName,omitempty"`
	// Número do cartão (REQUIRED)
	Number string `json:"number,omitempty"`
	// Mês de expiração (ex: 06) (REQUIRED)
	ExpiryMonth string `json:"expiryMonth,omitempty"`
	// Ano de expiração com 4 dígitos (ex: 2019) (REQUIRED)
	ExpiryYear string `json:"expiryYear,omitempty"`
	// Código de segurança (REQUIRED)
	Ccv string `json:"ccv,omitempty"`
}

type CreditCardHolderInfoRequest struct {
	// Nome do titular do cartão (REQUIRED)
	Name string `json:"name,omitempty"`
	// CPF ou CNPJ do titular do cartão (REQUIRED)
	CpfCnpj string `json:"cpfCnpj,omitempty"`
	// Email do titular do cartão (REQUIRED)
	Email string `json:"email,omitempty"`
	// Fone com DDD do titular do cartão (REQUIRED)
	Phone string `json:"phone,omitempty"`
	// Fone celular do titular do cartão
	MobilePhone string `json:"mobilePhone,omitempty"`
	// CEP do titular do cartão (REQUIRED)
	PostalCode string `json:"postalCode,omitempty"`
	// Número do endereço do titular do cartão (REQUIRED)
	AddressNumber string `json:"addressNumber,omitempty"`
	// Complemento do endereço do titular do cartão
	AddressComplement string `json:"addressComplement,omitempty"`
}

type CreditCardResponse struct {
	CreditCardNumber string `json:"creditCardNumber,omitempty"`
	CreditCardBrand  string `json:"creditCardBrand,omitempty"`
	CreditCardToken  string `json:"creditCardToken,omitempty"`
}

type CreditCardTokenizeResponse struct {
	CreditCardNumber string          `json:"creditCardNumber,omitempty"`
	CreditCardBrand  string          `json:"creditCardBrand,omitempty"`
	CreditCardToken  string          `json:"creditCardToken,omitempty"`
	Errors           []ErrorResponse `json:"errors,omitempty"`
}

type creditCard struct {
	env         Env
	accessToken string
}

type CreditCard interface {
	// Tokenize (Tokenização de cartão de crédito)
	//
	// Essa funcionalidade permite você cobrar de seus clientes recorrentemente sem a necessidade deles informarem todos
	// os dados de cartão de crédito novamente. Tudo isso de forma segura por meio de um token.
	//
	// # Resposta: 200
	//
	// CreditCardTokenizeResponse = not nil
	//
	// Error = nil
	//
	// CreditCardTokenizeResponse.IsSuccess() = true
	//
	// Possui os valores de resposta de sucesso segunda a documentação.
	//
	// # Resposta: 400/401/500
	//
	// CreditCardTokenizeResponse = not nil
	//
	// Error = nil
	//
	// CreditCardTokenizeResponse.IsFailure() = true
	//
	// Para qualquer outra resposta inesperada da API, possuímos o campo CreditCardTokenizeResponse.Errors preenchido com as informações
	// de erro, sendo 400 retornado da API Asaas com as instruções de requisição conforme a documentação,
	// diferente disso retornará uma mensagem padrão no index 0 do slice com campo ErrorResponse.Code retornando a
	// descrição status http (Ex: "401 Unauthorized") e no campo ErrorResponse.Description retornará com o valor
	// "response status code not expected".
	//
	// # Error
	//
	// CreditCardTokenizeResponse = nil
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
	// Tokenização de cartão de crédito: https://docs.asaas.com/reference/tokenizacao-de-cartao-de-credito
	Tokenize(ctx context.Context, body CreditCardTokenizeRequest) (*CreditCardTokenizeResponse, error)
}

func NewCreditCard(env Env, accessToken string) CreditCard {
	logWarning("CreditCard service running on", env.String())
	return creditCard{
		env:         env,
		accessToken: accessToken,
	}
}

func (c creditCard) Tokenize(ctx context.Context, body CreditCardTokenizeRequest) (*CreditCardTokenizeResponse,
	error) {
	req := NewRequest[CreditCardTokenizeResponse](ctx, c.env, c.accessToken)
	return req.make(http.MethodPost, "/v3/creditCard/tokenize", body)
}
