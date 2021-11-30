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

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("static/*.html"))
}

// func bannerHandler(w http.ResponseWriter, r *http.Request) {

// 	p := Banner{
// 		Title: "SELECT BANNERFILE\n",
// 		Ban1:  "Shadow\n",
// 		Ban2:  "Standard\n",
// 		Ban3:  "Thinkertoy\n",
// 	}

// 	//tpl.ExecuteTemplate(w, "index.html", nil)

// 	t, _ := template.ParseFiles("static/index.html")
// 	t.Execute(w, p)

// }

func indexHandler(w http.ResponseWriter, r *http.Request) {

	p := Banner{
		Title: "SELECT BANNERFILE\n",
		Ban1:  "Shadow\n",
		Ban2:  "Standard\n",
		Ban3:  "Thinkertoy\n",
	}

	//tpl.ExecuteTemplate(w, "index.html", nil)

	t, _ := template.ParseFiles("static/index.html")
	t.Execute(w, p)
}

func processHandler(w http.ResponseWriter, r *http.Request) {
	// if r.Method != "POST" {
	// 	http.Redirect(w, r, "/", http.StatusSeeOther)
	// 	return
	// }

	getban1 := r.FormValue("banner")
	getban2 := r.FormValue("banner")
	getban3 := r.FormValue("banner")
	tbox := r.FormValue("textbox")

	testReturn := struct {
		ban1    string
		ban2    string
		ban3    string
		textbox string
	}{
		ban1:    getban1,
		ban2:    getban2,
		ban3:    getban3,
		textbox: tbox,
	}

	fmt.Printf("%v", testReturn)
	fmt.Fprintf(w, testReturn.textbox)

	// s := []byte("helloworld")

	// t, _ := template.ParseFiles("static/process.html")
	// t.Execute(w, s)

}

func main() {
	// opening a port and listening for instructions
	//server := http.FileServer(http.Dir("./static"))

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/ascii-art", processHandler)
	//	http.HandleFunc("/banner/", bannerHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))

}
