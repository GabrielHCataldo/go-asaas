package asaas

import (
	"context"
	"fmt"
	"net/http"
	"os"
)

type CreateChargeRequest struct {
	// Identificador único do cliente no Asaas (REQUIRED)
	Customer string `json:"customer,omitempty" validate:"required"`
	// Forma de pagamento (Default: BillingTypeUndefined)
	BillingType BillingType `json:"billingType,omitempty" validate:"omitempty,enum"`
	// Valor da cobrança (REQUIRED)
	Value float64 `json:"value,omitempty" validate:"gte=0"`
	// Data de vencimento da cobrança (REQUIRED)
	DueDate *Date `json:"dueDate,omitempty" validate:"required,after_now"`
	// Descrição da cobrança (máx. 500 caracteres)
	Description string `json:"description,omitempty" validate:"omitempty,lte=500"`
	// Campo livre para busca
	ExternalReference string `json:"externalReference,omitempty"`
	// Informações de desconto
	Discount *DiscountRequest `json:"discount,omitempty"`
	// Informações de juros para pagamento após o vencimento
	Interest *InterestRequest `json:"interest,omitempty"`
	// Informações de multa para pagamento após o vencimento
	Fine *FineRequest `json:"fine,omitempty"`
	// Define se a cobrança será enviada via Correios
	PostalService bool `json:"postalService,omitempty"`
	// Configurações do split
	Split []SplitRequest `json:"split,omitempty"`
	// Informações de redirecionamento automático após pagamento na tela de fatura
	Callback *CallbackRequest `json:"callback,omitempty"`
	// Informações do cartão de crédito (REQUIRED se BillingType = BillingTypeCreditCard e se CreditCardToken não for informado)
	CreditCard *CreditCardRequest `json:"creditCard,omitempty"`
	// Informações do titular do cartão de crédito (REQUIRED se BillingType = BillingTypeCreditCard e se CreditCardToken não for informado)
	CreditCardHolderInfo *CreditCardHolderInfoRequest `json:"creditCardHolderInfo,omitempty"`
	// Token do cartão de crédito para uso da funcionalidade de tokenização de cartão de crédito
	CreditCardToken string `json:"creditCardToken,omitempty"`
	// Número de parcelas (somente no caso de cobrança parcelada)
	InstallmentCount int `json:"installmentCount,omitempty" validate:"omitempty,gte=2"`
	// Valor de cada parcela (somente no caso de cobrança parcelada)
	InstallmentValue float64 `json:"installmentValue,omitempty" validate:"omitempty,gt=0"`
	// Realizar apenas a Pré-Autorização da cobrança
	AuthorizeOnly bool `json:"authorizeOnly,omitempty"`
	// IP de onde o cliente está fazendo a compra. Não deve ser informado o IP do seu servidor
	RemoteIp string `json:"remoteIp,omitempty"`
}

type UpdateChargeRequest struct {
	Customer          string           `json:"customer,omitempty" validate:"required"`
	BillingType       BillingType      `json:"billingType,omitempty" validate:"required,enum"`
	Value             float64          `json:"value,omitempty" validate:"required,gt=0"`
	DueDate           Date             `json:"dueDate,omitempty" validate:"required"`
	Description       string           `json:"description,omitempty"`
	ExternalReference string           `json:"externalReference,omitempty"`
	Discount          *DiscountRequest `json:"discount,omitempty"`
	Interest          *InterestRequest `json:"interest,omitempty"`
	Fine              *FineRequest     `json:"fine,omitempty"`
	PostalService     bool             `json:"postalService,omitempty"`
	Split             []SplitRequest   `json:"split,omitempty"`
	Callback          *CallbackRequest `json:"callback,omitempty"`
	InstallmentCount  int              `json:"installmentCount,omitempty" validate:"omitempty,gte=2"`
	InstallmentValue  float64          `json:"installmentValue,omitempty" validate:"omitempty,gt=0"`
}

type GetAllChargesRequest struct {
	Customer              string        `json:"customer,omitempty"`
	Installment           string        `json:"installment,omitempty"`
	CustomerGroupName     string        `json:"customerGroupName,omitempty"`
	BillingType           BillingType   `json:"billingType,omitempty"`
	Status                ChargeStatus  `json:"status,omitempty"`
	Subscription          string        `json:"subscription,omitempty"`
	ExternalReference     string        `json:"externalReference,omitempty"`
	InvoiceStatus         InvoiceStatus `json:"invoiceStatus,omitempty"`
	EstimatedCreditDate   *Date         `json:"estimatedCreditDate,omitempty"`
	PixQrCodeId           string        `json:"pixQrCodeId,omitempty"`
	Anticipated           bool          `json:"anticipated,omitempty"`
	DateCreatedGe         *Date         `json:"dateCreated[ge],omitempty"`
	DateCreatedLe         *Date         `json:"dateCreated[le],omitempty"`
	EstimatedCreditDateGE *Date         `json:"estimatedCreditDate[ge],omitempty"`
	EstimatedCreditDateLE *Date         `json:"estimatedCreditDate[le],omitempty"`
	DueDateGE             *Date         `json:"dueDate[ge],omitempty"`
	DueDateLE             *Date         `json:"dueDate[le],omitempty"`
	User                  string        `json:"user,omitempty"`
	Offset                int           `json:"offset,omitempty"`
	Limit                 int           `json:"limit,omitempty"`
}

