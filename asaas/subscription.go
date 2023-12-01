package asaas

import (
	"context"
	"fmt"
	"net/http"
)

type CreateSubscriptionRequest struct {
	// Identificador único do cliente no Asaas (REQUIRED)
	Customer string `json:"customer,omitempty" validate:"required"`
	// Forma de pagamento (Default: BillingTypeUndefined)
	BillingType BillingType `json:"billingType,omitempty" validate:"required,enum"`
	// Valor da assinatura
	Value float64 `json:"value,omitempty" validate:"required,gt=0"`
	// Vencimento da primeira mensalidade
	NextDueDate Date `json:"nextDueDate,omitempty" validate:"required,after_now"`
	// Informações de desconto
	Discount *DiscountRequest `json:"discount,omitempty"`
	// Informações de juros para pagamento após o vencimento
	Interest *InterestRequest `json:"interest,omitempty"`
	// Informações de multa para pagamento após o vencimento
	Fine *FineRequest `json:"fine,omitempty"`
	// Periodicidade da cobrança (REQUIRED)
	Cycle SubscriptionCycle `json:"cycle,omitempty" validate:"required,enum"`
	// Descrição da assinatura (máx. 500 caracteres)
	Description string `json:"description,omitempty" validate:"omitempty,lte=500"`
	// Informações do cartão de crédito (REQUIRED se BillingType = BillingTypeCreditCard e se CreditCardToken não for informado)
	CreditCard *CreditCardRequest `json:"creditCard,omitempty"`
	// Informações do titular do cartão de crédito (REQUIRED se BillingType = BillingTypeCreditCard e se CreditCardToken não for informado)
	CreditCardHolderInfo *CreditCardHolderInfoRequest `json:"creditCardHolderInfo,omitempty"`
	// Token do cartão de crédito para uso da funcionalidade de tokenização de cartão de crédito
	CreditCardToken string `json:"creditCardToken,omitempty"`
	// Data limite para vencimento das mensalidades
	EndDate *Date `json:"endDate,omitempty" validate:"omitempty,after_now"`
	// Número máximo de mensalidades a serem geradas para esta assinatura
	MaxPayments int `json:"maxPayments,omitempty" validate:"omitempty,gt=0"`
	// Identificador da assinatura no seu sistema
	ExternalReference string `json:"externalReference,omitempty"`
	// Informações de split
	Split []SplitRequest `json:"split,omitempty"`
	// IP de onde o cliente está fazendo a compra. Não deve ser informado o IP do seu servidor (REQUIRED se BillingType = BillingTypeCreditCard)
	RemoteIp string `json:"remoteIp,omitempty"`
}

type UpdateSubscriptionRequest struct {
	// Forma de pagamento
	BillingType BillingType `json:"billingType,omitempty" validate:"omitempty,enum"`
	// Valor da assinatura
	Value float64 `json:"value,omitempty" validate:"omitempty,gt=0"`
	// Status da assinatura
	Status SubscriptionStatus `json:"status,omitempty" validate:"omitempty,enum"`
	// Vencimento da próxima mensalidade
	NextDueDate *Date `json:"nextDueDate,omitempty"`
	// Informações de desconto
	Discount *DiscountRequest `json:"discount,omitempty"`
	// Informações de juros para pagamento após o vencimento
	Interest *InterestRequest `json:"interest,omitempty"`
	// Informações de multa para pagamento após o vencimento
	Fine *FineRequest `json:"fine,omitempty"`
	// Periodicidade da cobrança
	Cycle SubscriptionCycle `json:"cycle,omitempty" validate:"omitempty,enum"`
	// Descrição da assinatura (máx. 500 caracteres)
	Description string `json:"description,omitempty" validate:"omitempty,lte=500"`
	// Informações do cartão de crédito (REQUIRED se BillingType = BillingTypeCreditCard e se CreditCardToken não for informado)
	CreditCard *CreditCardRequest `json:"creditCard,omitempty"`
	// Informações do titular do cartão de crédito (REQUIRED se BillingType = BillingTypeCreditCard e se CreditCardToken não for informado)
	CreditCardHolderInfo *CreditCardHolderInfoRequest `json:"creditCardHolderInfo,omitempty"`
	// Token do cartão de crédito para uso da funcionalidade de tokenização de cartão de crédito
	CreditCardToken string `json:"creditCardToken,omitempty"`
	// Data de validade da assinatura
	EndDate *Date `json:"endDate,omitempty"`
	// True para atualizar mensalidades já existentes com o novo valor ou forma de pagamento
	UpdatePendingPayments bool `json:"updatePendingPayments,omitempty"`
	// Identificador da assinatura no seu sistema
	ExternalReference string `json:"externalReference,omitempty"`
}

