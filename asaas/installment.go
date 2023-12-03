package asaas

import (
	"context"
	"fmt"
	"net/http"
)

type InstallmentPaymentBookRequest struct {
	// Filtrar pelo nome da coluna
	Sort SortPaymentBookField `json:"sort,omitempty"`
	// Ordenação da coluna
	Order Order `json:"order,omitempty"`
}

type InstallmentResponse struct {
	Id                    string                         `json:"id,omitempty"`
	Customer              string                         `json:"customer,omitempty"`
	Value                 float64                        `json:"value,omitempty"`
	NetValue              float64                        `json:"netValue,omitempty"`
	PaymentValue          float64                        `json:"paymentValue,omitempty"`
	InstallmentCount      int                            `json:"installmentCount,omitempty"`
	BillingType           BillingType                    `json:"billingType,omitempty"`
	PaymentDate           *Date                          `json:"paymentDate,omitempty"`
	Description           string                         `json:"description,omitempty"`
	ExpirationDay         int                            `json:"expirationDay,omitempty"`
	Deleted               bool                           `json:"deleted,omitempty"`
	PaymentLink           string                         `json:"paymentLink,omitempty"`
	TransactionReceiptUrl string                         `json:"transactionReceiptUrl,omitempty"`
	Chargeback            *InstallmentChargebackResponse `json:"chargeback,omitempty"`
	DateCreated           *Date                          `json:"dateCreated,omitempty"`
	Errors                []ErrorResponse                `json:"errors,omitempty"`
}

type InstallmentChargebackResponse struct {
	Status ChargebackStatus `json:"status,omitempty"`
	Reason ChargebackReason `json:"reason,omitempty"`
}

type UpdateInstallmentSplitsResponse struct {
	Splits []SplitResponse `json:"splits,omitempty"`
	Errors []ErrorResponse `json:"errors,omitempty"`
}

type installment struct {
	env         Env
	accessToken string
}

