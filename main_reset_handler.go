package main

import (
	"net/http"
)

func HandleReset(w http.ResponseWriter, r *http.Request, tableData *TableData) {
	tableData.Reset()
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
