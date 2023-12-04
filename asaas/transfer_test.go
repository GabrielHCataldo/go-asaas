package asaas

import (
	"context"
	"github.com/mvrilo/go-cpf"
	"testing"
	"time"
)

func TestTransferTransferToBank(t *testing.T) {
	accessToken := getEnvValue(EnvAccessToken)
	assertFatalStringBlank(t, accessToken)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nTransfer := NewTransfer(EnvSandbox, accessToken)
	resp, err := nTransfer.TransferToBank(ctx, TransferToBankRequest{
		Value: 10,
		BankAccount: BackAccountRequest{
			Bank: BankRequest{
				Code: "237",
			},
			AccountName:     "Conta do Bradesco",
			OwnerName:       "Unit test go",
			OwnerBirthDate:  NewDate(1999, 12, 12, time.Local),
			CpfCnpj:         cpf.Generate(),
			Agency:          "0001",
			Account:         "103913",
			AccountDigit:    "8",
			BankAccountType: BankAccountTypeChecking,
			Ispb:            "",
		},
		OperationType:     TransferOperationTypePix,
		PixAddressKey:     "",
		PixAddressKeyType: "",
		Description:       "",
		ScheduleDate:      Date{},
	})
	assertResponseSuccess(t, resp, err)
}

func TestTransferTransferToAsaas(t *testing.T) {
	accessToken := getEnvValue(EnvAccessToken)
	assertFatalStringBlank(t, accessToken)
	walletId := getEnvValue(EnvWalletIdSecondary)
	assertFatalStringBlank(t, walletId)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nTransfer := NewTransfer(EnvSandbox, accessToken)
	resp, err := nTransfer.TransferToAsaas(ctx, TransferToAssasRequest{
		Value:    10,
		WalletId: walletId,
	})
	assertResponseSuccess(t, resp, err)
}

func TestTransferCancelById(t *testing.T) {
	accessToken := getEnvValue(EnvAccessToken)
	assertFatalStringBlank(t, accessToken)
	initTransfer()
	transferId := getEnvValue(EnvTransferId)
	assertFatalStringBlank(t, transferId)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nTransfer := NewTransfer(EnvSandbox, accessToken)
	resp, err := nTransfer.CancelById(ctx, transferId)
	assertResponseSuccess(t, resp, err)
}

func TestTransferGetById(t *testing.T) {
	accessToken := getEnvValue(EnvAccessToken)
	assertFatalStringBlank(t, accessToken)
	initTransfer()
	transferId := getEnvValue(EnvTransferId)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nTransfer := NewTransfer(EnvSandbox, accessToken)
	resp, err := nTransfer.GetById(ctx, transferId)
	assertResponseSuccess(t, resp, err)
}

func TestTransferGetAll(t *testing.T) {
	accessToken := getEnvValue(EnvAccessToken)
	assertFatalStringBlank(t, accessToken)
	initTransfer()
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nTransfer := NewTransfer(EnvSandbox, accessToken)
	resp, err := nTransfer.GetAll(ctx, GetAllTransfersRequest{
		DateCreatedGe:  Date{},
		DateCreatedLe:  Date{},
		TransferDateGe: Date{},
		TransferDateLe: Date{},
		Type:           "",
	})
	assertResponseSuccess(t, resp, err)
}
