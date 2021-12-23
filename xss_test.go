package learngoweb

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

var TemplateAutoEscape http.HandlerFunc = func(rw http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(rw, "xss.html", map[string]interface{}{
		"Title": "Xss",
		"Body":  "<script>alert('Hello World')</script>",
	})
}

func TestTemplateAutoEscape(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/", nil)
	recorder := httptest.NewRecorder()

	TemplateAutoEscape(recorder, request)

	result := recorder.Result()
	body, _ := io.ReadAll(result.Body)

	fmt.Println(string(body))

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: TemplateAutoEscape,
	}

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}

var TemplateAutoEscapeDisable http.HandlerFunc = func(rw http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(rw, "xss.html", map[string]interface{}{
		"Title": "Xss",
		"Body":  template.HTML("<script>alert('Hello World')</script>"),
	})
}

func TestTemplateAutoEscapeDisable(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/", nil)
	recorder := httptest.NewRecorder()

	TemplateAutoEscapeDisable(recorder, request)

	result := recorder.Result()
	body, _ := io.ReadAll(result.Body)

	fmt.Println(string(body))

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: TemplateAutoEscapeDisable,
	}

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
