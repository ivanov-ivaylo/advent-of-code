package main

import (
	"fmt"
	"os"
	"strings"
)

func parseInput(inputData []byte) [][]int {
	inputDataRows := strings.Split(string(inputData), "\n")

	result := [][]int{}
	for _, dataRow := range inputDataRows {
		row := []int{}
		for _, charItem := range dataRow {
			if string(charItem) == "." {
				row = append(row, 0)
			} else if string(charItem) == "#" {
				row = append(row, 1)
			} else if string(charItem) == "O" {
				row = append(row, 2)
			}
		}
		result = append(result, row)
	}

	return result
}

func moveSingleStoneNorth(grid [][]int, x int, y int) {

	newX := x
	newY := y
	for i := x - 1; i >= 0; i-- {
		if grid[i][y] != 0 {
			break
		} else {
			newX = i
		}
	}
	grid[x][y] = 0
	grid[newX][newY] = 2
}

func moveSingleStoneWest(grid [][]int, x int, y int) {

	newX := x
	newY := y
	for i := y - 1; i >= 0; i-- {
		if grid[x][i] != 0 {
			break
		} else {
			newY = i
		}
	}
	grid[x][y] = 0
	grid[newX][newY] = 2
}

func moveSingleStoneSouth(grid [][]int, x int, y int) {

	newX := x
	newY := y
	for i := x + 1; i < len(grid); i++ {
		if grid[i][y] != 0 {
			break
		} else {
			newX = i
		}
	}
	grid[x][y] = 0
	grid[newX][newY] = 2
}

func moveSingleStoneEast(grid [][]int, x int, y int) {

	newX := x
	newY := y
	for i := y + 1; i < len(grid[0]); i++ {
		if grid[x][i] != 0 {
			break
		} else {
			newY = i
		}
	}
	grid[x][y] = 0
	grid[newX][newY] = 2
}

func moveStonesNorth(grid [][]int) {

	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[0]); y++ {
			if grid[x][y] == 2 { //stone
				moveSingleStoneNorth(grid, x, y)
			}
		}
	}
}

func moveStonesWest(grid [][]int) {

	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[0]); y++ {
			if grid[x][y] == 2 { //stone
				moveSingleStoneWest(grid, x, y)
			}
		}
	}
}

func moveStonesSouth(grid [][]int) {

	for x := len(grid) - 2; x >= 0; x-- {
		for y := 0; y < len(grid[0]); y++ {
			if grid[x][y] == 2 { //stone
				moveSingleStoneSouth(grid, x, y)
			}
		}
	}
}

func moveStonesEast(grid [][]int) {

	for x := 0; x < len(grid); x++ {
		for y := len(grid[0]) - 2; y >= 0; y-- {
			if grid[x][y] == 2 { //stone
				moveSingleStoneEast(grid, x, y)
			}
		}
	}
}

func calculateLoad(grid [][]int) int {

	res := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 2 { //stone
				res += len(grid) - i
			}
		}
	}
	return res
}

func printMatrix(matrix [][]int) {
	for _, row := range matrix {
		for _, element := range row {
			toPrint := "."
			if element == 1 {
				toPrint = "#"
			} else if element == 2 {
				toPrint = "O"
			}
			fmt.Print(toPrint, " ")
		}
		fmt.Println() // Move to the next line after printing a row
	}
	fmt.Println()
	fmt.Println()
}

func Part1(inputData []byte) int {

	grid := parseInput(inputData)

	moveStonesNorth(grid)

	//printMatrix(grid)

	return calculateLoad(grid)
}
func areGridsSame(grid1 [][]int, grid2 [][]int) bool {

	if len(grid1) != len(grid2) {
		return false
	}

	for i := 0; i < len(grid1); i++ {
		for j := 0; j < len(grid1[0]); j++ {
			if grid1[i][j] != grid2[i][j] {
				return false
			}
		}
	}
	return true
}

func matrixToStr(grid [][]int) string {
	str := ""
	for _, row := range grid {
		for _, c := range row {
			str += string(c)
		}
	}
	return str
}

func Part2(inputData []byte) int {
	grid := parseInput(inputData)

	gridsMap := map[string]int{}
	numberOfSpinCycles := 1000000000
	startCycleAt := -1
	endCycleAt := -1
	for i := 0; i < numberOfSpinCycles; i++ {
		moveStonesNorth(grid)
		moveStonesWest(grid)
		moveStonesSouth(grid)
		moveStonesEast(grid)

		state := matrixToStr(grid)
		if val, ok := gridsMap[state]; ok {
			startCycleAt = val
			endCycleAt = i
			break
		} else {
			gridsMap[state] = i
		}
	}
	numberOfSpinCyclesPart := numberOfSpinCycles - startCycleAt
	cycleLen := endCycleAt - startCycleAt

	newNumberOfSpinCycles := numberOfSpinCyclesPart % cycleLen
	for i := 1; i < newNumberOfSpinCycles; i++ {
		moveStonesNorth(grid)
		moveStonesWest(grid)
		moveStonesSouth(grid)
		moveStonesEast(grid)
	}

	return calculateLoad(grid)
}

func main() {

	pwd, _ := os.Getwd()
	inputData, err := os.ReadFile(pwd + "/2023/day14/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(Part1(inputData))
	fmt.Println(Part2(inputData))
}
