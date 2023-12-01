package asaas

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

type createPixKeyRequest struct {
	Type PixKeyType `json:"type,omitempty" validate:"required,enum"`
}

type CreatePixKeyStaticRequest struct {
	// Chave que receberá os pagamentos do QrCode
	AddressKey string `json:"addressKey,omitempty"`
	// Descrição do QrCode
	Description string `json:"description,omitempty"`
	// Valor do QrCode, caso não informado o pagador poderá escolher o valor
	Value float64 `json:"value,omitempty" validate:"omitempty,gt=0"`
	// Formato do QrCode
	Format QrCodeFormat `json:"format,omitempty" validate:"omitempty,enum"`
	// Data/Hora de expiração do QrCode, após desta data todos os pagamentos serão recusados.
	ExpirationDate *Datetime `json:"expirationDate,omitempty"`
	// Determina a data de expiração em segundos.
	ExpirationSeconds int `json:"expirationSeconds,omitempty" validate:"omitempty,gt=0"`
	// Define se o QrCode pode ser pago múltiplas vezes, caso não informado o valor padrão é true.
	AllowsMultiplePayments bool `json:"allowsMultiplePayments"`
}

type PayPixQrCodeRequest struct {
	// Payload do QRCode para pagamento (REQUIRED)
	QrCode PixQrCodeRequest `json:"qrCode,omitempty" validate:"required"`
	// Valor a ser pago (REQUIRED)
	Value float64 `json:"value,omitempty" validate:"required,gt=0"`
	// Descrição do pagamento
	Description string `json:"description,omitempty"`
	// Utilizada para realizar agendamento do pagamento
	ScheduleDate *Date `json:"scheduleDate,omitempty" validate:"omitempty,after_now"`
}

type PixQrCodeRequest struct {
	// Payload do QRCode (REQUIRED)
	Payload string `json:"payload,omitempty" validate:"required"`
	// Valor do troco (para QRCode Troco)
	ChangeValue float64 `json:"changeValue,omitempty" validate:"omitempty,gt=0"`
}

type GetAllPixKeysRequest struct {
	// Filtrar pelo status atual da chave
	Status PixKeyStatus `json:"status,omitempty"`
	// Filtrar por um ou mais status das chaves
	StatusList string `json:"statusList,omitempty"`
	// Elemento inicial da lista
	Offset int `json:"offset,omitempty"`
	// Número de elementos da lista (max: 100)
	Limit int `json:"limit,omitempty"`
}

type PixKeyResponse struct {
	Id                    string                `json:"id,omitempty"`
	Key                   string                `json:"key,omitempty"`
	Type                  PixKeyType            `json:"type,omitempty"`
	Status                PixKeyStatus          `json:"status,omitempty"`
	CanBeDeleted          bool                  `json:"canBeDeleted,omitempty"`
	CannotBeDeletedReason string                `json:"cannotBeDeletedReason,omitempty"`
	QrCode                *PixKeyQrCodeResponse `json:"qrCode,omitempty"`
	Errors                []ErrorResponse       `json:"errors,omitempty"`
	DateCreated           *Date                 `json:"dateCreated,omitempty"`
}

type PixTransactionResponse struct {
	Id                    string                      `json:"id,omitempty"`
	Payment               string                      `json:"payment,omitempty"`
	EndToEndIdentifier    string                      `json:"endToEndIdentifier,omitempty"`
	Type                  PixTransactionType          `json:"type,omitempty"`
	Status                PixTransactionStatus        `json:"status,omitempty"`
	Finality              PixTransactionFinality      `json:"finality,omitempty"`
	Value                 float64                     `json:"value,omitempty"`
	ChangeValue           float64                     `json:"changeValue,omitempty"`
	RefundedValue         float64                     `json:"refundedValue,omitempty"`
	EffectiveDate         *Date                       `json:"effectiveDate,omitempty"`
	ScheduledDate         *Date                       `json:"scheduledDate,omitempty"`
	OriginType            PixTransactionOriginType    `json:"originType,omitempty"`
	Description           string                      `json:"description,omitempty"`
	TransactionReceiptUrl string                      `json:"transactionReceiptUrl,omitempty"`
	RefusalReason         string                      `json:"refusalReason,omitempty"`
	CanBeCanceled         bool                        `json:"canBeCanceled,omitempty"`
	OriginalTransaction   string                      `json:"originalTransaction,omitempty"`
	ExternalAccount       *PixExternalAccountResponse `json:"externalAccount,omitempty"`
	QrCode                *PixQrCodeResponse          `json:"qrCode,omitempty"`
	Errors                []ErrorResponse             `json:"errors,omitempty"`
}

