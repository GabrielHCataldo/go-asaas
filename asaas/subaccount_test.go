package asaas

import (
	"context"
	"github.com/GabrielHCataldo/go-asaas/internal/util"
	"github.com/mvrilo/go-cpf"
	"os"
	"testing"
	"time"
)

func TestSubaccountCreate(t *testing.T) {
	accessToken := getEnvValue(EnvAccessToken)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nSubaccount := NewSubaccount(EnvSandbox, accessToken)
	resp, err := nSubaccount.Create(ctx, CreateSubaccountRequest{
		Name:          "Unit test go",
		Email:         util.GenerateEmail(),
		LoginEmail:    "",
		CpfCnpj:       cpf.Generate(),
		BirthDate:     NewDate(1999, 6, 12, time.Local),
		CompanyType:   "",
		Phone:         "",
		MobilePhone:   util.GenerateMobilePhone(),
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
	//initSubaccountDocument()
	accessToken := getEnvValue(EnvAccessToken)
	f, _ := os.Open(getEnvValue(EnvImageName))
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nSubaccount := NewSubaccount(EnvSandbox, accessToken)
	resp, err := nSubaccount.SendWhiteLabelDocument(ctx, "test", SendWhiteLabelDocumentRequest{
		Type:         SubaccountDocumentTypeCustom,
		DocumentFile: f,
	})
	assertResponseFailure(resp, err)
}

func TestSubaccountUpdateWhiteLabelDocumentSentById(t *testing.T) {
	//initSubaccount()
	accessToken := getEnvValue(EnvAccessToken)
	f, _ := os.Open(getEnvValue(EnvFileName))
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nSubaccount := NewSubaccount(EnvSandbox, accessToken)
	resp, err := nSubaccount.UpdateWhiteLabelDocumentSentById(ctx, "test", UpdateWhiteLabelDocumentSentRequest{
		DocumentFile: f,
	})
	assertResponseFailure(resp, err)
}

func TestSubaccountDeleteWhiteLabelDocumentSentById(t *testing.T) {
	//initSubaccountDocumentSent()
	accessToken := getEnvValue(EnvAccessToken)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nSubaccount := NewSubaccount(EnvSandbox, accessToken)
	resp, err := nSubaccount.DeleteWhiteLabelDocumentSentById(ctx, "test")
	assertResponseFailure(resp, err)
}

func TestSubaccountGetById(t *testing.T) {
	//initSubaccount()
	accessToken := getEnvValue(EnvAccessToken)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nSubaccount := NewSubaccount(EnvSandbox, accessToken)
	resp, err := nSubaccount.GetById(ctx, "test")
	assertResponseFailure(resp, err)
}

func TestSubaccountGetDocumentSentById(t *testing.T) {
	//initSubaccountDocumentSent()
	accessToken := getEnvValue(EnvAccessTokenSecondary)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nSubaccount := NewSubaccount(EnvSandbox, accessToken)
	resp, err := nSubaccount.GetDocumentSentById(ctx, "test")
	assertResponseFailure(resp, err)
}

func TestSubaccountGetAll(t *testing.T) {
	initSubaccount()
	accessToken := getEnvValue(EnvAccessToken)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nSubaccount := NewSubaccount(EnvSandbox, accessToken)
	resp, err := nSubaccount.GetAll(ctx, GetAllSubaccountsRequest{
		CpfCnpj:  "",
		Email:    "",
		Name:     "",
		WalletId: "",
		Offset:   0,
		Limit:    10,
	})
	assertResponseSuccess(t, resp, err)
}

func TestSubaccountGetPendingDocuments(t *testing.T) {
	accessToken := getEnvValue(EnvAccessToken)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nSubaccount := NewSubaccount(EnvSandbox, accessToken)
	resp, err := nSubaccount.GetPendingDocuments(ctx)
	assertResponseSuccess(t, resp, err)
}
