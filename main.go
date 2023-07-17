package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Welcome to my awesome site!</h1>")

}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Contact Page</h1><p>To get in touch, email me<a href=\"mailto:cyon256@icloud.com\"> cyon256@icloud.com</a>")
}

func page404(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(404)
	fmt.Fprint(w, "<h1>Page Not Found</h1>")
}

type router struct {
}

func (router router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, r.URL.Path)
	fmt.Fprintln(w, r.URL.RawPath)
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	default:
		page404(w, r)
	}
}

func main() {
	var router router
	fmt.Println("Server is running ")
	http.ListenAndServe("localhost:3000", router)
}
