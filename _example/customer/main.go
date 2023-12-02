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
	resp, err := customerAsaas.Create(context.TODO(), asaas.CustomerRequest{
		Name:                 "Go Asaas Test",
		CpfCnpj:              "85185238003",
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
	resp, err := customerAsaas.UpdateById(context.TODO(), "cus_000005799255", asaas.CustomerRequest{
		Name:                 "",
		CpfCnpj:              "",
		Email:                "test@test.com",
		Phone:                "",
		MobilePhone:          "48997576131",
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

func deleteCustomerById() {
	deleteResponse, err := customerAsaas.DeleteById(context.TODO(), "cus_000005791749")
	if err != nil {
		fmt.Println("error:", err)
	} else if deleteResponse.IsSuccess() {
		fmt.Println("success:", deleteResponse)
	} else {
		fmt.Println("failure:", deleteResponse.Errors)
	}
}

func getCustomerById() {
	resp, err := customerAsaas.GetById(context.TODO(), "cus_000005799255")
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
		Limit:             10,
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
