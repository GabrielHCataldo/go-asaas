package asaas

import (
	"context"
	"net/http"
	"os"
)

type UpdateFiscalInfoRequest struct {
	Email                    string   `json:"email,omitempty" validate:"required,email"`
	MunicipalInscription     string   `json:"municipalInscription,omitempty"`
	SimplesNacional          bool     `json:"simplesNacional"`
	CulturalProjectsPromoter bool     `json:"culturalProjectsPromoter,omitempty"`
	Cnae                     string   `json:"cnae,omitempty"`
	SpecialTaxRegime         string   `json:"specialTaxRegime,omitempty"`
	ServiceListItem          string   `json:"serviceListItem,omitempty"`
	RpsSerie                 string   `json:"rpsSerie,omitempty"`
	RpsNumber                int      `json:"rpsNumber,omitempty"`
	LoteNumber               int      `json:"loteNumber,omitempty"`
	Username                 string   `json:"username,omitempty"`
	Password                 string   `json:"password,omitempty"`
	AccessToken              string   `json:"accessToken,omitempty"`
	CertificateFile          *os.File `json:"certificateFile,omitempty"`
	CertificatePassword      string   `json:"certificatePassword,omitempty"`
}

type GetAllServicesRequest struct {
	Description string `json:"description,omitempty"`
	Offset      int    `json:"offset,omitempty"`
	Limit       int    `json:"limit,omitempty"`
}

type FiscalInfoResponse struct {
	SimplesNacional          bool            `json:"simplesNacional,omitempty"`
	RpsSerie                 string          `json:"rpsSerie,omitempty"`
	RpsNumber                int             `json:"rpsNumber,omitempty"`
	LoteNumber               int             `json:"loteNumber,omitempty"`
	Username                 string          `json:"username,omitempty"`
	PasswordSent             string          `json:"passwordSent,omitempty"`
	AccessTokenSent          string          `json:"accessTokenSent,omitempty"`
	CertificateSent          bool            `json:"certificateSent,omitempty"`
	SpecialTaxRegime         string          `json:"specialTaxRegime,omitempty"`
	Email                    string          `json:"email,omitempty"`
	ServiceListItem          string          `json:"serviceListItem,omitempty"`
	Cnae                     string          `json:"cnae,omitempty"`
	CulturalProjectsPromoter bool            `json:"culturalProjectsPromoter,omitempty"`
	MunicipalInscription     string          `json:"municipalInscription,omitempty"`
	UseNationalPortal        bool            `json:"useNationalPortal,omitempty"`
	Errors                   []ErrorResponse `json:"errors,omitempty"`
}

type MunicipalSettingsResponse struct {
	AuthenticationType       string                     `json:"authenticationType,omitempty"`
	SupportsCancellation     bool                       `json:"supportsCancellation,omitempty"`
	UsesSpecialTaxRegimes    bool                       `json:"usesSpecialTaxRegimes,omitempty"`
	UsesServiceListItem      bool                       `json:"usesServiceListItem,omitempty"`
	SpecialTaxRegimesList    []SpecialTaxRegimeResponse `json:"specialTaxRegimesList,omitempty"`
	MunicipalInscriptionHelp string                     `json:"municipalInscriptionHelp,omitempty"`
	SpecialTaxRegimeHelp     string                     `json:"specialTaxRegimeHelp,omitempty"`
	ServiceListItemHelp      string                     `json:"serviceListItemHelp,omitempty"`
	DigitalCertificatedHelp  string                     `json:"digitalCertificatedHelp,omitempty"`
	AccessTokenHelp          string                     `json:"accessTokenHelp,omitempty"`
	MunicipalServiceCodeHelp string                     `json:"municipalServiceCodeHelp,omitempty"`
}

type SpecialTaxRegimeResponse struct {
	Label string  `json:"label,omitempty"`
	Value float64 `json:"value,omitempty"`
}

type FiscalInfoServiceResponse struct {
	Id          string  `json:"id,omitempty"`
	Description string  `json:"description,omitempty"`
	IssTax      float64 `json:"issTax,omitempty"`
}

type fiscalInfo struct {
	env         Env
	accessToken string
}

type FiscalInfo interface {
	Update(ctx context.Context, body UpdateFiscalInfoRequest) (*FiscalInfoResponse, Error)
	Get(ctx context.Context) (*FiscalInfoResponse, Error)
	GetMunicipalSettings(ctx context.Context) (*MunicipalSettingsResponse, Error)
	GetAllServices(ctx context.Context, filter GetAllServicesRequest) (*Pageable[FiscalInfoServiceResponse], Error)
}

func NewFiscalInfo(env Env, accessToken string) FiscalInfo {
	logWarning("FiscalInfo service running on", env.String())
	return fiscalInfo{
		env:         env,
		accessToken: accessToken,
	}
}

func (f fiscalInfo) Update(ctx context.Context, body UpdateFiscalInfoRequest) (*FiscalInfoResponse, Error) {
	if err := Validate().Struct(body); err != nil {
		return nil, NewError(ErrorTypeValidation, err)
	}
	req := NewRequest[FiscalInfoResponse](ctx, f.env, f.accessToken)
	return req.makeMultipartForm(http.MethodPost, "/v3/fiscalInfo", body)
}

func (f fiscalInfo) Get(ctx context.Context) (*FiscalInfoResponse, Error) {
	req := NewRequest[FiscalInfoResponse](ctx, f.env, f.accessToken)
	return req.make(http.MethodGet, "/v3/fiscalInfo", nil)
}

func (f fiscalInfo) GetMunicipalSettings(ctx context.Context) (*MunicipalSettingsResponse, Error) {
	req := NewRequest[MunicipalSettingsResponse](ctx, f.env, f.accessToken)
	return req.make(http.MethodGet, "/v3/fiscalInfo/municipalOptions", nil)
}

func (f fiscalInfo) GetAllServices(ctx context.Context, filter GetAllServicesRequest) (
	*Pageable[FiscalInfoServiceResponse], Error) {
	req := NewRequest[Pageable[FiscalInfoServiceResponse]](ctx, f.env, f.accessToken)
	return req.make(http.MethodGet, "/v3/fiscalInfo/services", filter)
}
