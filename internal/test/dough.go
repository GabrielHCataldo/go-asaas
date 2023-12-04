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