type Installment interface {
	// UpdateSplitsById (Atualizar splits do parcelamento)
	//
	// # Resposta: 200
	//
	// UpdateInstallmentSplitsResponse = not nil
	//
	// Error = nil
	//
	// UpdateInstallmentSplitsResponse.IsSuccess() = true
	//
	// Possui os valores de resposta de sucesso segunda a documentação.
	//
	// # Resposta: 404
	//
	// UpdateInstallmentSplitsResponse = not nil
	//
	// Error = nil
	//
	// UpdateInstallmentSplitsResponse.IsNoContent() = true
	//
	// ID(s) informado no parâmetro não foi encontrado.
	//
	// # Resposta: 400/401/500
	//
	// UpdateInstallmentSplitsResponse = not nil
	//
	// Error = nil
	//
	// UpdateInstallmentSplitsResponse.IsFailure() = true
	//
	// Para qualquer outra resposta inesperada da API, possuímos o campo UpdateInstallmentSplitsResponse.Errors
	// preenchido com as informações de erro, sendo 400 retornado da API Asaas com as instruções de requisição conforme a documentação,
	// diferente disso retornará uma mensagem padrão no index 0 do slice com campo ErrorResponse.Code retornando a
	// descrição status http (Ex: "401 Unauthorized") e no campo ErrorResponse.Description retornará com o valor
	// "response status code not expected".
	//
	// # Error
	//
	// UpdateInstallmentSplitsResponse = nil
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
	// Atualizar splits do parcelamento: https://docs.asaas.com/reference/atualizar-split-do-parcelamento
	UpdateSplitsById(ctx context.Context, installmentId string, body []SplitRequest) (*UpdateInstallmentSplitsResponse,
		Error)
	// RefundById (Estornar parcelamento)
	//
	// É possível estornar um parcelamento via cartão de crédito recebido ou confirmado. Como já ocorre no processo de
	// estorno de uma cobrança avulsa por cartão de crédito, o saldo correspondente do parcelamento é debitado de sua
	// conta no Asaas e a cobrança é cancelada no cartão do seu cliente. O cancelamento pode levar até 10 dias úteis
	// para aparecer na fatura de seu cliente.
	//
	// # Resposta: 200
	//
	// InstallmentResponse = not nil
	//
	// Error = nil
	//
	// InstallmentResponse.IsSuccess() = true
	//
	// Possui os valores de resposta de sucesso segunda a documentação.
	//
	// # Resposta: 404
	//
	// InstallmentResponse = not nil
	//
	// Error = nil
	//
	// InstallmentResponse.IsNoContent() = true
	//
	// ID(s) informado no parâmetro não foi encontrado.
	//
	// # Resposta: 400/401/500
	//
	// InstallmentResponse = not nil
	//
	// Error = nil
	//
	// InstallmentResponse.IsFailure() = true
	//
	// Para qualquer outra resposta inesperada da API, possuímos o campo InstallmentResponse.Errors preenchido com as informações
	// de erro, sendo 400 retornado da API Asaas com as instruções de requisição conforme a documentação,
	// diferente disso retornará uma mensagem padrão no index 0 do slice com campo ErrorResponse.Code retornando a
	// descrição status http (Ex: "401 Unauthorized") e no campo ErrorResponse.Description retornará com o valor
	// "response status code not expected".
	//
	// # Error
	//
	// InstallmentResponse = nil
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
	// Estornar parcelamento: https://docs.asaas.com/reference/estornar-parcelamento
	RefundById(ctx context.Context, installmentId string) (*InstallmentResponse, Error)
	// DeleteById (Remover parcelamento)
	//
	// Somente é possível remover parcelamentos com cobranças aguardando pagamento ou vencidas e que não possuem
	// processo de antecipação ou de negativação.
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
	// Remover parcelamento: https://docs.asaas.com/reference/remover-parcelamento
	DeleteById(ctx context.Context, installmentId string) (*DeleteResponse, Error)
	// GetById (Recuperar um único parcelamento)
	//
	// O identificador único do parcelamento no Asaas pode ser obtido por meio do atributo installment,
	// retornado no momento da criação de uma cobrança parcelada.
	//
	// # Resposta: 200
	//
	// InstallmentResponse = not nil
	//
	// Error = nil
	//
	// InstallmentResponse.IsSuccess() = true
	//
	// Possui os valores de resposta de sucesso segunda a documentação.
	//
	// # Resposta: 404
	//
	// InstallmentResponse = not nil
	//
	// Error = nil
	//
	// InstallmentResponse.IsNoContent() = true
	//
	// ID(s) informado no parâmetro não foi encontrado.
	//
	// # Resposta: 401/500
	//
	// InstallmentResponse = not nil
	//
	// Error = nil
	//
	// InstallmentResponse.IsFailure() = true
	//
	// Para qualquer outra resposta inesperada da API, possuímos o campo InstallmentResponse.Errors preenchido com as informações
	// de erro, o index 0 do slice com campo ErrorResponse.Code retornando a descrição
	// status http (Ex: "401 Unauthorized") e no campo ErrorResponse.Description retornará com o valor
	// "response status code not expected".
	//
	// # Error
	//
	// InstallmentResponse = nil
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
	// Recuperar um único parcelamento: https://docs.asaas.com/reference/repurar-um-unico-parcelamento
	GetById(ctx context.Context, installmentId string) (*InstallmentResponse, Error)
	// GetPaymentBookById (Gerar carnê de parcelamento)
	//
	// Para gerar os carnês de um parcelamento em formato PDF, é necessário que você tenha o ID do
	// parcelamento retornado pelo Asaas.
	//
	// # Resposta: 200
	//
	// FileTextPlainResponse = not nil
	//
	// Error = nil
	//
	// FileTextPlainResponse.IsSuccess() = true
	//
	// Possui os valores de resposta de sucesso segunda a documentação.
	//
	// # Resposta: 404
	//
	// FileTextPlainResponse = not nil
	//
	// Error = nil
	//
	// FileTextPlainResponse.IsNoContent() = true
	//
	// ID(s) informado no parâmetro não foi encontrado.
	//
	// # Resposta: 400/401/500
	//
	// FileTextPlainResponse = not nil
	//
	// Error = nil
	//
	// FileTextPlainResponse.IsFailure() = true
	//
	// Para qualquer outra resposta inesperada da API, possuímos o campo FileTextPlainResponse.Errors preenchido
	// com as informações de erro, sendo 400 retornado da API Asaas com as instruções de requisição conforme a documentação,
	// diferente disso retornará uma mensagem padrão no index 0 do slice com campo ErrorResponse.Code retornando a
	// descrição status http (Ex: "401 Unauthorized") e no campo ErrorResponse.Description retornará com o valor
	// "response status code not expected".
	//
	// # Error
	//
	// FileTextPlainResponse = nil
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
	// Gerar carnê de parcelamento: https://docs.asaas.com/reference/gerar-carne-de-parcelamento
	GetPaymentBookById(ctx context.Context, installmentId string, filter InstallmentPaymentBookRequest) (
		*FileTextPlainResponse, Error)
	// GetAll (Listar parcelamentos)
	//
	// Diferente da recuperação de um parcelamento específico, este método retorna uma lista paginada com
	// todos seus parcelamentos.
	//
	// # Resposta: 200
	//
	// Pageable(InstallmentResponse) = not nil
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
	// Pageable(InstallmentResponse) = not nil
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
	// Pageable(InstallmentResponse) = nil
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
	// Listar parcelamentos: https://docs.asaas.com/reference/listar-parcelamentos
	GetAll(ctx context.Context, filter PageableDefaultRequest) (*Pageable[InstallmentResponse], Error)
}

