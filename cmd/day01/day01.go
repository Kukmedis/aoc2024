package main

import (
	"slices"
	"sort"
)

func calculateDistanceBetweenLists(oneList []int, anotherList []int) int {
	oneListCloned := slices.Clone(oneList)
	sort.Ints(oneListCloned)

	anotherListCloned := slices.Clone(anotherList)
	sort.Ints(anotherListCloned)

	sum := 0

	for i := 0; i < len(oneListCloned); i++ {
		sum = sum + abs(oneListCloned[i]-anotherListCloned[i])
	}

	return sum
}

func abs(num int) int {
	absNum := num
	if num < 0 {
		absNum = -num
	}
	return absNum
}

func mapToOccurances(list []int) map[int]int {
	occurances := make(map[int]int)
	for i := 0; i < len(list); i++ {
		occurances[list[i]] = occurances[list[i]] + 1
	}
	return occurances
}

func calculateSimilarity(oneList []int, anotherList []int) int {
	sum := 0
	occuranceMap := mapToOccurances(anotherList)
	for i := 0; i < len(oneList); i++ {
		sum = sum + (oneList[i] * occuranceMap[oneList[i]])
	}
	return sum
}
