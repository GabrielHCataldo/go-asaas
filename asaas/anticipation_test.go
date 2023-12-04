package asaas

import (
	"context"
	"testing"
	"time"
)

func TestAnticipationSimulate(t *testing.T) {
	accessToken := getEnvValue(EnvAccessToken)
	assertFatalStringBlank(t, accessToken)
	initCreditCardCharge(true, false)
	chargeId := getEnvValue(EnvCreditCardChargeId)
	assertFatalStringBlank(t, chargeId)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nAnticipation := NewAnticipation(EnvSandbox, accessToken)
	resp, errAsaas := nAnticipation.Simulate(ctx, AnticipationSimulateRequest{
		Payment:     chargeId,
		Installment: "",
	})
	assertResponseSuccess(t, resp, errAsaas)
}

func TestAnticipationRequest(t *testing.T) {
	accessToken := getEnvValue(EnvAccessToken)
	assertFatalStringBlank(t, accessToken)
	initCreditCardCharge(true, false)
	chargeId := getEnvValue(EnvCreditCardChargeId)
	assertFatalStringBlank(t, chargeId)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nAnticipation := NewAnticipation(EnvSandbox, accessToken)
	resp, errAsaas := nAnticipation.Request(ctx, AnticipationRequest{
		Payment:     chargeId,
		Installment: "",
		Documents:   nil,
	})
	assertResponseSuccess(t, resp, errAsaas)
}

func TestAnticipationAgreementSign(t *testing.T) {
	accessToken := getEnvValue(EnvAccessToken)
	assertFatalStringBlank(t, accessToken)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nAnticipation := NewAnticipation(EnvSandbox, accessToken)
	resp, errAsaas := nAnticipation.AgreementSign(ctx, AgreementSignRequest{
		Agreed: true,
	})
	assertResponseSuccess(t, resp, errAsaas)
}

func TestAnticipationGetById(t *testing.T) {
	accessToken := getEnvValue(EnvAccessToken)
	assertFatalStringBlank(t, accessToken)
	initAnticipation()
	anticipationId := getEnvValue(EnvAnticipationId)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nAnticipation := NewAnticipation(EnvSandbox, accessToken)
	resp, errAsaas := nAnticipation.GetById(ctx, anticipationId)
	assertResponseSuccess(t, resp, errAsaas)
}

func TestAnticipationGetLimits(t *testing.T) {
	accessToken := getEnvValue(EnvAccessToken)
	assertFatalStringBlank(t, accessToken)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nAnticipation := NewAnticipation(EnvSandbox, accessToken)
	resp, errAsaas := nAnticipation.GetLimits(ctx)
	assertResponseSuccess(t, resp, errAsaas)
}

func TestAnticipationGetAll(t *testing.T) {
	accessToken := getEnvValue(EnvAccessToken)
	assertFatalStringBlank(t, accessToken)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nAnticipation := NewAnticipation(EnvSandbox, accessToken)
	resp, errAsaas := nAnticipation.GetAll(ctx, GetAllAnticipationsRequest{
		Payment:     "",
		Installment: "",
		Status:      "",
		Offset:      0,
		Limit:       10,
	})
	assertResponseSuccess(t, resp, errAsaas)
}
