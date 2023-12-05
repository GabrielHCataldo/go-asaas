package main

import (
	"context"
	"fmt"
	"github.com/GabrielHCataldo/go-asaas/asaas"
	"os"
)

var customerAsaas asaas.Customer

func main() {
	customerAsaas = asaas.NewCustomer(asaas.EnvSandbox, os.Getenv("ASAAS_ACCESS_TOKEN"))
	createCustomer()
	updateCustomerById()
	getCustomerById()
	getAllCustomers()
	deleteCustomerById()
}

func createCustomer() {
	resp, err := customerAsaas.Create(context.TODO(), asaas.CreateCustomerRequest{
		Name:                 "",
		CpfCnpj:              "",
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
	if err != nil {
		fmt.Println("error:", err)
	} else if resp.IsSuccess() {
		fmt.Println("success:", resp)
	} else {
		fmt.Println("failure:", resp.Errors)
	}
}

func updateCustomerById() {
	resp, err := customerAsaas.UpdateById(context.TODO(), "", asaas.UpdateCustomerRequest{
		Name:                 "",
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
	if err != nil {
		fmt.Println("error:", err)
	} else if resp.IsSuccess() {
		fmt.Println("success:", resp)
	} else {
		fmt.Println("failure:", resp.Errors)
	}
}

func deleteCustomerById() {
	deleteResponse, err := customerAsaas.DeleteById(context.TODO(), "")
	if err != nil {
		fmt.Println("error:", err)
	} else if deleteResponse.IsSuccess() {
		fmt.Println("success:", deleteResponse)
	} else {
		fmt.Println("failure:", deleteResponse.Errors)
	}
}

func getCustomerById() {
	resp, err := customerAsaas.GetById(context.TODO(), "")
	if err != nil {
		fmt.Println("error:", err)
	} else if resp.IsSuccess() {
		fmt.Println("success:", resp)
	} else if resp.IsNoContent() {
		fmt.Println("no content:", resp)
	} else {
		fmt.Println("failure:", resp.Errors)
	}
}

func getAllCustomers() {
	resp, err := customerAsaas.GetAll(context.TODO(), asaas.GetAllCustomersRequest{
		Name:              "",
		Email:             "",
		CpfCnpj:           "",
		GroupName:         "",
		ExternalReference: "",
		Offset:            0,
		Limit:             0,
	})
	if err != nil {
		fmt.Println("error:", err)
	} else if resp.IsSuccess() {
		fmt.Println("success:", resp)
	} else if resp.IsNoContent() {
		fmt.Println("no content:", resp)
	} else {
		fmt.Println("failure:", resp.Errors)
	}
}
