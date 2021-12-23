package learngoweb

import (
	_ "embed"
	"fmt"
	"net/http"
	"testing"
)

var HandlerServeFile http.HandlerFunc = func(rw http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Get("name") != "" {
		http.ServeFile(rw, r, "resources/ok.html")
	} else {
		http.ServeFile(rw, r, "resources/notfound.html")
	}
}

//go:embed resources/ok.html
var resourceOK string

//go:embed resources/notfound.html
var resourceNOTFOUND string

var HandlerServeFileEmbed http.HandlerFunc = func(rw http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Get("name") != "" {
		fmt.Fprint(rw, resourceOK)
	} else {
		fmt.Fprint(rw, resourceNOTFOUND)
	}
}

func TestServeFile(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: HandlerServeFile,
	}

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}

func TestServeFileEmbed(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: HandlerServeFileEmbed,
	}

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
