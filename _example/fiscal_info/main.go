package main

import (
	"context"
	"fmt"
	"github.com/GabrielHCataldo/go-asaas/asaas"
	"os"
)

var fiscalInfoAsaas asaas.FiscalInfo

func main() {
	fiscalInfoAsaas = asaas.NewFiscalInfo(asaas.EnvSandbox, os.Getenv("ASAAS_ACCESS_TOKEN"))
	saveFiscalInfo()
	getFiscalInfo()
	getAllFiscalInfoServices()
}

func saveFiscalInfo() {
	resp, err := fiscalInfoAsaas.Save(context.TODO(), asaas.SaveFiscalInfoRequest{
		Email:                    "",
		MunicipalInscription:     nil,
		SimplesNacional:          nil,
		CulturalProjectsPromoter: nil,
		Cnae:                     nil,
		SpecialTaxRegime:         nil,
		ServiceListItem:          nil,
		RpsSerie:                 nil,
		RpsNumber:                nil,
		LoteNumber:               nil,
		Username:                 nil,
		Password:                 nil,
		AccessToken:              nil,
		CertificateFile:          nil,
		CertificatePassword:      nil,
	})
	if err != nil {
		fmt.Println("error:", err)
	} else if resp.IsSuccess() {
		fmt.Println("success:", resp)
	} else {
		fmt.Println("failure:", resp.Errors)
	}
}

func getFiscalInfo() {
	resp, err := fiscalInfoAsaas.Get(context.TODO())
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

func getAllFiscalInfoServices() {
	resp, err := fiscalInfoAsaas.GetAllServices(context.TODO(), asaas.GetAllServicesRequest{
		Description: "",
		Offset:      0,
		Limit:       0,
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
