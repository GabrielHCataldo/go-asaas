package asaas

import (
	"context"
	"github.com/GabrielHCataldo/go-asaas/internal/util"
	"os"
	"testing"
	"time"
)

func TestSubaccountCreate(t *testing.T) {
	accessToken := getEnvValue(EnvAccessTokenSecondary)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nSubaccount := NewSubaccount(EnvSandbox, accessToken)
	resp, err := nSubaccount.Create(ctx, CreateSubaccountRequest{
		Name:          "Unit test go",
		Email:         util.GenerateEmail(),
		LoginEmail:    "",
		CpfCnpj:       "81452811000125",
		BirthDate:     NewDate(1999, 6, 12, time.Local),
		CompanyType:   CompanyTypeLimited,
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
	initSubaccountDocument()
	accessToken := getEnvValue(EnvAccessTokenSecondary)
	subaccountDocumentId := getEnvValue(EnvSubaccountDocumentId)
	subaccountDocumentType := getEnvValue(EnvSubaccountDocumentType)
	f, _ := os.Open(getEnvValue(EnvImageName))
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nSubaccount := NewSubaccount(EnvSandbox, accessToken)
	resp, err := nSubaccount.SendWhiteLabelDocument(ctx, subaccountDocumentId, SendWhiteLabelDocumentRequest{
		Type:         SubaccountDocumentType(subaccountDocumentType),
		DocumentFile: f,
	})
	assertResponseSuccess(t, resp, err)
}

func TestSubaccountUpdateWhiteLabelDocumentSentById(t *testing.T) {
	initSubaccount()
	accessToken := getEnvValue(EnvAccessTokenSecondary)
	subaccountDocumentSentId := getEnvValue(EnvSubaccountDocumentSentId)
	f, _ := os.Open(getEnvValue(EnvFileName))
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nSubaccount := NewSubaccount(EnvSandbox, accessToken)
	resp, err := nSubaccount.UpdateWhiteLabelDocumentSentById(ctx, subaccountDocumentSentId,
		UpdateWhiteLabelDocumentSentRequest{
			DocumentFile: f,
		})
	assertResponseSuccess(t, resp, err)
}

func TestSubaccountDeleteWhiteLabelDocumentSentById(t *testing.T) {
	initSubaccountDocumentSent()
	accessToken := getEnvValue(EnvAccessTokenSecondary)
	subaccountDocumentSentId := getEnvValue(EnvSubaccountDocumentSentId)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nSubaccount := NewSubaccount(EnvSandbox, accessToken)
	resp, err := nSubaccount.DeleteWhiteLabelDocumentSentById(ctx, subaccountDocumentSentId)
	assertResponseSuccess(t, resp, err)
}

func TestSubaccountGetById(t *testing.T) {
	initSubaccount()
	accessToken := getEnvValue(EnvAccessTokenSecondary)
	subaccountId := getEnvValue(EnvSubaccountId)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nSubaccount := NewSubaccount(EnvSandbox, accessToken)
	resp, err := nSubaccount.GetById(ctx, subaccountId)
	assertResponseSuccess(t, resp, err)
}

func TestSubaccountGetDocumentSentById(t *testing.T) {
	initSubaccountDocumentSent()
	accessToken := getEnvValue(EnvAccessTokenSecondary)
	subaccountDocumentSentId := getEnvValue(EnvSubaccountDocumentSentId)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nSubaccount := NewSubaccount(EnvSandbox, accessToken)
	resp, err := nSubaccount.GetDocumentSentById(ctx, subaccountDocumentSentId)
	assertResponseSuccess(t, resp, err)
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
	initSubaccountDocumentSent()
	accessToken := getEnvValue(EnvAccessTokenSecondary)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nSubaccount := NewSubaccount(EnvSandbox, accessToken)
	resp, err := nSubaccount.GetPendingDocuments(ctx)
	assertResponseSuccess(t, resp, err)
}