type ChargeReceiveInCashRequest struct {
	PaymentDate    *Date   `json:"paymentDate,omitempty" validate:"required"`
	Value          float64 `json:"value,omitempty" validate:"required,gt=0"`
	NotifyCustomer bool    `json:"notifyCustomer,omitempty"`
}

type UploadChargeDocumentRequest struct {
	AvailableAfterPayment bool         `json:"availableAfterPayment,omitempty"`
	Type                  DocumentType `json:"type,omitempty" validate:"required,enum"`
	File                  *os.File     `json:"file,omitempty" validate:"required"`
}

type UpdateChargeDocumentDefinitionsRequest struct {
	AvailableAfterPayment bool         `json:"availableAfterPayment,omitempty"`
	Type                  DocumentType `json:"type,omitempty" validate:"required,enum"`
}

type CallbackRequest struct {
	SuccessUrl   string `json:"successUrl,omitempty" validate:"required,url"`
	AutoRedirect bool   `json:"autoRedirect,omitempty"`
}

type ChargeResponse struct {
	Id                    string              `json:"id,omitempty"`
	Customer              string              `json:"customer,omitempty"`
	Status                ChargeStatus        `json:"status,omitempty"`
	PaymentLink           string              `json:"paymentLink,omitempty"`
	DueDate               *Date               `json:"dueDate,omitempty"`
	Value                 float64             `json:"value,omitempty"`
	NetValue              float64             `json:"netValue,omitempty"`
	BillingType           BillingType         `json:"billingType,omitempty"`
	CanBePaidAfterDueDate bool                `json:"canBePaidAfterDueDate,omitempty"`
	PixTransaction        string              `json:"pixTransaction,omitempty"`
	Description           string              `json:"description,omitempty"`
	ExternalReference     string              `json:"externalReference,omitempty"`
	OriginalValue         string              `json:"originalValue,omitempty"`
	InterestValue         string              `json:"interestValue,omitempty"`
	OriginalDueDate       *Date               `json:"originalDueDate,omitempty"`
	PaymentDate           *Date               `json:"paymentDate,omitempty"`
	ClientPaymentDate     *Date               `json:"clientPaymentDate,omitempty"`
	InstallmentNumber     int                 `json:"installmentCount,omitempty"`
	TransactionReceiptUrl string              `json:"transactionReceiptUrl,omitempty"`
	NossoNumero           string              `json:"nossoNumero,omitempty"`
	InvoiceUrl            string              `json:"invoiceUrl,omitempty"`
	BankSlipUrl           string              `json:"bankSlipUrl,omitempty"`
	InvoiceNumber         string              `json:"invoiceNumber,omitempty"`
	CreditCard            *CreditCardResponse `json:"creditCard,omitempty"`
	Discount              *DiscountResponse   `json:"discount,omitempty"`
	Fine                  *FineResponse       `json:"fine,omitempty"`
	Interest              *InterestResponse   `json:"interest,omitempty"`
	Deleted               bool                `json:"deleted,omitempty"`
	PostalService         bool                `json:"postalService,omitempty"`
	Anticipated           bool                `json:"anticipated,omitempty"`
	Anticipable           bool                `json:"anticipable,omitempty"`
	Refunds               []RefundResponse    `json:"refunds,omitempty"`
	DateCreated           *Date               `json:"dateCreated,omitempty"`
	Errors                []ErrorResponse     `json:"errors,omitempty"`
}

type ChargeStatusResponse struct {
	Status ChargeStatus    `json:"status,omitempty"`
	Errors []ErrorResponse `json:"errors,omitempty"`
}

type IdentificationFieldResponse struct {
	IdentificationField string          `json:"identificationField,omitempty"`
	NossoNumero         string          `json:"nossoNumero,omitempty"`
	BarCode             string          `json:"barCode,omitempty"`
	Errors              []ErrorResponse `json:"errors,omitempty"`
}

type ChargePixQrCodeResponse struct {
	EncodedImage   string          `json:"encodedImage,omitempty"`
	Payload        string          `json:"payload,omitempty"`
	ExpirationDate *Date           `json:"expirationDate,omitempty"`
	Errors         []ErrorResponse `json:"errors,omitempty"`
}

type ChargeDocumentResponse struct {
	Id                    string              `json:"id,omitempty"`
	Name                  string              `json:"name,omitempty"`
	AvailableAfterPayment bool                `json:"availableAfterPayment,omitempty"`
	Type                  DocumentType        `json:"type,omitempty"`
	File                  *ChargeFileResponse `json:"file,omitempty"`
	Deleted               bool                `json:"deleted,omitempty"`
	Errors                []ErrorResponse     `json:"errors,omitempty"`
}