type PixCancelTransactionResponse struct {
	Id                    string                      `json:"id,omitempty"`
	Payment               string                      `json:"payment,omitempty"`
	EndToEndIdentifier    string                      `json:"endToEndIdentifier,omitempty"`
	Type                  PixTransactionType          `json:"type,omitempty"`
	Status                PixTransactionStatus        `json:"status,omitempty"`
	Finality              PixTransactionFinality      `json:"finality,omitempty"`
	Value                 float64                     `json:"value,omitempty"`
	ChangeValue           float64                     `json:"changeValue,omitempty"`
	RefundedValue         float64                     `json:"refundedValue,omitempty"`
	EffectiveDate         *Date                       `json:"effectiveDate,omitempty"`
	ScheduledDate         *Date                       `json:"scheduledDate,omitempty"`
	OriginType            PixTransactionOriginType    `json:"originType,omitempty"`
	Description           string                      `json:"description,omitempty"`
	TransactionReceiptUrl string                      `json:"transactionReceiptUrl,omitempty"`
	RefusalReason         string                      `json:"refusalReason,omitempty"`
	CanBeCanceled         bool                        `json:"canBeCanceled,omitempty"`
	OriginalTransaction   string                      `json:"originalTransaction,omitempty"`
	ExternalAccount       *PixExternalAccountResponse `json:"externalAccount,omitempty"`
	QrCode                string                      `json:"qrCode,omitempty"`
	Errors                []ErrorResponse             `json:"errors,omitempty"`
}

type DecodePixQrCodeResponse struct {
	Payload                     string                   `json:"payload,omitempty"`
	Type                        PixTransactionType       `json:"type,omitempty"`
	TransactionOriginType       PixTransactionOriginType `json:"transactionOriginType,omitempty"`
	PixKey                      string                   `json:"pix,omitempty"`
	ConciliationIdentifier      string                   `json:"conciliationIdentifier,omitempty"`
	EndToEndIdentifier          string                   `json:"endToEndIdentifier,omitempty"`
	DueDate                     *Date                    `json:"dueDate,omitempty"`
	ExpirationDate              *Date                    `json:"expirationDate,omitempty"`
	Finality                    PixTransactionFinality   `json:"finality,omitempty"`
	Value                       float64                  `json:"value,omitempty"`
	ChangeValue                 float64                  `json:"changeValue,omitempty"`
	Interest                    float64                  `json:"interest,omitempty"`
	Fine                        float64                  `json:"fine,omitempty"`
	Discount                    float64                  `json:"discount,omitempty"`
	TotalValue                  float64                  `json:"totalValue,omitempty"`
	CanBePaidWithDifferentValue bool                     `json:"canBePaidWithDifferentValue,omitempty"`
	CanBeModifyChangeValue      bool                     `json:"canBeModifyChangeValue,omitempty"`
	Receiver                    *PixReceiverResponse     `json:"receiver,omitempty"`
	Payer                       *PixPayerResponse        `json:"payer,omitempty"`
	Description                 string                   `json:"description,omitempty"`
	Errors                      []ErrorResponse          `json:"errors,omitempty"`
	CanBePaid                   bool                     `json:"canBePaid,omitempty"`
	CannotBePaidReason          string                   `json:"cannotBePaidReason,omitempty"`
}

type PixReceiverResponse struct {
	Ispb        int         `json:"ispb,omitempty"`
	IspbName    string      `json:"ispbName,omitempty"`
	Name        string      `json:"name,omitempty"`
	TradingName string      `json:"tradingName,omitempty"`
	CpfCnpj     string      `json:"cpfCnpj,omitempty"`
	PersonType  PersonType  `json:"personType,omitempty"`
	Agency      string      `json:"agency,omitempty"`
	Account     string      `json:"account,omitempty"`
	AccountType AccountType `json:"accountType,omitempty"`
}

type PixExternalAccountResponse struct {
	Ispb           int    `json:"ispb,omitempty"`
	IspbName       string `json:"ispbName,omitempty"`
	Name           string `json:"name,omitempty"`
	CpfCnpj        string `json:"cpfCnpj,omitempty"`
	AddressKey     string `json:"addressKey,omitempty"`
	AddressKeyType string `json:"addressKeyType,omitempty"`
}

