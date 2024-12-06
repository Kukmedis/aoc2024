package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEscape(t *testing.T) {
	startPosition := position{6, 4}
	assert.Equal(t, 41, escape(startPosition, getTestMap()))
}

func TestInfiniteNo(t *testing.T) {
	startPosition := position{6, 4}
	assert.True(t, isInfinite(startPosition, getInfiniteMap()))
}

func TestFiniteTricky(t *testing.T) {
	startPosition := position{6, 4}
	assert.False(t, isInfinite(startPosition, getFiniteMapTricky()))
}

func TestCalculateInfiniteMaps(t *testing.T) {
	startPosition := position{6, 4}
	assert.Equal(t, 6, calculateInfiniteMaps(startPosition, getTestMap()))
}

func getTestMap() [][]bool {
	return [][]bool{
		{false, false, false, false, true, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, true},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, true, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, true, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, true, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, true, false},
		{true, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, true, false, false, false},
	}
}

func getInfiniteMap() [][]bool {
	return [][]bool{
		{false, false, false, false, true, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, true},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, true, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, true, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, true, false, true, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, true, false},
		{true, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, true, false, false, false},
	}
}

func getFiniteMapTricky() [][]bool {
	return [][]bool{
		{false, false, false, false, true, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, true},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, true, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, true, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, true, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, true, false},
		{true, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, true, false, false, true},
	}
}
