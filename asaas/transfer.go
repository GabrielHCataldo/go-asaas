package asaas

import (
	"context"
	"fmt"
	"net/http"
)

type TransferToBankRequest struct {
	// Valor a ser transferido (REQUIRED)
	Value float64 `json:"value,omitempty"`
	// Informe os dados da conta caso seja uma transferência para conta bancária (REQUIRED)
	BankAccount       BackAccountRequest    `json:"bankAccount,omitempty"`
	OperationType     TransferOperationType `json:"operationType,omitempty"`
	PixAddressKey     string                `json:"pixAddressKey,omitempty"`
	PixAddressKeyType PixKeyType            `json:"pixAddressKeyType,omitempty"`
	Description       string                `json:"description,omitempty"`
	ScheduleDate      Date                  `json:"scheduleDate,omitempty"`
}

type TransferToAssasRequest struct {
	Value    float64 `json:"value,omitempty"`
	WalletId string  `json:"walletId,omitempty"`
}

type BackAccountRequest struct {
	// Informações da instituição bancária
	Bank BankRequest `json:"bank,omitempty"`
	// Nome da conta bancária
	AccountName string `json:"accountName,omitempty"`
	// Nome do proprietário da conta bancária (REQUIRED)
	OwnerName string `json:"ownerName,omitempty"`
	// Data de nascimento do proprietário da conta. Somente quando a conta bancária não pertencer ao mesmo CPF ou CNPJ da conta Asaas.
	OwnerBirthDate Date `json:"ownerBirthDate,omitempty"`
	// CPF ou CNPJ do proprietário da conta bancária (REQUIRED)
	CpfCnpj string `json:"cpfCnpj,omitempty"`
	// Número da agência sem dígito (REQUIRED)
	Agency string `json:"agency,omitempty"`
	// Número da conta bancária sem dígito (REQUIRED
	Account string `json:"account,omitempty"`
	// Dígito da conta bancária (REQUIRED
	AccountDigit string `json:"accountDigit,omitempty"`
	// Tipo da conta (REQUIRED)
	BankAccountType BankAccountType `json:"bankAccountType,omitempty"`
	// Identificador no Sistema de Pagamentos Brasileiro
	Ispb string `json:"ispb,omitempty"`
}

type BankRequest struct {
	// Código de compensação do banco no sistema bancário (REQUIRED)
	Code string `json:"code,omitempty"`
}

type GetAllTransfersRequest struct {
	// Filtrar pela data de criação inicial
	DateCreatedGe Date `json:"dateCreated[ge],omitempty"`
	// Filtrar pela data de criação final
	DateCreatedLe Date `json:"dateCreated[le],omitempty"`
	// Filtrar pela data inicial de efetivação de transferência
	TransferDateGe Date `json:"transferDate[ge],omitempty"`
	// Filtrar pela data final de efetivação de transferência
	TransferDateLe Date `json:"transferDate[le],omitempty"`
	// Filtrar por tipo da transferência
	Type TransferType `json:"type,omitempty"`
}

type TransferResponse struct {
	Id                    string                `json:"id,omitempty"`
	Type                  TransferType          `json:"type,omitempty"`
	Status                TransferStatus        `json:"status,omitempty"`
	Value                 float64               `json:"value,omitempty"`
	NetValue              float64               `json:"netValue,omitempty"`
	TransferFee           float64               `json:"transferFee,omitempty"`
	EffectiveDate         Date                  `json:"effectiveDate,omitempty"`
	EndToEndIdentifier    string                `json:"endToEndIdentifier,omitempty"`
	ScheduleDate          Date                  `json:"scheduleDate,omitempty"`
	Authorized            bool                  `json:"authorized,omitempty"`
	FailReason            string                `json:"failReason,omitempty"`
	WalletId              string                `json:"walletId,omitempty"`
	BackAccount           *BackAccountResponse  `json:"backAccount,omitempty"`
	TransactionReceiptUrl string                `json:"transactionReceiptUrl,omitempty"`
	OperationType         TransferOperationType `json:"operationType,omitempty"`
	Description           string                `json:"description,omitempty"`
	Errors                []ErrorResponse       `json:"errors,omitempty"`
	DateCreated           Date                  `json:"dateCreated,omitempty"`
}

