package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/Kukmedis/aoc2024/pkg/utils"
)

func main() {
	defer utils.TimeTrack(time.Now(), "Day 5")
	file, err := os.Open("inputs/day05.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	var orders []order
	var allPages [][]int
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		pages := strings.Split(line, "|")

		orders = append(orders, order{toInt(pages[0]), toInt(pages[1])})
	}
	for scanner.Scan() {
		line := scanner.Text()
		pagesStrings := strings.Split(line, ",")
		var pages []int
		for _, p := range pagesStrings {
			pages = append(pages, toInt(p))
		}
		allPages = append(allPages, pages)
	}
	fmt.Println(sumCorrectPagesMiddle(allPages, orders))
	fmt.Println(sumIncorrectSortedPagesMiddle(allPages, orders))
}

func toInt(value string) int {
	number, err := strconv.Atoi(value)
	if err != nil {
		panic(err)
	}
	return number
}
