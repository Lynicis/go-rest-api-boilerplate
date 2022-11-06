//go:build unit

package pact

import (
	"testing"

	"github.com/pact-foundation/pact-go/types"
	"github.com/stretchr/testify/assert"
)

func TestPact_CreateProvider(t *testing.T) {
	actualProvider := CreateProvider(
		"localhost.com",
		"TestProvider",
		"TestPactToken",
		nil,
	)

	expectedProvider := &types.VerifyRequest{}

	assert.IsType(t, expectedProvider, actualProvider)
}
