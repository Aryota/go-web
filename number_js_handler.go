package main

import (
	"net/http"
)

func HandleNumberJS(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "view/javascript/number.js")
}
