package asaas

import (
	"context"
	"fmt"
	"net/http"
	"os"
)

type CreateChargeRequest struct {
	// Identificador único do cliente no Asaas (REQUIRED)
	Customer string `json:"customer,omitempty"`
	// Forma de pagamento (Default: BillingTypeUndefined)
	BillingType BillingType `json:"billingType,omitempty"`
	// Valor da cobrança (REQUIRED)
	Value float64 `json:"value,omitempty"`
	// Data de vencimento da cobrança (REQUIRED)
	DueDate Date `json:"dueDate,omitempty"`
	// Descrição da cobrança (máx. 500 caracteres)
	Description string `json:"description,omitempty"`
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
	InstallmentCount int `json:"installmentCount,omitempty"`
	// Valor de cada parcela (somente no caso de cobrança parcelada)
	InstallmentValue float64 `json:"installmentValue,omitempty"`
	// Realizar apenas a Pré-Autorização da cobrança
	AuthorizeOnly bool `json:"authorizeOnly,omitempty"`
	// IP de onde o cliente está fazendo a compra. Não deve ser informado o IP do seu servidor (REQUIRED se BillingType = BillingTypeCreditCard)
	RemoteIp string `json:"remoteIp,omitempty"`
}

type PayWithCreditCardRequest struct {
	// Informações do cartão de crédito (REQUIRED se CreditCardToken não for informado)
	CreditCard *CreditCardRequest `json:"creditCard,omitempty"`
	// Informações do titular do cartão de crédito (REQUIRED se CreditCardToken não for informado)
	CreditCardHolderInfo *CreditCardHolderInfoRequest `json:"creditCardHolderInfo,omitempty"`
	// Token do cartão de crédito para uso da funcionalidade de tokenização de cartão de crédito. Caso informado, os campos acima não são obrigatórios.
	CreditCardToken string `json:"creditCardToken,omitempty"`
}

type UpdateChargeRequest struct {
	// Identificador único do cliente no Asaas
	Customer string `json:"customer,omitempty"`
	// Forma de pagamento
	BillingType BillingType `json:"billingType,omitempty"`
	// Valor da cobrança
	Value float64 `json:"value,omitempty"`
	// Data de vencimento da cobrança
	DueDate Date `json:"dueDate,omitempty"`
	// Descrição da cobrança (máx. 500 caracteres)
	Description *string `json:"description,omitempty"`
	// Campo livre para busca
	ExternalReference *string `json:"externalReference,omitempty"`
	// Informações de desconto
	Discount *DiscountRequest `json:"discount,omitempty"`
	// Informações de juros para pagamento após o vencimento
	Interest *InterestRequest `json:"interest,omitempty"`
	// Informações de multa para pagamento após o vencimento
	Fine *FineRequest `json:"fine,omitempty"`
	// Define se a cobrança será enviada via Correios
	PostalService *bool `json:"postalService,omitempty"`
	// Configurações do split
	Split []SplitRequest `json:"split,omitempty"`
	// Informações de redirecionamento automático após pagamento na tela de fatura
	Callback *CallbackRequest `json:"callback,omitempty"`
	// Número de parcelas (somente no caso de cobrança parcelada)
	InstallmentCount int `json:"installmentCount,omitempty"`
	// Valor de cada parcela (somente no caso de cobrança parcelada)
	InstallmentValue float64 `json:"installmentValue,omitempty"`
}

