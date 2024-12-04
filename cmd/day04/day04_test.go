package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsXmasSimple(t *testing.T) {
	value := [][]rune{{'X'}, {'M'}, {'A'}, {'S'}}
	assert.True(t, isXmas(value, 0, 0, 1, 0))
	assert.False(t, isXmas(value, 0, 0, 1, 1))
	assert.False(t, isXmas(value, 0, 0, 1, -1))
	assert.False(t, isXmas(value, 0, 0, 0, 1))
	assert.False(t, isXmas(value, 0, 0, 0, -1))
	assert.False(t, isXmas(value, 0, 0, -1, 1))
	assert.False(t, isXmas(value, 0, 0, -1, 0))
	assert.False(t, isXmas(value, 0, 0, -1, -1))
}

func TestIsXmasDiagonal(t *testing.T) {
	value := [][]rune{
		{'S', 'n', 'n', 'n'},
		{'n', 'A', 'n', 'n'},
		{'n', 'n', 'M', 'n'},
		{'n', 'n', 'n', 'X'},
	}
	assert.True(t, isXmas(value, 3, 3, -1, -1))
}

func TestCalculateXmasInAnyDirection(t *testing.T) {
	value := [][]rune{
		{'S', 'n', 'n', 'S'},
		{'n', 'A', 'n', 'A'},
		{'n', 'n', 'M', 'M'},
		{'S', 'A', 'M', 'X'},
	}
	assert.Equal(t, 3, calculateXmasInAnyDirection(value, 3, 3))
}

func TestCalculateAllXmasesSimple(t *testing.T) {
	value := [][]rune{
		{'S', 'n', 'n', 'S'},
		{'n', 'A', 'n', 'A'},
		{'n', 'n', 'M', 'M'},
		{'S', 'A', 'M', 'X'},
	}
	assert.Equal(t, 3, calculateAllXmases(value))
}

func TestCalculateAllXmasesExample(t *testing.T) {
	value := [][]rune{
		{'M', 'M', 'M', 'S', 'X', 'X', 'M', 'A', 'S', 'M'},
		{'M', 'S', 'A', 'M', 'X', 'M', 'S', 'M', 'S', 'A'},
		{'A', 'M', 'X', 'S', 'X', 'M', 'A', 'A', 'M', 'M'},
		{'M', 'S', 'A', 'M', 'A', 'S', 'M', 'S', 'M', 'X'},
		{'X', 'M', 'A', 'S', 'A', 'M', 'X', 'A', 'M', 'M'},
		{'X', 'X', 'A', 'M', 'M', 'X', 'X', 'A', 'M', 'A'},
		{'S', 'M', 'S', 'M', 'S', 'A', 'S', 'X', 'S', 'S'},
		{'S', 'A', 'X', 'A', 'M', 'A', 'S', 'A', 'A', 'A'},
		{'M', 'A', 'M', 'M', 'M', 'X', 'M', 'M', 'M', 'M'},
		{'M', 'X', 'M', 'X', 'A', 'X', 'M', 'A', 'S', 'X'},
	}
	assert.Equal(t, 18, calculateAllXmases(value))
}

func TestIsXmaxCross(t *testing.T) {
	value := [][]rune{
		{'M', '.', 'S'},
		{'.', 'A', '.'},
		{'M', '.', 'S'},
	}
	assert.True(t, isXmasCross(value, 1, 1))
}

func TestCalculateXmaxCrossesExample(t *testing.T) {
	value := [][]rune{
		{'M', 'M', 'M', 'S', 'X', 'X', 'M', 'A', 'S', 'M'},
		{'M', 'S', 'A', 'M', 'X', 'M', 'S', 'M', 'S', 'A'},
		{'A', 'M', 'X', 'S', 'X', 'M', 'A', 'A', 'M', 'M'},
		{'M', 'S', 'A', 'M', 'A', 'S', 'M', 'S', 'M', 'X'},
		{'X', 'M', 'A', 'S', 'A', 'M', 'X', 'A', 'M', 'M'},
		{'X', 'X', 'A', 'M', 'M', 'X', 'X', 'A', 'M', 'A'},
		{'S', 'M', 'S', 'M', 'S', 'A', 'S', 'X', 'S', 'S'},
		{'S', 'A', 'X', 'A', 'M', 'A', 'S', 'A', 'A', 'A'},
		{'M', 'A', 'M', 'M', 'M', 'X', 'M', 'M', 'M', 'M'},
		{'M', 'X', 'M', 'X', 'A', 'X', 'M', 'A', 'S', 'X'},
	}
	assert.Equal(t, 9, calculateXmasCrosses(value))
}

func TestIsXmaxCrossTricky(t *testing.T) {
	value := [][]rune{
		{'S', '.', 'M'},
		{'.', 'A', '.'},
		{'M', '.', 'S'},
	}
	assert.False(t, isXmasCross(value, 1, 1))
}
