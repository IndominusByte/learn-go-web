package learngoweb

import (
	"embed"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

var HandlerSimpleHtml http.HandlerFunc = func(rw http.ResponseWriter, r *http.Request) {
	templateString := "<html><body>{{.}}</body></html>"
	t := template.Must(template.New("SIMPLE").Parse(templateString))

	t.Execute(rw, "<script>alert('you have been pwned')</script>")
}

func TestSimpleHtml(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/", nil)
	recorder := httptest.NewRecorder()

	HandlerSimpleHtml(recorder, request)

	result := recorder.Result()
	body, _ := io.ReadAll(result.Body)

	fmt.Println(string(body))
}

var HandlerSimpleHTMLFile http.HandlerFunc = func(rw http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("templates/simple.html"))
	t.Execute(rw, "<script>alert('you have been pwned')</script>")
}

func TestSimpleHTMLFile(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/", nil)
	recorder := httptest.NewRecorder()

	HandlerSimpleHTMLFile(recorder, request)

	result := recorder.Result()
	body, _ := io.ReadAll(result.Body)

	fmt.Println(string(body))
}

var HandlerTemplateDirectory http.HandlerFunc = func(rw http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseGlob("templates/*.html"))
	t.ExecuteTemplate(rw, "hard.html", "<script>alert('you have been pwned')</script>")
}

func TestTemplateDir(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/", nil)
	recorder := httptest.NewRecorder()

	HandlerTemplateDirectory(recorder, request)

	result := recorder.Result()
	body, _ := io.ReadAll(result.Body)

	fmt.Println(string(body))
}

//go:embed templates/*.html
var templates embed.FS

var HandlerTemplateEmbed http.HandlerFunc = func(rw http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFS(templates, "templates/*.html"))
	t.ExecuteTemplate(rw, "hard.html", "<script>alert('you have been pwned')</script>")
}

func TestTemplateEmbed(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/", nil)
	recorder := httptest.NewRecorder()

	HandlerTemplateEmbed(recorder, request)

	result := recorder.Result()
	body, _ := io.ReadAll(result.Body)

	fmt.Println(string(body))
}
