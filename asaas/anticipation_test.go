package asaas

import (
	"context"
	"os"
	"testing"
	"time"
)

func TestAnticipationSimulate(t *testing.T) {
	initCreditCardCharge(false, false)
	accessToken := getEnvValue(EnvAccessToken)
	chargeId := getEnvValue(EnvCreditCardChargeId)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nAnticipation := NewAnticipation(EnvSandbox, accessToken)
	resp, err := nAnticipation.Simulate(ctx, AnticipationSimulateRequest{
		Payment:     chargeId,
		Installment: "",
	})
	assertResponseSuccess(t, resp, err)
}

func TestAnticipationRequest(t *testing.T) {
	initCreditCardCharge(false, false)
	accessToken := getEnvValue(EnvAccessToken)
	chargeId := getEnvValue(EnvCreditCardChargeId)
	f, _ := os.Open(getEnvValue(EnvFileName))
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nAnticipation := NewAnticipation(EnvSandbox, accessToken)
	resp, err := nAnticipation.Request(ctx, AnticipationRequest{
		Payment:     chargeId,
		Installment: "",
		Documents:   []*os.File{f},
	})
	assertResponseSuccess(t, resp, err)
}

func TestAnticipationAgreementSign(t *testing.T) {
	accessToken := getEnvValue(EnvAccessToken)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nAnticipation := NewAnticipation(EnvSandbox, accessToken)
	resp, err := nAnticipation.AgreementSign(ctx, AgreementSignRequest{
		Agreed: true,
	})
	assertResponseSuccess(t, resp, err)
}

func TestAnticipationGetById(t *testing.T) {
	initAnticipation()
	accessToken := getEnvValue(EnvAccessToken)
	anticipationId := getEnvValue(EnvAnticipationId)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nAnticipation := NewAnticipation(EnvSandbox, accessToken)
	resp, err := nAnticipation.GetById(ctx, anticipationId)
	assertResponseSuccess(t, resp, err)
}

func TestAnticipationGetLimits(t *testing.T) {
	accessToken := getEnvValue(EnvAccessToken)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nAnticipation := NewAnticipation(EnvSandbox, accessToken)
	resp, err := nAnticipation.GetLimits(ctx)
	assertResponseSuccess(t, resp, err)
}

func TestAnticipationGetAll(t *testing.T) {
	accessToken := getEnvValue(EnvAccessToken)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nAnticipation := NewAnticipation(EnvSandbox, accessToken)
	resp, err := nAnticipation.GetAll(ctx, GetAllAnticipationsRequest{
		Payment:     "",
		Installment: "",
		Status:      "",
		Offset:      0,
		Limit:       10,
	})
	assertResponseSuccess(t, resp, err)
}
