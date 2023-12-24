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
	Path           []Point
	Visited        map[string]int
}

func parseInput(inputData []byte) ([][]string, Point, Point) {
	inputDataRows := strings.Split(string(inputData), "\n")
	result := [][]string{}
	startPoint := Point{}
	endPoint := Point{}
	for x := 0; x < len(inputDataRows); x++ {
		row := []string{}
		for y := 0; y < len(inputDataRows[x]); y++ {
			if x == 0 && string(inputDataRows[x][y]) == "." {
				startPoint.X = x
				startPoint.Y = y
			} else if x == len(inputDataRows)-1 && string(inputDataRows[x][y]) == "." {
				endPoint.X = x
				endPoint.Y = y
			}
			row = append(row, string(inputDataRows[x][y]))
		}
		result = append(result, row)
	}
	return result, startPoint, endPoint
}

func isValidPoint(grid [][]string, point Point) bool {

	if point.X < 0 || point.X >= len(grid) || point.Y < 0 || point.Y >= len(grid[0]) ||
		grid[point.X][point.Y] == "#" {
		return false
	}

	for i := len(point.Path) - 1; i >= 0; i-- {
		if point.Path[i].X == point.X && point.Path[i].Y == point.Y {
			return false
		}
	}

	//if _, ok := point.Visited[getStrCoordinates(point)]; ok {
	//	return false
	//}

	return true
}

func isValidPoint2(grid [][]string, point Point) bool {

	if point.X < 0 || point.X >= len(grid) || point.Y < 0 || point.Y >= len(grid[0]) ||
		grid[point.X][point.Y] == "#" {
		return false
	}

	if _, ok := point.Visited[getStrCoordinates(point)]; ok {
		return false
	}

	return true
}

func getStrCoordinates(point Point) string {
	return strconv.Itoa(point.X) + ":" + strconv.Itoa(point.Y)
}

func getMapVisited(visited map[string]int, point Point) {

}

func findMaxNumberOfStepsToStart(grid [][]string, statPoint Point, endPoint Point) (int, []Point) {

	queue := []Point{statPoint}
	var currStep Point
	maxSteps := 0
	maxPath := []Point{}
	for len(queue) > 0 {

		currStep, queue = queue[0], queue[1:]

		if currStep.X == endPoint.X && currStep.Y == endPoint.Y {
			fmt.Println("Found Exit: ", currStep.CurrentPathLen, " Current Max: ", maxSteps)
			if currStep.CurrentPathLen > maxSteps {
				maxSteps = currStep.CurrentPathLen
				maxPath = currStep.Path
			}
		}

		p1 := Point{X: currStep.X - 1, Y: currStep.Y, CurrentPathLen: currStep.CurrentPathLen + 1, Path: append(currStep.Path, Point{X: currStep.X, Y: currStep.Y, CurrentPathLen: currStep.CurrentPathLen})}
		p2 := Point{X: currStep.X + 1, Y: currStep.Y, CurrentPathLen: currStep.CurrentPathLen + 1, Path: append(currStep.Path, Point{X: currStep.X, Y: currStep.Y, CurrentPathLen: currStep.CurrentPathLen})}
		p3 := Point{X: currStep.X, Y: currStep.Y + 1, CurrentPathLen: currStep.CurrentPathLen + 1, Path: append(currStep.Path, Point{X: currStep.X, Y: currStep.Y, CurrentPathLen: currStep.CurrentPathLen})}
		p4 := Point{X: currStep.X, Y: currStep.Y - 1, CurrentPathLen: currStep.CurrentPathLen + 1, Path: append(currStep.Path, Point{X: currStep.X, Y: currStep.Y, CurrentPathLen: currStep.CurrentPathLen})}

		if grid[currStep.X][currStep.Y] == ">" {
			p := Point{X: currStep.X, Y: currStep.Y + 1, CurrentPathLen: currStep.CurrentPathLen + 1, Path: append(currStep.Path, Point{X: currStep.X, Y: currStep.Y, CurrentPathLen: currStep.CurrentPathLen})}
			if isValidPoint(grid, p) {
				queue = append([]Point{p}, queue...)
			}
			continue
		} else if grid[currStep.X][currStep.Y] == "<" {
			p := Point{X: currStep.X, Y: currStep.Y - 1, CurrentPathLen: currStep.CurrentPathLen + 1, Path: append(currStep.Path, Point{X: currStep.X, Y: currStep.Y, CurrentPathLen: currStep.CurrentPathLen})}
			if isValidPoint(grid, p) {
				queue = append([]Point{p}, queue...)
			}
			continue
		} else if grid[currStep.X][currStep.Y] == "^" {
			p := Point{X: currStep.X - 1, Y: currStep.Y, CurrentPathLen: currStep.CurrentPathLen + 1, Path: append(currStep.Path, Point{X: currStep.X, Y: currStep.Y, CurrentPathLen: currStep.CurrentPathLen})}
			if isValidPoint(grid, p) {
				queue = append([]Point{p}, queue...)
			}
			continue
		} else if grid[currStep.X][currStep.Y] == "v" {
			p := Point{X: currStep.X + 1, Y: currStep.Y, CurrentPathLen: currStep.CurrentPathLen + 1, Path: append(currStep.Path, Point{X: currStep.X, Y: currStep.Y, CurrentPathLen: currStep.CurrentPathLen})}
			if isValidPoint(grid, p) {
				queue = append([]Point{p}, queue...)
			}
			continue
		}

		if isValidPoint(grid, p1) {
			queue = append([]Point{p1}, queue...)
		}
		if isValidPoint(grid, p2) {
			queue = append([]Point{p2}, queue...)
		}
		if isValidPoint(grid, p3) {
			queue = append([]Point{p3}, queue...)
		}
		if isValidPoint(grid, p4) {
			queue = append([]Point{p4}, queue...)
		}
	}

	return maxSteps, maxPath
}

