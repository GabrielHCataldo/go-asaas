package asaas

import (
	"encoding/json"
	"github.com/GabrielHCataldo/go-asaas/internal/util"
	"reflect"
	"testing"
)

func assertFatalErrorNonnull(t *testing.T, err error) {
	if err != nil {
		logErrorSkipCaller(4, err)
		t.Fatal()
	}
}

func assertFatalIsBlank(t *testing.T, v string) {
	if util.IsBlank(&v) {
		logErrorSkipCaller(4, "value is blank")
		t.Fatal()
	}
}

func assertSuccessNonnull(t *testing.T, v any) {
	if v != nil || !reflect.ValueOf(v).IsNil() {
		vJson, _ := json.Marshal(v)
		logDebugSkipCaller(4, "success: object nonnull", string(vJson))
	}
	t.Fail()
}

func assertResponseSuccess(t *testing.T, resp any, err any) {
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
		if iResp.IsSuccess() && !iResp.IsNoContent() && !iResp.IsFailure() {
			logDebugSkipCaller(4, "success: resp is success:", string(vJson))
		} else {
			logErrorSkipCaller(4, "unexpect: resp is failure:", string(vJson))
			t.Fail()
		}
	}
}

func assertResponseFailure(t *testing.T, resp any, err any) {
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
		if iResp.IsSuccess() || !iResp.IsFailure() || iResp.IsNoContent() {
			logErrorSkipCaller(4, "unexpect: resp is success: ", string(vJson))
			t.Fail()
		} else {
			logDebugSkipCaller(4, "success: resp is failure:", string(vJson))
		}
	}
}

func assertResponseNoContent(t *testing.T, resp any, err any) {
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
	if iResp.IsNoContent() && !iResp.IsSuccess() && !iResp.IsFailure() {
		logDebugSkipCaller(4, "success: resp is no content", string(vJson))
	} else {
		logErrorSkipCaller(4, "unexpect: resp has content ", string(vJson))
		t.Fail()
	}
}
