package main

import (
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type PageData struct {
	Title string
}

type NumberData struct {
	Number int
	First  int
	Second int
	Third  int
}

type TableData struct {
	Numbers []NumberData
}

var tableData TableData

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("view/main.html"))

		data := PageData{
			Title: "Checkout!",
		}

		err := tmpl.ExecuteTemplate(w, "main.html", data)
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

	http.HandleFunc("/process", processNumber)
	http.HandleFunc("/reset", resetTable)

	http.ListenAndServe(":8080", nil)
}

func generateRandomNumber() int {
	rand.Seed(time.Now().UnixNano())
	return 1 + rand.Intn(60)
}

func processNumber(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	numberStr := r.Form.Get("number")
	number, err := strconv.Atoi(numberStr)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	first := generateRandomNumber()
	second := generateRandomNumber()
	third := generateRandomNumber()

	newNumber := NumberData{
		Number: number,
		First:  first,
		Second: second,
		Third:  third,
	}

	tableData.Numbers = append(tableData.Numbers, newNumber)

	tmpl := template.Must(template.ParseFiles("view/table.html"))

	err = tmpl.ExecuteTemplate(w, "table", tableData)
	if err != nil {
		log.Println("Failed to render template:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func resetTable(w http.ResponseWriter, r *http.Request) {
	tableData.Numbers = nil
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
