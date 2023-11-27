package asaas

import (
	"context"
	berrors "errors"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

type createPixKeyRequest struct {
	Type PixKeyType `json:"type,omitempty" validate:"required,enum"`
}

type CreatePixKeyStaticRequest struct {
	AddressKey             string       `json:"addressKey,omitempty"`
	Description            string       `json:"description,omitempty"`
	Value                  float64      `json:"value,omitempty" validate:"omitempty,gt=0"`
	Format                 QrCodeFormat `json:"format,omitempty" validate:"omitempty,enum"`
	ExpirationDate         DateTime     `json:"expirationDate,omitempty"`
	ExpirationSeconds      int          `json:"expirationSeconds,omitempty" validate:"omitempty,gt=0"`
	AllowsMultiplePayments bool         `json:"allowsMultiplePayments"`
}

type PayPixQrCodeRequest struct {
	QrCode       PixQrCodeRequest `json:"qrCode,omitempty" validate:"required"`
	Value        float64          `json:"value,omitempty" validate:"required,gt=0"`
	Description  string           `json:"description,omitempty"`
	ScheduleDate Date             `json:"scheduleDate,omitempty"`
}

type PixQrCodeRequest struct {
	Payload     string  `json:"payload,omitempty" validate:"required"`
	ChangeValue float64 `json:"changeValue,omitempty" validate:"omitempty,gt=0"`
}

type GetAllPixKeysRequest struct {
	Status     PixKeyStatus `json:"status,omitempty"`
	StatusList string       `json:"statusList,omitempty"`
	Offset     int          `json:"offset,omitempty"`
	Limit      int          `json:"limit,omitempty"`
}

type PixKeyResponse struct {
	ID                    string                `json:"id,omitempty"`
	Key                   string                `json:"key,omitempty"`
	Type                  PixKeyType            `json:"type,omitempty"`
	Status                PixKeyStatus          `json:"status,omitempty"`
	CanBeDeleted          bool                  `json:"canBeDeleted,omitempty"`
	CannotBeDeletedReason string                `json:"cannotBeDeletedReason,omitempty"`
	QrCode                *PixKeyQrCodeResponse `json:"qrCode,omitempty"`
	Errors                []ErrorResponse       `json:"errors,omitempty"`
	DateCreated           DateTime              `json:"dateCreated,omitempty"`
}

type PixTransactionResponse struct {
	ID                    string                      `json:"id,omitempty"`
	EndToEndIdentifier    string                      `json:"endToEndIdentifier,omitempty"`
	Finality              PixTransactionFinality      `json:"finality,omitempty"`
	Value                 float64                     `json:"value,omitempty"`
	ChangeValue           float64                     `json:"changeValue,omitempty"`
	RefundedValue         float64                     `json:"refundedValue,omitempty"`
	EffectiveDate         *Date                       `json:"effectiveDate,omitempty"`
	ScheduledDate         *Date                       `json:"scheduledDate,omitempty"`
	Status                PixTransactionStatus        `json:"status,omitempty"`
	Type                  PixTransactionType          `json:"type,omitempty"`
	OriginType            PixTransactionOriginType    `json:"originType,omitempty"`
	Description           string                      `json:"description,omitempty"`
	TransactionReceiptUrl string                      `json:"transactionReceiptUrl,omitempty"`
	RefusalReason         string                      `json:"refusalReason,omitempty"`
	CanBeCanceled         bool                        `json:"canBeCanceled,omitempty"`
	OriginalTransaction   string                      `json:"originalTransaction,omitempty"`
	ExternalAccount       *PixExternalAccountResponse `json:"externalAccount,omitempty"`
	QrCode                *PixQrCodeResponse          `json:"qrCode,omitempty"`
	Payment               string                      `json:"payment,omitempty"`
	Errors                []ErrorResponse             `json:"errors,omitempty"`
}

type PixCancelTransactionResponse struct {
	ID                    string                      `json:"id,omitempty"`
	EndToEndIdentifier    string                      `json:"endToEndIdentifier,omitempty"`
	Finality              PixTransactionFinality      `json:"finality,omitempty"`
	Value                 float64                     `json:"value,omitempty"`
	ChangeValue           float64                     `json:"changeValue,omitempty"`
	RefundedValue         float64                     `json:"refundedValue,omitempty"`
	EffectiveDate         *Date                       `json:"effectiveDate,omitempty"`
	ScheduledDate         *Date                       `json:"scheduledDate,omitempty"`
	Status                PixTransactionStatus        `json:"status,omitempty"`
	Type                  PixTransactionType          `json:"type,omitempty"`
	OriginType            PixTransactionOriginType    `json:"originType,omitempty"`
	Description           string                      `json:"description,omitempty"`
	TransactionReceiptUrl string                      `json:"transactionReceiptUrl,omitempty"`
	RefusalReason         string                      `json:"refusalReason,omitempty"`
	CanBeCanceled         bool                        `json:"canBeCanceled,omitempty"`
	OriginalTransaction   string                      `json:"originalTransaction,omitempty"`
	ExternalAccount       *PixExternalAccountResponse `json:"externalAccount,omitempty"`
	QrCode                string                      `json:"qrCode,omitempty"`
	Payment               string                      `json:"payment,omitempty"`
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
	Ispb        int             `json:"ispb,omitempty"`
	IspbName    string          `json:"ispbName,omitempty"`
	Name        string          `json:"name,omitempty"`
	TradingName string          `json:"tradingName,omitempty"`
	CpfCnpj     string          `json:"cpfCnpj,omitempty"`
	PersonType  string          `json:"personType,omitempty"`
	Agency      string          `json:"agency,omitempty"`
	Account     string          `json:"account,omitempty"`
	AccountType BankAccountType `json:"accountType,omitempty"`
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
	ID                     string          `json:"id,omitempty"`
	EncodedImage           string          `json:"encodedImage,omitempty"`
	Payload                string          `json:"payload,omitempty"`
	AllowsMultiplePayments bool            `json:"allowsMultiplePayments,omitempty"`
	ExpirationDate         Date            `json:"expirationDate,omitempty"`
	Errors                 []ErrorResponse `json:"errors,omitempty"`
}

type pix struct {
	env         Env
	accessToken string
}

type Pix interface {
	PayQrCode(ctx context.Context, body PayPixQrCodeRequest) (*PixTransactionResponse, Error)
	DecodeQrCode(ctx context.Context, body PixQrCodeRequest) (*DecodePixQrCodeResponse, Error)
	CancelTransactionByID(ctx context.Context, pixTransactionID string) (*PixCancelTransactionResponse, Error)
	CreateKey(ctx context.Context) (*PixKeyResponse, Error)
	CreateStaticKey(ctx context.Context, body CreatePixKeyStaticRequest) (*QrCodeResponse, Error)
	DeleteKeyByID(ctx context.Context, pixKeyID string) (*DeleteResponse, Error)
	GetKeyByID(ctx context.Context, pixKeyID string) (*PixKeyResponse, Error)
	GetTransactionByID(ctx context.Context, pixTransactionID string) (*PixTransactionResponse, Error)
	GetAllTransactions(ctx context.Context) (*Pageable[PixTransactionResponse], Error)
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
	if err := p.validatePayQrCodeBodyRequest(body); err != nil {
		return nil, NewError(ERROR_VALIDATION, err)
	}
	req := NewRequest[PixTransactionResponse](ctx, p.env, p.accessToken)
	return req.make(http.MethodPost, "/v3/pix/qrCodes/pay", body)
}

func (p pix) DecodeQrCode(ctx context.Context, body PixQrCodeRequest) (*DecodePixQrCodeResponse, Error) {
	if err := Validate().Struct(body); err != nil {
		return nil, NewError(ERROR_VALIDATION, err)
	}
	req := NewRequest[DecodePixQrCodeResponse](ctx, p.env, p.accessToken)
	return req.make(http.MethodPost, "/v3/pix/qrCodes/decode", body)
}

func (p pix) CancelTransactionByID(ctx context.Context, pixTransactionID string) (*PixCancelTransactionResponse, Error) {
	req := NewRequest[PixCancelTransactionResponse](ctx, p.env, p.accessToken)
	return req.make(http.MethodPost, fmt.Sprintf("/v3/pix/transactions/%s/cancel", pixTransactionID), nil)
}

func (p pix) CreateKey(ctx context.Context) (*PixKeyResponse, Error) {
	req := NewRequest[PixKeyResponse](ctx, p.env, p.accessToken)
	return req.make(http.MethodPost, "/v3/pix/addressKeys", createPixKeyRequest{Type: PIX_KEY_TYPE_EVP})
}

func (p pix) CreateStaticKey(ctx context.Context, body CreatePixKeyStaticRequest) (*QrCodeResponse, Error) {
	if err := p.validateCreateStaticKeyBodyRequest(body); err != nil {
		return nil, NewError(ERROR_VALIDATION, err)
	} else if !body.Format.IsEnumValid() {
		body.Format = QR_CODE_FORMAT_ALL
	}
	req := NewRequest[QrCodeResponse](ctx, p.env, p.accessToken)
	return req.make(http.MethodPost, "/v3/pix/qrCodes/static", body)
}

func (p pix) DeleteKeyByID(ctx context.Context, pixKeyID string) (*DeleteResponse, Error) {
	req := NewRequest[DeleteResponse](ctx, p.env, p.accessToken)
	return req.make(http.MethodDelete, fmt.Sprintf("/v3/pix/addressKeys/%s", pixKeyID), nil)
}

func (p pix) GetTransactionByID(ctx context.Context, pixTransactionID string) (*PixTransactionResponse, Error) {
	req := NewRequest[PixTransactionResponse](ctx, p.env, p.accessToken)
	urlValues := url.Values{"id": []string{pixTransactionID}}
	return req.make(http.MethodGet, "/v3/pix/transactions?"+urlValues.Encode(), nil)
}

func (p pix) GetKeyByID(ctx context.Context, pixKeyID string) (*PixKeyResponse, Error) {
	req := NewRequest[PixKeyResponse](ctx, p.env, p.accessToken)
	return req.make(http.MethodGet, fmt.Sprintf("/v3/pix/addressKeys/%s", pixKeyID), nil)
}

func (p pix) GetAllTransactions(ctx context.Context) (*Pageable[PixTransactionResponse], Error) {
	req := NewRequest[Pageable[PixTransactionResponse]](ctx, p.env, p.accessToken)
	return req.make(http.MethodGet, "/v3/pix/transactions", nil)
}

func (p pix) GetAllKeys(ctx context.Context, filter GetAllPixKeysRequest) (*Pageable[PixKeyResponse], Error) {
	req := NewRequest[Pageable[PixKeyResponse]](ctx, p.env, p.accessToken)
	return req.make(http.MethodGet, "/v3/pix/addressKeys", filter)
}

func (p pix) validatePayQrCodeBodyRequest(body PayPixQrCodeRequest) error {
	if err := Validate().Struct(body); err != nil {
		return err
	} else if !body.ScheduleDate.IsZero() && time.Now().After(body.ScheduleDate.Time()) {
		return berrors.New("invalid scheduleDate")
	}
	return nil
}

func (p pix) validateCreateStaticKeyBodyRequest(body CreatePixKeyStaticRequest) error {
	if err := Validate().Struct(body); err != nil {
		return err
	} else if !body.ExpirationDate.IsZero() && time.Now().After(body.ExpirationDate.Time()) {
		return berrors.New("invalid expirationDate")
	}
	return nil
}
