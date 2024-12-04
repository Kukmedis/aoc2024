package main

import (
	"regexp"
	"strconv"
)

type mul struct {
	x int
	y int
}

func cleanDonts(text string) string {
	dontsRegex := `don't\(\).*?do\(\)`
	re := regexp.MustCompile(dontsRegex)
	return re.ReplaceAllString(text, "")
}

func findMultipliers(text string) []mul {
	multipliersRegex := `mul\((\d+),(\d+)\)`
	re := regexp.MustCompile(multipliersRegex)
	matches := re.FindAllStringSubmatch(text, -1)
	var result []mul
	for i := range matches {
		result = append(result, mul{toInt(matches[i][1]), toInt(matches[i][2])})
	}
	return result
}

func calculateMultipliers(multipliers []mul) int {
	sum := 0
	for _, multiplier := range multipliers {
		sum = sum + multiplier.x*multiplier.y
	}
	return sum
}

func toInt(value string) int {
	number, err := strconv.Atoi(value)
	if err != nil {
		panic(err)
	}
	return number
}
