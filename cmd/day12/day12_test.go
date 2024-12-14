package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalcPriceSimple(t *testing.T) {
	testInput := []string{
		"AAAA",
		"BBCD",
		"BBCC",
		"EEEC",
	}
	assert.Equal(t, 140, process(testInput, calcPrice))
}

func TestCalcPriceXO(t *testing.T) {
	testInput := []string{
		"OOOOO",
		"OXOXO",
		"OOOOO",
		"OXOXO",
		"OOOOO",
	}
	assert.Equal(t, 772, process(testInput, calcPrice))
}

func TestCalcPriceBig(t *testing.T) {
	testInput := []string{
		"RRRRIICCFF",
		"RRRRIICCCF",
		"VVRRRCCFFF",
		"VVRCCCJFFF",
		"VVVVCJJCFE",
		"VVIVCCJJEE",
		"VVIIICJJEE",
		"MIIIIIJJEE",
		"MIIISIJEEE",
		"MMMISSJEEE",
	}
	assert.Equal(t, 1930, process(testInput, calcPrice))
}

func TestCalcPriceSidesSimple(t *testing.T) {
	testInput := []string{
		"AAAA",
		"BBCD",
		"BBCC",
		"EEEC",
	}
	assert.Equal(t, 80, process(testInput, calcPriceWithSides))
}

func TestCalcPriceSidesE(t *testing.T) {
	testInput := []string{
		"EEEEE",
		"EXXXX",
		"EEEEE",
		"EXXXX",
		"EEEEE",
	}
	assert.Equal(t, 236, process(testInput, calcPriceWithSides))
}

func TestCalcPriceSidesLastExample(t *testing.T) {
	testInput := []string{
		"AAAAAA",
		"AAABBA",
		"AAABBA",
		"ABBAAA",
		"ABBAAA",
		"AAAAAA",
	}
	assert.Equal(t, 368, process(testInput, calcPriceWithSides))
}

func TestCalcPriceSidesBig(t *testing.T) {
	testInput := []string{
		"RRRRIICCFF",
		"RRRRIICCCF",
		"VVRRRCCFFF",
		"VVRCCCJFFF",
		"VVVVCJJCFE",
		"VVIVCCJJEE",
		"VVIIICJJEE",
		"MIIIIIJJEE",
		"MIIISIJEEE",
		"MMMISSJEEE",
	}
	assert.Equal(t, 1206, process(testInput, calcPriceWithSides))
}

func TestCalcPriceSidesDebugE(t *testing.T) {
	testInput := []string{
		"  E",
		" EE",
		" EE",
		" EE",
		"EEE",
		"EEE",
	}

	assert.Equal(t, 104, process(testInput, calcPriceWithSides))
}

func TestCalcNumberOfSidesPerSide(t *testing.T) {
	testInput := []cCell{
		{0, 5, true, false, true, false},
		{1, 5, true, true, true, false},
		{2, 5, true, true, false, false},
	}
	assert.Equal(t, 1, calcNumberOfSidesPerSide(testInput))
}

func TestCalcSides(t *testing.T) {
	assert.Equal(t, 4, calcSides([]cell{{0, 0, "A"}, {1, 0, "A"}, {2, 0, "A"}, {3, 0, "A"}}))
}

func TestCalcSidesS(t *testing.T) {
	assert.Equal(t, 8, calcSides([]cell{{0, 0, "A"}, {1, 0, "A"}, {1, 1, "A"}, {2, 1, "A"}}))
}