type GetAllSubscriptionsRequest struct {
	// Filtrar pelo Identificador único do cliente
	Customer string `json:"customer,omitempty"`
	// Filtrar pelo nome do grupo de cliente
	CustomerGroupName string `json:"customerGroupName,omitempty"`
	// Filtrar por forma de pagamento
	BillingType BillingType `json:"billingType,omitempty"`
	// Filtrar pelo status
	Status SubscriptionStatus `json:"status,omitempty"`
	// Envie true para retornar somente as assinaturas removidas
	DeletedOnly bool `json:"deletedOnly,omitempty"`
	// Envie true para recuperar também as assinaturas removidas
	IncludeDeleted bool `json:"includeDeleted,omitempty"`
	// Filtrar pelo Identificador do seu sistema
	ExternalReference string `json:"externalReference,omitempty"`
	// Ordem crescente ou decrescente
	Order Order `json:"order,omitempty"`
	// Por qual campo será ordenado
	Sort SortSubscriptionField `json:"sort,omitempty"`
	// Elemento inicial da lista
	Offset int `json:"offset,omitempty"`
	// Número de elementos da lista (max: 100)
	Limit int `json:"limit,omitempty"`
}

type GetAllSubscriptionInvoicesRequest struct {
	// Filtrar a partir de uma data de emissão
	EffectiveDateGE *Date `json:"effectiveDate[ge],omitempty"`
	// Filtrar até uma data de emissão
	EffectiveDateLE *Date `json:"effectiveDate[le],omitempty"`
	// Identificador da nota fiscal no seu sistema
	ExternalReference string `json:"externalReference,omitempty"`
	// Status da nota fiscal
	Status InvoiceStatus `json:"status,omitempty"`
	// Filtrar pelo identificador único do cliente
	Customer string `json:"customer,omitempty"`
	// Elemento inicial da lista
	Offset int `json:"offset,omitempty"`
	// Número de elementos da lista (max: 100)
	Limit int `json:"limit,omitempty"`
}

type GetAllChargesBySubscriptionRequest struct {
	// Filtrar por status das cobranças
	Status ChargeStatus `json:"status,omitempty"`
}

type SubscriptionPaymentBookRequest struct {
	// Mês final para geração do carnê (REQUIRED)
	Month int `json:"month,omitempty" validate:"required,gte=1,lte=12"`
	// Ano final para geração do carnê (REQUIRED)
	Year int `json:"year,omitempty" validate:"required,gt=0"`
	// Filtrar pelo nome da coluna
	Sort SortPaymentBookField `json:"sort,omitempty" validate:"omitempty,enum"`
	// Ordenação da coluna
	Order Order `json:"order,omitempty" validate:"omitempty,enum"`
}

type SubscriptionResponse struct {
	Id                string             `json:"id,omitempty"`
	Customer          string             `json:"customer,omitempty"`
	Status            SubscriptionStatus `json:"status,omitempty"`
	Refunds           []RefundResponse   `json:"refunds,omitempty"`
	BillingType       BillingType        `json:"billingType,omitempty"`
	Value             float64            `json:"value,omitempty"`
	NextDueDate       *Date              `json:"nextDueDate,omitempty"`
	Cycle             SubscriptionCycle  `json:"cycle,omitempty"`
	Discount          *DiscountResponse  `json:"discount,omitempty"`
	Interest          *InterestResponse  `json:"interest,omitempty"`
	Fine              *FineResponse      `json:"fine,omitempty"`
	Description       string             `json:"description,omitempty"`
	EndDate           *Date              `json:"endDate,omitempty"`
	MaxPayments       int                `json:"maxPayments,omitempty"`
	ExternalReference string             `json:"externalReference,omitempty"`
	Deleted           bool               `json:"deleted,omitempty"`
	Errors            []ErrorResponse    `json:"errors,omitempty"`
	DateCreated       *Date              `json:"dateCreated,omitempty"`
}

