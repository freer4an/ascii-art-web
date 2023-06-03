package handlers

import (
	"fmt"
	"net/http"
	"strconv"
)

func Download(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	// Set size of reponse body which is file.txt's size
	w.Header().Set("Content-Length", strconv.Itoa(len(test)))

	// Set Content-Type to .txt file
	w.Header().Set("Content-Type", "text/plain")

	// Set Content-Disposition to 'attachment' (indicating it should be downloaded)
	w.Header().Set("Content-Disposition", "attachment; filename=file.txt")
	// Copy the contents of the file to the response writer
	_, err := w.Write([]byte(test))
	if err != nil {
		fmt.Fprintf(w, "Something went wrong!")
	}
}
