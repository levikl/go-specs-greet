package go_specs_greet_test

import (
	"testing"

	gospecsgreet "github.com/levikl/go-specs-greet/domain/interactions"
	"github.com/levikl/go-specs-greet/specifications"
)

func TestGreet(t *testing.T) {
	specifications.GreetSpecification(
		t,
		specifications.GreetAdapter(gospecsgreet.Greet),
	)
}
