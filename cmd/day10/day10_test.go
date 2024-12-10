package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalcSimplePaths(t *testing.T) {
	testInput := []string{
		"0123",
		"1234",
		"8765",
		"9876",
	}
	assert.Equal(t, 1, calculateAllPathsTo9(parseMapInput(testInput)))
}

func TestCalcSimpleExamplePaths(t *testing.T) {
	testInput := []string{
		"...0...",
		"...1...",
		"...2...",
		"6543456",
		"7.....7",
		"8.....8",
		"9.....9",
	}
	assert.Equal(t, 2, calculateAllPathsTo9(parseMapInput(testInput)))
}

func TestCalcSimpleExamplePaths2(t *testing.T) {
	testInput := []string{
		"..90..9",
		"...1.98",
		"...2..7",
		"6543456",
		"765.987",
		"876....",
		"987....",
	}
	assert.Equal(t, 4, calculateAllPathsTo9(parseMapInput(testInput)))
}

func TestCalcAllPathsTo9(t *testing.T) {
	testInput := []string{
		"89010123",
		"78121874",
		"87430965",
		"96549874",
		"45678903",
		"32019012",
		"01329801",
		"10456732",
	}
	assert.Equal(t, 36, calculateAllPathsTo9(parseMapInput(testInput)))
}

func TestCalcRatingsAllPathsTo9Example1(t *testing.T) {
	testInput := []string{
		".....0.",
		"..4321.",
		"..5..2.",
		"..6543.",
		"..7..4.",
		"..8765.",
		"..9....",
	}
	assert.Equal(t, 3, calculateRatingsAllPathsTo9(parseMapInput(testInput)))
}

func TestCalcRatingsAllPathsTo9Example2(t *testing.T) {
	testInput := []string{
		"..90..9",
		"...1.98",
		"...2..7",
		"6543456",
		"765.987",
		"876....",
		"987....",
	}
	assert.Equal(t, 13, calculateRatingsAllPathsTo9(parseMapInput(testInput)))
}

func TestCalcRatingsAllPathsTo9ExampleBiggest(t *testing.T) {
	testInput := []string{
		"89010123",
		"78121874",
		"87430965",
		"96549874",
		"45678903",
		"32019012",
		"01329801",
		"10456732",
	}
	assert.Equal(t, 81, calculateRatingsAllPathsTo9(parseMapInput(testInput)))
}
