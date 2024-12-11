package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHowManyDigits(t *testing.T) {
	assert.Equal(t, 1, howManyDigits(9))
	assert.Equal(t, 6, howManyDigits(987654))
}

func TestSplitDigits(t *testing.T) {
	n1, n2 := splitIntoTwo(1001, 2)
	assert.Equal(t, 10, n1)
	assert.Equal(t, 1, n2)
	m1, m2 := splitIntoTwo(1000, 2)
	assert.Equal(t, 10, m1)
	assert.Equal(t, 0, m2)
	l1, l2 := splitIntoTwo(909909, 3)
	assert.Equal(t, 909, l1)
	assert.Equal(t, 909, l2)
	k1, k2 := splitIntoTwo(99, 1)
	assert.Equal(t, 9, k1)
	assert.Equal(t, 9, k2)
}

func TestSplitDigitsStr(t *testing.T) {
	n1, n2 := splitIntoTwoStr(1001)
	assert.Equal(t, 10, n1)
	assert.Equal(t, 1, n2)
	m1, m2 := splitIntoTwoStr(1000)
	assert.Equal(t, 10, m1)
	assert.Equal(t, 0, m2)
	l1, l2 := splitIntoTwoStr(909909)
	assert.Equal(t, 909, l1)
	assert.Equal(t, 909, l2)
	k1, k2 := splitIntoTwoStr(99)
	assert.Equal(t, 9, k1)
	assert.Equal(t, 9, k2)
}

func TestBlink(t *testing.T) {
	assert.Equal(t, []int{1, 2024, 1, 0, 9, 9, 2021976}, blink([]int{0, 1, 10, 99, 999}))
}

func TestBlinkExample(t *testing.T) {
	s1 := []int{125, 17}
	s2 := []int{253000, 1, 7}
	s3 := []int{253, 0, 2024, 14168}
	s4 := []int{512072, 1, 20, 24, 28676032}
	s5 := []int{512, 72, 2024, 2, 0, 2, 4, 2867, 6032}
	s6 := []int{1036288, 7, 2, 20, 24, 4048, 1, 4048, 8096, 28, 67, 60, 32}
	s7 := []int{2097446912, 14168, 4048, 2, 0, 2, 4, 40, 48, 2024, 40, 48, 80, 96, 2, 8, 6, 7, 6, 0, 3, 2}
	assert.Equal(t, s2, blink(s1))
	assert.Equal(t, s3, blink(s2))
	assert.Equal(t, s4, blink(s3))
	assert.Equal(t, s5, blink(s4))
	assert.Equal(t, s6, blink(s5))
	assert.Equal(t, s7, blink(s6))
}

func TestBlinkTimes(t *testing.T) {
	stones := []int{125, 17}
	assert.Equal(t, 22, len(blinkTimes(stones, 6)))
	assert.Equal(t, 55312, len(blinkTimes(stones, 25)))
}

func TestBlinkTimesRecur(t *testing.T) {
	stones := []int{125, 17}
	assert.Equal(t, 22, blinkRecurStones(stones, 6))
	assert.Equal(t, 55312, blinkRecurStones(stones, 25))
}
