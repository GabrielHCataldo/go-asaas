package asaas

import (
	"context"
	"testing"
	"time"
)

func TestSubscriptionCreate(t *testing.T) {
	initCustomer(false)
	accessToken := getEnvValue(EnvAccessToken)
	customerId := getEnvValue(EnvCustomerId)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	now := DateNow()
	nSubscription := NewSubscription(EnvSandbox, accessToken)
	resp, err := nSubscription.Create(ctx, CreateSubscriptionRequest{
		Customer:             customerId,
		BillingType:          BillingTypeBankSlip,
		Value:                5,
		NextDueDate:          NewDate(now.Year(), now.Month()+1, now.Day(), now.Location()),
		Discount:             nil,
		Interest:             nil,
		Fine:                 nil,
		Cycle:                SubscriptionCycleMonthly,
		Description:          "Unit test go",
		CreditCard:           nil,
		CreditCardHolderInfo: nil,
		CreditCardToken:      "",
		EndDate:              Date{},
		MaxPayments:          0,
		ExternalReference:    "",
		Split:                nil,
		RemoteIp:             "",
	})
	assertResponseSuccess(t, resp, err)
}

func TestSubscriptionCreateInvoiceSettingById(t *testing.T) {
	initSubscription()
	accessToken := getEnvValue(EnvAccessToken)
	subscriptionId := getEnvValue(EnvSubscriptionId)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nSubscription := NewSubscription(EnvSandbox, accessToken)
	resp, err := nSubscription.CreateInvoiceSettingById(ctx, subscriptionId, CreateInvoiceSettingRequest{
		MunicipalServiceId:   "",
		MunicipalServiceCode: "123",
		MunicipalServiceName: "Unit test go",
		UpdatePayment:        false,
		Deductions:           0,
		EffectiveDatePeriod:  EffectiveDatePeriodOnNextMonth,
		ReceivedOnly:         false,
		DaysBeforeDueDate:    0,
		Observations:         "Unit test go",
		Taxes:                InvoiceTaxesRequest{},
	})
	assertResponseSuccess(t, resp, err)
}

func TestSubscriptionUpdateById(t *testing.T) {
	initSubscription()
	accessToken := getEnvValue(EnvAccessToken)
	subscriptionId := getEnvValue(EnvSubscriptionId)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nSubscription := NewSubscription(EnvSandbox, accessToken)
	resp, err := nSubscription.UpdateById(ctx, subscriptionId, UpdateSubscriptionRequest{
		BillingType:           "",
		Value:                 0,
		Status:                SubscriptionStatusInactive,
		NextDueDate:           Date{},
		Discount:              nil,
		Interest:              nil,
		Fine:                  nil,
		Cycle:                 SubscriptionCycleBimonthly,
		Description:           Pointer("Unit test go"),
		EndDate:               Date{},
		UpdatePendingPayments: true,
		ExternalReference:     "",
	})
	assertResponseSuccess(t, resp, err)
}

func TestSubscriptionUpdateInvoiceSettingsById(t *testing.T) {
	initSubscriptionInvoiceSetting()
	accessToken := getEnvValue(EnvAccessToken)
	subscriptionId := getEnvValue(EnvSubscriptionId)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nSubscription := NewSubscription(EnvSandbox, accessToken)
	resp, err := nSubscription.UpdateInvoiceSettingsById(ctx, subscriptionId, UpdateInvoiceSettingRequest{
		Deductions:          nil,
		EffectiveDatePeriod: EffectiveDatePeriodOnDueDateMonth,
		ReceivedOnly:        nil,
		DaysBeforeDueDate:   nil,
		Observations:        nil,
		Taxes:               InvoiceTaxesRequest{},
	})
	assertResponseSuccess(t, resp, err)
}

func TestSubscriptionDeleteById(t *testing.T) {
	initSubscription()
	accessToken := getEnvValue(EnvAccessToken)
	subscriptionId := getEnvValue(EnvSubscriptionId)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nSubscription := NewSubscription(EnvSandbox, accessToken)
	resp, err := nSubscription.DeleteById(ctx, subscriptionId)
	assertResponseSuccess(t, resp, err)
}

