package main

import (
	"context"
	"fmt"
	"github.com/GabrielHCataldo/go-asaas/asaas"
	"os"
)

var transferAsaas asaas.Transfer

func main() {
	transferAsaas = asaas.NewTransfer(asaas.EnvSandbox, os.Getenv("ASAAS_ACCESS_TOKEN"))
	transferToBank()
	transferToAsaas()
	cancelById()
	getTransferById()
	getAllTransfer()
}

func transferToBank() {
	resp, err := transferAsaas.TransferToBank(context.TODO(), asaas.TransferToBankRequest{
		Value:             0,
		BankAccount:       asaas.BackAccountRequest{},
		OperationType:     "",
		PixAddressKey:     "",
		PixAddressKeyType: "",
		Description:       "",
		ScheduleDate:      asaas.Date{},
	})
	if err != nil {
		fmt.Println("error:", err)
	} else if resp.IsSuccess() {
		fmt.Println("success:", resp)
	} else {
		fmt.Println("failure:", resp.Errors)
	}
}

func transferToAsaas() {
	resp, err := transferAsaas.TransferToAsaas(context.TODO(), asaas.TransferToAssasRequest{
		Value:    0,
		WalletId: "",
	})
	if err != nil {
		fmt.Println("error:", err)
	} else if resp.IsSuccess() {
		fmt.Println("success:", resp)
	} else {
		fmt.Println("failure:", resp.Errors)
	}
}

func cancelById() {
	resp, err := transferAsaas.CancelById(context.TODO(), "")
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

func getTransferById() {
	resp, err := transferAsaas.GetById(context.TODO(), "")
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

func getAllTransfer() {
	resp, err := transferAsaas.GetAll(context.TODO(), asaas.GetAllTransfersRequest{
		DateCreatedGe:  asaas.Date{},
		DateCreatedLe:  asaas.Date{},
		TransferDateGe: asaas.Date{},
		TransferDateLe: asaas.Date{},
		Type:           "",
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
