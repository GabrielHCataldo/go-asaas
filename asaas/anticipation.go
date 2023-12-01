package asaas

import (
	"context"
	berrors "errors"
	"fmt"
	"github.com/GabrielHCataldo/go-asaas/internal/util"
	"net/http"
	"os"
)

type AnticipationRequest struct {
	// ID da cobrança a ser antecipada (REQUIRED se Installment não for informado)
	Payment string `json:"payment,omitempty"`
	// ID do parcelamento a ser antecipado (REQUIRED se Payment não for informado)
	Installment string `json:"installment,omitempty"`
	// Lista com uma ou mais notas fiscais eletrônicas, ou Contratos de Prestação de Serviços, com firma reconhecida em cartório
	Documents []*os.File `json:"documents,omitempty"`
}

type AnticipationSimulateRequest struct {
	// ID da cobrança a ser antecipada (REQUIRED se Installment não for informado)
	Payment string `json:"payment,omitempty"`
	// ID do parcelamento a ser antecipado (REQUIRED se Payment não for informado)
	Installment string `json:"installment,omitempty"`
}

type AgreementSignRequest struct {
	// Determina se concorda ou discorda
	Agreed bool `json:"agreed,omitempty"`
}

type GetAllAnticipationsRequest struct {
	Payment     string             `json:"payment,omitempty"`
	Installment string             `json:"installment,omitempty"`
	Status      AnticipationStatus `json:"status,omitempty"`
	Offset      int                `json:"offset,omitempty"`
	Limit       int                `json:"limit,omitempty"`
}

type AnticipationResponse struct {
	Id                string             `json:"id,omitempty"`
	Installment       string             `json:"installment,omitempty"`
	Payment           string             `json:"payment,omitempty"`
	Status            AnticipationStatus `json:"status,omitempty"`
	AnticipationDate  *Date              `json:"anticipationDate,omitempty"`
	DueDate           *Date              `json:"dueDate,omitempty"`
	RequestDate       *Date              `json:"requestDate,omitempty"`
	Fee               float64            `json:"fee,omitempty"`
	AnticipationDays  int                `json:"anticipationDays,omitempty"`
	NetValue          float64            `json:"netValue,omitempty"`
	Value             float64            `json:"value,omitempty"`
	TotalValue        float64            `json:"totalValue,omitempty"`
	DenialObservation string             `json:"denialObservation,omitempty"`
	Errors            []ErrorResponse    `json:"errors,omitempty"`
}

type AnticipationLimitsResponse struct {
	CreditCard AnticipationLimitResponse `json:"creditCard,omitempty"`
	BankSlip   AnticipationLimitResponse `json:"bankSlip,omitempty"`
	Errors     []ErrorResponse           `json:"errors,omitempty"`
}

type AnticipationLimitResponse struct {
	Total     float64 `json:"total,omitempty"`
	Available float64 `json:"available,omitempty"`
}

type AnticipationSimulateResponse struct {
	Payment                 string          `json:"payment,omitempty"`
	Installment             string          `json:"installment,omitempty"`
	AnticipationDate        *Date           `json:"anticipationDate,omitempty"`
	DueDate                 *Date           `json:"dueDate,omitempty"`
	Fee                     float64         `json:"fee,omitempty"`
	AnticipationDays        int             `json:"anticipationDays,omitempty"`
	NetValue                float64         `json:"netValue,omitempty"`
	Value                   float64         `json:"value,omitempty"`
	TotalValue              float64         `json:"totalValue,omitempty"`
	IsDocumentationRequired bool            `json:"isDocumentationRequired,omitempty"`
	Errors                  []ErrorResponse `json:"errors,omitempty"`
}

type AgreementSignResponse struct {
	Agreed bool            `json:"agreed,omitempty"`
	Errors []ErrorResponse `json:"errors,omitempty"`
}

type anticipation struct {
	env         Env
	accessToken string
}

