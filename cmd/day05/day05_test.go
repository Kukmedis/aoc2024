package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func getExampleOrder() []order {
	return []order{
		{47, 53},
		{97, 13},
		{97, 61},
		{97, 47},
		{75, 29},
		{61, 13},
		{75, 53},
		{29, 13},
		{97, 29},
		{53, 29},
		{61, 53},
		{97, 53},
		{61, 29},
		{47, 13},
		{75, 47},
		{97, 75},
		{47, 61},
		{75, 61},
		{47, 29},
		{75, 13},
		{53, 13},
	}
}

func TestIsOrderCorrect(t *testing.T) {

	assert.True(t, isOrderCorrect(1, 5, order{1, 5}))
	assert.True(t, isOrderCorrect(2, 3, order{1, 5}))
	assert.True(t, isOrderCorrect(1, 6, order{1, 5}))
	assert.True(t, isOrderCorrect(1, 0, order{1, 5}))
	assert.False(t, isOrderCorrect(1, 5, order{5, 1}))

}

func TestIsPageOrderedCorrectlyExample(t *testing.T) {
	assert.True(t, isPageOrderedCorrectly([]int{75, 47, 61, 53, 29}, getExampleOrder()))
	assert.True(t, isPageOrderedCorrectly([]int{97, 61, 53, 29, 13}, getExampleOrder()))
	assert.True(t, isPageOrderedCorrectly([]int{75, 29, 13}, getExampleOrder()))
	assert.False(t, isPageOrderedCorrectly([]int{75, 97, 47, 61, 53}, getExampleOrder()))
	assert.False(t, isPageOrderedCorrectly([]int{61, 13, 29}, getExampleOrder()))
	assert.False(t, isPageOrderedCorrectly([]int{97, 13, 75, 29, 47}, getExampleOrder()))
}

func TestSumCorrectPagesMiddle(t *testing.T) {
	testPages := [][]int{
		{75, 47, 61, 53, 29},
		{97, 61, 53, 29, 13},
		{75, 29, 13},
		{75, 97, 47, 61, 53},
		{61, 13, 29},
		{97, 13, 75, 29, 47}}
	assert.Equal(t, 143, sumCorrectPagesMiddle(testPages, getExampleOrder()))
}

func TestSortPage(t *testing.T) {
	assert.Equal(t, []int{97, 75, 47, 61, 53}, sortPage([]int{75, 97, 47, 61, 53}, getExampleOrder()))
	assert.Equal(t, []int{61, 29, 13}, sortPage([]int{61, 13, 29}, getExampleOrder()))
	assert.Equal(t, []int{97, 75, 47, 29, 13}, sortPage([]int{97, 13, 75, 29, 47}, getExampleOrder()))
}

func TestSumIncorrectSortedPagesMiddle(t *testing.T) {
	testPages := [][]int{
		{75, 47, 61, 53, 29},
		{97, 61, 53, 29, 13},
		{75, 29, 13},
		{75, 97, 47, 61, 53},
		{61, 13, 29},
		{97, 13, 75, 29, 47}}
	assert.Equal(t, 123, sumIncorrectSortedPagesMiddle(testPages, getExampleOrder()))
}
