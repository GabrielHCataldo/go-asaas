package asaas

import (
	"context"
	"testing"
	"time"
)

func TestFiscalInfoSave(t *testing.T) {
	accessToken := getEnvValue(EnvAccessToken)
	assertFatalStringBlank(t, accessToken)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nFiscalInfo := NewFiscalInfo(EnvSandbox, accessToken)
	resp, err := nFiscalInfo.Save(ctx, SaveFiscalInfoRequest{
		Email:                    "test@gmail.com",
		MunicipalInscription:     Pointer("15.54.74"),
		SimplesNacional:          Pointer(true),
		CulturalProjectsPromoter: nil,
		Cnae:                     Pointer("6201501"),
		SpecialTaxRegime:         nil,
		ServiceListItem:          nil,
		RpsSerie:                 nil,
		RpsNumber:                Pointer(21),
		LoteNumber:               nil,
		Username:                 nil,
		Password:                 Pointer("test"),
		AccessToken:              nil,
		CertificateFile:          nil,
		CertificatePassword:      nil,
	})
	assertResponseSuccess(t, resp, err)
}

func TestFiscalInfoGet(t *testing.T) {
	accessToken := getEnvValue(EnvAccessToken)
	assertFatalStringBlank(t, accessToken)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nFiscalInfo := NewFiscalInfo(EnvSandbox, accessToken)
	resp, err := nFiscalInfo.Get(ctx)
	assertResponseSuccess(t, resp, err)
}

func TestFiscalInfoGetMunicipalSettings(t *testing.T) {
	accessToken := getEnvValue(EnvAccessToken)
	assertFatalStringBlank(t, accessToken)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nFiscalInfo := NewFiscalInfo(EnvSandbox, accessToken)
	resp, err := nFiscalInfo.GetMunicipalSettings(ctx)
	assertResponseSuccess(t, resp, err)
}

func TestFiscalInfoGetAllServices(t *testing.T) {
	accessToken := getEnvValue(EnvAccessToken)
	assertFatalStringBlank(t, accessToken)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nFiscalInfo := NewFiscalInfo(EnvSandbox, accessToken)
	resp, err := nFiscalInfo.GetAllServices(ctx, GetAllServicesRequest{
		Description: "",
		Offset:      0,
		Limit:       10,
	})
	assertResponseSuccess(t, resp, err)
}