type Anticipation interface {
	// Request (Solicitar antecipação)
	//
	// É possível solicitar uma antecipação de um parcelamento ou de uma cobrança avulsa. Em casos de parcelamento,
	// onde a forma de pagamento é por cartão, a antecipação poderá ser feita para o parcelamento completo ou para cada
	// parcela individualmente, e quando a forma de pagamento é por boleto, a antecipação será obrigatoriamente para
	// cada parcela individualmente.
	//
	// Para solicitar uma antecipação de cobrança avulsa, informe o ID da cobrança para o campo AnticipationRequest.Payment.
	// Para solicitar uma antecipação de parcelamentos, informe o ID do parcelamento para o campo AnticipationRequest.Installment.
	//
	// Para determinar se o envio de notas fiscais eletrônicas ou contratos de prestação de serviços é obrigatório,
	// verifique a propriedade AnticipationSimulateResponse.IsDocumentationRequired retornada na func Simulate
	//
	// # Resposta: 200
	//
	// AnticipationResponse = not nil
	//
	// Error = nil
	//
	// AnticipationResponse.IsSuccess() = true
	//
	// Possui os valores de resposta de sucesso segunda a documentação.
	//
	// # Resposta: 400/401/500
	//
	// AnticipationResponse = not nil
	//
	// Error = nil
	//
	// AnticipationResponse.IsFailure() = true
	//
	// Para qualquer outra resposta inesperada da API, possuímos o campo AnticipationResponse.Errors preenchido com as informações
	// de erro, sendo 400 retornado da API Asaas com as instruções de requisição conforme a documentação,
	// diferente disso retornará uma mensagem padrão no index 0 do slice com campo ErrorResponse.Code retornando a
	// descrição status http (Ex: "401 Unauthorized") e no campo ErrorResponse.Description retornará com o valor
	// "response status code not expected".
	//
	// # Error
	//
	// AnticipationResponse = nil
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
	// Solicitar antecipação: https://docs.asaas.com/reference/solicitar-antecipacao
	Request(ctx context.Context, body AnticipationRequest) (*AnticipationResponse, Error)
	// Simulate (Simular antecipação)
	//
	// # Resposta: 200
	//
	// AnticipationSimulateResponse = not nil
	//
	// Error = nil
	//
	// AnticipationSimulateResponse.IsSuccess() = true
	//
	// Possui os valores de resposta de sucesso segunda a documentação.
	//
	// # Resposta: 400/401/500
	//
	// AnticipationSimulateResponse = not nil
	//
	// Error = nil
	//
	// AnticipationSimulateResponse.IsFailure() = true
	//
	// Para qualquer outra resposta inesperada da API, possuímos o campo AnticipationSimulateResponse.Errors preenchido com as informações
	// de erro, sendo 400 retornado da API Asaas com as instruções de requisição conforme a documentação,
	// diferente disso retornará uma mensagem padrão no index 0 do slice com campo ErrorResponse.Code retornando a
	// descrição status http (Ex: "401 Unauthorized") e no campo ErrorResponse.Description retornará com o valor
	// "response status code not expected".
	//
	// # Error
	//
	// AnticipationSimulateResponse = nil
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
	// Simular antecipação: https://docs.asaas.com/reference/simular-antecipacao
	Simulate(ctx context.Context, body AnticipationSimulateRequest) (*AnticipationSimulateResponse, Error)
	// AgreementSign (Concordar ou discordar do Aditivo aos Termos de Uso do ASAAS para contratação do Serviço de Antecipação)
	//
	// Possibilita concordar ou discordar do Aditivo aos Termos de Uso do ASAAS para contratação do Serviço de Antecipação.
	// Para prosseguir com uma antecipação primeiramente solicitamos que concorde com nosso termo de antecipação que
	// pode ser acessado nesse link https://ajuda.asaas.com/pt-BR/articles/1369992-termo-aditivo-antecipacao-de-recebiveis.
	//
	// Para concordar com o termo, informe o campo agreed como true e todas as antecipações aguardando assinatura
	// serão processadas. Para discordar do termo, informe o campo agreed como false e todas as antecipações aguardando
	// assinatura serão canceladas.
	//
	// # Resposta: 200
	//
	// AgreementSignResponse = not nil
	//
	// Error = nil
	//
	// AgreementSignResponse.IsSuccess() = true
	//
	// Possui os valores de resposta de sucesso segunda a documentação.
	//
	// # Resposta: 401/500
	//
	// AgreementSignResponse = not nil
	//
	// Error = nil
	//
	// AgreementSignResponse.IsFailure() = true
	//
	// Para qualquer outra resposta inesperada da API, possuímos o campo AgreementSignResponse.Errors
	// preenchido com as informações de erro, sendo 400 retornado da API Asaas com as instruções de requisição conforme a documentação,
	// diferente disso retornará uma mensagem padrão no index 0 do slice com campo ErrorResponse.Code retornando a
	// descrição status http (Ex: "401 Unauthorized") e no campo ErrorResponse.Description retornará com o valor
	// "response status code not expected".
	//
	// # Error
	//
	// AgreementSignResponse = nil
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
	// https://docs.asaas.com/reference/concordar-ou-discordar-do-aditivo-aos-termos-de-uso-do-asaas-para-contratacao-do-servico-de-antecipacao
	AgreementSign(ctx context.Context, body AgreementSignRequest) (*AgreementSignResponse, Error)
	// GetById (Recuperar uma única antecipação)
	//
	// Para recuperar uma antecipação é necessário que você tenha o ID que o Asaas retornou no momento da solicitação.
	//
	// # Resposta: 200
	//
	// AnticipationResponse = not nil
	//
	// Error = nil
	//
	// AnticipationResponse.IsSuccess() = true
	//
	// Possui os valores de resposta de sucesso segunda a documentação.
	//
	// # Resposta: 404
	//
	// AnticipationResponse = not nil
	//
	// Error = nil
	//
	// AnticipationResponse.IsNoContent() = true
	//
	// ID(s) informado no parâmetro não foi encontrado.
	//
	// # Resposta: 401/500
	//
	// AnticipationResponse = not nil
	//
	// Error = nil
	//
	// AnticipationResponse.IsFailure() = true
	//
	// Para qualquer outra resposta inesperada da API, possuímos o campo AnticipationResponse.Errors preenchido com as informações
	// de erro, o index 0 do slice com campo ErrorResponse.Code retornando a descrição
	// status http (Ex: "401 Unauthorized") e no campo ErrorResponse.Description retornará com o valor
	// "response status code not expected".
	//
	// # Error
	//
	// AnticipationResponse = nil
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
	// Recuperar uma única antecipação: https://docs.asaas.com/reference/recuperar-uma-unica-antecipacao
	GetById(ctx context.Context, anticipationId string) (*AnticipationResponse, Error)
	// GetLimits (Recuperar limites de antecipações)
	//
	// Permite você recuperar os limites de antecipações liberados na conta e também o limite disponível para antecipar.
	//
	// # Resposta: 200
	//
	// AnticipationLimitsResponse = not nil
	//
	// Error = nil
	//
	// AnticipationLimitsResponse.IsSuccess() = true
	//
	// Possui os valores de resposta de sucesso segunda a documentação.
	//
	// # Resposta: 401/500
	//
	// AnticipationLimitsResponse = not nil
	//
	// Error = nil
	//
	// AnticipationLimitsResponse.IsFailure() = true
	//
	// Para qualquer outra resposta inesperada da API, possuímos o campo AnticipationLimitsResponse.Errors preenchido com as informações
	// de erro, o index 0 do slice com campo ErrorResponse.Code retornando a descrição
	// status http (Ex: "401 Unauthorized") e no campo ErrorResponse.Description retornará com o valor
	// "response status code not expected".
	//
	// # Error
	//
	// AnticipationLimitsResponse = nil
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
	// Recuperar limites de antecipações: https://docs.asaas.com/reference/recuperar-limites-de-antecipacoes
	GetLimits(ctx context.Context) (*AnticipationLimitsResponse, Error)
	// GetAll (Listar antecipações)
	//
	// Diferente da recuperação de uma antecipação específica, este método retorna uma lista paginada com todas as
	// antecipações para o filtro informado.
	//
	// # Resposta: 200
	//
	// Pageable(AnticipationResponse) = not nil
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
	// Pageable(AnticipationResponse) = not nil
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
	// Pageable(AnticipationResponse) = nil
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
	// Listar antecipações: https://docs.asaas.com/reference/listar-antecipacoes
	GetAll(ctx context.Context, filter GetAllAnticipationsRequest) (*Pageable[AnticipationResponse], Error)
}

