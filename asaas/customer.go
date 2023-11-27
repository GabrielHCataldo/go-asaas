package asaas

import (
	"context"
	"fmt"
	"net/http"
)

type UpdateCustomerRequest struct {
	Name                 string `json:"name,omitempty" validate:"required,full_name"`
	CpfCnpj              string `json:"cpfCnpj,omitempty" validate:"required,document"`
	Email                string `json:"email,omitempty" validate:"omitempty,email"`
	Phone                string `json:"phone,omitempty" validate:"omitempty,phone"`
	MobilePhone          string `json:"mobilePhone,omitempty" validate:"omitempty,phone"`
	Address              string `json:"address,omitempty"`
	AddressNumber        string `json:"addressNumber,omitempty"`
	Complement           string `json:"complement,omitempty"`
	Province             string `json:"province,omitempty"`
	PostalCode           string `json:"postalCode,omitempty" validate:"omitempty,postal_code"`
	ExternalReference    string `json:"externalReference,omitempty"`
	NotificationDisabled bool   `json:"notificationDisabled,omitempty"`
	AdditionalEmails     string `json:"additionalEmails,omitempty"`
	MunicipalInscription string `json:"municipalInscription,omitempty"`
	StateInscription     string `json:"stateInscription,omitempty"`
	Observations         string `json:"observations,omitempty"`
	GroupName            string `json:"groupName,omitempty"`
	Company              string `json:"company,omitempty"`
}

type GetAllCustomersRequest struct {
	Name              string `json:"name,omitempty"`
	Email             string `json:"email,omitempty"`
	GroupName         string `json:"groupName,omitempty"`
	ExternalReference string `json:"externalReference,omitempty"`
	CpfCnpj           string `json:"cpfCnpj,omitempty"`
	Offset            int    `json:"offset,omitempty"`
	Limit             int    `json:"limit,omitempty"`
}

type CustomerResponse struct {
	ID                    string          `json:"id,omitempty"`
	Name                  string          `json:"name,omitempty"`
	Email                 string          `json:"email,omitempty"`
	Phone                 string          `json:"phone,omitempty"`
	MobilePhone           string          `json:"mobilePhone,omitempty"`
	Address               string          `json:"address,omitempty"`
	AddressNumber         string          `json:"addressNumber,omitempty"`
	Complement            string          `json:"complement,omitempty"`
	Province              string          `json:"province,omitempty"`
	PostalCode            string          `json:"postalCode,omitempty"`
	CpfCnpj               string          `json:"cpfCnpj,omitempty"`
	PersonType            string          `json:"personType,omitempty"`
	Deleted               bool            `json:"deleted,omitempty"`
	AdditionalEmails      string          `json:"additionalEmails,omitempty"`
	ExternalReference     string          `json:"externalReference,omitempty"`
	NotificationDisabled  bool            `json:"notificationDisabled,omitempty"`
	MunicipalInscription  string          `json:"municipalInscription,omitempty"`
	StateInscription      string          `json:"stateInscription,omitempty"`
	CanDelete             bool            `json:"canDelete,omitempty"`
	CannotBeDeletedReason string          `json:"cannotBeDeletedReason,omitempty"`
	CanEdit               bool            `json:"canEdit,omitempty"`
	CannotEditReason      string          `json:"cannotEditReason,omitempty"`
	ForeignCustomer       bool            `json:"foreignCustomer,omitempty"`
	City                  int             `json:"city,omitempty"`
	State                 string          `json:"state,omitempty"`
	Country               string          `json:"country,omitempty"`
	Observations          string          `json:"observations,omitempty"`
	Errors                []ErrorResponse `json:"errors,omitempty"`
	DateCreated           DateTime        `json:"dateCreated,omitempty"`
}

type customer struct {
	env         Env
	accessToken string
}

type Customer interface {
	Create(ctx context.Context, body UpdateCustomerRequest) (*CustomerResponse, Error)
	UpdateByID(ctx context.Context, customerID string, body UpdateCustomerRequest) (*CustomerResponse, Error)
	DeleteByID(ctx context.Context, customerID string) (*DeleteResponse, Error)
	RestoreByID(ctx context.Context, customerID string) (*CustomerResponse, Error)
	GetByID(ctx context.Context, customerID string) (*CustomerResponse, Error)
	GetAll(ctx context.Context, filter GetAllCustomersRequest) (*Pageable[CustomerResponse], Error)
}

func NewCustomer(env Env, accessToken string) Customer {
	logWarning("Customer service running on", env.String())
	return customer{
		env:         env,
		accessToken: accessToken,
	}
}

func (c customer) Create(ctx context.Context, body UpdateCustomerRequest) (*CustomerResponse, Error) {
	if err := Validate().Struct(body); err != nil {
		return nil, NewError(ERROR_VALIDATION, err)
	}
	req := NewRequest[CustomerResponse](ctx, c.env, c.accessToken)
	return req.make(http.MethodPost, "/v3/customers", body)
}

func (c customer) UpdateByID(ctx context.Context, customerID string, body UpdateCustomerRequest) (*CustomerResponse, Error) {
	if err := Validate().Struct(body); err != nil {
		return nil, NewError(ERROR_VALIDATION, err)
	}
	req := NewRequest[CustomerResponse](ctx, c.env, c.accessToken)
	return req.make(http.MethodPost, fmt.Sprintf("/v3/customers/%s", customerID), body)
}

func (c customer) DeleteByID(ctx context.Context, customerID string) (*DeleteResponse, Error) {
	req := NewRequest[DeleteResponse](ctx, c.env, c.accessToken)
	return req.make(http.MethodDelete, fmt.Sprintf("/v3/customers/%s", customerID), nil)
}

func (c customer) RestoreByID(ctx context.Context, customerID string) (*CustomerResponse, Error) {
	req := NewRequest[CustomerResponse](ctx, c.env, c.accessToken)
	return req.make(http.MethodPost, fmt.Sprintf("/v3/customers/%s", customerID), nil)
}

func (c customer) GetByID(ctx context.Context, customerID string) (*CustomerResponse, Error) {
	req := NewRequest[CustomerResponse](ctx, c.env, c.accessToken)
	return req.make(http.MethodGet, fmt.Sprintf("/v3/customers/%s", customerID), nil)
}

func (c customer) GetAll(ctx context.Context, filter GetAllCustomersRequest) (*Pageable[CustomerResponse], Error) {
	req := NewRequest[Pageable[CustomerResponse]](ctx, c.env, c.accessToken)
	return req.make(http.MethodGet, "/v3/customers", filter)
}
