package main

import (
	"fmt"
	"os"
	"strings"
)

type Point struct {
	X int
	Y int
}

func parseInput(inputData []byte) ([][]string, int, int) {
	inputDataRows := strings.Split(string(inputData), "\n")
	result := [][]string{}
	var startX int
	var startY int
	for i := 0; i < len(inputDataRows); i++ {
		row := []string{}
		for j := 0; j < len(inputDataRows[i]); j++ {
			if string(inputDataRows[i][j]) == "S" {
				startX = i
				startY = j
			}
			row = append(row, string(inputDataRows[i][j]))
		}
		result = append(result, row)
	}
	return result, startX, startY
}

func isStart(currX int, currY int, startX int, startY int) bool {
	if currX == startX && currY == startY {
		return true
	}
	return false
}

func getNextFromStart(maze [][]string, startX int, startY int) (int, int) {
	if maze[startX][startY+1] == "-" || maze[startX][startY+1] == "J" || maze[startX][startY+1] == "7" {
		return startX, startY + 1
	}
	if maze[startX+1][startY] == "|" || maze[startX+1][startY] == "L" || maze[startX+1][startY] == "J" {
		return startX + 1, startY
	}

	if maze[startX][startY-1] == "-" || maze[startX][startY-1] == "L" || maze[startX][startY-1] == "F" {
		return startX, startY + 1
	}
	if maze[startX-1][startY] == "|" || maze[startX-1][startY] == "7" || maze[startX-1][startY] == "F" {
		return startX, startY + 1
	}

	return 0, 0
}

func mazeTravel(maze [][]string, startX int, startY int) int {
	res := 1
	currX := startX + 1
	currY := startY
	currX, currY = getNextFromStart(maze, startX, startY)
	prevX := startX
	prevY := startY
	for !isStart(currX, currY, startX, startY) {
		if maze[currX][currY] == "-" {
			if prevY < currY {
				prevY = currY
				currY++
			} else {
				prevY = currY
				currY--
			}
		} else if maze[currX][currY] == "|" {
			if prevX < currX {
				prevX = currX
				currX++
			} else {
				prevX = currX
				currX--
			}
		} else if maze[currX][currY] == "L" {
			if currX == prevX {
				prevX = currX
				prevY = currY
				currX--

			} else if currY == prevY {
				prevX = currX
				prevY = currY
				currY++
			}
		} else if maze[currX][currY] == "J" {
			if currX == prevX {
				prevX = currX
				prevY = currY
				currX--
			} else {
				prevX = currX
				prevY = currY
				currY--
			}
		} else if maze[currX][currY] == "7" {
			if currX == prevX {
				prevX = currX
				prevY = currY
				currX++
			} else {
				prevX = currX
				prevY = currY
				currY--
			}
		} else if maze[currX][currY] == "F" {
			if currX == prevX {
				prevX = currX
				prevY = currY
				currX++
			} else {
				prevX = currX
				prevY = currY
				currY++
			}
		} else {
			continue
		}

		res++
	}

	return res

}
func copyMaze(maze [][]string) [][]string {
	res := [][]string{}
	for _, row := range maze {
		copyRow := []string{}
		for _, cell := range row {
			copyRow = append(copyRow, cell)
		}
		res = append(res, copyRow)
	}
	return res
}

