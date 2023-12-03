package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type RelevatedSet struct {
	RedCubes   int
	GreenCubes int
	BlueCubes  int
}

type Game struct {
	gameIndex int
	RelSets   []RelevatedSet
}

func parseInput(inputData []byte) ([]Game, error) {

	var result []Game
	inputDataRows := strings.Split(string(inputData), "\n")
	for _, dataRow := range inputDataRows {
		dataRowParts := strings.Split(dataRow, ":")
		if len(dataRowParts) != 2 {
			return result, fmt.Errorf("Invalid inputs too many parts")
		}
		labelPart := dataRowParts[0]
		bodyPart := dataRowParts[1]
		labelPart = strings.TrimSpace(strings.ReplaceAll(labelPart, "Game", ""))
		gameIndex, err := strconv.Atoi(labelPart)
		if err != nil {
			return result, err
		}
		currentGame := Game{gameIndex: gameIndex}
		bodySets := strings.Split(bodyPart, ";")
		for _, currSet := range bodySets {
			currSet = strings.TrimSpace(currSet)
			colorsSet := strings.Split(currSet, ",")
			currentRelevatedSet := RelevatedSet{}
			for _, currColor := range colorsSet {
				currColor = strings.TrimSpace(currColor)
				cubeDetails := strings.Split(currColor, " ")
				if len(cubeDetails) != 2 {
					return result, fmt.Errorf("Invalid color", cubeDetails)
				}
				cubeColor := strings.TrimSpace(cubeDetails[1])
				cubeCount, err := strconv.Atoi(strings.TrimSpace(cubeDetails[0]))
				if err != nil {
					return result, err
				}
				if cubeColor == "red" {
					currentRelevatedSet.RedCubes = cubeCount
				} else if cubeColor == "green" {
					currentRelevatedSet.GreenCubes = cubeCount
				} else {
					currentRelevatedSet.BlueCubes = cubeCount
				}
			}
			currentGame.RelSets = append(currentGame.RelSets, currentRelevatedSet)
		}
		result = append(result, currentGame)
	}
	return result, nil
}

func Part1(inputData []byte) int {
	redMax := 12
	greenMax := 13
	blueMax := 14

	games, err := parseInput(inputData)
	if err != nil {
		fmt.Println(err)
		return -1
	}

	result := 0
	for _, game := range games {
		validGame := true
		for _, gameSets := range game.RelSets {
			if gameSets.RedCubes > redMax || gameSets.GreenCubes > greenMax || gameSets.BlueCubes > blueMax {
				validGame = false
				break
			}
		}
		if validGame {
			result += game.gameIndex
		}
	}

	return result
}

func Part2(inputData []byte) int {
	games, err := parseInput(inputData)
	if err != nil {
		fmt.Println(err)
		return -1
	}

	result := 0
	for _, game := range games {
		minRedCubes := 0
		minGreenCubes := 0
		minBlueCubes := 0

		for _, gameSets := range game.RelSets {
			if gameSets.RedCubes > minRedCubes {
				minRedCubes = gameSets.RedCubes
			}
			if gameSets.GreenCubes > minGreenCubes {
				minGreenCubes = gameSets.GreenCubes
			}
			if gameSets.BlueCubes > minBlueCubes {
				minBlueCubes = gameSets.BlueCubes
			}
		}

		result += minRedCubes * minGreenCubes * minBlueCubes
	}
	return result
}

func main() {

	pwd, _ := os.Getwd()
	inputData, err := os.ReadFile(pwd + "/2023/day2/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	//fmt.Println(Part1(inputData))
	fmt.Println(Part2(inputData))
}
