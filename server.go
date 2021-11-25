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
	age := r.FormValue("age")
	fmt.Fprintf(w, "Age = %s\n", age)

}

func main() {
	// opening a port and listening for instructions
	server := http.FileServer(http.Dir("./static"))

	http.Handle("/", server)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", inputHandler)
}
