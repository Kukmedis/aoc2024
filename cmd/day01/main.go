package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("inputs/day01.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	var oneList []int
	var secondList []int
	for scanner.Scan() {
		line := scanner.Text()
		twoNumbers := strings.Split(line, "   ")
		oneList = append(oneList, toInt(twoNumbers[0]))
		secondList = append(secondList, toInt(twoNumbers[1]))
	}
	fmt.Println(calculateDistanceBetweenLists(oneList, secondList))
	fmt.Println(calculateSimilarity(oneList, secondList))
}

func toInt(value string) int {
	number, err := strconv.Atoi(value)
	if err != nil {
		panic(err)
	}
	return number
}
