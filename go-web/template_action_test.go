package go_web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TemplateActionIf(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/if.gohtml"))
	t.ExecuteTemplate(w, "if.gohtml", Page{
		Name: "Hafidz",
	})
}

func TemplateActionOperator(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/comparator.gohtml"))
	t.ExecuteTemplate(w, "comparator.gohtml", map[string]interface{}{
		"Title":      "Template Action Operator",
		"FinalValue": 90,
	})
}

func TemplateActionRange(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/range.gohtml"))
	t.ExecuteTemplate(w, "range.gohtml", map[string]interface{}{
		"Title": "Template Action Range",
		"Hobbies": []string{
			"Gaming", "Read", "Code",
		},
	})
}

func TemplateActionWith(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/with.gohtml"))
	t.ExecuteTemplate(w, "with.gohtml", map[string]interface{}{
		"Title": "Template Action With",
		"Name":  "Hafidz",
	})
}

func TestTemplateActionIf(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateActionIf(recorder, request)

	response := recorder.Result()
	bodyData, _ := io.ReadAll(response.Body)

	fmt.Println(string(bodyData))
}

func TestTemplateActionOperator(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateActionOperator(recorder, request)

	response := recorder.Result()
	bodyData, _ := io.ReadAll(response.Body)

	fmt.Println(string(bodyData))
}

func TestTemplateActionRange(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateActionRange(recorder, request)

	response := recorder.Result()
	bodyData, _ := io.ReadAll(response.Body)

	fmt.Println(string(bodyData))
}

func TestTemplateActionWith(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateActionWith(recorder, request)

	response := recorder.Result()
	bodyData, _ := io.ReadAll(response.Body)

	fmt.Println(string(bodyData))
}
