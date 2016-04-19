package example_test

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSomething(t *testing.T) {

	var a = "Hello"
	var b = "Hello"

	require.Equal(t, a, b, "The two words should be the same.")

}
