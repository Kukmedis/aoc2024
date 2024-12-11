package main

import (
	"fmt"
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

func blinkRecur(stone int, times int) {
	if times == 0 {
		lock.Lock()
		amount++
		fmt.Println(amount)
		lock.Unlock()
	} else if stone == 0 {
		wg.Add(1)
		go blinkRecur(1, times-1)
	} else if numOfDigits := howManyDigits(stone); numOfDigits%2 == 0 {
		stone1, stone2 := splitIntoTwo(stone, numOfDigits/2)
		wg.Add(1)
		go blinkRecur(stone1, times-1)
		wg.Add(1)
		go blinkRecur(stone2, times-1)
	} else {
		wg.Add(1)
		go blinkRecur(stone*2024, times-1)
	}
	wg.Done()
}

func blinkRecurStones(stones []int, times int) int {
	amount = 0
	for _, stone := range stones {
		wg.Add(1)
		blinkRecur(stone, times)
	}
	wg.Wait()
	return amount
}
