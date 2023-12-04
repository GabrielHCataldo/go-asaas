package asaas

import (
	"context"
	"net/http"
)

type SaveWebhookSettingRequest struct {
	// URL que receberá as informações de sincronização (REQUIRED)
	Url string `json:"url,omitempty"`
	// Email para receber as notificações em caso de erros na fila (REQUIRED)
	Email string `json:"email,omitempty"`
	// Versão utilizada da API. Utilize "3" para a versão v3 (REQUIRED)
	ApiVersion string `json:"apiVersion,omitempty"`
	// Habilitar ou não o webhook
	Enabled *bool `json:"enabled"`
	// Situação da fila de sincronização
	Interrupted *bool `json:"interrupted"`
	// Token de autenticação
	AuthToken string `json:"authToken,omitempty"`
}

type WebhookResponse struct {
	Type        WebhookType     `json:"type,omitempty"`
	Url         string          `json:"url,omitempty"`
	Email       string          `json:"email,omitempty"`
	ApiVersion  string          `json:"apiVersion,omitempty"`
	Enabled     bool            `json:"enabled,omitempty"`
	Interrupted bool            `json:"interrupted,omitempty"`
	AuthToken   string          `json:"authToken,omitempty"`
	Errors      []ErrorResponse `json:"errors,omitempty"`
}

type webhook struct {
	env         Env
	accessToken string
}

type Webhook interface {
	// SaveSetting (Criar ou atualizar configurações de webhook)
	//
	// # Resposta: 200
	//
	// WebhookResponse = not nil
	//
	// Error = nil
	//
	// WebhookResponse.IsSuccess() = true
	//
	// Possui os valores de resposta de sucesso segunda a documentação.
	//
	// # Resposta: 400/401/500
	//
	// WebhookResponse = not nil
	//
	// Error = nil
	//
	// WebhookResponse.IsFailure() = true
	//
	// Para qualquer outra resposta inesperada da API, possuímos o campo WebhookResponse.Errors preenchido com as informações
	// de erro, sendo 400 retornado da API Asaas com as instruções de requisição conforme a documentação,
	// diferente disso retornará uma mensagem padrão no index 0 do slice com campo ErrorResponse.Code retornando a
	// descrição status http (Ex: "401 Unauthorized") e no campo ErrorResponse.Description retornará com o valor
	// "response status code not expected".
	//
	// # Error
	//
	// WebhookResponse = nil
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
	// https://docs.asaas.com/reference/webhook-para-cobrancas-criar-ou-atualizar-configuracoes
	SaveSetting(ctx context.Context, typeWebhook WebhookType, body SaveWebhookSettingRequest) (*WebhookResponse, error)
	// GetSetting (Recuperar configurações)
	//
	// # Resposta: 200
	//
	// WebhookResponse = not nil
	//
	// Error = nil
	//
	// WebhookResponse.IsSuccess() = true
	//
	// Possui os valores de resposta de sucesso segunda a documentação.
	//
	// # Resposta: 404
	//
	// WebhookResponse = not nil
	//
	// Error = nil
	//
	// WebhookResponse.IsNoContent() = true
	//
	// ID(s) informado no parâmetro não foi encontrado.
	//
	// # Resposta: 401/500
	//
	// WebhookResponse = not nil
	//
	// Error = nil
	//
	// WebhookResponse.IsFailure() = true
	//
	// Para qualquer outra resposta inesperada da API, possuímos o campo WebhookResponse.Errors preenchido com as informações
	// de erro, o index 0 do slice com campo ErrorResponse.Code retornando a descrição
	// status http (Ex: "401 Unauthorized") e no campo ErrorResponse.Description retornará com o valor
	// "response status code not expected".
	//
	// # Error
	//
	// WebhookResponse = nil
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
	// https://docs.asaas.com/reference/webhook-para-cobran%C3%A7as-recuperar-configuracoes
	GetSetting(ctx context.Context, typeWebhook WebhookType) (*WebhookResponse, error)
}

func NewWebhook(env Env, accessToken string) Webhook {
	logWarning("Webhook service running on", env.String())
	return webhook{
		env:         env,
		accessToken: accessToken,
	}
}

func (w webhook) SaveSetting(ctx context.Context, typeWebhook WebhookType, body SaveWebhookSettingRequest) (
	*WebhookResponse, error) {
	req := NewRequest[WebhookResponse](ctx, w.env, w.accessToken)
	return req.make(http.MethodPost, "/v3/webhook"+typeWebhook.PathUrl(), body)
}

func (w webhook) GetSetting(ctx context.Context, typeWebhook WebhookType) (*WebhookResponse, error) {
	req := NewRequest[WebhookResponse](ctx, w.env, w.accessToken)
	return req.make(http.MethodGet, "/v3/webhook"+typeWebhook.PathUrl(), nil)
}
