package asaas

import (
	"context"
	"net/http"
	"os"
)

type CreateInvoiceCustomizationRequest struct {
	LogoBackgroundColor string   `json:"logoBackgroundColor,omitempty" validate:"required"`
	InfoBackgroundColor string   `json:"infoBackgroundColor,omitempty" validate:"required"`
	FontColor           string   `json:"fontColor,omitempty" validate:"required"`
	Enabled             bool     `json:"enabled,omitempty"`
	LogoFile            *os.File `json:"logoFile,omitempty"`
}

type UpdateAccountRequest struct {
	PersonType    PersonType  `json:"personType,omitempty" validate:"omitempty,enum"`
	CpfCnpj       string      `json:"cpfCnpj,omitempty" validate:"omitempty,document"`
	BirthDate     Date        `json:"birthDate,omitempty" validate:"omitempty,before_now"`
	CompanyType   CompanyType `json:"companyType,omitempty" validate:"omitempty,enum"`
	Email         string      `json:"email,omitempty" validate:"omitempty,email"`
	Phone         string      `json:"phone,omitempty" validate:"omitempty,phone"`
	MobilePhone   string      `json:"mobilePhone,omitempty" validate:"omitempty,phone"`
	Site          string      `json:"site,omitempty" validate:"omitempty,url"`
	PostalCode    string      `json:"postalCode,omitempty" validate:"omitempty,postal_code"`
	Address       string      `json:"address,omitempty"`
	AddressNumber string      `json:"addressNumber,omitempty"`
	Complement    string      `json:"complement,omitempty"`
	Province      string      `json:"province,omitempty"`
}

type DeleteWhiteLabelSubaccountRequest struct {
	RemoveReason string `json:"removeReason,omitempty" validate:"required"`
}

type GetFinancialTransactionsRequest struct {
	StartDate  *Date `json:"startDate,omitempty"`
	FinishDate *Date `json:"finishDate,omitempty"`
	Offset     int   `json:"offset,omitempty"`
	Limit      int   `json:"limit,omitempty"`
	Order      Order `json:"order,omitempty"`
}

type FinancialTransactionsResponse struct {
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

type GetPaymentStatisticRequest struct {
	Customer              string       `json:"customer,omitempty"`
	BillingType           BillingType  `json:"billingType,omitempty"`
	Status                ChargeStatus `json:"status,omitempty"`
	Anticipated           bool         `json:"anticipated,omitempty"`
	DueDateGe             *Date        `json:"dueDate[ge],omitempty"`
	DueDateLe             *Date        `json:"dueDate[le],omitempty"`
	DateCreatedGe         *Date        `json:"dateCreated[ge],omitempty"`
	DateCreatedLe         *Date        `json:"dateCreated[le],omitempty"`
	EstimatedCreditDateGe *Date        `json:"estimatedCreditDate[ge],omitempty"`
	EstimatedCreditDateLe *Date        `json:"estimatedCreditDate[le],omitempty"`
	ExternalReference     string       `json:"externalReference,omitempty"`
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

type (
	Account interface {
		CreateInvoiceCustomization(ctx context.Context, body CreateInvoiceCustomizationRequest) (
			*InvoiceCustomizationResponse, Error)
		Update(ctx context.Context, body UpdateAccountRequest) (*AccountResponse, Error)
		DeleteWhiteLabelSubaccount(ctx context.Context, body DeleteWhiteLabelSubaccountRequest) (
			*DeleteWhiteLabelSubaccountResponse, Error)
		Get(ctx context.Context) (*AccountResponse, Error)
		GetRegistrationStatus(ctx context.Context) (*AccountRegistrationStatusResponse, Error)
		GetBankInfo(ctx context.Context) (*AccountBankInfoResponse, Error)
		GetFees(ctx context.Context) (*AccountFeesResponse, Error)
		GetWallets(ctx context.Context) (*Pageable[AccountWalletResponse], Error)
		GetBalance(ctx context.Context) (*AccountBalanceResponse, Error)
		GetSplitStatistic(ctx context.Context) (*SplitStatisticResponse, Error)
		GetAccountStatement(ctx context.Context, filter GetFinancialTransactionsRequest) (
			*Pageable[FinancialTransactionsResponse], Error)
		GetPaymentStatistic(ctx context.Context, filter GetPaymentStatisticRequest) (*PaymentStatisticResponse, Error)
		GetInvoiceCustomization(ctx context.Context) (*InvoiceCustomizationResponse, Error)
	}
)

func NewAccount(env Env, accessToken string) Account {
	logWarning("Account service running on", env.String())
	return account{
		env:         env,
		accessToken: accessToken,
	}
}

func (a account) CreateInvoiceCustomization(ctx context.Context, body CreateInvoiceCustomizationRequest) (
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

func (a account) GetAccountStatement(ctx context.Context, filter GetFinancialTransactionsRequest) (
	*Pageable[FinancialTransactionsResponse], Error) {
	req := NewRequest[Pageable[FinancialTransactionsResponse]](ctx, a.env, a.accessToken)
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