type GetAllChargesRequest struct {
	// Filtrar pelo Identificador único do cliente
	Customer string `json:"customer,omitempty"`
	// Filtrar pelo Identificador único da assinatura
	Subscription string `json:"subscription,omitempty"`
	// Filtrar pelo Identificador único do parcelamento
	Installment string `json:"installment,omitempty"`
	// Filtrar pelo nome do grupo de cliente
	CustomerGroupName string `json:"customerGroupName,omitempty"`
	// Filtrar por forma de pagamento
	BillingType BillingType `json:"billingType,omitempty"`
	// Filtrar por status
	Status ChargeStatus `json:"status,omitempty"`
	// Filtrar pelo Identificador do seu sistema
	ExternalReference string `json:"externalReference,omitempty"`
	// Filtro para retornar cobranças que possuem ou não nota fiscal.
	InvoiceStatus InvoiceStatus `json:"invoiceStatus,omitempty"`
	// Filtrar pela data estimada de crédito.
	EstimatedCreditDate Date `json:"estimatedCreditDate,omitempty"`
	// Filtrar recebimentos originados de um QrCode estático utilizando o id gerado na hora da criação do QrCode.
	PixQrCodeId string `json:"pixQrCodeId,omitempty"`
	// Filtrar registros antecipados ou não
	Anticipated *bool `json:"anticipated,omitempty"`
	// Filtrar pela data de pagamento
	PaymentDate Date `json:"paymentDate,omitempty"`
	// Filtrar a partir da data de criação inicial
	DateCreatedGe Date `json:"dateCreated[ge],omitempty"`
	// Filtrar a partir da data de criação final
	DateCreatedLe Date `json:"dateCreated[le],omitempty"`
	// Filtrar a partir da data de recebimento inicial
	PaymentDateGe Date `json:"paymentDate[ge],omitempty"`
	// Filtrar a partir da data de recebimento final
	PaymentDateLe Date `json:"paymentDate[le],omitempty"`
	// Filtrar a partir da data estimada de crédito inicial
	EstimatedCreditDateGE Date `json:"estimatedCreditDate[ge],omitempty"`
	// Filtrar a partir da data estimada de crédito final
	EstimatedCreditDateLE Date `json:"estimatedCreditDate[le],omitempty"`
	// Filtrar a partir da data de vencimento inicial
	DueDateGe Date `json:"dueDate[ge],omitempty"`
	// Filtrar a partir da data de vencimento final
	DueDateLe Date `json:"dueDate[le],omitempty"`
	// Filtrar pelo endereço de e-mail do usuário que criou a cobrança.
	User string `json:"user,omitempty"`
	// Elemento inicial da lista
	Offset int `json:"offset,omitempty"`
	// Número de elementos da lista (max: 100)
	Limit int `json:"limit,omitempty"`
}

type ChargeReceiveInCashRequest struct {
	// Data em que o cliente efetuou o pagamento (REQUIRED)
	PaymentDate Date `json:"paymentDate,omitempty"`
	// Valor pago pelo cliente (REQUIRED)
	Value float64 `json:"value,omitempty"`
	// Enviar ou não notificação de pagamento confirmado para o cliente
	NotifyCustomer bool `json:"notifyCustomer,omitempty"`
}

type UploadChargeDocumentRequest struct {
	// true para disponibilizar o arquivo somente após o recebimento da cobrança
	AvailableAfterPayment bool `json:"availableAfterPayment"`
	// Tipo de documento (REQUIRED)
	Type DocumentType `json:"type,omitempty"`
	// Arquivo a ser anexado (REQUIRED)
	File *os.File `json:"file,omitempty"`
}

type UpdateChargeDocumentDefinitionsRequest struct {
	// Define se o arquivo será disponibilizado somente após o pagamento
	AvailableAfterPayment bool `json:"availableAfterPayment"`
	// Novo tipo do documento (REQUIRED)
	Type DocumentType `json:"type,omitempty"`
}

type CallbackRequest struct {
	// URL que o cliente será redirecionado após o pagamento com sucesso da fatura ou link de pagamento (REQUIRED)
	SuccessUrl string `json:"successUrl,omitempty"`
	// Definir se o cliente será redirecionado automaticamente ou será apenas informado com um botão para retornar ao site. O padrão é true, caso queira desativar informar false
	AutoRedirect *bool `json:"autoRedirect,omitempty"`
}

