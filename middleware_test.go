package golangweb

import (
	"fmt"
	"net/http"
	"testing"
)

type LogMiddleware struct {
	Handler http.Handler
}

func (middlware *LogMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("berfore execute handler")
	middlware.Handler.ServeHTTP(writer, request)
	fmt.Println("after execute handler")
}

type ErrorHandler struct {
	Handler http.Handler
}

func (errorHandler *ErrorHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("Terjadi errot")
			writer.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(writer, "Error :%s", err)
		}
	}()
	errorHandler.Handler.ServeHTTP(writer, request)
}
func TestMiddleware(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Handler Executed")
		fmt.Fprint(writer, "Hello Middleware")
	})
	mux.HandleFunc("/foo", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("foo Executed")
		fmt.Fprint(writer, "Hello Foo")
	})
	mux.HandleFunc("/panic", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("panic executed")
		panic("Ups")
	})
	logMiddleware := LogMiddleware{
		Handler: mux,
	}
	errorhandler := ErrorHandler{
		Handler: &logMiddleware,
	}
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: &errorhandler,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
