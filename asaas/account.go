package asaas

import (
	"context"
	"net/http"
	"os"
)

type SaveInvoiceCustomizationRequest struct {
	// Cor de fundo do logo (REQUIRED)
	LogoBackgroundColor string `json:"logoBackgroundColor,omitempty" validate:"required,color"`
	// Cor de fundo das suas informações (REQUIRED)
	InfoBackgroundColor string `json:"infoBackgroundColor,omitempty" validate:"required,color"`
	// Cor da fonte das suas informações (REQUIRED)
	FontColor string `json:"fontColor,omitempty" validate:"required"`
	// True para habilitar a personalização
	Enabled bool `json:"enabled,omitempty"`
	// Logo que aparecerá no topo da fatura
	LogoFile *os.File `json:"logoFile,omitempty"`
}

type UpdateAccountRequest struct {
	// Tipo de Pessoa
	PersonType PersonType `json:"personType,omitempty" validate:"omitempty,enum"`
	// CPF ou CNPJ do proprietário da conta
	CpfCnpj string `json:"cpfCnpj,omitempty" validate:"omitempty,document"`
	// Data de nascimento necessária caso as informações forem de pessoa física
	BirthDate Date `json:"birthDate,omitempty" validate:"omitempty,before_now"`
	// Tipo da empresa (somente quando Pessoa Jurídica)
	CompanyType CompanyType `json:"companyType,omitempty" validate:"omitempty,enum"`
	// Email da conta
	Email string `json:"email,omitempty" validate:"omitempty,email"`
	// Telefone
	Phone string `json:"phone,omitempty" validate:"omitempty,phone"`
	// Telefone celular
	MobilePhone string `json:"mobilePhone,omitempty" validate:"omitempty,phone"`
	// URL do site da conta
	Site string `json:"site,omitempty" validate:"omitempty,url"`
	// CEP do endereço
	PostalCode string `json:"postalCode,omitempty" validate:"omitempty,postal_code"`
	// Logradouro
	Address string `json:"address,omitempty"`
	// Número do endereço
	AddressNumber string `json:"addressNumber,omitempty"`
	// Complemento do endereço
	Complement string `json:"complement,omitempty"`
	// Bairro
	Province string `json:"province,omitempty"`
}

type DeleteWhiteLabelSubaccountRequest struct {
	// Motivo da remoção (REQUIRED)
	RemoveReason string `json:"removeReason,omitempty" validate:"required"`
}

type GetAccountStatementRequest struct {
	// Data inicial da lista
	StartDate *Date `json:"startDate,omitempty"`
	// Data final da lista
	FinishDate *Date `json:"finishDate,omitempty"`
	// Elemento inicial da lista
	Offset int `json:"offset,omitempty"`
	// Número de elementos da lista (max: 100)
	Limit int `json:"limit,omitempty"`
	// Ordenação do resultado
	Order Order `json:"order,omitempty"`
}

type GetPaymentStatisticRequest struct {
	// Filtrar pelo Identificador único do cliente
	Customer string `json:"customer,omitempty"`
	// Filtrar por forma de pagamento
	BillingType BillingType `json:"billingType,omitempty"`
	// Filtrar por status
	Status ChargeStatus `json:"status,omitempty"`
	// Filtrar registros antecipados ou não
	Anticipated bool `json:"anticipated,omitempty"`
	// Filtrar a partir da data de vencimento inicial
	DueDateGe *Date `json:"dueDate[ge],omitempty"`
	// Filtrar a partir da data de vencimento final
	DueDateLe *Date `json:"dueDate[le],omitempty"`
	// Filtrar a partir da data de criação inicial
	DateCreatedGe *Date `json:"dateCreated[ge],omitempty"`
	// Filtrar a partir da data de criação final
	DateCreatedLe *Date `json:"dateCreated[le],omitempty"`
	// Filtrar a partir da data estimada de crédito inicial
	EstimatedCreditDateGe *Date `json:"estimatedCreditDate[ge],omitempty"`
	// Filtrar a partir da data estimada de crédito final
	EstimatedCreditDateLe *Date `json:"estimatedCreditDate[le],omitempty"`
	// Filtrar pelo Identificador do seu sistema
	ExternalReference string `json:"externalReference,omitempty"`
}