type subscription struct {
	env         Env
	accessToken string
}

type Subscription interface {
	// Create (Criar nova assinatura)
	//
	// Ao criar a assinatura a primeira mensalidade será gerada vencendo na data enviada no parâmetro nextDueDate.
	//
	// # Resposta: 200
	//
	// SubscriptionResponse = not nil
	//
	// Error = nil
	//
	// SubscriptionResponse.IsSuccess() = true
	//
	// Possui os valores de resposta de sucesso segunda a documentação.
	//
	// # Resposta: 400/401/500
	//
	// SubscriptionResponse = not nil
	//
	// Error = nil
	//
	// SubscriptionResponse.IsFailure() = true
	//
	// Para qualquer outra resposta inesperada da API, possuímos o campo SubscriptionResponse.Errors preenchido com as informações
	// de erro, sendo 400 retornado da API Asaas com as instruções de requisição conforme a documentação,
	// diferente disso retornará uma mensagem padrão no index 0 do slice com campo ErrorResponse.Code retornando a
	// descrição status http (Ex: "401 Unauthorized") e no campo ErrorResponse.Description retornará com o valor
	// "response status code not expected".
	//
	// # Error
	//
	// SubscriptionResponse = nil
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
	// Criar nova assinatura: https://docs.asaas.com/reference/criar-nova-assinatura
	//
	// Criar assinatura com cartão de crédito: https://docs.asaas.com/reference/criar-assinatura-com-cartao-de-credito
	Create(ctx context.Context, body CreateSubscriptionRequest) (*SubscriptionResponse, Error)
	// CreateInvoiceSettingById (Criar configuração para emissão de Notas Fiscais)
	//
	// # Resposta: 200
	//
	// InvoiceSettingResponse = not nil
	//
	// Error = nil
	//
	// InvoiceSettingResponse.IsSuccess() = true
	//
	// Possui os valores de resposta de sucesso segunda a documentação.
	//
	// # Resposta: 404
	//
	// InvoiceSettingResponse = not nil
	//
	// Error = nil
	//
	// InvoiceSettingResponse.IsNoContent() = true
	//
	// ID(s) informado no parâmetro não foi encontrado.
	//
	// # Resposta: 400/401/500
	//
	// InvoiceSettingResponse = not nil
	//
	// Error = nil
	//
	// InvoiceSettingResponse.IsFailure() = true
	//
	// Para qualquer outra resposta inesperada da API, possuímos o campo InvoiceSettingResponse.Errors preenchido com as informações
	// de erro, sendo 400 retornado da API Asaas com as instruções de requisição conforme a documentação,
	// diferente disso retornará uma mensagem padrão no index 0 do slice com campo ErrorResponse.Code retornando a
	// descrição status http (Ex: "401 Unauthorized") e no campo ErrorResponse.Description retornará com o valor
	// "response status code not expected".
	//
	// # Error
	//
	// InvoiceSettingResponse = nil
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
	// Criar configuração para emissão de Notas Fiscais: https://docs.asaas.com/reference/criar-configuracao-para-emissao-de-notas-fiscais
	CreateInvoiceSettingById(ctx context.Context, subscriptionId string, body SaveInvoiceSettingRequest) (
		*InvoiceSettingResponse, Error)
	// UpdateById (Atualizar assinatura existente)
	//
	// Ao atualizar uma assinatura, o parâmetro nextDueDate permite indicar o vencimento da próxima mensalidade a
	// ser gerada, ou seja, não atualiza o vencimento da mensalidade já gerada.
	//
	// Além disso, ao atualizar o valor da assinatura ou forma de pagamento serão somente afetadas
	// mensalidades futuras. Para atualizar as mensalidades já existentes com a nova forma de pagamento e/ou novo valor,
	// é necessário passar o parâmetro updatePendingPayments: true.
	//
	// # Resposta: 200
	//
	// SubscriptionResponse = not nil
	//
	// Error = nil
	//
	// SubscriptionResponse.IsSuccess() = true
	//
	// Possui os valores de resposta de sucesso segunda a documentação.
	//
	// # Resposta: 404
	//
	// SubscriptionResponse = not nil
	//
	// Error = nil
	//
	// SubscriptionResponse.IsNoContent() = true
	//
	// ID(s) informado no parâmetro não foi encontrado.
	//
	// # Resposta: 400/401/500
	//
	// SubscriptionResponse = not nil
	//
	// Error = nil
	//
	// SubscriptionResponse.IsFailure() = true
	//
	// Para qualquer outra resposta inesperada da API, possuímos o campo SubscriptionResponse.Errors preenchido com as informações
	// de erro, sendo 400 retornado da API Asaas com as instruções de requisição conforme a documentação,
	// diferente disso retornará uma mensagem padrão no index 0 do slice com campo ErrorResponse.Code retornando a
	// descrição status http (Ex: "401 Unauthorized") e no campo ErrorResponse.Description retornará com o valor
	// "response status code not expected".
	//
	// # Error
	//
	// SubscriptionResponse = nil
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
	// Atualizar assinatura existente: https://docs.asaas.com/reference/atualizar-assinatura-existente
	UpdateById(ctx context.Context, subscriptionId string, body UpdateSubscriptionRequest) (*SubscriptionResponse, Error)
	// UpdateInvoiceSettingsById (Atualizar configuração para emissão de Notas Fiscais)
	//
	// A nova configuração apenas será aplicada nas notas fiscais das próximas cobranças da assinatura,
	// ou as que ainda não possuem uma nota fiscal criada.
	//
	// # Resposta: 200
	//
	// InvoiceSettingResponse = not nil
	//
	// Error = nil
	//
	// InvoiceSettingResponse.IsSuccess() = true
	//
	// Possui os valores de resposta de sucesso segunda a documentação.
	//
	// # Resposta: 404
	//
	// InvoiceSettingResponse = not nil
	//
	// Error = nil
	//
	// InvoiceSettingResponse.IsNoContent() = true
	//
	// ID(s) informado no parâmetro não foi encontrado.
	//
	// # Resposta: 400/401/500
	//
	// InvoiceSettingResponse = not nil
	//
	// Error = nil
	//
	// InvoiceSettingResponse.IsFailure() = true
	//
	// Para qualquer outra resposta inesperada da API, possuímos o campo InvoiceSettingResponse.Errors preenchido com as informações
	// de erro, sendo 400 retornado da API Asaas com as instruções de requisição conforme a documentação,
	// diferente disso retornará uma mensagem padrão no index 0 do slice com campo ErrorResponse.Code retornando a
	// descrição status http (Ex: "401 Unauthorized") e no campo ErrorResponse.Description retornará com o valor
	// "response status code not expected".
	//
	// # Error
	//
	// InvoiceSettingResponse = nil
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
	// Atualizar configuração para emissão de Notas Fiscais: https://docs.asaas.com/reference/atualizar-configuracao-para-emissao-de-notas-fiscais
	UpdateInvoiceSettingsById(ctx context.Context, subscriptionId string, body UpdateInvoiceSettingRequest) (
		*InvoiceSettingResponse, Error)
	// DeleteById (Remover assinatura)
	//
	// Ao remover uma assinatura, as mensalidades aguardando pagamento ou vencidas também são removidas.
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
	// Remover assinatura: https://docs.asaas.com/reference/remover-assinatura
	DeleteById(ctx context.Context, subscriptionId string) (*DeleteResponse, Error)
	// DeleteInvoiceSettingById (Remover configuração para emissão de Notas Fiscais)
	//
	// Ao remover a configuração, todas as notas fiscais agendadas para as cobranças desta assinatura serão canceladas
	// e sua geração automática será interrompida.
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
	// Remover configuração para emissão de Notas Fiscais: https://docs.asaas.com/reference/remover-configuracao-para-emissao-de-notas-fiscais
	DeleteInvoiceSettingById(ctx context.Context, subscriptionId string) (*DeleteResponse, Error)
	// GetById (Recuperar uma única assinatura)
	//
	// Para recuperar uma assinatura específica é necessário que você tenha o ID que o Asaas retornou no momento da criação dela.
	//
	// Para recuperar as cobranças de uma assinatura utilize GetAllChargesBySubscription
	//
	// # Resposta: 200
	//
	// SubscriptionResponse = not nil
	//
	// Error = nil
	//
	// SubscriptionResponse.IsSuccess() = true
	//
	// Possui os valores de resposta de sucesso segunda a documentação.
	//
	// # Resposta: 404
	//
	// SubscriptionResponse = not nil
	//
	// Error = nil
	//
	// SubscriptionResponse.IsNoContent() = true
	//
	// ID(s) informado no parâmetro não foi encontrado.
	//
	// # Resposta: 401/500
	//
	// SubscriptionResponse = not nil
	//
	// Error = nil
	//
	// SubscriptionResponse.IsFailure() = true
	//
	// Para qualquer outra resposta inesperada da API, possuímos o campo SubscriptionResponse.Errors preenchido com as informações
	// de erro, o index 0 do slice com campo ErrorResponse.Code retornando a descrição
	// status http (Ex: "401 Unauthorized") e no campo ErrorResponse.Description retornará com o valor
	// "response status code not expected".
	//
	// # Error
	//
	// SubscriptionResponse = nil
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
	// Recuperar uma única assinatura: https://docs.asaas.com/reference/recuperar-uma-unica-assinatura
	GetById(ctx context.Context, subscriptionId string) (*SubscriptionResponse, Error)
	// GetInvoiceSettingById (Recuperar configuração para emissão de Notas Fiscais)
	//
	// Para recuperar a configuração de emissão de notas fiscais de uma assinatura é necessário que você tenha criada
	// ela anteriormente e possua o ID da assinatura que o Asaas retornou no momento da criação dela.
	//
	// # Resposta: 200
	//
	// InvoiceSettingResponse = not nil
	//
	// Error = nil
	//
	// InvoiceSettingResponse.IsSuccess() = true
	//
	// Possui os valores de resposta de sucesso segunda a documentação.
	//
	// # Resposta: 404
	//
	// InvoiceSettingResponse = not nil
	//
	// Error = nil
	//
	// InvoiceSettingResponse.IsNoContent() = true
	//
	// ID(s) informado no parâmetro não foi encontrado.
	//
	// # Resposta: 401/500
	//
	// InvoiceSettingResponse = not nil
	//
	// Error = nil
	//
	// InvoiceSettingResponse.IsFailure() = true
	//
	// Para qualquer outra resposta inesperada da API, possuímos o campo InvoiceSettingResponse.Errors preenchido com as informações
	// de erro, o index 0 do slice com campo ErrorResponse.Code retornando a descrição
	// status http (Ex: "401 Unauthorized") e no campo ErrorResponse.Description retornará com o valor
	// "response status code not expected".
	//
	// # Error
	//
	// InvoiceSettingResponse = nil
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
	// Recuperar configuração para emissão de Notas Fiscais: https://docs.asaas.com/reference/recuperar-configuracao-para-emissao-de-notas-fiscais
	GetInvoiceSettingById(ctx context.Context, subscriptionId string) (*InvoiceSettingResponse, Error)
	// GetPaymentBookById (Gerar carnê de assinatura)
	//
	// Para gerar os carnês gerados a partir de uma assinatura em formato PDF, é necessário que você tenha o
	// ID da assinatura retornado pelo Asaas.
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
	// Gerar carnê de assinatura: https://docs.asaas.com/reference/gerar-carne-de-assinatura
	GetPaymentBookById(ctx context.Context, subscriptionId string, filter SubscriptionPaymentBookRequest) (
		*FileTextPlainResponse, Error)
	// GetAll (Listar assinaturas)
	//
	// Diferente da recuperação de uma assinatura específica, este método retorna uma lista paginada com todas
	// as assinaturas para os filtros informados.
	//
	// # Resposta: 200
	//
	// Pageable(SubscriptionResponse) = not nil
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
	// Pageable(SubscriptionResponse) = not nil
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
	// Pageable(SubscriptionResponse) = nil
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
	// Listar assinaturas: https://docs.asaas.com/reference/listar-assinaturas
	GetAll(ctx context.Context, filter GetAllSubscriptionsRequest) (*Pageable[SubscriptionResponse], Error)
	// GetAllChargesBySubscription (Listar cobranças)
	//
	// Para listar as cobranças geradas a partir de uma assinatura utilize esse endpoint.
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
	// Listar cobranças de uma assinatura: https://docs.asaas.com/reference/listar-cobrancas-de-uma-assinatura
	GetAllChargesBySubscription(ctx context.Context, subscriptionId string, filter GetAllChargesBySubscriptionRequest) (
		*Pageable[ChargeResponse], Error)
	// GetAllInvoicesBySubscription (Listar cobranças)
	//
	// Este método retorna uma lista paginada com todas as notas fiscais geradas a partir de cobranças da assinatura informada.
	//
	// É possível também filtrar o status e período de emissão das notas fiscais, como já apresentado na listagem notas fiscais
	//
	// # Resposta: 200
	//
	// Pageable(InvoiceResponse) = not nil
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
	// Pageable(InvoiceResponse) = not nil
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
	// Pageable(InvoiceResponse) = nil
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
	// Listar notas fiscais das cobranças de uma assinatura: https://docs.asaas.com/reference/listar-notas-fiscais-das-cobrancas-de-uma-assinatura
	GetAllInvoicesBySubscription(ctx context.Context, subscriptionId string, filter GetAllSubscriptionInvoicesRequest) (
		*Pageable[InvoiceResponse], Error)
}

