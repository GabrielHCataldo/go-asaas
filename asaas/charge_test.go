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
	req := &CreateChargeRequest{}
	err = json.Unmarshal(test.GetCreateChargeCreditCardRequestDefault(), req)
	AssertFatalErrorNonnull(t, err)
	nCharge := NewCharge(SANDBOX, *accessToken)
	resp, errAsaas := nCharge.Create(ctx, *req)
	AssertAsaasResponseSuccess(t, resp, errAsaas)
}

func TestChangeCreateCreditCardFailure(t *testing.T) {
	accessToken, err := test.GetAccessTokenByEnv()
	AssertFatalErrorNonnull(t, err)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	req := &CreateChargeRequest{}
	err = json.Unmarshal(test.GetCreateChargeCreditCardFailureRequestDefault(), req)
	AssertFatalErrorNonnull(t, err)
	nCharge := NewCharge(SANDBOX, *accessToken)
	resp, errAsaas := nCharge.Create(ctx, *req)
	AssertAsaasResponseFailure(t, resp, errAsaas)
}

func TestChangeCreatePixSuccess(t *testing.T) {
	accessToken, err := test.GetAccessTokenByEnv()
	AssertFatalErrorNonnull(t, err)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	req := &CreateChargeRequest{}
	err = json.Unmarshal(test.GetCreateChargePixRequestDefault(), req)
	AssertFatalErrorNonnull(t, err)
	nCharge := NewCharge(SANDBOX, *accessToken)
	resp, errAsaas := nCharge.Create(ctx, *req)
	AssertAsaasResponseSuccess(t, resp, errAsaas)
}

func TestChangeCreateBoletoSuccess(t *testing.T) {
	accessToken, err := test.GetAccessTokenByEnv()
	AssertFatalErrorNonnull(t, err)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	req := &CreateChargeRequest{}
	err = json.Unmarshal(test.GetCreateChargeBoletoRequestDefault(), req)
	AssertFatalErrorNonnull(t, err)
	nCharge := NewCharge(SANDBOX, *accessToken)
	resp, errAsaas := nCharge.Create(ctx, *req)
	AssertAsaasResponseSuccess(t, resp, errAsaas)
}

func TestChargeUploadDocumentByID(t *testing.T) {
	accessToken, err := test.GetAccessTokenByEnv()
	AssertFatalErrorNonnull(t, err)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	f, err := test.GetSimpleFile()
	AssertFatalErrorNonnull(t, err)
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
	AssertAsaasResponseSuccess(t, resp, errAsaas)
}

func TestChargeGetDocumentByIDSuccess(t *testing.T) {
	accessToken, err := test.GetAccessTokenByEnv()
	AssertFatalErrorNonnull(t, err)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nCharge := NewCharge(SANDBOX, *accessToken)
	resp, errAsaas := nCharge.GetDocumentByID(ctx, test.GetChargeIdDefault(), test.GetDocumentIdDefault())
	AssertAsaasResponseSuccess(t, resp, errAsaas)
}

func TestChargeGetDocumentByIDNoContent(t *testing.T) {
	accessToken, err := test.GetAccessTokenByEnv()
	AssertFatalErrorNonnull(t, err)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nCharge := NewCharge(SANDBOX, *accessToken)
	resp, errAsaas := nCharge.GetDocumentByID(ctx, test.GetChargeIdDefault(), "test")
	AssertAsaasResponseNoContent(t, resp, errAsaas)
}

func TestChargeGetAllDocumentsByIDSuccess(t *testing.T) {
	accessToken, err := test.GetAccessTokenByEnv()
	AssertFatalErrorNonnull(t, err)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nCharge := NewCharge(SANDBOX, *accessToken)
	resp, errAsaas := nCharge.GetAllDocumentsByID(ctx, test.GetChargeIdDefault(), PageableDefaultRequest{})
	AssertAsaasResponseSuccess(t, resp, errAsaas)
}

func TestChargeGetAllDocumentsByIDNoContent(t *testing.T) {
	accessToken, err := test.GetAccessTokenByEnv()
	AssertFatalErrorNonnull(t, err)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nCharge := NewCharge(SANDBOX, *accessToken)
	resp, errAsaas := nCharge.GetAllDocumentsByID(ctx, test.GetChargeIdDefault(), PageableDefaultRequest{})
	AssertAsaasResponseSuccess(t, resp, errAsaas)
}

func TestChargeGetAllSuccess(t *testing.T) {
	accessToken, err := test.GetAccessTokenByEnv()
	AssertFatalErrorNonnull(t, err)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nCharge := NewCharge(SANDBOX, *accessToken)
	resp, errAsaas := nCharge.GetAll(ctx, GetAllChargesRequest{})
	AssertAsaasResponseSuccess(t, resp, errAsaas)
}

func TestChargeGetAllNoContent(t *testing.T) {
	accessToken, err := test.GetAccessTokenByEnv()
	AssertFatalErrorNonnull(t, err)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nCharge := NewCharge(SANDBOX, *accessToken)
	resp, errAsaas := nCharge.GetAll(ctx, GetAllChargesRequest{
		Status: CHARGE_RECEIVED,
	})
	AssertAsaasResponseNoContent(t, resp, errAsaas)
}
