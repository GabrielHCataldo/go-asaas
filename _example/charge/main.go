package main

import (
	"context"
	"fmt"
	"github.com/GabrielHCataldo/go-asaas/asaas"
	"os"
	"time"
)

var chargeAsaas asaas.Charge

func main() {
	chargeAsaas = asaas.NewCharge(asaas.EnvSandbox, os.Getenv("ASAAS_ACCESS_TOKEN"))
	createChargePix()
	createChargeBill()
	createChargeUndefined()
	createChargeCreditCard()
	updateChargeById()
	getChargeById()
	getAllCustomers()
	deleteChargeById()
}

func createChargePix() {
	chargeResponse, err := chargeAsaas.Create(context.TODO(), asaas.CreateChargeRequest{
		Customer:    "cus_000005799255",
		BillingType: asaas.BillingTypePix,
		Value:       10,
		DueDate:     asaas.NewDate(2023, 12, 2, time.Local),
		Description: "Example pix charge",
	})
	if err != nil {
		fmt.Println("error:", err)
	} else if chargeResponse.IsSuccess() {
		fmt.Println("success:", chargeResponse)
	} else {
		fmt.Println("failure:", chargeResponse.Errors)
	}
}

func createChargeBill() {
	chargeResponse, err := chargeAsaas.Create(context.TODO(), asaas.CreateChargeRequest{
		Customer:    "cus_000005799255",
		BillingType: asaas.BillingTypeBill,
		Value:       10,
		DueDate:     asaas.NewDate(2023, 12, 2, time.Local),
		Description: "Example bill charge",
	})
	if err != nil {
		fmt.Println("error:", err)
	} else if chargeResponse.IsSuccess() {
		fmt.Println("success:", chargeResponse)
	} else {
		fmt.Println("failure:", chargeResponse.Errors)
	}
}

func createChargeCreditCard() {
	chargeResponse, err := chargeAsaas.Create(context.TODO(), asaas.CreateChargeRequest{
		Customer:    "cus_000005799255",
		BillingType: asaas.BillingTypeCreditCard,
		Value:       10,
		DueDate:     asaas.NewDate(2023, 12, 1, time.Local),
		Description: "Example bill charge",
		CreditCard: &asaas.CreditCardRequest{
			HolderName:  "unit test go",
			Number:      "4000000000000010",
			ExpiryMonth: "05",
			ExpiryYear:  "2035",
			Ccv:         "123",
		},
	})
	if err != nil {
		fmt.Println("error:", err)
	} else if chargeResponse.IsSuccess() {
		fmt.Println("success:", chargeResponse)
	} else {
		fmt.Println("failure:", chargeResponse.Errors)
	}
}

func createChargeUndefined() {
	chargeResponse, err := chargeAsaas.Create(context.TODO(), asaas.CreateChargeRequest{
		Customer:    "cus_000005799255",
		BillingType: asaas.BillingTypeUndefined,
		Value:       10,
		DueDate:     asaas.NewDate(2023, 12, 2, time.Local),
		Description: "Example bill charge",
	})
	if err != nil {
		fmt.Println("error:", err)
	} else if chargeResponse.IsSuccess() {
		fmt.Println("success:", chargeResponse)
	} else {
		fmt.Println("failure:", chargeResponse.Errors)
	}
}

func updateChargeById() {
	chargeResponse, err := chargeAsaas.UpdateById(context.TODO(), "pay_jxqnfvp1qt8qpf5s", asaas.UpdateChargeRequest{
		BillingType:       asaas.BillingTypeBill,
		Value:             5,
		Description:       "updated",
		ExternalReference: "test",
		InstallmentCount:  2,
		InstallmentValue:  2.5,
	})
	if err != nil {
		fmt.Println("error:", err)
	} else if chargeResponse.IsSuccess() {
		fmt.Println("success:", chargeResponse)
	} else {
		fmt.Println("failure:", chargeResponse.Errors)
	}
}

func deleteChargeById() {
	deleteResponse, err := chargeAsaas.DeleteById(context.TODO(), "pay_484rmiebm04419ey")
	if err != nil {
		fmt.Println("error:", err)
	} else if deleteResponse.IsSuccess() {
		fmt.Println("success:", deleteResponse)
	} else {
		fmt.Println("failure:", deleteResponse.Errors)
	}
}

func getChargeById() {
	chargeResponse, err := chargeAsaas.GetById(context.TODO(), "pay_jxqnfvp1qt8qpf5s")
	if err != nil {
		fmt.Println("error:", err)
	} else if chargeResponse.IsSuccess() {
		fmt.Println("success:", chargeResponse)
	} else {
		fmt.Println("failure:", chargeResponse.Errors)
	}
}

func getAllCustomers() {
	pageableResponse, err := chargeAsaas.GetAll(context.TODO(), asaas.GetAllChargesRequest{})
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
