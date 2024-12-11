package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/Kukmedis/aoc2024/pkg/utils"
)

func main() {
	defer utils.TimeTrack(time.Now(), "Day 11")
	file, err := os.Open("inputs/day11.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	text := scanner.Text()
	var stones []int
	for _, v := range strings.Split(text, " ") {
		stones = append(stones, utils.ToInt(v))
	}
	fmt.Println(len(blinkTimes(stones, 25)))
	fmt.Println(blinkRecurStones(stones, 75))
}
