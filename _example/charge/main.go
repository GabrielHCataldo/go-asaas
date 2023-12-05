package main

import (
	"context"
	"fmt"
	"github.com/GabrielHCataldo/go-asaas/asaas"
	"os"
)

var chargeAsaas asaas.Charge

func main() {
	chargeAsaas = asaas.NewCharge(asaas.EnvSandbox, os.Getenv("ASAAS_ACCESS_TOKEN"))
	createCharge()
	updateChargeById()
	getChargeById()
	getAllCustomers()
	deleteChargeById()
}

func createCharge() {
	resp, err := chargeAsaas.Create(context.TODO(), asaas.CreateChargeRequest{
		Customer:             "",
		BillingType:          "",
		Value:                0,
		DueDate:              asaas.Date{},
		Description:          "",
		ExternalReference:    "",
		Discount:             nil,
		Interest:             nil,
		Fine:                 nil,
		PostalService:        false,
		Split:                nil,
		Callback:             nil,
		CreditCard:           nil,
		CreditCardHolderInfo: nil,
		CreditCardToken:      "",
		InstallmentCount:     0,
		InstallmentValue:     0,
		AuthorizeOnly:        false,
		RemoteIp:             "",
	})
	if err != nil {
		fmt.Println("error:", err)
	} else if resp.IsSuccess() {
		fmt.Println("success:", resp)
	} else {
		fmt.Println("failure:", resp.Errors)
	}
}

func updateChargeById() {
	resp, err := chargeAsaas.UpdateById(context.TODO(), "", asaas.UpdateChargeRequest{
		Customer:          "",
		BillingType:       "",
		Value:             0,
		DueDate:           asaas.Date{},
		Description:       nil,
		ExternalReference: nil,
		Discount:          nil,
		Interest:          nil,
		Fine:              nil,
		PostalService:     nil,
		Split:             nil,
		Callback:          nil,
		InstallmentCount:  0,
		InstallmentValue:  0,
	})
	if err != nil {
		fmt.Println("error:", err)
	} else if resp.IsSuccess() {
		fmt.Println("success:", resp)
	} else {
		fmt.Println("failure:", resp.Errors)
	}
}

func deleteChargeById() {
	deleteResponse, err := chargeAsaas.DeleteById(context.TODO(), "")
	if err != nil {
		fmt.Println("error:", err)
	} else if deleteResponse.IsSuccess() {
		fmt.Println("success:", deleteResponse)
	} else {
		fmt.Println("failure:", deleteResponse.Errors)
	}
}

func getChargeById() {
	resp, err := chargeAsaas.GetById(context.TODO(), "")
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
	resp, err := chargeAsaas.GetAll(context.TODO(), asaas.GetAllChargesRequest{
		Customer:              "",
		Subscription:          "",
		Installment:           "",
		CustomerGroupName:     "",
		BillingType:           "",
		Status:                "",
		ExternalReference:     "",
		InvoiceStatus:         "",
		EstimatedCreditDate:   asaas.Date{},
		PixQrCodeId:           "",
		Anticipated:           asaas.Pointer(false),
		PaymentDate:           asaas.Date{},
		DateCreatedGe:         asaas.Date{},
		DateCreatedLe:         asaas.Date{},
		PaymentDateGe:         asaas.Date{},
		PaymentDateLe:         asaas.Date{},
		EstimatedCreditDateGE: asaas.Date{},
		EstimatedCreditDateLE: asaas.Date{},
		DueDateGe:             asaas.Date{},
		DueDateLe:             asaas.Date{},
		User:                  "",
		Offset:                0,
		Limit:                 0,
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
