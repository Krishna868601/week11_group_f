package main

import (
	"log"
	"net/http"
)

func main() {

	// handle `/` route
	fs := http.FileServer(http.Dir("./templates"))
	http.Handle("/", fs)
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("./static/js"))))

	// run server on port "80"
	log.Fatal(http.ListenAndServe(":80", nil))

}
