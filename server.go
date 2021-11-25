package main

import (
	"log"
	"net/http"
)

func main() {
	// opening a port and listening for instructions
	server := http.FileServer(http.Dir("./static"))

	http.Handle("/", server)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