type AccountStatementResponse struct {
	Id                   string           `json:"id,omitempty"`
	PaymentId            string           `json:"paymentId,omitempty"`
	TransferId           string           `json:"transferId,omitempty"`
	BillId               string           `json:"billId,omitempty"`
	InvoiceId            string           `json:"invoiceId,omitempty"`
	PaymentDunningId     string           `json:"paymentDunningId,omitempty"`
	CreditBureauReportId string           `json:"creditBureauReportId,omitempty"`
	Type                 FinanceTransType `json:"type,omitempty"`
	Value                float64          `json:"value,omitempty"`
	Balance              float64          `json:"balance,omitempty"`
	Date                 Date             `json:"date,omitempty"`
	Description          string           `json:"description,omitempty"`
}

type AccountBalanceResponse struct {
	Balance float64         `json:"balance"`
	Errors  []ErrorResponse `json:"errors,omitempty"`
}

type PaymentStatisticResponse struct {
	Quantity int             `json:"quantity"`
	Value    float64         `json:"value"`
	NetValue float64         `json:"netValue"`
	Errors   []ErrorResponse `json:"errors,omitempty"`
}

type SplitStatisticResponse struct {
	Income  float64         `json:"income"`
	Outcome float64         `json:"outcome"`
	Errors  []ErrorResponse `json:"errors,omitempty"`
}

type AccountResponse struct {
	Name          string          `json:"name,omitempty"`
	BirthDate     *Date           `json:"birthDate,omitempty"`
	CpfCnpj       string          `json:"cpfCnpj,omitempty"`
	Email         string          `json:"email,omitempty"`
	Phone         string          `json:"phone,omitempty"`
	Status        AccountStatus   `json:"status,omitempty"`
	PersonType    PersonType      `json:"personType,omitempty"`
	MobilePhone   string          `json:"mobilePhone,omitempty"`
	Site          string          `json:"site,omitempty"`
	CompanyName   string          `json:"companyName,omitempty"`
	CompanyType   CompanyType     `json:"companyType,omitempty"`
	PostalCode    string          `json:"postalCode,omitempty"`
	Address       string          `json:"address,omitempty"`
	AddressNumber string          `json:"addressNumber,omitempty"`
	Complement    string          `json:"complement,omitempty"`
	Province      string          `json:"province,omitempty"`
	City          CityResponse    `json:"city,omitempty"`
	DenialReason  string          `json:"denialReason,omitempty"`
	Errors        []ErrorResponse `json:"errors,omitempty"`
}

type CityResponse struct {
	Id           int    `json:"id,omitempty"`
	IbgeCode     string `json:"ibgeCode,omitempty"`
	Name         string `json:"name,omitempty"`
	DistrictCode string `json:"districtCode,omitempty"`
	District     string `json:"district,omitempty"`
	State        string `json:"state,omitempty"`
}

type InvoiceCustomizationResponse struct {
	LogoUrl             string          `json:"logoFile,omitempty"`
	LogoBackgroundColor string          `json:"logoBackgroundColor,omitempty"`
	InfoBackgroundColor string          `json:"infoBackgroundColor,omitempty"`
	FontColor           string          `json:"fontColor,omitempty"`
	Enabled             bool            `json:"enabled,omitempty"`
	Status              string          `json:"status,omitempty"`
	Observations        string          `json:"observations,omitempty"`
	Errors              []ErrorResponse `json:"errors,omitempty"`
}

type AccountBankInfoResponse struct {
	Agency       string          `json:"agency,omitempty"`
	Account      string          `json:"account,omitempty"`
	AccountDigit string          `json:"accountDigit,omitempty"`
	Errors       []ErrorResponse `json:"errors,omitempty"`
}

type AccountFeesResponse struct {
	Payment            AccountPaymentFeesResponse      `json:"payment"`
	Transfer           AccountTransferFeesResponse     `json:"transfer"`
	Notification       AccountNotificationFeesResponse `json:"notification"`
	CreditBureauReport AccountCreditBureauFeesResponse `json:"creditBureauReport"`
	Invoice            AccountInvoiceFeesResponse      `json:"invoice"`
	Anticipation       AccountAnticipationFeesResponse `json:"anticipation"`
	Errors             []ErrorResponse                 `json:"errors,omitempty"`
}

