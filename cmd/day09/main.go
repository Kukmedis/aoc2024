package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/Kukmedis/aoc2024/pkg/utils"
)

func main() {
	defer utils.TimeTrack(time.Now(), "Day 9")
	file, err := os.Open("inputs/day09.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	text := scanner.Text()
	fmt.Println(process(text))
	fmt.Println(processPart2(text))
}
