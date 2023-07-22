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
