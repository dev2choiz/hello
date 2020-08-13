package helloworld

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSay(t *testing.T) {
	// When
	r := Say()

	// Then
	assert.Equal(t, r, "Hello world !")
}
