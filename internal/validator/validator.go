package validator

import (
	"github.com/GabrielHCataldo/go-asaas/internal/enum"
	"github.com/GabrielHCataldo/go-asaas/internal/util"
	"github.com/go-playground/validator/v10"
	"log"
	"reflect"
	"time"
)

var customValidate *validator.Validate

func Validate() *validator.Validate {
	if customValidate != nil {
		return customValidate
	}
	customValidate = validator.New()
	err := customValidate.RegisterValidation("enum", validateEnum)
	if err != nil {
		log.Fatal(err)
	}
	err = customValidate.RegisterValidation("phone", validatePhone)
	if err != nil {
		log.Fatal(err)
	}
	err = customValidate.RegisterValidation("full_name", validateFullName)
	if err != nil {
		log.Fatal(err)
	}
	err = customValidate.RegisterValidation("password", validatePassword)
	if err != nil {
		log.Fatal(err)
	}
	err = customValidate.RegisterValidation("birth-date", validateBirthDate)
	if err != nil {
		log.Fatal(err)
	}
	err = customValidate.RegisterValidation("document", validateDocument)
	if err != nil {
		log.Fatal(err)
	}
	err = customValidate.RegisterValidation("postal_code", validatePostalCode)
	if err != nil {
		log.Fatal(err)
	}
	return customValidate
}

func validatePhone(fl validator.FieldLevel) bool {
	return util.IsPhoneNumber(fl.Field().String())
}

func validateFullName(fl validator.FieldLevel) bool {
	return util.ValidateFullName(fl.Field().String())
}

func validatePassword(fl validator.FieldLevel) bool {
	return util.ValidatePassword(fl.Field().String())
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
	value := fl.Field().Interface().(enum.BaseEnum)
	return value.IsEnumValid()
}
