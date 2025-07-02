package adapters

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/alecthomas/assert/v2"
	"github.com/docker/go-connections/nat"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

const (
	startupTimeout = 5 * time.Second
	dockerfileName = "Dockerfile"
)

func StartDockerRod(t testing.TB, rodPort string) (string, error) {
	t.Helper()

	ctx := context.Background()
	gcr := testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			Image:        "ghcr.io/go-rod/rod",
			ExposedPorts: []string{fmt.Sprintf("%s:%s", rodPort, rodPort)},
			WaitingFor: wait.ForListeningPort(nat.Port(rodPort)).
				WithStartupTimeout(startupTimeout),
		},
		Started: true,
	}

	container, err := testcontainers.GenericContainer(ctx, gcr)
	testcontainers.CleanupContainer(t, container)
	assert.NoError(t, err)

	return container.ContainerIP(ctx)
}

func StartDockerServer(
	t testing.TB,
	port string,
	binToBuild string,
) (string, error) {
	t.Helper()

	ctx := context.Background()
	gcr := testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			FromDockerfile: newTCDockerfile(binToBuild),
			ExposedPorts:   []string{fmt.Sprintf("%s:%s", port, port)},
			WaitingFor: wait.ForListeningPort(nat.Port(port)).
				WithStartupTimeout(startupTimeout),
		},
		Started: true,
	}
	container, err := testcontainers.GenericContainer(ctx, gcr)
	testcontainers.CleanupContainer(t, container)
	assert.NoError(t, err)

	return container.ContainerIP(ctx)
}

func newTCDockerfile(binToBuild string) testcontainers.FromDockerfile {
	return testcontainers.FromDockerfile{
		Context:    "../../.",
		Dockerfile: dockerfileName,
		BuildArgs: map[string]*string{
			"bin_to_build": &binToBuild,
		},
		PrintBuildLog: true,
	}
}