type ChargeResponse struct {
	Id                    string              `json:"id,omitempty"`
	Customer              string              `json:"customer,omitempty"`
	Installment           string              `json:"installment,omitempty"`
	Status                ChargeStatus        `json:"status,omitempty"`
	PaymentLink           string              `json:"paymentLink,omitempty"`
	DueDate               Date                `json:"dueDate,omitempty"`
	Value                 float64             `json:"value,omitempty"`
	NetValue              float64             `json:"netValue,omitempty"`
	BillingType           BillingType         `json:"billingType,omitempty"`
	CanBePaidAfterDueDate bool                `json:"canBePaidAfterDueDate,omitempty"`
	PixTransaction        string              `json:"pixTransaction,omitempty"`
	Description           string              `json:"description,omitempty"`
	ExternalReference     string              `json:"externalReference,omitempty"`
	OriginalValue         string              `json:"originalValue,omitempty"`
	InterestValue         string              `json:"interestValue,omitempty"`
	OriginalDueDate       Date                `json:"originalDueDate,omitempty"`
	PaymentDate           Date                `json:"paymentDate,omitempty"`
	ClientPaymentDate     Date                `json:"clientPaymentDate,omitempty"`
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
	DateCreated           Date                `json:"dateCreated,omitempty"`
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
	ExpirationDate Date            `json:"expirationDate,omitempty"`
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
	// O CreateChargeRequest.BillingType BillingTypeBankSlip habilita o pagamento em PIX e Boleto. Em BillingTypePix,
	// apenas o pagamento em PIX, e em BillingTypeCreditCard, em cartão de crédito e débito (na fatura).
	//
	// Não é possível gerar uma cobrança com dois BillingType diferentes (BillingTypePix e BillingTypeCreditCard,
	// por exemplo, para não gerar o boleto, apenas esses dois tipos de cobrança).
	//
	// Caso não queira receber pagamento em PIX ou em Cartão de débito, é possível desabilitar dentro de sua interface
	// em Minha Conta > Configuração > Configurações do Sistema.
	//
	// Caso queira desabilitar em subcontas white label, entre em contato com o nosso time de
	// integração em suporte-tecnico@asaas.com.br
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
	// error = not nil
	//
	// Se o parâmetro de retorno error não estiver nil quer dizer que ocorreu um erro inesperado
	// na lib go-asaas.
	//
	// Se isso acontecer por favor report o erro no repositório: https://github.com/GabrielHCataldo/go-asaas
	//
	// # DOCS
	//
	// Criar uma cobrança: https://docs.asaas.com/reference/criar-nova-cobranca
	//
	// Criar uma cobrança parcelada: https://docs.asaas.com/reference/criar-uma-cobranca-parcelada
	//
	// Criar cobrança com cartão de crédito: https://docs.asaas.com/reference/criar-nova-cobranca-com-cartao-de-credito
	Create(ctx context.Context, body CreateChargeRequest) (*ChargeResponse, error)
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
	// error = not nil
	//
	// Se o parâmetro de retorno error não estiver nil quer dizer que ocorreu um erro inesperado
	// na lib go-asaas.
	//
	// Se isso acontecer por favor report o erro no repositório: https://github.com/GabrielHCataldo/go-asaas
	//
	// # DOCS
	//
	// Pagar uma cobrança com cartão de crédito: https://docs.asaas.com/reference/pagar-uma-cobranca-com-cartao-de-credito
	PayWithCreditCard(ctx context.Context, chargeId string, body PayWithCreditCardRequest) (*ChargeResponse, error)
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
	// error = not nil
	//
	// Se o parâmetro de retorno error não estiver nil quer dizer que ocorreu um erro inesperado
	// na lib go-asaas.
	//
	// Se isso acontecer por favor report o erro no repositório: https://github.com/GabrielHCataldo/go-asaas
	//
	// # DOCS
	//
	// Atualizar cobrança existente: https://docs.asaas.com/reference/atualizar-cobranca-existente
	UpdateById(ctx context.Context, chargeId string, body UpdateChargeRequest) (*ChargeResponse, error)
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
	// error = not nil
	//
	// Se o parâmetro de retorno error não estiver nil quer dizer que ocorreu um erro inesperado
	// na lib go-asaas.
	//
	// Se isso acontecer por favor report o erro no repositório: https://github.com/GabrielHCataldo/go-asaas
	//
	// # DOCS
	//
	// Excluir cobrança: https://docs.asaas.com/reference/remover-cobranca
	DeleteById(ctx context.Context, chargeId string) (*DeleteResponse, error)
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
	// error = not nil
	//
	// Se o parâmetro de retorno error não estiver nil quer dizer que ocorreu um erro inesperado
	// na lib go-asaas.
	//
	// Se isso acontecer por favor report o erro no repositório: https://github.com/GabrielHCataldo/go-asaas
	//
	// # DOCS
	//
	// Restaurar cobrança removida: https://docs.asaas.com/reference/restaurar-cobranca-removida
	RestoreById(ctx context.Context, chargeId string) (*ChargeResponse, error)
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
	// error = not nil
	//
	// Se o parâmetro de retorno error não estiver nil quer dizer que ocorreu um erro inesperado
	// na lib go-asaas.
	//
	// Se isso acontecer por favor report o erro no repositório: https://github.com/GabrielHCataldo/go-asaas
	//
	// # DOCS
	//
	// Estornar cobrança: https://docs.asaas.com/reference/estornar-cobranca
	RefundById(ctx context.Context, chargeId string, body RefundRequest) (*ChargeResponse, error)
	// ReceiveInCashById (Confirmar recebimento em dinheiro)
	//
	// Quando um cliente fizer o pagamento de uma cobrança diretamente para você, sem que esse pagamento seja processado
	// pelo Asaas, utilize este método para definir a cobrança como recebida em dinheiro. Esta opção permite manter
	// seu histórico consistente no sistema, mas não gera saldo ou faz qualquer alteração financeira em sua conta.
	// Ao confirmar um recebimento em dinheiro de uma cobrança que possua uma negativação em andamento uma taxa de
	// ativação de serviço de negativação poderá ser cobrada. Verifique essa taxa no campo receivedInCashFeeValue
	// localizada no retorno do objeto NegativityResponse.
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
	// error = not nil
	//
	// Se o parâmetro de retorno error não estiver nil quer dizer que ocorreu um erro inesperado
	// na lib go-asaas.
	//
	// Se isso acontecer por favor report o erro no repositório: https://github.com/GabrielHCataldo/go-asaas
	//
	// # DOCS
	//
	// Confirmar recebimento em dinheiro: https://docs.asaas.com/reference/confirmar-recebimento-em-dinheiro
	ReceiveInCashById(ctx context.Context, chargeId string, body ChargeReceiveInCashRequest) (*ChargeResponse, error)
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
	// error = not nil
	//
	// Se o parâmetro de retorno error não estiver nil quer dizer que ocorreu um erro inesperado
	// na lib go-asaas.
	//
	// Se isso acontecer por favor report o erro no repositório: https://github.com/GabrielHCataldo/go-asaas
	//
	// # DOCS
	//
	// Desfazer confirmação de recebimento em dinheiro: https://docs.asaas.com/reference/desfazer-confirmacao-de-recebimento-em-dinheiro
	UndoReceivedInCashById(ctx context.Context, chargeId string) (*ChargeResponse, error)
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
	// error = not nil
	//
	// Se o parâmetro de retorno error não estiver nil quer dizer que ocorreu um erro inesperado
	// na lib go-asaas.
	//
	// Se isso acontecer por favor report o erro no repositório: https://github.com/GabrielHCataldo/go-asaas
	//
	// # DOCS
	//
	// Fazer upload de documentos da cobrança: https://docs.asaas.com/reference/fazer-upload-de-documentos-da-cobranca
	UploadDocumentById(ctx context.Context, chargeId string, body UploadChargeDocumentRequest) (*ChargeDocumentResponse,
		error)
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
	// error = not nil
	//
	// Se o parâmetro de retorno error não estiver nil quer dizer que ocorreu um erro inesperado
	// na lib go-asaas.
	//
	// Se isso acontecer por favor report o erro no repositório: https://github.com/GabrielHCataldo/go-asaas
	//
	// # DOCS
	//
	// Atualizar definições de um documento da cobrança: https://docs.asaas.com/reference/atualizar-definicoes-de-um-documento-da-cobranca
	UpdateDocumentDefinitionsById(ctx context.Context, chargeId, docId string, body UpdateChargeDocumentDefinitionsRequest) (
		*ChargeDocumentResponse, error)
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
	// error = not nil
	//
	// Se o parâmetro de retorno error não estiver nil quer dizer que ocorreu um erro inesperado
	// na lib go-asaas.
	//
	// Se isso acontecer por favor report o erro no repositório: https://github.com/GabrielHCataldo/go-asaas
	//
	// # DOCS
	//
	// Excluir documento de uma cobrança: https://docs.asaas.com/reference/excluir-documento-de-uma-cobranca
	DeleteDocumentById(ctx context.Context, chargeId, docId string) (*DeleteResponse, error)
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
	// ChargeResponse = not nil
	//
	// Error = nil
	//
	// ChargeResponse.IsNoContent() = true
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
	// error = not nil
	//
	// Se o parâmetro de retorno error não estiver nil quer dizer que ocorreu um erro inesperado
	// na lib go-asaas.
	//
	// Se isso acontecer por favor report o erro no repositório: https://github.com/GabrielHCataldo/go-asaas
	//
	// # DOCS
	//
	// Recuperar uma única cobrança: https://docs.asaas.com/reference/recuperar-uma-unica-cobranca
	GetById(ctx context.Context, chargeId string) (*ChargeResponse, error)
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
	// error = not nil
	//
	// Se o parâmetro de retorno error não estiver nil quer dizer que ocorreu um erro inesperado
	// na lib go-asaas.
	//
	// Se isso acontecer por favor report o erro no repositório: https://github.com/GabrielHCataldo/go-asaas
	//
	// # DOCS
	//
	// Recuperando limites de cobranças: https://docs.asaas.com/reference/recuperando-limites-de-cobrancças
	GetCreationLimit(ctx context.Context) (*ChargeCreationLimitResponse, error)
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
	// error = not nil
	//
	// Se o parâmetro de retorno error não estiver nil quer dizer que ocorreu um erro inesperado
	// na lib go-asaas.
	//
	// Se isso acontecer por favor report o erro no repositório: https://github.com/GabrielHCataldo/go-asaas
	//
	// # DOCS
	//
	// Recuperar uma única cobrança: https://docs.asaas.com/reference/recuperar-uma-unica-cobranca
	GetStatusById(ctx context.Context, chargeId string) (*ChargeStatusResponse, error)
	// GetIdentificationFieldById (Obter linha digitável do boleto)
	//
	// A linha digitável do boleto é a representação numérica do código de barras.
	// Essa informação pode ser disponibilizada ao seu cliente para pagamento do boleto diretamente no Internet Banking.
	// Ao gerar uma cobrança com as formas de pagamento BOLETO ou UNDEFINED, a linha digitável pode ser recuperada.
	// Para recuperar a linha digitável do boleto, é necessário informar o ID da cobrança que o Asaas retornou no momento
	// da criação. Como retorno, você receberá a linha digitável.
	//
	// # Resposta: 200
	//
	// IdentificationFieldResponse = not nil
	//
	// Error = nil
	//
	// IdentificationFieldResponse.IsSuccess() = true
	//
	// Possui os valores de resposta de sucesso segunda a documentação.
	//
	// # Resposta: 404
	//
	// IdentificationFieldResponse = not nil
	//
	// Error = nil
	//
	// IdentificationFieldResponse.IsNoContent() = true
	//
	// ID(s) informado no parâmetro não foi encontrado.
	//
	// # Resposta: 400/401/500
	//
	// IdentificationFieldResponse = not nil
	//
	// Error = nil
	//
	// IdentificationFieldResponse.IsFailure() = true
	//
	// Para qualquer outra resposta inesperada da API, possuímos o campo IdentificationFieldResponse.Errors preenchido
	// com as informações de erro, sendo 400 retornado da API Asaas com as instruções de requisição conforme a documentação,
	// diferente disso retornará uma mensagem padrão no index 0 do slice com campo ErrorResponse.Code retornando a
	// descrição status http (Ex: "401 Unauthorized") e no campo ErrorResponse.Description retornará com o valor
	// "response status code not expected".
	//
	// # Error
	//
	// IdentificationFieldResponse = nil
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
	// Obter linha digitável do boleto: https://docs.asaas.com/reference/obter-linha-digitavel-do-boleto
	GetIdentificationFieldById(ctx context.Context, chargeId string) (*IdentificationFieldResponse, error)
	// GetPixQrCodeById (Obter QR Code para pagamentos via pix)
	//
	// O recebimento via pix é um meio rápido, eficaz e seguro para que sua empresa receba as cobranças de seus clientes.
	// Ao gerar uma cobrança com as formas de pagamento PIX, BOLETO ou UNDEFINED o pagamento via Pix é habilitado.
	// Uma das maiores vantagens dessa forma de pagamento é que ocorre de forma instantânea, ou seja, assim que o
	// pagamento for realizado o saldo é disponibilizado em sua conta Asaas. Você pode ler mais sobre o Pix aqui.
	//
	// # Resposta: 200
	//
	// ChargePixQrCodeResponse = not nil
	//
	// Error = nil
	//
	// ChargePixQrCodeResponse.IsSuccess() = true
	//
	// Possui os valores de resposta de sucesso segunda a documentação.
	//
	// # Resposta: 404
	//
	// ChargePixQrCodeResponse = not nil
	//
	// Error = nil
	//
	// ChargePixQrCodeResponse.IsNoContent() = true
	//
	// ID(s) informado no parâmetro não foi encontrado.
	//
	// # Resposta: 400/401/500
	//
	// ChargePixQrCodeResponse = not nil
	//
	// Error = nil
	//
	// ChargePixQrCodeResponse.IsFailure() = true
	//
	// Para qualquer outra resposta inesperada da API, possuímos o campo ChargePixQrCodeResponse.Errors preenchido com as informações
	// de erro, sendo 400 retornado da API Asaas com as instruções de requisição conforme a documentação,
	// diferente disso retornará uma mensagem padrão no index 0 do slice com campo ErrorResponse.Code retornando a
	// descrição status http (Ex: "401 Unauthorized") e no campo ErrorResponse.Description retornará com o valor
	// "response status code not expected".
	//
	// # Error
	//
	// ChargePixQrCodeResponse = nil
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
	// Obter QR Code para pagamentos via pix: https://docs.asaas.com/reference/obter-qr-code-para-pagamentos-via-pix
	GetPixQrCodeById(ctx context.Context, chargeId string) (*ChargePixQrCodeResponse, error)
	// GetDocumentById (Recuperar um único documento da cobrança)
	//
	// Para recuperar um único documento e suas informações, é necessário que você tenha o ID que o Asaas retornou no
	// momento do upload do documento.
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
	// error = not nil
	//
	// Se o parâmetro de retorno error não estiver nil quer dizer que ocorreu um erro inesperado
	// na lib go-asaas.
	//
	// Se isso acontecer por favor report o erro no repositório: https://github.com/GabrielHCataldo/go-asaas
	//
	// # DOCS
	//
	// Recuperar um único documento da cobrança: https://docs.asaas.com/reference/recuperar-um-unico-documento-da-cobranca
	GetDocumentById(ctx context.Context, chargeId, docId string) (*ChargeDocumentResponse, error)
	// GetAll (Listar cobranças)
	//
	// Diferente da recuperação de uma cobrança específica, este método retorna uma lista paginada com todas as
	// cobranças para os filtros informados.
	//
	// # Resposta: 200
	//
	// Pageable(ChargeResponse) = not nil
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
	// Pageable(ChargeResponse) = not nil
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
	// Pageable(ChargeResponse) = nil
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
	// Listar cobranças: https://docs.asaas.com/reference/listar-cobrancas
	GetAll(ctx context.Context, filter GetAllChargesRequest) (*Pageable[ChargeResponse], error)
	// GetAllDocumentsById (Listar documentos de uma cobrança)
	//
	// Para listar os documentos de uma cobrança, é necessário que você tenha o ID da cobrança retornado pelo Asaas.
	//
	// # Resposta: 200
	//
	// Pageable(ChargeDocumentResponse) = not nil
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
	// Pageable(ChargeDocumentResponse) = not nil
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
	// Pageable(ChargeDocumentResponse) = nil
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
	// Listar documentos de uma cobrança: https://docs.asaas.com/reference/listar-documentos-de-uma-cobranca
	GetAllDocumentsById(ctx context.Context, chargeId string, filter PageableDefaultRequest) (
		*Pageable[ChargeDocumentResponse], error)
}