type AccountPaymentFeesResponse struct {
	BankSlip   AccountBankSlipFeesResponse   `json:"bankSlip"`
	CreditCard AccountCreditCardFeesResponse `json:"creditCard"`
	DebitCard  AccountDebitCardFeesResponse  `json:"debitCard"`
	Pix        AccountPixPaymentFeesResponse `json:"pix"`
}

type AccountTransferFeesResponse struct {
	MonthlyTransfersWithoutFee int                            `json:"monthlyTransfersWithoutFee"`
	Ted                        AccountTedTransferFeesResponse `json:"ted"`
	Pix                        AccountPixTransferFeesResponse `json:"pix"`
}

type AccountBankSlipFeesResponse struct {
	DefaultValue   float64 `json:"defaultValue"`
	DiscountValue  float64 `json:"discountValue"`
	ExpirationDate Date    `json:"expirationDate,omitempty"`
}

type AccountCreditCardFeesResponse struct {
	OperationValue                           float64 `json:"operationValue"`
	OneInstallmentPercentage                 float64 `json:"oneInstallmentPercentage"`
	UpToSixInstallmentsPercentage            float64 `json:"upToSixInstallmentsPercentage"`
	UpToTwelveInstallmentsPercentage         float64 `json:"upToTwelveInstallmentsPercentage"`
	DiscountOneInstallmentPercentage         float64 `json:"discountOneInstallmentPercentage"`
	DiscountUpToSixInstallmentsPercentage    float64 `json:"discountUpToSixInstallmentsPercentage"`
	DiscountUpToTwelveInstallmentsPercentage float64 `json:"discountUpToTwelveInstallmentsPercentage"`
	DiscountExpiration                       float64 `json:"discountExpiration"`
}

type AccountDebitCardFeesResponse struct {
	OperationValue    float64 `json:"operationValue"`
	DefaultPercentage float64 `json:"defaultPercentage"`
}

type AccountPixPaymentFeesResponse struct {
	FixedFeeValue                 float64 `json:"fixedFeeValue"`
	FixedFeeValueWithDiscount     float64 `json:"fixedFeeValueWithDiscount"`
	PercentageFee                 float64 `json:"percentageFee"`
	MinimumFeeValue               float64 `json:"minimumFeeValue"`
	MaximumFeeValue               float64 `json:"maximumFeeValue"`
	DiscountExpiration            Date    `json:"discountExpiration,omitempty"`
	MonthlyCreditsWithoutFee      float64 `json:"monthlyCreditsWithoutFee"`
	CreditsReceivedOfCurrentMonth float64 `json:"creditsReceivedOfCurrentMonth"`
}

type AccountTedTransferFeesResponse struct {
	FeeValue float64 `json:"feeValue"`
}

type AccountPixTransferFeesResponse struct {
	FeeValue       float64 `json:"feeValue"`
	DiscountValue  float64 `json:"discountValue"`
	ExpirationDate Date    `json:"expirationDate,omitempty"`
}

type AccountNotificationFeesResponse struct {
	PhoneCallFeeValue float64 `json:"phoneCallFeeValue"`
	WhatsAppFeeValue  float64 `json:"whatsAppFeeValue"`
	MessagingFeeValue float64 `json:"messagingFeeValue"`
}

type AccountCreditBureauFeesResponse struct {
	NaturalPersonFeeValue float64 `json:"naturalPersonFeeValue"`
	LegalPersonFeeValue   float64 `json:"legalPersonFeeValue"`
}

type AccountInvoiceFeesResponse struct {
	FeeValue float64 `json:"feeValue"`
}

type AccountAnticipationFeesResponse struct {
	CreditCard AccountAnticipationCreditCardFeesResponse `json:"creditCard"`
	BankSlip   AccountAnticipationBankSlipResponse       `json:"bankSlip"`
}

