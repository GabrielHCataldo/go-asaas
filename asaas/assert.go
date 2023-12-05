package asaas

import (
	"encoding/json"
	"testing"
)

func assertResponseSuccess(t *testing.T, resp any, err error) {
	if err != nil {
		logDebugSkipCaller(4, "failure: err is not nil:", err)
		t.Fail()
		return
	}
	iResp, ok := resp.(response)
	vJson, _ := json.Marshal(resp)
	if ok && iResp != nil && iResp.IsSuccess() && !iResp.IsNoContent() && !iResp.IsFailure() {
		logDebugSkipCaller(4, "success: resp is success:", string(vJson))
		return
	}
	logDebugSkipCaller(4, "failure: resp is not success:", string(vJson))
	t.Fail()
}

func assertResponseFailure(t *testing.T, resp any, err error) {
	if err != nil {
		logDebugSkipCaller(4, "failure: err is not nil:", err)
		return
	}
	iResp, ok := resp.(response)
	vJson, _ := json.Marshal(resp)
	if ok && iResp != nil && iResp.IsFailure() {
		logDebugSkipCaller(4, "success: resp is failure:", string(vJson))
		return
	}
	logDebugSkipCaller(4, "failure: resp is not failure:", string(vJson))
	t.Fail()
}

func assertResponseNoContent(t *testing.T, resp any, err error) {
	if err != nil {
		logDebugSkipCaller(4, "failure: err is not nil:", err)
		return
	}
	iResp, ok := resp.(response)
	vJson, _ := json.Marshal(resp)
	if ok && iResp != nil && iResp.IsNoContent() && !iResp.IsSuccess() && !iResp.IsFailure() {
		logDebugSkipCaller(4, "success: resp is no content", string(vJson))
		return
	}
	logDebugSkipCaller(4, "failure: resp is not no content:", string(vJson))
	t.Fail()
}
