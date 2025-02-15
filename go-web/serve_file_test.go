package go_web

import (
	_ "embed"
	"fmt"
	"net/http"
	"testing"
)

func ServeFile(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	if name != "" {
		http.ServeFile(w, r, "./resources/ok.html")
	} else {
		http.ServeFile(w, r, "./resources/notfound.html")
	}
}

//go:embed resources/ok.html
var resourceOK string

//go:embed resources/notfound.html
var resourceNotFound string

func ServeFileEmbed(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	if name != "" {
		fmt.Fprint(w, resourceOK)
	} else {
		fmt.Fprint(w, resourceNotFound)
	}
}

func TestServeFileServer(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(ServeFile),
	}

	if err := server.ListenAndServe(); err != nil {
		t.Error(err)
	}
}

func TestServeFileEmbed(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(ServeFileEmbed),
	}

	if err := server.ListenAndServe(); err != nil {
		t.Error(err)
	}
}