func NewInstallment(env Env, accessToken string) Installment {
	logWarning("Installment service running on", env.String())
	return installment{
		env:         env,
		accessToken: accessToken,
	}
}

func (i installment) UpdateSplitsById(ctx context.Context, installmentId string, body []SplitRequest) (
	*UpdateInstallmentSplitsResponse, Error) {
	req := NewRequest[UpdateInstallmentSplitsResponse](ctx, i.env, i.accessToken)
	return req.make(http.MethodPost, fmt.Sprintf("/v3/installments/%s/splits", installmentId), body)
}

func (i installment) RefundById(ctx context.Context, installmentId string) (*InstallmentResponse, Error) {
	req := NewRequest[InstallmentResponse](ctx, i.env, i.accessToken)
	return req.make(http.MethodPost, fmt.Sprintf("/v3/installments/%s/refund", installmentId), nil)
}

func (i installment) DeleteById(ctx context.Context, installmentId string) (*DeleteResponse, Error) {
	req := NewRequest[DeleteResponse](ctx, i.env, i.accessToken)
	return req.make(http.MethodDelete, fmt.Sprintf("/v3/installments/%s", installmentId), nil)
}

func (i installment) GetById(ctx context.Context, installmentId string) (*InstallmentResponse, Error) {
	req := NewRequest[InstallmentResponse](ctx, i.env, i.accessToken)
	return req.make(http.MethodGet, fmt.Sprintf("/v3/installments/%s", installmentId), nil)
}

func (i installment) GetPaymentBookById(ctx context.Context, installmentId string, filter InstallmentPaymentBookRequest) (*FileTextPlainResponse, Error) {
	req := NewRequest[FileTextPlainResponse](ctx, i.env, i.accessToken)
	return req.make(http.MethodGet, fmt.Sprintf("/v3/installments/%s/paymentBook", installmentId), filter)
}

func (i installment) GetAll(ctx context.Context, filter PageableDefaultRequest) (*Pageable[InstallmentResponse], Error) {
	req := NewRequest[Pageable[InstallmentResponse]](ctx, i.env, i.accessToken)
	return req.make(http.MethodGet, "/v3/installments", filter)
}
