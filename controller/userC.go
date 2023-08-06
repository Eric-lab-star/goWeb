package controller

import (
	"fmt"
	"net/http"
)

type User struct {
	Template Template
}

func (user User) Render(w http.ResponseWriter, r *http.Request) {
	user.Template.ExecuteTempl(w, nil)
}

func (user User) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, r.FormValue("email"))
	fmt.Fprintln(w, r.FormValue("password"))

}
