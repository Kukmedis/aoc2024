package main

import (
	"strings"
	"unicode"

	"github.com/Kukmedis/aoc2024/pkg/utils"
)

type button struct {
	x int
	y int
}

type machine struct {
	aButton button
	bButton button
	xPrice  int
	yPrice  int
}

type presses struct {
	aButton int
	bButton int
}

func parseInput(input []string) []machine {
	var machines []machine
	for i := 0; i < len(input); i = i + 4 {
		var machine machine
		machine.aButton.x = extractIntFrom(input[i], "X+")
		machine.aButton.y = extractIntFrom(input[i], "Y+")
		machine.bButton.x = extractIntFrom(input[i+1], "X+")
		machine.bButton.y = extractIntFrom(input[i+1], "Y+")
		machine.xPrice = extractIntFrom(input[i+2], "X=")
		machine.yPrice = extractIntFrom(input[i+2], "Y=")
		machines = append(machines, machine)
	}
	return machines
}

func parseInputBig(input []string) []machine {
	var machines []machine
	for i := 0; i < len(input); i = i + 4 {
		var machine machine
		machine.aButton.x = extractIntFrom(input[i], "X+")
		machine.aButton.y = extractIntFrom(input[i], "Y+")
		machine.bButton.x = extractIntFrom(input[i+1], "X+")
		machine.bButton.y = extractIntFrom(input[i+1], "Y+")
		machine.xPrice = extractIntFrom(input[i+2], "X=") + 10000000000000
		machine.yPrice = extractIntFrom(input[i+2], "Y=") + 10000000000000
		machines = append(machines, machine)
	}
	return machines
}

func extractIntFrom(input string, marker string) int {
	startIndex := strings.Index(input, marker) + len(marker)
	digitInput := input[startIndex:]
	for i := 0; i < len(digitInput); i++ {
		if !unicode.IsDigit(rune(digitInput[i])) {
			return utils.ToInt(digitInput[0:i])
		}
	}
	return utils.ToInt(digitInput)
}

func findPresses(machine machine) []presses {
	var foundPresses []presses
	mathPresses := calcPresses(machine)
	if (mathPresses.aButton*machine.aButton.x+mathPresses.bButton*machine.bButton.x == machine.xPrice) &&
		(mathPresses.aButton*machine.aButton.y+mathPresses.bButton*machine.bButton.y == machine.yPrice) {
		foundPresses = append(foundPresses, mathPresses)
	}
	return foundPresses
}

func calcPresses(machine machine) presses {
	b := (machine.yPrice*machine.aButton.x - machine.xPrice*machine.aButton.y) /
		(machine.bButton.y*machine.aButton.x - machine.bButton.x*machine.aButton.y)
	a := (machine.xPrice - b*machine.bButton.x) / machine.aButton.x
	return presses{a, b}
}

func findCheapestPrice(presses []presses) int {
	price := 0
	for _, p := range presses {
		if price == 0 {
			price = p.aButton*3 + p.bButton
		} else {
			price = min(price, p.aButton*3+p.bButton)
		}
	}
	return price
}

func process(machines []machine) int {
	totalPrice := 0
	for _, m := range machines {
		totalPrice = totalPrice + findCheapestPrice(findPresses(m))
	}
	return totalPrice
}
