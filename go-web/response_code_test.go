package go_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func ResponseCode(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	if name == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "name is empty")
	} else {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Hello %s", name)
	}
}

func TestResponseCodeValid(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "localhost:8080/user?name=Hafidz", nil)
	recorder := httptest.NewRecorder()

	ResponseCode(recorder, request)

	response := recorder.Result()
	statusCode := response.StatusCode
	responseBody, _ := io.ReadAll(response.Body)

	fmt.Println("Status Code:", statusCode)
	fmt.Println("Response Body:", string(responseBody))
}

func TestResponseCodeInvalid(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "localhost:8080/user", nil)
	recorder := httptest.NewRecorder()

	ResponseCode(recorder, request)

	response := recorder.Result()
	statusCode := response.StatusCode
	responseBody, _ := io.ReadAll(response.Body)

	fmt.Println("Status Code:", statusCode)
	fmt.Println("Response Body:", string(responseBody))
}
