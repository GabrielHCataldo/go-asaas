package asaas

import (
	"context"
	"fmt"
	"net/http"
	"os"
)

type CreateSubaccountRequest struct {
	Name          string                           `json:"name,omitempty" validate:"required"`
	Email         string                           `json:"email,omitempty" validate:"required,email"`
	LoginEmail    string                           `json:"loginEmail,omitempty" validate:"omitempty,email"`
	CpfCnpj       string                           `json:"cpfCnpj,omitempty" validate:"required,document"`
	BirthDate     *Date                            `json:"birthDate,omitempty" validate:"omitempty,before_now"`
	CompanyType   CompanyType                      `json:"companyType,omitempty" validate:"omitempty,enum"`
	Phone         string                           `json:"phone,omitempty" validate:"omitempty,phone"`
	MobilePhone   string                           `json:"mobilePhone,omitempty" validate:"required,phone"`
	Site          string                           `json:"site,omitempty" validate:"omitempty,url"`
	Address       string                           `json:"address,omitempty" validate:"required"`
	AddressNumber string                           `json:"addressNumber,omitempty" validate:"required"`
	Complement    string                           `json:"complement,omitempty"`
	Province      string                           `json:"province,omitempty"`
	PostalCode    string                           `json:"postalCode,omitempty" validate:"required,postal_code"`
	Webhooks      []CreateSubaccountWebhookRequest `json:"webhooks,omitempty"`
}

type GetAllSubaccountsRequest struct {
	CpfCnpj  string `json:"cpfCnpj,omitempty"`
	Email    string `json:"email,omitempty"`
	Name     string `json:"name,omitempty"`
	WalletId string `json:"walletId,omitempty"`
	Offset   int    `json:"offset,omitempty"`
	Limit    int    `json:"limit,omitempty"`
}

type CreateSubaccountWebhookRequest struct {
	Type        TypeOfWebhook `json:"type,omitempty" validate:"omitempty,enum"`
	Url         string        `json:"url,omitempty" validate:"required,url"`
	Email       string        `json:"email,omitempty" validate:"required,email"`
	ApiVersion  string        `json:"apiVersion,omitempty" validate:"required,numeric,max=4"`
	Enabled     bool          `json:"enabled"`
	Interrupted bool          `json:"interrupted"`
	AuthToken   string        `json:"authToken,omitempty"`
}

type SubaccountSendDocumentRequest struct {
	DocumentFile *os.File               `json:"documentFile,omitempty" validate:"required"`
	Type         SubaccountDocumentType `json:"type,omitempty" validate:"required,enum"`
}

type SubaccountUpdateDocumentSentRequest struct {
	DocumentFile *os.File `json:"documentFile,omitempty" validate:"required"`
}

type SubaccountResponse struct {
	Id            string                  `json:"id,omitempty"`
	Name          string                  `json:"name,omitempty"`
	PersonType    PersonType              `json:"personType,omitempty"`
	Email         string                  `json:"email,omitempty"`
	LoginEmail    string                  `json:"loginEmail,omitempty"`
	CpfCnpj       string                  `json:"cpfCnpj,omitempty"`
	BirthDate     *Date                   `json:"birthDate,omitempty"`
	CompanyType   CompanyType             `json:"companyType,omitempty"`
	Phone         string                  `json:"phone,omitempty"`
	MobilePhone   string                  `json:"mobilePhone,omitempty"`
	Site          string                  `json:"site,omitempty"`
	Address       string                  `json:"address,omitempty"`
	AddressNumber string                  `json:"addressNumber,omitempty"`
	Complement    string                  `json:"complement,omitempty"`
	Province      string                  `json:"province,omitempty"`
	PostalCode    string                  `json:"postalCode,omitempty"`
	City          int                     `json:"city,omitempty"`
	Country       string                  `json:"country,omitempty"`
	ApiKey        string                  `json:"apiKey,omitempty"`
	WalletId      string                  `json:"walletId,omitempty"`
	AccountNumber AccountBankInfoResponse `json:"accountNumber,omitempty"`
	Errors        []ErrorResponse         `json:"errors,omitempty"`
}

type SubaccountDocumentSentResponse struct {
	Id     string                   `json:"id,omitempty"`
	Status SubaccountDocumentStatus `json:"status,omitempty"`
	Errors []ErrorResponse          `json:"errors,omitempty"`
}

type SubaccountDocumentsResponse struct {
	RejectReasons string                       `json:"rejectReasons,omitempty"`
	Data          []SubaccountDocumentResponse `json:"data,omitempty"`
	Errors        []ErrorResponse              `json:"errors,omitempty"`
}

