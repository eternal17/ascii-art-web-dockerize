package main

import (
	"fmt"
	"log"
	"net/http"
)

func inputHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	entertext := r.FormValue("entertext")

	fmt.Fprintf(w, "UserInput = %s\n", entertext)
}

func main() {
	// opening a port and listening for instructions
	server := http.FileServer(http.Dir("./static"))

	http.Handle("/", server)
	http.HandleFunc("/results", inputHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
		///dd
	}
}
