package main

import (
	"slices"
)

type order struct {
	lower  int
	higher int
}

func isPageOrderedCorrectly(page []int, orders []order) bool {
	for i := 0; i < len(page)-1; i++ {
		if !isOrdersCorrect(page[i], page[i+1], orders) {
			return false
		}
	}
	return true
}

func isOrderCorrect(a int, b int, singleOrder order) bool {
	return !(singleOrder.higher == a && singleOrder.lower == b)
}

func isOrdersCorrect(a int, b int, orders []order) bool {
	for _, o := range orders {
		if !isOrderCorrect(a, b, o) {
			return false
		}
	}
	return true
}

func sumCorrectPagesMiddle(pages [][]int, orders []order) int {
	sum := 0
	for _, p := range pages {
		if isPageOrderedCorrectly(p, orders) {
			sum = sum + p[len(p)/2]
		}
	}
	return sum
}

func sumIncorrectSortedPagesMiddle(pages [][]int, orders []order) int {
	sum := 0
	for _, p := range pages {
		if !isPageOrderedCorrectly(p, orders) {
			sortedP := sortPage(p, orders)
			sum = sum + sortedP[len(sortedP)/2]
		}
	}
	return sum
}

func sortPage(page []int, orders []order) []int {
	sortedP := slices.Clone(page)
	slices.SortStableFunc(sortedP, func(a int, b int) int {
		if isOrdersCorrect(a, b, orders) {
			return -1
		} else {
			return 1
		}
	})
	return sortedP
}
