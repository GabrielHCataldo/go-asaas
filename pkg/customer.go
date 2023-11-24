package asaas

import (
	"context"
	"encoding/json"
	"go-asaas/internal/validator"
	"io"
	"net/http"
)

type customer struct {
	asaasEnv   Env
	accessCode string
}

type Customer interface {
	Create(ctx context.Context, body CreateCustomerRequest) (*CreateCustomerResponse, Error)
}

func NewCustomer(assasEnv Env, accessCode string) Customer {
	return customer{
		asaasEnv:   assasEnv,
		accessCode: accessCode,
	}
}

func (c customer) Create(ctx context.Context, body CreateCustomerRequest) (
	*CreateCustomerResponse, Error) {
	if err := validator.Validate().Struct(body); err != nil {
		return nil, NewByErrorType(ERROR_VALIDATION, err)
	} else if ctx.Err() != nil {
		return nil, NewByError(ctx.Err())
	}
	req, err := CreateHttpRequest(ctx, c.asaasEnv, c.accessCode, http.MethodPost, "/v3/customers", body)
	if err != nil {
		return nil, NewByError(err)
	}
	res, err := MakeHttpRequest(req)
	if err != nil {
		return nil, NewByError(err)
	}
	defer CloseBody(res.Body)
	if res.StatusCode == http.StatusOK || res.StatusCode == http.StatusBadRequest {
		respBodyBytes, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, NewByError(err)
		}
		createCustomerResponse := &CreateCustomerResponse{}
		err = json.Unmarshal(respBodyBytes, createCustomerResponse)
		if err != nil {
			return nil, NewByError(err)
		}
		return createCustomerResponse, nil
	}
	return nil, NewError(ERROR_UNEXPECTED, "response status code not expected: ", res.StatusCode)
}
