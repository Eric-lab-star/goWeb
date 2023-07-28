package controller

import (
	"net/http"

	"github.com/Eric-lab-star/webDev/views"
)

type User struct {
	Template Template
}

type Template struct {
	New views.Template
}

func (user User) Render(w http.ResponseWriter, r *http.Request) {
	user.Template.New.ExecuteTempl(w, nil)
}
