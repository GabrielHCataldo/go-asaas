package asaas

import (
	berrors "errors"
	"github.com/GabrielHCataldo/go-asaas/internal/util"
	"github.com/go-playground/validator/v10"
	"log"
	"reflect"
	"time"
)

var v *validator.Validate

func Validate() *validator.Validate {
	if v != nil {
		return v
	}
	v = validator.New()
	err := v.RegisterValidation("enum", validateEnum)
	if err != nil {
		log.Fatal(err)
	}
	err = v.RegisterValidation("phone", validatePhone)
	if err != nil {
		log.Fatal(err)
	}
	err = v.RegisterValidation("full_name", validateFullName)
	if err != nil {
		log.Fatal(err)
	}
	err = v.RegisterValidation("before_now", validateBeforeNow)
	if err != nil {
		log.Fatal(err)
	}
	err = v.RegisterValidation("after_now", validateAfterNow)
	if err != nil {
		log.Fatal(err)
	}
	err = v.RegisterValidation("document", validateDocument)
	if err != nil {
		log.Fatal(err)
	}
	err = v.RegisterValidation("postal_code", validatePostalCode)
	if err != nil {
		log.Fatal(err)
	}
	return v
}

func validatePhone(fl validator.FieldLevel) bool {
	return util.IsPhoneNumber(fl.Field().String())
}

func validateFullName(fl validator.FieldLevel) bool {
	return util.ValidateFullName(fl.Field().String())
}

func validateBeforeNow(fl validator.FieldLevel) bool {
	var timeValidate time.Time
	if fl.Field().Kind() == reflect.String {
		t, err := time.Parse(time.RFC3339, fl.Field().String())
		if err != nil {
			return false
		}
		timeValidate = t
	} else if fl.Field().Type().String() == "asaas.Date" {
		date, ok := fl.Field().Interface().(Date)
		if !ok {
			return false
		}
		timeValidate = time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
	} else {
		datetime, ok := fl.Field().Interface().(time.Time)
		if !ok {
			return false
		}
		timeValidate = datetime
	}
	return timeValidate.UTC().Before(time.Now().UTC())
}

func validateAfterNow(fl validator.FieldLevel) bool {
	var timeValidate time.Time
	if fl.Field().Kind() == reflect.String {
		t, err := time.Parse(time.RFC3339, fl.Field().String())
		if err != nil {
			return false
		}
		timeValidate = t
	} else if fl.Field().Type().String() == "asaas.Date" {
		date, ok := fl.Field().Interface().(Date)
		if !ok {
			return false
		}
		timeValidate = time.Date(date.Year(), date.Month(), date.Day(), 23, 59, 0, 0, date.Location())
	} else {
		datetime, ok := fl.Field().Interface().(time.Time)
		if !ok {
			return false
		}
		timeValidate = datetime
	}
	return timeValidate.UTC().After(time.Now().UTC())
}

func validateDocument(fl validator.FieldLevel) bool {
	return util.IsCpfCnpj(fl.Field().String())
}

func validatePostalCode(fl validator.FieldLevel) bool {
	return util.ValidatePostalCode(fl.Field().String())
}

func validateEnum(fl validator.FieldLevel) bool {
	value := fl.Field().Interface().(BaseEnum)
	return value.IsEnumValid()
}

func validateBillingBody(
	billingType BillingType,
	cCard *CreditCardRequest,
	cCardHolderInfoBody *CreditCardHolderInfoRequest,
	cCardToken,
	remoteIp string,
) error {
	switch billingType {
	case BillingTypeCreditCard:
		if util.IsBlank(&cCardToken) && (cCard == nil || cCardHolderInfoBody == nil) {
			return berrors.New("charge by credit card, enter the credit card or credit card token")
		} else if cCard != nil && !util.ValidateExpirationCreditCard(cCard.ExpiryYear, cCard.ExpiryMonth) {
			return berrors.New("expired card")
		} else if util.IsBlank(&remoteIp) && !util.ValidateIp(remoteIp) {
			return berrors.New("invalid remoteIp")
		}
		break
	}
	return nil
}
