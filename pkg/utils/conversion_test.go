package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToInts(t *testing.T) {
	assert.Equal(t, []int{0, 1, 2, 3}, ToInts("0 1 2 3", " "))
}
