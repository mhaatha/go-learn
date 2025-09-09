package main

import (
	"html/template"
	"log"
	"net/http"
	"path"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	filePath := path.Join("templates", "index.html")
	tmpl := template.Must(template.ParseFiles(filePath))

	data := map[string]interface{}{
		"Title": "Halo dari Go!",
		"Name":  "Hafidz",
	}

	tmpl.Execute(w, data)
}

func main() {
	http.HandleFunc("/", indexHandler)

	// Serving static file (CSS/JS)
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	log.Println("Server started at localhost:5410")
	http.ListenAndServe(":5410", nil)
}
