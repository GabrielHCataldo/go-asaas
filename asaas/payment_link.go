package asaas

import (
	"context"
	"fmt"
	"net/http"
	"os"
)

type CreatePaymentLinkRequest struct {
	// Nome do link de pagamentos (REQUIRED)
	Name string `json:"name,omitempty"`
	// Descrição do link de pagamentos
	Description string `json:"description,omitempty"`
	// Forma de pagamento (Default: BillingTypeUndefined)
	BillingType BillingType `json:"billingType,omitempty"`
	// Forma de cobrança (Default: ChargeTypeDetached)
	ChargeType ChargeType `json:"chargeType,omitempty"`
	// Data de encerramento, a partir desta data o seu link de pagamentos será desativado automaticamente
	EndDate Date `json:"endDate,omitempty"`
	// Valor do link de pagamentos, caso não informado o pagador poderá informar o quanto deseja pagar
	Value float64 `json:"value,omitempty"`
	// Caso seja possível o pagamento via boleto bancário, define a quantidade de dias úteis que o seu cliente poderá pagar o boleto após gerado
	DueDateLimitDays int `json:"dueDateLimitDays"`
	// Periodicidade da cobrança (REQUIRED se ChargeType = ChargeTypeRecurrent)
	SubscriptionCycle SubscriptionCycle `json:"subscriptionCycle,omitempty"`
	// Quantidade máxima de parcelas que seu cliente poderá parcelar o valor do link de pagamentos (REQUIRED se ChargeType = ChargeTypeInstallment)
	MaxInstallmentCount int `json:"maxInstallmentCount,omitempty"`
	// Define se os clientes cadastrados pelo link de pagamentos terão as notificações habilitadas. Caso não seja informado o valor padrão será true
	NotificationEnabled bool `json:"notificationEnabled,omitempty"`
	// Informações de redirecionamento automático após pagamento do link de pagamento
	Callback *CallbackRequest `json:"callback,omitempty"`
}

type UpdatePaymentLinkRequest struct {
	// Nome do link de pagamentos
	Name string `json:"name,omitempty"`
	// Descrição do link de pagamentos
	Description *string `json:"description,omitempty"`
	// Forma de pagamento
	BillingType BillingType `json:"billingType,omitempty"`
	// Forma de cobrança
	ChargeType ChargeType `json:"chargeType,omitempty"`
	// Data de encerramento, a partir desta data o seu link de pagamentos será desativado automaticamente
	EndDate Date `json:"endDate,omitempty"`
	// Valor do link de pagamentos, caso não informado o pagador poderá informar o quanto deseja pagar
	Value *float64 `json:"value,omitempty"`
	// Caso seja possível o pagamento via boleto bancário, define a quantidade de dias úteis que o seu cliente poderá pagar o boleto após gerado
	DueDateLimitDays int `json:"dueDateLimitDays,omitempty"`
	// Periodicidade da cobrança (REQUIRED se ChargeType = ChargeTypeRecurrent)
	SubscriptionCycle *SubscriptionCycle `json:"subscriptionCycle,omitempty"`
	// Quantidade máxima de parcelas que seu cliente poderá parcelar o valor do link de pagamentos (REQUIRED se ChargeType = ChargeTypeInstallment)
	MaxInstallmentCount int `json:"maxInstallmentCount,omitempty"`
	// Define se os clientes cadastrados pelo link de pagamentos terão as notificações habilitadas. Caso não seja informado o valor padrão será true
	NotificationEnabled *bool `json:"notificationEnabled,omitempty"`
	// Informações de redirecionamento automático após pagamento do link de pagamento
	Callback *CallbackRequest `json:"callback,omitempty"`
}

type GetAllPaymentLinksRequest struct {
	// Filtrar pelo nome do link de pagamento
	Name string `json:"name,omitempty"`
	// Filtrar por link de pagamentos ativos ou desativados
	Active *bool `json:"active,omitempty"`
	// True para recuperar também os links de pagamento removidos
	IncludeDeleted *bool `json:"includeDeleted,omitempty"`
	// Elemento inicial da lista
	Offset int `json:"offset,omitempty"`
	// Número de elementos da lista (max: 100)
	Limit int `json:"limit,omitempty"`
}

