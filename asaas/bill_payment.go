package asaas

import (
	"context"
	berrors "errors"
	"fmt"
	"github.com/GabrielHCataldo/go-asaas/internal/util"
	"net/http"
)

type CreateBillPaymentRequest struct {
	// Linha digitável do boleto (REQUIRED)
	IdentificationField string `json:"identificationField,omitempty" validate:"required"`
	// Data de agendamento do pagamento
	ScheduleDate *Date `json:"scheduleDate,omitempty" validate:"omitempty,after_now"`
	// Descrição do pagamento de conta
	Description string `json:"description,omitempty"`
	// Desconto atribuído ao pagamento
	Discount float64 `json:"discount,omitempty" validate:"omitempty,gte=0"`
	// Juros atribuído ao pagamento
	Interest float64 `json:"interest,omitempty" validate:"omitempty,gte=0"`
	// Multa atribuída ao pagamento
	Fine float64 `json:"fine,omitempty" validate:"omitempty,gte=0"`
	// Valor da conta caso seja do tipo que não possui essa informação (Ex: faturas de cartão de crédito)
	Value float64 `json:"value,omitempty" validate:"omitempty,gte=0"`
	// Data de vencimento da conta caso seja do tipo que não possui essa informação
	DueDate *Date `json:"dueDate,omitempty"`
}

type BillPaymentSimulateRequest struct {
	// Linha digitável do boleto (REQUIRED se BarCode não for informado)
	IdentificationField string `json:"identificationField,omitempty"`
	// Código de barras do boleto (REQUIRED se IdentificationField não for informado)
	BarCode string `json:"barCode,omitempty"`
}

type BillPaymentResponse struct {
	Id                    string            `json:"id,omitempty"`
	IdentificationField   string            `json:"identificationField,omitempty"`
	Status                BillPaymentStatus `json:"status,omitempty"`
	Discount              float64           `json:"discount,omitempty"`
	Interest              float64           `json:"interest,omitempty"`
	Fine                  float64           `json:"fine,omitempty"`
	Value                 float64           `json:"value,omitempty"`
	Fee                   float64           `json:"fee,omitempty" `
	Description           string            `json:"description,omitempty"`
	CompanyName           string            `json:"companyName,omitempty"`
	TransactionReceiptUrl string            `json:"transactionReceiptUrl,omitempty"`
	CanBeCancelled        bool              `json:"canBeCancelled,omitempty"`
	FailReasons           string            `json:"failReasons,omitempty"`
	DueDate               *Date             `json:"dueDate,omitempty"`
	ScheduleDate          *Date             `json:"scheduleDate,omitempty"`
	PaymentDate           *Date             `json:"paymentDate,omitempty"`
	Errors                []ErrorResponse   `json:"errors,omitempty"`
}

type BillPaymentSimulateResponse struct {
	MinimumScheduleDate *Date                           `json:"minimumScheduleDate,omitempty"`
	Fee                 float64                         `json:"fee,omitempty" `
	BankSlipInfo        BillPaymentBankSlipInfoResponse `json:"bankSlipInfo,omitempty" `
	Errors              []ErrorResponse                 `json:"errors,omitempty"`
}

type BillPaymentBankSlipInfoResponse struct {
	IdentificationField  string  `json:"identificationField,omitempty"`
	Value                float64 `json:"value,omitempty"`
	DueDate              *Date   `json:"dueDate,omitempty"`
	CompanyName          string  `json:"companyName,omitempty"`
	Bank                 string  `json:"bank,omitempty"`
	BeneficiaryName      string  `json:"beneficiaryName,omitempty"`
	BeneficiaryCpfCnpj   string  `json:"beneficiaryCpfCnpj,omitempty"`
	AllowChangeValue     bool    `json:"allowChangeValue,omitempty"`
	MinValue             float64 `json:"minValue,omitempty"`
	MaxValue             float64 `json:"maxValue,omitempty"`
	DiscountValue        float64 `json:"discountValue,omitempty"`
	InterestValue        float64 `json:"interestValue,omitempty"`
	FineValue            float64 `json:"fineValue,omitempty"`
	OriginalValue        float64 `json:"originalValue,omitempty"`
	TotalDiscountValue   float64 `json:"totalDiscountValue,omitempty"`
	TotalAdditionalValue float64 `json:"totalAdditionalValue,omitempty"`
	IsOverdue            bool    `json:"isOverdue,omitempty"`
}

type billPayment struct {
	env         Env
	accessToken string
}

