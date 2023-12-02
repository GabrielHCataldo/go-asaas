package main

import (
	"context"
	"fmt"
	"github.com/GabrielHCataldo/go-asaas/asaas"
	"os"
)

var accountAsaas asaas.Account

func main() {
	accountAsaas = asaas.NewAccount(asaas.EnvSandbox, os.Getenv("ASAAS_ACCESS_TOKEN"))
	updateAccount()
	getAccountBalance()
	getAccountStatement()
}

func getAccountBalance() {
	resp, err := accountAsaas.GetBalance(context.TODO())
	if err != nil {
		fmt.Println("error:", err)
	} else if resp.IsSuccess() {
		fmt.Println("success:", resp)
	} else {
		fmt.Println("failure:", resp.Errors)
	}
}

func updateAccount() {
	resp, err := accountAsaas.Update(context.TODO(), asaas.UpdateAccountRequest{
		PersonType:    asaas.PersonTypePhysical,
		CpfCnpj:       "29376892000101",
		BirthDate:     asaas.Date{},
		CompanyType:   asaas.CompanyTypeLimited,
		Email:         "xxxxxx@gmail.com",
		Phone:         "",
		MobilePhone:   "",
		Site:          "",
		PostalCode:    "69620-970",
		Address:       "Praça São Cristovão, s/n",
		AddressNumber: "10",
		Complement:    "",
		Province:      "Centro",
	})
	if err != nil {
		fmt.Println("error:", err)
	} else if resp.IsSuccess() {
		fmt.Println("success:", resp)
	} else {
		fmt.Println("failure:", resp.Errors)
	}
}

func getAccountStatement() {
	resp, err := accountAsaas.GetAccountStatement(context.TODO(), asaas.GetAccountStatementRequest{
		StartDate:  nil,
		FinishDate: nil,
		Offset:     0,
		Limit:      10,
		Order:      asaas.OrderDesc,
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