type PixQrCodeResponse struct {
	Payer                  *PixPayerResponse `json:"payer,omitempty"`
	ConciliationIdentifier string            `json:"conciliationIdentifier,omitempty"`
	OriginalValue          float64           `json:"originalValue,omitempty"`
	DueDate                *Date             `json:"dueDate,omitempty"`
	Interest               float64           `json:"interest,omitempty"`
	Fine                   float64           `json:"fine,omitempty"`
	Discount               float64           `json:"discount,omitempty"`
	ExpirationDate         *Date             `json:"expirationDate,omitempty"`
}

type PixPayerResponse struct {
	Name    string `json:"name,omitempty"`
	CpfCnpj string `json:"cpfCnpj,omitempty"`
}

type PixKeyQrCodeResponse struct {
	EncodedImage string `json:"encodedImage,omitempty"`
	Payload      string `json:"payload,omitempty"`
}

type QrCodeResponse struct {
	Id                     string          `json:"id,omitempty"`
	EncodedImage           string          `json:"encodedImage,omitempty"`
	Payload                string          `json:"payload,omitempty"`
	AllowsMultiplePayments bool            `json:"allowsMultiplePayments,omitempty"`
	ExpirationDate         *Date           `json:"expirationDate,omitempty"`
	Errors                 []ErrorResponse `json:"errors,omitempty"`
}

type pix struct {
	env         Env
	accessToken string
}

