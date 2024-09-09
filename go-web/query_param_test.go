package go_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Query
func Welcome(writer http.ResponseWriter, request *http.Request) {
	name := request.URL.Query().Get("name")
	if name == "" {
		fmt.Fprint(writer, "Welcome Guest!")
	} else {
		fmt.Fprintf(writer, "Welcome %s!", name)
	}
}

func TestQueryParameter(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/welcome?name=", nil)
	recorder := httptest.NewRecorder()

	Welcome(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)

	expectedResponse := "Welcome Athaya!"
	if expectedResponse != bodyString {
		t.Errorf("Expected '%s' but got '%s'", expectedResponse, bodyString)
	}
}
