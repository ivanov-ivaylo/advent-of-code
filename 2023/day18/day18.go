package main

import (
	"fmt"
	"math"
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

func findVolumeSize(moves []Move) int {

	move := map[string]Point{
		"R": {0, 1},
		"L": {0, -1},
		"U": {-1, 0},
		"D": {1, 0},
	}

	perimeter := 0
	x := 0
	y := 0
	borders := []Point{}
	for _, m := range moves {
		currPoint := move[m.Direction]
		borders = append(borders, Point{X: x, Y: y})
		perimeter += m.StepNum
		x += currPoint.X * m.StepNum
		y += currPoint.Y * m.StepNum
	}
	area := 0
	for i := 0; i < len(borders)-1; i++ {
		area += borders[i].X*borders[i+1].Y - borders[i+1].X*borders[i].Y
	}
	area = int(math.Abs(float64(area))) / 2

	interior_area := area - (perimeter / 2) + 1
	return interior_area + perimeter
}

func Part1(inputData []byte) int {

	moves := parseInput(inputData)

	return findVolumeSize(moves)
}

func getMovesFromColors(moves []Move) []Move {
	result := []Move{}

	for _, move := range moves {
		colorDistance := move.Color[0:5]
		colorDirection := move.Color[5:]

		distance, _ := strconv.ParseInt(colorDistance, 16, 64)
		dir := ""
		if colorDirection == "0" {
			dir = "R"
		}
		if colorDirection == "1" {
			dir = "D"
		}
		if colorDirection == "2" {
			dir = "L"
		}
		if colorDirection == "3" {
			dir = "U"
		}

		result = append(result, Move{StepNum: int(distance), Direction: dir})
	}

	return result
}

func Part2(inputData []byte) int {

	moves := parseInput(inputData)

	moves = getMovesFromColors(moves)

	return findVolumeSize(moves)
}

func main() {

	pwd, _ := os.Getwd()
	inputData, err := os.ReadFile(pwd + "/2023/day18/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(Part1(inputData))
	fmt.Println(Part2(inputData))
}