type BackAccountResponse struct {
	Bank           BankResponse `json:"bank,omitempty"`
	AccountName    string       `json:"accountName,omitempty"`
	OwnerName      string       `json:"ownerName,omitempty"`
	OwnerBirthDate Date         `json:"ownerBirthDate,omitempty"`
	CpfCnpj        string       `json:"cpfCnpj,omitempty"`
	Agency         string       `json:"agency,omitempty"`
	Account        string       `json:"account,omitempty"`
	AccountDigit   string       `json:"accountDigit,omitempty"`
	PixAddressKey  string       `json:"pixAddressKey,omitempty"`
}

type BankResponse struct {
	Ispb string `json:"ispb,omitempty"`
	Code string `json:"code,omitempty"`
	Name string `json:"name,omitempty"`
}

type transfer struct {
	env         Env
	accessToken string
}

type Transfer interface {
	// TransferToBank (Transferir para conta de outra Instituição ou chave pix)
	//
	// Com este endpoint você pode fazer uma transferência para conta bancária ou chave pix.
	//
	// # Resposta: 200
	//
	// TransferResponse = not nil
	//
	// Error = nil
	//
	// TransferResponse.IsSuccess() = true
	//
	// Possui os valores de resposta de sucesso segunda a documentação.
	//
	// # Resposta: 400/401/500
	//
	// TransferResponse = not nil
	//
	// Error = nil
	//
	// TransferResponse.IsFailure() = true
	//
	// Para qualquer outra resposta inesperada da API, possuímos o campo TransferResponse.Errors preenchido com as informações
	// de erro, sendo 400 retornado da API Asaas com as instruções de requisição conforme a documentação,
	// diferente disso retornará uma mensagem padrão no index 0 do slice com campo ErrorResponse.Code retornando a
	// descrição status http (Ex: "401 Unauthorized") e no campo ErrorResponse.Description retornará com o valor
	// "response status code not expected".
	//
	// # Error
	//
	// TransferResponse = nil
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
	// Transferir para conta de outra Instituição ou chave pix: https://docs.asaas.com/reference/transferir-para-conta-de-outra-instituicao-ou-chave-pix
	TransferToBank(ctx context.Context, body TransferToBankRequest) (*TransferResponse, error)
	// TransferToAsaas (Transferir para conta Asaas)
	//
	// Só é possível fazer transferência entre contas Asaas para contas que possuam vínculo entre si,
	// como conta-pai e conta-filha, ou duas contas-filhas de mesma conta-pai.
	//
	// # Resposta: 200
	//
	// TransferResponse = not nil
	//
	// Error = nil
	//
	// TransferResponse.IsSuccess() = true
	//
	// Possui os valores de resposta de sucesso segunda a documentação.
	//
	// # Resposta: 400/401/500
	//
	// TransferResponse = not nil
	//
	// Error = nil
	//
	// TransferResponse.IsFailure() = true
	//
	// Para qualquer outra resposta inesperada da API, possuímos o campo TransferResponse.Errors preenchido com as informações
	// de erro, sendo 400 retornado da API Asaas com as instruções de requisição conforme a documentação,
	// diferente disso retornará uma mensagem padrão no index 0 do slice com campo ErrorResponse.Code retornando a
	// descrição status http (Ex: "401 Unauthorized") e no campo ErrorResponse.Description retornará com o valor
	// "response status code not expected".
	//
	// # Error
	//
	// TransferResponse = nil
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
	// Transferir para conta Asaas: https://docs.asaas.com/reference/transferir-para-conta-asaas
	TransferToAsaas(ctx context.Context, body TransferToAssasRequest) (*TransferResponse, error)
	// CancelById (Cancelar uma transferência)
	//
	// # Resposta: 200
	//
	// TransferResponse = not nil
	//
	// Error = nil
	//
	// TransferResponse.IsSuccess() = true
	//
	// Possui os valores de resposta de sucesso segunda a documentação.
	//
	// # Resposta: 404
	//
	// TransferResponse = not nil
	//
	// Error = nil
	//
	// TransferResponse.IsNoContent() = true
	//
	// ID(s) informado no parâmetro não foi encontrado.
	//
	// # Resposta: 400/401/500
	//
	// TransferResponse = not nil
	//
	// Error = nil
	//
	// TransferResponse.IsFailure() = true
	//
	// Para qualquer outra resposta inesperada da API, possuímos o campo TransferResponse.Errors preenchido com as informações
	// de erro, sendo 400 retornado da API Asaas com as instruções de requisição conforme a documentação,
	// diferente disso retornará uma mensagem padrão no index 0 do slice com campo ErrorResponse.Code retornando a
	// descrição status http (Ex: "401 Unauthorized") e no campo ErrorResponse.Description retornará com o valor
	// "response status code not expected".
	//
	// # Error
	//
	// TransferResponse = nil
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
	// Cancelar uma transferência: https://docs.asaas.com/reference/cancelar-uma-transferencia
	CancelById(ctx context.Context, transferId string) (*TransferResponse, error)
	// GetById (Recuperar uma única cobrança)
	//
	// Para recuperar uma transferência específica é necessário que você tenha o ID que o Asaas retornou no
	// momento da sua criação.
	//
	// # Resposta: 200
	//
	// TransferResponse = not nil
	//
	// Error = nil
	//
	// TransferResponse.IsSuccess() = true
	//
	// Possui os valores de resposta de sucesso segunda a documentação.
	//
	// # Resposta: 404
	//
	// TransferResponse = not nil
	//
	// Error = nil
	//
	// TransferResponse.IsNoContent() = true
	//
	// ID(s) informado no parâmetro não foi encontrado.
	//
	// # Resposta: 401/500
	//
	// TransferResponse = not nil
	//
	// Error = nil
	//
	// TransferResponse.IsFailure() = true
	//
	// Para qualquer outra resposta inesperada da API, possuímos o campo TransferResponse.Errors preenchido com as informações
	// de erro, o index 0 do slice com campo ErrorResponse.Code retornando a descrição
	// status http (Ex: "401 Unauthorized") e no campo ErrorResponse.Description retornará com o valor
	// "response status code not expected".
	//
	// # Error
	//
	// TransferResponse = nil
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
	// Recuperar uma única transferência: https://docs.asaas.com/reference/recuperar-uma-unica-transferencia
	GetById(ctx context.Context, transferId string) (*TransferResponse, error)
	// GetAll (Listar transferências)
	//
	// Este método retorna uma lista paginada com todas as transferências para o filtro informado.
	//
	// # Resposta: 200
	//
	// Pageable(TransferResponse) = not nil
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
	// Pageable(TransferResponse) = not nil
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
	// Pageable(TransferResponse) = nil
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
	// Listar transferências: https://docs.asaas.com/reference/listar-transferencias
	GetAll(ctx context.Context, filter GetAllTransfersRequest) (*Pageable[TransferResponse], error)
}

