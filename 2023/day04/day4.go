package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Card struct {
	Index            int
	WinningNumbers   []int
	SuggestedNumbers []int
	Matches          int
}

func parseInput(inputData []byte) ([]Card, error) {
	var result []Card
	inputDataRows := strings.Split(string(inputData), "\n")
	for _, dataRow := range inputDataRows {
		dataRowParts := strings.Split(dataRow, ":")
		if len(dataRowParts) != 2 {
			return result, fmt.Errorf("Invalid inputs when split by :")
		}
		labelPart := dataRowParts[0]
		bodyPart := dataRowParts[1]
		labelPart = strings.TrimSpace(strings.ReplaceAll(labelPart, "Card ", ""))
		cardIndex, err := strconv.Atoi(labelPart)
		if err != nil {
			return result, err
		}
		currentCard := Card{Index: cardIndex}

		bodyParts := strings.Split(bodyPart, " | ")
		if len(bodyParts) != 2 {
			return result, fmt.Errorf("Invalid inputs too many parts")
		}

		winningNumbersPart := strings.Split(strings.TrimSpace(bodyParts[0]), " ")
		suggesterNumbersPart := strings.Split(strings.TrimSpace(bodyParts[1]), " ")

		for _, winNumber := range winningNumbersPart {
			if winNumber == "" {
				continue
			}
			winNum, err := strconv.Atoi(winNumber)
			if err != nil {
				return result, err
			}
			currentCard.WinningNumbers = append(currentCard.WinningNumbers, winNum)
		}
		for _, suggestNumber := range suggesterNumbersPart {
			if suggestNumber == "" {
				continue
			}
			suggestNum, err := strconv.Atoi(suggestNumber)
			if err != nil {
				return result, err
			}
			currentCard.SuggestedNumbers = append(currentCard.SuggestedNumbers, suggestNum)
		}

		result = append(result, currentCard)
	}
	return result, nil
}

func hasElement(element int, list []int) bool {
	for _, item := range list {
		if item == element {
			return true
		}
	}
	return false
}

func Part1(inputData []byte) int {

	cards, err := parseInput(inputData)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	result := 0
	for _, card := range cards {
		currentPrice := 0
		for _, suggestedNumber := range card.SuggestedNumbers {
			if hasElement(suggestedNumber, card.WinningNumbers) {
				if currentPrice == 0 {
					currentPrice = 1
				} else {
					currentPrice *= 2
				}
			}
		}
		result += currentPrice
	}

	return result
}

func Part2(inputData []byte) int {

	cards, err := parseInput(inputData)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	for i := 0; i < len(cards); i++ {
		for _, suggestedNumber := range cards[i].SuggestedNumbers {
			if hasElement(suggestedNumber, cards[i].WinningNumbers) {
				cards[i].Matches++
			}
		}
	}
	//use this map for faster access by index
	sourceMap := map[int]Card{}
	for _, card := range cards {
		sourceMap[card.Index] = card
	}
	var currCard Card
	result := 0
	for len(cards) > 0 {
		result++
		currCard, cards = cards[0], cards[1:]

		for i := currCard.Index; i < currCard.Index+currCard.Matches; i++ {
			if val, ok := sourceMap[i+1]; ok {
				copyElem := Card{Index: val.Index, WinningNumbers: val.WinningNumbers, SuggestedNumbers: val.SuggestedNumbers, Matches: val.Matches}
				cards = append(cards, copyElem)
			}
		}
	}

	return result

}

func main() {

	pwd, _ := os.Getwd()
	inputData, err := os.ReadFile(pwd + "/2023/day04/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	//fmt.Println(Part1(inputData))
	fmt.Println(Part2(inputData))
}
