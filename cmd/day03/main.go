package main

import (
	"fmt"
	"os"
)

func main() {
	fileContents, err := os.ReadFile("inputs/day03.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(calculateMultipliers(findMultipliers(string(fileContents))))
	fmt.Println(calculateMultipliers(findMultipliers(cleanDonts(string(fileContents)))))
}
