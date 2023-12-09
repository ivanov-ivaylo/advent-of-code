package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type HistorySeq struct {
	Numbers    [][]int
	NextNumber int
	PrevNumber int
}

func parseInput(inputData []byte) []HistorySeq {

	inputDataRows := strings.Split(string(inputData), "\n")
	result := []HistorySeq{}
	for _, dataRow := range inputDataRows {
		his := HistorySeq{}
		numbers := strings.Split(dataRow, " ")
		tmp := []int{}
		for _, num := range numbers {
			if number, err := strconv.Atoi(num); err == nil {
				tmp = append(tmp, number)
			}
		}
		his.Numbers = append(his.Numbers, tmp)
		result = append(result, his)
	}
	return result
}

func isAllZero(list []int) bool {
	for _, l := range list {
		if l != 0 {
			return false
		}
	}
	return true
}

func findExtrapolated(seqs []HistorySeq) {

	for i := 0; i < len(seqs); i++ {
		for !isAllZero(seqs[i].Numbers[len(seqs[i].Numbers)-1]) {
			last := seqs[i].Numbers[len(seqs[i].Numbers)-1]
			tmp := []int{}
			for j := 0; j < len(last)-1; j++ {
				tmp = append(tmp, last[j+1]-last[j])
			}
			seqs[i].Numbers = append(seqs[i].Numbers, tmp)
		}

		currNext := 0
		for j := len(seqs[i].Numbers) - 2; j >= 0; j-- {
			currNext += seqs[i].Numbers[j][len(seqs[i].Numbers[j])-1]
		}
		seqs[i].NextNumber = currNext

		currPrev := 0
		for j := len(seqs[i].Numbers) - 2; j >= 0; j-- {
			currPrev = seqs[i].Numbers[j][0] - currPrev
		}
		seqs[i].PrevNumber = currPrev
	}
}

func Part1(inputData []byte) int {

	seqs := parseInput(inputData)

	findExtrapolated(seqs)

	res := 0
	for _, seq := range seqs {
		res += seq.NextNumber
	}

	//fmt.Println(seqs)

	return res
}

func Part2(inputData []byte) int {

	seqs := parseInput(inputData)

	findExtrapolated(seqs)

	res := 0
	for _, seq := range seqs {
		res += seq.PrevNumber
	}

	return res
}

func main() {

	pwd, _ := os.Getwd()
	inputData, err := os.ReadFile(pwd + "/2023/day09/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(Part1(inputData))
	fmt.Println(Part2(inputData))
}
