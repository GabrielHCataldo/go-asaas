package asaas

import (
	"context"
	"fmt"
	"net/http"
)

type UpdateNotificationRequest struct {
	// Habilita/desabilita a notificação
	Enabled bool `json:"enabled,omitempty"`
	// Habilita/desabilita o email enviado para você
	EmailEnabledForProvider bool `json:"emailEnabledForProvider,omitempty"`
	// Habilita/desabilita o SMS enviado para você
	SmsEnabledForProvider bool `json:"smsEnabledForProvider,omitempty"`
	// Habilita/desabilita o email enviado para o seu cliente
	EmailEnabledForCustomer bool `json:"emailEnabledForCustomer,omitempty"`
	// Habilita/desabilita o SMS enviado para o seu cliente
	SmsEnabledForCustomer bool `json:"smsEnabledForCustomer,omitempty"`
	// Habilita/desabilita a notificação por voz enviada para o seu cliente
	PhoneCallEnabledForCustomer bool `json:"phoneCallEnabledForCustomer,omitempty"`
	// Habilita/desabilita a mensagem de WhatsApp para seu cliente
	WhatsappEnabledForCustomer bool `json:"whatsappEnabledForCustomer,omitempty"`
	// Especifica quantos dias antes do vencimento a notificação deve se enviada. Válido somente para o evento PAYMENT_DUEDATE_WARNING
	ScheduleOffset int `json:"scheduleOffset,omitempty"`
}

type UpdateManyNotificationsRequest struct {
	// Identificador único do cliente no Asaas (REQUIRED)
	Customer string `json:"customer,omitempty"`
	// Lista de informações das notificações
	Notifications []UpdateManyNotificationRequest `json:"notifications,omitempty"`
}

type UpdateManyNotificationRequest struct {
	// Identificador único da notificação (REQUIRED)
	Id string `json:"id,omitempty"`
	// Habilita/desabilita a notificação
	Enabled bool `json:"enabled,omitempty"`
	// Habilita/desabilita o email enviado para você
	EmailEnabledForProvider bool `json:"emailEnabledForProvider,omitempty"`
	// Habilita/desabilita o SMS enviado para você
	SmsEnabledForProvider bool `json:"smsEnabledForProvider,omitempty"`
	// Habilita/desabilita o email enviado para o seu cliente
	EmailEnabledForCustomer bool `json:"emailEnabledForCustomer,omitempty"`
	// Habilita/desabilita o SMS enviado para o seu cliente
	SmsEnabledForCustomer bool `json:"smsEnabledForCustomer,omitempty"`
	// Habilita/desabilita a notificação por voz enviada para o seu cliente
	PhoneCallEnabledForCustomer bool `json:"phoneCallEnabledForCustomer,omitempty"`
	// Habilita/desabilita a mensagem de WhatsApp para seu cliente
	WhatsappEnabledForCustomer bool `json:"whatsappEnabledForCustomer,omitempty"`
	// Especifica quantos dias antes do vencimento a notificação deve se enviada. Válido somente para o evento PAYMENT_DUEDATE_WARNING
	ScheduleOffset int `json:"scheduleOffset,omitempty"`
}

type NotificationResponse struct {
	Id                          string            `json:"id,omitempty"`
	Customer                    string            `json:"customer,omitempty"`
	Enabled                     bool              `json:"enabled,omitempty"`
	EmailEnabledForProvider     bool              `json:"emailEnabledForProvider,omitempty"`
	SmsEnabledForProvider       bool              `json:"smsEnabledForProvider,omitempty"`
	EmailEnabledForCustomer     bool              `json:"emailEnabledForCustomer,omitempty"`
	SmsEnabledForCustomer       bool              `json:"smsEnabledForCustomer,omitempty"`
	PhoneCallEnabledForCustomer bool              `json:"phoneCallEnabledForCustomer,omitempty"`
	WhatsappEnabledForCustomer  bool              `json:"whatsappEnabledForCustomer,omitempty"`
	Event                       NotificationEvent `json:"event,omitempty"`
	ScheduleOffset              int               `json:"scheduleOffset,omitempty"`
	Errors                      []ErrorResponse   `json:"errors,omitempty"`
	Deleted                     bool              `json:"deleted,omitempty"`
}

type UpdateManyNotificationsResponse struct {
	Notifications []NotificationResponse `json:"notifications,omitempty"`
	Errors        []ErrorResponse        `json:"errors,omitempty"`
}

type notification struct {
	env         Env
	accessToken string
}

