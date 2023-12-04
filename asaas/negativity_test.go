package asaas

import (
	"context"
	"os"
	"testing"
	"time"
)

func TestNegativityCreate(t *testing.T) {
	accessToken := getEnvValue(EnvAccessToken)
	assertFatalStringBlank(t, accessToken)
	initBankSlipCharge(false)
	chargeId := getEnvValue(EnvBankSlipChargeId)
	assertFatalStringBlank(t, chargeId)
	f, err := os.Open(getEnvValue(EnvFileName))
	assertFatalErrorNonnull(t, err)
	v, err := os.ReadFile(f.Name())
	assertFatalErrorNonnull(t, err)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nNegativity := NewNegativity(EnvSandbox, accessToken)
	resp, err := nNegativity.Create(ctx, CreateNegativityRequest{
		Payment:                chargeId,
		Type:                   NegativityTypeCreditBureau,
		Description:            "Unit test golang",
		CustomerName:           "Unit test golang",
		CustomerCpfCnpj:        "24971563792",
		CustomerPrimaryPhone:   "47999376637",
		CustomerSecondaryPhone: "",
		CustomerPostalCode:     "01310-000",
		CustomerAddress:        "Av. Paulista",
		CustomerAddressNumber:  "150",
		CustomerComplement:     "",
		CustomerProvince:       "Centro",
		Documents: &FileRequest{
			Name: f.Name(),
			Mime: FileMimeTypeText,
			Data: v,
		},
	})
	assertResponseSuccess(t, resp, err)
}

func TestNegativitySimulate(t *testing.T) {
	accessToken := getEnvValue(EnvAccessToken)
	assertFatalStringBlank(t, accessToken)
	initBankSlipCharge(false)
	chargeId := getEnvValue(EnvBankSlipChargeId)
	assertFatalStringBlank(t, chargeId)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nNegativity := NewNegativity(EnvSandbox, accessToken)
	resp, err := nNegativity.Simulate(ctx, chargeId)
	assertResponseSuccess(t, resp, err)
}

func TestNegativityResendDocumentsById(t *testing.T) {
	accessToken := getEnvValue(EnvAccessToken)
	assertFatalStringBlank(t, accessToken)
	initNegativity()
	negativityId := getEnvValue(EnvNegativityId)
	assertFatalStringBlank(t, negativityId)
	f, err := os.Open(getEnvValue(EnvFileName))
	assertFatalErrorNonnull(t, err)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nNegativity := NewNegativity(EnvSandbox, accessToken)
	resp, err := nNegativity.ResendDocumentsById(ctx, negativityId, NegativityResendDocumentsRequest{
		Documents: f,
	})
	assertResponseSuccess(t, resp, err)
}

func TestNegativityCancelById(t *testing.T) {
	accessToken := getEnvValue(EnvAccessToken)
	assertFatalStringBlank(t, accessToken)
	initNegativity()
	negativityId := getEnvValue(EnvNegativityId)
	assertFatalStringBlank(t, negativityId)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nNegativity := NewNegativity(EnvSandbox, accessToken)
	resp, err := nNegativity.CancelById(ctx, negativityId)
	assertResponseSuccess(t, resp, err)
}

func TestNegativityGetById(t *testing.T) {
	accessToken := getEnvValue(EnvAccessToken)
	assertFatalStringBlank(t, accessToken)
	initNegativity()
	negativityId := getEnvValue(EnvNegativityId)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nNegativity := NewNegativity(EnvSandbox, accessToken)
	resp, err := nNegativity.GetById(ctx, negativityId)
	assertResponseSuccess(t, resp, err)
}

func TestNegativityGetAll(t *testing.T) {
	accessToken := getEnvValue(EnvAccessToken)
	assertFatalStringBlank(t, accessToken)
	initNegativity()
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nNegativity := NewNegativity(EnvSandbox, accessToken)
	resp, err := nNegativity.GetAll(ctx, GetAllNegativitiesRequest{
		Status:           "",
		Type:             "",
		Payment:          "",
		RequestStartDate: Date{},
		RequestEndDate:   Date{},
		Offset:           0,
		Limit:            10,
	})
	assertResponseSuccess(t, resp, err)
}

func TestNegativityGetHistoryById(t *testing.T) {
	accessToken := getEnvValue(EnvAccessToken)
	assertFatalStringBlank(t, accessToken)
	initNegativity()
	negativityId := getEnvValue(EnvNegativityId)
	assertFatalStringBlank(t, negativityId)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nNegativity := NewNegativity(EnvSandbox, accessToken)
	resp, err := nNegativity.GetHistoryById(ctx, negativityId, PageableDefaultRequest{
		Offset: 0,
		Limit:  10,
	})
	assertResponseSuccess(t, resp, err)
}

func TestNegativityGetPaymentsById(t *testing.T) {
	accessToken := getEnvValue(EnvAccessToken)
	assertFatalStringBlank(t, accessToken)
	initNegativity()
	negativityId := getEnvValue(EnvNegativityId)
	assertFatalStringBlank(t, negativityId)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nNegativity := NewNegativity(EnvSandbox, accessToken)
	resp, err := nNegativity.GetHistoryById(ctx, negativityId, PageableDefaultRequest{
		Offset: 0,
		Limit:  10,
	})
	assertResponseSuccess(t, resp, err)
}

func TestNegativityGetChargesAvailableForDunning(t *testing.T) {
	accessToken := getEnvValue(EnvAccessToken)
	assertFatalStringBlank(t, accessToken)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nNegativity := NewNegativity(EnvSandbox, accessToken)
	resp, err := nNegativity.GetChargesAvailableForDunning(ctx, PageableDefaultRequest{
		Offset: 0,
		Limit:  10,
	})
	assertResponseSuccess(t, resp, err)
}
