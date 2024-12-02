package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("inputs/day02.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	var reports [][]int
	for scanner.Scan() {
		line := scanner.Text()
		digits := strings.Split(line, " ")
		var report []int
		for _, v := range digits {
			report = append(report, toInt(v))
		}
		reports = append(reports, report)
	}
	fmt.Println(calculateSafeReports(reports))
	fmt.Println(calculateSafeReportsWithDampener(reports))
}

func toInt(value string) int {
	number, err := strconv.Atoi(value)
	if err != nil {
		panic(err)
	}
	return number
}
