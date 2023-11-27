package asaas

import (
	"context"
	"net/http"
)

type CreditCardTokenizeRequest struct {
	Customer             string                      `json:"customer,omitempty" validate:"required"`
	CreditCard           CreditCardRequest           `json:"creditCard,omitempty" validate:"required"`
	CreditCardHolderInfo CreditCardHolderInfoRequest `json:"creditCardHolderInfo,omitempty" validate:"required"`
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
