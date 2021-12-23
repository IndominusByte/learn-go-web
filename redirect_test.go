package learngoweb

import (
	"fmt"
	"net/http"
	"testing"
)

var RedirectFrom http.HandlerFunc = func(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprint(rw, "Hello Redirect")
}

var RedirectTo http.HandlerFunc = func(rw http.ResponseWriter, r *http.Request) {
	http.Redirect(rw, r, "/redirect-from", http.StatusTemporaryRedirect)
}

var RedirectOut http.HandlerFunc = func(rw http.ResponseWriter, r *http.Request) {
	http.Redirect(rw, r, "https://creentproduction.com", http.StatusTemporaryRedirect)
}

func TestRedirect(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/redirect-to", RedirectTo)
	mux.HandleFunc("/redirect-from", RedirectFrom)
	mux.HandleFunc("/redirect-out", RedirectOut)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
