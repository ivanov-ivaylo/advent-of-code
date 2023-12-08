package main

import (
	"fmt"
	"os"
	"strings"
)

type Moves struct {
	Move []string
}

type Doc struct {
	Name  string
	Left  string
	Right string
}

func parseInput(inputData []byte) (map[string]Doc, Moves) {
	result := map[string]Doc{}
	resMoves := Moves{}
	inputDataRows := strings.Split(string(inputData), "\n\n")

	for _, move := range inputDataRows[0] {
		resMoves.Move = append(resMoves.Move, string(move))
	}

	dataRow2 := strings.Split(inputDataRows[1], "\n")

	for _, currDoc := range dataRow2 {
		row := strings.Split(currDoc, " = ")
		doc := Doc{Name: row[0]}
		parts := strings.Split(row[1], ", ")

		doc.Left = parts[0][1:]
		doc.Right = strings.TrimSuffix(parts[1], ")")

		result[row[0]] = doc
	}
	return result, resMoves
}

func Part1(inputData []byte) int {

	docs, moves := parseInput(inputData)

	currDoc := "AAA"
	steps := 0
	for currDoc != "ZZZ" {
		for _, move := range moves.Move {
			steps++
			if string(move) == "L" {
				currDoc = docs[currDoc].Left
			} else {
				currDoc = docs[currDoc].Right
			}
			if currDoc == "ZZZ" {
				break
			}
		}
	}

	//fmt.Println(currDoc)
	//fmt.Println(docs, moves)

	return steps
}

func findADocs(docs map[string]Doc) []string {

	starts := []string{}
	for _, currDoc := range docs {
		if string(currDoc.Name[len(currDoc.Name)-1]) == "A" {
			starts = append(starts, currDoc.Name)
		}
	}

	return starts
}

func isEndReached(currDocs []string) bool {

	for _, curr := range currDocs {
		if string(curr[len(curr)-1]) != "Z" {
			return false
		}
	}

	return true
}

func getLeastCommonMultiple(numbers []int) int {
	lcm := numbers[0]
	for i := 0; i < len(numbers); i++ {
		num1 := lcm
		num2 := numbers[i]
		gcd := 1
		for num2 != 0 {
			temp := num2
			num2 = num1 % num2
			num1 = temp
		}
		gcd = num1
		lcm = (lcm * numbers[i]) / gcd
	}

	return lcm
}

func Part2(inputData []byte) int {

	docs, moves := parseInput(inputData)

	currStarts := findADocs(docs)

	nums := []int{}
	for _, currStart := range currStarts {

		currDocs := []string{currStart}
		steps := 0
		for !isEndReached(currDocs) {
			for _, move := range moves.Move {
				steps++
				for i := 0; i < len(currDocs); i++ {
					if move == "L" {
						currDocs[i] = docs[currDocs[i]].Left
					} else {
						currDocs[i] = docs[currDocs[i]].Right
					}
				}
				//fmt.Println(currDocs)
				if isEndReached(currDocs) {
					break
				}
			}
		}

		nums = append(nums, steps)
	}

	lcm := getLeastCommonMultiple(nums)

	return lcm
}

func main() {

	pwd, _ := os.Getwd()
	inputData, err := os.ReadFile(pwd + "/2023/day08/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(Part1(inputData))
	fmt.Println(Part2(inputData))
}
