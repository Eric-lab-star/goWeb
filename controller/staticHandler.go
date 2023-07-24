package controller

import (
	"net/http"

	"github.com/Eric-lab-star/webDev/views"
)

func StaticHandler(tmpl views.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl.ExecuteTempl(w, nil)
	}
}

func FAQ(tmpl views.Template) http.HandlerFunc {
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