func NewSubscription(env Env, accessToken string) Subscription {
	logWarning("Subscription service running on", env.String())
	return subscription{
		env:         env,
		accessToken: accessToken,
	}
}

func (s subscription) Create(ctx context.Context, body CreateSubscriptionRequest) (*SubscriptionResponse, Error) {
	if err := s.validateCreateBodyRequest(body); err != nil {
		return nil, NewError(ErrorTypeValidation, err)
	}
	s.prepareCreateBodyRequest(&body)
	req := NewRequest[SubscriptionResponse](ctx, s.env, s.accessToken)
	return req.make(http.MethodPost, "/v3/subscriptions", body)
}

func (s subscription) CreateInvoiceSettingById(ctx context.Context, subscriptionId string, body SaveInvoiceSettingRequest) (
	*InvoiceSettingResponse, Error) {
	if err := Validate().Struct(body); err != nil {
		return nil, NewError(ErrorTypeValidation, err)
	}
	req := NewRequest[InvoiceSettingResponse](ctx, s.env, s.accessToken)
	return req.make(http.MethodPost, fmt.Sprintf("/v3/subscriptions/%s/invoiceSettings", subscriptionId), body)
}

func (s subscription) UpdateById(ctx context.Context, subscriptionId string, body UpdateSubscriptionRequest) (
	*SubscriptionResponse, Error) {
	if err := Validate().Struct(body); err != nil {
		return nil, NewError(ErrorTypeValidation, err)
	}
	req := NewRequest[SubscriptionResponse](ctx, s.env, s.accessToken)
	return req.make(http.MethodPut, fmt.Sprintf("/v3/subscriptions/%s", subscriptionId), body)
}

