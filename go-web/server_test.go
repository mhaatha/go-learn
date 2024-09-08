package go_web

import (
	"fmt"
	"net/http"
	"testing"
)

// Server
func TestServer(t *testing.T) {
	var server http.Server = http.Server{
		Addr: "localhost:8080",
	}

	var err error = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

// Handler
func TestHandler(t *testing.T) {
	var myHandler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
		_, err := fmt.Fprint(writer, "Hello Golang Web Server")
		if err != nil {
			panic(err)
		}
	}

	var server http.Server = http.Server{
		Addr:    "localhost:8080",
		Handler: myHandler,
	}

	var err error = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

// ServeMux
func TestServeMux(t *testing.T) {
	var mux *http.ServeMux = http.NewServeMux()
	mux.HandleFunc("/login", func(writer http.ResponseWriter, request *http.Request) {
		_, err := fmt.Fprint(writer, "Masukkan Username:")
		if err != nil {
			panic(err)
		}
	})
	mux.HandleFunc("/logout", func(writer http.ResponseWriter, request *http.Request) {
		_, err := fmt.Fprint(writer, "Anda berhasil logout")
		if err != nil {
			panic(err)
		}
	})
	mux.HandleFunc("/sama/", func(writer http.ResponseWriter, request *http.Request) {
		_, err := fmt.Fprint(writer, "Sama")
		if err != nil {
			panic(err)
		}
	})
	mux.HandleFunc("/sama/dengan", func(writer http.ResponseWriter, request *http.Request) {
		_, err := fmt.Fprint(writer, "=")
		if err != nil {
			panic(err)
		}
	})

	var server http.Server = http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	var err error = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

// Request
func TestRequest(t *testing.T) {
	var myHandler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
		_, err := fmt.Fprint(writer, request.Method)
		if err != nil {
			panic(err)
		}
		_, err = fmt.Fprint(writer, request.RequestURI)
		if err != nil {
			panic(err)
		}
	}

	var server http.Server = http.Server{
		Addr:    "localhost:8080",
		Handler: myHandler,
	}

	var err error = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
