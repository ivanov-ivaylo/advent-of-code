package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type File struct {
	isFile   bool
	Name     string
	Size     int
	Parent   string
	Children map[string]File
}

func parseInput(inputData []byte) map[string]File {
	inputDataRows := strings.Split(string(inputData), "\n")
	var dataRow string
	var currentFolder string
	var prevFolder string
	var lsItem string
	result := map[string]File{}
	result["/"] = File{isFile: false, Name: "/", Children: map[string]File{}}
	currentFolder = "/"
	for len(inputDataRows) > 0 {
		dataRow, inputDataRows = inputDataRows[0], inputDataRows[1:]
		if string(dataRow[0]) == "$" {
			command := dataRow[2:4]
			if command == "cd" {
				if dataRow[4:] == ".." {
					currentFolder = result[result[currentFolder].Parent].Name
				} else {
					prevFolder = currentFolder
					currentFolder = strings.TrimSpace(dataRow[4:])
					if _, ok := result[currentFolder]; !ok {
						result[currentFolder] = File{isFile: false, Name: currentFolder, Parent: prevFolder, Children: map[string]File{}}
					}
				}
			} else if command == "ls" {
				lsItem = ""
				for len(inputDataRows) > 0 {
					lsItem, inputDataRows = inputDataRows[0], inputDataRows[1:]
					if string(lsItem[0]) == "$" {
						inputDataRows = append([]string{lsItem}, inputDataRows...)
						break
					}
					if string(lsItem[0:3]) == "dir" {
						currDir := File{isFile: false, Name: lsItem[4:], Parent: currentFolder, Children: map[string]File{}}
						result[currentFolder].Children[lsItem[4:]] = currDir
					} else {
						parts := strings.Split(lsItem, " ")
						currFile := File{isFile: true, Name: parts[1], Parent: currentFolder}
						if number, err := strconv.Atoi(parts[0]); err == nil {
							currFile.Size = number
						}
						result[currentFolder].Children[parts[1]] = currFile
					}
				}

			}
		}
	}
	return result
}

func Part1(inputData []byte) int {

	tree := parseInput(inputData)

	fmt.Println(tree)
	return 0
}

func Part2(inputData []byte) int {
	return 0
}

func main() {

	pwd, _ := os.Getwd()
	inputData, err := os.ReadFile(pwd + "/2022/day07/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(Part1(inputData))
	//fmt.Println(Part2(inputData))
}
