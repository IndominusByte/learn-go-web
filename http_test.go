package learngoweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

var handlerHome http.HandlerFunc = func(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "Hello World")
}

var handlerHi http.HandlerFunc = func(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, r.RequestURI)
	fmt.Fprintf(rw, "Hi bro")
}

func TestHttp(t *testing.T) {
	var mux = http.NewServeMux()
	mux.HandleFunc("/", handlerHome)
	mux.HandleFunc("/hi", handlerHi)

	request := httptest.NewRequest(http.MethodGet, "/", nil)
	recorder := httptest.NewRecorder()

	handlerHome(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(response.StatusCode)
	fmt.Println(string(body))

	request = httptest.NewRequest(http.MethodGet, "/hi?hello=true", nil)
	recorder = httptest.NewRecorder()

	handlerHi(recorder, request)

	response = recorder.Result()
	body, _ = io.ReadAll(response.Body)
	fmt.Println(response.StatusCode)
	fmt.Println(string(body))
}