func NewCharge(env Env, accessCode string) Charge {
	logWarning("Charge service running on", env.String())
	return charge{
		env:         env,
		accessToken: accessCode,
	}
}

func (c charge) Create(ctx context.Context, body CreateChargeRequest) (*ChargeResponse, error) {
	c.prepareCreateBodyRequest(&body)
	req := NewRequest[ChargeResponse](ctx, c.env, c.accessToken)
	return req.make(http.MethodPost, "/v3/payments", body)
}

func (c charge) PayWithCreditCard(ctx context.Context, chargeId string, body PayWithCreditCardRequest) (*ChargeResponse,
	error) {
	req := NewRequest[ChargeResponse](ctx, c.env, c.accessToken)
	return req.make(http.MethodPost, fmt.Sprintf(`/v3/payments/%s/payWithCreditCard`, chargeId), body)
}

func (c charge) UpdateById(ctx context.Context, chargeId string, body UpdateChargeRequest) (*ChargeResponse,
	error) {
	req := NewRequest[ChargeResponse](ctx, c.env, c.accessToken)
	return req.make(http.MethodPut, fmt.Sprintf(`/v3/payments/%s`, chargeId), body)
}

func (c charge) DeleteById(ctx context.Context, chargeId string) (*DeleteResponse, error) {
	req := NewRequest[DeleteResponse](ctx, c.env, c.accessToken)
	return req.make(http.MethodDelete, fmt.Sprintf(`/v3/payments/%s`, chargeId), nil)
}

