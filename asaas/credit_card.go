package asaas

import (
	"context"
	"net/http"
)

type CreditCardRequest struct {
	HolderName  string `json:"holderName,omitempty"`
	Number      string `json:"number,omitempty"`
	ExpiryMonth string `json:"expiryMonth,omitempty"`
	ExpiryYear  string `json:"expiryYear,omitempty"`
	Ccv         string `json:"ccv,omitempty"`
}

type CreditCardHolderInfoRequest struct {
	Name              string `json:"name,omitempty"`
	CpfCnpj           string `json:"cpfCnpj,omitempty"`
	Email             string `json:"email,omitempty"`
	Phone             string `json:"phone,omitempty"`
	MobilePhone       string `json:"mobilePhone,omitempty"`
	PostalCode        string `json:"postalCode,omitempty"`
	AddressNumber     string `json:"addressNumber,omitempty"`
	AddressComplement string `json:"addressComplement,omitempty"`
}

type CreditCardTokenizeRequest struct {
	Customer             string                      `json:"customer,omitempty"`
	CreditCard           CreditCardRequest           `json:"creditCard,omitempty"`
	CreditCardHolderInfo CreditCardHolderInfoRequest `json:"creditCardHolderInfo,omitempty"`
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
	req := NewRequest[CreditCardTokenizeResponse](ctx, c.env, c.accessToken)
	return req.make(http.MethodPost, "/v3/creditCard/tokenize", body)
}
