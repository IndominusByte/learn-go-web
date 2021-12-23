package learngoweb

import (
	"embed"
	"io/fs"
	"net/http"
	"testing"
)

func TestFileServer(t *testing.T) {
	ServeDir := http.FileServer(http.Dir("resources"))

	mux := http.NewServeMux()
	// without stripprefix = resources/static/<file>
	mux.Handle("/static/", http.StripPrefix("/static/", ServeDir))

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}

//go:embed resources/*
var resources embed.FS

func TestFileServerEmbed(t *testing.T) {
	dir, _ := fs.Sub(resources, "resources")
	ServeDir := http.FileServer(http.FS(dir))

	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static/", ServeDir))

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}

}
