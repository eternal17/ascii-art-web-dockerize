package main

import (
	"bufio"
	"fmt"
	"html/template"
	"net/http"
	"os"
)

type Banner struct {
	Title string
	Ban1  string
	Ban2  string
	Ban3  string
}

var tpl *template.Template

func indexHandler(w http.ResponseWriter, r *http.Request) {
	p := Banner{
		Title: "SELECT BANNERFILE",
		Ban1:  "Shadow",
		Ban2:  "Standard",
		Ban3:  "Thinkertoy",
	}

	// handling any pages that are not the index or ascii-art (404 error)
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	tpl.ExecuteTemplate(w, "index.html", p)
}

func processHandler(w http.ResponseWriter, r *http.Request) {
	getban1 := r.FormValue("banner")
	getban2 := r.FormValue("banner")
	getban3 := r.FormValue("banner")
	tbox := r.FormValue("textbox")

	// handling bad request status code
	if len(getban1) == 0 && len(getban2) == 0 && len(getban3) == 0 || len(tbox) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

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
	//********************************************************

	fmt.Println("This is tban1:", testReturn.ban1)
	h := testReturn.ban1 + ".txt"
	fmt.Println("This is h:", h)

	var bannerfile string
	fmt.Println("This is tRB1:", testReturn.ban1)

	if testReturn.ban1 == "Thinkertoy" {
		bannerfile = "Thinkertoy.txt"
	} else if testReturn.ban1 == "Shadow" {
		bannerfile = "Shadow.txt"
	} else {
		bannerfile = "Standard.txt"
	}

	fmt.Println("This is Z:", bannerfile)

	file, err := os.Open(bannerfile)
	if err != nil {
		fmt.Println("Usage: go run . [STRING] [BANNER]")
		fmt.Println("EX: go run . something standard")
		os.Exit(0)
	}
	defer file.Close()

	scanned := bufio.NewScanner(file) // reading file
	scanned.Split(bufio.ScanLines)

	var lines []string

	for scanned.Scan() {
		lines = append(lines, scanned.Text())
	}

	file.Close()

	asciiChrs := make(map[int][]string)
	id := 31

	for _, line := range lines {
		if string(line) == "" {
			id++
		} else {
			asciiChrs[id] = append(asciiChrs[id], line)
		}
	}

	x := Newline(testReturn.textbox, asciiChrs)

	fmt.Fprintf(w, x)

	tpl.ExecuteTemplate(w, "process.html", x)
}

func main() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/ascii-art", processHandler)
	http.ListenAndServe(":8080", nil)
}

func Newline(n string, y map[int][]string) string {
	// prints horizontally
	var empty string

	// prints horizontally
	for j := 0; j < len(y[32]); j++ {
		var line string
		for _, letter := range n {
			line = line + string((y[int(letter)][j]))
		}
		empty += line + "\n"
		line = ""
	}
	return empty
}
