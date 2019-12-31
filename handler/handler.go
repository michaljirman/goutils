package handler

import (
	"io"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request){
	io.WriteString(w, "<html><body>Hello!</body></html>")
}
