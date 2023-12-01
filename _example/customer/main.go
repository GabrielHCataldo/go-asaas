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
	customerResponse, err := customerAsaas.Create(context.TODO(), asaas.CustomerRequest{
		Name:    "Go Asaas Test",
		CpfCnpj: "85185238003",
	})
	if err != nil {
		fmt.Println("error:", err)
	} else if customerResponse.IsSuccess() {
		fmt.Println("success:", customerResponse)
	} else {
		fmt.Println("failure:", customerResponse.Errors)
	}
}

func updateCustomerById() {
	customerResponse, err := customerAsaas.UpdateById(context.TODO(), "cus_000005799255", asaas.CustomerRequest{
		Email:       "test@test.com",
		MobilePhone: "48997576131",
	})
	if err != nil {
		fmt.Println("error:", err)
	} else if customerResponse.IsSuccess() {
		fmt.Println("success:", customerResponse)
	} else {
		fmt.Println("failure:", customerResponse.Errors)
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
	customerResponse, err := customerAsaas.GetById(context.TODO(), "cus_000005799255")
	if err != nil {
		fmt.Println("error:", err)
	} else if customerResponse.IsSuccess() {
		fmt.Println("success:", customerResponse)
	} else {
		fmt.Println("failure:", customerResponse.Errors)
	}
}

func getAllCustomers() {
	pageableResponse, err := customerAsaas.GetAll(context.TODO(), asaas.GetAllCustomersRequest{})
	if err != nil {
		fmt.Println("error:", err)
	} else if pageableResponse.IsSuccess() {
		fmt.Println("success:", pageableResponse)
	} else if pageableResponse.IsNoContent() {
		fmt.Println("no content:", pageableResponse)
	} else {
		fmt.Println("failure:", pageableResponse.Errors)
	}
}
