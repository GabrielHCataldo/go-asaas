package test

import (
	"errors"
	"github.com/GabrielHCataldo/go-asaas/internal/util"
	"io"
	"os"
)

const PrefixEnvAccessToken = "ASAAS_ACCESS_TOKEN"
const MessageAccessTokenRequired = "ASAAS_ACCESS_TOKEN env is required"

func GetAccessTokenByEnv() (*string, error) {
	accessToken := os.Getenv(PrefixEnvAccessToken)
	if util.IsBlank(&accessToken) {
		return nil, errors.New(MessageAccessTokenRequired)
	}
	return &accessToken, nil
}

func GetChargeIdDefault() string {
	return "pay_0190206161077023"
}

func GetChargePixRequestDefault() []byte {
	return []byte(`
		{
			"billingType": "PIX",
			"customer": "cus_000005791749",
			"dueDate": "2100-11-26",
			"value": 100,
			"description": "Cobrança via teste unitário em Goland",
		}
	`)
}

func GetChargeBoletoRequestDefault() []byte {
	return []byte(`
		{
			"customer": "cus_000005791749",
			"billingType": "BOLETO",
			"discount": {
				"value": 10,
				"dueDateLimitDays": 0
			},
			"interest": {
				"value": 2
			},
			"fine": {
				"value": 1
			},
			"dueDate": "2100-11-26",
			"value": 100,
			"description": "Cobrança via teste unitário em Goland",
		}
	`)
}

func GetChargeCreditCardRequestDefault() []byte {
	return []byte(`
		{
			"customer": "cus_000005791749",
			"billingType": "CREDIT_CARD",
			"value": 10.0,
			"dueDate": "2100-11-26",
			"description": "Cobrança via teste unitário em Goland",
			"creditCard": {
				"holderName": "unit test go",
				"number": "4000000000000010",
				"expiryMonth": "05",
				"expiryYear": "2035",
				"ccv": "318"
			},
			"creditCardHolderInfo": {
				"name": "Unit Test Go",
				"cpfCnpj": "24971563792",
				"email": "unittest@gmail.com",
				"phone": "4738010919",
				"mobilePhone": "47998781877",
				"postalCode": "89223-005",
				"addressNumber": "277"
			},
			"remoteIp": "116.213.42.532"
		}
	`)
}

func GetChargeCreditCardFailureRequestDefault() []byte {
	return []byte(`
		{
			"customer": "cus_000005791749",
			"billingType": "CREDIT_CARD",
			"value": 10.0,
			"dueDate": "2080-11-26",
			"description": "Cobrança via teste unitário em Goland",
			"creditCard": {
				"holderName": "unit test go",
				"number": "5184019740373151",
				"expiryMonth": "05",
				"expiryYear": "2035",
				"ccv": "318"
			},
			"creditCardHolderInfo": {
				"name": "Unit Test Go",
				"cpfCnpj": "24971563792",
				"email": "unittest@gmail.com",
				"phone": "4738010919",
				"mobilePhone": "47998781877",
				"postalCode": "89223-005",
				"addressNumber": "277"
			},
			"remoteIp": "191.253.125.194"
		}
	`)
}

func GetCreditCardFailureRequestDefault() []byte {
	return []byte(`
		{
			"customer": "cus_000005791749",
			"creditCard": {
				"holderName": "unit test go",
				"number": "5184019740373151",
				"expiryMonth": "05",
				"expiryYear": "2035",
				"ccv": "318"
			},
			"creditCardHolderInfo": {
				"name": "Unit Test Go",
				"cpfCnpj": "24971563792",
				"email": "unittest@gmail.com",
				"phone": "4738010919",
				"mobilePhone": "47998781877",
				"postalCode": "89223-005",
				"addressNumber": "277"
			}
		}
	`)
}

func GetCreditCardRequestDefault() []byte {
	return []byte(`
		{
			"customer": "cus_000005791749",
			"creditCard": {
				"holderName": "unit test go",
				"number": "4000000000000010",
				"expiryMonth": "05",
				"expiryYear": "2035",
				"ccv": "318"
			},
			"creditCardHolderInfo": {
				"name": "Unit Test Go",
				"cpfCnpj": "24971563792",
				"email": "unittest@gmail.com",
				"phone": "4738010919",
				"mobilePhone": "47998781877",
				"postalCode": "89223-005",
				"addressNumber": "277"
			}
		}
	`)
}

func GetSimpleFile() (*os.File, error) {
	f, err := os.Create("test.txt")
	if err != nil {
		return nil, err
	}
	_, err = io.WriteString(f, "unit test golang")
	if err != nil {
		return nil, err
	}
	err = f.Close()
	if err != nil {
		return nil, err
	}
	return os.Open("test.txt")
}