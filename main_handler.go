package main

import (
	"log"
	"net/http"
)

type PageData struct {
	Title string
}

func HandleMain(w http.ResponseWriter, r *http.Request, tableData *TableData) {
	tmpl := parseTemplate("view/main.html")

	if r.Method == "POST" {
		HandlePostRequest(w, r, tableData)
	}

	data := struct {
		PageData
		TableData
	}{
		PageData:  PageData{Title: "Checkout!"},
		TableData: *tableData,
	}

	err := tmpl.ExecuteTemplate(w, "main.html", data)
	if err != nil {
		log.Println("Failed to render template:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