type BillPayment interface {
	// Create (Criar um pagamento de conta)
	//
	// Permite criar um pagamento de conta por meio da linha digitável do boleto.
	//
	// Para agendar seu pagamento de conta, informe o campo CreateBillPaymentRequest.ScheduleDate com a data desejada para pagamento.
	// Ao escolher um dia não útil, o pagamento será realizado no próximo dia útil. Caso não informado, o pagamento
	// irá ocorrer no dia de vencimento do boleto.
	//
	// Caso solicite o dia atual, é necessário se atentar ao horário limite para pagamentos, que é de até 14h00.
	// Se o pedido for feito depois disso será pago apenas no dia útil seguinte.
	//
	// # Resposta: 200
	//
	// BillPaymentResponse = not nil
	//
	// Error = nil
	//
	// BillPaymentResponse.IsSuccess() = true
	//
	// Possui os valores de resposta de sucesso segunda a documentação.
	//
	// # Resposta: 400/401/500
	//
	// BillPaymentResponse = not nil
	//
	// Error = nil
	//
	// BillPaymentResponse.IsFailure() = true
	//
	// Para qualquer outra resposta inesperada da API, possuímos o campo BillPaymentResponse.Errors preenchido com as informações
	// de erro, sendo 400 retornado da API Asaas com as instruções de requisição conforme a documentação,
	// diferente disso retornará uma mensagem padrão no index 0 do slice com campo ErrorResponse.Code retornando a
	// descrição status http (Ex: "401 Unauthorized") e no campo ErrorResponse.Description retornará com o valor
	// "response status code not expected".
	//
	// # Error
	//
	// BillPaymentResponse = nil
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
	// Criar um pagamento de conta: https://docs.asaas.com/reference/criar-um-pagamento-de-conta
	Create(ctx context.Context, body CreateBillPaymentRequest) (*BillPaymentResponse, Error)
	// Simulate (Simular um pagamento de conta)
	//
	// Permite a simulação de um pagamento de conta por meio da linha digitável ou código de barras do boleto.
	//
	// # Resposta: 200
	//
	// BillPaymentSimulateResponse = not nil
	//
	// Error = nil
	//
	// BillPaymentSimulateResponse.IsSuccess() = true
	//
	// Possui os valores de resposta de sucesso segunda a documentação.
	//
	// # Resposta: 400/401/500
	//
	// BillPaymentSimulateResponse = not nil
	//
	// Error = nil
	//
	// BillPaymentSimulateResponse.IsFailure() = true
	//
	// Para qualquer outra resposta inesperada da API, possuímos o campo BillPaymentSimulateResponse.Errors preenchido com as informações
	// de erro, sendo 400 retornado da API Asaas com as instruções de requisição conforme a documentação,
	// diferente disso retornará uma mensagem padrão no index 0 do slice com campo ErrorResponse.Code retornando a
	// descrição status http (Ex: "401 Unauthorized") e no campo ErrorResponse.Description retornará com o valor
	// "response status code not expected".
	//
	// # Error
	//
	// BillPaymentSimulateResponse = nil
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
	// Simular um pagamento de conta: https://docs.asaas.com/reference/simular-um-pagamento-de-conta
	Simulate(ctx context.Context, body BillPaymentSimulateRequest) (*BillPaymentSimulateResponse, Error)
	// CancelById (Cancelar pagamento de contas)
	//
	// Permite o cancelamento do pagamento de conta. Utilize a propriedade BillPaymentResponse.CanBeCancelled do objeto
	// BillPaymentResponse para verificar se o pagamento de conta pode ser cancelado.
	//
	// Ao ser cancelado o pagamento da conta não será realizado.
	//
	// # Resposta: 200
	//
	// BillPaymentResponse = not nil
	//
	// Error = nil
	//
	// BillPaymentResponse.IsSuccess() = true
	//
	// Possui os valores de resposta de sucesso segunda a documentação.
	//
	// # Resposta: 400/401/500
	//
	// BillPaymentResponse = not nil
	//
	// Error = nil
	//
	// BillPaymentResponse.IsFailure() = true
	//
	// Para qualquer outra resposta inesperada da API, possuímos o campo BillPaymentResponse.Errors preenchido com as informações
	// de erro, sendo 400 retornado da API Asaas com as instruções de requisição conforme a documentação,
	// diferente disso retornará uma mensagem padrão no index 0 do slice com campo ErrorResponse.Code retornando a
	// descrição status http (Ex: "401 Unauthorized") e no campo ErrorResponse.Description retornará com o valor
	// "response status code not expected".
	//
	// # Error
	//
	// BillPaymentResponse = nil
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
	// Cancelar pagamento de contas: https://docs.asaas.com/reference/cancelar-pagamento-de-contas
	CancelById(ctx context.Context, billPaymentId string) (*BillPaymentResponse, Error)
	// GetById (Recuperar um único pagamento de conta)
	//
	// Para recuperar um pagamento de conta específico é necessário que você tenha o ID que o Asaas retornou no
	// momento da sua criação.
	//
	// # Resposta: 200
	//
	// BillPaymentResponse = not nil
	//
	// Error = nil
	//
	// BillPaymentResponse.IsSuccess() = true
	//
	// Possui os valores de resposta de sucesso segunda a documentação.
	//
	// # Resposta: 404
	//
	// BillPaymentResponse = not nil
	//
	// Error = nil
	//
	// BillPaymentResponse.IsNoContent() = true
	//
	// ID(s) informado no parâmetro não foi encontrado.
	//
	// # Resposta: 401/500
	//
	// BillPaymentResponse = not nil
	//
	// Error = nil
	//
	// BillPaymentResponse.IsFailure() = true
	//
	// Para qualquer outra resposta inesperada da API, possuímos o campo BillPaymentResponse.Errors preenchido com as informações
	// de erro, o index 0 do slice com campo ErrorResponse.Code retornando a descrição
	// status http (Ex: "401 Unauthorized") e no campo ErrorResponse.Description retornará com o valor
	// "response status code not expected".
	//
	// # Error
	//
	// BillPaymentResponse = nil
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
	// Recuperar um único pagamento de conta: https://docs.asaas.com/reference/recuperar-um-unico-pagamento-de-conta
	GetById(ctx context.Context, billPaymentId string) (*BillPaymentResponse, Error)
	// GetAll (Listar pagamento de contas)
	//
	// Diferente da recuperação de um pagamento de conta específico, este método retorna uma lista paginada
	// com todos os pagamentos de conta.
	//
	// # Resposta: 200
	//
	// Pageable(BillPaymentResponse) = not nil
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
	// Pageable(BillPaymentResponse) = not nil
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
	// Pageable(BillPaymentResponse) = nil
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
	// Listar pagamento de contas: https://docs.asaas.com/reference/listar-pagamento-de-contas
	GetAll(ctx context.Context, filter PageableDefaultRequest) (*Pageable[BillPaymentResponse], Error)
}

