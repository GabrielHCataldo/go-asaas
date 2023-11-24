package asaas

import (
	"context"
	"encoding/json"
	"github.com/GabrielHCataldo/go-asaas/internal/test"
	"testing"
	"time"
)

func TestChargeCreateError(t *testing.T) {
	accessToken, err := test.GetAccessTokenByEnv()
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()
	nCharge := NewCharge(SANDBOX, *accessToken)
	_, errAsaas := nCharge.Create(ctx, CreateChargeRequest{})
	if errAsaas == nil {
		t.Error("unexpect response: err is nil")
		t.Fail()
	}
	t.Log("result success error msg: ", errAsaas.Msg)
}

func TestChangeCreateCreditCardSuccess(t *testing.T) {
	accessToken, err := test.GetAccessTokenByEnv()
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nCharge := NewCharge(SANDBOX, *accessToken)
	resp, errAsaas := nCharge.Create(ctx, CreateChargeRequest{
		Customer:             getTestCustomerIdDefault(),
		BillingType:          CREDIT_CARD,
		Value:                10,
		DueDate:              Now(),
		Description:          getTestChargeDescriptionDefault(),
		CreditCard:           getTestCreditCardRequestDefault(),
		CreditCardHolderInfo: getTestCreditCardHolderInfoRequestDefault(),
		RemoteIP:             getTestRemoteIPDefault(),
	})
	if errAsaas != nil {
		t.Error("unexpect response: err is not nil:", errAsaas)
		t.Fail()
	} else if len(resp.Errors) > 0 {
		t.Error("unexpect response: errors response from Asaas:", resp.Errors)
		t.Fail()
	} else {
		respJson, _ := json.Marshal(resp)
		t.Log("result success response: ", string(respJson))
	}
}
