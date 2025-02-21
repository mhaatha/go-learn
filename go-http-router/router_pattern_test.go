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

func TestRouterPatternNamedParameter(t *testing.T) {
	router := httprouter.New()
	router.GET("/products/:productId/items/:itemId", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		productId := p.ByName("productId")
		itemId := p.ByName("itemId")
		responseText := "Product " + productId + " Item " + itemId
		fmt.Fprint(w, responseText)
	})

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/products/1/items/1", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()

	bodyData, _ := io.ReadAll(response.Body)
	assert.Equal(t, "Product 1 Item 1", string(bodyData))
}

func TestRouterPatternCatchAllParameter(t *testing.T) {
	router := httprouter.New()
	router.GET("/images/*image", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		image := p.ByName("image")
		response := "Image: " + image
		fmt.Fprint(w, response)
	})

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/images/small/profile.png", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()

	bodyData, _ := io.ReadAll(response.Body)
	assert.Equal(t, "Image: /small/profile.png", string(bodyData))
}
