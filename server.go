package main

import (
	"bufio"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strings"
)

type Banner struct {
	Title   string
	Ban1    string
	Ban2    string
	Ban3    string
	String1 string
	String2 string
}

// type toAscii struct {
// 	String1 string
// 	String2 string
// }

// type Output struct {
// 	Banner  Banner
// 	toAscii toAscii
// }

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

	var bannerfile string

	if testReturn.ban1 == "Thinkertoy" {
		bannerfile = "Thinkertoy.txt"
	} else if testReturn.ban1 == "Shadow" {
		bannerfile = "Shadow.txt"
	} else {
		bannerfile = "Standard.txt"
	}

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

	// convert textbox to bytes to figure out where linebreak is (10)
	b := []byte(testReturn.textbox)
	count := 0
	for _, num := range b {
		count++
		if num == 10 {
			break
		}
	}

	// checking if there is linebreak in string, returning the string seperated on 2 lines if there is
	// 2nd line is an empty string if there isnt a line break
	if strings.Contains(testReturn.textbox, "\n") {
		p := Banner{
			Title:   "SELECT BANNERFILE",
			Ban1:    "Shadow",
			Ban2:    "Standard",
			Ban3:    "Thinkertoy",
			String1: Newline(testReturn.textbox[:count-2], asciiChrs),
			String2: Newline(testReturn.textbox[count:], asciiChrs),
		}

		err := tpl.ExecuteTemplate(w, "index.html", p)
		fmt.Println(err)
		// if err := tpl.ExecuteTemplate(w, "index.html", Y); err != nil {
		// 	http.Error(w, err.Error(), http.StatusInternalServerError)
		// }
	} else {
		p := Banner{
			Title:   "SELECT BANNERFILE",
			Ban1:    "Shadow",
			Ban2:    "Standard",
			Ban3:    "Thinkertoy",
			String1: Newline(testReturn.textbox, asciiChrs),
			String2: "",
		}

		err := tpl.ExecuteTemplate(w, "index.html", p)
		fmt.Println(err)
		// if err := tpl.ExecuteTemplate(w , "index.html", Y); err != nil {
		// 	http.Error(w, err.Error(), http.StatusInternalServerError)
		// }
	}
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
