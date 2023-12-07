package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ElfSections struct {
	StartIndex       int
	EndIndex         int
	NumberOfSections int
}

type ElfPair struct {
	First           ElfSections
	Second          ElfSections
	FullyContains   bool
	PartialContains bool
}

func parseInput(inputData []byte) ([]ElfPair, error) {
	inputDataRows := strings.Split(string(inputData), "\n")
	result := []ElfPair{}

	for _, dataRow := range inputDataRows {
		elfs := strings.Split(dataRow, ",")
		if len(elfs) != 2 {
			return result, fmt.Errorf("Wrong input")
		}
		elf1 := strings.Split(elfs[0], "-")
		elf2 := strings.Split(elfs[1], "-")
		first := ElfSections{}
		if number, err := strconv.Atoi(elf1[0]); err == nil {
			first.StartIndex = number
		}
		if number2, err := strconv.Atoi(elf1[1]); err == nil {
			first.EndIndex = number2
		}
		if first.StartIndex > first.EndIndex {
			return result, fmt.Errorf("Wrong input")
		}
		first.NumberOfSections = first.EndIndex - first.StartIndex + 1
		second := ElfSections{}
		if number, err := strconv.Atoi(elf2[0]); err == nil {
			second.StartIndex = number
		}
		if number, err := strconv.Atoi(elf2[1]); err == nil {
			second.EndIndex = number
		}
		second.NumberOfSections = second.EndIndex - second.StartIndex + 1

		result = append(result, ElfPair{First: first, Second: second})
	}
	return result, nil
}

func setFullyContains(elfPairs []ElfPair) int {

	res := 0
	for i := 0; i < len(elfPairs); i++ {
		first := elfPairs[i].First
		second := elfPairs[i].Second
		if first.StartIndex <= second.StartIndex && second.EndIndex <= first.EndIndex {
			elfPairs[i].FullyContains = true
			res++
		} else if second.StartIndex <= first.StartIndex && first.EndIndex <= second.EndIndex {
			elfPairs[i].FullyContains = true
			res++
		}
	}
	return res
}

func setPartialContains(elfPairs []ElfPair) int {
	res := 0
	for i := 0; i < len(elfPairs); i++ {
		first := elfPairs[i].First
		second := elfPairs[i].Second

		if first.StartIndex <= second.EndIndex && second.StartIndex <= first.EndIndex {
			elfPairs[i].PartialContains = true
			res++
		}
	}
	return res
}

func Part1(inputData []byte) int {

	elfPairs, err := parseInput(inputData)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	res := setFullyContains(elfPairs)
	fmt.Println(elfPairs)
	return res
}

func Part2(inputData []byte) int {

	elfPairs, err := parseInput(inputData)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	res := setPartialContains(elfPairs)
	return res
}

func main() {

	pwd, _ := os.Getwd()
	inputData, err := os.ReadFile(pwd + "/2022/day04/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(Part1(inputData))
	fmt.Println(Part2(inputData))
}