func (s subscription) UpdateInvoiceSettingsById(ctx context.Context, subscriptionId string, body UpdateInvoiceSettingRequest) (
	*InvoiceSettingResponse, Error) {
	if err := Validate().Struct(body); err != nil {
		return nil, NewError(ErrorTypeValidation, err)
	}
	req := NewRequest[InvoiceSettingResponse](ctx, s.env, s.accessToken)
	return req.make(http.MethodPut, fmt.Sprintf("/v3/subscriptions/%s/invoiceSettings", subscriptionId), body)
}

func (s subscription) DeleteById(ctx context.Context, subscriptionId string) (*DeleteResponse, Error) {
	req := NewRequest[DeleteResponse](ctx, s.env, s.accessToken)
	return req.make(http.MethodDelete, fmt.Sprintf("/v3/subscriptions/%s", subscriptionId), nil)
}

func (s subscription) DeleteInvoiceSettingById(ctx context.Context, subscriptionId string) (*DeleteResponse, Error) {
	req := NewRequest[DeleteResponse](ctx, s.env, s.accessToken)
	return req.make(http.MethodDelete, fmt.Sprintf("/v3/subscriptions/%s/invoiceSettings", subscriptionId), nil)
}

func (s subscription) GetById(ctx context.Context, subscriptionId string) (*SubscriptionResponse, Error) {
	req := NewRequest[SubscriptionResponse](ctx, s.env, s.accessToken)
	return req.make(http.MethodGet, fmt.Sprintf("/v3/subscriptions/%s", subscriptionId), nil)
}

