package main

type location struct {
	x int
	y int
}

type mapSize struct {
	xLen int
	yLen int
}

type antinodeFunc func(location, location, mapSize) []location

func calculateAntinodes(antennaLocations []location, puzzleSize mapSize, calcFunc antinodeFunc) []location {
	var resultLocations []location
	for i := 0; i < len(antennaLocations); i++ {
		for j := i + 1; j < len(antennaLocations); j++ {
			antinodes := calcFunc(antennaLocations[i], antennaLocations[j], puzzleSize)
			resultLocations = append(resultLocations, antinodes...)
		}
	}
	return resultLocations
}

func calculateAntinodesPair(antenna1 location, antenna2 location, puzzleSize mapSize) []location {
	var resultLocations []location
	var leftAntenna location
	var rightAntenna location
	if antenna1.x <= antenna2.x {
		leftAntenna = antenna1
		rightAntenna = antenna2
	} else {
		leftAntenna = antenna2
		rightAntenna = antenna1
	}
	var antinode1 location
	var antinode2 location
	xDistance := rightAntenna.x - leftAntenna.x
	if leftAntenna.y >= rightAntenna.y {
		yDistance := leftAntenna.y - rightAntenna.y
		antinode1 = location{leftAntenna.x - xDistance, leftAntenna.y + yDistance}
		antinode2 = location{rightAntenna.x + xDistance, rightAntenna.y - yDistance}
	} else {
		yDistance := rightAntenna.y - leftAntenna.y
		antinode1 = location{leftAntenna.x - xDistance, leftAntenna.y - yDistance}
		antinode2 = location{rightAntenna.x + xDistance, rightAntenna.y + yDistance}
	}
	if doesAntinodeFit(antinode1, puzzleSize) {
		resultLocations = append(resultLocations, antinode1)
	}
	if doesAntinodeFit(antinode2, puzzleSize) {
		resultLocations = append(resultLocations, antinode2)
	}
	return resultLocations
}

func calculateAntinodesPairTFrequency(antenna1 location, antenna2 location, puzzleSize mapSize) []location {
	var resultLocations []location
	var leftAntenna location
	var rightAntenna location
	if antenna1.x <= antenna2.x {
		leftAntenna = antenna1
		rightAntenna = antenna2
	} else {
		leftAntenna = antenna2
		rightAntenna = antenna1
	}
	xDistance := rightAntenna.x - leftAntenna.x
	if leftAntenna.y >= rightAntenna.y {
		yDistance := leftAntenna.y - rightAntenna.y
		nextLeftAntinode := location{leftAntenna.x, leftAntenna.y}
		for nextLeftAntinode.x >= 0 && nextLeftAntinode.y >= 0 &&
			nextLeftAntinode.x < puzzleSize.xLen &&
			nextLeftAntinode.y < puzzleSize.yLen {
			resultLocations = append(resultLocations, nextLeftAntinode)
			nextLeftAntinode = location{nextLeftAntinode.x - xDistance, nextLeftAntinode.y + yDistance}
		}
		nextRightAntinode := location{rightAntenna.x, rightAntenna.y}
		for nextRightAntinode.x >= 0 && nextRightAntinode.y >= 0 &&
			nextRightAntinode.x < puzzleSize.xLen &&
			nextRightAntinode.y < puzzleSize.yLen {
			resultLocations = append(resultLocations, nextRightAntinode)
			nextRightAntinode = location{nextRightAntinode.x + xDistance, nextRightAntinode.y - yDistance}
		}
	} else {
		yDistance := rightAntenna.y - leftAntenna.y
		nextLeftAntinode := location{leftAntenna.x, leftAntenna.y}
		for nextLeftAntinode.x >= 0 && nextLeftAntinode.y >= 0 &&
			nextLeftAntinode.x < puzzleSize.xLen &&
			nextLeftAntinode.y < puzzleSize.yLen {
			resultLocations = append(resultLocations, nextLeftAntinode)
			nextLeftAntinode = location{nextLeftAntinode.x - xDistance, nextLeftAntinode.y - yDistance}
		}
		nextRightAntinode := location{rightAntenna.x, rightAntenna.y}
		for nextRightAntinode.x >= 0 && nextRightAntinode.y >= 0 &&
			nextRightAntinode.x < puzzleSize.xLen &&
			nextRightAntinode.y < puzzleSize.yLen {
			resultLocations = append(resultLocations, nextRightAntinode)
			nextRightAntinode = location{nextRightAntinode.x + xDistance, nextRightAntinode.y + yDistance}
		}
	}
	return resultLocations
}

func doesAntinodeFit(antinode location, puzzleSize mapSize) bool {
	return antinode.x >= 0 && antinode.x < puzzleSize.xLen && antinode.y >= 0 && antinode.y < puzzleSize.yLen
}

func countMapAntinodes(antinodeMap map[rune][]location, puzzleSize mapSize) int {
	return countMapAntinodesWithFunction(antinodeMap, puzzleSize, calculateAntinodesPair)
}

func countMapAntinodesTFrequency(antinodeMap map[rune][]location, puzzleSize mapSize) int {
	return countMapAntinodesWithFunction(antinodeMap, puzzleSize, calculateAntinodesPairTFrequency)
}

func countMapAntinodesWithFunction(antinodeMap map[rune][]location, puzzleSize mapSize, calcFunc antinodeFunc) int {
	allAntinodeLocations := make(map[location]bool)
	for _, k := range antinodeMap {
		for _, antinode := range calculateAntinodes(k, puzzleSize, calcFunc) {
			allAntinodeLocations[antinode] = true
		}
	}
	return len(allAntinodeLocations)
}
