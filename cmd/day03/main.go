package main

import (
	"fmt"
	"os"
	"time"

	"github.com/Kukmedis/aoc2024/pkg/utils"
)

func main() {
	defer utils.TimeTrack(time.Now(), "Day 4")
	fileContents, err := os.ReadFile("inputs/day03.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(calculateMultipliers(findMultipliers(string(fileContents))))
	fmt.Println(calculateMultipliers(findMultipliers(cleanDonts(string(fileContents)))))
}