func findMaxNumberOfStepsToStart2(grid [][]string, statPoint Point, endPoint Point) int {
	queue := []Point{statPoint}
	var currStep Point
	maxSteps := 0
	for len(queue) > 0 {

		currStep, queue = queue[0], queue[1:]

		if currStep.X == endPoint.X && currStep.Y == endPoint.Y {
			fmt.Println("Found Exit: ", currStep.CurrentPathLen, " Current Max: ", maxSteps)
			if currStep.CurrentPathLen > maxSteps {
				maxSteps = currStep.CurrentPathLen
			}
		}

		p1 := Point{X: currStep.X - 1, Y: currStep.Y, CurrentPathLen: currStep.CurrentPathLen + 1, Path: append(currStep.Path, Point{X: currStep.X, Y: currStep.Y, CurrentPathLen: currStep.CurrentPathLen})}
		p2 := Point{X: currStep.X + 1, Y: currStep.Y, CurrentPathLen: currStep.CurrentPathLen + 1, Path: append(currStep.Path, Point{X: currStep.X, Y: currStep.Y, CurrentPathLen: currStep.CurrentPathLen})}
		p3 := Point{X: currStep.X, Y: currStep.Y + 1, CurrentPathLen: currStep.CurrentPathLen + 1, Path: append(currStep.Path, Point{X: currStep.X, Y: currStep.Y, CurrentPathLen: currStep.CurrentPathLen})}
		p4 := Point{X: currStep.X, Y: currStep.Y - 1, CurrentPathLen: currStep.CurrentPathLen + 1, Path: append(currStep.Path, Point{X: currStep.X, Y: currStep.Y, CurrentPathLen: currStep.CurrentPathLen})}

		if isValidPoint(grid, p1) {
			queue = append([]Point{p1}, queue...)
		}
		if isValidPoint(grid, p2) {
			queue = append([]Point{p2}, queue...)
		}
		if isValidPoint(grid, p3) {
			queue = append([]Point{p3}, queue...)
		}
		if isValidPoint(grid, p4) {
			queue = append([]Point{p4}, queue...)
		}
	}

	return maxSteps
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

	grid, startPoint, endPoint := parseInput(inputData)

	res, _ := findMaxNumberOfStepsToStart(grid, startPoint, endPoint)

	return res
}

func Part2(inputData []byte) int {

	grid, startPoint, endPoint := parseInput(inputData)

	res := findMaxNumberOfStepsToStart2(grid, startPoint, endPoint)

	//for _, p := range maxPath {
	//	grid[p.X][p.Y] = "O"
	//}
	//
	//printMatrix(grid)
	//fmt.Println(grid, startPoint, endPoint)

	return res
}

func main() {

	pwd, _ := os.Getwd()
	inputData, err := os.ReadFile(pwd + "/2023/day23/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	//fmt.Println(Part1(inputData))
	fmt.Println(Part2(inputData))
}
