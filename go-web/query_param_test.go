package go_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func SayHello(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	if name == "" {
		fmt.Fprint(w, "Hello User")
	} else {
		fmt.Fprintf(w, "Hello %s", name)
	}
}

func MultipleParameter(w http.ResponseWriter, r *http.Request) {
	firstName := r.URL.Query().Get("first_name")
	lastName := r.URL.Query().Get("last_name")

	fmt.Fprintf(w, "Hello %s %s", firstName, lastName)
}

func MultipleValueParameter(w http.ResponseWriter, r *http.Request) {
	var query url.Values = r.URL.Query()
	var names []string = query["name"]

	fmt.Fprint(w, strings.Join(names, " "))
}

func TestQueryParameter(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "localhost:8080/hello?name=Hafidz", nil)
	recorder := httptest.NewRecorder()

	SayHello(recorder, request)

	response := recorder.Result()
	defer response.Body.Close()

	dataBody, _ := io.ReadAll(recorder.Body)
	fmt.Println(string(dataBody))
}

func TestMultipleQueryParameter(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "localhost:8080/hello?first_name=Hafidz&last_name=Athaya", nil)
	recorder := httptest.NewRecorder()

	MultipleParameter(recorder, request)

	response := recorder.Result()
	defer response.Body.Close()

	dataBody, _ := io.ReadAll(recorder.Body)
	fmt.Println(string(dataBody))
}

func TestMultipleValueParameter(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "localhost:8080/hello?name=Muhammad&name=Hafidz&name=Athaya", nil)
	recorder := httptest.NewRecorder()

	MultipleValueParameter(recorder, request)

	response := recorder.Result()
	defer response.Body.Close()

	dataBody, _ := io.ReadAll(recorder.Body)
	fmt.Println(string(dataBody))
}