type SendImagePaymentLinksRequest struct {
	// True para ser a imagem principal
	Main bool `json:"main,omitempty"`
	// Imagem a ser enviada
	Image *os.File `json:"image,omitempty"`
}

type PaymentLinkResponse struct {
	Id                  string            `json:"id,omitempty"`
	Name                string            `json:"name,omitempty"`
	Url                 string            `json:"url,omitempty"`
	Active              bool              `json:"active,omitempty"`
	BillingType         BillingType       `json:"billingType,omitempty"`
	ChargeType          ChargeType        `json:"chargeType,omitempty"`
	EndDate             Date              `json:"endDate,omitempty"`
	Value               float64           `json:"value,omitempty"`
	SubscriptionCycle   SubscriptionCycle `json:"subscriptionCycle,omitempty"`
	Description         string            `json:"description,omitempty"`
	MaxInstallmentCount int               `json:"maxInstallmentCount,omitempty"`
	DueDateLimitDays    int               `json:"dueDateLimitDays,omitempty"`
	NotificationEnabled bool              `json:"notificationEnabled,omitempty"`
	Errors              []ErrorResponse   `json:"errors,omitempty"`
}

type PaymentLinkImageResponse struct {
	Id     string               `json:"id,omitempty"`
	Main   bool                 `json:"main,omitempty"`
	Image  PaymentImageDataLink `json:"image,omitempty"`
	Errors []ErrorResponse      `json:"errors,omitempty"`
}

type PaymentImageDataLink struct {
	OriginalName string `json:"originalName,omitempty"`
	Size         int    `json:"size,omitempty"`
	Extension    string `json:"extension,omitempty"`
	PreviewUrl   string `json:"previewUrl,omitempty"`
	DownloadUrl  string `json:"downloadUrl,omitempty"`
}

type paymentLink struct {
	env         Env
	accessToken string
}