func (c charge) RestoreById(ctx context.Context, chargeId string) (*ChargeResponse, error) {
	req := NewRequest[ChargeResponse](ctx, c.env, c.accessToken)
	return req.make(http.MethodPost, fmt.Sprintf(`/v3/payments/%s/restore`, chargeId), nil)
}

func (c charge) RefundById(ctx context.Context, chargeId string, body RefundRequest) (
	*ChargeResponse, error) {
	req := NewRequest[ChargeResponse](ctx, c.env, c.accessToken)
	return req.make(http.MethodPost, fmt.Sprintf(`/v3/payments/%s/refund`, chargeId), body)
}

func (c charge) ReceiveInCashById(ctx context.Context, chargeId string, body ChargeReceiveInCashRequest) (
	*ChargeResponse, error) {
	req := NewRequest[ChargeResponse](ctx, c.env, c.accessToken)
	return req.make(http.MethodPost, fmt.Sprintf(`/v3/payments/%s/receiveInCash`, chargeId), body)
}

func (c charge) UndoReceivedInCashById(ctx context.Context, chargeId string) (*ChargeResponse, error) {
	req := NewRequest[ChargeResponse](ctx, c.env, c.accessToken)
	return req.make(http.MethodPost, fmt.Sprintf(`/v3/payments/%s/undoReceivedInCash`, chargeId), nil)
}

