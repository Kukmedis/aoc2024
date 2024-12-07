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
	defer utils.TimeTrack(time.Now(), "Day 7")
	file, err := os.Open("inputs/day07.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	var tests []testFixture
	for scanner.Scan() {
		row := strings.Split(scanner.Text(), ": ")
		tests = append(tests, testFixture{utils.ToInts(row[1], " "), utils.ToInt(row[0])})
	}
	fmt.Println(calculateCalibrationResult(tests))
	fmt.Println(calculateExtendedCalibrationResult(tests))
}
