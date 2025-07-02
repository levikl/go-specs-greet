package main

import (
	"log"
	"net/http"

	"github.com/levikl/go-specs-greet/adapters/webserver"
)

func main() {
	log.Fatal(http.ListenAndServe(":8081", webserver.NewHandler()))
}