type SubaccountDocumentResponse struct {
	Id          string                      `json:"id,omitempty"`
	Status      SubaccountDocumentStatus    `json:"status,omitempty"`
	Type        SubaccountDocumentType      `json:"type,omitempty"`
	Title       string                      `json:"title,omitempty"`
	Description string                      `json:"description,omitempty"`
	Responsible DocumentResponsibleResponse `json:"responsible,omitempty"`
	Documents   []any                       `json:"documents,omitempty"`
}

type DocumentResponsibleResponse struct {
	Name string `json:"name,omitempty"`
	Type string `json:"type,omitempty"`
}

type subaccount struct {
	env         Env
	accessToken string
}

type Subaccount interface {
	Create(ctx context.Context, body CreateSubaccountRequest) (*SubaccountResponse, Error)
	SendDocument(ctx context.Context, accountId string, body SubaccountSendDocumentRequest) (
		*SubaccountDocumentSentResponse, Error)
	UpdateDocumentSentByID(ctx context.Context, documentSentId string, body SubaccountUpdateDocumentSentRequest) (
		*SubaccountDocumentSentResponse, Error)
	DeleteDocumentSentByID(ctx context.Context, documentSentId string) (*DeleteResponse, Error)
	GetByID(ctx context.Context, subaccountId string) (*SubaccountResponse, Error)
	GetDocumentSentByID(ctx context.Context, documentSentId string) (*SubaccountDocumentSentResponse, Error)
	GetAll(ctx context.Context, filter GetAllSubaccountsRequest) (*Pageable[SubaccountResponse], Error)
	GetDocuments(ctx context.Context) (*SubaccountDocumentsResponse, Error)
}

func NewSubaccount(env Env, accessToken string) Subaccount {
	logWarning("Subaccount service running on", env.String())
	return subaccount{
		env:         env,
		accessToken: accessToken,
	}
}

func (s subaccount) Create(ctx context.Context, body CreateSubaccountRequest) (*SubaccountResponse, Error) {
	if err := Validate().Struct(body); err != nil {
		return nil, NewError(ErrorTypeValidation, err)
	}
	req := NewRequest[SubaccountResponse](ctx, s.env, s.accessToken)
	return req.make(http.MethodPost, "/v3/accounts", body)
}

func (s subaccount) SendDocument(ctx context.Context, accountId string, body SubaccountSendDocumentRequest) (
	*SubaccountDocumentSentResponse, Error) {
	if err := Validate().Struct(body); err != nil {
		return nil, NewError(ErrorTypeValidation, err)
	}
	req := NewRequest[SubaccountDocumentSentResponse](ctx, s.env, s.accessToken)
	return req.makeMultipartForm(http.MethodPost, fmt.Sprintf("/v3/myAccount/documents/%s", accountId), body)
}

func (s subaccount) UpdateDocumentSentByID(
	ctx context.Context,
	documentSentId string,
	body SubaccountUpdateDocumentSentRequest,
) (*SubaccountDocumentSentResponse, Error) {
	if err := Validate().Struct(body); err != nil {
		return nil, NewError(ErrorTypeValidation, err)
	}
	req := NewRequest[SubaccountDocumentSentResponse](ctx, s.env, s.accessToken)
	return req.makeMultipartForm(http.MethodPut, fmt.Sprintf("/v3/myAccount/documents/files/%s", documentSentId), body)
}

func (s subaccount) DeleteDocumentSentByID(ctx context.Context, documentSentId string) (*DeleteResponse, Error) {
	req := NewRequest[DeleteResponse](ctx, s.env, s.accessToken)
	return req.make(http.MethodDelete, fmt.Sprintf("/v3/myAccount/documents/files/%s", documentSentId), nil)
}

func (s subaccount) GetByID(ctx context.Context, subaccountId string) (*SubaccountResponse, Error) {
	req := NewRequest[SubaccountResponse](ctx, s.env, s.accessToken)
	return req.make(http.MethodGet, fmt.Sprintf("/v3/accounts/%s", subaccountId), nil)
}

func (s subaccount) GetDocumentSentByID(ctx context.Context, documentSentId string) (*SubaccountDocumentSentResponse,
	Error) {
	req := NewRequest[SubaccountDocumentSentResponse](ctx, s.env, s.accessToken)
	return req.make(http.MethodGet, fmt.Sprintf("/v3/myAccount/documents/files/%s", documentSentId), nil)
}

func (s subaccount) GetDocuments(ctx context.Context) (*SubaccountDocumentsResponse, Error) {
	req := NewRequest[SubaccountDocumentsResponse](ctx, s.env, s.accessToken)
	return req.make(http.MethodGet, "/v3/myAccount/documents", nil)
}

func (s subaccount) GetAll(ctx context.Context, filter GetAllSubaccountsRequest) (*Pageable[SubaccountResponse], Error) {
	req := NewRequest[Pageable[SubaccountResponse]](ctx, s.env, s.accessToken)
	return req.make(http.MethodGet, "/v3/accounts", filter)
}
