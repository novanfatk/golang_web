package golangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

func SimpleHtml(writer http.ResponseWriter, request *http.Request) {
	templateText := `<html><body>{{.}}</body></html>`
	// t, err := template.New("SIMPLE").Parse(templateText)
	// if err != nil {
	// 	panic(err)
	// }

	t := template.Must(template.New("SIMPLE").Parse(templateText))
	t.ExecuteTemplate(writer, "SIMPLE", "Hello Html Template")
}

func TestSimpleHtml(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	SimpleHtml(recorder, request)
	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func SimpleHtmlFile(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/simple.gohtml"))
	t.ExecuteTemplate(writer, "simple.gohtml", "Hello Html Template")
}

func TestSimpleHtmlFile(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	SimpleHtmlFile(recorder, request)
	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func SimpleHtmlDirectory(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseGlob("./templates/*.gohtml"))
	t.ExecuteTemplate(writer, "simple.gohtml", "Hello Html Template")
}

func TestSimpleHtmlDirectory(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	SimpleHtmlDirectory(recorder, request)
	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func SimpleHtmlEmbed(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFS(templates, "templates/*.gohtml"))
	t.ExecuteTemplate(writer, "simple.gohtml", "Hello Html Template")
}

func TestSimpleHtmlEmbed(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	SimpleHtmlEmbed(recorder, request)
	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
