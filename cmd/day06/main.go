package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/Kukmedis/aoc2024/pkg/utils"
)

func main() {
	defer utils.TimeTrack(time.Now(), "Day 6")
	file, err := os.Open("inputs/day06.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	var puzzleMap [][]bool
	iRow := 0
	var starPos position
	for scanner.Scan() {
		mapRow := []bool{}
		inputRow := scanner.Text()
		for iCol, c := range inputRow {
			if c == '#' {
				mapRow = append(mapRow, true)
			} else if c == '^' {
				starPos = position{iRow, iCol}
				mapRow = append(mapRow, false)
			} else {
				mapRow = append(mapRow, false)
			}
		}
		puzzleMap = append(puzzleMap, mapRow)
		iRow++
	}
	fmt.Println(escape(starPos, puzzleMap))
	fmt.Println(calculateInfiniteMaps(starPos, puzzleMap))
}
