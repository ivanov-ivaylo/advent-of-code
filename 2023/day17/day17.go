package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Tile struct {
	X                  int
	Y                  int
	Direction          string
	SameDirectionCount int
	CurrentPathHeat    int
}

func parseInput(inputData []byte) [][]int {
	inputDataRows := strings.Split(string(inputData), "\n")
	result := [][]int{}
	for _, dataRow := range inputDataRows {
		row := []int{}
		for i := 0; i < len(dataRow); i++ {
			if number, err := strconv.Atoi(string(dataRow[i])); err == nil {
				row = append(row, number)
			}
		}
		result = append(result, row)
	}

	return result
}

func getStrCoordinatesWithDirection(tile Tile) string {
	return strconv.Itoa(tile.X) + ":" + strconv.Itoa(tile.Y) + ":" + tile.Direction
}

func findMinHeatLoss(grid [][]int) int {

	visited := map[string]int{}

	res := math.MaxInt
	queue := []Tile{{0, 0, "E", 0, 0 - grid[0][0]}}
	var currTile Tile
	for len(queue) > 0 {
		currTile, queue = queue[0], queue[1:]

		if currTile.X < 0 || currTile.Y < 0 || currTile.X >= len(grid) || currTile.Y >= len(grid[0]) {
			continue
		}
		if _, ok := visited[getStrCoordinatesWithDirection(currTile)]; ok {
			continue
		} else {
			visited[getStrCoordinatesWithDirection(currTile)] = 1
		}

		currTile.CurrentPathHeat += grid[currTile.X][currTile.Y]
		//exit point
		if currTile.X == len(grid)-1 && currTile.Y == len(grid[0])-1 {
			if res > currTile.CurrentPathHeat {
				res = currTile.CurrentPathHeat - grid[currTile.X][currTile.Y]
			}
		}
		if currTile.CurrentPathHeat > res {
			continue
		}

		if currTile.Direction == "N" {
			queue = append(queue, Tile{currTile.X, currTile.Y - 1, "W", 0, currTile.CurrentPathHeat})
			queue = append(queue, Tile{currTile.X, currTile.Y + 1, "E", 0, currTile.CurrentPathHeat})
			if currTile.SameDirectionCount < 10 {
				queue = append(queue, Tile{currTile.X - 1, currTile.Y, "N", currTile.SameDirectionCount + 1, currTile.CurrentPathHeat})
			}
		} else if currTile.Direction == "E" {
			queue = append(queue, Tile{currTile.X - 1, currTile.Y, "N", 0, currTile.CurrentPathHeat})
			queue = append(queue, Tile{currTile.X + 1, currTile.Y, "S", 0, currTile.CurrentPathHeat})
			if currTile.SameDirectionCount < 10 {
				queue = append(queue, Tile{currTile.X, currTile.Y + 1, "E", currTile.SameDirectionCount + 1, currTile.CurrentPathHeat})
			}

		} else if currTile.Direction == "W" {
			queue = append(queue, Tile{currTile.X - 1, currTile.Y, "N", 0, currTile.CurrentPathHeat})
			queue = append(queue, Tile{currTile.X + 1, currTile.Y, "S", 0, currTile.CurrentPathHeat})
			if currTile.SameDirectionCount < 10 {
				queue = append(queue, Tile{currTile.X, currTile.Y - 1, "W", currTile.SameDirectionCount + 1, currTile.CurrentPathHeat})
			}

		} else if currTile.Direction == "S" {
			queue = append(queue, Tile{currTile.X, currTile.Y + 1, "E", 0, currTile.CurrentPathHeat})
			queue = append(queue, Tile{currTile.X, currTile.Y - 1, "W", 0, currTile.CurrentPathHeat})
			if currTile.SameDirectionCount < 10 {
				queue = append(queue, Tile{currTile.X + 1, currTile.Y, "S", currTile.SameDirectionCount + 1, currTile.CurrentPathHeat})
			}
		}
		sort.Slice(queue, func(i, j int) bool {
			return queue[i].CurrentPathHeat < queue[j].CurrentPathHeat
		})
	}

	return res
}

func Part1(inputData []byte) int {

	grid := parseInput(inputData)

	res := findMinHeatLoss(grid)

	//fmt.Println(grid)

	return res
}

func Part2(inputData []byte) int {
	return 0
}

func main() {

	pwd, _ := os.Getwd()
	inputData, err := os.ReadFile(pwd + "/2023/day17/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(Part1(inputData))
	//fmt.Println(Part2(inputData))
}
