package main

import (
	"context"
	"fmt"
	"github.com/GabrielHCataldo/go-asaas/asaas"
	"os"
)

var installmentAsaas asaas.Installment

func main() {
	installmentAsaas = asaas.NewInstallment(asaas.EnvSandbox, os.Getenv("ASAAS_ACCESS_TOKEN"))
	refundInstallmentById()
	getInstallmentById()
	getAllInstallments()
	deleteInstallmentById()
}

func refundInstallmentById() {
	resp, err := installmentAsaas.RefundById(context.TODO(), "")
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

func deleteInstallmentById() {
	resp, err := installmentAsaas.DeleteById(context.TODO(), "")
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

func getInstallmentById() {
	resp, err := installmentAsaas.GetById(context.TODO(), "")
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

func getAllInstallments() {
	resp, err := installmentAsaas.GetAll(context.TODO(), asaas.PageableDefaultRequest{
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
