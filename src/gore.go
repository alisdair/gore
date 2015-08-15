package main

import (
	"./link"
	"fmt"
	"net/http"
	"strings"
)

type errorHandler func(http.ResponseWriter, *http.Request) error

func (fn errorHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := fn(w, r); err != nil {
		http.Error(w, err.Error(), 400)
	}
}

func redirect(w http.ResponseWriter, r *http.Request) error {
	encoded := strings.TrimPrefix(r.URL.Path, "/r/")

	link, err := link.Parse(encoded)
	if err != nil {
		return err
	}

	_, err = fmt.Fprintln(w, link.Target)
	return err
}

func main() {
	http.Handle("/r/", errorHandler(redirect))
	http.ListenAndServe(":12345", nil)
}
