package main

import (
	"context"
	"fmt"
	"github.com/GabrielHCataldo/go-asaas/asaas"
	"os"
)

var subscriptionAsaas asaas.Subscription

func main() {
	subscriptionAsaas = asaas.NewSubscription(asaas.EnvSandbox, os.Getenv("ASAAS_ACCESS_TOKEN"))
	createSubscription()
	updateSubscriptionById()
	getSubscriptionById()
	getAllSubscription()
	deleteSubscriptionById()
}

func createSubscription() {
	resp, err := subscriptionAsaas.Create(context.TODO(), asaas.CreateSubscriptionRequest{
		Customer:             "",
		BillingType:          "",
		Value:                0,
		NextDueDate:          asaas.Date{},
		Discount:             nil,
		Interest:             nil,
		Fine:                 nil,
		Cycle:                "",
		Description:          "",
		CreditCard:           nil,
		CreditCardHolderInfo: nil,
		CreditCardToken:      "",
		EndDate:              asaas.Date{},
		MaxPayments:          0,
		ExternalReference:    "",
		Split:                nil,
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

func updateSubscriptionById() {
	resp, err := subscriptionAsaas.UpdateById(context.TODO(), "", asaas.UpdateSubscriptionRequest{
		BillingType:           "",
		Value:                 0,
		Status:                "",
		NextDueDate:           asaas.Date{},
		Discount:              nil,
		Interest:              nil,
		Fine:                  nil,
		Cycle:                 "",
		Description:           nil,
		EndDate:               asaas.Date{},
		UpdatePendingPayments: false,
		ExternalReference:     "",
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

func deleteSubscriptionById() {
	resp, err := subscriptionAsaas.DeleteById(context.TODO(), "")
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

func getSubscriptionById() {
	resp, err := subscriptionAsaas.GetById(context.TODO(), "")
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

func getAllSubscription() {
	resp, err := subscriptionAsaas.GetAll(context.TODO(), asaas.GetAllSubscriptionsRequest{
		Customer:          "",
		CustomerGroupName: "",
		BillingType:       "",
		Status:            "",
		DeletedOnly:       false,
		IncludeDeleted:    false,
		ExternalReference: "",
		Order:             "",
		Sort:              "",
		Offset:            0,
		Limit:             0,
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
