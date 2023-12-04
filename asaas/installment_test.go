package asaas

import (
	"context"
	"testing"
	"time"
)

func TestInstallmentUpdateSplitsById(t *testing.T) {
	accessToken := getEnvValue(EnvAccessToken)
	assertFatalStringBlank(t, accessToken)
	initBankSlipCharge(true)
	installmentId := getEnvValue(EnvChargeInstallmentId)
	assertFatalStringBlank(t, installmentId)
	walletIdSecondary := getEnvValue(EnvWalletIdSecondary)
	assertFatalStringBlank(t, walletIdSecondary)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nInstallment := NewInstallment(EnvSandbox, accessToken)
	resp, err := nInstallment.UpdateSplitsById(ctx, installmentId, []SplitRequest{
		{
			WalletId:        walletIdSecondary,
			FixedValue:      0,
			PercentualValue: 50,
			TotalFixedValue: 0,
		},
	})
	assertResponseSuccess(t, resp, err)
}

func TestInstallmentDeleteById(t *testing.T) {
	accessToken := getEnvValue(EnvAccessToken)
	assertFatalStringBlank(t, accessToken)
	initBankSlipCharge(true)
	installmentId := getEnvValue(EnvChargeInstallmentId)
	assertFatalStringBlank(t, installmentId)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nInstallment := NewInstallment(EnvSandbox, accessToken)
	resp, err := nInstallment.DeleteById(ctx, installmentId)
	assertResponseSuccess(t, resp, err)
}

func TestInstallmentRefundById(t *testing.T) {
	accessToken := getEnvValue(EnvAccessToken)
	assertFatalStringBlank(t, accessToken)
	initCreditCardCharge(false, true)
	installmentId := getEnvValue(EnvChargeInstallmentId)
	assertFatalStringBlank(t, installmentId)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nInstallment := NewInstallment(EnvSandbox, accessToken)
	resp, err := nInstallment.RefundById(ctx, installmentId)
	assertResponseSuccess(t, resp, err)
}

func TestInstallmentGetById(t *testing.T) {
	accessToken := getEnvValue(EnvAccessToken)
	assertFatalStringBlank(t, accessToken)
	initBankSlipCharge(true)
	installmentId := getEnvValue(EnvChargeInstallmentId)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nInstallment := NewInstallment(EnvSandbox, accessToken)
	resp, err := nInstallment.GetById(ctx, installmentId)
	assertResponseSuccess(t, resp, err)
}

func TestInstallmentGetPaymentBookById(t *testing.T) {
	accessToken := getEnvValue(EnvAccessToken)
	assertFatalStringBlank(t, accessToken)
	initBankSlipCharge(true)
	installmentId := getEnvValue(EnvChargeInstallmentId)
	assertFatalStringBlank(t, installmentId)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nInstallment := NewInstallment(EnvSandbox, accessToken)
	resp, err := nInstallment.GetPaymentBookById(ctx, installmentId, InstallmentPaymentBookRequest{
		Sort:  "",
		Order: "",
	})
	assertResponseFailure(t, resp, err)
}

func TestInstallmentGetAll(t *testing.T) {
	accessToken := getEnvValue(EnvAccessToken)
	assertFatalStringBlank(t, accessToken)
	initBankSlipCharge(true)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nInstallment := NewInstallment(EnvSandbox, accessToken)
	resp, errAsaas := nInstallment.GetAll(ctx, PageableDefaultRequest{})
	assertResponseSuccess(t, resp, errAsaas)
}
