package day1

import (
	"fmt"
	"path/filepath"
	"aoc2022/utils"
	"strconv"
	 "sort"
)

func part1(list []string) {
	var max = GetMax(list)
	fmt.Println("\tPart 1: ", max);
	// 70509
}

func part2(list []string) {
	var sum = GetLargestN(list, 3);
	fmt.Println("\tPart 2: ", sum)
	// 208567
}

func Run() {
	fmt.Println("--- Day1 ---")
	absPath, _ := filepath.Abs("./input/day1.txt")
	var input = utils.ReadFile(absPath)
	part1(input)
	part2(input)
}

var emptyString = "";

func GetMax(calList []string) int {
	var current = 0;
	var maxAmount = 0;
	for _, calorie := range calList {
		if (calorie == emptyString) {
			if (current > maxAmount) {
				maxAmount = current
			}
			current = 0;
		} else {
			parsedValue, err := strconv.Atoi(calorie);
			if err != nil {
        // ... handle error
        panic(err)
	    }
			current = current + parsedValue
		}
	}

	return maxAmount;
}

func GetLargestN(calList []string, n int) int {
	var current = 0;
	var sumList []int;

	for index, calorie := range calList {
		if (calorie == emptyString) {
			sumList = append(sumList, current);
			current = 0;
		} else {
			parsedValue, err := strconv.Atoi(calorie);
			if err != nil {
        // ... handle error
        panic(err)
	    }

			current = current + parsedValue
			if (index == (len(calList) - 1)) {
				sumList = append(sumList, current);
			}
		}
	}

	sort.Ints(sumList)

	i:=1;
	maxCalorieSum := 0
	for ;i <= n; i++ {
		maxCalorieSum = maxCalorieSum + sumList[len(sumList) - i];
	}

	return maxCalorieSum;
}
