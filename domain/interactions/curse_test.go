package interactions_test

import (
	"testing"

	"github.com/levikl/go-specs-greet/domain/interactions"
	"github.com/levikl/go-specs-greet/specifications"
)

func TestCurse(t *testing.T) {
	specifications.CurseSpecification(
		t,
		specifications.CurseAdapter(interactions.Curse),
	)
}
