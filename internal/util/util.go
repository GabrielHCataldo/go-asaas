package util

import (
	"github.com/klassmann/cpfcnpj"
	"github.com/nyaruka/phonenumbers"
	"path"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"time"
)

func IsPhoneNumber(value string) bool {
	num, err := phonenumbers.Parse(value, "BR")
	if err == nil {
		return phonenumbers.IsValidNumber(num)
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

func ValidateBirthDate(values time.Time) bool {
	return time.Now().After(values)
}

func ValidatePostalCode(v string) bool {
	regex := regexp.MustCompile(`^\d{5}-\d{3}?$`)
	return regex.MatchString(v)
}

func ValidateIP(v string) bool {
	regex := regexp.MustCompile(`\b((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\.|$)){4}\b`)
	return regex.MatchString(v)
}

func ValidateExpirationCreditCard(expiryYear, expiryMonth string) bool {
	exp, err := time.Parse("2006-01-02", expiryYear+"-"+expiryMonth+"-01")
	if err != nil || time.Now().UTC().After(exp.UTC()) {
		return false
	}
	return true
}

func IsBlank(value *string) bool {
	return value == nil || len(strings.TrimSpace(*value)) == 0
}

func IsNotBlank(value *string) bool {
	return !IsBlank(value)
}

func ReplaceAllSpacesRepeat(v string) string {
	re := regexp.MustCompile(`\s+`)
	out := re.ReplaceAllString(v, " ")
	return strings.TrimSpace(out)
}

func GetSystemInfo(skipCaller int) (fileName string, line string, funcName string) {
	pc := make([]uintptr, 10) // at least 1 entry needed
	runtime.Callers(skipCaller, pc)
	f := runtime.FuncForPC(pc[0])
	file, lineInt := f.FileLine(pc[0])
	fileBase := path.Base(file)

	nameFunc := path.Base(f.Name())
	splitFunc := strings.Split(nameFunc, ".")
	if len(splitFunc) >= 3 {
		nameFunc = strings.Replace(nameFunc, splitFunc[0]+".", "", 1)
	}
	return fileBase, strconv.Itoa(lineInt), nameFunc
}