type AccountAnticipationCreditCardFeesResponse struct {
	DetachedMonthlyFeeValue    float64 `json:"detachedMonthlyFeeValue"`
	InstallmentMonthlyFeeValue float64 `json:"installmentMonthlyFeeValue"`
}

type AccountAnticipationBankSlipResponse struct {
	MonthlyFeePercentage float64 `json:"monthlyFeePercentage"`
}

type AccountRegistrationStatusResponse struct {
	Id             string          `json:"id,omitempty"`
	CommercialInfo string          `json:"commercialInfo,omitempty"`
	Documentation  string          `json:"documentation,omitempty"`
	General        string          `json:"general,omitempty"`
	Errors         []ErrorResponse `json:"errors,omitempty"`
}

type AccountWalletResponse struct {
	Id string `json:"id,omitempty"`
}

type DeleteWhiteLabelSubaccountResponse struct {
	Observations string          `json:"observations,omitempty"`
	Errors       []ErrorResponse `json:"errors,omitempty"`
}

type account struct {
	env         Env
	accessToken string
}

type Account interface {
	// SaveInvoiceCustomization (Salvar personalização da fatura)
	//
	// Possibilita personalizar a fatura apresentada ao seu cliente com o logo e cores da sua empresa.
	// Após salva, a personalização é analisada e aprovada pela nossa equipe dentro de algumas horas.
	//
	// # Resposta: 200
	//
	// InvoiceCustomizationResponse = not nil
	//
	// Error = nil
	//
	// InvoiceCustomizationResponse.IsSuccess() = true
	//
	// Possui os valores de resposta de sucesso segunda a documentação.
	//
	// # Resposta: 400/401/500
	//
	// InvoiceCustomizationResponse = not nil
	//
	// Error = nil
	//
	// InvoiceCustomizationResponse.IsFailure() = true
	//
	// Para qualquer outra resposta inesperada da API, possuímos o campo InvoiceCustomizationResponse.Errors preenchido com as informações
	// de erro, sendo 400 retornado da API Asaas com as instruções de requisição conforme a documentação,
	// diferente disso retornará uma mensagem padrão no index 0 do slice com campo ErrorResponse.Code retornando a
	// descrição status http (Ex: "401 Unauthorized") e no campo ErrorResponse.Description retornará com o valor
	// "response status code not expected".
	//
	// # Error
	//
	// InvoiceCustomizationResponse = nil
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
	// Salvar personalização da fatura: https://docs.asaas.com/reference/salvar-personalizacao-da-fatura
	SaveInvoiceCustomization(ctx context.Context, body SaveInvoiceCustomizationRequest) (
		*InvoiceCustomizationResponse, Error)
	// Update (Atualizar dados comerciais)
	//
	// Dependendo das informações alteradas é possível que sua conta passe por uma nova análise, o que ocasionará
	// em um bloqueio temporário de algumas funcionalidades do sistema.
	//
	// # Resposta: 200
	//
	// AccountResponse = not nil
	//
	// Error = nil
	//
	// AccountResponse.IsSuccess() = true
	//
	// Possui os valores de resposta de sucesso segunda a documentação.
	//
	// # Resposta: 400/401/500
	//
	// AccountResponse = not nil
	//
	// Error = nil
	//
	// AccountResponse.IsFailure() = true
	//
	// Para qualquer outra resposta inesperada da API, possuímos o campo AccountResponse.Errors preenchido com as informações
	// de erro, sendo 400 retornado da API Asaas com as instruções de requisição conforme a documentação,
	// diferente disso retornará uma mensagem padrão no index 0 do slice com campo ErrorResponse.Code retornando a
	// descrição status http (Ex: "401 Unauthorized") e no campo ErrorResponse.Description retornará com o valor
	// "response status code not expected".
	//
	// # Error
	//
	// AccountResponse = nil
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
	// Atualizar dados comerciais: https://docs.asaas.com/reference/atualizar-dados-comerciais
	Update(ctx context.Context, body UpdateAccountRequest) (*AccountResponse, Error)
	// DeleteWhiteLabelSubaccount (Excluir subconta White Label)
	//
	// Ao excluir uma subconta no Asaas, ela perderá o acesso a todas as funcionalidades e todos os seus dados serão
	// removidos, incluindo cobranças, clientes e documentos.
	//
	// Não será possível recuperar a conta após o cancelamento.
	//
	// # Resposta: 200
	//
	// DeleteWhiteLabelSubaccountResponse = not nil
	//
	// Error = nil
	//
	// Se DeleteWhiteLabelSubaccountResponse.IsSuccess() for true quer dizer que foi excluída.
	//
	// Se caso DeleteWhiteLabelSubaccountResponse.IsFailure() for true quer dizer que não foi excluída.
	//
	// # Resposta: 400/401/500
	//
	// DeleteWhiteLabelSubaccountResponse = not nil
	//
	// Error = nil
	//
	// DeleteWhiteLabelSubaccountResponse.IsFailure() = true
	//
	// Para qualquer outra resposta inesperada da API, possuímos o campo DeleteWhiteLabelSubaccountResponse.Errors preenchido com as informações
	// de erro, sendo 400 retornado da API Asaas com as instruções de requisição conforme a documentação,
	// diferente disso retornará uma mensagem padrão no index 0 do slice com campo ErrorResponse.Code retornando a
	// descrição status http (Ex: "401 Unauthorized") e no campo ErrorResponse.Description retornará com o valor
	// "response status code not expected".
	//
	// # Error
	//
	// DeleteWhiteLabelSubaccountResponse = nil
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
	// Excluir subconta White Label: https://docs.asaas.com/reference/excluir-subconta-white-label
	DeleteWhiteLabelSubaccount(ctx context.Context, body DeleteWhiteLabelSubaccountRequest) (
		*DeleteWhiteLabelSubaccountResponse, Error)
	// Get (Recuperar dados comerciais)
	//
	// # Resposta: 200
	//
	// AccountResponse = not nil
	//
	// Error = nil
	//
	// AccountResponse.IsSuccess() = true
	//
	// Possui os valores de resposta de sucesso segunda a documentação.
	//
	// # Resposta: 401/500
	//
	// AccountResponse = not nil
	//
	// Error = nil
	//
	// AccountResponse.IsFailure() = true
	//
	// Para qualquer outra resposta inesperada da API, possuímos o campo AccountResponse.Errors preenchido com as informações
	// de erro, o index 0 do slice com campo ErrorResponse.Code retornando a descrição
	// status http (Ex: "401 Unauthorized") e no campo ErrorResponse.Description retornará com o valor
	// "response status code not expected".
	//
	// # Error
	//
	// AccountResponse = nil
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
	// Recuperar dados comerciais: https://docs.asaas.com/reference/recuperar-dados-comerciais
	Get(ctx context.Context) (*AccountResponse, Error)
	// GetRegistrationStatus (Recuperar taxas da conta)
	//
	// Com este endpoint é possível verificar os status de aprovação da conta.
	//
	// # Resposta: 200
	//
	// AccountRegistrationStatusResponse = not nil
	//
	// Error = nil
	//
	// AccountRegistrationStatusResponse.IsSuccess() = true
	//
	// Possui os valores de resposta de sucesso segunda a documentação.
	//
	// # Resposta: 401/500
	//
	// AccountRegistrationStatusResponse = not nil
	//
	// Error = nil
	//
	// AccountRegistrationStatusResponse.IsFailure() = true
	//
	// Para qualquer outra resposta inesperada da API, possuímos o campo AccountRegistrationStatusResponse.Errors preenchido com as informações
	// de erro, o index 0 do slice com campo ErrorResponse.Code retornando a descrição
	// status http (Ex: "401 Unauthorized") e no campo ErrorResponse.Description retornará com o valor
	// "response status code not expected".
	//
	// # Error
	//
	// AccountRegistrationStatusResponse = nil
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
	// Consultar situação cadastral da conta: https://docs.asaas.com/reference/consultar-situacao-cadastral-da-conta
	GetRegistrationStatus(ctx context.Context) (*AccountRegistrationStatusResponse, Error)
	// GetBankInfo (Recuperar número de conta no Asaas)
	//
	// # Resposta: 200
	//
	// AccountBankInfoResponse = not nil
	//
	// Error = nil
	//
	// AccountBankInfoResponse.IsSuccess() = true
	//
	// Possui os valores de resposta de sucesso segunda a documentação.
	//
	// # Resposta: 401/500
	//
	// AccountBankInfoResponse = not nil
	//
	// Error = nil
	//
	// AccountBankInfoResponse.IsFailure() = true
	//
	// Para qualquer outra resposta inesperada da API, possuímos o campo AccountBankInfoResponse.Errors preenchido com as informações
	// de erro, o index 0 do slice com campo ErrorResponse.Code retornando a descrição
	// status http (Ex: "401 Unauthorized") e no campo ErrorResponse.Description retornará com o valor
	// "response status code not expected".
	//
	// # Error
	//
	// AccountBankInfoResponse = nil
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
	// Recuperar número de conta no Asaas: https://docs.asaas.com/reference/recuperar-numero-de-conta-no-asaas
	GetBankInfo(ctx context.Context) (*AccountBankInfoResponse, Error)
	// GetFees (Recuperar taxas da conta)
	//
	// Através deste endpoint é possível verificar as taxas aplicadas na conta.
	//
	// # Resposta: 200
	//
	// AccountFeesResponse = not nil
	//
	// Error = nil
	//
	// AccountFeesResponse.IsSuccess() = true
	//
	// Possui os valores de resposta de sucesso segunda a documentação.
	//
	// # Resposta: 401/500
	//
	// AccountFeesResponse = not nil
	//
	// Error = nil
	//
	// AccountFeesResponse.IsFailure() = true
	//
	// Para qualquer outra resposta inesperada da API, possuímos o campo AccountFeesResponse.Errors preenchido com as informações
	// de erro, o index 0 do slice com campo ErrorResponse.Code retornando a descrição
	// status http (Ex: "401 Unauthorized") e no campo ErrorResponse.Description retornará com o valor
	// "response status code not expected".
	//
	// # Error
	//
	// AccountFeesResponse = nil
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
	// Recuperar taxas da conta: https://docs.asaas.com/reference/recuperar-taxas-da-conta
	GetFees(ctx context.Context) (*AccountFeesResponse, Error)
	// GetWallets (Recuperar uma lista de Wallets)
	//
	// Através deste endpoint é possível recuperar uma lista de WalletId de uma conta caso você já tenha uma conta
	// criada, mas não armazenou o WalletId.
	//
	// # Resposta: 200
	//
	// Pageable(AccountWalletResponse) = not nil
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
	// Pageable(AccountWalletResponse) = not nil
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
	// Pageable(AccountWalletResponse) = nil
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
	// Recuperar WalletId: https://docs.asaas.com/reference/recuperar-walletid
	GetWallets(ctx context.Context) (*Pageable[AccountWalletResponse], Error)
	// GetBalance (Recuperar saldo da conta)
	//
	// # Resposta: 200
	//
	// AccountBalanceResponse = not nil
	//
	// Error = nil
	//
	// AccountBalanceResponse.IsSuccess() = true
	//
	// Possui os valores de resposta de sucesso segunda a documentação.
	//
	// # Resposta: 401/500
	//
	// AccountBalanceResponse = not nil
	//
	// Error = nil
	//
	// AccountBalanceResponse.IsFailure() = true
	//
	// Para qualquer outra resposta inesperada da API, possuímos o campo AccountBalanceResponse.Errors preenchido com as informações
	// de erro, o index 0 do slice com campo ErrorResponse.Code retornando a descrição
	// status http (Ex: "401 Unauthorized") e no campo ErrorResponse.Description retornará com o valor
	// "response status code not expected".
	//
	// # Error
	//
	// AccountBalanceResponse = nil
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
	// Recuperar saldo da conta: https://docs.asaas.com/reference/recuperar-saldo-da-conta
	GetBalance(ctx context.Context) (*AccountBalanceResponse, Error)
	// GetAccountStatement (Recuperar extrato)
	//
	// Retorna uma lista de movimentações financeiras no período informado nos parâmetros.
	//
	// # Resposta: 200
	//
	// Pageable(AccountStatementResponse) = not nil
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
	// Pageable(AccountStatementResponse) = not nil
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
	// Pageable(AccountStatementResponse) = nil
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
	// Recuperar extrato: https://docs.asaas.com/reference/recuperar-extrato
	GetAccountStatement(ctx context.Context, filter GetAccountStatementRequest) (
		*Pageable[AccountStatementResponse], Error)
	// GetPaymentStatistic (Estatísticas de cobranças)
	//
	// # Resposta: 200
	//
	// PaymentStatisticResponse = not nil
	//
	// Error = nil
	//
	// PaymentStatisticResponse.IsSuccess() = true
	//
	// Possui os valores de resposta de sucesso segunda a documentação.
	//
	// # Resposta: 401/500
	//
	// PaymentStatisticResponse = not nil
	//
	// Error = nil
	//
	// PaymentStatisticResponse.IsFailure() = true
	//
	// Para qualquer outra resposta inesperada da API, possuímos o campo PaymentStatisticResponse.Errors preenchido com as informações
	// de erro, o index 0 do slice com campo ErrorResponse.Code retornando a descrição
	// status http (Ex: "401 Unauthorized") e no campo ErrorResponse.Description retornará com o valor
	// "response status code not expected".
	//
	// # Error
	//
	// PaymentStatisticResponse = nil
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
	// Estatísticas de cobranças: https://docs.asaas.com/reference/estatisticas-de-cobran%C3%A7as
	GetPaymentStatistic(ctx context.Context, filter GetPaymentStatisticRequest) (*PaymentStatisticResponse, Error)
	// GetSplitStatistic (Recuperar valores de split)
	//
	// # Resposta: 200
	//
	// SplitStatisticResponse = not nil
	//
	// Error = nil
	//
	// SplitStatisticResponse.IsSuccess() = true
	//
	// Possui os valores de resposta de sucesso segunda a documentação.
	//
	// # Resposta: 401/500
	//
	// SplitStatisticResponse = not nil
	//
	// Error = nil
	//
	// SplitStatisticResponse.IsFailure() = true
	//
	// Para qualquer outra resposta inesperada da API, possuímos o campo SplitStatisticResponse.Errors preenchido com as informações
	// de erro, o index 0 do slice com campo ErrorResponse.Code retornando a descrição
	// status http (Ex: "401 Unauthorized") e no campo ErrorResponse.Description retornará com o valor
	// "response status code not expected".
	//
	// # Error
	//
	// SplitStatisticResponse = nil
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
	// Recuperar valores de split: https://docs.asaas.com/reference/recuperar-valores-de-split
	GetSplitStatistic(ctx context.Context) (*SplitStatisticResponse, Error)
	// GetInvoiceCustomization (Recuperar configurações de personalização)
	//
	// # Resposta: 200
	//
	// InvoiceCustomizationResponse = not nil
	//
	// Error = nil
	//
	// InvoiceCustomizationResponse.IsSuccess() = true
	//
	// Possui os valores de resposta de sucesso segunda a documentação.
	//
	// # Resposta: 401/500
	//
	// InvoiceCustomizationResponse = not nil
	//
	// Error = nil
	//
	// InvoiceCustomizationResponse.IsFailure() = true
	//
	// Para qualquer outra resposta inesperada da API, possuímos o campo InvoiceCustomizationResponse.Errors preenchido com as informações
	// de erro, o index 0 do slice com campo ErrorResponse.Code retornando a descrição
	// status http (Ex: "401 Unauthorized") e no campo ErrorResponse.Description retornará com o valor
	// "response status code not expected".
	//
	// # Error
	//
	// InvoiceCustomizationResponse = nil
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
	// Recuperar configurações de personalização: https://docs.asaas.com/reference/recuperar-configuracoes-de-personalizacao
	GetInvoiceCustomization(ctx context.Context) (*InvoiceCustomizationResponse, Error)
}

