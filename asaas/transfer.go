package asaas

import (
	"context"
	"fmt"
	"net/http"
)

type TransferToBankRequest struct {
	Value             float64               `json:"value,omitempty" validate:"required,gt=0"`
	BankAccount       BackAccountRequest    `json:"bankAccount,omitempty" validate:"required"`
	OperationType     TransferOperationType `json:"operationType,omitempty" validate:"omitempty,enum"`
	PixAddressKey     string                `json:"pixAddressKey,omitempty"`
	PixAddressKeyType PixKeyType            `json:"pixAddressKeyType,omitempty" validate:"omitempty,enum"`
	Description       string                `json:"description,omitempty"`
	ScheduleDate      Date                  `json:"scheduleDate,omitempty" validate:"omitempty,after_now"`
}

type TransferToAssasRequest struct {
	Value    float64 `json:"value,omitempty" validate:"required,gt=0"`
	WalletID string  `json:"walletId,omitempty" validate:"required"`
}

type BackAccountRequest struct {
	Bank            BankRequest     `json:"bank,omitempty" validate:"required"`
	AccountName     string          `json:"accountName,omitempty"`
	OwnerName       string          `json:"ownerName,omitempty" validate:"required"`
	OwnerBirthDate  Date            `json:"ownerBirthDate,omitempty" validate:"omitempty,before_now"`
	CpfCnpj         string          `json:"cpfCnpj,omitempty" validate:"required,document"`
	Agency          string          `json:"agency,omitempty" validate:"required,numeric,max=5"`
	Account         string          `json:"account,omitempty" validate:"required,numeric,max=12"`
	AccountDigit    string          `json:"accountDigit,omitempty" validate:"required,numeric,max=2"`
	BankAccountType BankAccountType `json:"bankAccountType,omitempty" validate:"required,enum"`
	Ispb            string          `json:"ispb,omitempty"`
}

type BankRequest struct {
	Code string `json:"code,omitempty" validate:"required"`
}

type GetAllTransfersRequest struct {
	DateCreatedGe  Date `json:"dateCreated[ge],omitempty"`
	DateCreatedLe  Date `json:"dateCreated[le],omitempty"`
	TransferDateGe Date `json:"transferDate[ge],omitempty"`
	TransferDateLe Date `json:"transferDate[le],omitempty"`
	Type           Date `json:"type,omitempty"`
}

type TransferResponse struct {
	ID                    string                `json:"id,omitempty"`
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
	WalletID              string                `json:"walletId,omitempty"`
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
	TransferToBank(ctx context.Context, body TransferToBankRequest) (*TransferResponse, Error)
	TransferToAsaas(ctx context.Context, body TransferToAssasRequest) (*TransferResponse, Error)
	CancelByID(ctx context.Context, transferID string) (*TransferResponse, Error)
	GetByID(ctx context.Context, transferID string) (*TransferResponse, Error)
	GetAll(ctx context.Context, filter GetAllTransfersRequest) (*Pageable[TransferResponse], Error)
}

func NewTransfer(env Env, accessToken string) Transfer {
	logWarning("Transfer service running on", env.String())
	return transfer{
		env:         env,
		accessToken: accessToken,
	}
}

func (t transfer) TransferToBank(ctx context.Context, body TransferToBankRequest) (*TransferResponse, Error) {
	if err := Validate().Struct(body); err != nil {
		return nil, NewError(ErrorTypeValidation, err)
	}
	req := NewRequest[TransferResponse](ctx, t.env, t.accessToken)
	return req.make(http.MethodPost, "/v3/transfers", body)
}

func (t transfer) TransferToAsaas(ctx context.Context, body TransferToAssasRequest) (*TransferResponse,
	Error) {
	if err := Validate().Struct(body); err != nil {
		return nil, NewError(ErrorTypeValidation, err)
	}
	req := NewRequest[TransferResponse](ctx, t.env, t.accessToken)
	return req.make(http.MethodPost, "/v3/transfers", body)
}

func (t transfer) CancelByID(ctx context.Context, transferID string) (*TransferResponse, Error) {
	req := NewRequest[TransferResponse](ctx, t.env, t.accessToken)
	return req.make(http.MethodDelete, fmt.Sprintf("/v3/transfers/%s/cancel", transferID), nil)
}

func (t transfer) GetByID(ctx context.Context, transferID string) (*TransferResponse, Error) {
	req := NewRequest[TransferResponse](ctx, t.env, t.accessToken)
	return req.make(http.MethodGet, fmt.Sprintf("/v3/transfers/%s", transferID), nil)
}

func (t transfer) GetAll(ctx context.Context, filter GetAllTransfersRequest) (*Pageable[TransferResponse], Error) {
	req := NewRequest[Pageable[TransferResponse]](ctx, t.env, t.accessToken)
	return req.make(http.MethodGet, "/v3/transfers", filter)
}
