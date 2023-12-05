package main

import (
	"context"
	"fmt"
	"github.com/GabrielHCataldo/go-asaas/asaas"
	"os"
)

var subaccountAsaas asaas.Subaccount

func main() {
	subaccountAsaas = asaas.NewSubaccount(asaas.EnvSandbox, os.Getenv("ASAAS_ACCESS_TOKEN"))
	createSubaccount()
	getSubaccountById()
	getAllSubaccount()
}

func createSubaccount() {
	resp, err := subaccountAsaas.Create(context.TODO(), asaas.CreateSubaccountRequest{
		Name:          "",
		Email:         "",
		LoginEmail:    "",
		CpfCnpj:       "",
		BirthDate:     asaas.Date{},
		CompanyType:   "",
		Phone:         "",
		MobilePhone:   "",
		Site:          "",
		Address:       "",
		AddressNumber: "",
		Complement:    "",
		Province:      "",
		PostalCode:    "",
		Webhooks:      nil,
	})
	if err != nil {
		fmt.Println("error:", err)
	} else if resp.IsSuccess() {
		fmt.Println("success:", resp)
	} else {
		fmt.Println("failure:", resp.Errors)
	}
}

func getSubaccountById() {
	resp, err := subaccountAsaas.GetById(context.TODO(), "")
	if err != nil {
		fmt.Println("error:", err)
	} else if resp.IsSuccess() {
		fmt.Println("success:", resp)
	} else if resp.IsNoContent() {
		fmt.Println("no content:", resp)
	} else {
		fmt.Println("no content:", resp)
	}
}

func getAllSubaccount() {
	resp, err := subaccountAsaas.GetAll(context.TODO(), asaas.GetAllSubaccountsRequest{
		CpfCnpj:  "",
		Email:    "",
		Name:     "",
		WalletId: "",
		Offset:   0,
		Limit:    0,
	})
	if err != nil {
		fmt.Println("error:", err)
	} else if resp.IsSuccess() {
		fmt.Println("success:", resp)
	} else if resp.IsNoContent() {
		fmt.Println("no content:", resp)
	} else {
		fmt.Println("no content:", resp)
	}
}