type Notification interface {
	// UpdateById (Atualizar notificação existente)
	//
	// # Resposta: 200
	//
	// NotificationResponse = not nil
	//
	// Error = nil
	//
	// NotificationResponse.IsSuccess() = true
	//
	// Possui os valores de resposta de sucesso segunda a documentação.
	//
	// # Resposta: 404
	//
	// NotificationResponse = not nil
	//
	// Error = nil
	//
	// NotificationResponse.IsNoContent() = true
	//
	// ID(s) informado no parâmetro não foi encontrado.
	//
	// # Resposta: 400/401/500
	//
	// NotificationResponse = not nil
	//
	// Error = nil
	//
	// NotificationResponse.IsFailure() = true
	//
	// Para qualquer outra resposta inesperada da API, possuímos o campo NotificationResponse.Errors preenchido com as informações
	// de erro, sendo 400 retornado da API Asaas com as instruções de requisição conforme a documentação,
	// diferente disso retornará uma mensagem padrão no index 0 do slice com campo ErrorResponse.Code retornando a
	// descrição status http (Ex: "401 Unauthorized") e no campo ErrorResponse.Description retornará com o valor
	// "response status code not expected".
	//
	// # Error
	//
	// NotificationResponse = nil
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
	// Atualizar notificação existente: https://docs.asaas.com/reference/atualizar-notificacao-existente
	UpdateById(ctx context.Context, notificationId string, body UpdateNotificationRequest) (*NotificationResponse, Error)
	// UpdateManyByCustomer (Atualizar notificações existentes em lote)
	//
	// É possível personalizar várias notificações, independente do canal de comunicação que utilizar (email, SMS e voz)
	// e quem deve receber a notificação(você e/ou seu cliente) enviando qual o id do cliente e as notificações
	// a serem atualizadas.
	//
	// # Resposta: 200
	//
	// UpdateManyNotificationsResponse = not nil
	//
	// Error = nil
	//
	// UpdateManyNotificationsResponse.IsSuccess() = true
	//
	// Possui os valores de resposta de sucesso segunda a documentação.
	//
	// # Resposta: 400/401/404/500
	//
	// UpdateManyNotificationsResponse = not nil
	//
	// Error = nil
	//
	// UpdateManyNotificationsResponse.IsFailure() = true
	//
	// Para qualquer outra resposta inesperada da API, possuímos o campo UpdateManyNotificationsResponse.Errors
	// preenchido com as informações de erro, sendo 400 retornado da API Asaas com as instruções de requisição
	// conforme a documentação, diferente disso retornará uma mensagem padrão no index 0 do slice com campo
	// ErrorResponse.Code retornando a descrição status http (Ex: "401 Unauthorized") e no campo
	// ErrorResponse.Description retornará com o valor "response status code not expected".
	//
	// # Error
	//
	// UpdateManyNotificationsResponse = nil
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
	// Atualizar notificações existentes em lote: https://docs.asaas.com/reference/atualizar-notificacoes-existentes-em-lote
	UpdateManyByCustomer(ctx context.Context, body UpdateManyNotificationsRequest) (*UpdateManyNotificationsResponse,
		Error)
	// GetAllByCustomer (Recuperar notificações de um cliente)
	//
	// # Resposta: 200
	//
	// Pageable(NotificationResponse) = not nil
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
	// Pageable(NotificationResponse) = not nil
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
	// Pageable(NotificationResponse) = nil
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
	// Recuperar notificações de um cliente: https://docs.asaas.com/reference/recuperar-notificacoes-de-um-cliente
	GetAllByCustomer(ctx context.Context, customerId string) (*Pageable[NotificationResponse], Error)
}

func NewNotification(env Env, accessToken string) Notification {
	logWarning("Notification service running on", env.String())
	return notification{
		env:         env,
		accessToken: accessToken,
	}
}

func (n notification) UpdateById(ctx context.Context, notificationId string, body UpdateNotificationRequest) (
	*NotificationResponse, Error) {
	req := NewRequest[NotificationResponse](ctx, n.env, n.accessToken)
	return req.make(http.MethodPut, fmt.Sprintf("/v3/notifications/%s", notificationId), body)
}

func (n notification) UpdateManyByCustomer(ctx context.Context, body UpdateManyNotificationsRequest) (
	*UpdateManyNotificationsResponse, Error) {
	req := NewRequest[UpdateManyNotificationsResponse](ctx, n.env, n.accessToken)
	return req.make(http.MethodPut, "/v3/notifications/batch", body)
}

func (n notification) GetAllByCustomer(ctx context.Context, customerId string) (*Pageable[NotificationResponse], Error) {
	req := NewRequest[Pageable[NotificationResponse]](ctx, n.env, n.accessToken)
	return req.make(http.MethodGet, fmt.Sprintf("/v3/customers/%s/notifications", customerId), nil)
}
