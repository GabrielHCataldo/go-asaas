package main

import (
	"context"
	"fmt"
	"github.com/GabrielHCataldo/go-asaas/asaas"
	"os"
)

var negativityAsaas asaas.Negativity

func main() {
	negativityAsaas = asaas.NewNegativity(asaas.EnvSandbox, os.Getenv("ASAAS_ACCESS_TOKEN"))
	createNegativity()
	cancelNegativityById()
	getNegativityHistoryById()
	getChargesAvailableForDunning()
}

func createNegativity() {
	resp, err := negativityAsaas.Create(context.TODO(), asaas.CreateNegativityRequest{
		Payment:                "",
		Type:                   "",
		Description:            "",
		CustomerName:           "",
		CustomerCpfCnpj:        "",
		CustomerPrimaryPhone:   "",
		CustomerSecondaryPhone: "",
		CustomerPostalCode:     "",
		CustomerAddress:        "",
		CustomerAddressNumber:  "",
		CustomerComplement:     "",
		CustomerProvince:       "",
		Documents:              nil,
	})
	if err != nil {
		fmt.Println("error:", err)
	} else if resp.IsSuccess() {
		fmt.Println("success:", resp)
	} else {
		fmt.Println("failure:", resp.Errors)
	}
}

func cancelNegativityById() {
	resp, err := negativityAsaas.CancelById(context.TODO(), "")
	if err != nil {
		fmt.Println("error:", err)
	} else if resp.IsSuccess() {
		fmt.Println("success:", resp)
	} else if resp.IsNoContent() {
		fmt.Println("no content:", resp.Errors)
	} else {
		fmt.Println("failure:", resp.Errors)
	}
}

func getNegativityHistoryById() {
	resp, err := negativityAsaas.GetHistoryById(context.TODO(), "", asaas.PageableDefaultRequest{
		Offset: 0,
		Limit:  10,
	})
	if err != nil {
		fmt.Println("error:", err)
	} else if resp.IsSuccess() {
		fmt.Println("success:", resp)
	} else if resp.IsNoContent() {
		fmt.Println("no content:", resp.Errors)
	} else {
		fmt.Println("failure:", resp.Errors)
	}
}

func getChargesAvailableForDunning() {
	resp, err := negativityAsaas.GetChargesAvailableForDunning(context.TODO(), asaas.PageableDefaultRequest{
		Offset: 0,
		Limit:  10,
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
