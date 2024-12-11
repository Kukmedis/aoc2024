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
	m := map[int]int{0: 1, 1: 1, 10: 1, 99: 1, 999: 1}
	assert.Equal(t, map[int]int{1: 2, 2024: 1, 0: 1, 9: 2, 2021976: 1}, blink(m))
}

func TestBlinkExample(t *testing.T) {
	s1 := map[int]int{125: 1, 17: 1}
	s2 := map[int]int{253000: 1, 1: 1, 7: 1}
	s3 := map[int]int{253: 1, 0: 1, 2024: 1, 14168: 1}
	s4 := map[int]int{512072: 1, 1: 1, 20: 1, 24: 1, 28676032: 1}
	s5 := map[int]int{512: 1, 72: 1, 2024: 1, 2: 2, 0: 1, 4: 1, 2867: 1, 6032: 1}
	s6 := map[int]int{1036288: 1, 7: 1, 2: 1, 20: 1, 24: 1, 4048: 2, 1: 1, 8096: 1, 28: 1, 67: 1, 60: 1, 32: 1}
	assert.Equal(t, s2, blink(s1))
	assert.Equal(t, s3, blink(s2))
	assert.Equal(t, s4, blink(s3))
	assert.Equal(t, s5, blink(s4))
	assert.Equal(t, s6, blink(s5))
}

func TestBlinkTimes(t *testing.T) {
	stones := map[int]int{125: 1, 17: 1}
	assert.Equal(t, 22, blinkTimes(stones, 6))
	assert.Equal(t, 55312, blinkTimes(stones, 25))
}
