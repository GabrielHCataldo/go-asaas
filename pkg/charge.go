package asaas

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/GabrielHCataldo/go-asaas/internal/util"
	"github.com/GabrielHCataldo/go-asaas/internal/validator"
	"io"
	"log"
	"net/http"
	"time"
)

type charge struct {
	env        Env
	accessCode string
}

type Charge interface {
	Create(ctx context.Context, body CreateChargeRequest) (*CreateChargeResponse, Error)
}

func NewCharge(env Env, accessCode string) Charge {
	return charge{
		env:        env,
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
	bodyReq, _ := json.Marshal(body)
	log.Println("request body:", string(bodyReq))
	req, err := CreateHttpRequest(ctx, c.env, c.accessCode, http.MethodPost, "/v3/payments", body)
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
		log.Println("response body:", string(respBody))
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
	} else {
		dueDate := time.Date(body.DueDate.Time().Year(), body.DueDate.Month(), body.DueDate.Day(),
			23, 59, 0, 0, body.DueDate.Location())
		if time.Now().UTC().After(dueDate.UTC()) {
			return errors.New("invalid due date")
		}
	}
	switch body.BillingType {
	case CREDIT_CARD:
		if util.IsBlank(&body.CreditCardToken) && (body.CreditCard == nil || body.CreditCardHolderInfo == nil) {
			return errors.New("to charge by credit card, enter the credit card or credit card token")
		} else if body.CreditCard != nil {
			exp, err := time.Parse("2006-01-02", body.CreditCard.ExpiryYear+"-"+body.CreditCard.ExpiryMonth+"-01")
			if err != nil || time.Now().UTC().After(exp.UTC()) {
				return errors.New("expired card")
			}
		}
		break
	}
	return nil
}

func (c charge) prepareCreateBodyRequest(body *CreateChargeRequest) {
	body.DueDate = NewDate(body.DueDate.Year(), body.DueDate.Month(), body.DueDate.Day(),
		23, 59, 0, 0, body.DueDate.Location())
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
