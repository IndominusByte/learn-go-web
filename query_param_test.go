package learngoweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var Req *http.Request
var Res *http.Response
var Rec *httptest.ResponseRecorder

var SingleParam http.HandlerFunc = func(rw http.ResponseWriter, r *http.Request) {
	if name := r.URL.Query().Get("name"); name != "" {
		fmt.Fprintf(rw, "Hello %s", name)
	} else {
		fmt.Fprint(rw, "Hello Anonymous")
	}
}

var MultipleParam http.HandlerFunc = func(rw http.ResponseWriter, r *http.Request) {
	first_name := r.URL.Query().Get("first_name")
	last_name := r.URL.Query().Get("last_name")

	fmt.Fprintf(rw, "Hello %s %s", first_name, last_name)
}

var MultipleParamValue http.HandlerFunc = func(rw http.ResponseWriter, r *http.Request) {
	names := r.URL.Query()

	fmt.Fprintf(rw, "Helo Bro %s", strings.Join(names["names"], " "))
}

func TestParameters(t *testing.T) {
	Req = httptest.NewRequest(http.MethodGet, "/?name=oman", nil)
	Rec = httptest.NewRecorder()

	SingleParam(Rec, Req)

	Res = Rec.Result()
	body, _ := io.ReadAll(Res.Body)

	fmt.Println(string(body))

	Req = httptest.NewRequest(http.MethodGet, "/", nil)
	Rec = httptest.NewRecorder()

	SingleParam(Rec, Req)

	Res = Rec.Result()
	body, _ = io.ReadAll(Res.Body)

	fmt.Println(string(body))

	Req = httptest.NewRequest(http.MethodGet, "/?first_name=oman&last_name=pradipta", nil)
	Rec = httptest.NewRecorder()

	MultipleParam(Rec, Req)

	Res = Rec.Result()
	body, _ = io.ReadAll(Res.Body)

	fmt.Println(string(body))

	Req = httptest.NewRequest(http.MethodGet, "/?names=1&names=2&names=3", nil)
	Rec = httptest.NewRecorder()

	MultipleParamValue(Rec, Req)

	Res = Rec.Result()
	body, _ = io.ReadAll(Res.Body)

	fmt.Println(string(body))
}
