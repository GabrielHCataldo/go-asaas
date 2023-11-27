package asaas

import (
	"context"
	"fmt"
	"net/http"
)

type UpdateNotificationRequest struct {
	Enabled                     bool   `json:"enabled,omitempty"`
	EmailEnabledForProvider     bool   `json:"emailEnabledForProvider,omitempty"`
	SmsEnabledForProvider       bool   `json:"smsEnabledForProvider,omitempty"`
	EmailEnabledForCustomer     bool   `json:"emailEnabledForCustomer,omitempty"`
	SmsEnabledForCustomer       bool   `json:"smsEnabledForCustomer,omitempty"`
	PhoneCallEnabledForCustomer bool   `json:"phoneCallEnabledForCustomer,omitempty"`
	WhatsappEnabledForCustomer  bool   `json:"whatsappEnabledForCustomer,omitempty"`
	Event                       string `json:"event,omitempty"`
	ScheduleOffset              int    `json:"scheduleOffset,omitempty"`
}

type UpdateManyNotificationsRequest struct {
	Customer      string                          `json:"customer,omitempty" validate:"required"`
	Notifications []UpdateManyNotificationRequest `json:"notifications,omitempty" validate:"required"`
}

type UpdateManyNotificationRequest struct {
	ID                          string `json:"id,omitempty" validate:"required"`
	Enabled                     bool   `json:"enabled,omitempty"`
	EmailEnabledForProvider     bool   `json:"emailEnabledForProvider,omitempty"`
	SmsEnabledForProvider       bool   `json:"smsEnabledForProvider,omitempty"`
	EmailEnabledForCustomer     bool   `json:"emailEnabledForCustomer,omitempty"`
	SmsEnabledForCustomer       bool   `json:"smsEnabledForCustomer,omitempty"`
	PhoneCallEnabledForCustomer bool   `json:"phoneCallEnabledForCustomer,omitempty"`
	WhatsappEnabledForCustomer  bool   `json:"whatsappEnabledForCustomer,omitempty"`
	Event                       string `json:"event,omitempty"`
	ScheduleOffset              int    `json:"scheduleOffset,omitempty"`
}

type NotificationResponse struct {
	ID                          string            `json:"id,omitempty"`
	Customer                    string            `json:"customer,omitempty"`
	Enabled                     bool              `json:"enabled,omitempty"`
	EmailEnabledForProvider     bool              `json:"emailEnabledForProvider,omitempty"`
	SmsEnabledForProvider       bool              `json:"smsEnabledForProvider,omitempty"`
	EmailEnabledForCustomer     bool              `json:"emailEnabledForCustomer,omitempty"`
	SmsEnabledForCustomer       bool              `json:"smsEnabledForCustomer,omitempty"`
	PhoneCallEnabledForCustomer bool              `json:"phoneCallEnabledForCustomer,omitempty"`
	WhatsappEnabledForCustomer  bool              `json:"whatsappEnabledForCustomer,omitempty"`
	Event                       NotificationEvent `json:"event,omitempty"`
	ScheduleOffset              int               `json:"scheduleOffset,omitempty"`
	Errors                      []ErrorResponse   `json:"errors,omitempty"`
	Deleted                     bool              `json:"deleted,omitempty"`
}

type UpdateManyNotificationsResponse struct {
	Notifications []NotificationResponse `json:"notifications,omitempty"`
	Errors        []ErrorResponse        `json:"errors,omitempty"`
}

type notification struct {
	env         Env
	accessToken string
}

type Notification interface {
	UpdateByID(ctx context.Context, notificationID string, body UpdateNotificationRequest) (*NotificationResponse, Error)
	UpdateManyByCustomer(ctx context.Context, body UpdateManyNotificationsRequest) (*UpdateManyNotificationsResponse,
		Error)
	GetAllByCustomer(ctx context.Context, customerID string) (*Pageable[NotificationResponse], Error)
}

func NewNotification(env Env, accessToken string) Notification {
	return notification{
		env:         env,
		accessToken: accessToken,
	}
}

func (n notification) UpdateByID(ctx context.Context, notificationID string, body UpdateNotificationRequest) (
	*NotificationResponse, Error) {
	if err := Validate().Struct(body); err != nil {
		return nil, NewError(ERROR_VALIDATION, err)
	}
	req := NewRequest[NotificationResponse](ctx, n.env, n.accessToken)
	return req.make(http.MethodPut, fmt.Sprintf("/v3/notifications/%s", notificationID), body)
}

func (n notification) UpdateManyByCustomer(ctx context.Context, body UpdateManyNotificationsRequest) (
	*UpdateManyNotificationsResponse, Error) {
	if err := Validate().Struct(body); err != nil {
		return nil, NewError(ERROR_VALIDATION, err)
	}
	req := NewRequest[UpdateManyNotificationsResponse](ctx, n.env, n.accessToken)
	return req.make(http.MethodPut, "/v3/notifications/batch", body)
}

func (n notification) GetAllByCustomer(ctx context.Context, customerID string) (*Pageable[NotificationResponse], Error) {
	req := NewRequest[Pageable[NotificationResponse]](ctx, n.env, n.accessToken)
	return req.make(http.MethodGet, fmt.Sprintf("/v3/customers/%s/notifications", customerID), nil)
}
