package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type SeedInterval struct {
	Start int64
	End   int64
}

type Seed struct {
	Number       int64
	ToNumber     int64
	TargetNumber int64

	TargetInterval []SeedInterval
}

type Seeds struct {
	Seeds []Seed
}

type Converter struct {
	Start    int64
	MoveWith int64
	Len      int64
}

type Map struct {
	Converters []Converter
}

func parseInput(inputData []byte, version int) ([]Map, Seeds, error) {
	inputDataRows := strings.Split(string(inputData), "\n\n")

	result := []Map{}
	seedsRes := Seeds{}
	seedsRow := inputDataRows[0]
	seedsParts := strings.Split(seedsRow, " ")
	for s := 1; s < len(seedsParts); s++ {
		seed, err := strconv.ParseInt(seedsParts[s], 10, 64)
		if err != nil {
			return result, seedsRes, err
		}
		if version == 1 {
			seedsRes.Seeds = append(seedsRes.Seeds, Seed{Number: seed})
		} else {
			if s%2 == 1 {
				seedsRes.Seeds = append(seedsRes.Seeds, Seed{Number: seed})
			} else {
				seedsRes.Seeds[len(seedsRes.Seeds)-1].ToNumber = seedsRes.Seeds[len(seedsRes.Seeds)-1].Number + seed - 1
			}
			//if s%2 == 0 {
			//	prevSeed, _ := strconv.ParseInt(seedsParts[s-1], 10, 64)
			//	for k := prevSeed; k < prevSeed+seed; k++ {
			//		seedsRes.Seeds = append(seedsRes.Seeds, Seed{Number: k})
			//	}
			//}
		}

	}

	for i := 1; i < len(inputDataRows); i++ {
		dataRow := strings.Split(inputDataRows[i], "\n")
		currMap := Map{}
		for j := 1; j < len(dataRow); j++ {
			parts := strings.Split(dataRow[j], " ")
			if len(parts) != 3 {
				return result, seedsRes, fmt.Errorf("Error: parts wrong count")
			}
			end, err := strconv.ParseInt(parts[0], 10, 64)
			if err != nil {
				return result, seedsRes, err
			}
			start, err := strconv.ParseInt(parts[1], 10, 64)
			if err != nil {
				return result, seedsRes, err
			}
			lenConv, err := strconv.ParseInt(parts[2], 10, 64)
			if err != nil {
				return result, seedsRes, err
			}
			currConverter := Converter{Start: start, MoveWith: end - start, Len: lenConv}
			currMap.Converters = append(currMap.Converters, currConverter)
		}
		result = append(result, currMap)
	}

	return result, seedsRes, nil
}

func findSeedDestination(seeds Seeds, mapsConverter []Map) {

	for i := 0; i < len(seeds.Seeds); i++ {
		seedTarget := seeds.Seeds[i].Number
		for _, currMap := range mapsConverter {
			for _, currConv := range currMap.Converters {
				if seedTarget >= currConv.Start && seedTarget <= (currConv.Start+currConv.Len) {
					seedTarget += currConv.MoveWith
					break
				}
			}
		}
		seeds.Seeds[i].TargetNumber = seedTarget
	}
}

func findIntervalsFromConverter(currInterval SeedInterval, converter Converter) []SeedInterval {
	x := currInterval.Start
	y := currInterval.End
	a := converter.Start
	b := converter.Start + converter.Len - 1

	res := []SeedInterval{}
	if x > b || y < a {
		res = append(res, SeedInterval{Start: currInterval.Start, End: currInterval.End})
		return res
	}

	if x > a && x < b && y > b {
		res = append(res, SeedInterval{Start: currInterval.Start + converter.MoveWith, End: b + converter.MoveWith})
		res = append(res, SeedInterval{Start: b, End: y})
		return res
	}
	if x > a && x < b && y > a && y < b {
		res = append(res, SeedInterval{Start: currInterval.Start + converter.MoveWith, End: currInterval.End + converter.MoveWith})
		return res
	}
	if x < a && y > a && y < b {
		res = append(res, SeedInterval{Start: currInterval.Start, End: a})
		res = append(res, SeedInterval{Start: a + converter.MoveWith, End: y + converter.MoveWith})
		return res
	}
	if x < a {
		res = append(res, SeedInterval{Start: currInterval.Start, End: a})
		res = append(res, SeedInterval{Start: a + converter.MoveWith, End: b + converter.MoveWith})
		res = append(res, SeedInterval{Start: b, End: y})
		return res
	}

	return res

}

func findSeedDestination2(seeds Seeds, mapsConverter []Map) int64 {

	res := int64(math.MaxInt64)
	var currInterval SeedInterval
	for i := 0; i < len(seeds.Seeds); i++ {
		seeds.Seeds[i].TargetInterval = append(seeds.Seeds[i].TargetInterval, SeedInterval{Start: seeds.Seeds[i].Number, End: seeds.Seeds[i].ToNumber})
		for _, currMap := range mapsConverter {
			newIntervals := []SeedInterval{}
			for len(seeds.Seeds[i].TargetInterval) > 0 {
				currInterval, seeds.Seeds[i].TargetInterval = seeds.Seeds[i].TargetInterval[0], seeds.Seeds[i].TargetInterval[1:]
				for _, currConv := range currMap.Converters {
					newIntervals = append(newIntervals, findIntervalsFromConverter(currInterval, currConv)...)
				}
			}
			seeds.Seeds[i].TargetInterval = newIntervals
		}
	}

	for _, currSeed := range seeds.Seeds {
		for _, currInter := range currSeed.TargetInterval {
			if res > currInter.Start && currInter.Start > 0 {
				res = currInter.Start
			}
		}
	}

	return res
}

func Part1(inputData []byte) int64 {

	mapsConverters, seeds, err := parseInput(inputData, 1)
	if err != nil {
		fmt.Println(err)
		return 0
	}

	findSeedDestination(seeds, mapsConverters)

	//fmt.Println(seeds)
	result := seeds.Seeds[0].TargetNumber
	for _, seed := range seeds.Seeds {
		if result > seed.TargetNumber {
			result = seed.TargetNumber
		}
	}

	return result
}

// not finished
func Part2(inputData []byte) int64 {

	mapsConverters, seeds, err := parseInput(inputData, 2)
	if err != nil {
		fmt.Println(err)
		return 0
	}

	return findSeedDestination2(seeds, mapsConverters)

	findSeedDestination(seeds, mapsConverters)

	result := seeds.Seeds[0].TargetNumber
	for _, seed := range seeds.Seeds {
		if result > seed.TargetNumber {
			result = seed.TargetNumber
		}
	}

	return result - 1
}

func main() {

	pwd, _ := os.Getwd()
	inputData, err := os.ReadFile(pwd + "/2023/day05/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(Part1(inputData))
	//fmt.Println(Part2(inputData))
}
