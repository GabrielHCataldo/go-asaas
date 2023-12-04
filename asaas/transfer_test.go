package asaas

import (
	"context"
	"encoding/json"
	"github.com/GabrielHCataldo/go-asaas/internal/test"
	"testing"
	"time"
)

func TestTransferTransferToBankSuccess(t *testing.T) {
	accessToken := getEnvValue(EnvAccessToken)
	assertFatalStringBlank(t, accessToken)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	req := &TransferToBankRequest{}
	err := json.Unmarshal(test.GetTransferToBankRequestDefault(), req)
	assertFatalErrorNonnull(t, err)
	nTransfer := NewTransfer(EnvSandbox, accessToken)
	resp, errAsaas := nTransfer.TransferToBank(ctx, *req)
	assertResponseSuccess(t, resp, errAsaas)
}

func TestTransferTransferToBankFailure(t *testing.T) {
	accessToken := getEnvValue(EnvAccessToken)
	assertFatalStringBlank(t, accessToken)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	req := &TransferToBankRequest{}
	err := json.Unmarshal(test.GetTransferToBankFailureRequestDefault(), req)
	assertFatalErrorNonnull(t, err)
	nTransfer := NewTransfer(EnvSandbox, accessToken)
	resp, err := nTransfer.TransferToBank(ctx, *req)
	assertResponseFailure(t, resp, err)
}
