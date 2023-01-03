package handlers

import (
	"ascii-art-web/ascii-art/pkg/steps"
	"fmt"
	"html/template"
	"net/http"
)

type Ascii struct {
	Text   string
	Result string
}

func SubmitForm(w http.ResponseWriter, r *http.Request) {
	// if you want to go to '/ascii-art' page without submitting data
	// which means method will be "GET" handle 'error' page
	if r.Method != "POST" {
		ErrorHandler(w, r, http.StatusBadRequest)
		return
	}

	// parse and save index.html in tmpl
	tmpl, err := template.ParseFiles("./templates/index.html")

	// if error unhandled show 500 InternalServerError
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// get banner from HTML form -> <radio> --> checked radio button
	banner := r.FormValue("banner")

	// get text from HTML form -> <textarea> --> value of field
	text := r.FormValue("text")

	// by default banner is standard
	path := "./ascii-art/samples/" + banner + ".txt"

	// read the data of the banner that was selected
	samples, contents := steps.GetSamples(path)
	if samples == nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// get result
	result := steps.PrintSamples(text, contents, samples)

	// save result into data
	data := Ascii{Text: text, Result: result}

	// execute the HTML with new data
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println("OK 200")
}