func NewAccount(env Env, accessToken string) Account {
	logWarning("Account service running on", env.String())
	return account{
		env:         env,
		accessToken: accessToken,
	}
}

func (a account) SaveInvoiceCustomization(ctx context.Context, body SaveInvoiceCustomizationRequest) (
	*InvoiceCustomizationResponse, Error) {
	if err := Validate().Struct(body); err != nil {
		return nil, NewError(ErrorTypeValidation, err)
	}
	req := NewRequest[InvoiceCustomizationResponse](ctx, a.env, a.accessToken)
	return req.makeMultipartForm(http.MethodPost, "/v3/myAccount/paymentCheckoutConfig", body)
}

func (a account) Update(ctx context.Context, body UpdateAccountRequest) (*AccountResponse, Error) {
	if err := Validate().Struct(body); err != nil {
		return nil, NewError(ErrorTypeValidation, err)
	}
	req := NewRequest[AccountResponse](ctx, a.env, a.accessToken)
	return req.make(http.MethodPut, "/v3/myAccount/commercialInfo", body)
}

func (a account) DeleteWhiteLabelSubaccount(ctx context.Context, body DeleteWhiteLabelSubaccountRequest) (
	*DeleteWhiteLabelSubaccountResponse, Error) {
	if err := Validate().Struct(body); err != nil {
		return nil, NewError(ErrorTypeValidation, err)
	}
	req := NewRequest[DeleteWhiteLabelSubaccountResponse](ctx, a.env, a.accessToken)
	return req.make(http.MethodDelete, "/v3/myAccount", body)
}

