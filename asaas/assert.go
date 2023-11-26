package asaas

import (
	"encoding/json"
	"github.com/fatih/structs"
	"reflect"
	"testing"
)

func AssertFatalErrorNonnull(t *testing.T, err error) {
	if err != nil {
		logErrorSkipCaller(4, err)
		t.Fatal()
	}
}

func AssertSuccessNonnull(t *testing.T, v any) {
	r := reflect.ValueOf(v)
	if !r.IsNil() {
		logErrorSkipCaller(4, "unexpect result: object is nil")
		t.Fail()
		return
	}
	vJson, _ := json.Marshal(v)
	logDebugSkipCaller(4, "success result: object nonnull", string(vJson))
}

func AssertAsaasResponseSuccess(t *testing.T, resp any, err any) {
	r := reflect.ValueOf(resp)
	e := reflect.ValueOf(err)
	if !e.IsNil() {
		vJson, _ := json.Marshal(err)
		logErrorSkipCaller(4, "unexpect result: err Asaas is not nil:", string(vJson))
		t.Fail()
	} else if r.IsNil() {
		logErrorSkipCaller(4, "unexpect result: result is nil")
		t.Fail()
	} else {
		vJson, _ := json.Marshal(resp)
		respStruct := structs.New(resp)
		errorsField := respStruct.Field("Errors")
		if errorsField != nil {
			logErrorSkipCaller(4, "unexpect result: errors result from Asaas:", string(vJson))
			t.Fail()
		} else {
			logDebugSkipCaller(4, "success result:", string(vJson))
		}
	}
}

func AssertAsaasResponseFailure(t *testing.T, resp any, err any) {
	r := reflect.ValueOf(resp)
	e := reflect.ValueOf(err)
	if !e.IsNil() {
		vJson, _ := json.Marshal(err)
		logErrorSkipCaller(4, "unexpect result: err Asaas is not nil:", string(vJson))
		t.Fail()
	} else if r.IsNil() {
		logErrorSkipCaller(4, "unexpect result: result is nil")
		t.Fail()
	} else {
		vJson, _ := json.Marshal(resp)
		respStruct := structs.New(resp)
		errorsField := respStruct.Field("Errors")
		if errorsField == nil {
			logDebugSkipCaller(4, "unexpect result: errors result from Asaas is empty")
			t.Fail()
		} else {
			logDebugSkipCaller(4, "success result:", string(vJson))
		}
	}
}
