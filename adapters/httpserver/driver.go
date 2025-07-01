package httpserver

import (
	"io"
	"log"
	"net/http"
)

type Driver struct {
	BaseURL string
	Client  *http.Client
}

func (d Driver) Greet(name string) (string, error) {
	res, err := http.Get(d.BaseURL + "/greet?name=" + name)
	if err != nil {
		return "", err
	}
	defer func() {
		if err := res.Body.Close(); err != nil {
			log.Printf("failed to close response body: %v", err)
		}
	}()

	greeting, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(greeting), nil
}
