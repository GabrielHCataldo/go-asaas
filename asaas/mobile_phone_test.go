package asaas

import (
	"context"
	"github.com/GabrielHCataldo/go-asaas/internal/test"
	"testing"
	"time"
)

func TestMobilePhoneRecharge(t *testing.T) {
	accessToken, err := test.GetAccessTokenByEnv()
	assertFatalErrorNonnull(t, err)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nMobilePhone := NewMobilePhone(EnvSandbox, accessToken)
	resp, errAsaas := nMobilePhone.Recharge(ctx, MobilePhoneRechargeRequest{
		PhoneNumber: "47997576130",
		Value:       15,
	})
	assertResponseSuccess(t, resp, errAsaas)
}
