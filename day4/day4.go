package day4

import (
	"fmt"
	"path/filepath"
	"aoc2022/utils"
	"strings"
	"strconv"
)

func part1(multiLineInput []string) {
	var count = GetFullOverlapCount(multiLineInput)
	fmt.Println("\tPart 1: ", count);
}
// 494

func part2(multiLineInput []string) {
	var score = GetPartialOverlapCount(multiLineInput);
	fmt.Println("\tPart 2: ", score)
}
// 

func Run() {
	fmt.Println("--- Day2 ---")
	absPath, _ := filepath.Abs("./input/day4.txt")
	var multiLineInput = utils.ReadFile(absPath)
	part1(multiLineInput)
	part2(multiLineInput)
}

func createdArray(from int, to int) ([]int) {
	var arr = []int{};
	for i:= from; i <= to; i++ {
		arr = append(arr, i);
	}

	return arr;
}

func parseLine(line string) ([]int, []int) {
	var splitted = strings.Split(line, ",");

	var sectors1 = strings.Split(splitted[0], "-");
	begin1, err := strconv.Atoi(sectors1[0]);
	if err != nil {
    // ... handle error
    panic(err)
  }
  end1, err := strconv.Atoi(sectors1[1]);
	if err != nil {
    // ... handle error
    panic(err)
  }

	var sectorList1 = createdArray(begin1, end1)

	var sectors2 = strings.Split(splitted[1], "-");
	begin2, err := strconv.Atoi(sectors2[0]);
	if err != nil {
    // ... handle error
    panic(err)
  }
  end2, err := strconv.Atoi(sectors2[1]);
	if err != nil {
    // ... handle error
    panic(err)
  }
  var sectorList2 = createdArray(begin2, end2)
	return sectorList1, sectorList2
}

func contains(s []int, e int) bool {
    for _, a := range s {
        if a == e {
            return true
        }
    }
    return false
}


// PART 1
func IsFullOverlap(firstList []int, secondList []int) (bool) {
	var result = true;
	for _, sector := range firstList {
		if (!contains(secondList, sector)) {
			result = false;
		}
	}
	return result;
}

func GetFullOverlapCount(multiLineInput []string) (int) {
	count := 0
	for _, line := range multiLineInput {
		sectors1, sectors2 := parseLine(line)

		if (IsFullOverlap(sectors1, sectors2) || IsFullOverlap(sectors2, sectors1)) {
			count = count + 1;
		}
	}

	return count;
}

// PART 2
func IsPartialOverlap(firstList []int, secondList []int) (bool) {
	var result = false;
	for _, sector := range firstList {
		if (contains(secondList, sector)) {
			result = true;
		}
	}
	return result;
}

func GetPartialOverlapCount(multiLineInput []string) (int) {
	count := 0
	for _, line := range multiLineInput {
		sectors1, sectors2 := parseLine(line)

		if (IsPartialOverlap(sectors1, sectors2) || IsPartialOverlap(sectors2, sectors1)) {
			count = count + 1;
		}
	}

	return count;
}