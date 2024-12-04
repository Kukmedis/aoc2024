package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMultipliers(t *testing.T) {
	result := findMultipliers("xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))")
	assert.Equal(t, 4, len(result))
	assert.Equal(t, mul{2, 4}, result[0])
	assert.Equal(t, mul{5, 5}, result[1])
	assert.Equal(t, mul{11, 8}, result[2])
	assert.Equal(t, mul{8, 5}, result[3])
}

func TestCalculateMultipliers(t *testing.T) {
	value := []mul{mul{2, 4}, mul{5, 5}, mul{11, 8}, mul{8, 5}}
	assert.Equal(t, 161, calculateMultipliers(value))
}

func TestCleanDonts(t *testing.T) {
	value := "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"
	assert.Equal(t, "xmul(2,4)&mul[3,7]!^?mul(8,5))", cleanDonts(value))
}
