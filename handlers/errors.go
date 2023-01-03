package handlers

import (
	"fmt"
	"html/template"
	"net/http"
)

type ErrorMsg struct {
	ErrMsg string
}

func ErrorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	data := &ErrorMsg{}
	if status == http.StatusNotFound {
		err := "Page not found 404 " + r.URL.Path[1:]
		fmt.Println(err)
		data = &ErrorMsg{ErrMsg: err}
	}
	if status == http.StatusBadRequest {
		err := "Bad request 400"
		fmt.Println(err)
		data = &ErrorMsg{ErrMsg: err}
	}
	tmpl, err := template.ParseFiles("./templates/error.html")
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println(err)
		return
	}
}
