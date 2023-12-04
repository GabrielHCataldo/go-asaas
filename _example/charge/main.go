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
	resp, err := chargeAsaas.Create(context.TODO(), asaas.CreateChargeRequest{
		Customer:             "cus_000005799255",
		BillingType:          asaas.BillingTypePix,
		Value:                10,
		DueDate:              asaas.NewDate(2023, 12, 2, time.Local),
		Description:          "Example pix charge",
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
	})
	if err != nil {
		fmt.Println("error:", err)
	} else if resp.IsSuccess() {
		fmt.Println("success:", resp)
	} else {
		fmt.Println("failure:", resp.Errors)
	}
}

func createChargeBill() {
	resp, err := chargeAsaas.Create(context.TODO(), asaas.CreateChargeRequest{
		Customer:             "cus_000005799255",
		BillingType:          asaas.BillingTypeBankSlip,
		Value:                10,
		DueDate:              asaas.NewDate(2023, 12, 2, time.Local),
		Description:          "Example bill charge",
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
	})
	if err != nil {
		fmt.Println("error:", err)
	} else if resp.IsSuccess() {
		fmt.Println("success:", resp)
	} else {
		fmt.Println("failure:", resp.Errors)
	}
}

func createChargeCreditCard() {
	resp, err := chargeAsaas.Create(context.TODO(), asaas.CreateChargeRequest{
		Customer:          "cus_000005799255",
		BillingType:       asaas.BillingTypeCreditCard,
		Value:             10,
		DueDate:           asaas.NewDate(2023, 12, 1, time.Local),
		Description:       "Example bill charge",
		ExternalReference: "",
		Discount:          nil,
		Interest:          nil,
		Fine:              nil,
		PostalService:     false,
		Split:             nil,
		Callback:          nil,
		CreditCard: &asaas.CreditCardRequest{
			HolderName:  "unit test go",
			Number:      "4000000000000010",
			ExpiryMonth: "05",
			ExpiryYear:  "2035",
			Ccv:         "123",
		},
		CreditCardHolderInfo: &asaas.CreditCardHolderInfoRequest{
			Name:              "Example go",
			CpfCnpj:           "29376892000101",
			Email:             "example@gmail.com",
			Phone:             "4738010919",
			MobilePhone:       "47998781877",
			PostalCode:        "89223-005",
			AddressNumber:     "10",
			AddressComplement: "",
		},
		CreditCardToken:  "",
		InstallmentCount: 2,
		InstallmentValue: 5,
		AuthorizeOnly:    false,
		RemoteIp:         "",
	})
	if err != nil {
		fmt.Println("error:", err)
	} else if resp.IsSuccess() {
		fmt.Println("success:", resp)
	} else {
		fmt.Println("failure:", resp.Errors)
	}
}

func createChargeUndefined() {
	resp, err := chargeAsaas.Create(context.TODO(), asaas.CreateChargeRequest{
		Customer:             "cus_000005799255",
		BillingType:          asaas.BillingTypeUndefined,
		Value:                10,
		DueDate:              asaas.NewDate(2023, 12, 2, time.Local),
		Description:          "Example bill charge",
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
	resp, err := chargeAsaas.UpdateById(context.TODO(), "pay_jxqnfvp1qt8qpf5s", asaas.UpdateChargeRequest{
		Customer:          "",
		BillingType:       asaas.BillingTypeBankSlip,
		Value:             5,
		DueDate:           asaas.Date{},
		Description:       asaas.Pointer("updated"),
		ExternalReference: asaas.Pointer(""),
		Discount:          nil,
		Interest:          nil,
		Fine:              nil,
		PostalService:     asaas.Pointer(false),
		Split:             nil,
		Callback:          nil,
		InstallmentCount:  2,
		InstallmentValue:  2.5,
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
	resp, err := chargeAsaas.GetById(context.TODO(), "pay_jxqnfvp1qt8qpf5s")
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
		Limit:                 10,
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
