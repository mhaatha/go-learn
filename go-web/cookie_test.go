package go_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SetCookie(w http.ResponseWriter, r *http.Request) {
	cookie := new(http.Cookie)
	cookie.Name = "X-MHA-Name"
	cookie.Value = r.URL.Query().Get("name")
	cookie.Path = "/"

	http.SetCookie(w, cookie)
	fmt.Fprint(w, "Success create cookie")
}

func GetCookie(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("X-MHA-Name")
	if err != nil {
		http.Error(w, "can't get cookie", http.StatusBadRequest)
		return
	} else {
		name := cookie.Value
		fmt.Fprintf(w, "Hello %s", name)
	}
}

func TestCookie(t *testing.T) {
	mux := http.NewServeMux()

	mux.HandleFunc("/set-cookie", SetCookie)
	mux.HandleFunc("/get-cookie", GetCookie)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	if err := server.ListenAndServe(); err != nil {
		t.Errorf("can't start a server: %v", err)
	}
}

func TestSetCookie(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "localhost:8080/?name=Hafidz", nil)
	recorder := httptest.NewRecorder()

	SetCookie(recorder, request)

	cookies := recorder.Result().Cookies()

	for _, cookie := range cookies {
		fmt.Printf("Cookie %s: %s\n", cookie.Name, cookie.Value)
	}
}

func TestGetCookie(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "localhost:8080/", nil)
	cookie := new(http.Cookie)
	cookie.Name = "X-MHA-Name"
	cookie.Value = "Hafidz"
	request.AddCookie(cookie)

	recorder := httptest.NewRecorder()

	GetCookie(recorder, request)
	body, _ := io.ReadAll(recorder.Result().Body)

	fmt.Println(string(body))
}
