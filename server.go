package main

import (
	"bufio"
	"fmt"
	"html/template"
	"log"
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
	//********************************************************
	z := (testReturn.ban1 + ".txt")
	fmt.Println(z)

	file, err := os.Open("standard.txt")
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

	//	x := Newline(testReturn.textbox, asciiChrs)

	//	fmt.Println("This is X:", x)

	//fmt.Printf("%v", testReturn)
	//fmt.Fprintf(w, x)

	// s := []byte("helloworld")

	// t, _ := template.ParseFiles("static/process.html")
	// t.Execute(w, s)
	tpl.ExecuteTemplate(w, "process.html", nil)

}

func main() {
	// opening a port and listening for instructions
	//server := http.FileServer(http.Dir("./static"))

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/ascii-art", processHandler)
	//	http.HandleFunc("/banner/", bannerHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))

}

func Newline(n string, y map[int][]string) string {
	// prints horizontally
	var ascart []string

	for j := 0; j < len(y[32]); j++ {
		for _, letter := range n {
			//fmt.Print(y[int(letter)][j])
			ascart = append(ascart, (y[int(letter)][j]))
		}
		fmt.Println()
	}
	return "a" //strings.Join(ascart, "")
}
