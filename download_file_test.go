package learngoweb

import (
	"fmt"
	"net/http"
	"testing"
)

var DownloadFile http.HandlerFunc = func(rw http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("file")

	if name == "" {
		rw.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(rw, "upss file not found")
		return
	}

	rw.Header().Add("Content-Disposition", `attachment; filename="`+name+`"`)
	http.ServeFile(rw, r, "resources/"+name)
}

func TestDownloadFile(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: DownloadFile,
	}

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
