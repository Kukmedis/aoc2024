package main

type location struct {
	x, y int
}

type direction struct {
	x, y int
}

type priceLocation struct {
	location  location
	direction direction
}

func parseInput(input []string) [][]rune {
	var maze [][]rune
	for i := 0; i < len(input); i++ {
		var row []rune
		for _, c := range input[i] {
			row = append(row, c)
		}
		maze = append(maze, row)
	}
	return maze
}

func findCheapestPath(maze [][]rune) int {
	validPaths := make(map[int][]location)
	priceMap := make(map[priceLocation]int)
	start := find(maze, 'S')
	goal := find(maze, 'E')
	findPath(maze, start, goal, []location{start}, validPaths, priceMap)
	minPrice := 999999999999999
	for price := range validPaths {
		minPrice = min(price, minPrice)
	}
	return minPrice
}

func howManyBestTiles(maze [][]rune) int {
	validPaths := make(map[int][]location)
	priceMap := make(map[priceLocation]int)
	start := find(maze, 'S')
	goal := find(maze, 'E')
	findPath(maze, start, goal, []location{start}, validPaths, priceMap)
	minPrice := 999999999999999
	for price := range validPaths {
		minPrice = min(price, minPrice)
	}
	uniqueTiles := make(map[location]bool)

	for _, loc := range validPaths[minPrice] {
		uniqueTiles[loc] = true
	}
	return len(uniqueTiles)
}

func findPath(maze [][]rune, start location, goal location, path []location, validPaths map[int][]location, priceMap map[priceLocation]int) {
	if start == goal {
		price := calculatePrice(path)
		validPaths[price] = append(validPaths[price], path...)
		return
	}
	direction := getCurrentDirection(path)
	priceLocation := priceLocation{start, direction}
	price := calculatePrice(path)
	if priceMap[priceLocation] == 0 || priceMap[priceLocation] >= price {
		priceMap[priceLocation] = price
	} else {
		return
	}
	locEast := location{start.x + 1, start.y}
	locWest := location{start.x - 1, start.y}
	locNorth := location{start.x, start.y - 1}
	locSouth := location{start.x, start.y + 1}
	if !pathContains(path, locEast) && maze[locEast.y][locEast.x] != '#' {
		findPath(maze, locEast, goal, append(path, locEast), validPaths, priceMap)
	}
	if !pathContains(path, locWest) && maze[locWest.y][locWest.x] != '#' {
		findPath(maze, locWest, goal, append(path, locWest), validPaths, priceMap)
	}
	if !pathContains(path, locNorth) && maze[locNorth.y][locNorth.x] != '#' {
		findPath(maze, locNorth, goal, append(path, locNorth), validPaths, priceMap)
	}
	if !pathContains(path, locSouth) && maze[locSouth.y][locSouth.x] != '#' {
		findPath(maze, locSouth, goal, append(path, locSouth), validPaths, priceMap)
	}
}

func getCurrentDirection(path []location) direction {
	if len(path) < 2 {
		return direction{0, 0}
	} else {
		return direction{path[len(path)-2].x - path[len(path)-1].x, path[len(path)-2].y - path[len(path)-1].y}
	}
}

func calculatePrice(path []location) int {
	currentDir := direction{1, 0}
	sum := 0
	for i := 1; i < len(path); i++ {
		loc := path[i]
		previousLoc := path[i-1]
		nextLoc := location{previousLoc.x + currentDir.x, previousLoc.y + currentDir.y}
		if nextLoc != loc {
			currentDir = direction{loc.x - previousLoc.x, loc.y - previousLoc.y}
			sum = sum + 1000
		}
		previousLoc = loc
	}
	price := sum + len(path) - 1
	return price
}

func find(maze [][]rune, what rune) location {
	for y := 0; y < len(maze); y++ {
		for x := 0; x < len(maze[y]); x++ {
			if maze[y][x] == what {
				return location{x, y}
			}
		}
	}
	panic("Not found!")
}

func pathContains(path []location, location location) bool {
	for _, p := range path {
		if p == location {
			return true
		}
	}
	return false
}