type PaymentLink interface {
	// Create (Criar um link de pagamentos)
	//
	// Caso o seu cliente possa pagar via boleto bancário, será necessário informar o campo
	// CreatePaymentLinkRequest.DueDateLimitDays, que determina a quantidade de dias úteis para a realização do pagamento
	// após a geração do boleto.
	//
	// Se a forma de cobrança for Parcelamento, informe o campo CreatePaymentLinkRequest.MaxInstallmentCount para determinar
	// o limite máximo de parcelas que seu cliente poderá escolher para realizar o pagamento parcelado.
	//
	// Já se a forma de cobrança ser Assinatura, será necessário o envio do campo CreatePaymentLinkRequest.SubscriptionCycle
	// para determinar a periodicidade da geração das cobranças.
	//
	// Você poderá adicionar imagens ao seu link de pagamentos use SendImageById()
	//
	// # Resposta: 200
	//
	// PaymentLinkResponse = not nil
	//
	// Error = nil
	//
	// PaymentLinkResponse.IsSuccess() = true
	//
	// Possui os valores de resposta de sucesso segunda a documentação.
	//
	// # Resposta: 400/401/500
	//
	// PaymentLinkResponse = not nil
	//
	// Error = nil
	//
	// PaymentLinkResponse.IsFailure() = true
	//
	// Para qualquer outra resposta inesperada da API, possuímos o campo PaymentLinkResponse.Errors preenchido com as informações
	// de erro, sendo 400 retornado da API Asaas com as instruções de requisição conforme a documentação,
	// diferente disso retornará uma mensagem padrão no index 0 do slice com campo ErrorResponse.Code retornando a
	// descrição status http (Ex: "401 Unauthorized") e no campo ErrorResponse.Description retornará com o valor
	// "response status code not expected".
	//
	// # Error
	//
	// PaymentLinkResponse = nil
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
	// Criar um link de pagamentos: https://docs.asaas.com/reference/criar-um-link-de-pagamentos
	Create(ctx context.Context, body CreatePaymentLinkRequest) (*PaymentLinkResponse, error)
	// SendImageById (Adicionar uma imagem a um link de pagamentos)
	//
	// Permite adicionar imagens ao seu link de pagamentos.
	//
	// # Resposta: 200
	//
	// PaymentLinkImageResponse = not nil
	//
	// Error = nil
	//
	// PaymentLinkImageResponse.IsSuccess() = true
	//
	// Possui os valores de resposta de sucesso segunda a documentação.
	//
	// # Resposta: 404
	//
	// PaymentLinkImageResponse = not nil
	//
	// Error = nil
	//
	// PaymentLinkImageResponse.IsNoContent() = true
	//
	// ID(s) informado no parâmetro não foi encontrado.
	//
	// # Resposta: 400/401/500
	//
	// PaymentLinkImageResponse = not nil
	//
	// Error = nil
	//
	// PaymentLinkImageResponse.IsFailure() = true
	//
	// Para qualquer outra resposta inesperada da API, possuímos o campo PaymentLinkImageResponse.Errors preenchido com as informações
	// de erro, sendo 400 retornado da API Asaas com as instruções de requisição conforme a documentação,
	// diferente disso retornará uma mensagem padrão no index 0 do slice com campo ErrorResponse.Code retornando a
	// descrição status http (Ex: "401 Unauthorized") e no campo ErrorResponse.Description retornará com o valor
	// "response status code not expected".
	//
	// # Error
	//
	// PaymentLinkImageResponse = nil
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
	// Adicionar uma imagem a um link de pagamentos: https://docs.asaas.com/reference/adicionar-uma-imagem-a-um-link-de-pagamentos
	SendImageById(ctx context.Context, paymentLinkId string, body SendImagePaymentLinksRequest) (
		*PaymentLinkImageResponse, error)
	// UpdateById (Atualizar um link de pagamentos)
	//
	// Permite a atualização de um link de pagamentos já existente.
	//
	// # Resposta: 200
	//
	// PaymentLinkResponse = not nil
	//
	// Error = nil
	//
	// PaymentLinkResponse.IsSuccess() = true
	//
	// Possui os valores de resposta de sucesso segunda a documentação.
	//
	// # Resposta: 404
	//
	// PaymentLinkResponse = not nil
	//
	// Error = nil
	//
	// PaymentLinkResponse.IsNoContent() = true
	//
	// ID(s) informado no parâmetro não foi encontrado.
	//
	// # Resposta: 400/401/500
	//
	// PaymentLinkResponse = not nil
	//
	// Error = nil
	//
	// PaymentLinkResponse.IsFailure() = true
	//
	// Para qualquer outra resposta inesperada da API, possuímos o campo PaymentLinkResponse.Errors preenchido com as informações
	// de erro, sendo 400 retornado da API Asaas com as instruções de requisição conforme a documentação,
	// diferente disso retornará uma mensagem padrão no index 0 do slice com campo ErrorResponse.Code retornando a
	// descrição status http (Ex: "401 Unauthorized") e no campo ErrorResponse.Description retornará com o valor
	// "response status code not expected".
	//
	// # Error
	//
	// PaymentLinkResponse = nil
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
	// Atualizar um link de pagamentos: https://docs.asaas.com/reference/atualizar-um-link-de-pagamentos
	UpdateById(ctx context.Context, paymentLinkId string, body UpdatePaymentLinkRequest) (*PaymentLinkResponse, error)
	// UpdateImageAsMainById (Definir imagem principal do link de pagamentos)
	//
	// Permite a alteração da imagem principal do seu link de pagamentos.
	//
	// # Resposta: 200
	//
	// PaymentLinkImageResponse = not nil
	//
	// Error = nil
	//
	// PaymentLinkImageResponse.IsSuccess() = true
	//
	// Possui os valores de resposta de sucesso segunda a documentação.
	//
	// # Resposta: 404
	//
	// PaymentLinkImageResponse = not nil
	//
	// Error = nil
	//
	// PaymentLinkImageResponse.IsNoContent() = true
	//
	// ID(s) informado no parâmetro não foi encontrado.
	//
	// # Resposta: 401/500
	//
	// PaymentLinkImageResponse = not nil
	//
	// Error = nil
	//
	// PaymentLinkImageResponse.IsFailure() = true
	//
	// Para qualquer outra resposta inesperada da API, possuímos o campo PaymentLinkImageResponse.Errors preenchido com
	// as informações de erro, o index 0 do slice com campo ErrorResponse.Code retornando a descrição
	// status http (Ex: "401 Unauthorized") e no campo ErrorResponse.Description retornará com o valor
	// "response status code not expected".
	//
	// # Error
	//
	// PaymentLinkImageResponse = nil
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
	// Definir imagem principal do link de pagamentos: https://docs.asaas.com/reference/definir-imagem-principal-do-link-de-pagamentos
	UpdateImageAsMainById(ctx context.Context, paymentLinkId, imageId string) (*PaymentLinkImageResponse, error)
	// DeleteById (Remover um link de pagamentos)
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
	// # Resposta: 401/500
	//
	// DeleteResponse = not nil
	//
	// Error = nil
	//
	// DeleteResponse.IsFailure() = true
	//
	// Para qualquer outra resposta inesperada da API, possuímos o campo PaymentLinkImageResponse.Errors preenchido com
	// as informações de erro, o index 0 do slice com campo ErrorResponse.Code retornando a descrição
	// status http (Ex: "401 Unauthorized") e no campo ErrorResponse.Description retornará com o valor
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
	// Remover um link de pagamentos: https://docs.asaas.com/reference/remover-um-link-de-pagamentos
	DeleteById(ctx context.Context, paymentLinkId string) (*DeleteResponse, error)
	// DeleteImageById (Remover uma imagem do link de pagamentos)
	//
	// Permite a remoção de uma imagem do link de pagamentos.
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
	// # Resposta: 401/500
	//
	// DeleteResponse = not nil
	//
	// Error = nil
	//
	// DeleteResponse.IsFailure() = true
	//
	// Para qualquer outra resposta inesperada da API, possuímos o campo PaymentLinkImageResponse.Errors preenchido com
	// as informações de erro, o index 0 do slice com campo ErrorResponse.Code retornando a descrição
	// status http (Ex: "401 Unauthorized") e no campo ErrorResponse.Description retornará com o valor
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
	// Remover uma imagem do link de pagamentos: https://docs.asaas.com/reference/remover-uma-imagem-do-link-de-pagamentos
	DeleteImageById(ctx context.Context, paymentLinkId, imageId string) (*DeleteResponse, error)
	// RestoreById (Restaurar um link de pagamentos)
	//
	// Possibilita a restauração de um link de pagamentos removido.
	//
	// # Resposta: 200
	//
	// PaymentLinkResponse = not nil
	//
	// Error = nil
	//
	// PaymentLinkResponse.IsSuccess() = true
	//
	// Possui os valores de resposta de sucesso segunda a documentação.
	//
	// # Resposta: 404
	//
	// PaymentLinkResponse = not nil
	//
	// Error = nil
	//
	// PaymentLinkResponse.IsNoContent() = true
	//
	// ID(s) informado no parâmetro não foi encontrado.
	//
	// # Resposta: 400/401/500
	//
	// PaymentLinkResponse = not nil
	//
	// Error = nil
	//
	// PaymentLinkResponse.IsFailure() = true
	//
	// Para qualquer outra resposta inesperada da API, possuímos o campo PaymentLinkResponse.Errors preenchido com as informações
	// de erro, sendo 400 retornado da API Asaas com as instruções de requisição conforme a documentação,
	// diferente disso retornará uma mensagem padrão no index 0 do slice com campo ErrorResponse.Code retornando a
	// descrição status http (Ex: "401 Unauthorized") e no campo ErrorResponse.Description retornará com o valor
	// "response status code not expected".
	//
	// # Error
	//
	// PaymentLinkResponse = nil
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
	// Restaurar um link de pagamentos: https://docs.asaas.com/reference/restaurar-um-link-de-pagamentos
	RestoreById(ctx context.Context, paymentLinkId string) (*PaymentLinkResponse, error)
	// GetById (Recuperar um único link de pagamentos)
	//
	// Para recuperar um link de pagamentos específico é necessário que você tenha o ID que o Asaas retornou
	// no momento da sua criação.
	//
	// # Resposta: 200
	//
	// PaymentLinkResponse = not nil
	//
	// Error = nil
	//
	// PaymentLinkResponse.IsSuccess() = true
	//
	// Possui os valores de resposta de sucesso segunda a documentação.
	//
	// # Resposta: 404
	//
	// PaymentLinkResponse = not nil
	//
	// Error = nil
	//
	// PaymentLinkResponse.IsNoContent() = true
	//
	// ID(s) informado no parâmetro não foi encontrado.
	//
	// # Resposta: 401/500
	//
	// PaymentLinkResponse = not nil
	//
	// Error = nil
	//
	// PaymentLinkResponse.IsFailure() = true
	//
	// Para qualquer outra resposta inesperada da API, possuímos o campo PaymentLinkResponse.Errors preenchido com as informações
	// de erro, o index 0 do slice com campo ErrorResponse.Code retornando a descrição
	// status http (Ex: "401 Unauthorized") e no campo ErrorResponse.Description retornará com o valor
	// "response status code not expected".
	//
	// # Error
	//
	// PaymentLinkResponse = nil
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
	// Recuperar um único link de pagamentos: https://docs.asaas.com/reference/recuperar-um-unico-link-de-pagamentos
	GetById(ctx context.Context, paymentLinkId string) (*PaymentLinkResponse, error)
	// GetImageById (Recuperar uma única imagem do link de pagamentos)
	//
	// Para recuperar a imagem de um link de pagamentos específico é necessário que você tenha o ID que o Asaas retornou
	// no momento da criação dele.
	//
	// # Resposta: 200
	//
	// PaymentLinkImageResponse = not nil
	//
	// Error = nil
	//
	// PaymentLinkImageResponse.IsSuccess() = true
	//
	// Possui os valores de resposta de sucesso segunda a documentação.
	//
	// # Resposta: 404
	//
	// PaymentLinkImageResponse = not nil
	//
	// Error = nil
	//
	// PaymentLinkImageResponse.IsNoContent() = true
	//
	// ID(s) informado no parâmetro não foi encontrado.
	//
	// # Resposta: 401/500
	//
	// PaymentLinkImageResponse = not nil
	//
	// Error = nil
	//
	// PaymentLinkImageResponse.IsFailure() = true
	//
	// Para qualquer outra resposta inesperada da API, possuímos o campo PaymentLinkImageResponse.Errors preenchido com as informações
	// de erro, o index 0 do slice com campo ErrorResponse.Code retornando a descrição
	// status http (Ex: "401 Unauthorized") e no campo ErrorResponse.Description retornará com o valor
	// "response status code not expected".
	//
	// # Error
	//
	// PaymentLinkImageResponse = nil
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
	// Recuperar uma única imagem do link de pagamentos: https://docs.asaas.com/reference/recuperar-uma-unica-imagem-do-link-de-pagamentos
	GetImageById(ctx context.Context, paymentLinkId, imageId string) (*PaymentLinkImageResponse, error)
	// GetAll (Listar links de pagamentos)
	//
	// Diferente da recuperação de um link de pagamentos específico, este método retorna uma lista paginada com todos
	// os links de pagamentos para os filtros informados.
	//
	// # Resposta: 200
	//
	// Pageable(PaymentLinkResponse) = not nil
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
	// Pageable(PaymentLinkResponse) = not nil
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
	// Pageable(PaymentLinkResponse) = nil
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
	// Listar links de pagamentos: https://docs.asaas.com/reference/listar-links-de-pagamentos
	GetAll(ctx context.Context, filter GetAllPaymentLinksRequest) (*Pageable[PaymentLinkResponse], error)
	// GetImagesById (Listar imagens de um link de pagamentos)
	//
	// Este método retorna uma lista paginada com todas as imagens do link de pagamentos informado.
	//
	// # Resposta: 200
	//
	// Pageable(PaymentLinkImageResponse) = not nil
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
	// Pageable(PaymentLinkImageResponse) = not nil
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
	// Pageable(PaymentLinkImageResponse) = nil
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
	// Listar imagens de um link de pagamentos: https://docs.asaas.com/reference/listar-imagens-de-um-link-de-pagamentos
	GetImagesById(ctx context.Context, paymentLinkId string) (*Pageable[PaymentLinkImageResponse], error)
}

