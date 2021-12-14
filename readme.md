# ASCII ART WEB

#### A project by eternal17 & tb38r
---
</br>

### <ins>Table of Contents</ins>

1.    Description
2.    Authors
3.    Usage: how to run
4.    Implementation details: algorithm
</br>

---

### <ins>Description</ins>

Ascii art web uses Golang's HTML/Templates & net/HTTP packages to create an API that creates a static server & listens and responds to . An accompanying HTML file illustrates the client and provides the required inputs.
</br>
The primary aim is take a string from the user and return it to the client in Ascii form.
</br>

---

### <ins>Authors</ins>
tb38r: https://git.learn.01founders.co/tb38r
</br>
eternal17: https://git.learn.01founders.co/eternal17

---
</br>

### <ins>Usage: How to run</ins>
</br>

1. Within the project path: "go run .", to run the server. Server listens to localhost:8080.
2. In the web browser, go to localhost:8080
3. Here the user can chose between 3 banners, shadow, standard and thinkertoy.
4. In the text area below, the user can input up to 2 lines
5. After pressing submit, the ascii art version is returned with the banner chosen.
</br>
</br>
<img src="ascii.gif" alt="ascii gif" width="500" height="300">

---
</br>

### <ins>Implementation Details: Algorithm</ins>
</br>

First, a brief summary of the *imports used within this project. 
</br>
*Bufio implements a buffered I/O which allows for a range of manipulations. In this instance, we've used it to scan the text within a given file and returned the output as a slice of string.
</br>
*Fmt implements formatted I/O.
</br>
*Html/template implements data-driven templates for generating HTML output safe against code injection.
</br>
*Net/http provides HTTP client and server implementations
</br>
*Os import here is used to open  files as well as some minor error handling
</br>
Lastly, *Strings implements simple functions to manipulate UTF-8 encoded strings. We use its contains function to search for specific parameters of the user input.
</br>
__________________

</br>

The ListenandServe function on line 154 starts an HTTP server, listens on port 8080 incoming requests, and serves on '/ ' . (see Usage above)
</br>
</br>
The http.HandleFunc(handlers) functions in Lines 152 & 153 respond to the HTTP request and register the corresponding functions with the HTTP server.
(Handler functions are a convenient way of creating handlers which allow us to build web applications that are more modular.)
</br>
</br>
We initialised tp1 with the Must helper for effeciency & brevity.
</br>
</br>
Also of note, is the function Newline, called in from a previous project. Newline() takes in a slice of string(s) and returns it illustrated in ascii form. It's limited to the 128 characters of the ascii table.
</br>
</br>
______
</br>
Part 1 of the usage guidelines occurs on line 154. The range of banner options are stored within the struct p seen in line 25. The processHandler function (line 41) initially stores the requested banner type and string using the *Formvalue function. These are then stored within a struct to be parsed later for manipution. 
</br>
</br>
*Os (line 67) opens the desired banner type by appending ".txt" to it. The characters within are scanned and consequently stored within a map as a slice of string(s) (line 75 -95).
</br>
</br>
We use a simple loop (lines 98-103) to check if the user input has a line break and then pass either the whole string or a split version (contingent on the aforementioned check) to be executed. The index.html file has multiple pointers now with values passed through, the api  returns all of the above back to the client and awaits a new request. 

