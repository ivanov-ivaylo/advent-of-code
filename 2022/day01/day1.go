package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type ElfBag struct {
	TotalCal int
	BagItems []int
}

func parseInput(inputData []byte) ([]ElfBag, error) {

	inputDataRows := strings.Split(string(inputData), "\n\n")
	result := []ElfBag{}
	for _, dataRow := range inputDataRows {
		elfData := strings.Split(dataRow, "\n")
		elfBag := ElfBag{TotalCal: 0}
		for _, item := range elfData {
			if item == "" {
				continue
			}
			num, err := strconv.Atoi(item)
			if err != nil {
				fmt.Println(err)
				continue
			}
			elfBag.TotalCal += num
			elfBag.BagItems = append(elfBag.BagItems, num)
		}
		result = append(result, elfBag)
	}

	return result, nil
}

func Part1(inputData []byte) int {
	bags, _ := parseInput(inputData)

	result := 0
	for _, bag := range bags {
		if result < bag.TotalCal {
			result = bag.TotalCal
		}
	}

	return result
}

func Part2(inputData []byte) int {
	bags, _ := parseInput(inputData)

	result := 0

	sort.Slice(bags, func(i, j int) bool {
		return bags[i].TotalCal > bags[j].TotalCal
	})

	result += bags[0].TotalCal + bags[1].TotalCal + bags[2].TotalCal

	return result
}

func main() {

	pwd, _ := os.Getwd()
	inputData, err := os.ReadFile(pwd + "/2022/day01/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	//fmt.Println(Part1(inputData))
	fmt.Println(Part2(inputData))
}