func NewPaymentLink(env Env, accessToken string) PaymentLink {
	logWarning("PaymentLink service running on", env.String())
	return paymentLink{
		env:         env,
		accessToken: accessToken,
	}
}

func (p paymentLink) Create(ctx context.Context, body CreatePaymentLinkRequest) (*PaymentLinkResponse, error) {
	req := NewRequest[PaymentLinkResponse](ctx, p.env, p.accessToken)
	return req.make(http.MethodPost, "/v3/paymentLinks", body)
}

func (p paymentLink) SendImageById(ctx context.Context, paymentLinkId string, body SendImagePaymentLinksRequest) (
	*PaymentLinkImageResponse, error) {
	req := NewRequest[PaymentLinkImageResponse](ctx, p.env, p.accessToken)
	return req.makeMultipartForm(http.MethodPost, fmt.Sprintf("/v3/paymentLinks/%s/images", paymentLinkId), body)
}

func (p paymentLink) UpdateById(ctx context.Context, paymentLinkId string, body UpdatePaymentLinkRequest) (
	*PaymentLinkResponse, error) {
	req := NewRequest[PaymentLinkResponse](ctx, p.env, p.accessToken)
	return req.make(http.MethodPut, fmt.Sprintf("/v3/paymentLinks/%s", paymentLinkId), body)
}

