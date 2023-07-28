package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/Eric-lab-star/webDev/controller"
	"github.com/Eric-lab-star/webDev/views"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	tmpl := views.Must(views.ParseFS("layout.html", "home.gohtml"))
	r.Get("/", controller.StaticHandler(tmpl))

	tmpl = views.Must(views.ParseFS("layout.html", "contact.gohtml"))
	r.Get("/contact", controller.StaticHandler(tmpl))

	tmpl = views.Must(views.ParseFS("layout.html", "faq.gohtml"))
	r.Get("/faq", controller.FAQ(tmpl))

	tmpl = views.Must(views.ParseFS("404.gohtml"))
	r.NotFound(controller.StaticHandler(tmpl))

	user := controller.User{}
	user.Template.New = views.Must(views.ParseFS("layout.html", "signUp.html"))
	r.Get("/signup", user.Render)

	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	filesDir := http.Dir(filepath.Join(wd, "static"))
	FileServer(r, "/static", filesDir)

	fmt.Println("Server is running on http://localhost:3000 ")
	http.ListenAndServe("localhost:3000", r)
}

func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit any URL parameters.")
	}
	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", http.StatusPermanentRedirect).ServeHTTP)
		path += "/"
	}
	path += "*"
	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(root))
		fs.ServeHTTP(w, r)
	})
}
