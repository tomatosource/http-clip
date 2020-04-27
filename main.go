package main

import (
	"bytes"
	"io"
	"net/http"

	"github.com/atotto/clipboard"
)

func main() {
	handleCopy := func(w http.ResponseWriter, r *http.Request) {
		buf := new(bytes.Buffer)
		buf.ReadFrom(r.Body)
		in := buf.String()
		clipboard.WriteAll(in)
	}
	handlePaste := func(w http.ResponseWriter, _ *http.Request) {
		out, _ := clipboard.ReadAll()
		io.WriteString(w, out)
	}
	http.HandleFunc("/copy", handleCopy)
	http.HandleFunc("/paste", handlePaste)
	http.ListenAndServe(":8833", nil)
}
