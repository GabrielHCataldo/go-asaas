package asaas

import (
	"context"
	"testing"
	"time"
)

func TestCreditBureauGetReport(t *testing.T) {
	accessToken := getEnvValue(EnvAccessToken)
	assertFatalStringBlank(t, accessToken)
	initCustomer()
	customerId := getEnvValue(EnvCustomerId)
	assertFatalStringBlank(t, customerId)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
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
	accessToken := getEnvValue(EnvAccessToken)
	assertFatalStringBlank(t, accessToken)
	initCreditBureauReport()
	reportId := getEnvValue(EnvCreditBureauReportId)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nCreditBureau := NewCreditBureau(EnvSandbox, accessToken)
	resp, err := nCreditBureau.GetReportById(ctx, reportId)
	assertResponseSuccess(t, resp, err)
}

func TestCreditBureauGetAllReports(t *testing.T) {
	accessToken := getEnvValue(EnvAccessToken)
	assertFatalStringBlank(t, accessToken)
	initCreditBureauReport()
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nCreditBureau := NewCreditBureau(EnvSandbox, accessToken)
	resp, err := nCreditBureau.GetAllReports(ctx, GetAllReportsRequest{
		StartDate: Date{},
		EndDate:   Date{},
		Offset:    0,
		Limit:     10,
	})
	assertResponseNoContent(t, resp, err)
}
