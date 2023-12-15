package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type HashSeq struct {
	Seq    string
	Number int
}

type BoxItem struct {
	Word   string
	LenNum int
}

type Box struct {
	Items []BoxItem
	Index int
}

func parseInput(inputData []byte) []HashSeq {
	inputDataRows := strings.Split(string(inputData), ",")

	result := []HashSeq{}
	for _, hash := range inputDataRows {
		result = append(result, HashSeq{Seq: hash})
	}

	return result

}

func findHashForString(str string) int {
	currentValue := 0
	for j := 0; j < len(str); j++ {
		currentValue += int(str[j])
		currentValue = currentValue * 17
		currentValue = currentValue % 256
	}

	return currentValue
}

func findHashNumber(seqs []HashSeq) {

	for i := 0; i < len(seqs); i++ {
		seqs[i].Number = findHashForString(seqs[i].Seq)
	}
}

func initBoxes() []Box {
	result := []Box{}

	for i := 0; i < 256; i++ {
		result = append(result, Box{Index: i})
	}
	return result
}

func addToBoxes(box Box, hash string, lenNum int) {
	box.Items = append(box.Items, BoxItem{Word: hash, LenNum: lenNum})
}

func removeFromBoxes(box Box, hash string) {

	for j := 0; j < len(box.Items); j++ {
		if box.Items[j].Word == hash {
			box.Items = append(box.Items[:j], box.Items[j+1:]...)
			break
		}
	}
}

func Part1(inputData []byte) int {

	seqs := parseInput(inputData)
	findHashNumber(seqs)
	res := 0
	for i := 0; i < len(seqs); i++ {
		res += seqs[i].Number
	}

	return res
}

func calculateFocusPower(boxes []Box) int {
	focusPower := 0
	for i := 0; i < len(boxes); i++ {
		for j := 0; j < len(boxes[i].Items); j++ {
			curr := (i + 1) * (j + 1) * boxes[i].Items[j].LenNum
			focusPower += curr
		}
	}

	return focusPower
}

func Part2(inputData []byte) int {

	seqs := parseInput(inputData)
	boxes := initBoxes()

	for _, seq := range seqs {
		if strings.Contains(seq.Seq, "-") {
			partsDash := strings.Split(seq.Seq, "-")
			boxIndex := findHashForString(partsDash[0])
			for j := 0; j < len(boxes[boxIndex].Items); j++ {
				if boxes[boxIndex].Items[j].Word == partsDash[0] {
					boxes[boxIndex].Items = append(boxes[boxIndex].Items[:j], boxes[boxIndex].Items[j+1:]...)
					break
				}
			}
		} else if strings.Contains(seq.Seq, "=") {
			partsDash := strings.Split(seq.Seq, "=")
			boxIndex := findHashForString(partsDash[0])
			found := false
			for i := 0; i < len(boxes[boxIndex].Items); i++ {
				if boxes[boxIndex].Items[i].Word == partsDash[0] {
					found = true
					if number, err := strconv.Atoi(partsDash[1]); err == nil {
						boxes[boxIndex].Items[i].LenNum = number
					}
					break
				}
			}
			if !found {
				if number, err := strconv.Atoi(partsDash[1]); err == nil {
					boxes[boxIndex].Items = append(boxes[boxIndex].Items, BoxItem{Word: partsDash[0], LenNum: number})
				}
			}

		}
	}

	return calculateFocusPower(boxes)
}

func main() {

	pwd, _ := os.Getwd()
	inputData, err := os.ReadFile(pwd + "/2023/day15/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	//fmt.Println(Part1(inputData))
	fmt.Println(Part2(inputData))
}
