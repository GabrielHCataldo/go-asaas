package asaas

import (
	"context"
	"github.com/GabrielHCataldo/go-asaas/internal/test"
	"testing"
	"time"
)

func TestAccountSaveInvoiceCustomizationSuccess(t *testing.T) {
	accessToken, err := getAccessToken()
	assertFatalErrorNonnull(t, err)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nAccount := NewAccount(EnvSandbox, accessToken)
	f, err := test.GetSimpleImage()
	assertFatalErrorNonnull(t, err)
	resp, errAsaas := nAccount.SaveInvoiceCustomization(ctx, SaveInvoiceCustomizationRequest{
		LogoBackgroundColor: "#FFFFFF",
		InfoBackgroundColor: "#FF0000",
		FontColor:           "#000000",
		Enabled:             false,
		LogoFile:            f,
	})
	assertResponseSuccess(t, resp, errAsaas)
}

func TestAccountSaveInvoiceCustomizationError(t *testing.T) {
	nAccount := NewAccount(EnvSandbox, "")
	_, errAsaas := nAccount.SaveInvoiceCustomization(context.TODO(), SaveInvoiceCustomizationRequest{
		LogoBackgroundColor: "",
		InfoBackgroundColor: "",
		FontColor:           "",
		Enabled:             false,
		LogoFile:            nil,
	})
	assertSuccessNonnull(t, errAsaas)
}

func TestAccountUpdateSuccess(t *testing.T) {
	accessToken, err := getAccessToken()
	assertFatalErrorNonnull(t, err)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nAccount := NewAccount(EnvSandbox, accessToken)
	resp, errAsaas := nAccount.Update(ctx, UpdateAccountRequest{
		PersonType:    "",
		CpfCnpj:       "",
		BirthDate:     Date{},
		CompanyType:   "",
		Email:         "",
		Phone:         "",
		MobilePhone:   "",
		Site:          "https://sitetest.com.br",
		PostalCode:    "",
		Address:       "",
		AddressNumber: "",
		Complement:    "",
		Province:      "",
	})
	assertResponseSuccess(t, resp, errAsaas)
}

func TestAccountUpdateError(t *testing.T) {
	nAccount := NewAccount(EnvSandbox, "")
	_, errAsaas := nAccount.Update(context.TODO(), UpdateAccountRequest{
		PersonType:    "",
		CpfCnpj:       "",
		BirthDate:     Date{},
		CompanyType:   "",
		Email:         "",
		Phone:         "",
		MobilePhone:   "",
		Site:          "sitetestcom.br",
		PostalCode:    "",
		Address:       "",
		AddressNumber: "",
		Complement:    "",
		Province:      "",
	})
	assertSuccessNonnull(t, errAsaas)
}

func TestAccountDeleteWhiteLabelSubaccountSuccess(t *testing.T) {
	accessToken, err := getAccessToken()
	assertFatalErrorNonnull(t, err)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nAccount := NewAccount(EnvSandbox, accessToken)
	resp, errAsaas := nAccount.DeleteWhiteLabelSubaccount(ctx, DeleteWhiteLabelSubaccountRequest{
		RemoveReason: "unit test",
	})
	assertResponseSuccess(t, resp, errAsaas)
}

func TestAccountDeleteWhiteLabelSubaccountError(t *testing.T) {
	nAccount := NewAccount(EnvSandbox, "")
	_, errAsaas := nAccount.DeleteWhiteLabelSubaccount(context.TODO(), DeleteWhiteLabelSubaccountRequest{
		RemoveReason: "",
	})
	assertSuccessNonnull(t, errAsaas)
}

func TestAccountGetSuccess(t *testing.T) {
	accessToken, err := getAccessToken()
	assertFatalErrorNonnull(t, err)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nAccount := NewAccount(EnvSandbox, accessToken)
	resp, errAsaas := nAccount.Get(ctx)
	assertResponseSuccess(t, resp, errAsaas)
}

func TestAccountGetError(t *testing.T) {
	nAccount := NewAccount(EnvSandbox, "")
	resp, errAsaas := nAccount.Get(context.TODO())
	assertResponseFailure(t, resp, errAsaas)
}

func TestAccountGetRegistrationStatusSuccess(t *testing.T) {
	accessToken, err := getAccessToken()
	assertFatalErrorNonnull(t, err)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nAccount := NewAccount(EnvSandbox, accessToken)
	resp, errAsaas := nAccount.GetRegistrationStatus(ctx)
	assertResponseSuccess(t, resp, errAsaas)
}

func TestAccountGetRegistrationStatusError(t *testing.T) {
	nAccount := NewAccount(EnvSandbox, "")
	resp, errAsaas := nAccount.Get(context.TODO())
	assertResponseFailure(t, resp, errAsaas)
}

func TestAccountGetBankInfoSuccess(t *testing.T) {
	accessToken, err := getAccessToken()
	assertFatalErrorNonnull(t, err)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nAccount := NewAccount(EnvSandbox, accessToken)
	resp, errAsaas := nAccount.GetBankInfo(ctx)
	assertResponseSuccess(t, resp, errAsaas)
}