func NewAnticipation(env Env, accessToken string) Anticipation {
	logWarning("Anticipation service running on", env.String())
	return anticipation{
		env:         env,
		accessToken: accessToken,
	}
}

func (a anticipation) Request(ctx context.Context, body AnticipationRequest) (*AnticipationResponse, Error) {
	if err := a.validateBodyRequest(&body.Payment, &body.Installment); err != nil {
		return nil, NewError(ErrorTypeValidation, err)
	}
	req := NewRequest[AnticipationResponse](ctx, a.env, a.accessToken)
	return req.makeMultipartForm(http.MethodPost, "/v3/anticipations", body)
}

func (a anticipation) Simulate(ctx context.Context, body AnticipationSimulateRequest) (*AnticipationSimulateResponse,
	Error) {
	if err := a.validateBodyRequest(&body.Payment, &body.Installment); err != nil {
		return nil, NewError(ErrorTypeValidation, err)
	}
	req := NewRequest[AnticipationSimulateResponse](ctx, a.env, a.accessToken)
	return req.make(http.MethodPost, "/v3/anticipations/simulate", body)
}

func (a anticipation) AgreementSign(ctx context.Context, body AgreementSignRequest) (*AgreementSignResponse, Error) {
	req := NewRequest[AgreementSignResponse](ctx, a.env, a.accessToken)
	return req.make(http.MethodPost, "/v3/anticipations/agreement/sign", body)
}

func (a anticipation) GetById(ctx context.Context, anticipationId string) (*AnticipationResponse, Error) {
	req := NewRequest[AnticipationResponse](ctx, a.env, a.accessToken)
	return req.make(http.MethodGet, fmt.Sprintf("/v3/anticipations/%s", anticipationId), nil)
}

func (a anticipation) GetLimits(ctx context.Context) (*AnticipationLimitsResponse, Error) {
	req := NewRequest[AnticipationLimitsResponse](ctx, a.env, a.accessToken)
	return req.make(http.MethodGet, "/v3/anticipations/limits", nil)
}

func (a anticipation) GetAll(ctx context.Context, filter GetAllAnticipationsRequest) (
	*Pageable[AnticipationResponse], Error) {
	req := NewRequest[Pageable[AnticipationResponse]](ctx, a.env, a.accessToken)
	return req.make(http.MethodGet, "/v3/anticipations", filter)
}

func (a anticipation) validateBodyRequest(payment, installment *string) error {
	if util.IsBlank(payment) && util.IsBlank(installment) {
		return berrors.New("inform payment or installment")
	}
	return nil
}
