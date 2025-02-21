package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestNotFound(t *testing.T) {
	router := httprouter.New()
	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Not Found")
	})

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()

	bodyData, _ := io.ReadAll(response.Body)
	assert.Equal(t, "Not Found", string(bodyData))
}

func TestNotFoundServer(t *testing.T) {
	router := httprouter.New()
	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		myTemplates.ExecuteTemplate(w, "not_found.gohtml", struct {
			Title   string
			Error   string
			Message string
			Button  string
		}{
			Title:   "Page Not Found",
			Error:   "Oops! Page Not Found",
			Message: "The page you are looking for may have moved or may not exist.",
			Button:  "Back to Home 🏠",
		})
	})
	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprint(w, "Hello httprouter")
	})

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	server.ListenAndServe()
}
