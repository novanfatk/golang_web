package golangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func SayHello(writer http.ResponseWriter, request *http.Request) {
	name := request.URL.Query().Get("name")
	if name == "" {
		fmt.Fprint(writer, "hello")
	} else {
		fmt.Fprintf(writer, "hello %s", name)
	}
}

func TestQueryParameter(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?name=Novan", nil)
	recorder := httptest.NewRecorder()

	SayHello(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))

}

func MutipleQueryParameter(writer http.ResponseWriter, request *http.Request) {
	firstName := request.URL.Query().Get("first_name")
	lastName := request.URL.Query().Get("last_name")

	fmt.Fprintf(writer, "hello %s %s", firstName, lastName)

}

func TestMutipleQueryParameter(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?first_name=Novan&last_name=fatkhurrohman", nil)
	recorder := httptest.NewRecorder()

	MutipleQueryParameter(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}

func MultipleQueryParameterValue(writer http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	names := query["name"]

	fmt.Fprint(writer, strings.Join(names, " "))

}

func TestMultipleQueryParameterValue(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?name=Novan&name=fatkhurrohman&name=golang", nil)
	recorder := httptest.NewRecorder()

	MultipleQueryParameterValue(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}