type Pix interface {
	// PayQrCode (Pagar um QRCode)
	//
	// # Resposta: 200
	//
	// PixTransactionResponse = not nil
	//
	// Error = nil
	//
	// PixTransactionResponse.IsSuccess() = true
	//
	// Possui os valores de resposta de sucesso segunda a documentação.
	//
	// # Resposta: 400/401/500
	//
	// PixTransactionResponse = not nil
	//
	// Error = nil
	//
	// PixTransactionResponse.IsFailure() = true
	//
	// Para qualquer outra resposta inesperada da API, possuímos o campo PixTransactionResponse.Errors preenchido com as informações
	// de erro, sendo 400 retornado da API Asaas com as instruções de requisição conforme a documentação,
	// diferente disso retornará uma mensagem padrão no index 0 do slice com campo ErrorResponse.Code retornando a
	// descrição status http (Ex: "401 Unauthorized") e no campo ErrorResponse.Description retornará com o valor
	// "response status code not expected".
	//
	// # Error
	//
	// PixTransactionResponse = nil
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
	// Pagar um QRCode: https://docs.asaas.com/reference/pagar-um-qrcode
	PayQrCode(ctx context.Context, body PayPixQrCodeRequest) (*PixTransactionResponse, Error)
	// DecodeQrCode (Decodificar um QRCode para pagamento)
	//
	//  Permite decodificar um QRCode através de seu payload.
	//
	// # Resposta: 200
	//
	// DecodePixQrCodeResponse = not nil
	//
	// Error = nil
	//
	// DecodePixQrCodeResponse.IsSuccess() = true
	//
	// Possui os valores de resposta de sucesso segunda a documentação.
	//
	// # Resposta: 400/401/500
	//
	// DecodePixQrCodeResponse = not nil
	//
	// Error = nil
	//
	// DecodePixQrCodeResponse.IsFailure() = true
	//
	// Para qualquer outra resposta inesperada da API, possuímos o campo DecodePixQrCodeResponse.Errors preenchido com as informações
	// de erro, sendo 400 retornado da API Asaas com as instruções de requisição conforme a documentação,
	// diferente disso retornará uma mensagem padrão no index 0 do slice com campo ErrorResponse.Code retornando a
	// descrição status http (Ex: "401 Unauthorized") e no campo ErrorResponse.Description retornará com o valor
	// "response status code not expected".
	//
	// # Error
	//
	// DecodePixQrCodeResponse = nil
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
	// Decodificar um QRCode para pagamento: https://docs.asaas.com/reference/decodificar-um-qrcode-para-pagamento
	DecodeQrCode(ctx context.Context, body PixQrCodeRequest) (*DecodePixQrCodeResponse, Error)
	// CancelTransactionById (Cancelar uma transação agendada)
	//
	//  Permite decodificar um QRCode através de seu payload.
	//
	// # Resposta: 200
	//
	// PixCancelTransactionResponse = not nil
	//
	// Error = nil
	//
	// PixCancelTransactionResponse.IsSuccess() = true
	//
	// Possui os valores de resposta de sucesso segunda a documentação.
	//
	// # Resposta: 400/401/500
	//
	// PixCancelTransactionResponse = not nil
	//
	// Error = nil
	//
	// PixCancelTransactionResponse.IsFailure() = true
	//
	// Para qualquer outra resposta inesperada da API, possuímos o campo PixCancelTransactionResponse.Errors preenchido com as informações
	// de erro, sendo 400 retornado da API Asaas com as instruções de requisição conforme a documentação,
	// diferente disso retornará uma mensagem padrão no index 0 do slice com campo ErrorResponse.Code retornando a
	// descrição status http (Ex: "401 Unauthorized") e no campo ErrorResponse.Description retornará com o valor
	// "response status code not expected".
	//
	// # Error
	//
	// PixCancelTransactionResponse = nil
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
	// Cancelar uma transação agendada: https://docs.asaas.com/reference/cancelar-uma-transacao-agendada
	CancelTransactionById(ctx context.Context, pixTransactionId string) (*PixCancelTransactionResponse, Error)
	// CreateKey (Criar uma chave)
	//
	// Permite a manipulação de chaves aleatórias da sua conta Asaas.
	//
	// # Resposta: 200
	//
	// PixKeyResponse = not nil
	//
	// Error = nil
	//
	// PixKeyResponse.IsSuccess() = true
	//
	// Possui os valores de resposta de sucesso segunda a documentação.
	//
	// # Resposta: 400/401/500
	//
	// PixKeyResponse = not nil
	//
	// Error = nil
	//
	// PixKeyResponse.IsFailure() = true
	//
	// Para qualquer outra resposta inesperada da API, possuímos o campo PixKeyResponse.Errors preenchido com as informações
	// de erro, sendo 400 retornado da API Asaas com as instruções de requisição conforme a documentação,
	// diferente disso retornará uma mensagem padrão no index 0 do slice com campo ErrorResponse.Code retornando a
	// descrição status http (Ex: "401 Unauthorized") e no campo ErrorResponse.Description retornará com o valor
	// "response status code not expected".
	//
	// # Error
	//
	// PixKeyResponse = nil
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
	// Criar uma chave: https://docs.asaas.com/reference/criar-uma-chave
	CreateKey(ctx context.Context) (*PixKeyResponse, Error)
	// CreateStaticKey (Criar QRCode estático)
	//
	// Permite criar um QrCode estático para uma determinada chave. Caso não informado o campo valor, o pagador poderá
	// escolher o valor a ser pago.
	//
	// # Resposta: 200
	//
	// QrCodeResponse = not nil
	//
	// Error = nil
	//
	// QrCodeResponse.IsSuccess() = true
	//
	// Possui os valores de resposta de sucesso segunda a documentação.
	//
	// # Resposta: 400/401/500
	//
	// QrCodeResponse = not nil
	//
	// Error = nil
	//
	// QrCodeResponse.IsFailure() = true
	//
	// Para qualquer outra resposta inesperada da API, possuímos o campo QrCodeResponse.Errors preenchido com as informações
	// de erro, sendo 400 retornado da API Asaas com as instruções de requisição conforme a documentação,
	// diferente disso retornará uma mensagem padrão no index 0 do slice com campo ErrorResponse.Code retornando a
	// descrição status http (Ex: "401 Unauthorized") e no campo ErrorResponse.Description retornará com o valor
	// "response status code not expected".
	//
	// # Error
	//
	// QrCodeResponse = nil
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
	// Criar QRCode estático: https://docs.asaas.com/reference/criar-qrcode-estaticoe
	CreateStaticKey(ctx context.Context, body CreatePixKeyStaticRequest) (*QrCodeResponse, Error)
	// DeleteKeyById (Remover chave)
	//
	// # Resposta: 200
	//
	// PixKeyResponse = not nil
	//
	// Error = nil
	//
	// Se PixKeyResponse.IsSuccess() for true quer dizer que foi excluída.
	//
	// Se caso PixKeyResponse.IsFailure() for true quer dizer que não foi excluída.
	//
	// # Resposta: 404
	//
	// PixKeyResponse = not nil
	//
	// Error = nil
	//
	// PixKeyResponse.IsNoContent() = true
	//
	// ID(s) informado no parâmetro não foi encontrado.
	//
	// # Resposta: 400/401/500
	//
	// PixKeyResponse = not nil
	//
	// Error = nil
	//
	// PixKeyResponse.IsFailure() = true
	//
	// Para qualquer outra resposta inesperada da API, possuímos o campo PixKeyResponse.Errors preenchido com as informações
	// de erro, sendo 400 retornado da API Asaas com as instruções de requisição conforme a documentação,
	// diferente disso retornará uma mensagem padrão no index 0 do slice com campo ErrorResponse.Code retornando a
	// descrição status http (Ex: "401 Unauthorized") e no campo ErrorResponse.Description retornará com o valor
	// "response status code not expected".
	//
	// # Error
	//
	// PixKeyResponse = nil
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
	// Remover chave: https://docs.asaas.com/reference/remover-chave
	DeleteKeyById(ctx context.Context, pixKeyId string) (*PixKeyResponse, Error)
	// GetKeyById (Recuperar uma única chave)
	//
	// # Resposta: 200
	//
	// PixKeyResponse = not nil
	//
	// Error = nil
	//
	// PixKeyResponse.IsSuccess() = true
	//
	// Possui os valores de resposta de sucesso segunda a documentação.
	//
	// # Resposta: 404
	//
	// PixKeyResponse = not nil
	//
	// Error = nil
	//
	// PixKeyResponse.IsNoContent() = true
	//
	// ID(s) informado no parâmetro não foi encontrado.
	//
	// # Resposta: 401/500
	//
	// PixKeyResponse = not nil
	//
	// Error = nil
	//
	// PixKeyResponse.IsFailure() = true
	//
	// Para qualquer outra resposta inesperada da API, possuímos o campo PixKeyResponse.Errors preenchido com as informações
	// de erro, o index 0 do slice com campo ErrorResponse.Code retornando a descrição
	// status http (Ex: "401 Unauthorized") e no campo ErrorResponse.Description retornará com o valor
	// "response status code not expected".
	//
	// # Error
	//
	// PixKeyResponse = nil
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
	// Recuperar uma única chave: https://docs.asaas.com/reference/recuperar-uma-unica-chave
	GetKeyById(ctx context.Context, pixKeyId string) (*PixKeyResponse, Error)
	// GetTransactionById (Recuperar uma única transação)
	//
	// Para recuperar uma transação específica é necessário que você tenha o ID que o Asaas retornou no momento da
	// criação dela.
	//
	// # Resposta: 200
	//
	// PixTransactionResponse = not nil
	//
	// Error = nil
	//
	// PixTransactionResponse.IsSuccess() = true
	//
	// Possui os valores de resposta de sucesso segunda a documentação.
	//
	// # Resposta: 404
	//
	// PixTransactionResponse = not nil
	//
	// Error = nil
	//
	// PixTransactionResponse.IsNoContent() = true
	//
	// ID(s) informado no parâmetro não foi encontrado.
	//
	// # Resposta: 401/500
	//
	// PixTransactionResponse = not nil
	//
	// Error = nil
	//
	// PixTransactionResponse.IsFailure() = true
	//
	// Para qualquer outra resposta inesperada da API, possuímos o campo PixTransactionResponse.Errors preenchido com as informações
	// de erro, o index 0 do slice com campo ErrorResponse.Code retornando a descrição
	// status http (Ex: "401 Unauthorized") e no campo ErrorResponse.Description retornará com o valor
	// "response status code not expected".
	//
	// # Error
	//
	// PixTransactionResponse = nil
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
	// Recuperar uma única transação: https://docs.asaas.com/reference/recuperar-uma-unica-transacao
	GetTransactionById(ctx context.Context, pixTransactionId string) (*PixTransactionResponse, Error)
	// GetAllTransactions (Listar transações)
	//
	// # Resposta: 200
	//
	// Pageable(PixTransactionResponse) = not nil
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
	// Pageable(PixTransactionResponse) = not nil
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
	// Pageable(PixTransactionResponse) = nil
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
	// Listar transações: https://docs.asaas.com/reference/listar-transacoes
	GetAllTransactions(ctx context.Context) (*Pageable[PixTransactionResponse], Error)
	// GetAllKeys (Listar chaves)
	//
	// Podemos listar todas as chaves cadastradas na nossa conta ou somente as que estão em um determinado status.
	//
	// # Resposta: 200
	//
	// Pageable(PixKeyResponse) = not nil
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
	// Pageable(PixKeyResponse) = not nil
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
	// Pageable(PixKeyResponse) = nil
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
	// Listar chaves: https://docs.asaas.com/reference/listar-chaves
	GetAllKeys(ctx context.Context, filter GetAllPixKeysRequest) (*Pageable[PixKeyResponse], Error)
}

