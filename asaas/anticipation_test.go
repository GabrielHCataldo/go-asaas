package asaas

import (
	"context"
	"testing"
	"time"
)

func TestAnticipationSimulateSuccess(t *testing.T) {
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

func TestAnticipationSimulateError(t *testing.T) {
	nAnticipation := NewAnticipation(EnvSandbox, "")
	resp, err := nAnticipation.Simulate(context.TODO(), AnticipationSimulateRequest{})
	assertResponseFailure(t, resp, err)
}

func TestAnticipationRequestSuccess(t *testing.T) {
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

func TestAnticipationRequestError(t *testing.T) {
	nAnticipation := NewAnticipation(EnvSandbox, "")
	resp, err := nAnticipation.Request(context.TODO(), AnticipationRequest{})
	assertResponseFailure(t, resp, err)
}

func TestAnticipationAgreementSignSuccess(t *testing.T) {
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

func TestAnticipationAgreementSignError(t *testing.T) {
	nAnticipation := NewAnticipation(EnvSandbox, "")
	resp, err := nAnticipation.AgreementSign(context.TODO(), AgreementSignRequest{})
	assertResponseFailure(t, resp, err)
}

func TestAnticipationGetByIdSuccess(t *testing.T) {
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

func TestAnticipationGetByIdError(t *testing.T) {
	nAnticipation := NewAnticipation(EnvSandbox, "")
	resp, err := nAnticipation.GetById(context.TODO(), "")
	assertResponseFailure(t, resp, err)
}

func TestAnticipationGetLimitsSuccess(t *testing.T) {
	accessToken := getEnvValue(EnvAccessToken)
	assertFatalStringBlank(t, accessToken)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nAnticipation := NewAnticipation(EnvSandbox, accessToken)
	resp, errAsaas := nAnticipation.GetLimits(ctx)
	assertResponseSuccess(t, resp, errAsaas)
}

func TestAnticipationGetLimitsError(t *testing.T) {
	nAnticipation := NewAnticipation(EnvSandbox, "")
	resp, errAsaas := nAnticipation.GetLimits(context.TODO())
	assertResponseFailure(t, resp, errAsaas)
}

func TestAnticipationGetAllSuccess(t *testing.T) {
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

func TestAnticipationGetAllError(t *testing.T) {
	nAnticipation := NewAnticipation(EnvSandbox, "")
	resp, errAsaas := nAnticipation.GetAll(context.TODO(), GetAllAnticipationsRequest{})
	assertResponseFailure(t, resp, errAsaas)
}
