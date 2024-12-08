package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/Kukmedis/aoc2024/pkg/utils"
)

func main() {
	defer utils.TimeTrack(time.Now(), "Day 8")
	file, err := os.Open("inputs/day08.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	iRow := 0
	var maxColumn int
	mapOfAntennas := make(map[rune][]location)
	for scanner.Scan() {
		for iColumn, v := range scanner.Text() {
			if v != '.' {
				mapOfAntennas[v] = append(mapOfAntennas[v], location{iColumn, iRow})
			}
			maxColumn = iColumn
		}
		iRow++
	}
	puzzleSize := mapSize{maxColumn + 1, iRow}
	fmt.Println(countMapAntinodes(mapOfAntennas, puzzleSize))
	fmt.Println(countMapAntinodesTFrequency(mapOfAntennas, puzzleSize))
}
