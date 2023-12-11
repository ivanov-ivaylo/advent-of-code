package main

import (
	"fmt"
	"os"
	"strings"
)

func parseInput(inputData []byte) ([][]int, int) {
	inputDataRows := strings.Split(string(inputData), "\n")

	result := [][]int{}
	counter := 1
	for _, dataRow := range inputDataRows {
		row := []int{}
		for _, charItem := range dataRow {
			if string(charItem) == "#" {
				row = append(row, counter)
				counter++
			} else {
				row = append(row, 0)
			}
		}
		result = append(result, row)
	}

	return result, counter - 1
}

func containsItem(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func expandGalaxy(galaxy [][]int) [][]int {
	result := [][]int{}
	columnsToExpand := []int{}
	for y := 0; y < len(galaxy[0]); y++ {
		hasStar := false
		for x := 0; x < len(galaxy); x++ {
			if galaxy[x][y] != 0 {
				hasStar = true
				break
			}
		}
		if !hasStar {
			columnsToExpand = append(columnsToExpand, y)
		}
	}

	for i := 0; i < len(galaxy); i++ {
		isEmptyRow := true
		newRow := []int{}
		for j := 0; j < len(galaxy[0]); j++ {
			newRow = append(newRow, galaxy[i][j])
			if containsItem(columnsToExpand, j) {
				newRow = append(newRow, galaxy[i][j])
			}
			if galaxy[i][j] != 0 {
				isEmptyRow = false
			}
		}
		if isEmptyRow {
			newRowCopy := []int{}
			newRowCopy = append([]int(nil), newRow...)
			result = append(result, newRowCopy)
		}
		result = append(result, newRow)

	}

	return result
}

func findExpandedGalaxyParts(galaxy [][]int) ([]int, []int) {
	columnsToExpand := []int{}
	for y := 0; y < len(galaxy[0]); y++ {
		hasStar := false
		for x := 0; x < len(galaxy); x++ {
			if galaxy[x][y] != 0 {
				hasStar = true
				break
			}
		}
		if !hasStar {
			columnsToExpand = append(columnsToExpand, y)
		}
	}

	rowsToExpand := []int{}
	for i := 0; i < len(galaxy); i++ {
		isEmptyRow := true
		for j := 0; j < len(galaxy[0]); j++ {
			if galaxy[i][j] != 0 {
				isEmptyRow = false
				break
			}
		}
		if isEmptyRow {
			rowsToExpand = append(rowsToExpand, i)
		}
	}

	return rowsToExpand, columnsToExpand
}

func findIndexes(galaxy [][]int, star int) (int, int) {
	for i := 0; i < len(galaxy); i++ {
		for j := 0; j < len(galaxy[0]); j++ {
			if galaxy[i][j] == star {
				return i, j
			}
		}
	}
	return -1, -1
}

func findDistance(galaxy [][]int, star1 int, star2 int) int {

	x1, y1 := findIndexes(galaxy, star1)
	x2, y2 := findIndexes(galaxy, star2)

	d1 := x1 - x2
	d2 := y1 - y2
	if d1 < 0 {
		d1 = d1 * -1
	}
	if d2 < 0 {
		d2 = d2 * -1
	}

	return d1 + d2

}

func findDistance2(galaxy [][]int, star1 int, star2 int, rowsToExpand []int, columnsToExpand []int) int {
	x1, y1 := findIndexes(galaxy, star1)
	x2, y2 := findIndexes(galaxy, star2)

	distance := 0
	expandDistance := 1000000
	if x1 < x2 {
		for i := x1 + 1; i <= x2; i++ {
			if containsItem(rowsToExpand, i) {
				distance += expandDistance
			} else {
				distance += 1
			}
		}
	} else {
		for i := x2 + 1; i <= x1; i++ {
			if containsItem(rowsToExpand, i) {
				distance += expandDistance
			} else {
				distance += 1
			}
		}
	}
	if y1 < y2 {
		for j := y1 + 1; j <= y2; j++ {
			if containsItem(columnsToExpand, j) {
				distance += expandDistance
			} else {
				distance += 1
			}
		}
	} else {
		for j := y2 + 1; j <= y1; j++ {
			if containsItem(columnsToExpand, j) {
				distance += expandDistance
			} else {
				distance += 1
			}
		}
	}

	return distance
}

func Part1(inputData []byte) int {

	galaxy, starsCount := parseInput(inputData)
	galaxy = expandGalaxy(galaxy)

	res := 0
	for i := 1; i <= starsCount; i++ {
		for j := i + 1; j <= starsCount; j++ {
			res += findDistance(galaxy, i, j)
		}
	}

	return res
}

func Part2(inputData []byte) int {

	galaxy, starsCount := parseInput(inputData)

	rowsToExpand, columnsToExpand := findExpandedGalaxyParts(galaxy)

	res := 0
	for i := 1; i <= starsCount; i++ {
		for j := i + 1; j <= starsCount; j++ {
			res += findDistance2(galaxy, i, j, rowsToExpand, columnsToExpand)
		}
	}

	return res
}

func main() {

	pwd, _ := os.Getwd()
	inputData, err := os.ReadFile(pwd + "/2023/day11/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(Part1(inputData))
	fmt.Println(Part2(inputData))
}
