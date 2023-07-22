package main

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/Eric-lab-star/webDev/controller"
	"github.com/Eric-lab-star/webDev/views"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	tmpl := views.Must(views.Parse(filepath.Join("templates", "home.gohtml")))
	r.Get("/", controller.StaticHandler(tmpl))
	tmpl = views.Must(views.Parse(filepath.Join("templates", "contact.gohtml")))
	r.Get("/contact", controller.StaticHandler(tmpl))
	tmpl = views.Must(views.Parse(filepath.Join("templates", "faq.gohtml")))
	r.Get("/faq", controller.StaticHandler(tmpl))
	tmpl = views.Must(views.Parse(filepath.Join("templates", "404.gohtml")))
	r.NotFound(controller.StaticHandler(tmpl))
	fmt.Println("Server is running on http://localhost:3000 ")
	http.ListenAndServe("localhost:3000", r)
}
