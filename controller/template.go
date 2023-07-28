package controller

import "net/http"

type Template interface {
	ExecuteTempl(w http.ResponseWriter, data interface{})
}
