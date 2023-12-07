package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Hand struct {
	Cards      string
	Bid        int
	CardsPower []int
	Type       int
	Rank       int
}

func parseInput(inputData []byte) ([]Hand, error) {

	result := []Hand{}
	inputDataRows := strings.Split(string(inputData), "\n")
	for _, dataRow := range inputDataRows {

		parts := strings.Split(dataRow, " ")
		hand := Hand{Cards: parts[0]}
		if number, err := strconv.Atoi(parts[1]); err == nil {
			hand.Bid = number
		} else {
			return result, fmt.Errorf("Bad input format")
		}
		result = append(result, hand)
	}
	return result, nil
}

func isFiveOfAKind(s string) bool {
	if strings.Count(s, string(s[0])) == 5 {
		return true
	}
	return false
}

func isFourOfAKind(s string) bool {
	if strings.Count(s, string(s[0])) == 4 || strings.Count(s, string(s[1])) == 4 {
		return true
	}
	return false
}

func isFullHouse(s string) bool {
	s1 := string(s[0])
	s2 := ""
	if s[0] != s[1] {
		s2 = string(s[1])
	} else if s[0] != s[2] {
		s2 = string(s[2])
	} else if s[0] != s[3] {
		s2 = string(s[3])
	}

	if len(s2) > 0 {
		if strings.Count(s, s1) == 2 || strings.Count(s, s1) == 3 {
			if strings.Count(s, s1) == 2 && strings.Count(s, s2) == 3 {
				return true
			}
			if strings.Count(s, s1) == 3 && strings.Count(s, s2) == 2 {
				return true
			}
		}
	}

	return false
}

func isThreeOfKind(s string) bool {

	charCount := make(map[rune]int)

	for _, char := range s {
		charCount[char]++
	}

	for _, count := range charCount {
		if count == 3 {
			return true
		}
	}

	return false
}

func isTwoPair(s string) bool {

	charCount := make(map[rune]int)

	for _, char := range s {
		charCount[char]++
	}
	pairCount := 0
	for _, count := range charCount {
		if count == 2 {
			pairCount++
		}
	}
	return pairCount == 2
}
func isOnePair(s string) bool {

	charCount := make(map[rune]int)

	for _, char := range s {
		charCount[char]++
	}
	for _, count := range charCount {
		if count == 2 {
			return true
		}
	}

	return false
}

func findHandsCardsPower(hands []Hand, version int) {
	for i := 0; i < len(hands); i++ {
		power := []int{}
		for _, letter := range hands[i].Cards {
			if number, err := strconv.Atoi(string(letter)); err == nil {
				power = append(power, number)
			} else {
				if string(letter) == "T" {
					power = append(power, 10)
				} else if string(letter) == "T" {
					power = append(power, 11)
				} else if string(letter) == "J" {
					if version == 1 {
						power = append(power, 12)
					} else {
						power = append(power, 1)
					}

				} else if string(letter) == "Q" {
					power = append(power, 13)
				} else if string(letter) == "K" {
					power = append(power, 14)
				} else if string(letter) == "A" {
					power = append(power, 15)
				}
			}
		}
		hands[i].CardsPower = power
	}
}

func getTypeFromCards(cards string) int {
	if isFiveOfAKind(cards) {
		return 7
	} else if isFourOfAKind(cards) {
		return 6
	} else if isFullHouse(cards) {
		return 5
	} else if isThreeOfKind(cards) {
		return 4
	} else if isTwoPair(cards) {
		return 3
	} else if isOnePair(cards) {
		return 2
	}

	return 1
}

func findHansType(hands []Hand) {
	for i := 0; i < len(hands); i++ {
		hands[i].Type = getTypeFromCards(hands[i].Cards)
	}
}

func findThePossibleHighestRank(cards string) int {
	res := 0
	for i := 2; i <= 9; i++ {
		index := strconv.Itoa(i)
		newCards := strings.ReplaceAll(cards, "J", index)
		newRank := getTypeFromCards(newCards)
		if newRank > res {
			res = newRank
		}
	}
	newCards := strings.ReplaceAll(cards, "J", "T")
	newRank := getTypeFromCards(newCards)
	if newRank > res {
		res = newRank
	}
	newCards = strings.ReplaceAll(cards, "J", "Q")
	newRank = getTypeFromCards(newCards)
	if newRank > res {
		res = newRank
	}
	newCards = strings.ReplaceAll(cards, "J", "K")
	newRank = getTypeFromCards(newCards)
	if newRank > res {
		res = newRank
	}
	newCards = strings.ReplaceAll(cards, "J", "A")
	newRank = getTypeFromCards(newCards)
	if newRank > res {
		res = newRank
	}
	return res
}

func findHansTypeWithJ(hands []Hand) {
	for i := 0; i < len(hands); i++ {
		if strings.Contains(hands[i].Cards, "J") {
			hands[i].Type = findThePossibleHighestRank(hands[i].Cards)
		} else {
			hands[i].Type = getTypeFromCards(hands[i].Cards)
		}
	}
}

func findHandsRank(hands []Hand) {
	rank := 1
	for i := 0; i < len(hands); i++ {
		hands[i].Rank = rank
		rank++
	}
}

func Part1(inputData []byte) int {

	hands, err := parseInput(inputData)
	if err != nil {
		fmt.Println(err)
		return 0
	}

	findHandsCardsPower(hands, 1)
	findHansType(hands)

	sort.Slice(hands, func(i, j int) bool {
		if hands[i].Type == hands[j].Type {
			for index := 0; index < len(hands[i].CardsPower); index++ {
				if hands[i].CardsPower[index] == hands[j].CardsPower[index] {
					continue
				}
				return hands[i].CardsPower[index] < hands[j].CardsPower[index]
			}
		}

		return hands[i].Type < hands[j].Type
	})
	findHandsRank(hands)

	res := 0
	for _, hand := range hands {
		res += hand.Bid * hand.Rank
	}

	//fmt.Println(hands)

	return res
}

func Part2(inputData []byte) int {
	hands, err := parseInput(inputData)
	if err != nil {
		fmt.Println(err)
		return 0
	}

	findHandsCardsPower(hands, 2)
	findHansTypeWithJ(hands)

	sort.Slice(hands, func(i, j int) bool {
		if hands[i].Type == hands[j].Type {
			for index := 0; index < len(hands[i].CardsPower); index++ {
				if hands[i].CardsPower[index] == hands[j].CardsPower[index] {
					continue
				}
				return hands[i].CardsPower[index] < hands[j].CardsPower[index]
			}
		}

		return hands[i].Type < hands[j].Type
	})
	findHandsRank(hands)

	res := 0
	for _, hand := range hands {
		res += hand.Bid * hand.Rank
	}

	//fmt.Println(hands)

	return res
}

func main() {

	pwd, _ := os.Getwd()
	inputData, err := os.ReadFile(pwd + "/2023/day07/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(Part1(inputData))
	fmt.Println(Part2(inputData))
}
