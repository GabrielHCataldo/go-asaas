package asaas

import (
	"context"
	"testing"
	"time"
)

func TestCreditBureauGetReport(t *testing.T) {
	initCustomer(false)
	accessToken := getEnvValue(EnvAccessToken)
	customerId := getEnvValue(EnvCustomerId)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nCreditBureau := NewCreditBureau(EnvSandbox, accessToken)
	resp, err := nCreditBureau.GetReport(ctx, GetReportRequest{
		Customer: customerId,
		CpfCnpj:  "",
		State:    "SP",
	})
	assertResponseSuccess(t, resp, err)
}

func TestCreditBureauGetReportById(t *testing.T) {
	initCreditBureauReport()
	accessToken := getEnvValue(EnvAccessToken)
	reportId := getEnvValue(EnvCreditBureauReportId)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nCreditBureau := NewCreditBureau(EnvSandbox, accessToken)
	resp, err := nCreditBureau.GetReportById(ctx, reportId)
	assertResponseSuccess(t, resp, err)
}

func TestCreditBureauGetAllReports(t *testing.T) {
	initCreditBureauReport()
	accessToken := getEnvValue(EnvAccessToken)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nCreditBureau := NewCreditBureau(EnvSandbox, accessToken)
	resp, err := nCreditBureau.GetAllReports(ctx, GetAllReportsRequest{
		StartDate: Date{},
		EndDate:   Date{},
		Offset:    0,
		Limit:     10,
	})
	assertResponseNoContent(resp, err)
}