func (p paymentLink) UpdateImageAsMainById(ctx context.Context, paymentLinkId, imageId string) (
	*PaymentLinkImageResponse, error) {
	req := NewRequest[PaymentLinkImageResponse](ctx, p.env, p.accessToken)
	return req.make(http.MethodPut, fmt.Sprintf("/v3/paymentLinks/%s/images/%s/setAsMain", paymentLinkId, imageId), nil)
}

func (p paymentLink) RestoreById(ctx context.Context, paymentLinkId string) (*PaymentLinkResponse, error) {
	req := NewRequest[PaymentLinkResponse](ctx, p.env, p.accessToken)
	return req.make(http.MethodPost, fmt.Sprintf("/v3/paymentLinks/%s/restore", paymentLinkId), nil)
}

func (p paymentLink) DeleteById(ctx context.Context, paymentLinkId string) (*DeleteResponse, error) {
	req := NewRequest[DeleteResponse](ctx, p.env, p.accessToken)
	return req.make(http.MethodDelete, fmt.Sprintf("/v3/paymentLinks/%s", paymentLinkId), nil)
}

func (p paymentLink) DeleteImageById(ctx context.Context, paymentLinkId, imageId string) (*DeleteResponse, error) {
	req := NewRequest[DeleteResponse](ctx, p.env, p.accessToken)
	return req.make(http.MethodDelete, fmt.Sprintf("/v3/paymentLinks/%s/images/%s", paymentLinkId, imageId), nil)
}

