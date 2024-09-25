package go_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func HelloHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "Hello World")
}

// HTTP Test
func TestHttp(t *testing.T) {
	// Test for request using NewRequest(method, target, body)
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello", nil)
	// Test for response
	recorder := httptest.NewRecorder()

	// Call the handler
	HelloHandler(recorder, request)

	// Get the response body
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)

	// Check the response body
	expectedResponse := "Hello World"
	if expectedResponse != bodyString {
		t.Errorf("Expected '%s' but got '%s'", expectedResponse, bodyString)
	}
}
