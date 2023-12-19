package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Part struct {
	X     int
	M     int
	A     int
	S     int
	Total int
}

type Rule struct {
	isValid  bool
	Letter   string
	Number   int
	Type     string //GT, LT
	ExitName string //workflow name
}

type Workflow struct {
	Name  string
	Rules []Rule
}

func parseInput(inputData []byte) (map[string]Workflow, []Part) {
	inputDataRows := strings.Split(string(inputData), "\n\n")
	workflowParts := strings.Split(inputDataRows[0], "\n")
	parts := strings.Split(inputDataRows[1], "\n")

	resultWf := map[string]Workflow{}
	resultParts := []Part{}
	for _, workflow := range workflowParts {
		wfParts := strings.Split(workflow, "{")
		wfBody := strings.Trim(wfParts[1], "}")
		wfBodyParts := strings.Split(wfBody, ",")
		currWorkflow := Workflow{Name: wfParts[0]}
		for _, wfBodyItem := range wfBodyParts {
			if strings.Contains(wfBodyItem, ":") {
				tmpPrt := strings.Split(wfBodyItem, ":")
				var rulePrt []string
				var typeRule string
				if strings.Contains(tmpPrt[0], "<") {
					rulePrt = strings.Split(tmpPrt[0], "<")
					typeRule = "LT"
				} else if strings.Contains(tmpPrt[0], ">") {
					rulePrt = strings.Split(tmpPrt[0], ">")
					typeRule = "GT"
				}
				var numPtr int
				if number, err := strconv.Atoi(rulePrt[1]); err == nil {
					numPtr = number
				}
				currWorkflow.Rules = append(currWorkflow.Rules, Rule{isValid: true, Letter: rulePrt[0], Number: numPtr, Type: typeRule, ExitName: tmpPrt[1]})

			} else {
				currWorkflow.Rules = append(currWorkflow.Rules, Rule{isValid: true, ExitName: wfBodyItem})
			}
		}
		resultWf[currWorkflow.Name] = currWorkflow
	}

	for _, part := range parts {
		currPart := Part{}
		trimPart := strings.Trim(part, "{}")
		ruleList := strings.Split(trimPart, ",")
		for _, ruleItem := range ruleList {
			ruleItemPrt := strings.Split(ruleItem, "=")
			var ruleItemNum int
			if number, err := strconv.Atoi(ruleItemPrt[1]); err == nil {
				ruleItemNum = number
				currPart.Total += number
			}
			if ruleItemPrt[0] == "x" {
				currPart.X = ruleItemNum
			} else if ruleItemPrt[0] == "m" {
				currPart.M = ruleItemNum
			} else if ruleItemPrt[0] == "a" {
				currPart.A = ruleItemNum
			} else if ruleItemPrt[0] == "s" {
				currPart.S = ruleItemNum
			}
		}
		resultParts = append(resultParts, currPart)
	}

	return resultWf, resultParts
}

func isPartAccepted(workflow map[string]Workflow, part Part) bool {

	queue := []Workflow{workflow["in"]}
	var currWf Workflow
	for len(queue) > 0 {

		currWf, queue = queue[0], queue[1:]

		for _, rule := range currWf.Rules {
			partNum := -1
			if rule.Letter == "x" {
				partNum = part.X
			} else if rule.Letter == "m" {
				partNum = part.M
			} else if rule.Letter == "a" {
				partNum = part.A
			} else if rule.Letter == "s" {
				partNum = part.S
			}
			matchRule := false
			if rule.Type == "LT" {
				if partNum < rule.Number {
					matchRule = true
				}
			} else if rule.Type == "GT" {
				if partNum > rule.Number {
					matchRule = true
				}
			} else {
				matchRule = true
			}
			if matchRule {
				if rule.ExitName == "R" {
					return false
				} else if rule.ExitName == "A" {
					return true
				} else {
					queue = append(queue, workflow[rule.ExitName])
					break
				}
			}

		}
	}

	return false
}

func Part1(inputData []byte) int {

	workflows, parts := parseInput(inputData)

	res := 0
	for _, part := range parts {
		if isPartAccepted(workflows, part) {
			res += part.Total
		}
	}

	return res
}

func Part2(inputData []byte) int {

	workflows, _ := parseInput(inputData)

	res := 0
	for x := 1; x <= 4000; x++ {
		for m := 1; m <= 4000; m++ {
			for a := 1; a <= 4000; a++ {
				for s := 1; s <= 4000; s++ {
					part := Part{X: x, M: m, A: a, S: s}
					if isPartAccepted(workflows, part) {
						res++
					}
				}
			}
		}

	}

	return res
}

func main() {

	pwd, _ := os.Getwd()
	inputData, err := os.ReadFile(pwd + "/2023/day19/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	//fmt.Println(Part1(inputData))
	fmt.Println(Part2(inputData))
}
