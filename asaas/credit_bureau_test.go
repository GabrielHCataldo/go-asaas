package asaas

import (
	"context"
	"testing"
	"time"
)

func TestCreditBureauGetReportSuccess(t *testing.T) {
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

func TestCreditBureauGetReportError(t *testing.T) {
	nCreditBureau := NewCreditBureau(EnvSandbox, "")
	resp, err := nCreditBureau.GetReport(context.TODO(), GetReportRequest{})
	assertResponseFailure(t, resp, err)
}

func TestCreditBureauGetReportByIdSuccess(t *testing.T) {
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

func TestCreditBureauGetReportByIdError(t *testing.T) {
	nCreditBureau := NewCreditBureau(EnvSandbox, "")
	resp, err := nCreditBureau.GetReportById(context.TODO(), "")
	assertResponseFailure(t, resp, err)
}

func TestCreditBureauGetAllReportsSuccess(t *testing.T) {
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

func TestCreditBureauGetAllReportsError(t *testing.T) {
	nCreditBureau := NewCreditBureau(EnvSandbox, "")
	resp, err := nCreditBureau.GetAllReports(context.TODO(), GetAllReportsRequest{})
	assertResponseFailure(t, resp, err)
}
