package objects

import (
	"gotest.tools/assert"
	"testing"
)

func TestHasInvitation(t *testing.T) {
	// given
	bag := Bag{}

	// when
	invitation := bag.HasInvitation()

	// then
	assert.Assert(t, !invitation)
}