func (c charge) UploadDocumentById(ctx context.Context, chargeId string, body UploadChargeDocumentRequest) (
	*ChargeDocumentResponse, error) {
	req := NewRequest[ChargeDocumentResponse](ctx, c.env, c.accessToken)
	return req.makeMultipartForm(http.MethodPost, fmt.Sprintf(`/v3/payments/%s/documents`, chargeId), body)
}

func (c charge) UpdateDocumentDefinitionsById(
	ctx context.Context,
	chargeId,
	docId string,
	body UpdateChargeDocumentDefinitionsRequest,
) (*ChargeDocumentResponse, error) {
	req := NewRequest[ChargeDocumentResponse](ctx, c.env, c.accessToken)
	return req.make(http.MethodPut, fmt.Sprintf(`/v3/payments/%s/documents/%v`, chargeId, docId), body)
}

func (c charge) DeleteDocumentById(ctx context.Context, chargeId, docId string) (*DeleteResponse, error) {
	req := NewRequest[DeleteResponse](ctx, c.env, c.accessToken)
	return req.make(http.MethodDelete, fmt.Sprintf(`/v3/payments/%s/documents/%v`, chargeId, docId), nil)
}

func (c charge) GetCreationLimit(ctx context.Context) (*ChargeCreationLimitResponse, error) {
	req := NewRequest[ChargeCreationLimitResponse](ctx, c.env, c.accessToken)
	return req.make(http.MethodGet, "/v3/payments/limits", nil)
}

