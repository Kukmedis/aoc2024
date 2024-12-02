package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateDistanceBetweenListsExample(t *testing.T) {
	a := []int{3, 4, 2, 1, 3, 3}
	b := []int{4, 3, 5, 3, 9, 3}
	assert.Equal(t, 11, calculateDistanceBetweenLists(a, b))
}

func TestCalculateDistanceBetweenListsSimple(t *testing.T) {
	a := []int{1, 1}
	b := []int{1, 2}
	assert.Equal(t, 1, calculateDistanceBetweenLists(a, b))
}

func TestOccurances(t *testing.T) {
	testList := []int{0, 0, 1, 1, 1, 2}
	expectedMap := map[int]int{0: 2, 1: 3, 2: 1}
	assert.True(t, reflect.DeepEqual(mapToOccurances(testList), expectedMap))
}

func TestCalculateSimilaritySimple(t *testing.T) {
	a := []int{1, 2, 3}
	b := []int{1, 1, 2}
	assert.Equal(t, 4, calculateSimilarity(a, b))
}

func TestCalculateSimilarityExample(t *testing.T) {
	a := []int{3, 4, 2, 1, 3, 3}
	b := []int{4, 3, 5, 3, 9, 3}
	assert.Equal(t, 31, calculateSimilarity(a, b))
}
