package asaas

import (
	"context"
	"encoding/json"
	"github.com/GabrielHCataldo/go-asaas/internal/test"
	"testing"
	"time"
)

func TestCustomerCreate(t *testing.T) {
	accessToken, err := test.GetAccessTokenByEnv()
	AssertFatalErrorNonnull(t, err)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nCustomer := NewCustomer(SANDBOX, *accessToken)
	resp, errAsaas := nCustomer.Create(ctx, CreateCustomerRequest{
		Name:                 "Unit Test Golang",
		CpfCnpj:              "24971563792",
		Email:                "unittest@gmail.com",
		Phone:                "",
		MobilePhone:          "",
		Address:              "",
		AddressNumber:        "",
		Complement:           "",
		Province:             "",
		PostalCode:           "",
		ExternalReference:    "",
		NotificationDisabled: false,
		AdditionalEmails:     "",
		MunicipalInscription: "",
	})
	if errAsaas != nil {
		logDebug("unexpect response: err is not nil:", errAsaas)
		t.Fail()
	} else {
		respJson, _ := json.Marshal(resp)
		logDebug("result success response:", string(respJson))
	}
}
