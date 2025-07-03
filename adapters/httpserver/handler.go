package httpserver

import (
	"fmt"
	"log"
	"net/http"

	"github.com/levikl/go-specs-greet/domain/interactions"
)

const (
	greetPath = "/greet"
	cursePath = "/curse"
)

func NewHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc(greetPath, replyWith(interactions.Greet))
	mux.HandleFunc(cursePath, replyWith(interactions.Curse))
	return mux
}

func replyWith(f func(name string) (interaction string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if _, err := fmt.Fprint(w, f(name)); err != nil {
			log.Printf("failed to write to ResponseWriter: %v", err)
		}
	}
}
