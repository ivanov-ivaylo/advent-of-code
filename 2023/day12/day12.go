package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Spring struct {
	Numbers      []int
	UnknownCount int
	Row          []string
}

func parseInput(inputData []byte) []Spring {
	inputDataRows := strings.Split(string(inputData), "\n")

	result := []Spring{}
	for _, dataRow := range inputDataRows {
		row := strings.Split(dataRow, " ")
		spring := Spring{}
		for _, letter := range row[0] {
			if string(letter) == "?" {
				spring.UnknownCount++
			}
			spring.Row = append(spring.Row, string(letter))
		}
		parts := strings.Split(row[1], ",")
		for _, num := range parts {
			if number, err := strconv.Atoi(num); err == nil {
				spring.Numbers = append(spring.Numbers, number)
			}
		}
		result = append(result, spring)
	}

	return result
}

func getAllPossibleUnknownVariants(N int) []string {
	result := []string{}
	generateBinaryNumbers("", N, &result)
	return result
}

func generateBinaryNumbers(prefix string, remainingBits int, result *[]string) {
	if remainingBits == 0 {
		*result = append(*result, prefix)
		return
	}

	generateBinaryNumbers(prefix+"0", remainingBits-1, result)
	generateBinaryNumbers(prefix+"1", remainingBits-1, result)
}

func isVariantValid(spring Spring, variantUnknown string) bool {

	row := []string{}
	row = append([]string(nil), spring.Row...)
	numbers := []int{}
	numbers = append([]int(nil), spring.Numbers...)

	for _, letter := range variantUnknown {
		for i := 0; i < len(row); i++ {
			if row[i] == "?" {
				if string(letter) == "0" {
					row[i] = "."
				} else if string(letter) == "1" {
					row[i] = "#"
				}
				break
			}
		}
	}

	rowParts := strings.Split(strings.Join(row, ""), ".")
	var num int
	for _, item := range rowParts {
		if item == "" {
			continue
		}
		if len(numbers) == 0 {
			return false
		}
		num, numbers = numbers[0], numbers[1:]
		if len(item) != num {
			return false
		}
	}
	if len(numbers) > 0 {
		return false
	}

	return true
}

func unfoldRecord(record string) string {
	var res strings.Builder
	for i := 0; i < len(record)*5; i++ {
		if i != 0 && i%len(record) == 0 {
			res.WriteByte('?')
		}
		res.WriteByte(record[i%len(record)])
	}

	return res.String()
}

func unfoldGroup(group []int) []int {
	var res []int
	for i := 0; i < len(group)*5; i++ {
		res = append(res, group[i%len(group)])
	}

	return res
}

func solve(record string, group []int) int {
	var cache [][]int
	for i := 0; i < len(record); i++ {
		cache = append(cache, make([]int, len(group)+1))
		for j := 0; j < len(group)+1; j++ {
			cache[i][j] = -1
		}
	}

	return dp(0, 0, record, group, cache)
}

func dp(i, j int, record string, group []int, cache [][]int) int {
	if i >= len(record) {
		if j < len(group) {
			return 0
		}
		return 1
	}

	if cache[i][j] != -1 {
		return cache[i][j]
	}

	res := 0
	if record[i] == '.' {
		res = dp(i+1, j, record, group, cache)
	} else {
		if record[i] == '?' {
			res += dp(i+1, j, record, group, cache)
		}
		if j < len(group) {
			count := 0
			for k := i; k < len(record); k++ {
				if count > group[j] || record[k] == '.' || count == group[j] && record[k] == '?' {
					break
				}
				count += 1
			}

			if count == group[j] {
				if i+count < len(record) && record[i+count] != '#' {
					res += dp(i+count+1, j+1, record, group, cache)
				} else {
					res += dp(i+count, j+1, record, group, cache)
				}
			}
		}
	}

	cache[i][j] = res
	return res
}

func parse(input string) ([]string, [][]int) {
	var records []string
	var groups [][]int

	for _, line := range strings.Split(strings.ReplaceAll(input, "\r\n", "\n"), "\n") {
		parts := strings.Split(line, " ")
		records = append(records, parts[0])
		var group []int
		for _, num := range strings.Split(parts[1], ",") {
			num, _ := strconv.Atoi(num)
			group = append(group, num)
		}
		groups = append(groups, group)
	}

	return records, groups
}

func Part1(inputData []byte) int {

	springs := parseInput(inputData)

	res := 0
	for _, spring := range springs {
		arrangements := getAllPossibleUnknownVariants(spring.UnknownCount)
		currentRes := 0
		for _, variant := range arrangements {
			if isVariantValid(spring, variant) {
				currentRes++
			}
		}
		res += currentRes
	}

	//fmt.Println(springs)

	return res
}

func Part2(inputData []byte) int {
	records, groups := parse(string(inputData))
	res := 0
	for i := range records {
		res += solve(unfoldRecord(records[i]), unfoldGroup(groups[i]))
	}

	return res
}

func main() {

	pwd, _ := os.Getwd()
	inputData, err := os.ReadFile(pwd + "/2023/day12/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(Part1(inputData))
	fmt.Println(Part2(inputData))
}
