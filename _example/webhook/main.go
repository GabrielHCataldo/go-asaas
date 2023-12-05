package main

import (
	"context"
	"fmt"
	"github.com/GabrielHCataldo/go-asaas/asaas"
	"os"
)

var webhookAsaas asaas.Webhook

func main() {
	webhookAsaas = asaas.NewWebhook(asaas.EnvSandbox, os.Getenv("ASAAS_ACCESS_TOKEN"))
	saveWebhookSetting()
	getWebhookSetting()
}

func saveWebhookSetting() {
	resp, err := webhookAsaas.SaveSetting(context.TODO(), asaas.WebhookTypePayment, asaas.SaveWebhookSettingRequest{
		Url:         "",
		Email:       "",
		ApiVersion:  "",
		Enabled:     nil,
		Interrupted: nil,
		AuthToken:   "",
	})
	if err != nil {
		fmt.Println("error:", err)
	} else if resp.IsSuccess() {
		fmt.Println("success:", resp)
	} else {
		fmt.Println("failure:", resp.Errors)
	}
}

func getWebhookSetting() {
	resp, err := webhookAsaas.GetSetting(context.TODO(), asaas.WebhookTypePayment)
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
