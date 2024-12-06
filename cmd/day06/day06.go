package main

import (
	"slices"
)

type position struct {
	y int
	x int
}

type direction struct {
	y int
	x int
}

type positionWithDirection struct {
	pos position
	dir direction
}

func escape(startPosition position, puzzleMap [][]bool) int {
	right := direction{0, 1}
	left := direction{0, -1}
	up := direction{-1, 0}
	down := direction{1, 0}
	currentDirection := up
	currentPosition := startPosition
	visited := make(map[position]bool)
	ylen := len(puzzleMap)
	xlen := len(puzzleMap[0])
	for {
		visited[currentPosition] = true
		nextPosition := position{currentPosition.y + currentDirection.y, currentPosition.x + currentDirection.x}
		if nextPosition.x < 0 || nextPosition.x >= xlen || nextPosition.y < 0 || nextPosition.y >= ylen {
			break
		} else if !puzzleMap[nextPosition.y][nextPosition.x] {
			currentPosition = nextPosition
		} else {
			if currentDirection == right {
				currentDirection = down
			} else if currentDirection == down {
				currentDirection = left
			} else if currentDirection == left {
				currentDirection = up
			} else if currentDirection == up {
				currentDirection = right
			}
		}
	}
	return len(visited)
}

func isInfinite(startPosition position, puzzleMap [][]bool) bool {
	right := direction{0, 1}
	left := direction{0, -1}
	up := direction{-1, 0}
	down := direction{1, 0}
	currentDirection := up
	currentPosition := startPosition
	visited := make(map[positionWithDirection]bool)
	ylen := len(puzzleMap)
	xlen := len(puzzleMap[0])
	for {
		if visited[positionWithDirection{currentPosition, currentDirection}] {
			return true
		}
		visited[positionWithDirection{currentPosition, currentDirection}] = true
		nextPosition := position{currentPosition.y + currentDirection.y, currentPosition.x + currentDirection.x}
		if nextPosition.x < 0 || nextPosition.x >= xlen || nextPosition.y < 0 || nextPosition.y >= ylen {
			return false
		} else if !puzzleMap[nextPosition.y][nextPosition.x] {
			currentPosition = nextPosition
		} else {
			if currentDirection == right {
				currentDirection = down
			} else if currentDirection == down {
				currentDirection = left
			} else if currentDirection == left {
				currentDirection = up
			} else if currentDirection == up {
				currentDirection = right
			}
		}
	}
}

func calculateInfiniteMaps(startPosition position, puzzleMap [][]bool) int {
	sum := 0
	for r, row := range puzzleMap {
		for c := range row {
			newPos := position{r, c}
			if startPosition != newPos {
				newPuzzleMap := copyPuzzleMap(puzzleMap)
				newPuzzleMap[r][c] = true
				if isInfinite(startPosition, newPuzzleMap) {
					sum++
				}
			}
		}
	}
	return sum
}

func copyPuzzleMap(puzzleMap [][]bool) [][]bool {
	var newMap [][]bool
	for _, row := range puzzleMap {
		newMap = append(newMap, slices.Clone(row))
	}
	return newMap
}
