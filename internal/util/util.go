package util

import (
	"encoding/json"
	"fmt"
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

func GetValueByReflectField(f reflect.Value) any {
	if f.IsZero() || !f.IsValid() {
		return nil
	}
	k := f.Kind()
	if x, ok := f.Interface().(*os.File); ok {
		return x
	} else if f.Kind() == reflect.Pointer {
		f = f.Elem()
		k = f.Kind()
	}
	switch k {
	case reflect.String:
		return f.String()
	case reflect.Int:
		return f.Int()
	case reflect.Bool:
		return f.Bool()
	case reflect.Float32, reflect.Float64:
		return f.Float()
	default:
		return f.Interface()
	}
}

func GetJsonFieldNameByReflectField(f reflect.StructField) string {
	sk := strings.Split(f.Tag.Get("json"), ",")
	return sk[0]
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
	if len(v) > 0 {
		var data map[string]any
		return json.Unmarshal(v, &data) == nil
	}
	return false
}

func ConvertToString(v any) string {
	if v == nil {
		return ""
	} else if i, iOk := v.(int); iOk {
		return strconv.Itoa(i)
	} else if i32, i32Ok := v.(int32); i32Ok {
		return strconv.Itoa(int(i32))
	} else if i64, i64Ok := v.(int64); i64Ok {
		return strconv.Itoa(int(i64))
	} else if b, bOk := v.(bool); bOk {
		return strconv.FormatBool(b)
	} else if f32, f32Ok := v.(float32); f32Ok {
		return strconv.FormatFloat(float64(f32), 'f', -1, 32)
	} else if f64, f64Ok := v.(float64); f64Ok {
		return strconv.FormatFloat(f64, 'f', -1, 64)
	} else if s, sOk := v.(string); sOk {
		return s
	}
	return fmt.Sprintf(`"%s"`, v)
}
