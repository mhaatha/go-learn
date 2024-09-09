package go_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func ResponseCode(writer http.ResponseWriter, request *http.Request) {
	name := request.URL.Query().Get("name")

	if name == "" {
		writer.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(writer, "name cannot empty")
	} else {
		writer.WriteHeader(http.StatusOK)
		fmt.Fprintf(writer, "Hello %s", name)
	}
}

func TestResponseCode(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/rescode?name=Gustavo", nil)
	recorder := httptest.NewRecorder()

	ResponseCode(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)

	expected := "Hello Gustavo"
	if bodyString != expected {
		t.Errorf("Expected '%s' but got '%s'", expected, bodyString)
	}
	expected = "200 OK"
	if response.Status != expected {
		t.Errorf("Expected status '%s' but got '%s'", expected, response.Status)
	}
}
