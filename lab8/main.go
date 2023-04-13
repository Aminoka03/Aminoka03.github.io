package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func infoPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		first_name := r.FormValue("first_name")
		first_name1 := r.FormValue("first_name1")
		last_name := r.FormValue("last_name")
		last_name1 := r.FormValue("last_name1")
		info := map[string]interface{}{"first_name": first_name, "first_name1": first_name1, "last_name": last_name, "last_name1": last_name1}

		fmt.Print(info["last_name1"])

		tmpl, _ := template.ParseFiles("static/wedding.html")

		tmpl.Execute(w, info)

		return
	}
	tmpl, _ := template.ParseFiles("static/info.html")

	tmpl.Execute(w, nil)
}

func weddingPage(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("static/wedding.html")

	tmpl.Execute(w, nil)
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))

	http.Handle("/", fileServer)
	http.HandleFunc("/info", infoPage)
	http.HandleFunc("/wedding", weddingPage)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