func (c charge) GetById(ctx context.Context, chargeId string) (*ChargeResponse, error) {
	req := NewRequest[ChargeResponse](ctx, c.env, c.accessToken)
	return req.make(http.MethodGet, fmt.Sprintf(`/v3/payments/%s`, chargeId), nil)
}

func (c charge) GetStatusById(ctx context.Context, chargeId string) (*ChargeStatusResponse, error) {
	req := NewRequest[ChargeStatusResponse](ctx, c.env, c.accessToken)
	return req.make(http.MethodGet, fmt.Sprintf(`/v3/payments/%s/status`, chargeId), nil)
}

func (c charge) GetIdentificationFieldById(ctx context.Context, chargeId string) (*IdentificationFieldResponse,
	error) {
	req := NewRequest[IdentificationFieldResponse](ctx, c.env, c.accessToken)
	return req.make(http.MethodGet, fmt.Sprintf(`/v3/payments/%s/identificationField`, chargeId), nil)
}

func (c charge) GetPixQrCodeById(ctx context.Context, chargeId string) (*ChargePixQrCodeResponse, error) {
	req := NewRequest[ChargePixQrCodeResponse](ctx, c.env, c.accessToken)
	return req.make(http.MethodGet, fmt.Sprintf(`/v3/payments/%s/pixQrCode`, chargeId), nil)
}

