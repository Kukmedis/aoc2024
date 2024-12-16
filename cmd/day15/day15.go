package main

import (
	"fmt"
	"sort"
)

type box struct {
	x1, x2, y int
}

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

func processWide(input []string) int {
	var data [][]rune
	space := 0
	for i := 0; input[i] != ""; i++ {
		var row []rune
		for _, c := range input[i] {
			if c == '#' {
				row = append(row, '#')
				row = append(row, '#')
			} else if c == '@' {
				row = append(row, '@')
				row = append(row, '.')
			} else if c == '.' {
				row = append(row, '.')
				row = append(row, '.')
			} else if c == 'O' {
				row = append(row, '[')
				row = append(row, ']')
			}

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
		moveWide(data, d)
	}
	return calcScore(data)
}

func moveWide(data [][]rune, direction rune) {
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
		boxes := make(map[box]bool)
		if canPushBoxesUp(data, robotX, robotY-1, boxes) {
			var boxesSlice []box
			for key := range boxes {
				boxesSlice = append(boxesSlice, key)
			}
			pushBoxesUp(data, boxesSlice, robotX, robotY)
		}
	} else if direction == 'v' {
		boxes := make(map[box]bool)
		if canPushBoxesDown(data, robotX, robotY+1, boxes) {
			var boxesSlice []box
			for key := range boxes {
				boxesSlice = append(boxesSlice, key)
			}
			pushBoxesDown(data, boxesSlice, robotX, robotY)
		}
	}

}

func canPushBoxesUp(data [][]rune, x int, y int, boxes map[box]bool) bool {
	if data[y][x] == '[' {
		boxes[box{x, x + 1, y}] = true
		return canPushBoxesUp(data, x, y+1, boxes) && canPushBoxesUp(data, x+1, y+1, boxes)
	} else if data[y][x] == ']' {
		boxes[box{x - 1, x, y, y}] = true
		return canPushBoxesUp(data, x, y+1, boxes) && canPushBoxesUp(data, x-1, y+1, boxes)
	} else if data[y][x] == '#' {
		return false
	}
	return true
}

func canPushBoxesDown(data [][]rune, x int, y int, boxes map[box]bool) bool {
	if data[y][x] == '[' {
		boxes[box{x, x + 1, y}] = true
		return canPushBoxesDown(data, x, y-1, boxes) && canPushBoxesDown(data, x+1, y-1, boxes)
	} else if data[y][x] == ']' {
		boxes[box{x - 1, x, y}] = true
		return canPushBoxesDown(data, x, y-1, boxes) && canPushBoxesDown(data, x-1, y-1, boxes)
	} else if data[y][x] == '#' {
		return false
	}
	return true
}

func pushBoxesUp(data [][]rune, boxes []box, x int, y int) {
	sort.Slice(boxes, func(i int, j int) bool {
		return boxes[i].y < boxes[j].y
	})
	for _, b := range boxes {
		data[b.y-1][b.x1] = '['
		data[b.y-1][b.x2] = ']'
		data[b.y][b.x2] = '.'
		data[b.y][b.x1] = '.'
	}
	data[y-1][x] = '@'
	data[y][x] = '.'
}

func pushBoxesDown(data [][]rune, boxes []box, x int, y int) {
	sort.Slice(boxes, func(i int, j int) bool {
		return boxes[i].y > boxes[j].y
	})
	for _, b := range boxes {
		data[b.y+1][b.x1] = '['
		data[b.y+1][b.x2] = ']'
		data[b.y][b.x2] = '.'
		data[b.y][b.x1] = '.'
	}
	data[y+1][x] = '@'
	data[y][x] = '.'
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
