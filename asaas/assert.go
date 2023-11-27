package asaas

import (
	"encoding/json"
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
		logErrorSkipCaller(4, "unexpect: object is nil")
		t.Fail()
		return
	}
	vJson, _ := json.Marshal(v)
	logDebugSkipCaller(4, "success: object nonnull", string(vJson))
}

func AssertAsaasResponseSuccess(t *testing.T, resp any, err any) {
	r := reflect.ValueOf(resp)
	e := reflect.ValueOf(err)
	iResp, ok := resp.(response)
	if !e.IsNil() {
		vJson, _ := json.Marshal(err)
		logErrorSkipCaller(4, "unexpect: err asaas is not nil:", string(vJson))
		t.Fail()
	} else if r.IsNil() || !ok {
		logErrorSkipCaller(4, "unexpect: resp is nil or not response interface implemented")
		t.Fail()
	} else {
		vJson, _ := json.Marshal(resp)
		if iResp.IsSuccess() {
			logDebugSkipCaller(4, "success: resp is success:", string(vJson))
		} else {
			logErrorSkipCaller(4, "unexpect: resp is failure:", string(vJson))
			t.Fail()
		}
	}
}

func AssertAsaasResponseFailure(t *testing.T, resp any, err any) {
	r := reflect.ValueOf(resp)
	e := reflect.ValueOf(err)
	iResp, ok := resp.(response)
	if !e.IsNil() {
		vJson, _ := json.Marshal(err)
		logErrorSkipCaller(4, "unexpect: err asaas is not nil:", string(vJson))
		t.Fail()
	} else if r.IsNil() || !ok {
		logErrorSkipCaller(4, "unexpect: resp is nil or not response interface implemented")
		t.Fail()
	} else {
		vJson, _ := json.Marshal(resp)
		if iResp.IsSuccess() {
			logErrorSkipCaller(4, "unexpect: resp is success: ", string(vJson))
			t.Fail()
		} else {
			logDebugSkipCaller(4, "success: resp is failure:", string(vJson))
		}
	}
}

func AssertAsaasResponseNoContent(t *testing.T, resp any, err any) {
	r := reflect.ValueOf(resp)
	e := reflect.ValueOf(err)
	iResp, ok := resp.(response)
	if !e.IsNil() {
		vJson, _ := json.Marshal(err)
		logErrorSkipCaller(4, "unexpect: err asaas is not nil:", string(vJson))
		t.Fail()
		return
	} else if r.IsNil() || !ok {
		logErrorSkipCaller(4, "unexpect: resp is nil or not response interface implemented")
		t.Fail()
		return
	}
	vJson, _ := json.Marshal(resp)
	if iResp.IsNoContent() {
		logDebugSkipCaller(4, "success: resp is no content", string(vJson))
	} else {
		logErrorSkipCaller(4, "unexpect: resp has content ", string(vJson))
		t.Fail()
	}
}
