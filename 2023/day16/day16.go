package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Tile struct {
	Direction string
	X         int
	Y         int
}

func parseInput(inputData []byte) [][]string {
	inputDataRows := strings.Split(string(inputData), "\n")

	result := [][]string{}
	for _, dataRow := range inputDataRows {
		row := []string{}
		for _, charItem := range dataRow {
			row = append(row, string(charItem))
		}
		result = append(result, row)
	}
	return result
}

func getStrCoordinatesWithDirection(tile Tile) string {
	return strconv.Itoa(tile.X) + ":" + strconv.Itoa(tile.Y) + ":" + tile.Direction
}

func getStrCoordinates(tile Tile) string {
	return strconv.Itoa(tile.X) + ":" + strconv.Itoa(tile.Y)
}

func findEnergizesTiles(grid [][]string, startDirection string, startX int, startY int) int {
	encountered := map[string]int{}
	visited := map[string]int{}

	queue := []Tile{{Direction: startDirection, X: startX, Y: startY}}
	var currTile Tile
	for len(queue) > 0 {
		currTile, queue = queue[0], queue[1:]

		if currTile.X < 0 || currTile.Y < 0 || currTile.X >= len(grid) || currTile.Y >= len(grid[0]) {
			continue
		}
		if _, ok := encountered[getStrCoordinatesWithDirection(currTile)]; ok {
			continue
		}
		encountered[getStrCoordinatesWithDirection(currTile)] = 1

		c := grid[currTile.X][currTile.Y]
		visited[getStrCoordinates(currTile)] = 1

		if c == "|" && (currTile.Direction == "E" || currTile.Direction == "W") {
			queue = append(queue, Tile{Direction: "N", X: currTile.X - 1, Y: currTile.Y})
			queue = append(queue, Tile{Direction: "S", X: currTile.X + 1, Y: currTile.Y})
			continue
		}
		if c == "-" && (currTile.Direction == "N" || currTile.Direction == "S") {
			queue = append(queue, Tile{Direction: "E", X: currTile.X, Y: currTile.Y + 1})
			queue = append(queue, Tile{Direction: "W", X: currTile.X, Y: currTile.Y - 1})
			continue
		}
		if c == "/" {
			if currTile.Direction == "N" {
				queue = append(queue, Tile{Direction: "E", X: currTile.X, Y: currTile.Y + 1})
			} else if currTile.Direction == "E" {
				queue = append(queue, Tile{Direction: "N", X: currTile.X - 1, Y: currTile.Y})
			} else if currTile.Direction == "W" {
				queue = append(queue, Tile{Direction: "S", X: currTile.X + 1, Y: currTile.Y})
			} else if currTile.Direction == "S" {
				queue = append(queue, Tile{Direction: "W", X: currTile.X, Y: currTile.Y - 1})
			}
		} else if c == "\\" {
			if currTile.Direction == "N" {
				queue = append(queue, Tile{Direction: "W", X: currTile.X, Y: currTile.Y - 1})
			} else if currTile.Direction == "E" {
				queue = append(queue, Tile{Direction: "S", X: currTile.X + 1, Y: currTile.Y})
			} else if currTile.Direction == "W" {
				queue = append(queue, Tile{Direction: "N", X: currTile.X - 1, Y: currTile.Y})
			} else if currTile.Direction == "S" {
				queue = append(queue, Tile{Direction: "E", X: currTile.X, Y: currTile.Y + 1})
			}
		} else {
			if currTile.Direction == "N" {
				queue = append(queue, Tile{Direction: currTile.Direction, X: currTile.X - 1, Y: currTile.Y})
			} else if currTile.Direction == "E" {
				queue = append(queue, Tile{Direction: currTile.Direction, X: currTile.X, Y: currTile.Y + 1})
			} else if currTile.Direction == "W" {
				queue = append(queue, Tile{Direction: currTile.Direction, X: currTile.X, Y: currTile.Y - 1})
			} else if currTile.Direction == "S" {
				queue = append(queue, Tile{Direction: currTile.Direction, X: currTile.X + 1, Y: currTile.Y})
			}
		}
	}

	return len(visited)
}

func Part1(inputData []byte) int {

	grid := parseInput(inputData)

	return findEnergizesTiles(grid, "E", 0, 0)
}

func Part2(inputData []byte) int {

	grid := parseInput(inputData)

	res := 0
	for i := 0; i < len(grid[0]); i++ {
		curr := findEnergizesTiles(grid, "S", 0, i)
		if curr > res {
			res = curr
		}
	}
	for i := 0; i < len(grid); i++ {
		curr := findEnergizesTiles(grid, "E", i, 0)
		if curr > res {
			res = curr
		}
	}
	for i := 0; i < len(grid[0]); i++ {
		curr := findEnergizesTiles(grid, "N", len(grid)-1, i)
		if curr > res {
			res = curr
		}
	}

	for i := 0; i < len(grid); i++ {
		curr := findEnergizesTiles(grid, "W", i, len(grid[0]))
		if curr > res {
			res = curr
		}
	}

	return res
}

func main() {

	pwd, _ := os.Getwd()
	inputData, err := os.ReadFile(pwd + "/2023/day16/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(Part1(inputData))
	fmt.Println(Part2(inputData))
}
