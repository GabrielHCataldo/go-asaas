package asaas

import (
	"context"
	"os"
	"testing"
	"time"
)

func TestNegativityCreate(t *testing.T) {
	initBankSlipCharge(false)
	accessToken := getEnvValue(EnvAccessToken)
	chargeId := getEnvValue(EnvBankSlipChargeId)
	f, _ := os.Open(getEnvValue(EnvFileName))
	v, _ := os.ReadFile(f.Name())
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
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
	initBankSlipCharge(false)
	accessToken := getEnvValue(EnvAccessToken)
	chargeId := getEnvValue(EnvBankSlipChargeId)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nNegativity := NewNegativity(EnvSandbox, accessToken)
	resp, err := nNegativity.Simulate(ctx, chargeId)
	assertResponseSuccess(t, resp, err)
}

func TestNegativityResendDocumentsById(t *testing.T) {
	initNegativity()
	accessToken := getEnvValue(EnvAccessToken)
	negativityId := getEnvValue(EnvNegativityId)
	f, _ := os.Open(getEnvValue(EnvFileName))
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nNegativity := NewNegativity(EnvSandbox, accessToken)
	resp, err := nNegativity.ResendDocumentsById(ctx, negativityId, NegativityResendDocumentsRequest{
		Documents: f,
	})
	assertResponseSuccess(t, resp, err)
}

func TestNegativityCancelById(t *testing.T) {
	initNegativity()
	accessToken := getEnvValue(EnvAccessToken)
	negativityId := getEnvValue(EnvNegativityId)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nNegativity := NewNegativity(EnvSandbox, accessToken)
	resp, err := nNegativity.CancelById(ctx, negativityId)
	assertResponseSuccess(t, resp, err)
}

func TestNegativityGetById(t *testing.T) {
	initNegativity()
	accessToken := getEnvValue(EnvAccessToken)
	negativityId := getEnvValue(EnvNegativityId)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nNegativity := NewNegativity(EnvSandbox, accessToken)
	resp, err := nNegativity.GetById(ctx, negativityId)
	assertResponseSuccess(t, resp, err)
}

func TestNegativityGetAll(t *testing.T) {
	initNegativity()
	accessToken := getEnvValue(EnvAccessToken)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
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
	initNegativity()
	accessToken := getEnvValue(EnvAccessToken)
	negativityId := getEnvValue(EnvNegativityId)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nNegativity := NewNegativity(EnvSandbox, accessToken)
	resp, err := nNegativity.GetHistoryById(ctx, negativityId, PageableDefaultRequest{
		Offset: 0,
		Limit:  10,
	})
	assertResponseSuccess(t, resp, err)
}

func TestNegativityGetPaymentsById(t *testing.T) {
	initNegativity()
	accessToken := getEnvValue(EnvAccessToken)
	negativityId := getEnvValue(EnvNegativityId)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nNegativity := NewNegativity(EnvSandbox, accessToken)
	resp, err := nNegativity.GetPaymentsById(ctx, negativityId, PageableDefaultRequest{
		Offset: 0,
		Limit:  10,
	})
	assertResponseSuccess(t, resp, err)
}

func TestNegativityGetChargesAvailableForDunning(t *testing.T) {
	accessToken := getEnvValue(EnvAccessToken)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nNegativity := NewNegativity(EnvSandbox, accessToken)
	resp, err := nNegativity.GetChargesAvailableForDunning(ctx, PageableDefaultRequest{
		Offset: 0,
		Limit:  10,
	})
	assertResponseSuccess(t, resp, err)
}
