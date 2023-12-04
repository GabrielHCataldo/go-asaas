package asaas

import (
	"context"
	"encoding/json"
	"github.com/GabrielHCataldo/go-asaas/internal/test"
	"os"
	"testing"
	"time"
)

func TestNegativityCreateSuccess(t *testing.T) {
	accessToken := getEnvValue(EnvAccessToken)
	assertFatalIsBlank(t, accessToken)
	f, err := os.Open(getEnvValue(EnvFileName))
	assertFatalErrorNonnull(t, err)
	v, err := os.ReadFile(f.Name())
	assertFatalErrorNonnull(t, err)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	req := &CreateNegativityRequest{}
	err = json.Unmarshal(test.GetCreateNegativitySuccess(), req)
	assertFatalErrorNonnull(t, err)
	req.Documents = &FileRequest{
		Name: f.Name(),
		Mime: FileMimeTypeText,
		Data: v,
	}
	nNegativity := NewNegativity(EnvSandbox, accessToken)
	resp, errAsaas := nNegativity.Create(ctx, *req)
	assertResponseSuccess(t, resp, errAsaas)
}

func TestNegativityGetChargesAvailableForDunning(t *testing.T) {
	accessToken := getEnvValue(EnvAccessToken)
	assertFatalIsBlank(t, accessToken)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nNegativity := NewNegativity(EnvSandbox, accessToken)
	resp, errAsaas := nNegativity.GetChargesAvailableForDunning(ctx, PageableDefaultRequest{})
	assertResponseSuccess(t, resp, errAsaas)
}
