package main

import (
	"net/http"
	"strconv"
)

func HandlePostRequest(w http.ResponseWriter, r *http.Request, tableData *TableData) {
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
}
