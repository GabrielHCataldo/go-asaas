package test

import (
	"io"
	"os"
	"strconv"
	"time"
)

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
