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
	Parent   *File
	Children map[string]File
}

func parseInput(inputData []byte) (File, error) {
	inputDataRows := strings.Split(string(inputData), "\n")
	var dataRow string
	var lsItem string
	var currentFolder *File
	result := File{isFile: false, Name: "/", Children: map[string]File{}}
	currentFolder = &result
	for len(inputDataRows) > 0 {
		dataRow, inputDataRows = inputDataRows[0], inputDataRows[1:]
		if string(dataRow[0]) == "$" {
			command := dataRow[2:4]
			currentName := strings.TrimSpace(dataRow[4:])
			if command == "cd" {
				if currentName == ".." {
					currentFolder = currentFolder.Parent
				} else if currentName == "/" {
					continue
				} else {
					if cFile, ok := currentFolder.Children[currentName]; ok {
						currentFolder = &cFile
					} else {
						return result, fmt.Errorf("Unknown folder")
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
						currentFolder.Children[lsItem[4:]] = currDir
					} else {
						parts := strings.Split(lsItem, " ")
						currFile := File{isFile: true, Name: parts[1], Parent: currentFolder}
						if number, err := strconv.Atoi(parts[0]); err == nil {
							currFile.Size = number
						}
						currentFolder.Children[currFile.Name] = currFile
					}
				}
			}
		}
	}
	return result, nil
}

func findFolderSizes(file *File) int {

	if file.isFile {
		return file.Size
	}
	size := 0
	for childFileKey := range file.Children {
		childFile := file.Children[childFileKey]

		if childFile.isFile {
			size += childFile.Size
		} else {
			folderSize := findFolderSizes(&childFile)
			if entry, ok := file.Children[childFileKey]; ok {
				entry.Size = folderSize
				file.Children[childFileKey] = entry
			}
			size += folderSize
		}
	}
	file.Size = size

	return size
}

func Part1(inputData []byte) int {

	tree, err := parseInput(inputData)
	if err != nil {
		fmt.Println(err)
		return 0
	}

	findFolderSizes(&tree)

	//queue
	queue := []File{}
	queue = append(queue, tree)
	atMostSize := 100000
	res := 0
	var file File
	for len(queue) > 0 {
		file, queue = queue[0], queue[1:]
		if !file.isFile && file.Size < atMostSize {
			res += file.Size
		}
		for _, items := range file.Children {
			if !items.isFile {
				queue = append(queue, items)
			}
		}

	}

	return res
}

func Part2(inputData []byte) int {
	tree, err := parseInput(inputData)
	if err != nil {
		fmt.Println(err)
		return 0
	}

	findFolderSizes(&tree)

	spaceToClean := 30000000 - (70000000 - tree.Size)
	if spaceToClean < 0 {
		fmt.Println("Error to clean negative")
		return 0
	}

	//queue
	queue := []File{}
	queue = append(queue, tree)
	res := 0
	delta := spaceToClean
	var file File
	for len(queue) > 0 {
		file, queue = queue[0], queue[1:]
		if !file.isFile && file.Size > spaceToClean {
			if file.Size-spaceToClean < delta {
				res = file.Size
				delta = file.Size - spaceToClean
			}
		}
		for _, items := range file.Children {
			if !items.isFile {
				queue = append(queue, items)
			}
		}

	}

	return res
}

func main() {

	pwd, _ := os.Getwd()
	inputData, err := os.ReadFile(pwd + "/2022/day07/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	//fmt.Println(Part1(inputData))
	fmt.Println(Part2(inputData))
}
