package main

import (
	"net/http"

	"github.com/levikl/go-specs-greet/adapters/httpserver"
)

func main() {
	handler := http.HandlerFunc(httpserver.Handler)
	http.ListenAndServe(":8080", handler) // nolint:errcheck
}
