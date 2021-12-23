package learngoweb

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

var HandlerTemplateLayout http.HandlerFunc = func(rw http.ResponseWriter, r *http.Request) {
	// t := template.Must(template.ParseFiles(
	// 	"templates/layouts/header.html",
	// 	"templates/layouts/footer.html",
	// 	"templates/layouts/layout.html",
	// 	"templates/layouts/content.html",
	// ))
	t := template.Must(template.ParseGlob("templates/layouts/*.html"))

	t.ExecuteTemplate(rw, "layout", map[string]string{
		"Title": "Layout Bro",
		"Name":  "oman",
	})
}

func TestTemplateLayout(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/", nil)
	recorder := httptest.NewRecorder()

	HandlerTemplateLayout(recorder, request)

	result := recorder.Result()
	body, _ := io.ReadAll(result.Body)

	fmt.Println(string(body))
}
