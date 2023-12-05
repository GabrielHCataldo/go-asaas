package util

import (
	"encoding/json"
	"os"
	"path"
	"reflect"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"time"
)

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

func GetValueByReflect(f reflect.Value) any {
	if f.IsZero() || !f.IsValid() {
		return nil
	}
	v := f.Interface()
	if x, ok := v.(*os.File); ok {
		v = x
	} else if f.Kind() == reflect.Pointer {
		v = f.Elem().Interface()
	}
	return v
}

func GetJsonFieldNameByReflect(f reflect.StructField) string {
	sk := strings.Split(f.Tag.Get("json"), ",")
	if len(sk) > 0 && !IsBlank(&sk[0]) && sk[0] != "-" {
		return sk[0]
	}
	return ""
}

func GenerateEmail() string {
	randomNumber := strconv.Itoa(int(time.Now().UnixNano()))
	return "unit" + randomNumber + "@gmail.com"
}

func GenerateMobilePhone() string {
	randomNumber := strconv.Itoa(int(time.Now().UnixNano()))
	return "4799" + randomNumber[len(randomNumber)-7:]
}

func IsJson(v []byte) bool {
	var data map[string]any
	return json.Unmarshal(v, &data) == nil
}
