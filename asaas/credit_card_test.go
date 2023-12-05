package asaas

import (
	"context"
	"testing"
	"time"
)

func TestCreditCardTokenize(t *testing.T) {
	initCustomer(false)
	accessToken := getEnvValue(EnvAccessToken)
	customerId := getEnvValue(EnvCustomerId)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nCreditCard := NewCreditCard(EnvSandbox, accessToken)
	resp, err := nCreditCard.Tokenize(ctx, CreditCardTokenizeRequest{
		Customer: customerId,
		CreditCard: CreditCardRequest{
			HolderName:  "unit test go",
			Number:      "4000000000000010",
			ExpiryMonth: "05",
			ExpiryYear:  "2035",
			Ccv:         "318",
		},
		CreditCardHolderInfo: CreditCardHolderInfoRequest{
			Name:              "Unit Test Go",
			CpfCnpj:           "24971563792",
			Email:             "unittest@gmail.com",
			Phone:             "4738010919",
			MobilePhone:       "47998781877",
			PostalCode:        "89223-005",
			AddressNumber:     "277",
			AddressComplement: "",
		},
	})
	assertResponseSuccess(t, resp, err)
}
