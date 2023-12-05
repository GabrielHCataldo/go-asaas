package asaas

import (
	"context"
	"testing"
	"time"
)

func TestInstallmentUpdateSplitsById(t *testing.T) {
	initBankSlipCharge(true)
	accessToken := getEnvValue(EnvAccessToken)
	installmentId := getEnvValue(EnvChargeInstallmentId)
	walletIdSecondary := getEnvValue(EnvWalletIdSecondary)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nInstallment := NewInstallment(EnvSandbox, accessToken)
	resp, err := nInstallment.UpdateSplitsById(ctx, installmentId, UpdateSplitsRequest{
		Splits: []SplitRequest{
			{
				WalletId:        walletIdSecondary,
				FixedValue:      0,
				PercentualValue: 50,
				TotalFixedValue: 0,
			},
		},
	})
	assertResponseSuccess(t, resp, err)
}

func TestInstallmentDeleteById(t *testing.T) {
	initBankSlipCharge(true)
	accessToken := getEnvValue(EnvAccessToken)
	installmentId := getEnvValue(EnvChargeInstallmentId)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nInstallment := NewInstallment(EnvSandbox, accessToken)
	resp, err := nInstallment.DeleteById(ctx, installmentId)
	assertResponseSuccess(t, resp, err)
}

func TestInstallmentRefundById(t *testing.T) {
	initCreditCardCharge(true, false)
	accessToken := getEnvValue(EnvAccessToken)
	installmentId := getEnvValue(EnvChargeInstallmentId)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nInstallment := NewInstallment(EnvSandbox, accessToken)
	resp, err := nInstallment.RefundById(ctx, installmentId)
	assertResponseSuccess(t, resp, err)
}

func TestInstallmentGetById(t *testing.T) {
	initBankSlipCharge(true)
	accessToken := getEnvValue(EnvAccessToken)
	installmentId := getEnvValue(EnvChargeInstallmentId)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nInstallment := NewInstallment(EnvSandbox, accessToken)
	resp, err := nInstallment.GetById(ctx, installmentId)
	assertResponseSuccess(t, resp, err)
}

func TestInstallmentGetPaymentBookById(t *testing.T) {
	initBankSlipCharge(true)
	accessToken := getEnvValue(EnvAccessToken)
	installmentId := getEnvValue(EnvChargeInstallmentId)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nInstallment := NewInstallment(EnvSandbox, accessToken)
	resp, err := nInstallment.GetPaymentBookById(ctx, installmentId, InstallmentPaymentBookRequest{
		Sort:  "",
		Order: "",
	})
	assertResponseFailure(resp, err)
}

func TestInstallmentGetAll(t *testing.T) {
	initBankSlipCharge(true)
	accessToken := getEnvValue(EnvAccessToken)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nInstallment := NewInstallment(EnvSandbox, accessToken)
	resp, err := nInstallment.GetAll(ctx, PageableDefaultRequest{})
	assertResponseSuccess(t, resp, err)
}
