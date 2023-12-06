package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

type Rucksack struct {
	Section1        string
	Section2        string
	SameItems       []string
	HighestPriority int
}

func parseInput(inputData []byte) []Rucksack {
	inputDataRows := strings.Split(string(inputData), "\n")

	result := []Rucksack{}
	for _, dataRow := range inputDataRows {
		sack := Rucksack{Section1: dataRow[0 : len(dataRow)/2], Section2: dataRow[len(dataRow)/2:]}
		result = append(result, sack)
	}

	return result
}

func findSameChars(sacks []Rucksack) {

	for i := 0; i < len(sacks); i++ {
		for _, letter := range sacks[i].Section1 {
			if strings.Contains(sacks[i].Section2, string(letter)) {
				if !slices.Contains(sacks[i].SameItems, string(letter)) {
					score := int(letter)
					if int(letter) < 97 {
						score = int(letter) - 65 //A
						score += 27
					} else {
						score = int(letter) - 97 //A
						score += 1
					}

					sacks[i].SameItems = append(sacks[i].SameItems, string(letter))
					sacks[i].HighestPriority = score
				}
			}
		}

	}
}

func Part1(inputData []byte) int {

	sacks := parseInput(inputData)

	findSameChars(sacks)

	res := 0
	for _, sack := range sacks {
		res += sack.HighestPriority
	}
	return res
}

func Part2(inputData []byte) int {
	return 0
}

func main() {

	pwd, _ := os.Getwd()
	inputData, err := os.ReadFile(pwd + "/2022/day03/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(Part1(inputData))
	//fmt.Println(Part2(inputData))
}
