package learngoweb

import (
	"fmt"
	"net/http"
	"testing"
)

type logMiddleware struct {
	Handle http.Handler
}

func (mid *logMiddleware) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("log middleware executed")
	mid.Handle.ServeHTTP(rw, r)
}

type errorMiddleware struct {
	Handle http.Handler
}

func (mid *errorMiddleware) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	defer func() {
		if msg := recover(); msg != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(rw, "error %s", msg)
		}
	}()

	fmt.Println("error middleware executed")
	mid.Handle.ServeHTTP(rw, r)
}

func TestMiddleware(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprint(rw, "hello middleware")
	})
	mux.HandleFunc("/foo", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprint(rw, "hello foo")
	})
	mux.HandleFunc("/panic", func(rw http.ResponseWriter, r *http.Request) {
		panic("error nih bro")
	})

	log := logMiddleware{Handle: mux}
	er := errorMiddleware{Handle: &log}

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: &er,
	}

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