func NewPix(env Env, accessToken string) Pix {
	logWarning("Pix service running on", env.String())
	return pix{
		env:         env,
		accessToken: accessToken,
	}
}

func (p pix) PayQrCode(ctx context.Context, body PayPixQrCodeRequest) (*PixTransactionResponse, Error) {
	if err := Validate().Struct(body); err != nil {
		return nil, NewError(ErrorTypeValidation, err)
	}
	req := NewRequest[PixTransactionResponse](ctx, p.env, p.accessToken)
	return req.make(http.MethodPost, "/v3/pix/qrCodes/pay", body)
}

func (p pix) DecodeQrCode(ctx context.Context, body PixQrCodeRequest) (*DecodePixQrCodeResponse, Error) {
	if err := Validate().Struct(body); err != nil {
		return nil, NewError(ErrorTypeValidation, err)
	}
	req := NewRequest[DecodePixQrCodeResponse](ctx, p.env, p.accessToken)
	return req.make(http.MethodPost, "/v3/pix/qrCodes/decode", body)
}

func (p pix) CancelTransactionById(ctx context.Context, pixTransactionId string) (*PixCancelTransactionResponse, Error) {
	req := NewRequest[PixCancelTransactionResponse](ctx, p.env, p.accessToken)
	return req.make(http.MethodPost, fmt.Sprintf("/v3/pix/transactions/%s/cancel", pixTransactionId), nil)
}

