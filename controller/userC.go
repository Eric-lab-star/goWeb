package controller

import (
	"fmt"
	"net/http"
)

type User struct {
	Template Template
}

func (user User) Render(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email string
	}
	data.Email = r.FormValue("email")
	user.Template.ExecuteTempl(w, data)
}

func (user User) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, r.FormValue("email"))
	fmt.Fprintln(w, r.FormValue("password"))

}
