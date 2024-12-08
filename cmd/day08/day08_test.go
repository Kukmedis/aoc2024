package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateAntinodesPair(t *testing.T) {
	assert.ElementsMatch(t, []location{{3, 1}, {6, 7}}, calculateAntinodesPair(location{4, 3}, location{5, 5}, mapSize{10, 10}))
	assert.ElementsMatch(t, []location{{0, 2}}, calculateAntinodesPair(location{4, 3}, location{8, 4}, mapSize{10, 10}))
	assert.ElementsMatch(t, []location{{2, 6}}, calculateAntinodesPair(location{8, 4}, location{5, 5}, mapSize{10, 10}))
}

func TestCalculateAntinodesPairEmpty(t *testing.T) {
	assert.ElementsMatch(t, []location{}, calculateAntinodesPair(location{1, 1}, location{9, 9}, mapSize{10, 10}))
}

func TestCalculateAntinodes(t *testing.T) {
	assert.ElementsMatch(t,
		[]location{{3, 1}, {6, 7}, {0, 2}, {2, 6}},
		calculateAntinodes([]location{{4, 3}, {5, 5}, {8, 4}}, mapSize{10, 10}, calculateAntinodesPair))
}

func TestCountMapAntinodes(t *testing.T) {
	testMap := map[rune][]location{
		'0': {{8, 1}, {5, 2}, {7, 3}, {4, 4}},
		'A': {{6, 5}, {8, 8}, {9, 9}}}
	assert.Equal(t, 14, countMapAntinodes(testMap, mapSize{12, 12}))
}

func TestCalculateAntinodesPairTFrequency(t *testing.T) {
	assert.ElementsMatch(t, []location{{0, 0}, {3, 1}, {6, 2}, {9, 3}},
		calculateAntinodesPairTFrequency(location{0, 0}, location{3, 1}, mapSize{10, 10}))
	assert.ElementsMatch(t, []location{{0, 0}, {1, 2}, {2, 4}, {3, 6}, {4, 8}},
		calculateAntinodesPairTFrequency(location{0, 0}, location{1, 2}, mapSize{10, 10}))
	assert.ElementsMatch(t, []location{{1, 2}, {3, 1}, {5, 0}},
		calculateAntinodesPairTFrequency(location{1, 2}, location{3, 1}, mapSize{10, 10}))
}

func TestCountMapAntinodesTFrequency(t *testing.T) {
	testMap := map[rune][]location{
		'0': {{8, 1}, {5, 2}, {7, 3}, {4, 4}},
		'A': {{6, 5}, {8, 8}, {9, 9}}}
	assert.Equal(t, 34, countMapAntinodesTFrequency(testMap, mapSize{12, 12}))
}
