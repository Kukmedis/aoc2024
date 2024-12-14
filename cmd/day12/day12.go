package main

import (
	"maps"
	"sort"

	"github.com/Kukmedis/aoc2024/pkg/utils"
)

type cell struct {
	x    int
	y    int
	char string
}

type cCell struct {
	x int
	y int
	n bool
	w bool
	e bool
	s bool
}

type priceFunc func(map[int][]cell) int

func groupChar(cells []cell) map[int][]cell {
	grouping := make(map[int][]cell)
	for _, c := range cells {
		var appended bool
		maxKey := 0
		for key, g := range grouping {
			if adjacentToGroup(c, g) {
				grouping[key] = append(g, c)
				appended = true
				break
			}
			maxKey = max(maxKey, key)
		}
		if !appended {
			grouping[maxKey+1] = []cell{c}
		}
	}
	for key, g1 := range maps.Clone(grouping) {
		for anotherKey, g2 := range maps.Clone(grouping) {
			if areGroupsAdjacent(g1, g2) && key != anotherKey {
				grouping[key] = append(grouping[key], grouping[anotherKey]...)
				delete(grouping, anotherKey)
			}
		}
	}
	return grouping
}

func adjacent(c1, c2 cell) bool {
	return c1.char == c2.char && ((c1.x == c2.x && utils.Abs(c1.y-c2.y) == 1) || (c1.y == c2.y && utils.Abs(c1.x-c2.x) == 1))
}

func adjacentCcell(c1, c2 cCell) bool {
	return (c1.x == c2.x && utils.Abs(c1.y-c2.y) == 1) || (c1.y == c2.y && utils.Abs(c1.x-c2.x) == 1)
}

func adjacentToGroup(c cell, cells []cell) bool {
	for _, gc := range cells {
		if adjacent(c, gc) {
			return true
		}
	}
	return false
}

func areGroupsAdjacent(cells1 []cell, cells2 []cell) bool {
	for _, c1 := range cells1 {
		for _, c2 := range cells2 {
			if adjacent(c1, c2) {
				return true
			}
		}
	}
	return false
}

func toCells(text string, row int) []cell {
	var cells []cell
	for x, c := range text {
		if string(c) != " " {
			var cl cell
			cl.y = row
			cl.x = x
			cl.char = string(c)
			cells = append(cells, cl)
		}

	}
	return cells
}

func calcPerimiter(cells []cell) int {
	perimeter := 0
	for _, c := range cells {
		perimeter = perimeter + (4 - howManyNeighbors(c, cells))
	}
	return perimeter
}

func howManyNeighbors(c cell, cells []cell) int {
	neighbors := 0
	for _, gc := range cells {
		if adjacent(c, gc) {
			neighbors++
		}
	}
	return neighbors
}

func calcPrice(cellMap map[int][]cell) int {
	sum := 0
	for _, cells := range cellMap {
		price := calcPerimiter(cells) * len(cells)
		sum = sum + price
	}
	return sum
}

func calcPriceWithSides(cellMap map[int][]cell) int {
	sum := 0
	for _, cells := range cellMap {
		size := len(cells)
		sides := calcSides(cells)
		price := sides * size
		sum = sum + price
	}
	return sum
}

func calcSides(cells []cell) int {
	cCells := convertToCcells(cells)
	nCells := getAllN(cCells)
	wCells := getAllW(cCells)
	eCells := getAllE(cCells)
	sCells := getAllS(cCells)
	return calcNumberOfSidesPerSide(nCells) + calcNumberOfSidesPerSide(wCells) +
		calcNumberOfSidesPerSide(eCells) + calcNumberOfSidesPerSide(sCells)
}

func calcNumberOfSidesPerSide(ccells []cCell) int {
	sum := len(ccells)
	for i, cc := range ccells {
		if doesHaveNeighbor(cc, ccells[i:]) {
			sum--
		}
	}
	return max(1, sum)
}

func doesHaveNeighbor(cCell cCell, cCells []cCell) bool {
	for _, cc := range cCells {
		if adjacentCcell(cCell, cc) {
			return true
		}
	}
	return false
}

func convertToCcells(cells []cell) []cCell {
	var cCells []cCell
	for _, c := range cells {
		var cCell cCell
		cCell.x = c.x
		cCell.y = c.y
		for _, c2 := range cells {
			if c2.x == c.x && c2.y-c.y == 1 {
				cCell.s = true
				break
			}
		}
		for _, c2 := range cells {
			if c2.x == c.x && c2.y-c.y == -1 {
				cCell.n = true
				break
			}
		}
		for _, c2 := range cells {
			if c2.x-c.x == 1 && c2.y == c.y {
				cCell.e = true
				break
			}
		}
		for _, c2 := range cells {
			if c2.x-c.x == -1 && c2.y == c.y {
				cCell.w = true
				break
			}
		}
		cCells = append(cCells, cCell)
	}
	return cCells
}

func getAllN(cCells []cCell) []cCell {
	var ncells []cCell
	for _, cc := range cCells {
		if !cc.n {
			ncells = append(ncells, cc)
		}
	}
	sort.Slice(ncells, func(i, j int) bool {
		return ncells[i].x < ncells[j].x
	})
	return ncells
}
func getAllW(cCells []cCell) []cCell {
	var ncells []cCell
	for _, cc := range cCells {
		if !cc.w {
			ncells = append(ncells, cc)
		}
	}
	sort.Slice(ncells, func(i, j int) bool {
		return ncells[i].y < ncells[j].y
	})
	return ncells
}

func getAllE(cCells []cCell) []cCell {
	var ncells []cCell
	for _, cc := range cCells {
		if !cc.e {
			ncells = append(ncells, cc)
		}
	}
	sort.Slice(ncells, func(i, j int) bool {
		return ncells[i].y < ncells[j].y
	})
	return ncells
}

func getAllS(cCells []cCell) []cCell {
	var ncells []cCell
	for _, cc := range cCells {
		if !cc.s {
			ncells = append(ncells, cc)
		}
	}
	sort.Slice(ncells, func(i, j int) bool {
		return ncells[i].x < ncells[j].x
	})
	return ncells
}

func convToCells(texts []string) []cell {
	y := 0
	var cells []cell
	for _, text := range texts {
		cells = append(cells, toCells(text, y)...)
		y++
	}
	return cells
}

func process(input []string, priceFunc priceFunc) int {
	cells := convToCells(input)
	grouped := groupChar(cells)
	return priceFunc(grouped)
}
