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

func TestParams(t *testing.T) {
	router := httprouter.New()
	router.GET("/products/:id", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		productId := p.ByName("id")
		responseText := "Product " + productId
		fmt.Fprint(w, responseText)
	})

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/products/2", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()

	bodyData, _ := io.ReadAll(response.Body)
	assert.Equal(t, "Product 2", string(bodyData))
}
