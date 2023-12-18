package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Move struct {
	Direction string
	StepNum   int
	Color     string
}

type Point struct {
	X int
	Y int
}

func parseInput(inputData []byte) []Move {
	inputDataRows := strings.Split(string(inputData), "\n")

	result := []Move{}
	for _, dataRow := range inputDataRows {
		rowParts := strings.Split(dataRow, " ")
		currMove := Move{Direction: rowParts[0]}
		if number, err := strconv.Atoi(rowParts[1]); err == nil {
			currMove.StepNum = number
		}
		color := strings.Trim(rowParts[2], "()")
		color = strings.Trim(color, "#")
		currMove.Color = color

		result = append(result, currMove)
	}

	return result
}

func getGrid(sizeX int, sizeY int) [][]int {
	grid := [][]int{}

	for i := 0; i < sizeX; i++ {
		row := []int{}
		for j := 0; j < sizeY; j++ {
			row = append(row, 0)
		}
		grid = append(grid, row)
	}
	return grid
}

func walkPath(grid [][]int, moves []Move, startX int, startY int) int {

	grid[startX][startY] = 1
	//currDirection := "N"
	currX := startX
	currY := startY
	perimeter := 0
	for _, move := range moves {

		//printMatrix(grid)

		perimeter += move.StepNum
		if move.Direction == "U" {
			for i := 1; i <= move.StepNum; i++ {
				currX--
				grid[currX][currY] = 1
			}
		} else if move.Direction == "R" {
			for i := 1; i <= move.StepNum; i++ {
				currY++
				grid[currX][currY] = 1
			}
		} else if move.Direction == "L" {
			for i := 1; i <= move.StepNum; i++ {
				currY--
				grid[currX][currY] = 1
			}
		} else if move.Direction == "D" {
			for i := 1; i <= move.StepNum; i++ {
				currX++
				grid[currX][currY] = 1
			}
		}
	}

	return perimeter
}

func getStrCoordinates(point Point) string {
	return strconv.Itoa(point.X) + ":" + strconv.Itoa(point.Y)
}

func isValidItem(grid [][]int, X, Y int) (bool, int) {

	if grid[X][Y] == 1 {
		return true, -1
	}
	startPoint := Point{X: X, Y: Y}
	queue := []Point{startPoint}
	visited := map[string]int{}
	visited[getStrCoordinates(startPoint)] = 1
	var currPoint Point
	for len(queue) > 0 {
		currPoint, queue = queue[0], queue[1:]

		p1 := Point{currPoint.X + 1, currPoint.Y}
		p2 := Point{currPoint.X - 1, currPoint.Y}
		p3 := Point{currPoint.X, currPoint.Y + 1}
		p4 := Point{currPoint.X, currPoint.Y - 1}

		if p1.X < 0 || p1.X > len(grid) || p1.Y < 0 || p1.Y >= len(grid[0]) ||
			p2.X < 0 || p2.X > len(grid) || p2.Y < 0 || p2.Y >= len(grid[0]) ||
			p3.X < 0 || p3.X > len(grid) || p3.Y < 0 || p3.Y >= len(grid[0]) ||
			p4.X < 0 || p4.X > len(grid) || p4.Y < 0 || p4.Y >= len(grid[0]) {
			return false, -1
		}

		if grid[p1.X][p1.Y] == 0 {
			strP1 := getStrCoordinates(p1)
			if _, ok := visited[strP1]; !ok {
				queue = append(queue, p1)
				visited[strP1] = 1
			}
		}
		if grid[p2.X][p2.Y] == 0 {
			strP2 := getStrCoordinates(p2)
			if _, ok := visited[strP2]; !ok {
				queue = append(queue, p2)
				visited[strP2] = 1
			}
		}
		if grid[p3.X][p3.Y] == 0 {
			strP3 := getStrCoordinates(p3)
			if _, ok := visited[strP3]; !ok {
				queue = append(queue, p3)
				visited[strP3] = 1
			}
		}
		if grid[p4.X][p4.Y] == 0 {
			strP4 := getStrCoordinates(p4)
			if _, ok := visited[strP4]; !ok {
				queue = append(queue, p4)
				visited[strP4] = 1
			}
		}
	}

	return true, len(visited)
}

func findVolumeSize(grid [][]int, perimeter int) int {

	res := 0
	for i := 0; i < len(grid); i++ {
		leftIndex := -1
		rightIndex := -1
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] != 0 {
				leftIndex = j
				break
			}
		}

		for j := len(grid[0]) - 1; j >= 0; j-- {
			if grid[i][j] != 0 {
				rightIndex = j
				break
			}
		}
		if leftIndex >= 0 && rightIndex >= 0 {
			for k := leftIndex; k <= rightIndex; k++ {
				isValid, insideSize := isValidItem(grid, i, k)
				if isValid {
					res++
					if insideSize > 0 {
						return perimeter + insideSize
					}
				}

			}
		}
	}
	return res
}

func isValidItem2(grid [][]int, X, Y int) (bool, bool) {
	if grid[X][Y] == 1 {
		return true, false
	}

	leftNum := 0
	for i := 0; i < Y; i++ {
		if grid[X][i] == 1 {
			leftNum++
			for j := i + 1; j < Y; j++ {
				i = j
				if grid[X][j] == 0 {
					i = j - 1
					break
				}
			}

		}
	}
	rightNum := 0
	for i := len(grid[0]) - 1; i > Y; i-- {
		if grid[X][i] == 1 {
			rightNum++
			for j := i - 1; j > Y; j-- {
				i = j
				if grid[X][j] == 0 {
					i = j + 1
					break
				}
			}
		}
	}

	upNum := 0
	for i := 0; i < X; i++ {
		if grid[i][Y] == 1 {
			upNum++
			for j := i + 1; j > X; j++ {
				i = j
				if grid[i][Y] == 0 {
					i = j - 1
					break
				}
			}
		}
	}

	downNum := 0
	for i := len(grid) - 1; i > X; i-- {
		if grid[i][Y] == 1 {
			downNum++
			for j := i - 1; j > X; j-- {
				i = j
				if grid[i][Y] == 0 {
					i = j + 1
					break
				}
			}
		}
	}

	if leftNum%2 == 1 && rightNum%2 == 1 && upNum%2 == 1 && downNum%2 == 1 {
		return true, false
	}

	return false, leftNum == 0 || rightNum == 0 || upNum == 0 || downNum == 0
}

func printMatrix(matrix [][]int) {
	for _, row := range matrix {
		for _, element := range row {
			fmt.Print(element, " ")
		}
		fmt.Println() // Move to the next line after printing a row
	}
}

func Part1(inputData []byte) int {

	moves := parseInput(inputData)
	gridSize := 800
	startP := 400
	grid := getGrid(gridSize, gridSize)
	startX := startP
	startY := startP

	perimeter := walkPath(grid, moves, startX, startY)

	res := findVolumeSize(grid, perimeter)

	return res
}

func Part2(inputData []byte) int {
	return 0
}

func main() {

	pwd, _ := os.Getwd()
	inputData, err := os.ReadFile(pwd + "/2023/day18/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(Part1(inputData))
	//fmt.Println(Part2(inputData))
}
