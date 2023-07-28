package views

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/Eric-lab-star/webDev/templates"
)

type Template struct {
	htmlTmpl *template.Template
}

func Must(t Template, err error) Template {
	if err != nil {
		panic(err)
	}
	return t
}

func ParseFS(patterns ...string) (Template, error) {
	tmpl, err := template.ParseFS(templates.FS, patterns...)
	if err != nil {
		err = fmt.Errorf("parsing template error %w", err)
		return Template{}, err
	}
	return Template{tmpl}, nil
}

func Parse(path string) (Template, error) {
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		return Template{}, fmt.Errorf("Error Parsing Template: %v", err)
	}
	return Template{tmpl}, nil
}

func (t Template) ExecuteTempl(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err := t.htmlTmpl.Execute(w, data)
	if err != nil { // makes error when passed in data field doesn't exist
		log.Printf("executing template: %v", err)
		http.Error(w, "There was an error executing the template", http.StatusInternalServerError)
		return
	}
}
