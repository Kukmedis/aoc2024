package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/Kukmedis/aoc2024/pkg/utils"
)

func main() {
	defer utils.TimeTrack(time.Now(), "Day 10")
	file, err := os.Open("inputs/day10.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	var input []string
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	parsedMapInput := parseMapInput(input)
	fmt.Println(calculateAllPathsTo9(parsedMapInput))
	fmt.Println(calculateRatingsAllPathsTo9(parsedMapInput))
}
