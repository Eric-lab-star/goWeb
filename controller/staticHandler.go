package controller

import (
	"net/http"
)

func StaticHandler(tmpl Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl.ExecuteTempl(w, nil)
	}
}

func FAQ(tmpl Template) http.HandlerFunc {
	data := []struct {
		Question string
		Answer   string
	}{
		{"what is your name", "My name is Kim"},
		{"How old are you", "I am 27"},
		{"Where do you live", "I am living in seoul"},
	}

	return func(w http.ResponseWriter, r *http.Request) {
		tmpl.ExecuteTempl(w, data)
	}
}
