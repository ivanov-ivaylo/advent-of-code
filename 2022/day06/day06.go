package main

import (
	"fmt"
	"os"
)

func isAllDifferent4(x, y, z, t string) bool {
	if x != y && x != z && x != t && y != z && y != t && z != t {
		return true
	}
	return false
}

func allCharsDifferent(s string) bool {
	charSet := make(map[rune]bool)

	for _, char := range s {
		if charSet[char] {
			return false // If the character is already present, not all characters are different
		}
		charSet[char] = true
	}

	return true // All characters are different
}

func Part1(inputData []byte) int {

	input := string(inputData)

	for i := 0; i < len(input); i++ {
		if isAllDifferent4(string(input[i]), string(input[i+1]), string(input[i+2]), string(input[i+3])) {
			return i + 4
		}
	}

	return 0
}

func Part2(inputData []byte) int {

	input := string(inputData)

	for i := 0; i < len(input); i++ {
		subString := input[i : i+14]
		if allCharsDifferent(subString) {
			return i + 14
		}
	}

	return 0
}

func main() {

	pwd, _ := os.Getwd()
	inputData, err := os.ReadFile(pwd + "/2022/day06/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(Part1(inputData))
	fmt.Println(Part2(inputData))
}