type ChargeFileResponse struct {
	PublicId     string `json:"publicId,omitempty"`
	OriginalName string `json:"originalName,omitempty"`
	Size         int    `json:"size,omitempty"`
	Extension    string `json:"extension,omitempty"`
	PreviewUrl   string `json:"previewUrl,omitempty"`
	DownloadUrl  string `json:"downloadUrl,omitempty"`
}

type ChargeCreationLimitResponse struct {
	Creation ChargeCreationResponse `json:"creation,omitempty"`
	Errors   []ErrorResponse        `json:"errors,omitempty"`
}

type ChargeCreationResponse struct {
	Daily DailyCreationLimitResponse `json:"daily,omitempty"`
}

type DailyCreationLimitResponse struct {
	Limit      int  `json:"limit,omitempty"`
	Used       int  `json:"used,omitempty"`
	WasReached bool `json:"wasReached,omitempty"`
}

type charge struct {
	env         Env
	accessToken string
}

type Charge interface {
	// Create (Cria uma nova cobrança)
	//
	// É possível escolher entre as formas de pagamento com boleto, cartão de crédito,
	// pix ou permitir que o cliente escolha a forma que desejar.
	//
	// # Resposta: 200
	//
	// ChargeResponse = not nil
	//
	// Error = nil
	//
	// ChargeResponse.IsSuccess() = true
	//
	// Possui os valores de resposta de sucesso segunda a documentação.
	//
	// # Resposta: 400/401/500
	//
	// ChargeResponse = not nil
	//
	// Error = nil
	//
	// ChargeResponse.IsFailure() = true
	//
	// Para qualquer outra resposta inesperada da API, possuímos o campo ChargeResponse.Errors preenchido com as informações
	// de erro, sendo 400 retornado da API Asaas com as instruções de requisição conforme a documentação,
	// diferente disso retornará uma mensagem padrão no index 0 do slice com campo ErrorResponse.Code retornando a
	// descrição status http (Ex: "401 Unauthorized") e no campo ErrorResponse.Description retornará com o valor
	// "response status code not expected".
	//
	// # Error
	//
	// ChargeResponse = nil
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
	// Criar uma cobrança: https://docs.asaas.com/reference/criar-nova-cobranca
	//
	// Criar uma cobrança parcelada: https://docs.asaas.com/reference/criar-uma-cobranca-parcelada
	//
	// Criar cobrança com cartão de crédito: https://docs.asaas.com/reference/criar-nova-cobranca-com-cartao-de-credito
	Create(ctx context.Context, body CreateChargeRequest) (*ChargeResponse, Error)
	// PayWithCreditCard (Pagar uma cobrança com cartão de crédito)
	//
	// Este endpoint paga uma cobrança com o cartão de crédito informado na hora que você chamá-lo.
	// Não é possível agendar um pagamento
	//
	// # Resposta: 200
	//
	// ChargeResponse = not nil
	//
	// Error = nil
	//
	// ChargeResponse.IsSuccess() = true
	//
	// Possui os valores de resposta de sucesso segunda a documentação.
	//
	// # Resposta: 404
	//
	// ChargeResponse = not nil
	//
	// Error = nil
	//
	// ChargeResponse.IsNoContent() = true
	//
	// ID(s) informado no parâmetro não foi encontrado.
	//
	// # Resposta: 400/401/500
	//
	// ChargeResponse = not nil
	//
	// Error = nil
	//
	// ChargeResponse.IsFailure() = true
	//
	// Para qualquer outra resposta inesperada da API, possuímos o campo ChargeResponse.Errors preenchido com as informações
	// de erro, sendo 400 retornado da API Asaas com as instruções de requisição conforme a documentação,
	// diferente disso retornará uma mensagem padrão no index 0 do slice com campo ErrorResponse.Code retornando a
	// descrição status http (Ex: "401 Unauthorized") e no campo ErrorResponse.Description retornará com o valor
	// "response status code not expected".
	//
	// # Error
	//
	// ChargeResponse = nil
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
	// Pagar uma cobrança com cartão de crédito: https://docs.asaas.com/reference/pagar-uma-cobranca-com-cartao-de-credito
	PayWithCreditCard(ctx context.Context, chargeId string, body CreditCardRequest) (*ChargeResponse, Error)
	// UpdateById (Atualizar cobrança existente)
	//
	// Somente é possível atualizar cobranças aguardando pagamento ou vencidas. Uma vez criada, não é possível alterar
	// o cliente ao qual a cobrança pertence.
	//
	// # Resposta: 200
	//
	// ChargeResponse = not nil
	//
	// Error = nil
	//
	// ChargeResponse.IsSuccess() = true
	//
	// Possui os valores de resposta de sucesso segunda a documentação.
	//
	// # Resposta: 404
	//
	// ChargeResponse = not nil
	//
	// Error = nil
	//
	// ChargeResponse.IsNoContent() = true
	//
	// ID(s) informado no parâmetro não foi encontrado.
	//
	// # Resposta: 400/401/500
	//
	// ChargeResponse = not nil
	//
	// Error = nil
	//
	// ChargeResponse.IsFailure() = true
	//
	// Para qualquer outra resposta inesperada da API, possuímos o campo ChargeResponse.Errors preenchido com as informações
	// de erro, sendo 400 retornado da API Asaas com as instruções de requisição conforme a documentação,
	// diferente disso retornará uma mensagem padrão no index 0 do slice com campo ErrorResponse.Code retornando a
	// descrição status http (Ex: "401 Unauthorized") e no campo ErrorResponse.Description retornará com o valor
	// "response status code not expected".
	//
	// # Error
	//
	// ChargeResponse = nil
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
	// Atualizar cobrança existente: https://docs.asaas.com/reference/atualizar-cobranca-existente
	UpdateById(ctx context.Context, chargeId string, body UpdateChargeRequest) (*ChargeResponse, Error)
	// DeleteById (Excluir cobrança)
	//
	// Somente cobranças aguardando pagamento ou vencidas podem ser removidas. Ao excluir uma cobrança,
	// nenhuma nova notificação será enviada e seu cliente não poderá mais pagá-la.
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
	// Excluir cobrança: https://docs.asaas.com/reference/remover-cobranca
	DeleteById(ctx context.Context, chargeId string) (*DeleteResponse, Error)
	// RestoreById (Restaurar cobrança removida)
	//
	// # Resposta: 200
	//
	// ChargeResponse = not nil
	//
	// Error = nil
	//
	// ChargeResponse.IsSuccess() = true
	//
	// Possui os valores de resposta de sucesso segunda a documentação.
	//
	// # Resposta: 404
	//
	// ChargeResponse = not nil
	//
	// Error = nil
	//
	// ChargeResponse.IsNoContent() = true
	//
	// ID(s) informado no parâmetro não foi encontrado.
	//
	// # Resposta: 400/401/500
	//
	// ChargeResponse = not nil
	//
	// Error = nil
	//
	// ChargeResponse.IsFailure() = true
	//
	// Para qualquer outra resposta inesperada da API, possuímos o campo ChargeResponse.Errors preenchido com as informações
	// de erro, sendo 400 retornado da API Asaas com as instruções de requisição conforme a documentação,
	// diferente disso retornará uma mensagem padrão no index 0 do slice com campo ErrorResponse.Code retornando a
	// descrição status http (Ex: "401 Unauthorized") e no campo ErrorResponse.Description retornará com o valor
	// "response status code not expected".
	//
	// # Error
	//
	// ChargeResponse = nil
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
	// Restaurar cobrança removida: https://docs.asaas.com/reference/restaurar-cobranca-removida
	RestoreById(ctx context.Context, chargeId string) (*ChargeResponse, Error)
	// RefundById (Estornar cobrança)
	//
	// É possível estornar cobranças via cartão de crédito recebidas ou confirmadas. Ao fazer isto o saldo correspondente
	// é debitado de sua conta no Asaas e a cobrança cancelada no cartão do seu cliente. O cancelamento pode levar até 10
	// dias úteis para aparecer na fatura de seu cliente. Cobranças recebidas via Pix, permitem o estorno integral ou
	// vários estornos parciais. A soma desses estornos não poderão ultrapassar o valor total da cobrança recebida.
	//
	// # Resposta: 200
	//
	// ChargeResponse = not nil
	//
	// Error = nil
	//
	// ChargeResponse.IsSuccess() = true
	//
	// Possui os valores de resposta de sucesso segunda a documentação.
	//
	// # Resposta: 404
	//
	// ChargeResponse = not nil
	//
	// Error = nil
	//
	// ChargeResponse.IsNoContent() = true
	//
	// ID(s) informado no parâmetro não foi encontrado.
	//
	// # Resposta: 400/401/500
	//
	// ChargeResponse = not nil
	//
	// Error = nil
	//
	// ChargeResponse.IsFailure() = true
	//
	// Para qualquer outra resposta inesperada da API, possuímos o campo ChargeResponse.Errors preenchido com as informações
	// de erro, sendo 400 retornado da API Asaas com as instruções de requisição conforme a documentação,
	// diferente disso retornará uma mensagem padrão no index 0 do slice com campo ErrorResponse.Code retornando a
	// descrição status http (Ex: "401 Unauthorized") e no campo ErrorResponse.Description retornará com o valor
	// "response status code not expected".
	//
	// # Error
	//
	// ChargeResponse = nil
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
	// Estornar cobrança: https://docs.asaas.com/reference/estornar-cobranca
	RefundById(ctx context.Context, chargeId string, body RefundRequest) (*ChargeResponse, Error)
	// ReceiveInCashById (Confirmar recebimento em dinheiro)
	//
	// Quando um cliente fizer o pagamento de uma cobrança diretamente para você, sem que esse pagamento seja processado
	// pelo Asaas, utilize este método para definir a cobrança como recebida em dinheiro. Esta opção permite manter
	// seu histórico consistente no sistema, mas não gera saldo ou faz qualquer alteração financeira em sua conta.
	// Ao confirmar um recebimento em dinheiro de uma cobrança que possua uma negativação em andamento uma taxa de
	// ativação de serviço de negativação poderá ser cobrada. Verifique essa taxa no campo receivedInCashFeeValue
	// localizada no retorno do objeto de negativação.
	//
	// # Resposta: 200
	//
	// ChargeResponse = not nil
	//
	// Error = nil
	//
	// ChargeResponse.IsSuccess() = true
	//
	// Possui os valores de resposta de sucesso segunda a documentação.
	//
	// # Resposta: 404
	//
	// ChargeResponse = not nil
	//
	// Error = nil
	//
	// ChargeResponse.IsNoContent() = true
	//
	// ID(s) informado no parâmetro não foi encontrado.
	//
	// # Resposta: 400/401/500
	//
	// ChargeResponse = not nil
	//
	// Error = nil
	//
	// ChargeResponse.IsFailure() = true
	//
	// Para qualquer outra resposta inesperada da API, possuímos o campo ChargeResponse.Errors preenchido com as informações
	// de erro, sendo 400 retornado da API Asaas com as instruções de requisição conforme a documentação,
	// diferente disso retornará uma mensagem padrão no index 0 do slice com campo ErrorResponse.Code retornando a
	// descrição status http (Ex: "401 Unauthorized") e no campo ErrorResponse.Description retornará com o valor
	// "response status code not expected".
	//
	// # Error
	//
	// ChargeResponse = nil
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
	// Confirmar recebimento em dinheiro: https://docs.asaas.com/reference/confirmar-recebimento-em-dinheiro
	ReceiveInCashById(ctx context.Context, chargeId string, body ChargeReceiveInCashRequest) (*ChargeResponse, Error)
	// UndoReceivedInCashById (Desfazer confirmação de recebimento em dinheiro)
	//
	// Permite desfazer uma cobrança marcada como recebida em dinheiro.
	//
	// # Resposta: 200
	//
	// ChargeResponse = not nil
	//
	// Error = nil
	//
	// ChargeResponse.IsSuccess() = true
	//
	// Possui os valores de resposta de sucesso segunda a documentação.
	//
	// # Resposta: 404
	//
	// ChargeResponse = not nil
	//
	// Error = nil
	//
	// ChargeResponse.IsNoContent() = true
	//
	// ID(s) informado no parâmetro não foi encontrado.
	//
	// # Resposta: 400/401/500
	//
	// ChargeResponse = not nil
	//
	// Error = nil
	//
	// ChargeResponse.IsFailure() = true
	//
	// Para qualquer outra resposta inesperada da API, possuímos o campo ChargeResponse.Errors preenchido com as informações
	// de erro, sendo 400 retornado da API Asaas com as instruções de requisição conforme a documentação,
	// diferente disso retornará uma mensagem padrão no index 0 do slice com campo ErrorResponse.Code retornando a
	// descrição status http (Ex: "401 Unauthorized") e no campo ErrorResponse.Description retornará com o valor
	// "response status code not expected".
	//
	// # Error
	//
	// ChargeResponse = nil
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
	// Desfazer confirmação de recebimento em dinheiro: https://docs.asaas.com/reference/desfazer-confirmacao-de-recebimento-em-dinheiro
	UndoReceivedInCashById(ctx context.Context, chargeId string) (*ChargeResponse, Error)
	// UploadDocumentById (Fazer upload de documentos da cobrança)
	//
	// Permite anexar um documento dentro da cobrança, que será disponibilizado ao pagador diretamente na fatura Asaas
	// para download.
	//
	// # Resposta: 200
	//
	// ChargeDocumentResponse = not nil
	//
	// Error = nil
	//
	// ChargeDocumentResponse.IsSuccess() = true
	//
	// Possui os valores de resposta de sucesso segunda a documentação.
	//
	// # Resposta: 404
	//
	// ChargeDocumentResponse = not nil
	//
	// Error = nil
	//
	// ChargeDocumentResponse.IsNoContent() = true
	//
	// ID(s) informado no parâmetro não foi encontrado.
	//
	// # Resposta: 400/401/500
	//
	// ChargeDocumentResponse = not nil
	//
	// Error = nil
	//
	// ChargeDocumentResponse.IsFailure() = true
	//
	// Para qualquer outra resposta inesperada da API, possuímos o campo ChargeDocumentResponse.Errors preenchido
	// com as informações de erro, sendo 400 retornado da API Asaas com as instruções de requisição conforme a documentação,
	// diferente disso retornará uma mensagem padrão no index 0 do slice com campo ErrorResponse.Code retornando a
	// descrição status http (Ex: "401 Unauthorized") e no campo ErrorResponse.Description retornará com o valor
	// "response status code not expected".
	//
	// # Error
	//
	// ChargeDocumentResponse = nil
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
	// Fazer upload de documentos da cobrança: https://docs.asaas.com/reference/fazer-upload-de-documentos-da-cobranca
	UploadDocumentById(ctx context.Context, chargeId string, body UploadChargeDocumentRequest) (*ChargeDocumentResponse, Error)
	// UpdateDocumentDefinitionsById (Atualizar definições de um documento da cobrança)
	//
	// Permite atualizar tipo de arquivo e definição de disponibilização do arquivo após o pagamento de um documento
	// anexado em uma cobrança.
	//
	// # Resposta: 200
	//
	// ChargeDocumentResponse = not nil
	//
	// Error = nil
	//
	// ChargeDocumentResponse.IsSuccess() = true
	//
	// Possui os valores de resposta de sucesso segunda a documentação.
	//
	// # Resposta: 404
	//
	// ChargeDocumentResponse = not nil
	//
	// Error = nil
	//
	// ChargeDocumentResponse.IsNoContent() = true
	//
	// ID(s) informado no parâmetro não foi encontrado.
	//
	// # Resposta: 401/500
	//
	// ChargeDocumentResponse = not nil
	//
	// Error = nil
	//
	// ChargeDocumentResponse.IsFailure() = true
	//
	// Para qualquer outra resposta inesperada da API, possuímos o campo ChargeDocumentResponse.Errors preenchido com
	// as informações de erro, o index 0 do slice com campo ErrorResponse.Code retornando a descrição
	// status http (Ex: "401 Unauthorized") e no campo ErrorResponse.Description retornará com o valor
	// "response status code not expected".
	//
	// # Error
	//
	// ChargeDocumentResponse = nil
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
	// Atualizar definições de um documento da cobrança: https://docs.asaas.com/reference/atualizar-definicoes-de-um-documento-da-cobranca
	UpdateDocumentDefinitionsById(ctx context.Context, chargeId, docId string, body UpdateChargeDocumentDefinitionsRequest) (
		*ChargeDocumentResponse, Error)
	// DeleteDocumentById (Excluir documento de uma cobrança)
	//
	// Para excluir o documento de uma cobrança, é necessário que você tenha o ID que o Asaas retornou no momento do
	// upload do documento.
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
	// Excluir documento de uma cobrança: https://docs.asaas.com/reference/excluir-documento-de-uma-cobranca
	DeleteDocumentById(ctx context.Context, chargeId, docId string) (*DeleteResponse, Error)
	// GetById (Recuperar uma única cobrança)
	//
	// Para recuperar uma cobrança específica é necessário que você tenha o ID que o Asaas retornou no momento da
	// criação dela.
	//
	// # Resposta: 200
	//
	// ChargeResponse = not nil
	//
	// Error = nil
	//
	// ChargeResponse.IsSuccess() = true
	//
	// Possui os valores de resposta de sucesso segunda a documentação.
	//
	// # Resposta: 404
	//
	// ChargeDocumentResponse = not nil
	//
	// Error = nil
	//
	// ChargeDocumentResponse.IsNoContent() = true
	//
	// ID(s) informado no parâmetro não foi encontrado.
	//
	// # Resposta: 401/500
	//
	// ChargeResponse = not nil
	//
	// Error = nil
	//
	// ChargeResponse.IsFailure() = true
	//
	// Para qualquer outra resposta inesperada da API, possuímos o campo ChargeResponse.Errors preenchido com as informações
	// de erro, o index 0 do slice com campo ErrorResponse.Code retornando a descrição
	// status http (Ex: "401 Unauthorized") e no campo ErrorResponse.Description retornará com o valor
	// "response status code not expected".
	//
	// # Error
	//
	// ChargeResponse = nil
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
	// Recuperar uma única cobrança: https://docs.asaas.com/reference/recuperar-uma-unica-cobranca
	GetById(ctx context.Context, chargeId string) (*ChargeResponse, Error)
	// GetCreationLimit (Recuperando limites de cobranças)
	//
	// Permite você recuperar o limite de criação de cobranças configurado na conta, quantidade de cobranças criadas no
	// dia e se é possível criar mais cobranças.
	//
	// # Resposta: 200
	//
	// ChargeCreationLimitResponse = not nil
	//
	// Error = nil
	//
	// ChargeCreationLimitResponse.IsSuccess() = true
	//
	// Possui os valores de resposta de sucesso segunda a documentação.
	//
	// # Resposta: 401/500
	//
	// ChargeCreationLimitResponse = not nil
	//
	// Error = nil
	//
	// ChargeCreationLimitResponse.IsFailure() = true
	//
	// Para qualquer outra resposta inesperada da API, possuímos o campo ChargeCreationLimitResponse.Errors preenchido
	// com as informações de erro, o index 0 do slice com campo ErrorResponse.Code retornando a descrição
	// status http (Ex: "401 Unauthorized") e no campo ErrorResponse.Description retornará com o valor
	// "response status code not expected".
	//
	// # Error
	//
	// ChargeCreationLimitResponse = nil
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
	// Recuperando limites de cobranças: https://docs.asaas.com/reference/recuperando-limites-de-cobrancças
	GetCreationLimit(ctx context.Context) (*ChargeCreationLimitResponse, Error)
	// GetStatusById (Recuperar status de uma cobrança)
	//
	// # Resposta: 200
	//
	// ChargeStatusResponse = not nil
	//
	// Error = nil
	//
	// ChargeStatusResponse.IsSuccess() = true
	//
	// Possui os valores de resposta de sucesso segunda a documentação.
	//
	// # Resposta: 404
	//
	// ChargeStatusResponse = not nil
	//
	// Error = nil
	//
	// ChargeStatusResponse.IsNoContent() = true
	//
	// ID(s) informado no parâmetro não foi encontrado.
	//
	// # Resposta: 401/500
	//
	// ChargeStatusResponse = not nil
	//
	// Error = nil
	//
	// ChargeStatusResponse.IsFailure() = true
	//
	// Para qualquer outra resposta inesperada da API, possuímos o campo ChargeStatusResponse.Errors preenchido com as
	// informações de erro, o index 0 do slice com campo ErrorResponse.Code retornando a descrição
	// status http (Ex: "401 Unauthorized") e no campo ErrorResponse.Description retornará com o valor
	// "response status code not expected".
	//
	// # Error
	//
	// ChargeStatusResponse = nil
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
	// Recuperar uma única cobrança: https://docs.asaas.com/reference/recuperar-uma-unica-cobranca
	GetStatusById(ctx context.Context, chargeId string) (*ChargeStatusResponse, Error)
	GetIdentificationFieldById(ctx context.Context, chargeId string) (*IdentificationFieldResponse, Error)
	GetPixQrCodeById(ctx context.Context, chargeId string) (*ChargePixQrCodeResponse, Error)
	GetDocumentById(ctx context.Context, chargeId, docId string) (*ChargeDocumentResponse, Error)
	GetAllDocumentsById(ctx context.Context, chargeId string, filter PageableDefaultRequest) (
		*Pageable[ChargeDocumentResponse], Error)
	GetAll(ctx context.Context, filter GetAllChargesRequest) (*Pageable[ChargeResponse], Error)
}