func (s subscription) GetInvoiceSettingById(ctx context.Context, subscriptionId string) (*InvoiceSettingResponse, Error) {
	req := NewRequest[InvoiceSettingResponse](ctx, s.env, s.accessToken)
	return req.make(http.MethodGet, fmt.Sprintf("/v3/subscriptions/%s/invoiceSettings", subscriptionId), nil)
}

func (s subscription) GetPaymentBookById(ctx context.Context, subscriptionId string, filter SubscriptionPaymentBookRequest) (
	*FileTextPlainResponse, Error) {
	if err := Validate().Struct(filter); err != nil {
		return nil, NewError(ErrorTypeValidation, err)
	}
	req := NewRequest[FileTextPlainResponse](ctx, s.env, s.accessToken)
	return req.make(http.MethodGet, fmt.Sprintf("/v3/subscriptions/%s/paymentBook", subscriptionId), filter)
}

func (s subscription) GetAllChargesBySubscription(
	ctx context.Context,
	subscriptionId string,
	filter GetAllChargesBySubscriptionRequest) (*Pageable[ChargeResponse], Error) {
	req := NewRequest[Pageable[ChargeResponse]](ctx, s.env, s.accessToken)
	return req.make(http.MethodGet, fmt.Sprintf("/v3/subscriptions/%s/payments", subscriptionId), filter)
}

func (s subscription) GetAllInvoicesBySubscription(ctx context.Context, subscriptionId string,
	filter GetAllSubscriptionInvoicesRequest) (*Pageable[InvoiceResponse], Error) {
	req := NewRequest[Pageable[InvoiceResponse]](ctx, s.env, s.accessToken)
	return req.make(http.MethodGet, fmt.Sprintf("/v3/subscriptions/%s/invoices", subscriptionId), filter)
}

func (s subscription) GetAll(ctx context.Context, filter GetAllSubscriptionsRequest) (*Pageable[SubscriptionResponse],
	Error) {
	req := NewRequest[Pageable[SubscriptionResponse]](ctx, s.env, s.accessToken)
	return req.make(http.MethodGet, "/v3/subscriptions", filter)
}

func (s subscription) validateCreateBodyRequest(body CreateSubscriptionRequest) error {
	if err := Validate().Struct(body); err != nil {
		return err
	}
	return validateBillingBody(body.BillingType, body.CreditCard, body.CreditCardHolderInfo, body.CreditCardToken,
		body.RemoteIp)
}

func (s subscription) prepareCreateBodyRequest(body *CreateSubscriptionRequest) {
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
