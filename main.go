package main

import (
	"ascii-art-web/ascii"
	"fmt"
	"net/http"
	"strings"
	"text/template"
)

var tmpl *template.Template

type PageData struct {
	Result string
}

func main() {
	// Parse all the .html files
	var err error
	tmpl, err = template.ParseGlob("templates/*.html")
	if err != nil {
		fmt.Println("Error at parsing", err)
	}

	// Handle file in style folder
	fileServer := http.FileServer(http.Dir("./style/"))
	http.Handle("/style/", http.StripPrefix("/style/", fileServer))

	// Handle all the .html files
	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/ascii", AsciiHandler)

	// Launch the server
	fmt.Println("(http://localhost:8080) - Server started on port", ":8080")
	http.ListenAndServe(":8080", nil)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// handle error 404
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	// Execute the home.html file
	tmpl.ExecuteTemplate(w, "home.html", nil)
}

func AsciiHandler(w http.ResponseWriter, r *http.Request) {
	// Execute the AsciiArt function, with the text and banner chose
	text := r.PostFormValue("text")
	text = strings.ReplaceAll(text, "\r", "\n")
	banner := r.PostFormValue("template")
	result := ascii.AsciiArt(text, banner)
	data := PageData{Result: result}
	err := tmpl.ExecuteTemplate(w, "home.html", data)
	if err != nil {
		//Handle error 500
		http.Error(w, "500 Status Internal Server Error", http.StatusInternalServerError)
		return
	}
}