func NewCharge(env Env, accessCode string) Charge {
	logWarning("Charge service running on", env.String())
	return charge{
		env:         env,
		accessToken: accessCode,
	}
}

func (c charge) Create(ctx context.Context, body CreateChargeRequest) (*ChargeResponse, Error) {
	if err := c.validateCreateBodyRequest(body); err != nil {
		return nil, NewError(ErrorTypeValidation, err)
	}
	c.prepareCreateBodyRequest(&body)
	req := NewRequest[ChargeResponse](ctx, c.env, c.accessToken)
	return req.make(http.MethodPost, "/v3/payments", body)
}

func (c charge) PayWithCreditCard(ctx context.Context, chargeId string, body CreditCardRequest) (*ChargeResponse,
	Error) {
	if err := Validate().Struct(body); err != nil {
		return nil, NewError(ErrorTypeValidation, err)
	}
	req := NewRequest[ChargeResponse](ctx, c.env, c.accessToken)
	return req.make(http.MethodPost, fmt.Sprintf(`/v3/payments/%s/payWithCreditCard`, chargeId), body)
}

func (c charge) UpdateById(ctx context.Context, chargeId string, body UpdateChargeRequest) (*ChargeResponse,
	Error) {
	if err := Validate().Struct(body); err != nil {
		return nil, NewError(ErrorTypeValidation, err)
	}
	req := NewRequest[ChargeResponse](ctx, c.env, c.accessToken)
	return req.make(http.MethodPut, fmt.Sprintf(`/v3/payments/%s`, chargeId), body)
}

