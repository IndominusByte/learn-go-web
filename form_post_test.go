package learngoweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var HandlerPost http.HandlerFunc = func(rw http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		panic(err)
	}

	first_name := r.PostForm.Get("first_name")
	last_name := r.PostForm.Get("last_name")

	fmt.Fprintf(rw, "Hello %s %s", first_name, last_name)
}

func TestPostForm(t *testing.T) {
	requestBody := strings.NewReader("first_name=oman&last_name=pradipta")
	request := httptest.NewRequest(http.MethodPost, "/", requestBody)
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	response := httptest.NewRecorder()

	HandlerPost(response, request)

	result := response.Result()
	body, _ := io.ReadAll(result.Body)

	fmt.Println(string(body))
}
