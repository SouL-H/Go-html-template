package main

import (
	"html/template"
	"io/ioutil"
	"net/http"
)

type Main struct {
	Title   string
	Header  string
	Err     string
	Message string
}

func loadFile(file string) (string, error) {
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	main := Main{
		Title:   "Hello World",
		Header:  "Golang Example",
		Err:     "Not err",
		Message: "First Go Web Page",
	}
	p, _ := template.ParseFiles("index.html")
	p.Execute(w, main)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
