package main

import (
	"context"
	"fmt"
	"github.com/GabrielHCataldo/go-asaas/asaas"
	"os"
)

var invoiceAsaas asaas.Invoice

func main() {
	invoiceAsaas = asaas.NewInvoice(asaas.EnvSandbox, os.Getenv("ASAAS_ACCESS_TOKEN"))
	scheduleInvoice()
	authorizeInvoiceById()
	cancelInvoiceById()
	getInvoiceById()
	getAllInvoices()
}

func scheduleInvoice() {
	resp, err := invoiceAsaas.Schedule(context.TODO(), asaas.ScheduleInvoiceRequest{
		Payment:              "",
		Installment:          "",
		Customer:             "",
		ServiceDescription:   "",
		Observations:         "",
		ExternalReference:    "",
		Value:                0,
		Deductions:           0,
		EffectiveDate:        asaas.Date{},
		MunicipalServiceId:   "",
		MunicipalServiceCode: "",
		MunicipalServiceName: "",
		UpdatePayment:        false,
		Taxes:                asaas.InvoiceTaxesRequest{},
	})
	if err != nil {
		fmt.Println("error:", err)
	} else if resp.IsSuccess() {
		fmt.Println("success:", resp)
	} else {
		fmt.Println("failure:", resp.Errors)
	}
}

func authorizeInvoiceById() {
	resp, err := invoiceAsaas.AuthorizeById(context.TODO(), "")
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

func cancelInvoiceById() {
	resp, err := invoiceAsaas.CancelById(context.TODO(), "")
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

func getInvoiceById() {
	resp, err := invoiceAsaas.GetById(context.TODO(), "")
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

func getAllInvoices() {
	resp, err := invoiceAsaas.GetAll(context.TODO(), asaas.GetAllInvoicesRequest{
		EffectiveDateGE:   asaas.Date{},
		EffectiveDateLE:   asaas.Date{},
		Payment:           "",
		Installment:       "",
		Customer:          "",
		ExternalReference: "",
		Status:            "",
		Offset:            0,
		Limit:             10,
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
