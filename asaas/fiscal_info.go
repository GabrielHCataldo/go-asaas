package asaas

import (
	"context"
	"net/http"
	"os"
)

type SaveFiscalInfoRequest struct {
	// Email para notificações de notas fiscais (REQUIRED)
	Email string `json:"email,omitempty" validate:"required,email"`
	// Inscrição municipal da empresa
	MunicipalInscription string `json:"municipalInscription,omitempty"`
	// Indica se a empresa é optante pelo simples nacional
	SimplesNacional bool `json:"simplesNacional"`
	// Identifica se a empresa é classificada como incentivador cultural
	CulturalProjectsPromoter bool `json:"culturalProjectsPromoter,omitempty"`
	// Código CNAE
	Cnae string `json:"cnae,omitempty"`
	// Identificador do regime especial de tributação. Empresas do simples nacional geralmente optam pelo Microempresa Municipal
	SpecialTaxRegime string `json:"specialTaxRegime,omitempty"`
	// Item da lista de serviço, conforme http://www.planalto.gov.br/ccivil_03/leis/LCP/Lcp116.htm
	ServiceListItem string `json:"serviceListItem,omitempty"`
	// Número de Série utilizado pela sua empresa para emissão de notas fiscais. Na maioria das cidades o número de série utilizado é '1' ou 'E'
	RpsSerie string `json:"rpsSerie,omitempty"`
	// Número do RPS utilizado na última nota fiscal emitida pela sua empresa. Se a sua última NF emitida tem RPS igual a '100', esse campo deve ser preenchido com '101'. Se você nunca emitiu notas fiscais pelo site da sua prefeitura, informe '1' nesse campo
	RpsNumber int `json:"rpsNumber,omitempty"`
	// Número do Lote utilizado na última nota fiscal emitida pela sua empresa. Se o último lote utilizado na sua prefeitura for '25', esse campo deve ser preenchido com '26'. Informe esse campo apenas se sua prefeitura exigir a utilização de lotes
	LoteNumber int `json:"loteNumber,omitempty"`
	// Usuário para acesso ao site da prefeitura da sua cidade
	Username string `json:"username,omitempty"`
	// Senha para acesso ao site da prefeitura
	Password string `json:"password,omitempty"`
	// Token de acesso ao site da prefeitura (Caso o acesso ao site da sua prefeitura seja através por Token)
	AccessToken string `json:"accessToken,omitempty"`
	// Arquivo (.pfx ou .p12) do certificado digital da empresa (Caso o acesso ao site da sua prefeitura via certificado digital)
	CertificateFile *os.File `json:"certificateFile,omitempty"`
	// Senha do certificado digital enviado (Caso o acesso ao site da sua prefeitura via certificado digital)
	CertificatePassword string `json:"certificatePassword,omitempty"`
}

