package asaas

import (
	"context"
	"github.com/mvrilo/go-cpf"
	"testing"
	"time"
)

func TestCustomerCreate(t *testing.T) {
	accessToken := getEnvValue(EnvAccessToken)
	assertFatalStringBlank(t, accessToken)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nCustomer := NewCustomer(EnvSandbox, accessToken)
	resp, err := nCustomer.Create(ctx, CreateCustomerRequest{
		Name:                 "Unit test go",
		CpfCnpj:              cpf.Generate(),
		Email:                "",
		Phone:                "",
		MobilePhone:          "",
		Address:              "",
		AddressNumber:        "",
		Complement:           "",
		Province:             "",
		PostalCode:           "",
		ExternalReference:    "",
		NotificationDisabled: false,
		AdditionalEmails:     "",
		MunicipalInscription: "",
		StateInscription:     "",
		Observations:         "",
		GroupName:            "",
		Company:              "",
	})
	assertResponseSuccess(t, resp, err)
}

func TestCustomerUpdateById(t *testing.T) {
	accessToken := getEnvValue(EnvAccessToken)
	assertFatalStringBlank(t, accessToken)
	initCustomer()
	customerId := getEnvValue(EnvCustomerId)
	assertFatalStringBlank(t, customerId)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nCustomer := NewCustomer(EnvSandbox, accessToken)
	resp, err := nCustomer.UpdateById(ctx, customerId, UpdateCustomerRequest{
		Name:                 "Unit test go updated",
		CpfCnpj:              nil,
		Email:                nil,
		Phone:                nil,
		MobilePhone:          nil,
		Address:              nil,
		AddressNumber:        nil,
		Complement:           nil,
		Province:             nil,
		PostalCode:           nil,
		ExternalReference:    nil,
		NotificationDisabled: nil,
		AdditionalEmails:     nil,
		MunicipalInscription: nil,
		StateInscription:     nil,
		Observations:         nil,
		GroupName:            nil,
		Company:              nil,
	})
	assertResponseSuccess(t, resp, err)
}

func TestCustomerDeleteById(t *testing.T) {
	accessToken := getEnvValue(EnvAccessToken)
	assertFatalStringBlank(t, accessToken)
	initCustomer()
	customerId := getEnvValue(EnvCustomerId)
	assertFatalStringBlank(t, customerId)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nCustomer := NewCustomer(EnvSandbox, accessToken)
	resp, err := nCustomer.DeleteById(ctx, customerId)
	assertResponseSuccess(t, resp, err)
}

func TestCustomerRestoreById(t *testing.T) {
	accessToken := getEnvValue(EnvAccessToken)
	assertFatalStringBlank(t, accessToken)
	initCustomerDeleted()
	customerId := getEnvValue(EnvChargeDeletedId)
	assertFatalStringBlank(t, customerId)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nCustomer := NewCustomer(EnvSandbox, accessToken)
	resp, err := nCustomer.RestoreById(ctx, customerId)
	assertResponseSuccess(t, resp, err)
}

func TestCustomerGetById(t *testing.T) {
	accessToken := getEnvValue(EnvAccessToken)
	assertFatalStringBlank(t, accessToken)
	initCustomer()
	customerId := getEnvValue(EnvChargeDeletedId)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nCustomer := NewCustomer(EnvSandbox, accessToken)
	resp, err := nCustomer.GetById(ctx, customerId)
	assertResponseSuccess(t, resp, err)
}

func TestCustomerGetAll(t *testing.T) {
	accessToken := getEnvValue(EnvAccessToken)
	assertFatalStringBlank(t, accessToken)
	initCustomer()
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nCustomer := NewCustomer(EnvSandbox, accessToken)
	resp, err := nCustomer.GetAll(ctx, GetAllCustomersRequest{
		Name:              "",
		Email:             "",
		CpfCnpj:           "",
		GroupName:         "",
		ExternalReference: "",
		Offset:            0,
		Limit:             10,
	})
	assertResponseSuccess(t, resp, err)
}
