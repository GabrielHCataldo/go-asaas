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
		PersonType:    "",
		CpfCnpj:       "",
		BirthDate:     asaas.Date{},
		CompanyType:   nil,
		Email:         "",
		Phone:         nil,
		MobilePhone:   nil,
		Site:          nil,
		PostalCode:    "",
		Address:       nil,
		AddressNumber: nil,
		Complement:    nil,
		Province:      nil,
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
		StartDate:  asaas.Date{},
		FinishDate: asaas.Date{},
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