func (a account) Get(ctx context.Context) (*AccountResponse, Error) {
	req := NewRequest[AccountResponse](ctx, a.env, a.accessToken)
	return req.make(http.MethodGet, "/v3/myAccount/commercialInfo", nil)
}

func (a account) GetRegistrationStatus(ctx context.Context) (*AccountRegistrationStatusResponse, Error) {
	req := NewRequest[AccountRegistrationStatusResponse](ctx, a.env, a.accessToken)
	return req.make(http.MethodGet, "/v3/myAccount/status", nil)
}

func (a account) GetBankInfo(ctx context.Context) (*AccountBankInfoResponse, Error) {
	req := NewRequest[AccountBankInfoResponse](ctx, a.env, a.accessToken)
	return req.make(http.MethodGet, "/v3/myAccount/accountNumber", nil)
}

func (a account) GetFees(ctx context.Context) (*AccountFeesResponse, Error) {
	req := NewRequest[AccountFeesResponse](ctx, a.env, a.accessToken)
	return req.make(http.MethodGet, "/v3/myAccount/fees", nil)
}

func (a account) GetWallets(ctx context.Context) (*Pageable[AccountWalletResponse], Error) {
	req := NewRequest[Pageable[AccountWalletResponse]](ctx, a.env, a.accessToken)
	return req.make(http.MethodGet, "/v3/wallets", nil)
}

