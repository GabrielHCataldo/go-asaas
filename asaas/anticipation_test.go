package asaas

import (
	"context"
	"testing"
	"time"
)

func init() {
	initCustomer()
	initCreditCardCharge()
}

func TestAnticipationSimulateSuccess(t *testing.T) {
	accessToken, err := getAccessToken()
	assertFatalErrorNonnull(t, err)
	chargeId, err := getCreditCardChargeId()
	assertFatalErrorNonnull(t, err)
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
	_, errAsaas := nAnticipation.Simulate(context.TODO(), AnticipationSimulateRequest{})
	assertSuccessNonnull(t, errAsaas)
}

func TestAnticipationRequestSuccess(t *testing.T) {
	accessToken, err := getAccessToken()
	assertFatalErrorNonnull(t, err)
	chargeId, err := getCreditCardChargeId()
	assertFatalErrorNonnull(t, err)
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
	_, errAsaas := nAnticipation.Request(context.TODO(), AnticipationRequest{})
	assertSuccessNonnull(t, errAsaas)
}

func TestAnticipationAgreementSignSuccess(t *testing.T) {
	accessToken, err := getAccessToken()
	assertFatalErrorNonnull(t, err)
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
	_, errAsaas := nAnticipation.AgreementSign(context.TODO(), AgreementSignRequest{})
	assertSuccessNonnull(t, errAsaas)
}

func TestAnticipationGetByIdSuccess(t *testing.T) {
	accessToken, err := getAccessToken()
	assertFatalErrorNonnull(t, err)
	initAnticipation()
	anticipationId, _ := getAnticipationId()
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nAnticipation := NewAnticipation(EnvSandbox, accessToken)
	resp, errAsaas := nAnticipation.GetById(ctx, anticipationId)
	assertResponseSuccess(t, resp, errAsaas)
}

func TestAnticipationGetByIdError(t *testing.T) {
	nAnticipation := NewAnticipation(EnvSandbox, "")
	_, errAsaas := nAnticipation.GetById(context.TODO(), "test")
	assertSuccessNonnull(t, errAsaas)
}

func TestAnticipationGetLimitsSuccess(t *testing.T) {
	accessToken, err := getAccessToken()
	assertFatalErrorNonnull(t, err)
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
	accessToken, err := getAccessToken()
	assertFatalErrorNonnull(t, err)
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
