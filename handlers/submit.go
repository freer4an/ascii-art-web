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
	Banner string
}

var test string

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
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// get banner from HTML form -> <radio> --> checked radio button
	banner := r.FormValue("banners")
	if !r.Form.Has("banners") {
		ErrorHandler(w, r, http.StatusBadRequest)
		return
	}

	// get text from HTML form -> <textarea> --> value of field
	text := r.FormValue("text")
	if !r.Form.Has("text") {
		ErrorHandler(w, r, http.StatusBadRequest)
		return
	}

	// by default banner is standard
	path := "./ascii-art/samples/" + banner + ".txt"

	// read the data of the banner that was selected
	samples := steps.GetSamples(path)

	// if banner is not exists
	if samples == nil {
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// get result and status which is returns 1 if string contains only ASCII characters else 0
	result, status := steps.PrintSamples(text, samples)
	if status == 0 {
		ErrorHandler(w, r, http.StatusBadRequest)
		return
	}
	test = result

	// save result into data
	data := Ascii{Text: text, Result: result, Banner: banner}

	// execute the HTML with new data
	err = tmpl.Execute(w, data)
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}
	fmt.Println("OK 200")
}
