package asaas

import (
	"context"
	"testing"
	"time"
)

func TestMobilePhoneRecharge(t *testing.T) {
	accessToken := getEnvValue(EnvAccessToken)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nMobilePhone := NewMobilePhone(EnvSandbox, accessToken)
	resp, err := nMobilePhone.Recharge(ctx, MobilePhoneRechargeRequest{
		PhoneNumber: "47997576131",
		Value:       20,
	})
	assertResponseSuccess(t, resp, err)
}

func TestMobilePhoneCancelRechargeById(t *testing.T) {
	initMobilePhoneRecharge()
	accessToken := getEnvValue(EnvAccessToken)
	rechargeId := getEnvValue(EnvMobilePhoneRechargeId)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nMobilePhone := NewMobilePhone(EnvSandbox, accessToken)
	resp, err := nMobilePhone.CancelRechargeById(ctx, rechargeId)
	assertResponseSuccess(t, resp, err)
}

func TestMobilePhoneGetRechargeById(t *testing.T) {
	initMobilePhoneRecharge()
	accessToken := getEnvValue(EnvAccessToken)
	rechargeId := getEnvValue(EnvMobilePhoneRechargeId)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nMobilePhone := NewMobilePhone(EnvSandbox, accessToken)
	resp, err := nMobilePhone.GetRechargeById(ctx, rechargeId)
	assertResponseSuccess(t, resp, err)
}

func TestMobilePhoneGetAllRecharges(t *testing.T) {
	initMobilePhoneRecharge()
	accessToken := getEnvValue(EnvAccessToken)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nMobilePhone := NewMobilePhone(EnvSandbox, accessToken)
	resp, err := nMobilePhone.GetAllRecharges(ctx, PageableDefaultRequest{
		Offset: 0,
		Limit:  10,
	})
	assertResponseSuccess(t, resp, err)
}

func TestMobilePhoneGetProviderByPhoneNumber(t *testing.T) {
	accessToken := getEnvValue(EnvAccessToken)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nMobilePhone := NewMobilePhone(EnvSandbox, accessToken)
	resp, err := nMobilePhone.GetProviderByPhoneNumber(ctx, "47997576131")
	assertResponseSuccess(t, resp, err)
}
