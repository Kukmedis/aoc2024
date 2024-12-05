package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/Kukmedis/aoc2024/pkg/utils"
)

func main() {
	defer utils.TimeTrack(time.Now(), "Day 4")
	file, err := os.Open("inputs/day04.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	var letters [][]rune
	for scanner.Scan() {
		letterLine := []rune{}
		line := scanner.Text()
		for _, c := range line {
			letterLine = append(letterLine, c)
		}
		letters = append(letters, letterLine)
	}
	fmt.Println(calculateAllXmases(letters))
	fmt.Println(calculateXmasCrosses(letters))
}
