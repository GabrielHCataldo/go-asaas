package main

import (
	"context"
	"fmt"
	"github.com/GabrielHCataldo/go-asaas/asaas"
	"os"
)

var negativityAsaas asaas.Notification

func main() {
	negativityAsaas = asaas.NewNotification(asaas.EnvSandbox, os.Getenv("ASAAS_ACCESS_TOKEN"))
	updateNotificationById()
	updateNotificationManyByCustomer()
	getAllNotificationByCustomer()
}

func updateNotificationById() {
	resp, err := negativityAsaas.UpdateById(context.TODO(), "", asaas.UpdateNotificationRequest{
		Enabled:                     false,
		EmailEnabledForProvider:     false,
		SmsEnabledForProvider:       false,
		EmailEnabledForCustomer:     false,
		SmsEnabledForCustomer:       false,
		PhoneCallEnabledForCustomer: false,
		WhatsappEnabledForCustomer:  false,
		ScheduleOffset:              0,
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

func updateNotificationManyByCustomer() {
	resp, err := negativityAsaas.UpdateManyByCustomer(context.TODO(), asaas.UpdateManyNotificationsRequest{
		Customer: "",
		Notifications: []asaas.UpdateManyNotificationRequest{
			{
				Id:                          "",
				Enabled:                     false,
				EmailEnabledForProvider:     false,
				SmsEnabledForProvider:       false,
				EmailEnabledForCustomer:     false,
				SmsEnabledForCustomer:       false,
				PhoneCallEnabledForCustomer: false,
				WhatsappEnabledForCustomer:  false,
				ScheduleOffset:              0,
			},
		},
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

func getAllNotificationByCustomer() {
	resp, err := negativityAsaas.GetAllByCustomer(context.TODO(), "")
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