func (c charge) DeleteById(ctx context.Context, chargeId string) (*DeleteResponse, Error) {
	req := NewRequest[DeleteResponse](ctx, c.env, c.accessToken)
	return req.make(http.MethodDelete, fmt.Sprintf(`/v3/payments/%s`, chargeId), nil)
}

func (c charge) RestoreById(ctx context.Context, chargeId string) (*ChargeResponse, Error) {
	req := NewRequest[ChargeResponse](ctx, c.env, c.accessToken)
	return req.make(http.MethodPost, fmt.Sprintf(`/v3/payments/%s/restore`, chargeId), nil)
}

func (c charge) RefundById(ctx context.Context, chargeId string, body RefundRequest) (
	*ChargeResponse, Error) {
	req := NewRequest[ChargeResponse](ctx, c.env, c.accessToken)
	return req.make(http.MethodPost, fmt.Sprintf(`/v3/payments/%s/refund`, chargeId), body)
}

func (c charge) ReceiveInCashById(ctx context.Context, chargeId string, body ChargeReceiveInCashRequest) (
	*ChargeResponse, Error) {
	if err := Validate().Struct(body); err != nil {
		return nil, NewError(ErrorTypeValidation, err)
	}
	req := NewRequest[ChargeResponse](ctx, c.env, c.accessToken)
	return req.make(http.MethodPost, fmt.Sprintf(`/v3/payments/%s/receiveInCash`, chargeId), body)
}

