package learngoweb

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type Page struct {
	Name string
}

func (page *Page) SayHello(name string) string {
	return "Hi " + name + " My name is " + page.Name
}

var HandlerTemplateFunctionStruct http.HandlerFunc = func(rw http.ResponseWriter, r *http.Request) {
	t := template.Must(template.New("FUNCTION").Parse(`{{.SayHello "Budi"}}`))
	t.ExecuteTemplate(rw, "FUNCTION", &Page{
		Name: "Oman",
	})
}

func TestTemplateFunctionStruct(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/", nil)
	recorder := httptest.NewRecorder()

	HandlerTemplateFunctionStruct(recorder, request)

	result := recorder.Result()
	body, _ := io.ReadAll(result.Body)

	fmt.Println(string(body))
}

var HandlerTemplateFunctionGlobal http.HandlerFunc = func(rw http.ResponseWriter, r *http.Request) {
	t := template.Must(template.New("FUNCTION").Parse(`{{len .Name}}`))
	t.Execute(rw, map[string]string{
		"Name": "oman",
	})
}

func TestTemplateFunctionGlobal(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/", nil)
	recorder := httptest.NewRecorder()

	HandlerTemplateFunctionGlobal(recorder, request)

	result := recorder.Result()
	body, _ := io.ReadAll(result.Body)

	fmt.Println(string(body))
}

var HandlerTemplateFunctionGlobalCreate http.HandlerFunc = func(rw http.ResponseWriter, r *http.Request) {
	t := template.New("FUNCTION")
	t.Funcs(template.FuncMap{
		"upper": func(value string) string {
			return strings.ToUpper(value)
		},
	})
	t = template.Must(t.Parse(`{{upper .Name}}`))
	t.Execute(rw, map[string]string{
		"Name": "pradipta",
	})
}

func TestTemplateFunctionGlobalCreate(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/", nil)
	recorder := httptest.NewRecorder()

	HandlerTemplateFunctionGlobalCreate(recorder, request)

	result := recorder.Result()
	body, _ := io.ReadAll(result.Body)

	fmt.Println(string(body))
}

var HandlerTemplateFunctionPipeline http.HandlerFunc = func(rw http.ResponseWriter, r *http.Request) {
	t := template.New("FUNCTION")
	t.Funcs(template.FuncMap{
		"sayhello": func(name string) string {
			return "Hello " + name
		},
		"upper": func(value string) string {
			return strings.ToUpper(value)
		},
	})
	t = template.Must(t.Parse(`{{sayhello .Name | upper | printf "%s"}}`))
	t.Execute(rw, map[string]string{
		"Name": "pradipta",
	})
}

func TestTemplateFunctionPipeline(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/", nil)
	recorder := httptest.NewRecorder()

	HandlerTemplateFunctionPipeline(recorder, request)

	result := recorder.Result()
	body, _ := io.ReadAll(result.Body)

	fmt.Println(string(body))
}
