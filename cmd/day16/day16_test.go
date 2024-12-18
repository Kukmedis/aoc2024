package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindCheapestPath(t *testing.T) {
	testInput := []string{
		"###############",
		"#.......#....E#",
		"#.#.###.#.###.#",
		"#.....#.#...#.#",
		"#.###.#####.#.#",
		"#.#.#.......#.#",
		"#.#.#####.###.#",
		"#...........#.#",
		"###.#.#####.#.#",
		"#...#.....#.#.#",
		"#.#.#.###.#.#.#",
		"#.....#...#.#.#",
		"#.###.#.#.#.#.#",
		"#S..#.....#...#",
		"###############",
	}
	assert.Equal(t, 7036, findCheapestPath(parseInput(testInput)))
}

func TestFindCheapestPath2(t *testing.T) {
	testInput := []string{
		"#################",
		"#...#...#...#..E#",
		"#.#.#.#.#.#.#.#.#",
		"#.#.#.#...#...#.#",
		"#.#.#.#.###.#.#.#",
		"#...#.#.#.....#.#",
		"#.#.#.#.#.#####.#",
		"#.#...#.#.#.....#",
		"#.#.#####.#.###.#",
		"#.#.#.......#...#",
		"#.#.###.#####.###",
		"#.#.#...#.....#.#",
		"#.#.#.#####.###.#",
		"#.#.#.........#.#",
		"#.#.#.#########.#",
		"#S#.............#",
		"#################",
	}
	assert.Equal(t, 11048, findCheapestPath(parseInput(testInput)))
}

func TestCalculateBestTiles(t *testing.T) {
	testInput := []string{
		"###############",
		"#.......#....E#",
		"#.#.###.#.###.#",
		"#.....#.#...#.#",
		"#.###.#####.#.#",
		"#.#.#.......#.#",
		"#.#.#####.###.#",
		"#...........#.#",
		"###.#.#####.#.#",
		"#...#.....#.#.#",
		"#.#.#.###.#.#.#",
		"#.....#...#.#.#",
		"#.###.#.#.#.#.#",
		"#S..#.....#...#",
		"###############",
	}
	assert.Equal(t, 45, howManyBestTiles(parseInput(testInput)))
}

func TestCalculateBestTiles2(t *testing.T) {
	testInput := []string{
		"#################",
		"#...#...#...#..E#",
		"#.#.#.#.#.#.#.#.#",
		"#.#.#.#...#...#.#",
		"#.#.#.#.###.#.#.#",
		"#...#.#.#.....#.#",
		"#.#.#.#.#.#####.#",
		"#.#...#.#.#.....#",
		"#.#.#####.#.###.#",
		"#.#.#.......#...#",
		"#.#.###.#####.###",
		"#.#.#...#.....#.#",
		"#.#.#.#####.###.#",
		"#.#.#.........#.#",
		"#.#.#.#########.#",
		"#S#.............#",
		"#################",
	}
	assert.Equal(t, 64, howManyBestTiles(parseInput(testInput)))
}

func TestMultiPathSplit(t *testing.T) {
	maze := [][]rune{
		{'#', '#', '#', '#', '#', '#', '#', '#', '#'},
		{'#', '#', '#', '.', '.', '.', '.', 'E', '#'},
		{'#', '#', '#', '.', '#', '.', '#', '#', '#'},
		{'#', '#', '#', '.', '#', '.', '#', '#', '#'},
		{'#', 'S', '.', '.', '.', '.', '#', '#', '#'},
		{'#', '#', '#', '#', '#', '#', '#', '#', '#'},
	}
	paths := make(map[int][]location)
	priceMap := make(map[priceLocation]int)
	findPath(maze, location{1, 4}, location{7, 1}, []location{{1, 4}}, paths, priceMap)
	assert.Equal(t, 20, len(paths[2009]))
}