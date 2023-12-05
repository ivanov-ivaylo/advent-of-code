package main

import (
	"fmt"
	"os"
	"strings"
)

type Duel struct {
	OpponentChoose string
	MyChoose       string
	Points         int
}

func parseInput(inputData []byte) []Duel {
	inputDataRows := strings.Split(string(inputData), "\n")

	result := []Duel{}
	for _, dataRow := range inputDataRows {
		parts := strings.Split(dataRow, " ")
		item := Duel{OpponentChoose: parts[0], MyChoose: parts[1]}
		result = append(result, item)
	}

	return result
}

func getPoints(opponent string, my string) int {
	points := 0
	if my == "A" {
		points += 1
	} else if my == "B" {
		points += 2
	} else {
		points += 3
	}

	if opponent == my {
		points += 3
	} else {
		if (opponent == "A" && my == "B") || (opponent == "B" && my == "C") || (opponent == "C" && my == "A") {
			points += 6
		}
	}

	return points
}

func CalculateDuelScore(duels []Duel, X string, Y string, Z string) int {

	total := 0
	for i := 0; i < len(duels); i++ {
		myChoose := "A"
		if duels[i].MyChoose == "X" {
			myChoose = X
		}
		if duels[i].MyChoose == "Y" {
			myChoose = Y
		}
		if duels[i].MyChoose == "Z" {
			myChoose = Z
		}
		duels[i].Points = getPoints(duels[i].OpponentChoose, myChoose)
		total += duels[i].Points
	}

	return total
}

func CalculateDuelScore2(duels []Duel) int {

	res := 0
	for _, duel := range duels {
		if duel.MyChoose == "X" { //lose
			if duel.OpponentChoose == "A" {
				res += getPoints(duel.OpponentChoose, "C")
			} else if duel.OpponentChoose == "B" {
				res += getPoints(duel.OpponentChoose, "A")
			} else {
				res += getPoints(duel.OpponentChoose, "B")
			}
		} else if duel.MyChoose == "Y" { // draw
			res += getPoints(duel.OpponentChoose, duel.OpponentChoose)
		} else { // win
			if duel.OpponentChoose == "A" {
				res += getPoints(duel.OpponentChoose, "B")
			} else if duel.OpponentChoose == "B" {
				res += getPoints(duel.OpponentChoose, "C")
			} else {
				res += getPoints(duel.OpponentChoose, "A")
			}
		}
	}
	return res
}

func Part1(inputData []byte) int {

	duels := parseInput(inputData)

	result := 0
	//X, Y, Z -> A, B, C
	total1 := CalculateDuelScore(duels, "A", "B", "C")
	if result < total1 {
		result = total1
	}
	total2 := CalculateDuelScore(duels, "A", "C", "B")
	if result < total2 {
		result = total2
	}
	total3 := CalculateDuelScore(duels, "B", "A", "C")
	if result < total3 {
		result = total3
	}
	total4 := CalculateDuelScore(duels, "B", "C", "A")
	if result < total4 {
		result = total4
	}
	total5 := CalculateDuelScore(duels, "C", "A", "B")
	if result < total5 {
		result = total5
	}
	total6 := CalculateDuelScore(duels, "C", "B", "A")
	if result < total6 {
		result = total6
	}
	fmt.Println(total1, total2, total3, total4, total5, total6)
	return 0
}

func Part2(inputData []byte) int {

	duels := parseInput(inputData)

	return CalculateDuelScore2(duels)

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
