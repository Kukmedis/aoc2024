package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/Kukmedis/aoc2024/pkg/utils"
)

func main() {
	defer utils.TimeTrack(time.Now(), "Day 13")
	file, err := os.Open("inputs/day13.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	var input []string
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	fmt.Println(process(parseInput(input)))
	fmt.Println(process(parseInputBig(input)))
}