func NewTransfer(env Env, accessToken string) Transfer {
	logWarning("Transfer service running on", env.String())
	return transfer{
		env:         env,
		accessToken: accessToken,
	}
}

func (t transfer) TransferToBank(ctx context.Context, body TransferToBankRequest) (*TransferResponse, error) {
	req := NewRequest[TransferResponse](ctx, t.env, t.accessToken)
	return req.make(http.MethodPost, "/v3/transfers", body)
}

func (t transfer) TransferToAsaas(ctx context.Context, body TransferToAssasRequest) (*TransferResponse,
	error) {
	req := NewRequest[TransferResponse](ctx, t.env, t.accessToken)
	return req.make(http.MethodPost, "/v3/transfers", body)
}

func (t transfer) CancelById(ctx context.Context, transferId string) (*TransferResponse, error) {
	req := NewRequest[TransferResponse](ctx, t.env, t.accessToken)
	return req.make(http.MethodDelete, fmt.Sprintf("/v3/transfers/%s/cancel", transferId), nil)
}

func (t transfer) GetById(ctx context.Context, transferId string) (*TransferResponse, error) {
	req := NewRequest[TransferResponse](ctx, t.env, t.accessToken)
	return req.make(http.MethodGet, fmt.Sprintf("/v3/transfers/%s", transferId), nil)
}

func (t transfer) GetAll(ctx context.Context, filter GetAllTransfersRequest) (*Pageable[TransferResponse], error) {
	req := NewRequest[Pageable[TransferResponse]](ctx, t.env, t.accessToken)
	return req.make(http.MethodGet, "/v3/transfers", filter)
}