func (a account) GetBalance(ctx context.Context) (*AccountBalanceResponse, Error) {
	req := NewRequest[AccountBalanceResponse](ctx, a.env, a.accessToken)
	return req.make(http.MethodGet, "/v3/finance/balance", nil)
}

func (a account) GetSplitStatistic(ctx context.Context) (*SplitStatisticResponse, Error) {
	req := NewRequest[SplitStatisticResponse](ctx, a.env, a.accessToken)
	return req.make(http.MethodGet, "/v3/finance/split/statistics", nil)
}

func (a account) GetAccountStatement(ctx context.Context, filter GetAccountStatementRequest) (
	*Pageable[AccountStatementResponse], Error) {
	req := NewRequest[Pageable[AccountStatementResponse]](ctx, a.env, a.accessToken)
	return req.make(http.MethodGet, "/v3/financialTransactions", filter)
}

func (a account) GetPaymentStatistic(ctx context.Context, filter GetPaymentStatisticRequest) (
	*PaymentStatisticResponse, Error) {
	req := NewRequest[PaymentStatisticResponse](ctx, a.env, a.accessToken)
	return req.make(http.MethodGet, "/v3/payment/statistics", filter)
}

func (a account) GetInvoiceCustomization(ctx context.Context) (*InvoiceCustomizationResponse, Error) {
	req := NewRequest[InvoiceCustomizationResponse](ctx, a.env, a.accessToken)
	return req.make(http.MethodGet, "/v3/myAccount/paymentCheckoutConfig", nil)
}
