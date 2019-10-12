/**
 * In this simple example, we are creating a simple handler and assigning
 * it as the default handler for all requests to our http server.
 * As a result, all request will be redirected to this handler and
 * thus the same "Hello World" message will get displayed.
 */

package main

import (
	"fmt"
	"net/http"
)

type MyHandler struct{}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func main() {
	handler := MyHandler{}
	server := http.Server{
		Addr: "127.0.0.1:8080",
		Handler: &handler,
	}
	server.ListenAndServe()
}