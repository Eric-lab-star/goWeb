package controller

import (
	"net/http"
)

type User struct {
	Template Template
}

func (user User) Render(w http.ResponseWriter, r *http.Request) {
	user.Template.ExecuteTempl(w, nil)
}
