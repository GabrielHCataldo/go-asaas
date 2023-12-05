package main

import (
	"context"
	"fmt"
	"github.com/GabrielHCataldo/go-asaas/asaas"
	"os"
)

var anticipationAsaas asaas.Anticipation

func main() {
	anticipationAsaas = asaas.NewAnticipation(asaas.EnvSandbox, os.Getenv("ASAAS_ACCESS_TOKEN"))
	simulateAnticipation()
	requestAnticipation()
	getAnticipationById()
	getAllAnticipations()
}

func simulateAnticipation() {
	resp, err := anticipationAsaas.Simulate(context.TODO(), asaas.AnticipationSimulateRequest{
		Payment:     "",
		Installment: "",
	})
	if err != nil {
		fmt.Println("error:", err)
	} else if resp.IsSuccess() {
		fmt.Println("success:", resp)
	} else {
		fmt.Println("failure:", resp.Errors)
	}
}

func requestAnticipation() {
	resp, err := anticipationAsaas.Request(context.TODO(), asaas.AnticipationRequest{
		Payment:     "",
		Installment: "",
		Documents:   nil,
	})
	if err != nil {
		fmt.Println("error:", err)
	} else if resp.IsSuccess() {
		fmt.Println("success:", resp)
	} else {
		fmt.Println("failure:", resp.Errors)
	}
}

func getAnticipationById() {
	resp, err := anticipationAsaas.GetById(context.TODO(), "")
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

func getAllAnticipations() {
	resp, err := anticipationAsaas.GetAll(context.TODO(), asaas.GetAllAnticipationsRequest{
		Payment:     "",
		Installment: "",
		Status:      "",
		Offset:      0,
		Limit:       10,
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
