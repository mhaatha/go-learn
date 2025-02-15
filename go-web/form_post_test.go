package go_web

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func FormPost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "request body is not a valid form", http.StatusBadRequest)
		return
	}

	firstName := r.PostForm.Get("first_name")
	lastName := r.PostForm.Get("last_name")

	fmt.Fprintf(w, "Hello %s %s", firstName, lastName)
}

func FormJSON(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "unable to parsing json from request body", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "JSON Data: %v", user)
}

func TestFormPost(t *testing.T) {
	requestBody := strings.NewReader("first_name=Hafidz&last_name=Athaya")
	request := httptest.NewRequest(http.MethodPost, "localhost:8080/hello", requestBody)
	request.Header.Set("content-type", "application/x-www-form-urlencoded")
	recorder := httptest.NewRecorder()

	FormPost(recorder, request)

	response := recorder.Result()
	bodyData, _ := io.ReadAll(response.Body)

	fmt.Println(string(bodyData))
}

func TestFormJSON(t *testing.T) {
	user := User{
		FirstName: "Hafidz",
		LastName:  "Athaya",
	}
	b, err := json.Marshal(user)
	if err != nil {
		t.Errorf("can't marshal a struct into JSON: %v", err)
	}

	requestBody := strings.NewReader(string(b))
	request := httptest.NewRequest(http.MethodPost, "localhost:8080/helloJSON", requestBody)
	request.Header.Set("content-type", "application/json")
	recorder := httptest.NewRecorder()

	FormJSON(recorder, request)

	response := recorder.Result()
	bodyData, _ := io.ReadAll(response.Body)

	fmt.Println(string(bodyData))
}
