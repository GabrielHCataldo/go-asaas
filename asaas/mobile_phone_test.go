package asaas

import (
	"context"
	"testing"
	"time"
)

func TestMobilePhoneRecharge(t *testing.T) {
	accessToken := getEnvValue(EnvAccessToken)
	assertFatalIsBlank(t, accessToken)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nMobilePhone := NewMobilePhone(EnvSandbox, accessToken)
	resp, errAsaas := nMobilePhone.Recharge(ctx, MobilePhoneRechargeRequest{
		PhoneNumber: "47997576130",
		Value:       15,
	})
	assertResponseSuccess(t, resp, errAsaas)
}
