package asaas

import (
	"context"
	"os"
	"testing"
	"time"
)

func TestAccountSaveInvoiceCustomization(t *testing.T) {
	accessToken := getEnvValue(EnvAccessTokenSecondary)
	assertFatalStringBlank(t, accessToken)
	f, err := os.Open(getEnvValue(EnvImageName))
	assertFatalErrorNonnull(t, err)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nAccount := NewAccount(EnvSandbox, accessToken)
	resp, errAsaas := nAccount.SaveInvoiceCustomization(ctx, SaveInvoiceCustomizationRequest{
		LogoBackgroundColor: "#FFFFFF",
		InfoBackgroundColor: "#FF0000",
		FontColor:           "#000000",
		Enabled:             false,
		LogoFile:            f,
	})
	assertResponseSuccess(t, resp, errAsaas)
}

func TestAccountUpdate(t *testing.T) {
	accessToken := getEnvValue(EnvAccessTokenSecondary)
	assertFatalStringBlank(t, accessToken)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nAccount := NewAccount(EnvSandbox, accessToken)
	resp, errAsaas := nAccount.Update(ctx, UpdateAccountRequest{
		PersonType:    "",
		CpfCnpj:       "",
		BirthDate:     Date{},
		CompanyType:   nil,
		Email:         nil,
		Phone:         nil,
		MobilePhone:   nil,
		Site:          Pointer("https://sitetest.com.br"),
		PostalCode:    "",
		Address:       nil,
		AddressNumber: nil,
		Complement:    nil,
		Province:      nil,
	})
	assertResponseSuccess(t, resp, errAsaas)
}

func TestAccountDeleteWhiteLabelSubaccount(t *testing.T) {
	accessToken := getEnvValue(EnvAccessTokenSecondary)
	assertFatalStringBlank(t, accessToken)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nAccount := NewAccount(EnvSandbox, accessToken)
	resp, errAsaas := nAccount.DeleteWhiteLabelSubaccount(ctx, DeleteWhiteLabelSubaccountRequest{
		RemoveReason: "unit test",
	})
	assertResponseSuccess(t, resp, errAsaas)
}

func TestAccountGet(t *testing.T) {
	accessToken := getEnvValue(EnvAccessTokenSecondary)
	assertFatalStringBlank(t, accessToken)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nAccount := NewAccount(EnvSandbox, accessToken)
	resp, errAsaas := nAccount.Get(ctx)
	assertResponseSuccess(t, resp, errAsaas)
}

func TestAccountGetRegistrationStatus(t *testing.T) {
	accessToken := getEnvValue(EnvAccessTokenSecondary)
	assertFatalStringBlank(t, accessToken)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nAccount := NewAccount(EnvSandbox, accessToken)
	resp, errAsaas := nAccount.GetRegistrationStatus(ctx)
	assertResponseSuccess(t, resp, errAsaas)
}

func TestAccountGetBankInfo(t *testing.T) {
	accessToken := getEnvValue(EnvAccessTokenSecondary)
	assertFatalStringBlank(t, accessToken)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nAccount := NewAccount(EnvSandbox, accessToken)
	resp, errAsaas := nAccount.GetBankInfo(ctx)
	assertResponseSuccess(t, resp, errAsaas)
}

func TestAccountGetFees(t *testing.T) {
	accessToken := getEnvValue(EnvAccessTokenSecondary)
	assertFatalStringBlank(t, accessToken)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nAccount := NewAccount(EnvSandbox, accessToken)
	resp, errAsaas := nAccount.GetFees(ctx)
	assertResponseSuccess(t, resp, errAsaas)
}

func TestAccountGetWallets(t *testing.T) {
	accessToken := getEnvValue(EnvAccessTokenSecondary)
	assertFatalStringBlank(t, accessToken)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nAccount := NewAccount(EnvSandbox, accessToken)
	resp, errAsaas := nAccount.GetWallets(ctx)
	assertResponseSuccess(t, resp, errAsaas)
}

func TestAccountGetBalance(t *testing.T) {
	accessToken := getEnvValue(EnvAccessTokenSecondary)
	assertFatalStringBlank(t, accessToken)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nAccount := NewAccount(EnvSandbox, accessToken)
	resp, errAsaas := nAccount.GetBalance(ctx)
	assertResponseSuccess(t, resp, errAsaas)
}

func TestAccountGetAccountStatement(t *testing.T) {
	accessToken := getEnvValue(EnvAccessTokenSecondary)
	assertFatalStringBlank(t, accessToken)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nAccount := NewAccount(EnvSandbox, accessToken)
	resp, errAsaas := nAccount.GetAccountStatement(ctx, GetAccountStatementRequest{
		StartDate:  Date{},
		FinishDate: Date{},
		Offset:     0,
		Limit:      10,
		Order:      "",
	})
	assertResponseSuccess(t, resp, errAsaas)
}

func TestAccountGetPaymentStatistic(t *testing.T) {
	accessToken := getEnvValue(EnvAccessTokenSecondary)
	assertFatalStringBlank(t, accessToken)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nAccount := NewAccount(EnvSandbox, accessToken)
	resp, errAsaas := nAccount.GetPaymentStatistic(ctx, GetPaymentStatisticRequest{
		Customer:              "",
		BillingType:           "",
		Status:                "",
		Anticipated:           nil,
		DueDateGe:             Date{},
		DueDateLe:             Date{},
		DateCreatedGe:         Date{},
		DateCreatedLe:         Date{},
		EstimatedCreditDateGe: Date{},
		EstimatedCreditDateLe: Date{},
		ExternalReference:     "",
	})
	assertResponseSuccess(t, resp, errAsaas)
}

func TestAccountGetSplitStatistic(t *testing.T) {
	accessToken := getEnvValue(EnvAccessTokenSecondary)
	assertFatalStringBlank(t, accessToken)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nAccount := NewAccount(EnvSandbox, accessToken)
	resp, errAsaas := nAccount.GetSplitStatistic(ctx)
	assertResponseSuccess(t, resp, errAsaas)
}

func TestAccountGetInvoiceCustomization(t *testing.T) {
	accessToken := getEnvValue(EnvAccessTokenSecondary)
	assertFatalStringBlank(t, accessToken)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nAccount := NewAccount(EnvSandbox, accessToken)
	resp, errAsaas := nAccount.GetInvoiceCustomization(ctx)
	assertResponseSuccess(t, resp, errAsaas)
}
