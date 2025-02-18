package go_web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TemplateAutoEscape(w http.ResponseWriter, r *http.Request) {
	myTemplates.ExecuteTemplate(w, "post.gohtml", map[string]interface{}{
		"Title": "Template Auto Escape",
		"Body":  "<p>This is body<script>alert('You got hacked')</script></p>",
	})
}

func TemplateAutoEscapeDisabled(w http.ResponseWriter, r *http.Request) {
	myTemplates.ExecuteTemplate(w, "post.gohtml", map[string]interface{}{
		"Title": "Template Auto Escape",
		"Body":  template.HTML("<p>This is body<script>alert('You got hacked')</script></p>"),
	})
}

func TemplateXSS(w http.ResponseWriter, r *http.Request) {
	myTemplates.ExecuteTemplate(w, "post.gohtml", map[string]interface{}{
		"Title": "Template Auto Escape",
		"Body":  template.HTML(r.URL.Query().Get("body")),
	})
}

func TestTemplateAutoEscape(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateAutoEscape(recorder, request)

	response := recorder.Result()
	bodyData, _ := io.ReadAll(response.Body)

	fmt.Println(string(bodyData))
}

func TestTemplateAutoEscapeServer(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/auto-escape", TemplateAutoEscape)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	if err := server.ListenAndServe(); err != nil {
		t.Error(err)
	}
}

func TestTemplateAutoEscapeDisabledServer(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/auto-escape", TemplateAutoEscapeDisabled)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	if err := server.ListenAndServe(); err != nil {
		t.Error(err)
	}
}

func TestTemplateXSSServer(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/xss", TemplateXSS)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	if err := server.ListenAndServe(); err != nil {
		t.Error(err)
	}
}