func (c charge) UndoReceivedInCashById(ctx context.Context, chargeId string) (*ChargeResponse, Error) {
	req := NewRequest[ChargeResponse](ctx, c.env, c.accessToken)
	return req.make(http.MethodPost, fmt.Sprintf(`/v3/payments/%s/undoReceivedInCash`, chargeId), nil)
}

func (c charge) UploadDocumentById(ctx context.Context, chargeId string, body UploadChargeDocumentRequest) (
	*ChargeDocumentResponse, Error) {
	if err := Validate().Struct(body); err != nil {
		return nil, NewError(ErrorTypeValidation, err)
	}
	req := NewRequest[ChargeDocumentResponse](ctx, c.env, c.accessToken)
	return req.makeMultipartForm(http.MethodPost, fmt.Sprintf(`/v3/payments/%s/documents`, chargeId), body)
}

func (c charge) UpdateDocumentDefinitionsById(
	ctx context.Context,
	chargeId,
	docId string,
	body UpdateChargeDocumentDefinitionsRequest,
) (*ChargeDocumentResponse, Error) {
	if err := Validate().Struct(body); err != nil {
		return nil, NewError(ErrorTypeValidation, err)
	}
	req := NewRequest[ChargeDocumentResponse](ctx, c.env, c.accessToken)
	return req.make(http.MethodPut, fmt.Sprintf(`/v3/payments/%s/documents/%v`, chargeId, docId), body)
}

