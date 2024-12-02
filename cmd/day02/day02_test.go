package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsReportSafe(t *testing.T) {
	assert.Equal(t, true, isReportSafe([]int{7, 6, 4, 2, 1}))
	assert.Equal(t, false, isReportSafe([]int{1, 2, 7, 8, 9}))
	assert.Equal(t, false, isReportSafe([]int{9, 7, 6, 2, 1}))
	assert.Equal(t, false, isReportSafe([]int{1, 3, 2, 4, 5}))
	assert.Equal(t, false, isReportSafe([]int{8, 6, 4, 4, 1}))
	assert.Equal(t, true, isReportSafe([]int{1, 3, 6, 7, 9}))
}

func TestCalculateSafeReports(t *testing.T) {
	assert.Equal(t, 2, calculateSafeReports([][]int{{7, 6, 4, 2, 1}, {1, 2, 7, 8, 9}, {9, 7, 6, 2, 1}, {1, 3, 2, 4, 5}, {8, 6, 4, 4, 1}, {1, 3, 6, 7, 9}}))
}

func TestIsReportSafeWithDampener(t *testing.T) {
	assert.Equal(t, true, isReportSafeWithDampener([]int{7, 6, 4, 2, 1}))
	assert.Equal(t, false, isReportSafeWithDampener([]int{1, 2, 7, 8, 9}))
	assert.Equal(t, false, isReportSafeWithDampener([]int{9, 7, 6, 2, 1}))
	assert.Equal(t, true, isReportSafeWithDampener([]int{1, 3, 2, 4, 5}))
	assert.Equal(t, true, isReportSafeWithDampener([]int{8, 6, 4, 4, 1}))
	assert.Equal(t, true, isReportSafeWithDampener([]int{1, 3, 6, 7, 9}))
}

func TestIsReportSafeWithDampenerSneaky(t *testing.T) {
	assert.Equal(t, true, isReportSafeWithDampener([]int{83, 81, 82, 83, 85, 87, 90, 92}))
}

func TestCalculateSafeReportsWithDampener(t *testing.T) {
	assert.Equal(t, 4, calculateSafeReportsWithDampener([][]int{{7, 6, 4, 2, 1}, {1, 2, 7, 8, 9}, {9, 7, 6, 2, 1}, {1, 3, 2, 4, 5}, {8, 6, 4, 4, 1}, {1, 3, 6, 7, 9}}))
}
