package learngoweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

var HandlerSetCookie http.HandlerFunc = func(rw http.ResponseWriter, r *http.Request) {
	ck := &http.Cookie{
		Name:  "X-PZN-Name",
		Value: r.URL.Query().Get("name"),
		Path:  "/",
	}

	http.SetCookie(rw, ck)

	fmt.Fprint(rw, "success create cookie")
}

var HandlerGetCookie http.HandlerFunc = func(rw http.ResponseWriter, r *http.Request) {
	ck, err := r.Cookie("X-PZN-Name")
	if err != nil {
		fmt.Fprint(rw, "cookie required")
		return
	}

	fmt.Fprintf(rw, "Welcome back %s", ck.Value)
}

func TestServeCookie(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/set-cookie", HandlerSetCookie)
	mux.HandleFunc("/get-cookie", HandlerGetCookie)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}

func TestSetCookie(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/?name=oman", nil)
	recorder := httptest.NewRecorder()

	HandlerSetCookie(recorder, request)

	result := recorder.Result()
	body, _ := io.ReadAll(result.Body)

	fmt.Println(string(body))

	for _, cookie := range result.Cookies() {
		fmt.Println(cookie.Name, cookie.Value)
	}
}

func TestGetCookie(t *testing.T) {
	// set cookie
	request := httptest.NewRequest(http.MethodGet, "/?name=omanpradipta", nil)
	recorder := httptest.NewRecorder()

	HandlerSetCookie(recorder, request)

	result := recorder.Result()

	// get cookie
	request = httptest.NewRequest(http.MethodGet, "/", nil)
	recorder = httptest.NewRecorder()

	for _, cookie := range result.Cookies() {
		var ck *http.Cookie

		ck = &http.Cookie{
			Name:  cookie.Name,
			Value: cookie.Value,
			Path:  cookie.Path,
		}

		request.AddCookie(ck)
	}

	HandlerGetCookie(recorder, request)

	result = recorder.Result()
	body, _ := io.ReadAll(result.Body)

	fmt.Println(string(body))
}
