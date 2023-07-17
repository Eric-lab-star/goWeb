package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Home </h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Contact Page</h1><p>To get in touch, email me<a href=\"mailto:cyon256@icloud.com\"> cyon256@icloud.com</a>")
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
	fmt.Println("Server is running ")
	http.ListenAndServe("localhost:3000", r)
}
