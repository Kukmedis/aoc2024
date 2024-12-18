package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/Kukmedis/aoc2024/pkg/utils"
)

func main() {
	defer utils.TimeTrack(time.Now(), "Day 16")
	file, err := os.Open("inputs/day16.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	var input []string
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	fmt.Println(findCheapestPath(parseInput(input)))
	fmt.Println(howManyBestTiles(parseInput(input)))
}
