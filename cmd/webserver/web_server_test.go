package main_test

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert/v2"

	"github.com/levikl/go-specs-greet/adapters"
	"github.com/levikl/go-specs-greet/adapters/webserver"
	"github.com/levikl/go-specs-greet/specifications"
)

var (
	rodPort       = "7317"
	webserverPort = "8081"
)

func TestGreeterWeb(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	_, err := adapters.StartDockerRod(t, rodPort)
	assert.NoError(t, err)

	webserverIP, err := adapters.StartDockerServer(t, webserverPort, "webserver")
	assert.NoError(t, err)

	driver, cleanup := webserver.NewDriver(fmt.Sprintf("http://%s:%s", webserverIP, webserverPort))

	t.Cleanup(func() {
		assert.NoError(t, cleanup())
	})

	specifications.GreetSpecification(t, driver)
	specifications.CurseSpecification(t, driver)
}
