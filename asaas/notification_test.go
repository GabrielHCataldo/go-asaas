package asaas

import (
	"context"
	"testing"
	"time"
)

func TestNotificationUpdateById(t *testing.T) {
	initNotification()
	accessToken := getEnvValue(EnvAccessToken)
	notificationId := getEnvValue(EnvNotificationId)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
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
	initNotification()
	accessToken := getEnvValue(EnvAccessToken)
	customerId := getEnvValue(EnvCustomerId)
	notificationId := getEnvValue(EnvNotificationId)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
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
	initCreditCardCharge(false, true)
	accessToken := getEnvValue(EnvAccessToken)
	customerId := getEnvValue(EnvCustomerId)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nNotification := NewNotification(EnvSandbox, accessToken)
	resp, err := nNotification.GetAllByCustomer(ctx, customerId)
	assertResponseSuccess(t, resp, err)
}