type GetAllServicesRequest struct {
	// Nome do serviço
	Description string `json:"description,omitempty"`
	// Elemento inicial da lista
	Offset int `json:"offset,omitempty"`
	// Número de elementos da lista (max: 100)
	Limit int `json:"limit,omitempty"`
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
	Errors                   []ErrorResponse            `json:"errors,omitempty"`
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
	// Save (Criar e atualizar informações fiscais)
	//
	// # Resposta: 200
	//
	// FiscalInfoResponse = not nil
	//
	// Error = nil
	//
	// FiscalInfoResponse.IsSuccess() = true
	//
	// Possui os valores de resposta de sucesso segunda a documentação.
	//
	// # Resposta: 400/401/500
	//
	// FiscalInfoResponse = not nil
	//
	// Error = nil
	//
	// FiscalInfoResponse.IsFailure() = true
	//
	// Para qualquer outra resposta inesperada da API, possuímos o campo FiscalInfoResponse.Errors preenchido com as informações
	// de erro, sendo 400 retornado da API Asaas com as instruções de requisição conforme a documentação,
	// diferente disso retornará uma mensagem padrão no index 0 do slice com campo ErrorResponse.Code retornando a
	// descrição status http (Ex: "401 Unauthorized") e no campo ErrorResponse.Description retornará com o valor
	// "response status code not expected".
	//
	// # Error
	//
	// FiscalInfoResponse = nil
	//
	// Error = not nil
	//
	// Se o campo ErrorAsaas.Type tiver com valor ErrorTypeValidation quer dizer que não passou pela validação dos
	// parâmetros informados segundo a documentação.
	// Por fim se o campo ErrorAsaas.Type tiver com valor ErrorTypeUnexpected quer dizer que ocorreu um erro inesperado
	// na lib go-asaas.
	//
	// Para obter mais detalhes confira as colunas:
	//
	// ErrorAsaas.Msg (mensagem do erro),
	//
	// ErrorAsaas.File (Arquivo aonde ocorreu o erro),
	//
	// ErrorAsaas.Line (Linha aonde ocorreu o erro)
	//
	// Caso ocorra um erro inesperado por favor report o erro no repositório: https://github.com/GabrielHCataldo/go-asaas
	//
	// # DOCS
	//
	// Criar e atualizar informações fiscais: https://docs.asaas.com/reference/criar-e-atualizar-informacoes-fiscais
	Save(ctx context.Context, body SaveFiscalInfoRequest) (*FiscalInfoResponse, Error)
	// Get (Recuperar informações fiscais)
	//
	// Permite verificar as configurações para emissão de notas fiscais. Caso ainda não tenha sido cadastrada,
	// será retornado HTTP 404.
	//
	// # Resposta: 200/404
	//
	// FiscalInfoResponse = not nil
	//
	// Error = nil
	//
	// Se FiscalInfoResponse.IsSuccess() for true quer dizer que possui os valores de resposta de sucesso segunda a documentação.
	//
	// Se FiscalInfoResponse.IsNoContent() for true quer dizer não recuperou nenhum dado.
	//
	// # Resposta: 401/500
	//
	// FiscalInfoResponse = not nil
	//
	// Error = nil
	//
	// FiscalInfoResponse.IsFailure() = true
	//
	// Para qualquer outra resposta inesperada da API, possuímos o campo FiscalInfoResponse.Errors preenchido com as informações
	// de erro, o index 0 do slice com campo ErrorResponse.Code retornando a descrição
	// status http (Ex: "401 Unauthorized") e no campo ErrorResponse.Description retornará com o valor
	// "response status code not expected".
	//
	// # Error
	//
	// FiscalInfoResponse = nil
	//
	// Error = not nil
	//
	// Se o campo ErrorAsaas.Type tiver com valor ErrorTypeValidation quer dizer que não passou pela validação dos
	// parâmetros informados segundo a documentação.
	// Por fim se o campo ErrorAsaas.Type tiver com valor ErrorTypeUnexpected quer dizer que ocorreu um erro inesperado
	// na lib go-asaas.
	//
	// Para obter mais detalhes confira as colunas:
	//
	// ErrorAsaas.Msg (mensagem do erro),
	//
	// ErrorAsaas.File (Arquivo aonde ocorreu o erro),
	//
	// ErrorAsaas.Line (Linha aonde ocorreu o erro)
	//
	// Caso ocorra um erro inesperado por favor report o erro no repositório: https://github.com/GabrielHCataldo/go-asaas
	//
	// # DOCS
	//
	// Recuperar informações fiscais: https://docs.asaas.com/reference/recuperar-informacoes-fiscais
	Get(ctx context.Context) (*FiscalInfoResponse, Error)
	// GetMunicipalSettings (Listar configurações municipais)
	//
	// Algumas configurações para emissão de notas fiscais dependem da prefeitura onde a nota fiscal é emitida.
	//
	// # Resposta: 200
	//
	// MunicipalSettingsResponse = not nil
	//
	// Error = nil
	//
	// MunicipalSettingsResponse.IsSuccess() = true
	//
	// Possui os valores de resposta de sucesso segunda a documentação.
	//
	// # Resposta: 400/401/500
	//
	// MunicipalSettingsResponse = not nil
	//
	// Error = nil
	//
	// MunicipalSettingsResponse.IsFailure() = true
	//
	// Para qualquer outra resposta inesperada da API, possuímos o campo MunicipalSettingsResponse.Errors preenchido com as informações
	// de erro, sendo 400 retornado da API Asaas com as instruções de requisição conforme a documentação,
	// diferente disso retornará uma mensagem padrão no index 0 do slice com campo ErrorResponse.Code retornando a
	// descrição status http (Ex: "401 Unauthorized") e no campo ErrorResponse.Description retornará com o valor
	// "response status code not expected".
	//
	// # Error
	//
	// MunicipalSettingsResponse = nil
	//
	// Error = not nil
	//
	// Se o campo ErrorAsaas.Type tiver com valor ErrorTypeValidation quer dizer que não passou pela validação dos
	// parâmetros informados segundo a documentação.
	// Por fim se o campo ErrorAsaas.Type tiver com valor ErrorTypeUnexpected quer dizer que ocorreu um erro inesperado
	// na lib go-asaas.
	//
	// Para obter mais detalhes confira as colunas:
	//
	// ErrorAsaas.Msg (mensagem do erro),
	//
	// ErrorAsaas.File (Arquivo aonde ocorreu o erro),
	//
	// ErrorAsaas.Line (Linha aonde ocorreu o erro)
	//
	// Caso ocorra um erro inesperado por favor report o erro no repositório: https://github.com/GabrielHCataldo/go-asaas
	//
	// # DOCS
	//
	// Listar configurações municipais: https://docs.asaas.com/reference/listar-configuracoes-municipais
	GetMunicipalSettings(ctx context.Context) (*MunicipalSettingsResponse, Error)
	// GetAllServices (Listar serviços municipais)
	//
	// # Resposta: 200
	//
	// Pageable(FiscalInfoServiceResponse) = not nil
	//
	// Error = nil
	//
	// Se Pageable.IsSuccess() for true quer dizer que retornaram os dados conforme a documentação.
	// Se Pageable.IsNoContent() for true quer dizer que retornou os dados vazio.
	//
	// Error = nil
	//
	// Pageable.IsNoContent() = true
	//
	// Pageable.Data retornou vazio.
	//
	// # Resposta: 401/500
	//
	// Pageable(FiscalInfoServiceResponse) = not nil
	//
	// Error = nil
	//
	// Pageable.IsFailure() = true
	//
	// Para qualquer outra resposta inesperada da API, possuímos o campo Pageable.Errors preenchido com
	// as informações de erro, o index 0 do slice com campo ErrorResponse.Code retornando a descrição
	// status http (Ex: "401 Unauthorized") e no campo ErrorResponse.Description retornará com o valor
	// "response status code not expected".
	//
	// # Error
	//
	// Pageable(FiscalInfoServiceResponse) = nil
	//
	// Error = not nil
	//
	// Se o campo ErrorAsaas.Type tiver com valor ErrorTypeValidation quer dizer que não passou pela validação dos
	// parâmetros informados segundo a documentação.
	// Por fim se o campo ErrorAsaas.Type tiver com valor ErrorTypeUnexpected quer dizer que ocorreu um erro inesperado
	// na lib go-asaas.
	//
	// Para obter mais detalhes confira as colunas:
	//
	// ErrorAsaas.Msg (mensagem do erro),
	//
	// ErrorAsaas.File (Arquivo aonde ocorreu o erro),
	//
	// ErrorAsaas.Line (Linha aonde ocorreu o erro)
	//
	// Caso ocorra um erro inesperado por favor report o erro no repositório: https://github.com/GabrielHCataldo/go-asaas
	//
	// # DOCS
	//
	// Listar serviços municipais: https://docs.asaas.com/reference/listar-servicos-municipais
	GetAllServices(ctx context.Context, filter GetAllServicesRequest) (*Pageable[FiscalInfoServiceResponse], Error)
}

func NewFiscalInfo(env Env, accessToken string) FiscalInfo {
	logWarning("FiscalInfo service running on", env.String())
	return fiscalInfo{
		env:         env,
		accessToken: accessToken,
	}
}

func (f fiscalInfo) Save(ctx context.Context, body SaveFiscalInfoRequest) (*FiscalInfoResponse, Error) {
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
