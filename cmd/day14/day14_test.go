package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseInput(t *testing.T) {
	testInput := []string{
		"p=0,4 v=3,-3",
		"p=6,3 v=-1,-3",
		"p=10,3 v=-1,2",
		"p=2,0 v=2,-1",
		"p=0,0 v=1,3",
		"p=3,0 v=-2,-2",
		"p=7,6 v=-1,-3",
		"p=3,0 v=-1,-2",
		"p=9,3 v=2,3",
		"p=7,3 v=-1,2",
		"p=2,4 v=2,-3",
		"p=9,5 v=-3,-3",
	}
	expectedOutput := []robot{
		{0, 4, 3, -3},
		{6, 3, -1, -3},
		{10, 3, -1, 2},
		{2, 0, 2, -1},
		{0, 0, 1, 3},
		{3, 0, -2, -2},
		{7, 6, -1, -3},
		{3, 0, -1, -2},
		{9, 3, 2, 3},
		{7, 3, -1, 2},
		{2, 4, 2, -3},
		{9, 5, -3, -3}}
	assert.Equal(t, expectedOutput, parseInput(testInput))
}

func TestTick(t *testing.T) {
	robots := []robot{
		{9, 9, 2, 2},
		{1, 1, -2, -2}}
	assert.Equal(t, []robot{{1, 1, 2, 2}, {9, 9, -2, -2}}, tick(robots, 10, 10))
}

func TestTick100(t *testing.T) {
	robots := []robot{
		{0, 4, 3, -3},
		{6, 3, -1, -3},
		{10, 3, -1, 2},
		{2, 0, 2, -1},
		{0, 0, 1, 3},
		{3, 0, -2, -2},
		{7, 6, -1, -3},
		{3, 0, -1, -2},
		{9, 3, 2, 3},
		{7, 3, -1, 2},
		{2, 4, 2, -3},
		{9, 5, -3, -3}}
	after100Ticks := tickNumber(100, robots, 11, 7)

	expectedRobots := []robot{
		{3, 5, 3, -3},
		{5, 4, -1, -3},
		{9, 0, -1, 2},
		{4, 5, 2, -1},
		{1, 6, 1, 3},
		{1, 3, -2, -2},
		{6, 0, -1, -3},
		{2, 3, -1, -2},
		{0, 2, 2, 3},
		{6, 0, -1, 2},
		{4, 5, 2, -3},
		{6, 6, -3, -3}}
	assert.Equal(t, expectedRobots, after100Ticks)
}

func TestCalcSafety(t *testing.T) {
	robots := []robot{
		{3, 5, 3, -3},
		{5, 4, -1, -3},
		{9, 0, -1, 2},
		{4, 5, 2, -1},
		{1, 6, 1, 3},
		{1, 3, -2, -2},
		{6, 0, -1, -3},
		{2, 3, -1, -2},
		{0, 2, 2, 3},
		{6, 0, -1, 2},
		{4, 5, 2, -3},
		{6, 6, -3, -3}}
	assert.Equal(t, 12, calcSafety(robots, 11, 7))
}
