package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Banner struct {
	Title string
	Ban1  string
	Ban2  string
	Ban3  string
}

func bannerHandler(w http.ResponseWriter, r *http.Request) {
	p := Banner{
		Title: "SELECT BANNERFILE\n",
		Ban1:  "Shadow\n",
		Ban2:  "Standard\n",
		Ban3:  "Thinkertoy\n",
	}
	t, _ := template.ParseFiles("static/index.html")
	t.Execute(w, p)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	fmt.Fprintf(w, "<h1>Step by Step</h1>")
}

func main() {
	// opening a port and listening for instructions
	//server := http.FileServer(http.Dir("./static"))

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/banner", bannerHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)

	}
}
