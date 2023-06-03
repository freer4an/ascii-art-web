package handlers

import (
	"fmt"
	"html/template"
	"net/http"
)

type ErrorMsg struct {
	ErrStatus int
	ErrMsg    string
}

func ErrorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	data := &ErrorMsg{}

	switch status {

	case http.StatusNotFound:
		err := "Page not found 404 " + "(" + r.URL.Path[1:] + ")"
		fmt.Println(err)
		data = &ErrorMsg{ErrStatus: status, ErrMsg: err}

	case http.StatusBadRequest:
		err := "Bad request 400"
		fmt.Println(err)
		data = &ErrorMsg{ErrStatus: status, ErrMsg: err}

	case http.StatusInternalServerError:
		err := "Internal Server Error 500"
		fmt.Println(err)
		data = &ErrorMsg{ErrStatus: status, ErrMsg: err}
	}

	tmpl, err := template.ParseFiles("./templates/error.html")
	if err != nil {
		err := "Internal Server Error 500"
		fmt.Fprint(w, err)
		fmt.Println(err)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		err := "Internal Server Error 500"
		fmt.Fprint(w, err)
		fmt.Println(err)
		return
	}
}
