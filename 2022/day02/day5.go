package main

import (
	"fmt"
	"os"
)

type Duel struct {
	OpponentChoose string
	MyChoose       string
	A_point        int
	B_point        int
	C_point        int
}

//func parseInput(inputData []byte) []Duel {
//	inputDataRows := strings.Split(string(inputData), "\n")
//
//	result := []Duel{}
//	for _, dataRow := range inputDataRows {
//		parts := strings.Split(dataRow, " ")
//		//item := Duel{OpponentChoose: parts[0], My}
//	}
//
//	return result
//}

func Part1(inputData []byte) int {
	return 0
}

func Part2(inputData []byte) int {
	return 0
}

func main() {

	pwd, _ := os.Getwd()
	inputData, err := os.ReadFile(pwd + "/2022/day02/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	//fmt.Println(Part1(inputData))
	fmt.Println(Part2(inputData))
}
