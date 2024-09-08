package go_web

import (
	"fmt"
	"net/http"
	"testing"
)

func TestServer(t *testing.T) {
	var server http.Server = http.Server{
		Addr: "localhost:8080",
	}

	var err error = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

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
