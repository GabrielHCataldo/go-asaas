package asaas

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func CreateHttpRequest(
	ctx context.Context,
	asaasEnv Env,
	accessCode string,
	method string,
	path string,
	payload any,
) (*http.Request, error) {
	var payloadToSend io.Reader
	if payload != nil {
		payloadBytes, err := json.Marshal(payload)
		if err != nil {
			return nil, err
		}
		payloadToSend = bytes.NewReader(payloadBytes)
	}
	req, _ := http.NewRequestWithContext(ctx, method, asaasEnv.BaseURL()+path, payloadToSend)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("access_token", accessCode)
	return req, nil
}

func MakeHttpRequest(req *http.Request) (*http.Response, error) {
	client := &http.Client{}
	return client.Do(req)
}

func CloseBody(Body io.ReadCloser) {
	err := Body.Close()
	if err != nil {
		log.Fatal("error close read body:", err)
	}
}
