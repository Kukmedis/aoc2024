package main

import (
	"github.com/Kukmedis/aoc2024/pkg/utils"
)

type location struct {
	x      int
	y      int
	height int
}

func parseMapInput(input []string) map[int][]location {
	result := make(map[int][]location)
	for y, row := range input {
		for x, char := range row {
			height := 0
			if char == '.' {
				height = -9
			} else {
				height = utils.ToInt(string(char))
			}

			result[height] = append(result[height], location{x, y, height})
		}
	}
	return result
}

func transformToAdjancencyList(locations map[int][]location) map[location][]location {
	result := make(map[location][]location)
	for key, value := range locations {
		for _, location := range value {
			for _, nextLocation := range locations[key+1] {
				if adjacent(location, nextLocation) {
					result[location] = append(result[location], nextLocation)
				}
			}
		}
	}
	return result
}

func isTherePathFromTo(start, finish location, adjancencyMap map[location][]location) bool {
	if start == finish {
		return true
	} else if len(adjancencyMap[start]) == 0 {
		return false
	} else {
		for _, nextLocation := range adjancencyMap[start] {
			if isTherePathFromTo(nextLocation, finish, adjancencyMap) {
				return true
			}
		}
	}
	return false
}

func calculateAllPathsTo9(input map[int][]location) int {
	sum := 0
	adjacencyMap := transformToAdjancencyList(input)
	for _, start := range input[0] {
		for _, finish := range input[9] {
			if isTherePathFromTo(start, finish, adjacencyMap) {
				sum++
			}
		}
	}
	return sum
}

func adjacent(l1, l2 location) bool {
	xDistance := utils.Abs(l1.x - l2.x)
	yDistance := utils.Abs(l1.y - l2.y)
	return (xDistance == 1 && yDistance == 0) || (xDistance == 0 && yDistance == 1)
}

func howManyPathsFromTo(start, finish location, adjancencyMap map[location][]location) int {
	sum := 0
	if start == finish {
		return 1
	} else if len(adjancencyMap[start]) == 0 {
		return 0
	} else {
		for _, nextLocation := range adjancencyMap[start] {
			sum = sum + howManyPathsFromTo(nextLocation, finish, adjancencyMap)
		}
	}
	return sum
}

func calculateRatingsAllPathsTo9(input map[int][]location) int {
	sum := 0
	adjacencyMap := transformToAdjancencyList(input)
	for _, start := range input[0] {
		for _, finish := range input[9] {
			sum = sum + howManyPathsFromTo(start, finish, adjacencyMap)
		}
	}
	return sum
}