func (c charge) DeleteDocumentById(ctx context.Context, chargeId, docId string) (*DeleteResponse, Error) {
	req := NewRequest[DeleteResponse](ctx, c.env, c.accessToken)
	return req.make(http.MethodDelete, fmt.Sprintf(`/v3/payments/%s/documents/%v`, chargeId, docId), nil)
}

func (c charge) GetCreationLimit(ctx context.Context) (*ChargeCreationLimitResponse, Error) {
	req := NewRequest[ChargeCreationLimitResponse](ctx, c.env, c.accessToken)
	return req.make(http.MethodGet, "/v3/payments/limits", nil)
}

func (c charge) GetById(ctx context.Context, chargeId string) (*ChargeResponse, Error) {
	req := NewRequest[ChargeResponse](ctx, c.env, c.accessToken)
	return req.make(http.MethodGet, fmt.Sprintf(`/v3/payments/%s`, chargeId), nil)
}

func (c charge) GetStatusById(ctx context.Context, chargeId string) (*ChargeStatusResponse, Error) {
	req := NewRequest[ChargeStatusResponse](ctx, c.env, c.accessToken)
	return req.make(http.MethodGet, fmt.Sprintf(`/v3/payments/%s/status`, chargeId), nil)
}

func (c charge) GetIdentificationFieldById(ctx context.Context, chargeId string) (*IdentificationFieldResponse,
	Error) {
	req := NewRequest[IdentificationFieldResponse](ctx, c.env, c.accessToken)
	return req.make(http.MethodGet, fmt.Sprintf(`/v3/payments/%s/identificationField`, chargeId), nil)
}

