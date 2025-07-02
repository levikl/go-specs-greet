package interactions_test

import (
	"testing"

	"github.com/alecthomas/assert/v2"

	"github.com/levikl/go-specs-greet/domain/interactions"
	"github.com/levikl/go-specs-greet/specifications"
)

func TestGreet(t *testing.T) {
	specifications.GreetSpecification(
		t,
		specifications.GreetAdapter(interactions.Greet),
	)

	t.Run("it defaults `name` to `\"World\"` when it is empty", func(t *testing.T) {
		assert.Equal(t, "Hello, World", interactions.Greet(""))
	})
}
