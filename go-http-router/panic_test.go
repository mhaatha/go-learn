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

func TestPanicHandler(t *testing.T) {
	router := httprouter.New()
	router.PanicHandler = func(w http.ResponseWriter, r *http.Request, i interface{}) {
		fmt.Fprint(w, "Panic: ", i)
	}
	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		panic("Upss")
	})

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()

	bodyData, _ := io.ReadAll(response.Body)
	assert.Equal(t, "Panic: Upss", string(bodyData))
}

func TestPanicHandlerServer(t *testing.T) {
	router := httprouter.New()
	router.PanicHandler = func(w http.ResponseWriter, r *http.Request, i interface{}) {
		errName, _ := i.(string)

		myTemplates.ExecuteTemplate(w, "server_busy.gohtml", struct {
			Title   string
			Error   string
			Message string
			Button  string
		}{
			Title:   errName,
			Error:   fmt.Sprintf("🚧 %s 🚧", errName),
			Message: "Please wait a moment and try again later.",
			Button:  "Try again 🔄",
		})
	}
	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		panic("Server is busy")
	})

	server := http.Server{
		Addr:    ":3000",
		Handler: router,
	}

	server.ListenAndServe()
}
