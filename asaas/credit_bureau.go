package asaas

import (
	"context"
	berrors "errors"
	"fmt"
	"github.com/GabrielHCataldo/go-asaas/internal/util"
	"net/http"
)

type GetReportRequest struct {
	Customer string `json:"customer,omitempty"`
	CpfCnpj  string `json:"cpfCnpj,omitempty" validate:"omitempty,document"`
	State    string `json:"state,omitempty"` //todo -> passar para enum ou criar uma validacao
}

type GetAllReportsRequest struct {
	StartDate *Date `json:"startDate,omitempty"`
	EndDate   *Date `json:"endDate,omitempty"`
	Offset    int   `json:"offset,omitempty"`
	Limit     int   `json:"limit,omitempty"`
}

type CreditBureauReportResponse struct {
	ID          string          `json:"id,omitempty"`
	Customer    string          `json:"customer,omitempty"`
	CpfCnpj     string          `json:"cpfCnpj,omitempty"`
	State       string          `json:"state,omitempty"`
	DownloadUrl string          `json:"downloadUrl,omitempty"`
	DateCreated *Date           `json:"dateCreated,omitempty"`
	Errors      []ErrorResponse `json:"errors,omitempty"`
}

type creditBureau struct {
	env         Env
	accessToken string
}

type CreditBureau interface {
	GetReport(ctx context.Context, body GetReportRequest) (*CreditBureauReportResponse, Error)
	GetReportByID(ctx context.Context, creditBureauReportID string) (*CreditBureauReportResponse, Error)
	GetAllReports(ctx context.Context, filter GetAllReportsRequest) (*Pageable[CreditBureauReportResponse], Error)
}

func NewCreditBureau(env Env, accessToken string) CreditBureau {
	logWarning("CreditBureau service running on", env.String())
	return creditBureau{
		env:         env,
		accessToken: accessToken,
	}
}

func (c creditBureau) GetReport(ctx context.Context, body GetReportRequest) (*CreditBureauReportResponse, Error) {
	if err := c.validateBodyReportRequest(body); err != nil {
		return nil, NewError(ErrorTypeValidation, err)
	}
	req := NewRequest[CreditBureauReportResponse](ctx, c.env, c.accessToken)
	return req.make(http.MethodPost, "/v3/creditBureauReport", body)
}

func (c creditBureau) GetReportByID(ctx context.Context, creditBureauReportID string) (*CreditBureauReportResponse, Error) {
	req := NewRequest[CreditBureauReportResponse](ctx, c.env, c.accessToken)
	return req.make(http.MethodGet, fmt.Sprintf("/v3/creditBureauReport/%s", creditBureauReportID), nil)
}

func (c creditBureau) GetAllReports(ctx context.Context, filter GetAllReportsRequest) (
	*Pageable[CreditBureauReportResponse], Error) {
	req := NewRequest[Pageable[CreditBureauReportResponse]](ctx, c.env, c.accessToken)
	return req.make(http.MethodGet, "/v3/creditBureauReport", filter)
}

func (c creditBureau) validateBodyReportRequest(body GetReportRequest) error {
	if err := Validate().Struct(body); err != nil {
		return err
	} else if util.IsBlank(&body.Customer) && util.IsBlank(&body.State) {
		return berrors.New("state is required if customer is empty")
	}
	return nil
}
