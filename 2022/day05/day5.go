package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Container struct {
	Crates []string
}

type Move struct {
	CountToMove      int
	SourceIndex      int
	DestinationIndex int
}

func parseInput(inputData []byte) ([]Move, []Container) {

	conteiners := []Container{}
	conteiners = append(conteiners, Container{Crates: []string{"B", "S", "J", "Z", "V", "D", "G"}})      //1
	conteiners = append(conteiners, Container{Crates: []string{"P", "V", "G", "M", "S", "Z"}})           //2
	conteiners = append(conteiners, Container{Crates: []string{"F", "Q", "T", "W", "S", "B", "L", "C"}}) //3
	conteiners = append(conteiners, Container{Crates: []string{"Q", "V", "R", "M", "W", "G", "J", "H"}}) //4
	conteiners = append(conteiners, Container{Crates: []string{"D", "M", "F", "N", "S", "L", "C"}})      //5
	conteiners = append(conteiners, Container{Crates: []string{"D", "C", "G", "R"}})                     //6
	conteiners = append(conteiners, Container{Crates: []string{"Q", "S", "D", "J", "R", "T", "G", "H"}}) //7
	conteiners = append(conteiners, Container{Crates: []string{"V", "F", "P"}})                          //8
	conteiners = append(conteiners, Container{Crates: []string{"J", "T", "S", "R", "D"}})                //9

	//conteiners = append(conteiners, Container{Crates: []string{"N", "Z"}})
	//conteiners = append(conteiners, Container{Crates: []string{"D", "C", "M"}})
	//conteiners = append(conteiners, Container{Crates: []string{"P"}})

	inputDataRows := strings.Split(string(inputData), "\n")
	moves := []Move{}
	for _, dataRow := range inputDataRows {
		row := strings.ReplaceAll(dataRow, "move ", "")
		row = strings.ReplaceAll(row, "from ", "")
		row = strings.ReplaceAll(row, "to ", "")

		parts := strings.Split(row, " ")
		move := Move{}
		if number, err := strconv.Atoi(parts[0]); err == nil {
			move.CountToMove = number
		}
		if number, err := strconv.Atoi(parts[1]); err == nil {
			move.SourceIndex = number
		}
		if number, err := strconv.Atoi(parts[2]); err == nil {
			move.DestinationIndex = number
		}
		moves = append(moves, move)
	}

	return moves, conteiners
}

func Part1(inputData []byte) int {

	moves, containers := parseInput(inputData)

	var currCrate string
	for _, move := range moves {
		for i := 0; i < move.CountToMove; i++ {
			//queue
			currCrate, containers[move.SourceIndex-1].Crates = containers[move.SourceIndex-1].Crates[0], containers[move.SourceIndex-1].Crates[1:]
			containers[move.DestinationIndex-1].Crates = append([]string{currCrate}, containers[move.DestinationIndex-1].Crates...)
		}
	}

	fmt.Println(containers)

	return 0
}

func Part2(inputData []byte) int {

	moves, containers := parseInput(inputData)
	var currCrates []string

	for _, move := range moves {
		currCrates = []string{}
		currCrates = append(currCrates, containers[move.SourceIndex-1].Crates[0:move.CountToMove]...)
		containers[move.SourceIndex-1].Crates = containers[move.SourceIndex-1].Crates[move.CountToMove:]

		containers[move.DestinationIndex-1].Crates = append(currCrates, containers[move.DestinationIndex-1].Crates...)
	}

	res := ""
	for _, container := range containers {
		res += container.Crates[0]
	}

	fmt.Println(res)
	return 0
}

func main() {

	pwd, _ := os.Getwd()
	inputData, err := os.ReadFile(pwd + "/2022/day05/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	//fmt.Println(Part1(inputData))
	fmt.Println(Part2(inputData))
}