func TestAccountGetBankInfoError(t *testing.T) {
	nAccount := NewAccount(EnvSandbox, "")
	resp, errAsaas := nAccount.GetBankInfo(context.TODO())
	assertResponseFailure(t, resp, errAsaas)
}

func TestAccountGetFeesSuccess(t *testing.T) {
	accessToken, err := getAccessToken()
	assertFatalErrorNonnull(t, err)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nAccount := NewAccount(EnvSandbox, accessToken)
	resp, errAsaas := nAccount.GetFees(ctx)
	assertResponseSuccess(t, resp, errAsaas)
}

func TestAccountGetFeesError(t *testing.T) {
	nAccount := NewAccount(EnvSandbox, "")
	resp, errAsaas := nAccount.GetFees(context.TODO())
	assertResponseFailure(t, resp, errAsaas)
}

func TestAccountGetWalletsSuccess(t *testing.T) {
	accessToken, err := getAccessToken()
	assertFatalErrorNonnull(t, err)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nAccount := NewAccount(EnvSandbox, accessToken)
	resp, errAsaas := nAccount.GetWallets(ctx)
	assertResponseSuccess(t, resp, errAsaas)
}

func TestAccountGetWalletsError(t *testing.T) {
	nAccount := NewAccount(EnvSandbox, "")
	resp, errAsaas := nAccount.GetWallets(context.TODO())
	assertResponseFailure(t, resp, errAsaas)
}

func TestAccountGetBalanceSuccess(t *testing.T) {
	accessToken, err := getAccessToken()
	assertFatalErrorNonnull(t, err)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nAccount := NewAccount(EnvSandbox, accessToken)
	resp, errAsaas := nAccount.GetBalance(ctx)
	assertResponseSuccess(t, resp, errAsaas)
}

func TestAccountGetBalanceError(t *testing.T) {
	nAccount := NewAccount(EnvSandbox, "")
	resp, errAsaas := nAccount.GetWallets(context.TODO())
	assertResponseFailure(t, resp, errAsaas)
}

func TestAccountGetAccountStatementSuccess(t *testing.T) {
	accessToken, err := getAccessToken()
	assertFatalErrorNonnull(t, err)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nAccount := NewAccount(EnvSandbox, accessToken)
	resp, errAsaas := nAccount.GetAccountStatement(ctx, GetAccountStatementRequest{
		StartDate:  nil,
		FinishDate: nil,
		Offset:     0,
		Limit:      10,
		Order:      "",
	})
	assertResponseSuccess(t, resp, errAsaas)
}

func TestAccountGetAccountStatementError(t *testing.T) {
	nAccount := NewAccount(EnvSandbox, "")
	resp, errAsaas := nAccount.GetAccountStatement(context.TODO(), GetAccountStatementRequest{})
	assertResponseFailure(t, resp, errAsaas)
}

func TestAccountGetPaymentStatisticSuccess(t *testing.T) {
	accessToken, err := getAccessToken()
	assertFatalErrorNonnull(t, err)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nAccount := NewAccount(EnvSandbox, accessToken)
	resp, errAsaas := nAccount.GetPaymentStatistic(ctx, GetPaymentStatisticRequest{
		Customer:              "",
		BillingType:           "",
		Status:                "",
		Anticipated:           false,
		DueDateGe:             nil,
		DueDateLe:             nil,
		DateCreatedGe:         nil,
		DateCreatedLe:         nil,
		EstimatedCreditDateGe: nil,
		EstimatedCreditDateLe: nil,
		ExternalReference:     "",
	})
	assertResponseSuccess(t, resp, errAsaas)
}

func TestAccountGetPaymentStatisticError(t *testing.T) {
	nAccount := NewAccount(EnvSandbox, "")
	resp, errAsaas := nAccount.GetPaymentStatistic(context.TODO(), GetPaymentStatisticRequest{})
	assertResponseFailure(t, resp, errAsaas)
}

func TestAccountGetSplitStatisticSuccess(t *testing.T) {
	accessToken, err := getAccessToken()
	assertFatalErrorNonnull(t, err)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nAccount := NewAccount(EnvSandbox, accessToken)
	resp, errAsaas := nAccount.GetSplitStatistic(ctx)
	assertResponseSuccess(t, resp, errAsaas)
}

func TestAccountGetSplitStatisticError(t *testing.T) {
	nAccount := NewAccount(EnvSandbox, "")
	resp, errAsaas := nAccount.GetSplitStatistic(context.TODO())
	assertResponseFailure(t, resp, errAsaas)
}

func TestAccountGetInvoiceCustomizationSuccess(t *testing.T) {
	accessToken, err := getAccessToken()
	assertFatalErrorNonnull(t, err)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nAccount := NewAccount(EnvSandbox, accessToken)
	resp, errAsaas := nAccount.GetInvoiceCustomization(ctx)
	assertResponseSuccess(t, resp, errAsaas)
}

func TestAccountGetInvoiceCustomizationError(t *testing.T) {
	nAccount := NewAccount(EnvSandbox, "")
	resp, errAsaas := nAccount.GetInvoiceCustomization(context.TODO())
	assertResponseFailure(t, resp, errAsaas)
}
