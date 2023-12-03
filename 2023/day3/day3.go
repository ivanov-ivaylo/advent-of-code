package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type NumberLocation struct {
	StartX int
	StartY int
	EndX   int
	EndY   int
}

type Point struct {
	X int
	Y int
}

type ScanNumber struct {
	StartIndex int
	EndIndex   int
	Number     int
}

type StarNumbers struct {
	Numbers []int
}

func readInput(inputData []byte) ([][]string, []string) {
	inputDataRows := strings.Split(string(inputData), "\n")

	var result [][]string
	for _, dataRow := range inputDataRows {
		var row []string
		for _, symbol := range dataRow {
			row = append(row, string(symbol))
		}
		result = append(result, row)
	}
	return result, inputDataRows
}

func isValidPoint(x int, y int, lenX int, lenY int) bool {
	if x >= 0 && x < lenX && y >= 0 && y < lenY {
		return true
	}
	return false
}

func getSurroundingPoints(locationNumber NumberLocation, lenX int, lenY int) []Point {
	var result []Point
	//prefix
	pref := Point{X: locationNumber.StartX - 1, Y: locationNumber.StartY}
	if isValidPoint(pref.X, pref.Y, lenX, lenY) {
		result = append(result, pref)
	}

	//suffix
	suf := Point{X: locationNumber.EndX + 1, Y: locationNumber.EndY}
	if isValidPoint(suf.X, suf.Y, lenX, lenY) {
		result = append(result, suf)
	}

	for i := -1; i <= (locationNumber.EndX-locationNumber.StartX)+1; i++ {

		upPoint := Point{X: locationNumber.StartX + i, Y: locationNumber.StartY - 1}
		downPoint := Point{X: locationNumber.StartX + i, Y: locationNumber.StartY + 1}

		if isValidPoint(upPoint.X, upPoint.Y, lenX, lenY) {
			result = append(result, upPoint)
		}
		if isValidPoint(downPoint.X, downPoint.Y, lenX, lenY) {
			result = append(result, downPoint)
		}
	}

	return result
}

func isSpecialSymbol(s string) bool {
	if s == "." {
		return false
	}
	if isNumber(s) {
		return false
	}
	return true
}

func isPartNumber(locationNumber NumberLocation, input [][]string) bool {
	lenX := len(input[0])
	lenY := len(input)
	targetPoints := getSurroundingPoints(locationNumber, lenX, lenY)

	for _, point := range targetPoints {
		if isSpecialSymbol(input[point.Y][point.X]) {
			return true
		}
	}

	return false
}
func getStarNumber(number int, locationNumber NumberLocation, input [][]string, starMap map[string]StarNumbers) {
	lenX := len(input[0])
	lenY := len(input)
	targetPoints := getSurroundingPoints(locationNumber, lenX, lenY)

	for _, point := range targetPoints {
		if input[point.Y][point.X] == "*" {
			key := strconv.Itoa(point.Y) + ":" + strconv.Itoa(point.X)
			newStarNumbers := StarNumbers{Numbers: []int{number}}
			if _, ok := starMap[key]; ok {
				newStarNumbers.Numbers = append(starMap[key].Numbers, number)
			}
			starMap[key] = newStarNumbers
		}
	}
}

func isNumber(s string) bool {
	if _, err := strconv.Atoi(s); err == nil {
		return true
	}

	return false
}

func findNumbersInRow(str string) []ScanNumber {

	re := regexp.MustCompile("\\d+")
	matches := re.FindAllStringIndex(str, -1)

	var result []ScanNumber
	for _, match := range matches {
		startIndex := match[0]
		endIndex := match[1] - 1
		numberStr := str[startIndex : endIndex+1]

		number, err := strconv.Atoi(numberStr)
		if err != nil {
			fmt.Println("Error converting string to integer:", err)
			continue
		}
		result = append(result, ScanNumber{StartIndex: startIndex, EndIndex: endIndex, Number: number})
	}

	return result
}

func Part1(inputData []byte) int {

	inputMatrix, inputDataRows := readInput(inputData)
	result := 0

	for i := 0; i < len(inputDataRows); i++ {
		foundNumbers := findNumbersInRow(inputDataRows[i])

		for _, foundNum := range foundNumbers {
			locNumber := NumberLocation{StartX: foundNum.StartIndex, StartY: i, EndX: foundNum.EndIndex, EndY: i}
			isPart := isPartNumber(locNumber, inputMatrix)
			if isPart {
				result += foundNum.Number
			}
		}
	}

	return result
}

func Part2(inputData []byte) int {
	inputMatrix, inputDataRows := readInput(inputData)
	result := 0

	starPoints := map[string]StarNumbers{}

	for i := 0; i < len(inputDataRows); i++ {
		foundNumbers := findNumbersInRow(inputDataRows[i])

		for _, foundNum := range foundNumbers {
			locNumber := NumberLocation{StartX: foundNum.StartIndex, StartY: i, EndX: foundNum.EndIndex, EndY: i}
			getStarNumber(foundNum.Number, locNumber, inputMatrix, starPoints)
		}
	}

	for _, starSymbol := range starPoints {
		if len(starSymbol.Numbers) == 2 {
			result += starSymbol.Numbers[0] * starSymbol.Numbers[1]
		}
	}

	return result
}

func main() {

	pwd, _ := os.Getwd()
	inputData, err := os.ReadFile(pwd + "/2023/day3/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	//fmt.Println(Part1(inputData))
	fmt.Println(Part2(inputData))
}
