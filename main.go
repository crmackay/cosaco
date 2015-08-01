package main

import (
	"bytes"
	//"fmt"
	"html/template"
	"log"
	"net/http"
	//"runtime"
	"time"
)

/*

semi static handler:
	- parse template
	- execute template with data into a Buffer
	- write the buffer to ResponseWriter upon Request
	- watch for changes to the data
	- update buffer accordingly

database (still not sure which one):
	- write method
		- trigger cache update
	- read methods
	- backup method
	- structs or map from data
	- version, schematic
	- reflection of what's in the database


*/

type testStruct struct {
	Name  string
	Value string
}

var t = template.Must(template.ParseFiles("index.html", "header.html"))
var teamPage = new(bytes.Buffer)

func executeTemplate(b *bytes.Buffer) {

	data := testStruct{Name: "myName", Value: "myValue"}

	t.Execute(b, data)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tm := time.Now().Format(time.RFC1123)
	w.Write([]byte("The time is: " + tm))
}

func teamsHandler(w http.ResponseWriter, r *http.Request) {
	w.Write(teamPage.Bytes())
}

func main() {

	//runtime.GOMAXPROCS(runtime.NumCPU())
	//fmt.Println("running on ", runtime.NumCPU(), " CPUs")

	executeTemplate(teamPage)

	mux := http.NewServeMux()

	mux.HandleFunc("/", homeHandler)

	mux.HandleFunc("/teams/", teamsHandler)

	log.Println("Listening...")
	http.ListenAndServe(":3000", mux)

}
