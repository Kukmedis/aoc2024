package main

import (
	"math"
	"strconv"
	"sync"

	"github.com/Kukmedis/aoc2024/pkg/utils"
)

var (
	mu     sync.Mutex
	amount int
)

func howManyDigits(num int) int {
	result := 0
	for num > 0 {
		num = num / 10
		result++
	}
	return result
}

func splitIntoTwo(num int, wheToSplit int) (int, int) {
	splitDigits := int(math.Pow10(wheToSplit))

	return num / (splitDigits), num % splitDigits
}

func splitIntoTwoStr(num int) (int, int) {
	strNum := strconv.Itoa(num)
	divLen := len(strNum) / 2
	return utils.ToInt(strNum[:divLen]), utils.ToInt(strNum[divLen:])
}

func blink(stones []int) []int {
	var resultSlice []int
	for _, stone := range stones {
		if stone == 0 {
			resultSlice = append(resultSlice, 1)
		} else if numOfDigits := howManyDigits(stone); numOfDigits%2 == 0 {
			stone1, stone2 := splitIntoTwo(stone, numOfDigits/2)
			resultSlice = append(resultSlice, stone1)
			resultSlice = append(resultSlice, stone2)
		} else {
			resultSlice = append(resultSlice, stone*2024)
		}
	}
	return resultSlice
}

func blinkTimes(stones []int, times int) []int {
	blinked := stones
	for range times {
		blinked = blink(blinked)
	}
	return blinked
}

func blinkRecur(stone int, times int) int {
	if times == 0 {
		return 1
	}
	if stone == 0 {
		return blinkRecur(1, times-1)
	} else if numOfDigits := howManyDigits(stone); numOfDigits%2 == 0 {
		stone1, stone2 := splitIntoTwo(stone, numOfDigits/2)
		return blinkRecur(stone1, times-1) + blinkRecur(stone2, times-1)
	} else {
		return blinkRecur(stone*2024, times-1)
	}
}

func blinkRecurStones(stones []int, times int) int {
	sum := 0
	for _, stone := range stones {
		sum = sum + blinkRecur(stone, times)
	}
	return sum
}
