package learngoweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

var RequestHeader http.HandlerFunc = func(rw http.ResponseWriter, r *http.Request) {
	header := r.Header
	fmt.Fprint(rw, header.Get("content-type"))
}

var ResponseHeader http.HandlerFunc = func(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("X-Powered-By", "mentimun-mentah")
	fmt.Fprint(rw, "OK")
}

func TestHeader(t *testing.T) {
	request := httptest.NewRequest(http.MethodPost, "/", nil)
	request.Header.Add("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	RequestHeader(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))

	request = httptest.NewRequest(http.MethodPost, "/", nil)
	recorder = httptest.NewRecorder()

	ResponseHeader(recorder, request)

	response = recorder.Result()
	body, _ = io.ReadAll(response.Body)

	fmt.Println(string(body))
	fmt.Println(response.Header.Get("x-powered-by"))
}