func NewBillPayment(env Env, accessToken string) BillPayment {
	logWarning("BillPayment service running on", env.String())
	return billPayment{
		env:         env,
		accessToken: accessToken,
	}
}

func (b billPayment) Create(ctx context.Context, body CreateBillPaymentRequest) (*BillPaymentResponse, Error) {
	if err := Validate().Struct(body); err != nil {
		return nil, NewError(ErrorTypeValidation, err)
	}
	req := NewRequest[BillPaymentResponse](ctx, b.env, b.accessToken)
	return req.make(http.MethodPost, "/v3/bill", body)
}

func (b billPayment) Simulate(ctx context.Context, body BillPaymentSimulateRequest) (*BillPaymentSimulateResponse, Error) {
	if err := b.validateBodySimulateRequest(&body.IdentificationField, &body.BarCode); err != nil {
		return nil, NewError(ErrorTypeValidation, "inform identificationField or barCode")
	}
	req := NewRequest[BillPaymentSimulateResponse](ctx, b.env, b.accessToken)
	return req.make(http.MethodPost, "/v3/bill/simulate", body)
}

func (b billPayment) CancelById(ctx context.Context, billPaymentId string) (*BillPaymentResponse, Error) {
	req := NewRequest[BillPaymentResponse](ctx, b.env, b.accessToken)
	return req.make(http.MethodPost, fmt.Sprintf("/v3/bill/%s/cancel", billPaymentId), nil)
}

func (b billPayment) GetById(ctx context.Context, billPaymentId string) (*BillPaymentResponse, Error) {
	req := NewRequest[BillPaymentResponse](ctx, b.env, b.accessToken)
	return req.make(http.MethodPost, fmt.Sprintf("/v3/bill/%s", billPaymentId), nil)
}

func (b billPayment) GetAll(ctx context.Context, filter PageableDefaultRequest) (*Pageable[BillPaymentResponse], Error) {
	req := NewRequest[Pageable[BillPaymentResponse]](ctx, b.env, b.accessToken)
	return req.make(http.MethodPost, "/v3/bill", filter)
}

func (b billPayment) validateBodySimulateRequest(identificationField, barCode *string) error {
	if util.IsBlank(barCode) && util.IsBlank(identificationField) {
		return berrors.New("inform barCode or identificationField")
	}
	return nil
}
