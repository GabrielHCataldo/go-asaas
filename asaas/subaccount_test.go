package asaas

import (
	"context"
	"github.com/mvrilo/go-cpf"
	"os"
	"testing"
	"time"
)

func TestSubaccountCreate(t *testing.T) {
	accessToken := getEnvValue(EnvAccessTokenSecondary)
	assertFatalStringBlank(t, accessToken)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nSubaccount := NewSubaccount(EnvSandbox, accessToken)
	resp, err := nSubaccount.Create(ctx, CreateSubaccountRequest{
		Name:          "Unit test go",
		Email:         "unittestgo@gmail.com",
		LoginEmail:    "",
		CpfCnpj:       cpf.Generate(),
		BirthDate:     NewDate(1999, 1, 21, time.Local),
		CompanyType:   "",
		Phone:         "",
		MobilePhone:   "47997576131",
		Site:          "",
		Address:       "Rua Maria de Souza Maba",
		AddressNumber: "123",
		Complement:    "",
		Province:      "Fortaleza",
		PostalCode:    "89056-220",
		Webhooks:      nil,
	})
	assertResponseSuccess(t, resp, err)
}

func TestSubaccountSendWhiteLabelDocument(t *testing.T) {
	accessToken := getEnvValue(EnvAccessTokenSecondary)
	assertFatalStringBlank(t, accessToken)
	initSubaccount()
	subaccountId := getEnvValue(EnvSubaccountId)
	assertFatalStringBlank(t, subaccountId)
	f, err := os.Open(getEnvValue(EnvFileName))
	assertFatalErrorNonnull(t, err)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nSubaccount := NewSubaccount(EnvSandbox, accessToken)
	resp, err := nSubaccount.SendWhiteLabelDocument(ctx, subaccountId, SubaccountSendDocumentRequest{
		Type:         SubaccountDocumentTypeCustom,
		DocumentFile: f,
	})
	assertResponseSuccess(t, resp, err)
}

func TestSubaccountUpdateWhiteLabelDocumentSentById(t *testing.T) {
	accessToken := getEnvValue(EnvAccessTokenSecondary)
	assertFatalStringBlank(t, accessToken)
	initSubaccount()
	subaccountId := getEnvValue(EnvSubaccountId)
	assertFatalStringBlank(t, subaccountId)
	f, err := os.Open(getEnvValue(EnvFileName))
	assertFatalErrorNonnull(t, err)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nSubaccount := NewSubaccount(EnvSandbox, accessToken)
	resp, err := nSubaccount.UpdateWhiteLabelDocumentSentById(ctx, subaccountId, UpdateWhiteLabelDocumentSentRequest{
		DocumentFile: f,
	})
	assertResponseSuccess(t, resp, err)
}

func TestSubaccountDeleteWhiteLabelDocumentSentById(t *testing.T) {
	accessToken := getEnvValue(EnvAccessTokenSecondary)
	assertFatalStringBlank(t, accessToken)
	initSubaccountDocument()
	subaccountDocumentSendId := getEnvValue(EnvSubaccountDocumentSentId)
	assertFatalStringBlank(t, subaccountDocumentSendId)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nSubaccount := NewSubaccount(EnvSandbox, accessToken)
	resp, err := nSubaccount.DeleteWhiteLabelDocumentSentById(ctx, subaccountDocumentSendId)
	assertResponseSuccess(t, resp, err)
}

func TestSubaccountGetById(t *testing.T) {
	accessToken := getEnvValue(EnvAccessTokenSecondary)
	assertFatalStringBlank(t, accessToken)
	initSubaccount()
	subaccountId := getEnvValue(EnvSubaccountId)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nSubaccount := NewSubaccount(EnvSandbox, accessToken)
	resp, err := nSubaccount.GetById(ctx, subaccountId)
	assertResponseSuccess(t, resp, err)
}

func TestSubaccountGetDocumentSentById(t *testing.T) {
	accessToken := getEnvValue(EnvAccessTokenSecondary)
	assertFatalStringBlank(t, accessToken)
	initSubaccountDocument()
	subaccountDocumentSendId := getEnvValue(EnvSubaccountDocumentSentId)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nSubaccount := NewSubaccount(EnvSandbox, accessToken)
	resp, err := nSubaccount.GetDocumentSentById(ctx, subaccountDocumentSendId)
	assertResponseSuccess(t, resp, err)
}

func TestSubaccountGetAll(t *testing.T) {
	accessToken := getEnvValue(EnvAccessToken)
	assertFatalStringBlank(t, accessToken)
	initSubaccount()
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nSubaccount := NewSubaccount(EnvSandbox, accessToken)
	resp, errAsaas := nSubaccount.GetAll(ctx, GetAllSubaccountsRequest{
		CpfCnpj:  "",
		Email:    "",
		Name:     "",
		WalletId: "",
		Offset:   0,
		Limit:    10,
	})
	assertResponseSuccess(t, resp, errAsaas)
}

func TestSubaccountGetPendingDocuments(t *testing.T) {
	accessToken := getEnvValue(EnvAccessTokenSecondary)
	assertFatalStringBlank(t, accessToken)
	initSubaccountDocument()
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nSubaccount := NewSubaccount(EnvSandbox, accessToken)
	resp, err := nSubaccount.GetPendingDocuments(ctx)
	assertResponseSuccess(t, resp, err)
}
