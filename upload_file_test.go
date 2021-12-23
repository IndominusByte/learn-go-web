package learngoweb

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/IndominusByte/magicimage"
)

var Upload http.HandlerFunc = func(rw http.ResponseWriter, r *http.Request) {
	magic := magicimage.New(r, 32<<20)
	// magic.SetMaxFileSize(1 << 20)
	// err := magic.ValidateSingleImage("file")
	err := magic.ValidateMultipleImage("files")
	if err != nil {
		fmt.Fprint(rw, err)
	}

	err = magic.SaveImages(300, 200, "files", false)
	if err != nil {
		fmt.Fprint(rw, err)
	}
}

func TestUpload(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/upload", Upload)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
