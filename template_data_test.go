package learngoweb

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

var HandlerTemplateDataStruct http.HandlerFunc = func(rw http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseGlob("templates/*.html"))
	t.ExecuteTemplate(rw, "name.html", struct {
		Title string
		Name  string
		Other struct {
			Address string
		}
	}{
		Title: "Belajar",
		Name:  "Bodo Amat",
		Other: struct{ Address string }{
			Address: "purigading",
		},
	})
}

func TestTemplateDataStruct(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/", nil)
	recorder := httptest.NewRecorder()

	HandlerTemplateDataStruct(recorder, request)

	result := recorder.Result()
	body, _ := io.ReadAll(result.Body)

	fmt.Println(string(body))
}

var HandlerTemplateDataMap http.HandlerFunc = func(rw http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseGlob("templates/*.html"))
	t.ExecuteTemplate(rw, "name.html", map[string]interface{}{
		"Title": "Belajar",
		"Name":  "Bodo Amat",
		"Other": map[string]string{
			"Address": "Purigading Bro",
		},
	})
}

func TestTemplateDataMap(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/", nil)
	recorder := httptest.NewRecorder()

	HandlerTemplateDataMap(recorder, request)

	result := recorder.Result()
	body, _ := io.ReadAll(result.Body)

	fmt.Println(string(body))
}
