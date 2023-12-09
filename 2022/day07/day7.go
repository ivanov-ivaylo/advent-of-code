package main

import (
	"fmt"
	"os"
)

type File struct {
	isFile   bool
	Name     string
	Size     int
	Children map[string]File
}

//func parseInput(inputData []byte) map[string] {
//
//}

func Part1(inputData []byte) int {
	return 0
}

func Part2(inputData []byte) int {
	return 0
}

func main() {

	pwd, _ := os.Getwd()
	inputData, err := os.ReadFile(pwd + "/2022/day07/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(Part1(inputData))
	//fmt.Println(Part2(inputData))
}
