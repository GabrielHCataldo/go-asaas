package asaas

import (
	"context"
	"encoding/json"
	"github.com/GabrielHCataldo/go-asaas/internal/test"
	"os"
	"testing"
	"time"
)

func TestChargeCreateError(t *testing.T) {
	accessToken, err := test.GetAccessTokenByEnv()
	AssertFatalErrorNonnull(t, err)
	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()
	nCharge := NewCharge(SANDBOX, *accessToken)
	_, errAsaas := nCharge.Create(ctx, CreateChargeRequest{})
	AssertSuccessNonnull(t, errAsaas)
}

func TestChangeCreateCreditCardSuccess(t *testing.T) {
	accessToken, err := test.GetAccessTokenByEnv()
	AssertFatalErrorNonnull(t, err)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	var chargeReq *CreateChargeRequest
	err = json.Unmarshal(test.GetChargeCreditCardRequestDefault(), chargeReq)
	AssertFatalErrorNonnull(t, err)
	nCharge := NewCharge(SANDBOX, *accessToken)
	resp, errAsaas := nCharge.Create(ctx, *chargeReq)
	AssertAsaasResponseSuccess(t, resp, errAsaas)
}

func TestChangeCreateCreditCardFailure(t *testing.T) {
	accessToken, err := test.GetAccessTokenByEnv()
	AssertFatalErrorNonnull(t, err)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	chargeReq := &CreateChargeRequest{}
	err = json.Unmarshal(test.GetChargeCreditCardFailureRequestDefault(), chargeReq)
	AssertFatalErrorNonnull(t, err)
	nCharge := NewCharge(SANDBOX, *accessToken)
	resp, errAsaas := nCharge.Create(ctx, *chargeReq)
	AssertAsaasResponseFailure(t, resp, errAsaas)
}

func TestChangeCreatePixSuccess(t *testing.T) {
	accessToken, err := test.GetAccessTokenByEnv()
	AssertFatalErrorNonnull(t, err)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	var chargeReq *CreateChargeRequest
	err = json.Unmarshal(test.GetChargePixRequestDefault(), chargeReq)
	AssertFatalErrorNonnull(t, err)
	nCharge := NewCharge(SANDBOX, *accessToken)
	resp, errAsaas := nCharge.Create(ctx, *chargeReq)
	AssertAsaasResponseSuccess(t, resp, errAsaas)
}

func TestChangeCreateBoletoSuccess(t *testing.T) {
	accessToken, err := test.GetAccessTokenByEnv()
	AssertFatalErrorNonnull(t, err)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	var chargeReq *CreateChargeRequest
	err = json.Unmarshal(test.GetChargeBoletoRequestDefault(), chargeReq)
	AssertFatalErrorNonnull(t, err)
	nCharge := NewCharge(SANDBOX, *accessToken)
	resp, errAsaas := nCharge.Create(ctx, *chargeReq)
	AssertAsaasResponseSuccess(t, resp, errAsaas)
}

func TestChargeUploadDocumentByID(t *testing.T) {
	accessToken, err := test.GetAccessTokenByEnv()
	AssertFatalErrorNonnull(t, err)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	f, err := test.GetSimpleFile()
	defer func(name string) {
		err = os.Remove(name)
		if err != nil {
			logError("error remove file test:", err)
		}
	}(f.Name())
	nCharge := NewCharge(SANDBOX, *accessToken)
	resp, errAsaas := nCharge.UploadDocumentByID(ctx, test.GetChargeIdDefault(), UploadDocumentRequest{
		AvailableAfterPayment: false,
		Type:                  DOCUMENT,
		File:                  f,
	})
	AssertAsaasResponseFailure(t, resp, errAsaas)
}

func TestChargeGetAll(t *testing.T) {
	accessToken, err := test.GetAccessTokenByEnv()
	AssertFatalErrorNonnull(t, err)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nCharge := NewCharge(SANDBOX, *accessToken)
	resp, errAsaas := nCharge.GetAll(ctx, GetAllChargesRequest{
		Status: CHARGE_CONFIRMED,
	})
	AssertAsaasResponseSuccess(t, resp, errAsaas)
}
