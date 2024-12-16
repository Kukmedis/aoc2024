package main

import (
	"fmt"
)

func process(input []string) int {
	var data [][]rune
	space := 0
	for i := 0; input[i] != ""; i++ {
		var row []rune
		for _, c := range input[i] {
			row = append(row, c)
		}
		data = append(data, row)
		space = i + 1
	}

	var directions []rune

	for i := space + 1; i < len(input); i++ {
		for _, c := range input[i] {
			directions = append(directions, c)
		}
	}

	for _, d := range directions {
		move(data, d)
	}
	return calcScore(data)
}

func move(data [][]rune, direction rune) {
	// print(data)
	// fmt.Println("Moving", string(direction))
	robotX, robotY := findRobot(data)
	var path []*rune
	if direction == '>' {
		for i := robotX; data[robotY][i] != '#'; i++ {
			path = append(path, &data[robotY][i])
		}
		push(path)
	} else if direction == '<' {
		for i := robotX; data[robotY][i] != '#'; i-- {
			path = append(path, &data[robotY][i])
		}
		push(path)
	} else if direction == '^' {
		for i := robotY; data[i][robotX] != '#'; i-- {
			path = append(path, &data[i][robotX])
		}
		push(path)
	} else if direction == 'v' {
		for i := robotY; data[i][robotX] != '#'; i++ {
			path = append(path, &data[i][robotX])
		}
		push(path)
	}

}

func findRobot(data [][]rune) (int, int) {
	for y, row := range data {
		for x, c := range row {
			if c == '@' {
				return x, y
			}
		}
	}
	panic("Oh nein!")
}

func push(path []*rune) {
	howMuch := -100
	for i := 0; i < len(path); i++ {
		if *path[i] == '.' {
			howMuch = i
			break
		}
	}
	if howMuch == -100 {
		return
	}
	for i := howMuch; i > 0; i-- {
		*path[i] = *path[i-1]
	}
	*path[0] = '.'
}

func calcScore(data [][]rune) int {
	sum := 0
	for y, row := range data {
		for x, c := range row {
			if c == 'O' {
				sum = y*100 + sum + x
			}
		}
	}
	return sum
}

func print(data [][]rune) {
	fmt.Println("=======================")
	for _, d := range data {
		fmt.Println(string(d))
	}
	fmt.Println("=======================")
}
