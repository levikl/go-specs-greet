package main_test

import (
	"fmt"
	"testing"

	"github.com/levikl/go-specs-greet/adapters"
	"github.com/levikl/go-specs-greet/adapters/grpcserver"
	"github.com/levikl/go-specs-greet/specifications"
)

func TestGreeterServer(t *testing.T) {
	var (
		port   = "50051"
		driver = grpcserver.Driver{Addr: fmt.Sprintf("localhost:%s", port)}
	)

	adapters.StartDockerServer(t, port, "grpcserver")
	specifications.GreetSpecification(t, &driver)
}
