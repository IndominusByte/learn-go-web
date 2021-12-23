package learngoweb

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

var tmpl = template.Must(template.ParseGlob("templates/*.html"))

var HandlerTemplateCache http.HandlerFunc = func(rw http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(rw, "name.html", map[string]interface{}{
		"Title": "If",
		"Name":  "omann",
	})
}

func TestTemplateCache(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/", nil)
	recorder := httptest.NewRecorder()

	HandlerTemplateCache(recorder, request)

	result := recorder.Result()
	body, _ := io.ReadAll(result.Body)

	fmt.Println(string(body))
}
