package main

import (
	"./link"
	"html/template"
	"io"
	"net/http"
	"strings"
)

type errorHandler func(http.ResponseWriter, *http.Request) error

func (fn errorHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := fn(w, r); err != nil {
		w.Header().Set("Content-Type", "text/plain")
		http.Error(w, err.Error(), 400)
	}
}

func respond(wr io.Writer, url string) error {
	t, err := template.New("").Parse(`Click to visit <a href="{{.}}">{{.}}</a>`)
	if err != nil {
		return err
	}

	return t.Execute(wr, url)
}

func redirect(w http.ResponseWriter, r *http.Request) error {
	encoded := strings.TrimPrefix(r.URL.Path, "/r/")

	link, err := link.Parse(encoded)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "text/html")
	return respond(w, link.Target)
}

func main() {
	http.Handle("/r/", errorHandler(redirect))
	http.ListenAndServe(":12345", nil)
}
