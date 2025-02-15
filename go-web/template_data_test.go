package go_web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

type Page struct {
	Title   string
	Name    string
	Address Address
}

type Address struct {
	Street string
}

func TemplateDataMap(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/name.gohtml"))
	t.ExecuteTemplate(w, "name.gohtml", map[string]interface{}{
		"Title": "Template Data Map",
		"Name":  "Hafidz",
		"Address": map[string]interface{}{
			"Street": "Jalan in Aja Dulu V2",
		},
	})
}

func TemplateDataStruct(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/name.gohtml"))
	t.ExecuteTemplate(w, "name.gohtml", Page{
		Title: "Template Data Struct",
		Name:  "Hafidz",
		Address: Address{
			Street: "Jalan in Aja Dulu",
		},
	})
}

func TestTemplateDataMap(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateDataMap(recorder, request)

	response := recorder.Result()
	bodyData, _ := io.ReadAll(response.Body)

	fmt.Println(string(bodyData))
}

func TestTemplateDataStruct(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateDataStruct(recorder, request)

	response := recorder.Result()
	bodyData, _ := io.ReadAll(response.Body)

	fmt.Println(string(bodyData))
}
