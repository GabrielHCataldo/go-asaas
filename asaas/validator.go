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
	err = v.RegisterValidation("birth-date", validateBirthDate)
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

func validateBirthDate(fl validator.FieldLevel) bool {
	var timeValidate time.Time
	if fl.Field().Kind() == reflect.String {
		t, err := time.Parse(time.RFC3339, fl.Field().String())
		if err != nil {
			return false
		}
		timeValidate = t
	} else {
		datetime, ok := fl.Field().Interface().(time.Time)
		if !ok {
			return false
		}
		timeValidate = datetime
	}
	return util.ValidateBirthDate(timeValidate)
}

func validateDocument(fl validator.FieldLevel) bool {
	return util.IsCPF(fl.Field().String())
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
	remoteIP string,
) error {
	switch billingType {
	case CREDIT_CARD:
		if util.IsBlank(&cCardToken) && (cCard == nil || cCardHolderInfoBody == nil) {
			return berrors.New("charge by credit card, enter the credit card or credit card token")
		} else if cCard != nil && !util.ValidateExpirationCreditCard(cCard.ExpiryYear, cCard.ExpiryMonth) {
			return berrors.New("expired card")
		} else if util.IsBlank(&remoteIP) && !util.ValidateIP(remoteIP) {
			return berrors.New("invalid remoteIp")
		}
		break
	}
	return nil
}
