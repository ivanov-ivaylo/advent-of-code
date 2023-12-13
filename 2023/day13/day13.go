package main

import (
	"fmt"
	"os"
	"strings"
)

type Reflection struct {
	Field        [][]int
	NumberOfRows int
	NumberOfCols int
}

func parseInput(inputData []byte) []Reflection {
	inputDataRows := strings.Split(string(inputData), "\n\n")
	result := []Reflection{}

	for _, pattern := range inputDataRows {
		gridRows := strings.Split(pattern, "\n")
		refl := Reflection{}
		for _, row := range gridRows {
			fieldRow := []int{}
			for _, charItem := range row {
				if string(charItem) == "." {
					fieldRow = append(fieldRow, 0)
				} else {
					fieldRow = append(fieldRow, 1)
				}
			}
			refl.Field = append(refl.Field, fieldRow)
		}
		result = append(result, refl)
	}
	return result
}

func findReflectionNumbers(reflectionCol *Reflection, reflectionRow *Reflection) {

	field := reflectionRow.Field
	for x := 0; x < len(field)-1; x++ {
		hasReflection := true
		for y := 0; y < len(field[0]); y++ {
			if field[x][y] != field[x+1][y] {
				hasReflection = false
				break
			}
		}
		if hasReflection {
			reflectionRow.NumberOfRows = x + 1
			break
		}
	}

	field = reflectionCol.Field
	for y := 0; y < len(field[0])-1; y++ {
		hasReflection := true
		for x := 0; x < len(field); x++ {
			if field[x][y] != field[x][y+1] {
				hasReflection = false
				break
			}
		}
		if hasReflection {
			reflectionCol.NumberOfCols = y + 1
			break
		}
	}

}

func Part1(inputData []byte) int {

	reflections := parseInput(inputData)

	res := 0
	for i := 0; i < len(reflections); i = i + 2 {
		findReflectionNumbers(&reflections[i], &reflections[i+1])
		res += reflections[i+1].NumberOfRows*100 + reflections[i].NumberOfCols
	}

	fmt.Println(reflections)

	return res
}

func Part2(inputData []byte) int {
	return 0
}

func main() {

	pwd, _ := os.Getwd()
	inputData, err := os.ReadFile(pwd + "/2023/day13/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(Part1(inputData))
	//fmt.Println(Part2(inputData))
}
