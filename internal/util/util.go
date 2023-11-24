package util

import (
	"github.com/klassmann/cpfcnpj"
	"github.com/nyaruka/phonenumbers"
	"regexp"
	"strings"
	"time"
)

func IsPhoneNumber(value string) bool {
	num, err := phonenumbers.Parse(value, "BR")
	if err == nil {
		return phonenumbers.IsValidNumberForRegion(num, "BR")
	}
	return false
}

func IsCPF(value string) bool {
	return cpfcnpj.ValidateCPF(cpfcnpj.Clean(value))
}

func ValidateFullName(value string) bool {
	regex := regexp.MustCompile(`^([a-zA-Z]{2,}\s[a-zA-Z]+'?-?[a-zA-Z]+\s?([a-zA-Z]+)?)`)
	return regex.MatchString(value)
}

func ValidatePassword(value string) bool {
	return len(value) >= 6
}

func ValidateBirthDate(values time.Time) bool {
	return time.Now().After(values)
}

func IsBlank(value *string) bool {
	return value == nil || len(strings.TrimSpace(*value)) == 0
}

func IsNotBlank(value *string) bool {
	return !IsBlank(value)
}
