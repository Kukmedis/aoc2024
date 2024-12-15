package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseInput(t *testing.T) {
	testInput := []string{
		"Button A: X+94, Y+34",
		"Button B: X+22, Y+67",
		"Prize: X=8400, Y=5400",
		"",
		"Button A: X+26, Y+66",
		"Button B: X+67, Y+21",
		"Prize: X=12748, Y=12176",
		"",
		"Button A: X+17, Y+86",
		"Button B: X+84, Y+37",
		"Prize: X=7870, Y=6450",
		"",
		"Button A: X+69, Y+23",
		"Button B: X+27, Y+71",
		"Prize: X=18641, Y=10279",
	}
	expectedOutput := []machine{
		{button{94, 34}, button{22, 67}, 8400, 5400},
		{button{26, 66}, button{67, 21}, 12748, 12176},
		{button{17, 86}, button{84, 37}, 7870, 6450},
		{button{69, 23}, button{27, 71}, 18641, 10279},
	}
	assert.Equal(t, expectedOutput, parseInput(testInput))
}

func TestFindPresses(t *testing.T) {
	testMachine := machine{button{94, 34}, button{22, 67}, 8400, 5400}
	foundCombinations := findPresses(testMachine)
	assert.Equal(t, []presses{{80, 40}}, foundCombinations)
}

func TestFindCheapest(t *testing.T) {
	testMachine := machine{button{94, 34}, button{22, 67}, 8400, 5400}
	cheapest := findCheapestPrice(findPresses(testMachine))
	assert.Equal(t, 280, cheapest)
}

func TestFullExample(t *testing.T) {
	testMachines := []machine{
		{button{94, 34}, button{22, 67}, 8400, 5400},
		{button{26, 66}, button{67, 21}, 12748, 12176},
		{button{17, 86}, button{84, 37}, 7870, 6450},
		{button{69, 23}, button{27, 71}, 18641, 10279},
	}
	assert.Equal(t, 480, process(testMachines))
}

func TestParseInputBig(t *testing.T) {
	testInput := []string{
		"Button A: X+94, Y+34",
		"Button B: X+22, Y+67",
		"Prize: X=8400, Y=5400",
		"",
		"Button A: X+26, Y+66",
		"Button B: X+67, Y+21",
		"Prize: X=12748, Y=12176",
		"",
		"Button A: X+17, Y+86",
		"Button B: X+84, Y+37",
		"Prize: X=7870, Y=6450",
		"",
		"Button A: X+69, Y+23",
		"Button B: X+27, Y+71",
		"Prize: X=18641, Y=10279",
	}
	expectedOutput := []machine{
		{button{94, 34}, button{22, 67}, 10000000008400, 10000000005400},
		{button{26, 66}, button{67, 21}, 10000000012748, 10000000012176},
		{button{17, 86}, button{84, 37}, 10000000007870, 10000000006450},
		{button{69, 23}, button{27, 71}, 10000000018641, 10000000010279},
	}
	assert.Equal(t, expectedOutput, parseInputBig(testInput))
}
