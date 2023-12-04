package asaas

import (
	"context"
	"testing"
	"time"
)

func TestMobilePhoneRecharge(t *testing.T) {
	accessToken := getEnvValue(EnvAccessToken)
	assertFatalStringBlank(t, accessToken)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nMobilePhone := NewMobilePhone(EnvSandbox, accessToken)
	resp, errAsaas := nMobilePhone.Recharge(ctx, MobilePhoneRechargeRequest{
		PhoneNumber: "47997576131",
		Value:       20,
	})
	assertResponseSuccess(t, resp, errAsaas)
}

func TestMobilePhoneCancelRechargeById(t *testing.T) {
	accessToken := getEnvValue(EnvAccessToken)
	assertFatalStringBlank(t, accessToken)
	initMobilePhoneRecharge()
	rechargeId := getEnvValue(EnvMobilePhoneRechargeId)
	assertFatalStringBlank(t, rechargeId)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nMobilePhone := NewMobilePhone(EnvSandbox, accessToken)
	resp, errAsaas := nMobilePhone.CancelRechargeById(ctx, rechargeId)
	assertResponseSuccess(t, resp, errAsaas)
}

func TestMobilePhoneGetRechargeById(t *testing.T) {
	accessToken := getEnvValue(EnvAccessToken)
	assertFatalStringBlank(t, accessToken)
	initMobilePhoneRecharge()
	rechargeId := getEnvValue(EnvMobilePhoneRechargeId)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nMobilePhone := NewMobilePhone(EnvSandbox, accessToken)
	resp, errAsaas := nMobilePhone.GetRechargeById(ctx, rechargeId)
	assertResponseSuccess(t, resp, errAsaas)
}

func TestMobilePhoneGetAllRecharges(t *testing.T) {
	accessToken := getEnvValue(EnvAccessToken)
	assertFatalStringBlank(t, accessToken)
	initMobilePhoneRecharge()
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nMobilePhone := NewMobilePhone(EnvSandbox, accessToken)
	resp, errAsaas := nMobilePhone.GetAllRecharges(ctx, PageableDefaultRequest{
		Offset: 0,
		Limit:  10,
	})
	assertResponseSuccess(t, resp, errAsaas)
}

func TestMobilePhoneGetProviderByPhoneNumber(t *testing.T) {
	accessToken := getEnvValue(EnvAccessToken)
	assertFatalStringBlank(t, accessToken)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nMobilePhone := NewMobilePhone(EnvSandbox, accessToken)
	resp, errAsaas := nMobilePhone.GetProviderByPhoneNumber(ctx, "47997576131")
	assertResponseSuccess(t, resp, errAsaas)
}
