package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	X              int
	Y              int
	CurrentPathLen int
}

func parseInput(inputData []byte) ([][]string, Point) {
	inputDataRows := strings.Split(string(inputData), "\n")
	result := [][]string{}
	startPoint := Point{CurrentPathLen: 0}
	for x := 0; x < len(inputDataRows); x++ {
		row := []string{}
		for y := 0; y < len(inputDataRows[x]); y++ {
			if string(inputDataRows[x][y]) == "S" {
				startPoint.X = x
				startPoint.Y = y
				row = append(row, ".")
			} else {
				row = append(row, string(inputDataRows[x][y]))
			}
		}
		result = append(result, row)
	}
	return result, startPoint
}
func isValidPoint(grid [][]string, point Point) bool {

	if point.X < 0 || point.X >= len(grid) || point.Y < 0 || point.Y >= len(grid[0]) ||
		grid[point.X][point.Y] == "#" {
		return false
	}

	return true
}

func getStrCoordinates(point Point) string {
	return strconv.Itoa(point.X) + ":" + strconv.Itoa(point.Y)
}

func findNumberOfStepsToStart(grid [][]string, statPoint Point, endPoint Point, maxNumberOfSteps int) int {

	queue := []Point{statPoint}
	var currStep Point
	visited := map[string]int{}
	for len(queue) > 0 {

		currStep, queue = queue[0], queue[1:]
		if currStep.X == endPoint.X && currStep.Y == endPoint.Y {
			return currStep.CurrentPathLen
		}
		p1 := Point{X: currStep.X - 1, Y: currStep.Y, CurrentPathLen: currStep.CurrentPathLen + 1}
		p2 := Point{X: currStep.X + 1, Y: currStep.Y, CurrentPathLen: currStep.CurrentPathLen + 1}
		p3 := Point{X: currStep.X, Y: currStep.Y + 1, CurrentPathLen: currStep.CurrentPathLen + 1}
		p4 := Point{X: currStep.X, Y: currStep.Y - 1, CurrentPathLen: currStep.CurrentPathLen + 1}

		if currStep.CurrentPathLen >= maxNumberOfSteps {
			continue
		}
		if _, ok := visited[getStrCoordinates(currStep)]; ok {
			continue
		}
		visited[getStrCoordinates(currStep)] = 1

		if isValidPoint(grid, p1) {
			queue = append(queue, p1)
		}
		if isValidPoint(grid, p2) {
			queue = append(queue, p2)
		}
		if isValidPoint(grid, p3) {
			queue = append(queue, p3)
		}
		if isValidPoint(grid, p4) {
			queue = append(queue, p4)
		}
	}

	return -1
}

func findDestinationPoints(grid [][]string, statPoint Point, numberOfSteps int) map[string]Point {

	queue := []Point{statPoint}
	result := map[string]Point{}
	var currStep Point
	for len(queue) > 0 {

		currStep, queue = queue[0], queue[1:]
		if currStep.CurrentPathLen == numberOfSteps {
			result[getStrCoordinates(currStep)] = currStep
			continue
		}
		p1 := Point{X: currStep.X - 1, Y: currStep.Y, CurrentPathLen: currStep.CurrentPathLen + 1}
		p2 := Point{X: currStep.X + 1, Y: currStep.Y, CurrentPathLen: currStep.CurrentPathLen + 1}
		p3 := Point{X: currStep.X, Y: currStep.Y + 1, CurrentPathLen: currStep.CurrentPathLen + 1}
		p4 := Point{X: currStep.X, Y: currStep.Y - 1, CurrentPathLen: currStep.CurrentPathLen + 1}

		if isValidPoint(grid, p1) {
			queue = append(queue, p1)
		}
		if isValidPoint(grid, p2) {
			queue = append(queue, p2)
		}
		if isValidPoint(grid, p3) {
			queue = append(queue, p3)
		}
		if isValidPoint(grid, p4) {
			queue = append(queue, p4)
		}
	}

	return result
}

func printMatrix(matrix [][]string) {
	for _, row := range matrix {
		for _, element := range row {
			fmt.Print(element, " ")
		}
		fmt.Println() // Move to the next line after printing a row
	}
}

func Part1(inputData []byte) int {

	grid, startPoint := parseInput(inputData)

	numOfSteps := 64
	res := 0
	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[0]); y++ {
			if grid[x][y] == "." {
				rangeSteps := findNumberOfStepsToStart(grid, startPoint, Point{X: x, Y: y}, numOfSteps)
				if rangeSteps < 0 {
					continue
				}
				if rangeSteps <= numOfSteps {
					if numOfSteps%2 == 0 && rangeSteps%2 == 0 {
						res++
					}
					if numOfSteps%2 == 1 && rangeSteps%2 == 1 {
						res++
					}
				}

			}
		}
	}

	return res
}

func Part2(inputData []byte) int {
	grid, startPoint := parseInput(inputData)

	numOfSteps := 26501365
	res := 0
	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[0]); y++ {
			if grid[x][y] == "." {
				rangeSteps := findNumberOfStepsToStart(grid, startPoint, Point{X: x, Y: y}, numOfSteps)
				if rangeSteps < 0 {
					continue
				}
				if rangeSteps <= numOfSteps {
					if numOfSteps%2 == 0 && rangeSteps%2 == 0 {
						res++
					}
					if numOfSteps%2 == 1 && rangeSteps%2 == 1 {
						res++
					}
				}

			}
		}
	}

	return res
}

func main() {

	pwd, _ := os.Getwd()
	inputData, err := os.ReadFile(pwd + "/2023/day21/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	//fmt.Println(Part1(inputData))
	fmt.Println(Part2(inputData))
}