func (c charge) GetPixQrCodeById(ctx context.Context, chargeId string) (*ChargePixQrCodeResponse, Error) {
	req := NewRequest[ChargePixQrCodeResponse](ctx, c.env, c.accessToken)
	return req.make(http.MethodGet, fmt.Sprintf(`/v3/payments/%s/pixQrCode`, chargeId), nil)
}

func (c charge) GetDocumentById(ctx context.Context, chargeId, docId string) (*ChargeDocumentResponse, Error) {
	req := NewRequest[ChargeDocumentResponse](ctx, c.env, c.accessToken)
	return req.make(http.MethodGet, fmt.Sprintf(`/v3/payments/%s/documents/%v`, chargeId, docId), nil)
}

func (c charge) GetAllDocumentsById(ctx context.Context, chargeId string, filter PageableDefaultRequest) (
	*Pageable[ChargeDocumentResponse], Error) {
	req := NewRequest[Pageable[ChargeDocumentResponse]](ctx, c.env, c.accessToken)
	return req.make(http.MethodGet, fmt.Sprintf(`/v3/payments/%s/documents`, chargeId), filter)
}

func (c charge) GetAll(ctx context.Context, filter GetAllChargesRequest) (
	*Pageable[ChargeResponse], Error) {
	req := NewRequest[Pageable[ChargeResponse]](ctx, c.env, c.accessToken)
	return req.make(http.MethodGet, "/v3/payments", filter)
}

func (c charge) validateCreateBodyRequest(body CreateChargeRequest) error {
	if err := Validate().Struct(body); err != nil {
		return err
	}
	return validateBillingBody(body.BillingType, body.CreditCard, body.CreditCardHolderInfo, body.CreditCardToken,
		body.RemoteIp)
}

func (c charge) prepareCreateBodyRequest(body *CreateChargeRequest) {
	if !body.BillingType.IsEnumValid() {
		body.BillingType = BillingTypeUndefined
	}
	switch body.BillingType {
	case BillingTypeCreditCard:
		if body.Fine != nil {
			body.Fine.DueDateLimitDays = 0
		}
		break
	default:
		body.CreditCard = nil
		body.CreditCardHolderInfo = nil
		body.CreditCardToken = ""
		body.RemoteIp = ""
	}
}
