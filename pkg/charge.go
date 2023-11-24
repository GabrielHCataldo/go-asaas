package asaas

import (
	"context"
	"encoding/json"
	"errors"
	"go-asaas/internal/util"
	"go-asaas/internal/validator"
	"io"
	"net/http"
	"time"
)

type charge struct {
	assasEnv   Env
	accessCode string
}

type Charge interface {
	Create(ctx context.Context, body CreateChargeRequest) (*CreateChargeResponse, Error)
}

func NewCharge(assasEnv Env, accessCode string) Charge {
	return charge{
		assasEnv:   assasEnv,
		accessCode: accessCode,
	}
}

func (c charge) Create(ctx context.Context, body CreateChargeRequest) (*CreateChargeResponse, Error) {
	if err := c.validateCreateBodyRequest(body); err != nil {
		return nil, NewByErrorType(ERROR_VALIDATION, err)
	} else if ctx.Err() != nil {
		return nil, NewByError(ctx.Err())
	}
	c.prepareCreateBodyRequest(&body)
	req, err := CreateHttpRequest(ctx, c.assasEnv, c.accessCode, http.MethodPost, "/v3/payments", body)
	if err != nil {
		return nil, NewByError(err)
	}
	res, err := MakeHttpRequest(req)
	defer CloseBody(res.Body)
	if err != nil {
		return nil, NewByError(err)
	}
	if res.StatusCode == http.StatusOK || res.StatusCode == http.StatusBadRequest {
		respBody, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, NewByError(err)
		}
		chargeResponse := &CreateChargeResponse{}
		err = json.Unmarshal(respBody, chargeResponse)
		if err != nil {
			return nil, NewByError(err)
		}
		return chargeResponse, nil
	}
	return nil, NewError(ERROR_UNEXPECTED, "response status code not expected: ", res.StatusCode)
}

func (c charge) validateCreateBodyRequest(body CreateChargeRequest) error {
	if err := validator.Validate().Struct(body); err != nil {
		return err
	} else if time.Now().After(body.DueDate) {
		return errors.New("invalid due date")
	}
	switch body.BillingType {
	case CREDIT_CARD:
		if util.IsBlank(&body.CreditCardToken) && (body.CreditCard == nil || body.CreditCardHolderInfo == nil) {
			return errors.New("to charge by credit card, enter the credit card or credit card token")
		} else if body.CreditCard != nil {
			exp, err := time.Parse("1999-01-01", body.CreditCard.ExpiryYear+"-"+body.CreditCard.ExpiryMonth+"-01")
			if err != nil || time.Now().After(exp) {
				return errors.New("expired card")
			}
		}
		break
	}
	return nil
}

func (c charge) prepareCreateBodyRequest(body *CreateChargeRequest) {
	switch body.BillingType {
	case CREDIT_CARD:
		if body.Fine != nil {
			body.Fine.DueDateLimitDays = 0
		}
		break
	default:
		body.CreditCard = nil
		body.CreditCardHolderInfo = nil
		body.CreditCardToken = ""
	}
}
