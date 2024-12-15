package utils

import (
	"strconv"
	"strings"
	"unicode"
)

func ToInt(value string) int {
	number, err := strconv.Atoi(value)
	if err != nil {
		panic(err)
	}
	return number
}

func ToIntFromRune(value string) int {
	number, err := strconv.Atoi(value)
	if err != nil {
		panic(err)
	}
	return number
}

func ToInts(value string, separator string) []int {
	strings.Split(value, separator)
	var result []int
	for _, v := range strings.Split(value, separator) {
		result = append(result, ToInt(v))
	}
	return result
}

func ExtractIntFrom(input string, marker string) int {
	startIndex := strings.Index(input, marker) + len(marker)
	digitInput := input[startIndex:]
	for i := 0; i < len(digitInput); i++ {
		if !unicode.IsDigit(rune(digitInput[i])) && digitInput[i] != '-' {
			return ToInt(digitInput[0:i])
		}
	}
	return ToInt(digitInput)
}

func ExtractIntFromLast(input string, marker string) int {
	startIndex := strings.LastIndex(input, marker) + len(marker)
	digitInput := input[startIndex:]
	for i := 0; i < len(digitInput); i++ {
		if !unicode.IsDigit(rune(digitInput[i])) && digitInput[i] != '-' {
			return ToInt(digitInput[0:i])
		}
	}
	return ToInt(digitInput)
}
