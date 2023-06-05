package main

import (
	"html/template"
	"log"
	"net/http"
)

type PageData struct {
	Title string
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("view/main.html"))

		data := PageData{
			Title: "Checkout!",
		}

		err := tmpl.Execute(w, data)
		if err != nil {
			log.Println("Failed to render template:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	})

	http.HandleFunc("/javascript/number.js", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "view/javascript/number.js")
	})

	fs := http.FileServer(http.Dir("view/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":8080", nil)
}
