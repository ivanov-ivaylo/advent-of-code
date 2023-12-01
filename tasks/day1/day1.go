package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func getNumberFromRow(row string) int {
	firstDigit := ""
	lastDigit := ""
	for i := 0; i < len(row); i++ {
		if _, err := strconv.Atoi(string(row[i])); err == nil {
			firstDigit = string(row[i])
			break
		}
	}
	for i := len(row) - 1; i >= 0; i-- {
		if _, err := strconv.Atoi(string(row[i])); err == nil {
			lastDigit = string(row[i])
			break
		}
	}

	if number, err := strconv.Atoi(firstDigit + lastDigit); err == nil {
		return number
	}

	return 0
}

type Digits struct {
	word   string
	number string
}

func getNumberFromRowWithWords(row string) int {

	data := []Digits{
		{word: "one", number: "1"},
		{word: "two", number: "2"},
		{word: "three", number: "3"},
		{word: "four", number: "4"},
		{word: "five", number: "5"},
		{word: "six", number: "6"},
		{word: "seven", number: "7"},
		{word: "eight", number: "8"},
		{word: "nine", number: "9"},
	}

	lowFound := Digits{number: "", word: ""}
	highFound := Digits{number: "", word: ""}
	currentLowIndex := math.MaxInt32
	currentHighIndex := -1
	for _, digit := range data {

		firstIndexWord := strings.Index(row, digit.word)
		firstIndexNumber := strings.Index(row, digit.number)
		lastIndexWord := strings.LastIndex(row, digit.word)
		lastIndexNumber := strings.LastIndex(row, digit.number)
		if firstIndexWord >= 0 && firstIndexWord < currentLowIndex {
			currentLowIndex = firstIndexWord
			lowFound = digit
		}
		if firstIndexNumber >= 0 && firstIndexNumber < currentLowIndex {
			currentLowIndex = firstIndexNumber
			lowFound = digit
		}
		if lastIndexWord >= 0 && lastIndexWord > currentHighIndex {
			currentHighIndex = lastIndexWord
			highFound = digit
		}
		if lastIndexNumber >= 0 && lastIndexNumber > currentHighIndex {
			currentHighIndex = lastIndexNumber
			highFound = digit
		}
	}

	if number, err := strconv.Atoi(lowFound.number + highFound.number); err == nil {
		return number
	}

	return 0
}

func Part1(inputData []byte) int {

	inputDataRows := strings.Split(string(inputData), "\n")
	totalNumber := 0
	for _, dataRow := range inputDataRows {
		totalNumber += getNumberFromRow(dataRow)
	}

	return totalNumber
}

func Part2(inputData []byte) int {

	inputDataRows := strings.Split(string(inputData), "\n")
	totalNumber := 0
	for _, dataRow := range inputDataRows {
		totalNumber += getNumberFromRowWithWords(dataRow)
	}

	return totalNumber
}

func main() {

	pwd, _ := os.Getwd()
	inputData, err := os.ReadFile(pwd + "/tasks/day1/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	//fmt.Println(Part1(inputData))
	fmt.Println(Part2(inputData))
}
