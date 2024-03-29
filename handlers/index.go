package handlers

import (
	"fmt"
	"html/template"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	// parse and save index.html in tmpl
	tmpl, err := template.ParseFiles("./templates/index.html")
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}
	// to make sure that request method is GET
	if r.Method != "GET" {
		ErrorHandler(w, r, http.StatusBadRequest)
	}

	// if you manually want to go to non-existent addresses handle 'error' page
	if r.URL.Path != "/" {
		ErrorHandler(w, r, http.StatusNotFound)
		return
	}

	// just execute the HTML data without changes
	err = tmpl.Execute(w, nil)
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	fmt.Println("OK 200")
}