func mazeTravelStar(maze [][]string, startX int, startY int) int {
	steps := 0
	currX := startX + 1
	currY := startY
	currX, currY = getNextFromStart(maze, startX, startY)
	prevX := startX
	prevY := startY
	visited := []Point{}
	visited = append(visited, Point{X: startX, Y: startY})
	for !isStart(currX, currY, startX, startY) {
		visited = append(visited, Point{X: currX, Y: currY})
		if maze[currX][currY] == "-" {
			if prevY < currY {
				prevY = currY
				currY++
			} else {
				prevY = currY
				currY--
			}
		} else if maze[currX][currY] == "|" {
			if prevX < currX {
				prevX = currX
				currX++
			} else {
				prevX = currX
				currX--
			}
		} else if maze[currX][currY] == "L" {
			if currX == prevX {
				prevX = currX
				prevY = currY
				currX--

			} else if currY == prevY {
				prevX = currX
				prevY = currY
				currY++
			}
		} else if maze[currX][currY] == "J" {
			if currX == prevX {
				prevX = currX
				prevY = currY
				currX--
			} else {
				prevX = currX
				prevY = currY
				currY--
			}
		} else if maze[currX][currY] == "7" {
			if currX == prevX {
				prevX = currX
				prevY = currY
				currX++
			} else {
				prevX = currX
				prevY = currY
				currY--
			}
		} else if maze[currX][currY] == "F" {
			if currX == prevX {
				prevX = currX
				prevY = currY
				currX++
			} else {
				prevX = currX
				prevY = currY
				currY++
			}
		} else {
			continue
		}
		steps += 1
	}

	for _, point := range visited {
		maze[point.X][point.Y] = "*"
		//fmt.Println()
		//fmt.Println()
		//printMatrix(maze)
		//fmt.Println()
		//time.Sleep(1 * time.Second)
	}
	res := 0
	printMatrix(maze)

	for i := 0; i < len(maze); i++ {
		for j := 0; j < len(maze[i]); j++ {
			if maze[i][j] == "*" {
				continue
			}
			if !hasPathOutside(copyMaze(maze), i, j) {
				res++
			}
		}
	}

	printMatrix(maze)

	return res

}
func printMatrix(matrix [][]string) {
	for _, row := range matrix {
		for _, element := range row {
			fmt.Print(element, " ")
		}
		fmt.Println() // Move to the next line after printing a row
	}
}

func isOut(p1 Point, p2 Point, p3 Point, p4 Point, lenX int, lenY int) bool {
	if p1.X < 0 || p1.X >= lenX || p1.Y < 0 || p1.Y >= lenY {
		return true
	}
	if p2.X < 0 || p2.X >= lenX || p2.Y < 0 || p2.Y >= lenY {
		return true
	}
	if p3.X < 0 || p3.X >= lenX || p3.Y < 0 || p3.Y >= lenY {
		return true
	}
	if p4.X < 0 || p4.X >= lenX || p4.Y < 0 || p4.Y >= lenY {
		return true
	}

	return false
}

func hasPathOutside(maze [][]string, X int, Y int) bool {

	queue := []Point{{X, Y}}
	var point Point
	for len(queue) > 0 {
		point, queue = queue[0], queue[1:]

		p1 := Point{point.X - 1, point.Y}
		p2 := Point{point.X + 1, point.Y}
		p3 := Point{point.X, point.Y - 1}
		p4 := Point{point.X, point.Y + 1}

		if isOut(p1, p2, p3, p4, len(maze), len(maze[0])) {
			return true
		}
		if maze[p1.X][p1.Y] != "*" {
			queue = append(queue, p1)
			maze[p1.X][p1.Y] = "*"
		}
		if maze[p2.X][p2.Y] != "*" {
			queue = append(queue, p2)
			maze[p2.X][p2.Y] = "*"
		}
		if maze[p3.X][p3.Y] != "*" {
			queue = append(queue, p3)
			maze[p3.X][p3.Y] = "*"
		}
		if maze[p4.X][p4.Y] != "*" {
			queue = append(queue, p4)
			maze[p4.X][p4.Y] = "*"
		}
	}

	return false
}

func Part1(inputData []byte) int {

	maze, startX, startY := parseInput(inputData)

	res := mazeTravel(maze, startX, startY)

	//for _, m := range maze {
	//	fmt.Println(m)
	//}
	//fmt.Println(startX)
	//fmt.Println(startY)

	return res / 2
}

//not finished
func Part2(inputData []byte) int {

	maze, startX, startY := parseInput(inputData)

	res := mazeTravelStar(maze, startX, startY)
	return res
}

func main() {

	pwd, _ := os.Getwd()
	inputData, err := os.ReadFile(pwd + "/2023/day10/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(Part1(inputData))
	//fmt.Println(Part2(inputData))
}