func (c charge) GetDocumentById(ctx context.Context, chargeId, docId string) (*ChargeDocumentResponse, error) {
	req := NewRequest[ChargeDocumentResponse](ctx, c.env, c.accessToken)
	return req.make(http.MethodGet, fmt.Sprintf(`/v3/payments/%s/documents/%v`, chargeId, docId), nil)
}

func (c charge) GetAll(ctx context.Context, filter GetAllChargesRequest) (
	*Pageable[ChargeResponse], error) {
	req := NewRequest[Pageable[ChargeResponse]](ctx, c.env, c.accessToken)
	return req.make(http.MethodGet, "/v3/payments", filter)
}

func (c charge) GetAllDocumentsById(ctx context.Context, chargeId string, filter PageableDefaultRequest) (
	*Pageable[ChargeDocumentResponse], error) {
	req := NewRequest[Pageable[ChargeDocumentResponse]](ctx, c.env, c.accessToken)
	return req.make(http.MethodGet, fmt.Sprintf(`/v3/payments/%s/documents`, chargeId), filter)
}

func (c charge) prepareCreateBodyRequest(body *CreateChargeRequest) {
	if body.BillingType != BillingTypeCreditCard {
		body.CreditCard = nil
		body.CreditCardHolderInfo = nil
		body.CreditCardToken = ""
		body.RemoteIp = ""
	}
}
