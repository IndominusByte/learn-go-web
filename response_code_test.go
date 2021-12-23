package learngoweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

var HandlerResponseCode http.HandlerFunc = func(rw http.ResponseWriter, r *http.Request) {
	if name := r.URL.Query().Get("name"); name != "" {
		rw.WriteHeader(http.StatusOK)
		fmt.Fprintf(rw, "Hello %s", name)
		return
	}

	rw.WriteHeader(http.StatusBadRequest)
	fmt.Fprint(rw, "name is required!")
}

func TestResponseCodeInvalid(t *testing.T) {
	request := httptest.NewRequest(http.MethodPost, "/", nil)
	response := httptest.NewRecorder()

	HandlerResponseCode(response, request)

	result := response.Result()
	body, _ := io.ReadAll(result.Body)

	fmt.Println(result.StatusCode)
	fmt.Println(result.Status)
	fmt.Println(string(body))
}

func TestResponseCodeValid(t *testing.T) {
	request := httptest.NewRequest(http.MethodPost, "/?name=oman", nil)
	response := httptest.NewRecorder()

	HandlerResponseCode(response, request)

	result := response.Result()
	body, _ := io.ReadAll(result.Body)

	fmt.Println(result.StatusCode)
	fmt.Println(result.Status)
	fmt.Println(string(body))
}
