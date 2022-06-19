package golangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func ResponseCode(writer http.ResponseWriter, request *http.Request) {
	name := request.URL.Query().Get("name")

	if name == "" {
		writer.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(writer, "name is empety")
	} else {
		fmt.Fprintf(writer, "hello %s", name)
	}
}

func TestResponseCodeInvalid(t *testing.T) {
	request := httptest.NewRequest("Get", "http://localhost:8080", nil)
	recoder := httptest.NewRecorder()

	ResponseCode(recoder, request)

	response := recoder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(response.StatusCode)
	fmt.Println(response.Status)
	fmt.Println(string(body))

}

func TestResponseCodeSuccess(t *testing.T) {
	request := httptest.NewRequest("Get", "http://localhost:8080/?name=novan", nil)
	recoder := httptest.NewRecorder()

	ResponseCode(recoder, request)

	response := recoder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(response.StatusCode)
	fmt.Println(response.Status)
	fmt.Println(string(body))

}
