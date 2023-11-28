package asaas

import (
	"context"
	berrors "errors"
	"fmt"
	"github.com/GabrielHCataldo/go-asaas/internal/util"
	"net/http"
	"os"
)

type AnticipationRequest struct {
	Payment     string     `json:"payment,omitempty"`
	Installment string     `json:"installment,omitempty"`
	Documents   []*os.File `json:"documents,omitempty"`
}

type AnticipationSimulateRequest struct {
	Payment     string `json:"payment,omitempty"`
	Installment string `json:"installment,omitempty"`
}

type AgreementSignRequest struct {
	Agreed bool `json:"agreed,omitempty"`
}

type GetAllAnticipationsRequest struct {
	Payment     string             `json:"payment,omitempty"`
	Installment string             `json:"installment,omitempty"`
	Status      AnticipationStatus `json:"status,omitempty"`
	Offset      int                `json:"offset,omitempty"`
	Limit       int                `json:"limit,omitempty"`
}

type AnticipationResponse struct {
	ID                string             `json:"id,omitempty"`
	Installment       string             `json:"installment,omitempty"`
	Payment           string             `json:"payment,omitempty"`
	Status            AnticipationStatus `json:"status,omitempty"`
	AnticipationDate  Date               `json:"anticipationDate,omitempty"`
	DueDate           Date               `json:"dueDate,omitempty"`
	RequestDate       Date               `json:"requestDate,omitempty"`
	Fee               float64            `json:"fee,omitempty"`
	AnticipationDays  int                `json:"anticipationDays,omitempty"`
	NetValue          float64            `json:"netValue,omitempty"`
	Value             float64            `json:"value,omitempty"`
	TotalValue        float64            `json:"totalValue,omitempty"`
	DenialObservation string             `json:"denialObservation,omitempty"`
	Errors            []ErrorResponse    `json:"errors,omitempty"`
}

type AnticipationLimitsResponse struct {
	CreditCard AnticipationLimitResponse `json:"creditCard,omitempty"`
	BankSlip   AnticipationLimitResponse `json:"bankSlip,omitempty"`
}

type AnticipationLimitResponse struct {
	Total     float64 `json:"total,omitempty"`
	Available float64 `json:"available,omitempty"`
}

type AnticipationSimulateResponse struct {
	Installment             string          `json:"installment,omitempty"`
	Payment                 string          `json:"payment,omitempty"`
	AnticipationDate        Date            `json:"anticipationDate,omitempty"`
	DueDate                 Date            `json:"dueDate,omitempty"`
	Fee                     float64         `json:"fee,omitempty"`
	AnticipationDays        int             `json:"anticipationDays,omitempty"`
	NetValue                float64         `json:"netValue,omitempty"`
	Value                   float64         `json:"value,omitempty"`
	TotalValue              float64         `json:"totalValue,omitempty"`
	IsDocumentationRequired bool            `json:"isDocumentationRequired,omitempty"`
	Errors                  []ErrorResponse `json:"errors,omitempty"`
}

type AgreementSignResponse struct {
	Agreed bool `json:"agreed,omitempty"`
}

type anticipation struct {
	env         Env
	accessToken string
}

type Anticipation interface {
	Request(ctx context.Context, body AnticipationRequest) (*AnticipationResponse, Error)
	Simulate(ctx context.Context, body AnticipationSimulateRequest) (*AnticipationSimulateResponse, Error)
	AgreementSign(ctx context.Context, body AgreementSignRequest) (*AgreementSignResponse, Error)
	GetByID(ctx context.Context, anticipationID string) (*AnticipationResponse, Error)
	GetLimits(ctx context.Context) (*AnticipationLimitsResponse, Error)
	GetAll(ctx context.Context, filter GetAllAnticipationsRequest) (*Pageable[AnticipationResponse], Error)
}

func NewAnticipation(env Env, accessToken string) Anticipation {
	logWarning("Anticipation service running on", env.String())
	return anticipation{
		env:         env,
		accessToken: accessToken,
	}
}

func (a anticipation) Request(ctx context.Context, body AnticipationRequest) (*AnticipationResponse, Error) {
	if err := a.validateBodyRequest(&body.Payment, &body.Installment); err != nil {
		return nil, NewError(ErrorTypeValidation, err)
	}
	req := NewRequest[AnticipationResponse](ctx, a.env, a.accessToken)
	return req.makeMultipartForm(http.MethodPost, "/v3/anticipations", body)
}

func (a anticipation) Simulate(ctx context.Context, body AnticipationSimulateRequest) (*AnticipationSimulateResponse,
	Error) {
	if err := a.validateBodyRequest(&body.Payment, &body.Installment); err != nil {
		return nil, NewError(ErrorTypeValidation, err)
	}
	req := NewRequest[AnticipationSimulateResponse](ctx, a.env, a.accessToken)
	return req.make(http.MethodPost, "/v3/anticipations/simulate", body)
}

func (a anticipation) AgreementSign(ctx context.Context, body AgreementSignRequest) (*AgreementSignResponse, Error) {
	req := NewRequest[AgreementSignResponse](ctx, a.env, a.accessToken)
	return req.make(http.MethodPost, "/v3/anticipations/agreement/sign", body)
}

func (a anticipation) GetByID(ctx context.Context, anticipationID string) (*AnticipationResponse, Error) {
	req := NewRequest[AnticipationResponse](ctx, a.env, a.accessToken)
	return req.make(http.MethodGet, fmt.Sprintf("/v3/anticipations/%s", anticipationID), nil)
}

func (a anticipation) GetLimits(ctx context.Context) (*AnticipationLimitsResponse, Error) {
	req := NewRequest[AnticipationLimitsResponse](ctx, a.env, a.accessToken)
	return req.make(http.MethodGet, "/v3/anticipations/limits", nil)
}

func (a anticipation) GetAll(ctx context.Context, filter GetAllAnticipationsRequest) (
	*Pageable[AnticipationResponse], Error) {
	req := NewRequest[Pageable[AnticipationResponse]](ctx, a.env, a.accessToken)
	return req.make(http.MethodGet, "/v3/anticipations", filter)
}

func (a anticipation) validateBodyRequest(payment, installment *string) error {
	if util.IsBlank(payment) && util.IsBlank(installment) {
		return berrors.New("inform payment or installment")
	}
	return nil
}
