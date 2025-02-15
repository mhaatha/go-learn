package go_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func RequestHeader(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("content-type")
	fmt.Fprint(w, contentType)
}

func ResponseHeader(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("X-Powered-By", "Team 5")
	fmt.Fprint(w, "OK")
}

func TestRequestHeader(t *testing.T) {
	request := httptest.NewRequest(http.MethodPost, "localhost:8080", nil)
	request.Header.Add("content-type", "application/json")
	recorder := httptest.NewRecorder()

	RequestHeader(recorder, request)

	response := recorder.Result()
	bodyData, _ := io.ReadAll(response.Body)

	fmt.Println(string(bodyData))
}

func TestResponseHeader(t *testing.T) {
	request := httptest.NewRequest(http.MethodPost, "localhost:8080", nil)
	recorder := httptest.NewRecorder()

	ResponseHeader(recorder, request)

	response := recorder.Result()
	bodyData, _ := io.ReadAll(response.Body)
	poweredBy := response.Header.Get("X-Powered-By")

	fmt.Println(string(bodyData))
	fmt.Println(poweredBy)
}
