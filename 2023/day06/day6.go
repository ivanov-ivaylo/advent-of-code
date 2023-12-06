package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Race struct {
	Time              int
	MaxDistance       int
	NumberOfWaysToWin int
}

func parseInput(inputData []byte) ([]Race, error) {

	result := []Race{}
	inputDataRows := strings.Split(string(inputData), "\n")
	if len(inputDataRows) != 2 {
		return result, fmt.Errorf("Wrong input, inputDataRows")
	}
	timeParts := strings.Split(inputDataRows[0], " ")
	for _, timePart := range timeParts {
		if number, err := strconv.Atoi(timePart); err == nil {
			race := Race{Time: number}
			result = append(result, race)
		}
	}

	distanceParts := strings.Split(inputDataRows[1], " ")
	index := 0
	for _, distancePart := range distanceParts {
		if number, err := strconv.Atoi(distancePart); err == nil {
			result[index].MaxDistance = number
			index++
		}
	}

	return result, nil

}

func findNumberOfWins(races []Race) {

	for i := 0; i < len(races); i++ {
		numberOfWins := 0
		for d := 0; d < races[i].Time; d++ {
			speed := d
			time := races[i].Time - d
			distance := speed * time
			if distance > races[i].MaxDistance {
				numberOfWins++
			}
		}
		races[i].NumberOfWaysToWin = numberOfWins
	}
}

func convertInput(races []Race) Race {

	time := ""
	distance := ""

	for _, race := range races {
		t := strconv.Itoa(race.Time)
		time += t
		d := strconv.Itoa(race.MaxDistance)
		distance += d
	}
	res := Race{}
	if number, err := strconv.Atoi(time); err == nil {
		res.Time = number
	}
	if number, err := strconv.Atoi(distance); err == nil {
		res.MaxDistance = number
	}

	return res

}

func Part1(inputData []byte) int {

	races, err := parseInput(inputData)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	findNumberOfWins(races)

	res := 1
	for _, race := range races {
		res *= race.NumberOfWaysToWin
	}

	//fmt.Println(races)

	return res
}

func Part2(inputData []byte) int {

	races, err := parseInput(inputData)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	race := convertInput(races)

	races = []Race{race}
	findNumberOfWins(races)

	res := 1
	for _, currRace := range races {
		res *= currRace.NumberOfWaysToWin
	}

	return res
}

func main() {

	pwd, _ := os.Getwd()
	inputData, err := os.ReadFile(pwd + "/2023/day06/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	//fmt.Println(Part1(inputData))
	fmt.Println(Part2(inputData))
}