func (p paymentLink) GetById(ctx context.Context, paymentLinkId string) (*PaymentLinkResponse, error) {
	req := NewRequest[PaymentLinkResponse](ctx, p.env, p.accessToken)
	return req.make(http.MethodGet, fmt.Sprintf("/v3/paymentLinks/%s", paymentLinkId), nil)
}

func (p paymentLink) GetImageById(ctx context.Context, paymentLinkId, imageId string) (*PaymentLinkImageResponse, error) {
	req := NewRequest[PaymentLinkImageResponse](ctx, p.env, p.accessToken)
	return req.make(http.MethodGet, fmt.Sprintf("/v3/paymentLinks/%s/images/%s", paymentLinkId, imageId), nil)
}

func (p paymentLink) GetAll(ctx context.Context, filter GetAllPaymentLinksRequest) (*Pageable[PaymentLinkResponse], error) {
	req := NewRequest[Pageable[PaymentLinkResponse]](ctx, p.env, p.accessToken)
	return req.make(http.MethodGet, "/v3/paymentLinks", filter)
}

func (p paymentLink) GetImagesById(ctx context.Context, paymentLinkId string) (*Pageable[PaymentLinkImageResponse], error) {
	req := NewRequest[Pageable[PaymentLinkImageResponse]](ctx, p.env, p.accessToken)
	return req.make(http.MethodGet, fmt.Sprintf("/v3/paymentLinks/%s/images", paymentLinkId), nil)
}
