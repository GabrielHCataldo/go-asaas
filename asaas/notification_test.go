package asaas

import (
	"context"
	"testing"
	"time"
)

func TestNotificationUpdateById(t *testing.T) {
	accessToken := getEnvValue(EnvAccessToken)
	assertFatalStringBlank(t, accessToken)
	initNotification()
	notificationId := getEnvValue(EnvNotificationId)
	assertFatalStringBlank(t, notificationId)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nNotification := NewNotification(EnvSandbox, accessToken)
	resp, err := nNotification.UpdateById(ctx, notificationId, UpdateNotificationRequest{
		Enabled:                     Pointer(true),
		EmailEnabledForProvider:     nil,
		SmsEnabledForProvider:       nil,
		EmailEnabledForCustomer:     nil,
		SmsEnabledForCustomer:       nil,
		PhoneCallEnabledForCustomer: nil,
		WhatsappEnabledForCustomer:  nil,
		ScheduleOffset:              nil,
	})
	assertResponseSuccess(t, resp, err)
}

func TestNotificationUpdateManyByCustomer(t *testing.T) {
	accessToken := getEnvValue(EnvAccessToken)
	assertFatalStringBlank(t, accessToken)
	initNotification()
	customerId := getEnvValue(EnvCustomerId)
	assertFatalStringBlank(t, customerId)
	notificationId := getEnvValue(EnvNotificationId)
	assertFatalStringBlank(t, notificationId)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nNotification := NewNotification(EnvSandbox, accessToken)
	resp, err := nNotification.UpdateManyByCustomer(ctx, UpdateManyNotificationsRequest{
		Customer: customerId,
		Notifications: []UpdateManyNotificationRequest{
			{
				Id:                          notificationId,
				Enabled:                     Pointer(true),
				EmailEnabledForProvider:     nil,
				SmsEnabledForProvider:       nil,
				EmailEnabledForCustomer:     nil,
				SmsEnabledForCustomer:       nil,
				PhoneCallEnabledForCustomer: nil,
				WhatsappEnabledForCustomer:  nil,
				ScheduleOffset:              nil,
			},
		},
	})
	assertResponseSuccess(t, resp, err)
}

func TestNotificationGetAllByCustomer(t *testing.T) {
	accessToken := getEnvValue(EnvAccessToken)
	assertFatalStringBlank(t, accessToken)
	initCreditCardCharge(false, false)
	customerId := getEnvValue(EnvCustomerId)
	assertFatalStringBlank(t, customerId)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nNotification := NewNotification(EnvSandbox, accessToken)
	resp, err := nNotification.GetAllByCustomer(ctx, customerId)
	assertResponseSuccess(t, resp, err)
}
