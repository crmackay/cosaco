package main

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"time"
)

type testStruct struct {
	Name  string
	Value string
}

func prepTemplates() *bytes.Buffer {
	t, err := template.ParseFiles("index.html", "header.html")
	if err != nil {
		log.Println(err)
	}
	b := new(bytes.Buffer)

	data := testStruct{Name: "myName", Value: "myValue"}
	t.Execute(b, data)

	return b
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tm := time.Now().Format(time.RFC1123)
	w.Write([]byte("The time is: " + tm))
}

func teamsHandler(w http.ResponseWriter, r *http.Request) {
	t := prepTemplates()
	w.Write(t.Bytes())
}

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", homeHandler)

	mux.HandleFunc("/teams/", teamsHandler)

	log.Println("Listening...")
	http.ListenAndServe(":3000", mux)

}
