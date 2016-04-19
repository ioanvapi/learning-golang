package example1_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSomething(t *testing.T) {

	assert := assert.New(t)
	var a = "Hello"
	var b = "Hello"

	assert.Equal(a, b, "The two words should be the same.")

}
