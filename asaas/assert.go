package asaas

import (
	"encoding/json"
	"testing"
)

func assertResponseSuccess(t *testing.T, resp any, err error) {
	iResp, ok := resp.(response)
	vJson, _ := json.Marshal(resp)
	if err == nil && ok && iResp != nil && iResp.IsSuccess() && !iResp.IsNoContent() && !iResp.IsFailure() {
		logDebugSkipCaller(4, "success: resp is success:", string(vJson))
		return
	}
	logDebugSkipCaller(4, "failure: resp is not success:", string(vJson), "err:", err)
	t.Fail()
}

func assertResponseFailure(resp any, err error) {
	iResp, ok := resp.(response)
	vJson, _ := json.Marshal(resp)
	if err == nil && ok && iResp != nil && iResp.IsFailure() && !iResp.IsSuccess() && !iResp.IsNoContent() {
		logDebugSkipCaller(4, "success: resp is failure:", string(vJson))
	}
}

func assertResponseNoContent(resp any, err error) {
	iResp, ok := resp.(response)
	vJson, _ := json.Marshal(resp)
	if err == nil && ok && iResp != nil && iResp.IsNoContent() && !iResp.IsSuccess() && !iResp.IsFailure() {
		logDebugSkipCaller(4, "success: resp is no content", string(vJson))
	}
}
