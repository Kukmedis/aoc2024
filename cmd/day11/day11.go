package main

import (
	"maps"
	"math"
	"strconv"
	"sync"

	"github.com/Kukmedis/aoc2024/pkg/utils"
)

var wg sync.WaitGroup
var amount int
var lock sync.Mutex

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

func blink(stones map[int]int) map[int]int {
	resultMap := make(map[int]int)
	for stone, number := range stones {
		if stone == 0 {
			resultMap[1] = number + resultMap[1]
		} else if numOfDigits := howManyDigits(stone); numOfDigits%2 == 0 {
			stone1, stone2 := splitIntoTwo(stone, numOfDigits/2)
			resultMap[stone1] = number + resultMap[stone1]
			resultMap[stone2] = number + resultMap[stone2]
		} else {
			multiplied := stone * 2024
			resultMap[multiplied] = number + resultMap[multiplied]
		}
	}
	return resultMap
}

func blinkTimes(stones map[int]int, times int) int {
	blinked := maps.Clone(stones)
	for range times {
		blinked = blink(blinked)
	}
	sum := 0
	for _, v := range blinked {
		sum = sum + v
	}
	wg.Wait()
	return amount
}