func (p pix) CreateKey(ctx context.Context) (*PixKeyResponse, Error) {
	req := NewRequest[PixKeyResponse](ctx, p.env, p.accessToken)
	return req.make(http.MethodPost, "/v3/pix/addressKeys", createPixKeyRequest{Type: PixKeyTypeEvp})
}

func (p pix) CreateStaticKey(ctx context.Context, body CreatePixKeyStaticRequest) (*QrCodeResponse, Error) {
	if err := Validate().Struct(body); err != nil {
		return nil, NewError(ErrorTypeValidation, err)
	} else if !body.Format.IsEnumValid() {
		body.Format = QrCodeFormatAll
	}
	req := NewRequest[QrCodeResponse](ctx, p.env, p.accessToken)
	return req.make(http.MethodPost, "/v3/pix/qrCodes/static", body)
}

func (p pix) DeleteKeyById(ctx context.Context, pixKeyId string) (*PixKeyResponse, Error) {
	req := NewRequest[PixKeyResponse](ctx, p.env, p.accessToken)
	return req.make(http.MethodDelete, fmt.Sprintf("/v3/pix/addressKeys/%s", pixKeyId), nil)
}

func (p pix) GetTransactionById(ctx context.Context, pixTransactionId string) (*PixTransactionResponse, Error) {
	req := NewRequest[PixTransactionResponse](ctx, p.env, p.accessToken)
	urlValues := url.Values{"id": []string{pixTransactionId}}
	return req.make(http.MethodGet, "/v3/pix/transactions?"+urlValues.Encode(), nil)
}

func (p pix) GetKeyById(ctx context.Context, pixKeyId string) (*PixKeyResponse, Error) {
	req := NewRequest[PixKeyResponse](ctx, p.env, p.accessToken)
	return req.make(http.MethodGet, fmt.Sprintf("/v3/pix/addressKeys/%s", pixKeyId), nil)
}

func (p pix) GetAllTransactions(ctx context.Context) (*Pageable[PixTransactionResponse], Error) {
	req := NewRequest[Pageable[PixTransactionResponse]](ctx, p.env, p.accessToken)
	return req.make(http.MethodGet, "/v3/pix/transactions", nil)
}

func (p pix) GetAllKeys(ctx context.Context, filter GetAllPixKeysRequest) (*Pageable[PixKeyResponse], Error) {
	req := NewRequest[Pageable[PixKeyResponse]](ctx, p.env, p.accessToken)
	return req.make(http.MethodGet, "/v3/pix/addressKeys", filter)
}
