package main_test

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert/v2"

	"github.com/levikl/go-specs-greet/adapters"
	"github.com/levikl/go-specs-greet/adapters/grpcserver"
	"github.com/levikl/go-specs-greet/specifications"
)

var port = "50051"

func TestGreeterServer(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	driver := grpcserver.Driver{Addr: fmt.Sprintf("localhost:%s", port)}

	t.Cleanup(driver.Close)

	_, err := adapters.StartDockerServer(t, port, "grpcserver")
	assert.NoError(t, err)

	specifications.GreetSpecification(t, &driver)
	specifications.CurseSpecification(t, &driver)
}
