package main

import (
	"fmt"
	"slices"
)

func isReportSafe(report []int) bool {
	reportSize := len(report)
	safe := true
	increasing := false
	decreasing := false
	for i := 0; i < reportSize-1; i++ {
		if 0 < report[i]-report[i+1] && report[i]-report[i+1] < 4 {
			decreasing = true
		} else if -4 < report[i]-report[i+1] && report[i]-report[i+1] < 0 {
			increasing = true
		} else {
			safe = false
			break
		}
		if increasing && decreasing {
			safe = false
			break
		}
	}
	return safe
}

func isReportSafeWithDampener(report []int) bool {
	reportSize := len(report)
	safe := true
	increasing := false
	decreasing := false

	for i := 0; i < reportSize-1; i++ {
		if 0 < report[i]-report[i+1] && report[i]-report[i+1] < 4 {
			decreasing = true
		} else if -4 < report[i]-report[i+1] && report[i]-report[i+1] < 0 {
			increasing = true
		} else {
			safe = isReportSafe(removeIndex(report, i)) || isReportSafe(removeIndex(report, i+1))
			break
		}
		if increasing && decreasing {
			safe = isReportSafe(removeIndex(report, i)) || isReportSafe(removeIndex(report, i+1)) || isReportSafe(removeIndex(report, i-1))
			break
		}
	}
	if !safe {
		fmt.Println("Is Report Safe With Dampener", report, safe)
	}
	return safe
}

func calculateSafeReports(reports [][]int) int {
	sum := 0
	for i := 0; i < len(reports); i++ {
		if isReportSafe(reports[i]) {
			sum++
		}
	}
	return sum
}

func calculateSafeReportsWithDampener(reports [][]int) int {
	sum := 0
	for i := 0; i < len(reports); i++ {
		if isReportSafeWithDampener(reports[i]) {
			sum++
		}
	}
	return sum
}

func removeIndex(s []int, index int) []int {
	clonedS := slices.Clone(s)
	return append(clonedS[:index], clonedS[index+1:]...)
}
