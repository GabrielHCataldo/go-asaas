package main

import (
	"context"
	"fmt"
	"github.com/GabrielHCataldo/go-asaas/asaas"
	"os"
)

var creditBureauAsaas asaas.CreditBureau

func main() {
	creditBureauAsaas = asaas.NewCreditBureau(asaas.EnvSandbox, os.Getenv("ASAAS_ACCESS_TOKEN"))
	getCreditBureauReport()
	getCreditBureauReportById()
	getAllCreditBureauReports()
}

func getCreditBureauReport() {
	resp, err := creditBureauAsaas.GetReport(context.TODO(), asaas.GetReportRequest{
		Customer: "cus_000005791749",
		CpfCnpj:  "",
		State:    "SP",
	})
	if err != nil {
		fmt.Println("error:", err)
	} else if resp.IsSuccess() {
		fmt.Println("success:", resp)
	} else {
		fmt.Println("failure:", resp.Errors)
	}
}

func getCreditBureauReportById() {
	resp, err := creditBureauAsaas.GetReportById(context.TODO(), "")
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

func getAllCreditBureauReports() {
	resp, err := creditBureauAsaas.GetAllReports(context.TODO(), asaas.GetAllReportsRequest{
		StartDate: nil,
		EndDate:   nil,
		Offset:    0,
		Limit:     10,
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
