package go_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Request Header
func RequestHeader(writer http.ResponseWriter, request *http.Request) {
	contentType := request.Header.Get("content-type")
	fmt.Fprint(writer, contentType)
}

func TestRequestHeader(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	RequestHeader(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)

	expectedResponse := "application/json"
	if expectedResponse != bodyString {
		t.Errorf("Expected '%s' but got '%s'", expectedResponse, bodyString)
	}
}

// Response Header
func ResponseHeader(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("X-Powered-Bye", "mhaatha")
	fmt.Fprint(writer, "OK")
}

func TestResponseHeader(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	ResponseHeader(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)

	expectedResponse := "OK"
	if expectedResponse != bodyString {
		t.Errorf("Expected '%s' but got '%s'", expectedResponse, bodyString)
	}

	headerResponse := response.Header.Get("X-Powered-Bye")
	if headerResponse != "mhaatha" {
		t.Errorf("Expected 'mhaatha' but got '%s'", headerResponse)
	}
}
