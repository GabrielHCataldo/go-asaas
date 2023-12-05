package asaas

import (
	"context"
	"os"
	"testing"
	"time"
)

func TestAccountSaveInvoiceCustomization(t *testing.T) {
	accessToken := getEnvValue(EnvAccessTokenSecondary)
	f, _ := os.Open(getEnvValue(EnvImageName))
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nAccount := NewAccount(EnvSandbox, accessToken)
	resp, err := nAccount.SaveInvoiceCustomization(ctx, SaveInvoiceCustomizationRequest{
		LogoBackgroundColor: "#FFFFFF",
		InfoBackgroundColor: "#FF0000",
		FontColor:           "#000000",
		Enabled:             false,
		LogoFile:            f,
	})
	assertResponseSuccess(t, resp, err)
}

func TestAccountUpdate(t *testing.T) {
	accessToken := getEnvValue(EnvAccessTokenSecondary)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nAccount := NewAccount(EnvSandbox, accessToken)
	resp, err := nAccount.Update(ctx, UpdateAccountRequest{
		PersonType:    PersonTypePhysical,
		CpfCnpj:       "02104996643",
		BirthDate:     NewDate(1999, 6, 2, time.Local),
		CompanyType:   nil,
		Email:         "gabrielcataldo20@gmail.com",
		Phone:         nil,
		MobilePhone:   nil,
		Site:          Pointer("https://sitetest.com.br"),
		PostalCode:    "69159-970",
		Address:       nil,
		AddressNumber: nil,
		Complement:    nil,
		Province:      nil,
	})
	assertResponseSuccess(t, resp, err)
}

func TestAccountDeleteWhiteLabelSubaccount(t *testing.T) {
	initSubaccountDocumentSent()
	accessToken := getEnvValue(EnvAccessTokenSecondary)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nAccount := NewAccount(EnvSandbox, accessToken)
	resp, err := nAccount.DeleteWhiteLabelSubaccount(ctx, DeleteWhiteLabelSubaccountRequest{
		RemoveReason: "unit test",
	})
	assertResponseFailure(resp, err)
}

func TestAccountGet(t *testing.T) {
	accessToken := getEnvValue(EnvAccessTokenSecondary)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nAccount := NewAccount(EnvSandbox, accessToken)
	resp, err := nAccount.Get(ctx)
	assertResponseSuccess(t, resp, err)
}

func TestAccountGetRegistrationStatus(t *testing.T) {
	accessToken := getEnvValue(EnvAccessTokenSecondary)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nAccount := NewAccount(EnvSandbox, accessToken)
	resp, err := nAccount.GetRegistrationStatus(ctx)
	assertResponseSuccess(t, resp, err)
}

func TestAccountGetBankInfo(t *testing.T) {
	accessToken := getEnvValue(EnvAccessTokenSecondary)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nAccount := NewAccount(EnvSandbox, accessToken)
	resp, err := nAccount.GetBankInfo(ctx)
	assertResponseSuccess(t, resp, err)
}

func TestAccountGetFees(t *testing.T) {
	accessToken := getEnvValue(EnvAccessTokenSecondary)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nAccount := NewAccount(EnvSandbox, accessToken)
	resp, err := nAccount.GetFees(ctx)
	assertResponseSuccess(t, resp, err)
}

func TestAccountGetWallets(t *testing.T) {
	accessToken := getEnvValue(EnvAccessTokenSecondary)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nAccount := NewAccount(EnvSandbox, accessToken)
	resp, err := nAccount.GetWallets(ctx)
	assertResponseSuccess(t, resp, err)
}

func TestAccountGetBalance(t *testing.T) {
	accessToken := getEnvValue(EnvAccessTokenSecondary)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nAccount := NewAccount(EnvSandbox, accessToken)
	resp, err := nAccount.GetBalance(ctx)
	assertResponseSuccess(t, resp, err)
}

func TestAccountGetAccountStatement(t *testing.T) {
	accessToken := getEnvValue(EnvAccessTokenSecondary)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nAccount := NewAccount(EnvSandbox, accessToken)
	resp, err := nAccount.GetAccountStatement(ctx, GetAccountStatementRequest{
		StartDate:  Date{},
		FinishDate: Date{},
		Offset:     0,
		Limit:      10,
		Order:      "",
	})
	assertResponseSuccess(t, resp, err)
}

func TestAccountGetPaymentStatistic(t *testing.T) {
	accessToken := getEnvValue(EnvAccessTokenSecondary)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nAccount := NewAccount(EnvSandbox, accessToken)
	resp, err := nAccount.GetPaymentStatistic(ctx, GetPaymentStatisticRequest{
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
	assertResponseSuccess(t, resp, err)
}

func TestAccountGetSplitStatistic(t *testing.T) {
	accessToken := getEnvValue(EnvAccessTokenSecondary)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nAccount := NewAccount(EnvSandbox, accessToken)
	resp, err := nAccount.GetSplitStatistic(ctx)
	assertResponseSuccess(t, resp, err)
}

func TestAccountGetInvoiceCustomization(t *testing.T) {
	accessToken := getEnvValue(EnvAccessTokenSecondary)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nAccount := NewAccount(EnvSandbox, accessToken)
	resp, err := nAccount.GetInvoiceCustomization(ctx)
	assertResponseSuccess(t, resp, err)
}
