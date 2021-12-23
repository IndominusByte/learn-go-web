package learngoweb

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

var HandlerTemplateActionIf http.HandlerFunc = func(rw http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("templates/action/if.html"))
	t.ExecuteTemplate(rw, "if.html", map[string]interface{}{
		"Title": "If",
		"Name":  "Nyoman",
	})
}

func TestTemplateIf(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/", nil)
	recorder := httptest.NewRecorder()

	HandlerTemplateActionIf(recorder, request)

	result := recorder.Result()
	body, _ := io.ReadAll(result.Body)

	fmt.Println(string(body))
}

var HandlerTemplateActionCompare http.HandlerFunc = func(rw http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("templates/action/compare.html"))
	t.ExecuteTemplate(rw, "compare.html", map[string]interface{}{
		"Title": "Compare",
		"Value": 60,
	})
}

func TestTemplateCompare(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/", nil)
	recorder := httptest.NewRecorder()

	HandlerTemplateActionCompare(recorder, request)

	result := recorder.Result()
	body, _ := io.ReadAll(result.Body)

	fmt.Println(string(body))
}

var HandlerTemplateActionRange http.HandlerFunc = func(rw http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("templates/action/range.html"))
	t.ExecuteTemplate(rw, "range.html", map[string]interface{}{
		"Title": "Range",
		"Persons": [...]map[string]interface{}{
			{
				"Name":    "Oman",
				"Hobbies": [...]string{"membaca", "mewarnai"},
			},
			{
				"Name":    "Pradipta",
				"Hobbies": [...]string{"coding", "hacking"},
			},
		},
	})
}

func TestTemplateRange(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/", nil)
	recorder := httptest.NewRecorder()

	HandlerTemplateActionRange(recorder, request)

	result := recorder.Result()
	body, _ := io.ReadAll(result.Body)

	fmt.Println(string(body))
}

var HandlerTemplateActionWith http.HandlerFunc = func(rw http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("templates/action/with.html"))
	t.ExecuteTemplate(rw, "with.html", map[string]interface{}{
		"Title": "With",
		"Address": map[string]string{
			"State": "Indonesia",
			"City":  "Bali",
		},
	})
}

func TestTemplateWith(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/", nil)
	recorder := httptest.NewRecorder()

	HandlerTemplateActionWith(recorder, request)

	result := recorder.Result()
	body, _ := io.ReadAll(result.Body)

	fmt.Println(string(body))
}
