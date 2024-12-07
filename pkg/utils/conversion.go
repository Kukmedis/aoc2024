package utils

import (
	"strconv"
	"strings"
)

func ToInt(value string) int {
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
