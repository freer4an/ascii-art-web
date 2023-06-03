package main

import (
	"ascii-art-web/handlers"
	"fmt"
	"net/http"
)

func main() {
	port := "8080"

	// serve CSS file
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// handle 'home' page
	http.HandleFunc("/", handlers.Index)

	// handle 'result' page
	http.HandleFunc("/ascii-art", handlers.SubmitForm)

	// handle 'download' page
	http.HandleFunc("/download", handlers.Download)

	fmt.Println("Listening server: http://localhost:" + port + "/")
	fmt.Println("Press Ctrl + C to kill running port")

	// listen the local server at 'port'
	http.ListenAndServe(":"+port, nil)
}
