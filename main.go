package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func executeTempl(w http.ResponseWriter, filePath string) {
	tmp, err := template.ParseFiles(filePath)
	if err != nil {
		log.Printf("parsing template: %v", err)
		http.Error(w, "There was an error parsing the template", http.StatusInternalServerError)
		return
	}
	err = tmp.Execute(w, nil)
	if err != nil { // makes error when passed in data field doesn't exist
		log.Printf("executing template: %v", err)
		http.Error(w, "There was an error executing the template", http.StatusInternalServerError)
		return
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tmplPath := filepath.Join("templates", "home.gohtml")
	executeTempl(w, tmplPath)
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tmplPath := filepath.Join("templates", "contact.gohtml")
	executeTempl(w, tmplPath)
}

func faqPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>FAQ</h1>\n<ul><li>hey!</li></ul>")
}

func page404(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "<h1>Page Not Found </h1>", http.StatusNotFound)
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.Get("/faq", faqPage)
	r.NotFound(page404)
	fmt.Println("Server is running on http://localhost:3000 ")
	http.ListenAndServe("localhost:3000", r)
}
