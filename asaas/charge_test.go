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
	assertFatalErrorNonnull(t, err)
	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()
	nCharge := NewCharge(SANDBOX, *accessToken)
	_, errAsaas := nCharge.Create(ctx, CreateChargeRequest{})
	assertSuccessNonnull(t, errAsaas)
}

func TestChangeCreateCreditCardSuccess(t *testing.T) {
	accessToken, err := test.GetAccessTokenByEnv()
	assertFatalErrorNonnull(t, err)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	req := &CreateChargeRequest{}
	err = json.Unmarshal(test.GetCreateChargeCreditCardRequestDefault(), req)
	assertFatalErrorNonnull(t, err)
	nCharge := NewCharge(SANDBOX, *accessToken)
	resp, errAsaas := nCharge.Create(ctx, *req)
	assertResponseSuccess(t, resp, errAsaas)
}

func TestChangeCreateCreditCardFailure(t *testing.T) {
	accessToken, err := test.GetAccessTokenByEnv()
	assertFatalErrorNonnull(t, err)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	req := &CreateChargeRequest{}
	err = json.Unmarshal(test.GetCreateChargeCreditCardFailureRequestDefault(), req)
	assertFatalErrorNonnull(t, err)
	nCharge := NewCharge(SANDBOX, *accessToken)
	resp, errAsaas := nCharge.Create(ctx, *req)
	assertResponseFailure(t, resp, errAsaas)
}

func TestChangeCreatePixSuccess(t *testing.T) {
	accessToken, err := test.GetAccessTokenByEnv()
	assertFatalErrorNonnull(t, err)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	req := &CreateChargeRequest{}
	err = json.Unmarshal(test.GetCreateChargePixRequestDefault(), req)
	assertFatalErrorNonnull(t, err)
	nCharge := NewCharge(SANDBOX, *accessToken)
	resp, errAsaas := nCharge.Create(ctx, *req)
	assertResponseSuccess(t, resp, errAsaas)
}

func TestChangeCreateBoletoSuccess(t *testing.T) {
	accessToken, err := test.GetAccessTokenByEnv()
	assertFatalErrorNonnull(t, err)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	req := &CreateChargeRequest{}
	err = json.Unmarshal(test.GetCreateChargeBoletoRequestDefault(), req)
	assertFatalErrorNonnull(t, err)
	nCharge := NewCharge(SANDBOX, *accessToken)
	resp, errAsaas := nCharge.Create(ctx, *req)
	assertResponseSuccess(t, resp, errAsaas)
}

func TestChargeUploadDocumentByID(t *testing.T) {
	accessToken, err := test.GetAccessTokenByEnv()
	assertFatalErrorNonnull(t, err)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	f, err := test.GetSimpleFile()
	assertFatalErrorNonnull(t, err)
	defer func(name string) {
		err = os.Remove(name)
		if err != nil {
			logError("error remove file test:", err)
		}
	}(f.Name())
	nCharge := NewCharge(SANDBOX, *accessToken)
	resp, errAsaas := nCharge.UploadDocumentByID(ctx, test.GetChargeIdDefault(), UploadChargeDocumentRequest{
		AvailableAfterPayment: false,
		Type:                  DOCUMENT,
		File:                  f,
	})
	assertResponseSuccess(t, resp, errAsaas)
}

func TestChargeGetDocumentByIDSuccess(t *testing.T) {
	accessToken, err := test.GetAccessTokenByEnv()
	assertFatalErrorNonnull(t, err)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nCharge := NewCharge(SANDBOX, *accessToken)
	resp, errAsaas := nCharge.GetDocumentByID(ctx, test.GetChargeIdDefault(), test.GetDocumentIdDefault())
	assertResponseSuccess(t, resp, errAsaas)
}

func TestChargeGetDocumentByIDNoContent(t *testing.T) {
	accessToken, err := test.GetAccessTokenByEnv()
	assertFatalErrorNonnull(t, err)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nCharge := NewCharge(SANDBOX, *accessToken)
	resp, errAsaas := nCharge.GetDocumentByID(ctx, test.GetChargeIdDefault(), "test")
	assertResponseNoContent(t, resp, errAsaas)
}

func TestChargeGetCreationLimit(t *testing.T) {
	accessToken, err := test.GetAccessTokenByEnv()
	assertFatalErrorNonnull(t, err)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nCharge := NewCharge(SANDBOX, *accessToken)
	resp, errAsaas := nCharge.GetCreationLimit(ctx)
	assertResponseSuccess(t, resp, errAsaas)
}

func TestChargeGetAllDocumentsByIDSuccess(t *testing.T) {
	accessToken, err := test.GetAccessTokenByEnv()
	assertFatalErrorNonnull(t, err)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nCharge := NewCharge(SANDBOX, *accessToken)
	resp, errAsaas := nCharge.GetAllDocumentsByID(ctx, test.GetChargeIdDefault(), PageableDefaultRequest{})
	assertResponseSuccess(t, resp, errAsaas)
}

func TestChargeGetAllDocumentsByIDNoContent(t *testing.T) {
	accessToken, err := test.GetAccessTokenByEnv()
	assertFatalErrorNonnull(t, err)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nCharge := NewCharge(SANDBOX, *accessToken)
	resp, errAsaas := nCharge.GetAllDocumentsByID(ctx, test.GetChargeIdDefault(), PageableDefaultRequest{})
	assertResponseSuccess(t, resp, errAsaas)
}

func TestChargeGetAllSuccess(t *testing.T) {
	accessToken, err := test.GetAccessTokenByEnv()
	assertFatalErrorNonnull(t, err)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nCharge := NewCharge(SANDBOX, *accessToken)
	resp, errAsaas := nCharge.GetAll(ctx, GetAllChargesRequest{})
	assertResponseSuccess(t, resp, errAsaas)
}

func TestChargeGetAllNoContent(t *testing.T) {
	accessToken, err := test.GetAccessTokenByEnv()
	assertFatalErrorNonnull(t, err)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nCharge := NewCharge(SANDBOX, *accessToken)
	resp, errAsaas := nCharge.GetAll(ctx, GetAllChargesRequest{
		Status: CHARGE_RECEIVED,
	})
	assertResponseNoContent(t, resp, errAsaas)
}
