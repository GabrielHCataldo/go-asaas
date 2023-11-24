package test

import (
	"errors"
	"github.com/GabrielHCataldo/go-asaas/internal/util"
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
