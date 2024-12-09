package main

import (
	"slices"
	"strconv"

	"github.com/Kukmedis/aoc2024/pkg/utils"
)

func parseInput(input string) []string {
	var result []string
	id := 0
	for i := 0; i < len(input); i++ {
		times := utils.ToInt(string(input[i]))
		if i%2 == 0 {
			for range times {
				result = append(result, strconv.Itoa(id))
			}
			id++
		} else {
			for range times {
				result = append(result, ".")
			}
		}
	}
	return result
}

func moveRightMost(input []string) ([]string, bool) {
	for i := len(input) - 1; i >= 0; i-- {
		if (input[i]) != "." {
			for j := 0; j < i; j++ {
				if (input[j]) == "." {
					return swapData(input, j, i), true
				}
			}
		}
	}
	return input, false
}

func moveRightMostMany(input []string, idToMove int) []string {
	for i := len(input) - 1; i >= 0; i-- {
		if (input[i]) == strconv.Itoa(idToMove) {
			dataSize := getDataSize(input, i)
			for j := 0; j < i; j++ {
				emptySize := getEmptySize(input, j)
				if (input[j]) == "." && dataSize <= emptySize {
					swapped := slices.Clone(input)
					for k := range dataSize {
						swapped = swapData(swapped, j+k, i-k)
					}
					return swapped
				}
			}
			return input
		} else if input[i] == "." {
			continue
		} else if utils.ToInt(input[i]) <= idToMove {
			break
		}
	}
	return input
}

func swapData(input []string, first int, second int) []string {
	newSlice := slices.Clone(input)
	firstElement := newSlice[first]
	secondElement := newSlice[second]
	newSlice[first] = secondElement
	newSlice[second] = firstElement
	return newSlice
}

func calculateChecksum(input []string) uint64 {
	var sum uint64
	for i := 0; i < len(input); i++ {
		char := string(input[i])
		if char == "." {
			continue
		}
		sum = sum + uint64(i)*uint64(utils.ToInt(string(input[i])))
	}
	return sum
}

func process(input string) uint64 {
	parsedInput := parseInput(input)
	oneIterationMore, moved := moveRightMost(parsedInput)
	for moved {
		oneIterationMore, moved = moveRightMost(oneIterationMore)
	}
	return calculateChecksum(oneIterationMore)
}

func processPart2(input string) uint64 {
	parsedInput := parseInput(input)
	startId := findBiggestId(parsedInput)
	calcInput := parsedInput
	for i := startId; i >= 0; i-- {
		calcInput = moveRightMostMany(calcInput, i)
	}
	return calculateChecksum(calcInput)
}

func getDataSize(input []string, index int) int {
	size := 1
	for i := index - 1; i >= 0; i-- {
		if input[i] == input[index] {
			size++
		} else {
			return size
		}
	}
	return size
}

func getEmptySize(input []string, index int) int {
	size := 1
	for i := index + 1; i < len(input); i++ {
		if input[i] == input[index] {
			size++
		} else {
			return size
		}
	}
	return size
}

func findBiggestId(input []string) int {
	for i := len(input) - 1; i >= 0; i-- {
		if input[i] != "." {
			return utils.ToInt(input[i])
		}
	}
	return 0
}
