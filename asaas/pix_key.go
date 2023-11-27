package asaas

import (
	"context"
	berrors "errors"
	"fmt"
	"net/http"
	"time"
)

type createPixKeyRequest struct {
	Type PixKeyType `json:"type,omitempty" validate:"required,enum"`
}

type CreatePixKeyStaticRequest struct {
	AddressKey             string       `json:"addressKey,omitempty"`
	Description            string       `json:"description,omitempty"`
	Value                  float64      `json:"value,omitempty" validate:"omitempty,gt=0"`
	Format                 QRCodeFormat `json:"format,omitempty" validate:"omitempty,enum"`
	ExpirationDate         DateTime     `json:"expirationDate,omitempty"`
	ExpirationSeconds      int          `json:"expirationSeconds,omitempty" validate:"omitempty,gt=0"`
	AllowsMultiplePayments bool         `json:"allowsMultiplePayments"`
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
	QrCode                *PixKeyQRCodeResponse `json:"qrCode,omitempty"`
	Errors                []ErrorResponse       `json:"errors,omitempty"`
	DateCreated           DateTime              `json:"dateCreated,omitempty"`
}

type PixKeyQRCodeResponse struct {
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

type pixKey struct {
	env         Env
	accessToken string
}

type PixKey interface {
	Create(ctx context.Context) (*PixKeyResponse, Error)
	CreateStatic(ctx context.Context, body CreatePixKeyStaticRequest) (*QrCodeResponse, Error)
	DeleteByID(ctx context.Context, pixKeyID string) (*DeleteResponse, Error)
	GetByID(ctx context.Context, pixKeyID string) (*PixKeyResponse, Error)
	GetAll(ctx context.Context, filter GetAllPixKeysRequest) (*Pageable[PixKeyResponse], Error)
}

func NewPixKey(env Env, accessToken string) PixKey {
	logWarning("PixKey service running on", env.String())
	return pixKey{
		env:         env,
		accessToken: accessToken,
	}
}

func (p pixKey) Create(ctx context.Context) (*PixKeyResponse, Error) {
	req := NewRequest[PixKeyResponse](ctx, p.env, p.accessToken)
	return req.make(http.MethodPost, "/v3/pix/addressKeys", createPixKeyRequest{Type: PIX_KEY_TYPE_EVP})
}

func (p pixKey) CreateStatic(ctx context.Context, body CreatePixKeyStaticRequest) (*QrCodeResponse, Error) {
	if err := p.validateCreateStaticBodyRequest(body); err != nil {
		return nil, NewError(ERROR_VALIDATION, err)
	} else if !body.Format.IsEnumValid() {
		body.Format = QR_CODE_FORMAT_ALL
	}
	req := NewRequest[QrCodeResponse](ctx, p.env, p.accessToken)
	return req.make(http.MethodPost, "/v3/pix/qrCodes/static", body)
}

func (p pixKey) DeleteByID(ctx context.Context, pixKeyID string) (*DeleteResponse, Error) {
	req := NewRequest[DeleteResponse](ctx, p.env, p.accessToken)
	return req.make(http.MethodDelete, fmt.Sprintf("/v3/pix/addressKeys/%s", pixKeyID), nil)
}

func (p pixKey) GetByID(ctx context.Context, pixKeyID string) (*PixKeyResponse, Error) {
	req := NewRequest[PixKeyResponse](ctx, p.env, p.accessToken)
	return req.make(http.MethodGet, fmt.Sprintf("/v3/pix/addressKeys/%s", pixKeyID), nil)
}

func (p pixKey) GetAll(ctx context.Context, filter GetAllPixKeysRequest) (*Pageable[PixKeyResponse], Error) {
	req := NewRequest[Pageable[PixKeyResponse]](ctx, p.env, p.accessToken)
	return req.make(http.MethodGet, "/v3/pix/addressKeys", filter)
}

func (p pixKey) validateCreateStaticBodyRequest(body CreatePixKeyStaticRequest) error {
	if err := Validate().Struct(body); err != nil {
		return err
	} else if !body.ExpirationDate.IsZero() && time.Now().After(body.ExpirationDate.Time()) {
		return berrors.New("invalid expirationDate")
	}
	return nil
}
