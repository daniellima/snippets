package assert

import (
	"net/http/httptest"
	"testing"
)

func Status(t *testing.T, expectedStatus int, recorder *httptest.ResponseRecorder) {
	actualStatus := recorder.Result().StatusCode

	if expectedStatus != actualStatus {
		t.Errorf("Status %v is different from expected status %v", actualStatus, expectedStatus)
	}
}

func Body(t *testing.T, expectedBody string, recorder *httptest.ResponseRecorder) {
	actualBody := recorder.Body.String()

	if expectedBody != actualBody {
		t.Errorf("Body \"%v\" is different from expected body \"%v\"", actualBody, expectedBody)
	}
}
