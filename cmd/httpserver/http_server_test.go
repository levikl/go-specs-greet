package main_test

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/alecthomas/assert/v2"

	"github.com/levikl/go-specs-greet/adapters"
	"github.com/levikl/go-specs-greet/adapters/httpserver"
	"github.com/levikl/go-specs-greet/specifications"
)

var port = "8080"

func TestGreeterServer(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	driver := httpserver.Driver{
		BaseURL: fmt.Sprintf("http://localhost:%s", port),
		Client: &http.Client{
			Timeout: 1 * time.Second,
		},
	}

	_, err := adapters.StartDockerServer(t, port, "httpserver")
	assert.NoError(t, err)

	specifications.GreetSpecification(t, &driver)
	specifications.CurseSpecification(t, &driver)
}
