package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Kukmedis/aoc2024/pkg/utils"
)

type robot struct {
	px int
	py int
	vx int
	vy int
}

func parseInput(input []string) []robot {
	var robots []robot
	for _, line := range input {
		var robot robot
		robot.px = utils.ExtractIntFrom(line, "p=")
		robot.py = utils.ExtractIntFrom(line, ",")
		robot.vx = utils.ExtractIntFrom(line, "v=")
		robot.vy = utils.ExtractIntFromLast(line, ",")
		robots = append(robots, robot)
	}
	return robots
}

func tick(robots []robot, xSize int, ySize int) []robot {
	var robotsAfterTick []robot
	for _, r := range robots {
		var robotAfterTick robot
		robotAfterTick.vx = r.vx
		robotAfterTick.vy = r.vy
		robotAfterTick.px = (r.px + r.vx + xSize) % xSize
		robotAfterTick.py = (r.py + r.vy + ySize) % ySize
		robotsAfterTick = append(robotsAfterTick, robotAfterTick)
	}
	return robotsAfterTick
}

func tickNumber(number int, robots []robot, xSize int, ySize int) []robot {
	ticked := robots
	for range number {
		ticked = tick(ticked, xSize, ySize)
	}
	return ticked
}

func calcSafety(robots []robot, xSize int, ySize int) int {
	q1, q2, q3, q4 := 0, 0, 0, 0
	midX := xSize / 2
	midY := ySize / 2
	for _, robot := range robots {
		if robot.px < midX && robot.py < midY {
			q1++
		} else if robot.px > midX && robot.py < midY {
			q2++
		} else if robot.px < midX && robot.py > midY {
			q3++
		} else if robot.px > midX && robot.py > midY {
			q4++
		}
	}
	return q1 * q2 * q3 * q4
}

func printPicture(robots []robot, xSize int, ySize int, iteration int) {
	var picture []string
	for range ySize {
		var row string
		for range xSize {
			row = row + "."
		}
		picture = append(picture, row)
	}
	for _, r := range robots {
		picture[r.py] = replaceAtIndex(picture[r.py], '1', r.px)
	}
	hasStraight := false
	for _, p := range picture {
		hasStraight = strings.Contains(p, "11111111")
		if hasStraight {
			break
		}
	}
	if hasStraight {
		fmt.Println("======================", iteration, "=============================")
		for _, p := range picture {
			fmt.Println(p)
		}
		fmt.Println("======================", iteration, "=============================")
		input := bufio.NewScanner(os.Stdin)
		input.Scan()
	}

}

func replaceAtIndex(in string, r rune, i int) string {
	out := []rune(in)
	out[i] = r
	return string(out)
}
