package asaas

import (
	"context"
	"testing"
	"time"
)

func TestCreditBureauGetAllReportsNoContent(t *testing.T) {
	accessToken, err := getAccessToken()
	assertFatalErrorNonnull(t, err)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nCreditBureau := NewCreditBureau(EnvSandbox, accessToken)
	resp, errAsaas := nCreditBureau.GetAllReports(ctx, GetAllReportsRequest{})
	assertResponseNoContent(t, resp, errAsaas)
}
