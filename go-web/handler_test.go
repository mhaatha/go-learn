package go_web

import (
	"fmt"
	"net/http"
	"testing"
)

func TestHandlerFunction(t *testing.T) {
	var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World")
	}

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: handler,
	}

	if err := server.ListenAndServe(); err != nil {
		t.Error(err)
	}
}

func TestServeMuxFunction(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World")
	})
	mux.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hi")
	})
	mux.HandleFunc("/images/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Images")
	})
	mux.HandleFunc("/images/thumbnails", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hi")
	})

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	if err := server.ListenAndServe(); err != nil {
		t.Error(err)
	}
}

func TestRequestHandler(t *testing.T) {
	var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, r.Method)
		fmt.Fprintln(w, r.RequestURI)
		fmt.Fprintln(w, r.URL.Path)
	}

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: handler,
	}

	if err := server.ListenAndServe(); err != nil {
		t.Error(err)
	}
}
