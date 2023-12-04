package test

import (
	"io"
	"os"
	"strconv"
	"time"
)

func GetCustomerIdDefault() string {
	return "cus_000005791749"
}

func GetPaymentLinkIdDefault() string {
	return "le5jqxz8as7pgwn9"
}

func GetSimpleFile() (*os.File, error) {
	randomKey := strconv.FormatInt(time.Now().Unix()+int64(time.Now().Nanosecond()), 10)
	nameFile := "test " + randomKey + ".txt"
	f, err := os.Create(nameFile)
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
	return os.Open(nameFile)
}

func GetSimpleImage() (*os.File, error) {
	return os.Open("../gopher-asaas.png")
}

func GetTransferToBankRequestDefault() []byte {
	return []byte(`
		{
			"value": 1,
			"bankAccount": {
				"bank": {
					"code": "237"
				},
				"accountName": "Conta do Bradesco",
				"ownerName": "GC",
				"ownerBirthDate": "1999-01-21",
				"cpfCnpj": "45991590000108",
				"type": "PAYMENT_ACCOUNT",
				"agency": "0001",
				"account": "103913",
				"accountDigit": "8",
				"bankAccountType": "CONTA_CORRENTE"
			},
			"operationType": "PIX",
			"scheduleDate": null,
			"description": "Test via Postman"
		}
`)
}

func GetTransferToBankFailureRequestDefault() []byte {
	return []byte(`
		{
			"value": 0,
			"bankAccount": {
				"ownerName": "GC",
				"ownerBirthDate": "1999-01-21",
				"cpfCnpj": "45991590000108",
				"type": "PAYMENT_ACCOUNT",
				"agency": "0001",
				"account": "103913",
				"bankAccountType": "CONTA_CORRENTE"
			},
			"description": "Test via Postman"
		}
`)
}

func GetCreateNegativitySuccess() []byte {
	return []byte(`
		{
			"payment": "pay_8129071930338672",
			"type": "CREDIT_BUREAU",
			"description": "Unit test Golang",
			"customerName": "Marcelo Almeida",
			"customerCpfCnpj": "24971563792",
			"customerPrimaryPhone": "47999376637",
			"customerPostalCode": "01310-000",
			"customerAddress": "Av. Paulista",
			"customerAddressNumber": "150",
			"customerProvince": "Centro"
		}
`)
}
