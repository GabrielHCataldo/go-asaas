package asaas

import (
	"context"
	"net/http"
)

type CreditCardRequest struct {
	HolderName  string `json:"holderName,omitempty" validate:"required,full_name"`
	Number      string `json:"number,omitempty" validate:"required,numeric,min=10,max=19"`
	ExpiryMonth string `json:"expiryMonth,omitempty" validate:"required,numeric,len=2"`
	ExpiryYear  string `json:"expiryYear,omitempty" validate:"required,numeric,len=4"`
	CCV         string `json:"ccv,omitempty" validate:"required,numeric,min=3,max=4"`
}

type CreditCardHolderInfoRequest struct {
	Name              string `json:"name,omitempty" validate:"required,full_name"`
	CpfCnpj           string `json:"cpfCnpj,omitempty" validate:"required,document"`
	Email             string `json:"email,omitempty" validate:"required,email"`
	Phone             string `json:"phone,omitempty" validate:"required,phone"`
	MobilePhone       string `json:"mobilePhone,omitempty" validate:"omitempty,phone"`
	PostalCode        string `json:"postalCode,omitempty" validate:"required,postal_code"`
	AddressNumber     string `json:"addressNumber,omitempty" validate:"required,numeric"`
	AddressComplement string `json:"addressComplement,omitempty"`
}

type CreditCardTokenizeRequest struct {
	Customer             string                      `json:"customer,omitempty" validate:"required"`
	CreditCard           CreditCardRequest           `json:"creditCard,omitempty" validate:"required"`
	CreditCardHolderInfo CreditCardHolderInfoRequest `json:"creditCardHolderInfo,omitempty" validate:"required"`
}

type CreditCardResponse struct {
	CreditCardNumber string `json:"creditCardNumber,omitempty"`
	CreditCardBrand  string `json:"creditCardBrand,omitempty"`
	CreditCardToken  string `json:"creditCardToken,omitempty"`
}

type CreditCardTokenizeResponse struct {
	CreditCardNumber string          `json:"creditCardNumber,omitempty"`
	CreditCardBrand  string          `json:"creditCardBrand,omitempty"`
	CreditCardToken  string          `json:"creditCardToken,omitempty"`
	Errors           []ErrorResponse `json:"errors,omitempty"`
}

type creditCard struct {
	env         Env
	accessToken string
}

type CreditCard interface {
	Tokenize(ctx context.Context, body CreditCardTokenizeRequest) (*CreditCardTokenizeResponse, Error)
}

func NewCreditCard(env Env, accessToken string) CreditCard {
	logWarning("CreditCard service running on", env.String())
	return creditCard{
		env:         env,
		accessToken: accessToken,
	}
}

func (c creditCard) Tokenize(ctx context.Context, body CreditCardTokenizeRequest) (*CreditCardTokenizeResponse,
	Error) {
	if err := Validate().Struct(body); err != nil {
		return nil, NewError(ERROR_VALIDATION, err)
	}
	req := NewRequest[CreditCardTokenizeResponse](ctx, c.env, c.accessToken)
	return req.make(http.MethodPost, "/v3/creditCard/tokenize", body)
}
