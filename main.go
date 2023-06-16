package main

import (
	"net/http"
)

func main() {
	tableData := &TableData{}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		HandleMain(w, r, tableData)
	})

	http.HandleFunc("/javascript/number.js", HandleNumberJS)

	http.HandleFunc("/reset", func(w http.ResponseWriter, r *http.Request) {
		HandleReset(w, r, tableData)
	})

	fs := http.FileServer(http.Dir("view/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":8080", nil)
}
