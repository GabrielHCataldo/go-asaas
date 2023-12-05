package main

import (
	"context"
	"fmt"
	"github.com/GabrielHCataldo/go-asaas/asaas"
	"os"
)

var mobilePhoneAsaas asaas.MobilePhone

func main() {
	mobilePhoneAsaas = asaas.NewMobilePhone(asaas.EnvSandbox, os.Getenv("ASAAS_ACCESS_TOKEN"))
	rechargeMobilePhone()
	cancelRechargeMobilePhoneById()
	getMobilePhoneProviderByPhoneNumber()
	getAllMobilePhoneRecharges()
}

func rechargeMobilePhone() {
	resp, err := mobilePhoneAsaas.Recharge(context.TODO(), asaas.MobilePhoneRechargeRequest{
		PhoneNumber: "",
		Value:       0,
	})
	if err != nil {
		fmt.Println("error:", err)
	} else if resp.IsSuccess() {
		fmt.Println("success:", resp)
	} else {
		fmt.Println("failure:", resp.Errors)
	}
}

func cancelRechargeMobilePhoneById() {
	resp, err := mobilePhoneAsaas.CancelRechargeById(context.TODO(), "")
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

func getMobilePhoneProviderByPhoneNumber() {
	resp, err := mobilePhoneAsaas.GetProviderByPhoneNumber(context.TODO(), "")
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

func getAllMobilePhoneRecharges() {
	resp, err := mobilePhoneAsaas.GetAllRecharges(context.TODO(), asaas.PageableDefaultRequest{
		Offset: 0,
		Limit:  0,
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