func TestSubscriptionDeleteInvoiceSettingById(t *testing.T) {
	initSubscriptionInvoiceSetting()
	accessToken := getEnvValue(EnvAccessToken)
	subscriptionId := getEnvValue(EnvSubscriptionId)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nSubscription := NewSubscription(EnvSandbox, accessToken)
	resp, err := nSubscription.DeleteInvoiceSettingById(ctx, subscriptionId)
	assertResponseSuccess(t, resp, err)
}

func TestSubscriptionGetById(t *testing.T) {
	initSubscription()
	accessToken := getEnvValue(EnvAccessToken)
	subscriptionId := getEnvValue(EnvSubscriptionId)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nSubscription := NewSubscription(EnvSandbox, accessToken)
	resp, err := nSubscription.GetById(ctx, subscriptionId)
	assertResponseSuccess(t, resp, err)
}

func TestSubscriptionGetInvoiceSettingById(t *testing.T) {
	initSubscriptionInvoiceSetting()
	accessToken := getEnvValue(EnvAccessToken)
	subscriptionId := getEnvValue(EnvSubscriptionId)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nSubscription := NewSubscription(EnvSandbox, accessToken)
	resp, err := nSubscription.GetInvoiceSettingById(ctx, subscriptionId)
	assertResponseSuccess(t, resp, err)
}

func TestSubscriptionGetPaymentBookById(t *testing.T) {
	initSubscription()
	accessToken := getEnvValue(EnvAccessToken)
	subscriptionId := getEnvValue(EnvSubscriptionId)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	now := DateNow()
	nSubscription := NewSubscription(EnvSandbox, accessToken)
	resp, err := nSubscription.GetPaymentBookById(ctx, subscriptionId, SubscriptionPaymentBookRequest{
		Month: int(now.Month()),
		Year:  now.Year(),
		Sort:  "",
		Order: "",
	})
	assertResponseFailure(resp, err)
}

func TestSubscriptionGetAll(t *testing.T) {
	initSubscription()
	accessToken := getEnvValue(EnvAccessToken)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nSubscription := NewSubscription(EnvSandbox, accessToken)
	resp, err := nSubscription.GetAll(ctx, GetAllSubscriptionsRequest{
		Customer:          "",
		CustomerGroupName: "",
		BillingType:       "",
		Status:            "",
		DeletedOnly:       false,
		IncludeDeleted:    false,
		ExternalReference: "",
		Order:             "",
		Sort:              "",
		Offset:            0,
		Limit:             10,
	})
	assertResponseSuccess(t, resp, err)
}

func TestSubscriptionGetAllChargesBySubscription(t *testing.T) {
	initSubscription()
	accessToken := getEnvValue(EnvAccessToken)
	subscriptionId := getEnvValue(EnvSubscriptionId)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nSubscription := NewSubscription(EnvSandbox, accessToken)
	resp, err := nSubscription.GetAllChargesBySubscription(ctx, subscriptionId, GetAllChargesBySubscriptionRequest{
		Status: "",
	})
	assertResponseSuccess(t, resp, err)
}

func TestSubscriptionGetAllInvoicesBySubscription(t *testing.T) {
	initSubscription()
	accessToken := getEnvValue(EnvAccessToken)
	subscriptionId := getEnvValue(EnvSubscriptionId)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nSubscription := NewSubscription(EnvSandbox, accessToken)
	resp, err := nSubscription.GetAllInvoicesBySubscription(ctx, subscriptionId, GetAllSubscriptionInvoicesRequest{
		EffectiveDateGE:   Date{},
		EffectiveDateLE:   Date{},
		ExternalReference: "",
		Status:            "",
		Customer:          "",
		Offset:            0,
		Limit:             10,
	})
	assertResponseNoContent(resp, err)
}
