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
	Badge           string
	BadgePriority   int
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

func getPriority(letter int) int {
	score := letter
	if int(letter) < 97 {
		score = int(letter) - 65 //A
		score += 27
	} else {
		score = int(letter) - 97 //a
		score += 1
	}

	return score
}

func findSameChars(sacks []Rucksack) {

	for i := 0; i < len(sacks); i++ {
		for _, letter := range sacks[i].Section1 {
			if strings.Contains(sacks[i].Section2, string(letter)) {
				if !slices.Contains(sacks[i].SameItems, string(letter)) {
					sacks[i].SameItems = append(sacks[i].SameItems, string(letter))
					sacks[i].HighestPriority = getPriority(int(letter))
				}
			}
		}

	}
}

func findBadgeForRucksack(sacks []Rucksack) int {

	res := 0
	for i := 0; i < len(sacks); i += 3 {
		item1 := sacks[i].Section1 + sacks[i].Section2
		item2 := sacks[i+1].Section1 + sacks[i+1].Section2
		item3 := sacks[i+2].Section1 + sacks[i+2].Section2

		for _, item := range item1 {
			if strings.Contains(item2, string(item)) && strings.Contains(item3, string(item)) {
				sacks[i].Badge = string(item)
				sacks[i+1].Badge = string(item)
				sacks[i+2].Badge = string(item)
				sacks[i].BadgePriority = getPriority(int(item))
				sacks[i+1].BadgePriority = getPriority(int(item))
				sacks[i+2].BadgePriority = getPriority(int(item))
				res += getPriority(int(item))
				break
			}
		}
	}

	return res
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

	sacks := parseInput(inputData)
	findSameChars(sacks)

	return findBadgeForRucksack(sacks)
}

func main() {

	pwd, _ := os.Getwd()
	inputData, err := os.ReadFile(pwd + "/2022/day03/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	//fmt.Println(Part1(inputData))
	fmt.Println(Part2(inputData))
}
